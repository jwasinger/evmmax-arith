package evmmax_arith

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"math/big"
	"math/bits"
)

const maxModulusSize = 96 // 768 bits maximum modulus width

// 256 * 768 bit buffer for writing values out before mutating register space
var outputWriteBuf [3072]uint64

type FieldContext struct {
	Modulus           []uint64
	R2                []uint64
	modInv            uint64
	useMontgomeryRepr bool // true if values are represented in montgomery form internally

	scratchSpace []uint64
	AddSubCost   uint64
	MulCost      uint64

	addMod addOrSubFunc
	subMod addOrSubFunc
	mulMod mulFunc

	one                   []uint64
	modulusInt            *big.Int
	elemSize              uint
	scratchSpaceElemCount uint
}

const (
	FallBackOnly = iota
	MulModAsm    = iota
	AllAsm       = iota
)

func isModulusBinary(modulus *big.Int) bool {
	if modulus.Bit(modulus.BitLen()-1) == 1 && new(big.Int).SetBit(modulus, modulus.BitLen()-1, 0).Cmp(big.NewInt(0)) == 0 {
		return true
	}
	return false
}

// NewFieldContext instantiates a field context with a given big-endian modulus, number of field elements
func NewFieldContext(modBytes []byte, scratchSize int, _ int) (*FieldContext, error) {
	if len(modBytes) > maxModulusSize {
		return nil, errors.New("modulus cannot be greater than 768 bits")
	}
	if len(modBytes) == 0 {
		return nil, errors.New("modulus must be non-empty")
	}
	if modBytes[0] == 0 {
		return nil, errors.New("most significant byte of modulus must not be zero")
	}
	if scratchSize == 0 {
		return nil, errors.New("scratch space must have non-zero size")
	}
	if scratchSize > 256 {
		return nil, errors.New("scratch space can allocate a maximum of 256 field elements")
	}

	mod := new(big.Int).SetBytes(modBytes)
	paddedSize := int(math.Ceil(float64(len(modBytes))/8.0)) * 8
	if isModulusBinary(mod) {
		return &FieldContext{
			Modulus:               bytesToLimbs(modBytes),
			mulMod:                MulModBinary,
			addMod:                AddModBinary,
			subMod:                SubModBinary,
			scratchSpace:          make([]uint64, (paddedSize/8)*scratchSize),
			scratchSpaceElemCount: uint(scratchSize),
			modulusInt:            mod,
			elemSize:              uint(paddedSize),
			useMontgomeryRepr:     false,
		}, nil
	}
	if modBytes[len(modBytes)-1]%2 == 0 {
		fmt.Println(mod.Text(2))
		return nil, errors.New("modulus cannot be even")
	}
	modInv := negModInverse(mod.Uint64())

	r2 := new(big.Int).Lsh(big.NewInt(1), uint(paddedSize)*8*2)
	r2.Mod(r2, mod)

	r2Bytes := r2.Bytes()
	if len(modBytes) < paddedSize {
		modBytes = append(make([]byte, paddedSize-len(modBytes)), modBytes...)
	}
	if len(r2Bytes) < paddedSize {
		r2Bytes = append(make([]byte, paddedSize-len(r2Bytes)), r2Bytes...)
	}

	one := make([]uint64, paddedSize/8)
	one[0] = 1

	m := FieldContext{
		Modulus:               bytesToLimbs(modBytes),
		modInv:                modInv,
		R2:                    bytesToLimbs(r2Bytes),
		mulMod:                mulmodPreset[paddedSize/8-1],
		addMod:                addmodPreset[paddedSize/8-1],
		subMod:                submodPreset[paddedSize/8-1],
		scratchSpace:          make([]uint64, (paddedSize/8)*scratchSize),
		scratchSpaceElemCount: uint(scratchSize),
		one:                   one,
		modulusInt:            mod,
		elemSize:              uint(paddedSize),
		useMontgomeryRepr:     true,
	}

	/*
		switch preset384bit {
		case FallBackOnly:
			break
		case MulModAsm:
			if paddedSize/8 == 6 {
				m.mulMod = MulMont384_asm
			}
		case AllAsm:
			if paddedSize/8 == 6 {
				m.mulMod = MulMont384_asm
				m.addMod = AddMod384_asm
				m.subMod = SubMod384_asm
			}
		default:
			panic("invalid parameter for 384-bit preset")
		}
	*/

	return &m, nil
}

func (f *FieldContext) NumElems() uint {
	return f.scratchSpaceElemCount
}

func (f *FieldContext) AllocedSize() uint {
	return uint(len(f.scratchSpace) * 8)
}

// elem size in bytes
func (f *FieldContext) ElemSize() uint {
	return f.elemSize
}

