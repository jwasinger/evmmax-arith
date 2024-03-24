package evmmax_arith

import (
	"encoding/binary"
	"errors"
	"math"
	"math/big"
)

const limbSize = 8

type ModulusState struct {
	Modulus []uint64
	R2      []uint64
	modInv  uint64

	// TODO: make this uint
	scratchSpace []uint64
	AddSubCost   uint64
	MulCost      uint64

	AddMod arithFunc
	SubMod arithFunc
	MulMod arithFunc

	one        []uint64
	modulusInt *big.Int
}

func bytesToLimbs(b []byte) []uint64 {
	limbs := make([]uint64, len(b)/8)
	for i := 0; i < len(b)/8; i++ {
		limbs[i] = binary.BigEndian.Uint64(b[i*8 : (i+1)*8])
	}
	// reverse to little-endian limb ordering
	for i, j := 0, len(limbs)-1; i < j; i, j = i+1, j-1 {
		limbs[i], limbs[j] = limbs[j], limbs[i]
	}
	return limbs
}

func NewModulusState(modBytes []byte, scratchSize int) (*ModulusState, error) {
	// TODO: will move validation into EVM
	if len(modBytes) >= 96 {
		return nil, errors.New("modulus cannot be greater than 768 bits")
	}
	if modBytes[len(modBytes)-1]%2 == 0 {
		return nil, errors.New("modulus cannot be even")
	}
	if modBytes[0] == 0 {
		return nil, errors.New("modulus must be entirely occupied")
	}
	if scratchSize > 256 {
		return nil, errors.New("scratch space can be 256-sized max")
	}
	mod := new(big.Int).SetBytes(modBytes)
	paddedSize := int(math.Ceil(float64(len(modBytes))/8.0)) * 8
	modInv := new(big.Int).ModInverse(big.NewInt(-int64(mod.Uint64())), new(big.Int).Lsh(big.NewInt(1), 64)).Uint64()

	r2 := new(big.Int).Lsh(big.NewInt(1), uint(paddedSize)*8*2)
	r2.Mod(r2, mod)

	r2Bytes := r2.Bytes()
	if len(modBytes) < paddedSize {
		modBytes = append(modBytes, make([]byte, paddedSize-len(modBytes))...)
	}
	if len(r2Bytes) < paddedSize {
		r2Bytes = append(r2Bytes, make([]byte, paddedSize-len(r2Bytes))...)
	}

	one := make([]uint64, paddedSize/8)
	one[len(one)-1] = 1

	m := ModulusState{
		Modulus:      bytesToLimbs(modBytes),
		modInv:       modInv,
		R2:           bytesToLimbs(r2Bytes),
		MulMod:       Preset[paddedSize/8-1],
		scratchSpace: make([]uint64, (paddedSize/8)*scratchSize),
		one:          one,
		modulusInt:   mod,
	}
	return &m, nil
}
func (m *ModulusState) Store(dst, count int, from []byte) error {
	elemSize := len(m.Modulus)
	dstIdx := dst * elemSize
	for srcIdx := 0; srcIdx < elemSize*8*count; srcIdx += elemSize * 8 {
		// convert the big-endian bytes to little-endian limbs, descending-significance ordered
		val := bytesToLimbs(from[srcIdx : srcIdx+elemSize*8])
		if !lt(val, m.Modulus) {
			return errors.New("value must be less than modulus")
		}
		// convert to Montgomery form
		m.MulMod(m.modInv,
			m.Modulus,
			m.scratchSpace[dstIdx:dstIdx+elemSize],
			val,
			m.R2)
		dstIdx++
	}
	return nil
}

func (m *ModulusState) Load(dst []byte, from, count int) {
	elemSize := len(m.Modulus)
	var dstIdx int
	for srcIdx := from; srcIdx < from+count; srcIdx++ {
		res := make([]uint64, elemSize)
		// convert from Montgomery to canonical form
		m.MulMod(m.modInv, m.Modulus, res, m.scratchSpace[srcIdx:srcIdx+elemSize], m.one)
		// swap each limb to big endian (the result in dst is a big-endian number)
		for i := 0; i < elemSize; i++ {
			binary.BigEndian.PutUint64(dst[dstIdx+i*8:dstIdx+(i+1)*8], res[i])
		}
		dstIdx += elemSize * 8
	}
}
