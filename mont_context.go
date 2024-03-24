package evmmax_arith

import (
	"errors"
	"fmt"
	"math"
	"math/big"
)

const limbSize = 8

type ModulusState struct {
	Modulus      []byte
	R2           []byte
	modInv       uint64
	scratchSpace []byte
	AddSubCost   uint64
	MulCost      uint64

	AddMod arithFunc
	SubMod arithFunc
	MulMod arithFunc

	one []byte
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

	one := make([]byte, paddedSize)
	one[paddedSize-1] = 1

	// TODO: represent scratch space as array of uints internally (?)
	m := ModulusState{
		Modulus:      modBytes,
		modInv:       modInv,
		R2:           r2Bytes,
		MulMod:       Preset[paddedSize/8-1],
		scratchSpace: make([]byte, paddedSize*scratchSize),
		one:          one,
	}
	return &m, nil
}
func (m *ModulusState) Store(dst, count int, from []byte) error {
	elemSize := len(m.Modulus)
	dstIdx := dst
	for srcIdx := 0; srcIdx < elemSize*count; srcIdx += elemSize {
		if !lte(from[srcIdx:srcIdx+elemSize], m.Modulus) {
			return errors.New("value must be less than modulus")
		}
		fmt.Println("converting to mont")
		fmt.Printf("modulus=%x\n", m.Modulus)
		fmt.Printf("modinv=%x\n", m.modInv)
		fmt.Printf("arg=%x\n", from[srcIdx:(srcIdx+1)*elemSize])
		fmt.Printf("r2=%x\n", m.R2)

		// convert to Montgomery form
		m.MulMod(m.modInv,
			m.Modulus,
			m.scratchSpace[dstIdx*elemSize:(dstIdx+1)*elemSize],
			from[srcIdx:(srcIdx+1)*elemSize],
			m.R2)
		fmt.Printf("result=%x\n", m.scratchSpace[dstIdx*elemSize:(dstIdx+1)*elemSize])
		dstIdx++
	}
	return nil
}

func (m *ModulusState) Load(dst []byte, from, count int) {
	dstIdx := 0
	elemSize := len(m.Modulus)
	for srcIdx := from * elemSize; srcIdx < (from+count)*elemSize; srcIdx += elemSize {
		// convert from Montgomery to canonical form
		m.MulMod(m.modInv, m.Modulus, dst[dstIdx:dstIdx+elemSize], m.scratchSpace[srcIdx:srcIdx+elemSize], m.one)
	}
}