// compute -mod ** -1 % 1 << 64 .
// from (paper), used in go-stdlib
func negModInverse(mod uint64) uint64 {
	k0 := 2 - mod
	t := mod - 1
	for i := 1; i < bits.UintSize; i <<= 1 {
		t *= t
		k0 *= (t + 1)
	}
	k0 = -k0
	return k0
}

// Note: manually inlining the arith funcs here into the opcode handler seems to give overall ~6-7% performance increase on g2 mul
// benchmark
func (m *FieldContext) MulMod(out, outStride, x, xStride, y, yStride, count uint) {
	elemSize := uint(len(m.Modulus))

	for i := uint(0); i < count; i++ {
		xSrc := (x + i*xStride) * elemSize
		ySrc := (y + i*yStride) * elemSize
		dst := (out + i*outStride) * elemSize
		m.mulMod(outputWriteBuf[dst:dst+elemSize],
			m.scratchSpace[xSrc:xSrc+elemSize],
			m.scratchSpace[ySrc:ySrc+elemSize],
			m.Modulus,
			m.modInv)
	}
	copy(m.scratchSpace[out:out+elemSize*count], outputWriteBuf[out:out+elemSize*count])
}

func (m *FieldContext) SubMod(out, outStride, x, xStride, y, yStride, count uint) {
	elemSize := uint(len(m.Modulus))

	for i := uint(0); i < count; i++ {
		xSrc := (x + i*xStride) * elemSize
		ySrc := (y + i*yStride) * elemSize
		dst := (out + i*outStride) * elemSize
		m.subMod(outputWriteBuf[dst:dst+elemSize],
			m.scratchSpace[xSrc:xSrc+elemSize],
			m.scratchSpace[ySrc:ySrc+elemSize],
			m.Modulus)
	}
	copy(m.scratchSpace[out:out+elemSize*count], outputWriteBuf[out:out+elemSize*count])
}

func (m *FieldContext) AddMod(out, outStride, x, xStride, y, yStride, count uint) {
	elemSize := uint(len(m.Modulus))

	for i := uint(0); i < count; i++ {
		xSrc := (x + i*xStride) * elemSize
		ySrc := (y + i*yStride) * elemSize
		dst := (out + i*outStride) * elemSize
		m.addMod(outputWriteBuf[dst:dst+elemSize],
			m.scratchSpace[xSrc:xSrc+elemSize],
			m.scratchSpace[ySrc:ySrc+elemSize],
			m.Modulus)
	}
	copy(m.scratchSpace[out:out+elemSize*count], outputWriteBuf[out:out+elemSize*count])
}

func (m *FieldContext) Store(dst, count uint, from []byte) error {
	elemSize := uint(len(m.Modulus))
	dstIdx := dst * elemSize

	if dstIdx+count > m.scratchSpaceElemCount {
		return errors.New("out of bounds field element store")
	}
	for i := uint(0); i < count; i++ {
		srcIdx := i * elemSize * 8
		dstIdx := dst*elemSize + i*elemSize

		// convert the big-endian bytes to little-endian limbs, descending-significance ordered
		val := bytesToLimbs(from[srcIdx : srcIdx+elemSize*8])
		if !lt(val, m.Modulus) {
			return fmt.Errorf("value (%+v) must be less than modulus (%+v)", val, m.Modulus)
		}

		if m.useMontgomeryRepr {
			// convert to Montgomery form
			m.mulMod(m.scratchSpace[dstIdx:dstIdx+elemSize],
				val,
				m.R2,
				m.Modulus,
				m.modInv)
		} else {
			copy(m.scratchSpace[dstIdx:dstIdx+elemSize], val[:])
		}
		dstIdx++
	}
	return nil
}

func (m *FieldContext) Load(dst []byte, from, count int) {
	elemSize := len(m.Modulus)
	var dstIdx int
	for srcIdx := from; srcIdx < from+count; srcIdx++ {
		res := make([]uint64, elemSize)
		if m.useMontgomeryRepr {
			// convert from Montgomery to canonical form
			m.mulMod(res, m.scratchSpace[srcIdx*elemSize:(srcIdx+1)*elemSize], m.one, m.Modulus, m.modInv)
		} else {
			copy(res[:], m.scratchSpace[srcIdx*elemSize:(srcIdx+1)*elemSize])
		}
		// swap each limb to big endian (the result in dst is a big-endian number)
		for i := 0; i < elemSize; i++ {
			binary.BigEndian.PutUint64(dst[dstIdx+i*8:dstIdx+(i+1)*8], res[len(res)-(i+1)])
		}
		dstIdx += elemSize * 8
	}
}
