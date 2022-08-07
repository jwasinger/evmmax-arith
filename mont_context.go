package mont_arith

import (
    "fmt"
	"errors"
	"math/big"
)

const limbSize = 8

// TODO rename to FieldPreset?
type Field struct {
	// TODO make most of these private and the arith operations methods of this struct
	Modulus               []uint64
	ModulusNonInterleaved *big.Int // just here for convenience XXX better naming

	MontParamInterleaved    uint64
	MontParamNonInterleaved *big.Int

	NumLimbs uint

	r    *big.Int
	rInv *big.Int

	// mask for mod by R: 0xfff...fff - (1 << NumLimbs * 64) - 1
	mask *big.Int

    MulMont arithFunc
    AddMod arithFunc
    SubMod arithFunc

    preset ArithPreset
}

func (m *Field) RVal() *big.Int {
    r := big.NewInt(1)
    r.Lsh(r, limbSize * m.NumLimbs * 8)
    return r
}

func (m *Field) RInv() *big.Int {
    r := m.RVal()
    r.ModInverse(r, m.ModulusNonInterleaved)
    return r
}

func (m *Field) ToMont(val []uint64) []uint64 {
	dst_val := new(big.Int)
	src_val := LimbsToInt(val)
	dst_val.Mul(src_val, m.RVal())
	dst_val.Mod(dst_val, m.ModulusNonInterleaved)

    return IntToLimbs(dst_val, m.NumLimbs)
}

func (m *Field) ToNorm(val []uint64) []uint64 {
	dst_val := new(big.Int)
	src_val := LimbsToInt(val)
	dst_val.Mul(src_val, m.RInv())
	dst_val.Mod(dst_val, m.ModulusNonInterleaved)

	return IntToLimbs(dst_val, m.NumLimbs)
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

        preset,
	}

	return &result
}

func (m *Field) ModIsSet() bool {
	return m.NumLimbs != 0
}

func (m *Field) ValueSize() uint {
	return uint(len(m.Modulus))
}

func (m *Field) SetMod(mod []uint64) error {
	// XXX proper handling without hardcoding
	var limbCount uint = uint(len(mod))
	if limbCount == 0 || limbCount > 12 {
        fmt.Println("1")
		return errors.New("invalid modulus length")
	} else if mod[0] % 2 == 0 {
        fmt.Println(mod)
        fmt.Println("2")
        return errors.New("modulus cannot be even")
    }

    if mod[limbCount - 1] == 0 {
        fmt.Printf("modErr = %x\n", mod)
        return errors.New("modulus must occupy all limbs")
    }


	modInt := LimbsToInt(mod)
    negModInt := new(big.Int)
    negModInt.Neg(modInt)
    m.ModulusNonInterleaved = modInt

    modInv := new(big.Int)
    smallBase, _ := new(big.Int).SetString("18446744073709551616", 10)
    modInv.ModInverse(negModInt, smallBase)

    m.Modulus = make([]uint64, limbCount)
    copy(m.Modulus, mod[:])
	m.NumLimbs = limbCount

	m.MontParamInterleaved = modInv.Uint64()

    m.MulMont = m.preset.MulMontImpls[limbCount - 1]
    m.AddMod = m.preset.AddModImpls[limbCount - 1]
    m.SubMod = m.preset.SubModImpls[limbCount - 1]

	return nil
}
