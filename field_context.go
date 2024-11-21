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

// FieldContext represents a modulus, an allocated space of reduced field
// elements, and any internal state necessary to perform efficient modular
// addition subtraction and multiplication on elements within the space,
// load/store them to/from the space.
type FieldContext struct {
	Modulus []uint64
	R2      []uint64
	modInv  uint64

	useMontgomeryRepr bool // true if values are represented in montgomery form internally
	isModulusBinary   bool

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

// returns true if the modulus is a power of two
func isModulusBinary(modulus *big.Int) bool {
	if modulus.Bit(modulus.BitLen()-1) == 1 && new(big.Int).SetBit(modulus, modulus.BitLen()-1, 0).Cmp(big.NewInt(0)) == 0 {
		return true
	}
	return false
}

// NewFieldContext instantiates a field context with a given big-endian modulus, number of field elements
func NewFieldContext(modBytes []byte, scratchSize int) (*FieldContext, error) {
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
			isModulusBinary:       true,
		}, nil
	}
	if modBytes[len(modBytes)-1]%2 == 0 {
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

	return &m, nil
}

// IsModulusBinary returns whether the modulus is a power of two
func (f *FieldContext) IsModulusBinary() bool {
	return f.isModulusBinary
}

// NumElems returns the number of field elements allocated in this context
func (f *FieldContext) NumElems() uint {
	return f.scratchSpaceElemCount
}

// AllocedSize returns the size of the field elements in bytes, where each
// field element's size is the size of the modulus padded to be a multiple
// of 64 bits.
func (f *FieldContext) AllocedSize() uint {
	return uint(len(f.scratchSpace) * 8)
}

// ElemSize returns the size of field elements: the size of the modulus padded
// to the nearest multiple of 64 bits.
func (f *FieldContext) ElemSize() uint {
	return f.elemSize
}

// compute -mod ** -1 % 1 << 64 .  computed via hensel lifting
// as per the go standard library (TODO: link to the code and paper)
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

// MulMod computes 'count' modular multiplications, pairwise multiplying values
// from offsets [x, x+xStride, x+xStride*2, ..., x+xStride*(count - 1)]
// and [y, y+yStride, y+yStride*2, ..., y+yStride*(count - 1)]
// placing the result in [out, out+outStride, out+outStride*2, ..., out+outStride*(count - 1)].
//
// inputs/outputs can overlap without affecting the result.  it is not validated
// that inputs are within bounds.
func (m *FieldContext) MulMod(out, outStride, x, xStride, y, yStride, count uint) {
	elemSize := uint(len(m.Modulus))

	// perform the multiplications
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
	// copy the result from the intermediate scratch buffer back into the context's field element space
	for i := out; i < out+(count*outStride); i += outStride {
		offset := i * elemSize
		copy(m.scratchSpace[offset:offset+elemSize], outputWriteBuf[offset:offset+elemSize])
	}
}

// SubMod computes 'count' modular subtractions, pairwise subtracting values
// at offsets [x, x+xStride, x+xStride*2, ..., x+xStride*(count - 1)]
// and [y, y+yStride, y+yStride*2, ..., y+yStride*(count - 1)]
// placing the result in [out, out+outStride, out+outStride*2, ..., out+outStride*(count - 1)].
//
// inputs/outputs can overlap without affecting the result.  it is not validated
// that inputs are within bounds.
func (m *FieldContext) SubMod(out, outStride, x, xStride, y, yStride, count uint) {
	elemSize := uint(len(m.Modulus))

	// perform the subtractions
	for i := uint(0); i < count; i++ {
		xSrc := (x + i*xStride) * elemSize
		ySrc := (y + i*yStride) * elemSize
		dst := (out + i*outStride) * elemSize
		m.subMod(outputWriteBuf[dst:dst+elemSize],
			m.scratchSpace[xSrc:xSrc+elemSize],
			m.scratchSpace[ySrc:ySrc+elemSize],
			m.Modulus)
	}
	// copy the results from the intermediate scratch buffer back into the context's field element space
	for i := out; i < out+(count*outStride); i += outStride {
		offset := i * elemSize
		copy(m.scratchSpace[offset:offset+elemSize], outputWriteBuf[offset:offset+elemSize])
	}
}

// AddMod computes 'count' modular additions, pairwise adding values
// at offsets [x, x+xStride, x+xStride*2, ..., x+xStride*(count - 1)]
// and [y, y+yStride, y+yStride*2, ..., y+yStride*(count - 1)]
// placing the result in [out, out+outStride, out+outStride*2, ..., out+outStride*(count - 1)].
//
// inputs/outputs can overlap without affecting the result.  it is not validated
// that inputs are within bounds.
func (m *FieldContext) AddMod(out, outStride, x, xStride, y, yStride, count uint) {
	elemSize := uint(len(m.Modulus))

	// perform the additions
	for i := uint(0); i < count; i++ {
		xSrc := (x + i*xStride) * elemSize
		ySrc := (y + i*yStride) * elemSize
		dst := (out + i*outStride) * elemSize
		m.addMod(outputWriteBuf[dst:dst+elemSize],
			m.scratchSpace[xSrc:xSrc+elemSize],
			m.scratchSpace[ySrc:ySrc+elemSize],
			m.Modulus)
	}
	// copy the results from the intermediate scratch buffer back into the context's field element space
	for i := out; i < out+(count*outStride); i += outStride {
		offset := i * elemSize
		copy(m.scratchSpace[offset:offset+elemSize], outputWriteBuf[offset:offset+elemSize])
	}
}

// Store takes a byte slice representing 'count' field elements, each of which
// is sized to the modulus length padded to the nearest 64 bits.  It places them
// in the allocated field element space starting at offset dst.
//
// does not perform bounds checks on the inputs.  Checks that each field element in 'from'
// is reduced by the modulus.
func (m *FieldContext) Store(dst, count uint, from []byte) error {
	elemSize := uint(len(m.Modulus))

	for i := uint(0); i < count; i++ {
		srcIdx := i * elemSize * 8
		dstIdx := dst*elemSize + i*elemSize

		// swap big-endian bytes to ascending-significance-ordered little-endian limbs internal repr
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

// Load loads 'count' number of field elements starting at from, and placing
// them into dst.
//
// does not perform any validity checks on the inputs.
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
		// swap to descending-significance (big-endian) limb ordering
		for i := 0; i < elemSize; i++ {
			binary.BigEndian.PutUint64(dst[dstIdx+i*8:dstIdx+(i+1)*8], res[len(res)-(i+1)])
		}
		dstIdx += elemSize * 8
	}
}
