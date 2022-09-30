package mont_arith

import (
	"errors"
	"math/big"
    "encoding/binary"
)

const limbSize = 8

// TODO rename to Context
type Field struct {
	// TODO make most of these private and the arith operations methods of this struct
	Modulus               []uint64
	ModulusNonInterleaved *big.Int // just here for convenience XXX better naming

	MontParamInterleaved    uint64
	MontParamNonInterleaved *big.Int

	NumLimbs uint

	r    *big.Int
	rInv *big.Int

	rSquared []uint64

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

func (m *Field) RSquared() []uint64 {
	rSquared := make([]uint64, m.NumLimbs)
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

func uint64_array_to_le_bytes(val []uint64) []byte {
	res := make([]byte, len(val) * 8)
	for i := 0; i < len(val); i++ {
        binary.LittleEndian.PutUint64(res[i*8:(i+1)*8], val[i])
	}

	return res
}

func le_bytes_to_uint64_array(val []byte) []uint64 {
    res := make([]uint64, len(val) / 8)
	for i := 0; i < len(val) / 8; i++ {
        res[i] = binary.LittleEndian.Uint64(val[i*8:(i+1)*8])
    }
    return res
}

// TODO this should not do allocation/copying.  should be just as fast as mulmont
func (m *Field) ToMont(val []uint64) ([]uint64, error) {
    // TODO ensure val is less than modulus
    out_bytes := make([]byte, m.NumLimbs * 8)
	input_bytes := uint64_array_to_le_bytes(val)
	r_squared_bytes := uint64_array_to_le_bytes(m.RSquared())

	if err := m.MulMont(m, out_bytes, input_bytes, r_squared_bytes); err != nil {
		return nil, err
	}
    return le_bytes_to_uint64_array(out_bytes), nil
}

func (m *Field) ToNorm(val []uint64) ([]uint64, error) {
    // TODO ensure val is less than the modulus?
    out_bytes := make([]byte, m.NumLimbs * 8)
	input_bytes := uint64_array_to_le_bytes(val)
    one := make([]byte, len(val))
    one[0] = 1

	if err := m.MulMont(m, out_bytes, input_bytes, one); err != nil {
		return nil, err
	}

    return le_bytes_to_uint64_array(out_bytes), nil
}

func NewField(preset ArithPreset) *Field {
	result := Field{
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

const karatsubaThreshold = 64

func (m *Field) SetMod(mod []uint64) error {
	var limbCount uint = uint(len(mod))

	if mod[0]%2 == 0 {
		return errors.New("modulus cannot be even")
	}

	if mod[limbCount-1] == 0 {
		return errors.New("modulus must occupy all limbs")
	}

	modInt := LimbsToInt(mod)
	rSquared := big.NewInt(1)
	rSquared.Lsh(rSquared, 64*limbCount)
	rSquared.Mod(rSquared, modInt)
	rSquared.Mul(rSquared, rSquared)
	rSquared.Mod(rSquared, modInt)

	/*
		rSquared = rSquared.Mul(rVal, rVal)
		rSquared = rSquared.Mod(rSquared, modInt)
	*/

	m.rSquared = IntToLimbs(rSquared, limbCount)

	// TODO place interleaved/non-interleaved mont parameters in their own unnamed structs
	if limbCount < m.preset.mulMontCIOSCutoff {
		// want to compute r_val - (mod & (r_val - 1))
		littleRVal, _ := new(big.Int).SetString("18446744073709551616", 10)

		negModInt := new(big.Int)
		negModInt.SetUint64(mod[0])
		negModInt.Sub(littleRVal, negModInt)
		modInv := new(big.Int)
		modInv.ModInverse(negModInt, littleRVal)

		m.MontParamInterleaved = modInv.Uint64()
	} else {
		m.ModulusNonInterleaved = modInt
		rVal := big.NewInt(1)
		rVal.Lsh(rVal, 64*limbCount)
		negModInt := new(big.Int)
		negModInt.Neg(modInt)
		m.MontParamNonInterleaved = new(big.Int)
		m.MontParamNonInterleaved.ModInverse(negModInt, rVal)
		m.mask = big.NewInt(1)
		m.mask.Lsh(m.mask, 64*limbCount)
		m.mask.Sub(m.mask, big.NewInt(1))
	}

	m.Modulus = make([]uint64, limbCount)
	copy(m.Modulus, mod[:])
	m.NumLimbs = limbCount
	m.ElementSize = uint64(limbCount) * 4

    var genericMulMontCutoff uint = 64
    if limbCount >= genericMulMontCutoff {
        m.MulMont = MulMontNonInterleaved
        m.AddMod = AddModGeneric
        m.SubMod = SubModGeneric
    } else {
        m.MulMont = m.preset.MulMontImpls[limbCount-1]

        // TODO fix
        m.AddMod = m.preset.AddModImpls[limbCount-1]
        m.SubMod = m.preset.SubModImpls[limbCount-1]
    }

	return nil
}
