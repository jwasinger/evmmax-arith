package mont_arith

import (
	"encoding/binary"
	"errors"
	"math/big"
    "fmt"
)

const limbSize = 8

// TODO rename to Context
type Field struct {
	// TODO make most of these private and the arith operations methods of this struct
	Modulus               []byte
	ModulusNonInterleaved *big.Int // just here for convenience XXX better naming
    ModulusLimbs          []uint64

	MontParamInterleaved    uint64
	MontParamNonInterleaved *big.Int

	NumLimbs uint

	r    *big.Int
	rInv *big.Int

	rSquared []byte

	// mask for mod by R: 0xfff...fff - (1 << NumLimbs * 64) - 1
	mask *big.Int

	MulMont     arithFunc
	AddMod      arithFunc
	SubMod      arithFunc
	MulMontCost uint64
	AddModCost  uint64
	SubModCost  uint64
	SetModCost  uint64

	ElementSize uint64

	preset ArithPreset
}

func (m *Field) RSquared() []byte {
	rSquared := make([]byte, m.NumLimbs * 8)
	copy(rSquared, m.rSquared)
	return rSquared
}

func (m *Field) RVal() *big.Int {
	r := big.NewInt(1)
	r.Lsh(r, limbSize*m.NumLimbs*8)
	return r
}

func (m *Field) RInv() *big.Int {
	r := m.RVal()
	r.ModInverse(r, m.ModulusNonInterleaved)
	return r
}

func (m *Field) ModInv() *big.Int {
	rVal := m.RVal()
	result := new(big.Int)
	result.Set(m.ModulusNonInterleaved)
	result.Neg(result)
	result.ModInverse(result, rVal)
	return result
}

// TODO this should not do allocation/copying.  should be just as fast as mulmont
func (m *Field) ToMont(val []byte) ([]byte, error) {
	// TODO ensure val is less than modulus
	out_bytes := make([]byte, m.NumLimbs*8)
	r_squared_bytes := m.RSquared()

	if err := m.MulMont(m, out_bytes, val, r_squared_bytes); err != nil {
		return nil, err
	}
	return out_bytes, nil
}

func (m *Field) ToNorm(val []byte) ([]byte, error) {
	// TODO ensure val is less than the modulus?
	out_bytes := make([]byte, m.NumLimbs*8)
	one := make([]byte, len(val))
	one[0] = 1

	if err := m.MulMont(m, out_bytes, val, one); err != nil {
		return nil, err
	}

	return out_bytes, nil
}

func NewField(preset ArithPreset) *Field {
	result := Field{
		nil,
		nil,
        nil,

		0,
		nil,

		0,
		nil,
		nil,
		nil,

		nil,

		nil,
		nil,
		nil,

		0,
		0,
		0,
		0,
		0,

		preset,
	}

	return &result
}

func (m *Field) ModIsSet() bool {
	return m.NumLimbs != 0
}

// compute montgomery parameters given big-endian modulus bytes
func (m *Field) SetMod(mod []byte) error {
	var limbCount uint = uint(len(mod))

    if mod[len(mod) - 1] % 2 == 0 {
		return errors.New("modulus cannot be even")
	}

    mod = PadBytes8(mod)
    fmt.Printf("mod is %+x (len=%d)\n", mod, len(mod))

    // TODO pad mod

	modInt := new(big.Int).SetBytes(mod)
	rSquared := big.NewInt(1)
	rSquared.Lsh(rSquared, 64*limbCount)
	rSquared.Mod(rSquared, modInt)
	rSquared.Mul(rSquared, rSquared)
	rSquared.Mod(rSquared, modInt)

	/*
		rSquared = rSquared.Mul(rVal, rVal)
		rSquared = rSquared.Mod(rSquared, modInt)
	*/

	m.rSquared = rSquared.Bytes()

	// want to compute r_val - (mod & (r_val - 1))
	littleRVal, _ := new(big.Int).SetString("18446744073709551616", 10)

    fmt.Printf("mod is %+x (len=%d)\n", mod, len(mod))
    fmt.Println(len(mod) - 8)
    fmt.Println(len(mod) - 1)
    mod_uint64 := binary.BigEndian.Uint64(mod[len(mod) - 8: len(mod)])
    fmt.Println("fin")

	negModInt := new(big.Int)
	negModInt.SetUint64(mod_uint64)
	negModInt.Sub(littleRVal, negModInt)
	modInv := new(big.Int)
	modInv.ModInverse(negModInt, littleRVal)

	m.MontParamInterleaved = modInv.Uint64()

	m.ModulusNonInterleaved = modInt
	rVal := big.NewInt(1)
	rVal.Lsh(rVal, 64*limbCount)
	negModInt = new(big.Int)
	negModInt.Neg(modInt)
	m.MontParamNonInterleaved = new(big.Int)
	m.MontParamNonInterleaved.ModInverse(negModInt, rVal)
	m.mask = big.NewInt(1)
	m.mask.Lsh(m.mask, 64*limbCount)
	m.mask.Sub(m.mask, big.NewInt(1))

    m.Modulus = mod
	m.NumLimbs = uint(len(m.Modulus))
	m.ElementSize = uint64(m.NumLimbs) * 8
    m.ModulusLimbs = BytesToLimbs(m.Modulus)

	var genericMulMontCutoff uint = 64
	if m.NumLimbs >= genericMulMontCutoff {
		m.MulMont = MulMontNonInterleaved
		m.AddMod = AddModGeneric
		m.SubMod = SubModGeneric
	} else {
		m.MulMont = m.preset.MulMontImpls[limbCount-1]

		// TODO fix (TODO ?? what was I thinking was wrong here?)
		m.AddMod = m.preset.AddModImpls[limbCount-1]
		m.SubMod = m.preset.SubModImpls[limbCount-1]
	}

	return nil
}
