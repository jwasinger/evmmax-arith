package mont_arith

import (
	"errors"
	"math/big"
)

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
	return m.r
}

func (m *Field) RInv() *big.Int {
	return m.rInv
}

func (m *Field) ToMont(val []uint64) []uint64 {
	dst_val := new(big.Int)
	src_val := LimbsToInt(val)
	dst_val.Mul(src_val, m.r)
	dst_val.Mod(dst_val, LimbsToInt(m.Modulus))

    return IntToLimbs(dst_val, m.NumLimbs)
}

func (m *Field) ToNorm(val []uint64) []uint64 {
	dst_val := new(big.Int)
	src_val := LimbsToInt(val)
	dst_val.Mul(src_val, m.rInv)
	dst_val.Mod(dst_val, LimbsToInt(m.Modulus))

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

func (m *Field) GTEMod(x, y []uint64) bool {
    for i := 0; i < int(m.NumLimbs); i++ {
        if x[i] > m.Modulus[i] || y[i] > m.Modulus[i]  {
            return true
        }
    }
    return false
}

func (m *Field) ModIsSet() bool {
	return m.NumLimbs != 0
}

func (m *Field) ValueSize() uint {
	return uint(len(m.Modulus))
}

func (m *Field) SetMod(mod []uint64) error {
	// XXX proper handling without hardcoding
	if len(mod) == 0 || len(mod) > 12 {
		return errors.New("invalid modulus length")
	} else if mod[0] % 2 == 0 {
        return errors.New("modulus cannot be even")
    }

	var limbCount uint = uint(len(mod))
	var limbSize uint = 8

	// r val chosen as max representable value for limbCount + 1: 0x1000...000
	rVal := new(big.Int)
	rVal.Lsh(big.NewInt(1), limbCount*limbSize*8)

	rValMask := new(big.Int)
	rValMask.Sub(rVal, big.NewInt(1))

	modInt := LimbsToInt(mod)
	montParamNonInterleaved := new(big.Int)
	montParamNonInterleaved = montParamNonInterleaved.Mul(modInt, big.NewInt(-1))
	montParamNonInterleaved.Mod(montParamNonInterleaved, rVal)

	if montParamNonInterleaved.ModInverse(montParamNonInterleaved, rVal) == nil {
		return errors.New("modinverse failed")
	}

	rInv := new(big.Int)
	if rInv.ModInverse(rVal, modInt) == nil {
		return errors.New("modinverse to compute rInv failed")
	}

	m.NumLimbs = limbCount
	m.r = rVal
	m.rInv = rInv
	m.mask = rValMask

	// mod % (1 << limb_count_bits)  == mod % (1 << limb_count_bytes * 8)
	m.ModulusNonInterleaved = modInt

	m.Modulus = IntToLimbs(modInt, m.NumLimbs)

	m.MontParamNonInterleaved = montParamNonInterleaved
	m.MontParamInterleaved = montParamNonInterleaved.Uint64()

    m.MulMont = m.preset.MulMontImpls[len(mod) - 1]
    m.AddMod = m.preset.AddModImpls[len(mod) - 1]
    m.SubMod = m.preset.SubModImpls[len(mod) - 1]

	return nil
}
