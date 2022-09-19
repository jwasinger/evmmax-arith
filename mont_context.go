package mont_arith

import (
	"errors"
	"math/big"
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

	MulMont arithFunc
	AddMod  arithFunc
	SubMod  arithFunc
    MulMontCost uint64
    AddModCost uint64
    SubModCost uint64
    SetModCost uint64

    ElementSize uint64

	preset ArithPreset
}

/*
// TODO: ToMont that does MulMont(val, rSquared)
func (m *Field) ToMont(result, val []uint64) {
    
}
*/

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

func (m *Field) ToMont(val []uint64) []uint64 {
	dst_val := new(big.Int)
	src_val := LimbsToInt(val)
	dst_val.Mul(src_val, m.RVal())
    // recompute ModulusNonInterleaved as the set preset may not have set it in SetMod
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
	if limbCount > m.preset.MaxLimbCount() {
		return errors.New("modulus limb count greater than max")
	}

	if mod[0]%2 == 0 {
		return errors.New("modulus cannot be even")
	}

	if mod[limbCount-1] == 0 {
		return errors.New("modulus must occupy all limbs")
	}

	modInt := LimbsToInt(mod)
    m.ModulusNonInterleaved = new(big.Int)
    m.ModulusNonInterleaved.Set(modInt)
	negModInt := new(big.Int)
	negModInt.Neg(modInt)

	rVal := big.NewInt(1)
	rVal.Lsh(rVal, 64*limbCount)

    rSquared := new(big.Int)
	rSquared = rSquared.Mul(rVal, rVal)
	rSquared = rSquared.Mod(rSquared, modInt)

	m.rSquared = IntToLimbs(rSquared, limbCount)

    // TODO place interleaved/non-interleaved mont parameters in their own unnamed structs
    if limbCount <= m.preset.mulMontCIOSCutoff {
        modInv := new(big.Int)
        smallBase, _ := new(big.Int).SetString("18446744073709551616", 10)
        modInv.ModInverse(negModInt, smallBase)

        m.MontParamInterleaved = modInv.Uint64()
    } else {
        m.MontParamNonInterleaved = new(big.Int)
        m.MontParamNonInterleaved.ModInverse(negModInt, rVal)
        m.mask = big.NewInt(1)
        m.mask.Lsh(m.mask, 64 * limbCount)
        m.mask.Sub(m.mask, big.NewInt(1))
    }

	m.Modulus = make([]uint64, limbCount)
	copy(m.Modulus, mod[:])
	m.NumLimbs = limbCount
    m.ElementSize = uint64(limbCount) * 4

	m.MulMont = m.preset.MulMontImpls[limbCount-1]
	m.AddMod = m.preset.AddModImpls[limbCount-1]
	m.SubMod = m.preset.SubModImpls[limbCount-1]

	return nil
}
