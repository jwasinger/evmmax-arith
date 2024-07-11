package evmmax_arith

import (
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
)

func MaxModulus(limbCount int) []uint64 {
	mod := make([]uint64, limbCount, limbCount)

	for i := 0; i < limbCount; i++ {
		mod[i] = math.MaxUint64
	}
	return mod
}

func limbsToInt(limbs []uint64) *big.Int {
	numBytes := make([]byte, len(limbs)*8)
	for i, limb := range limbs {
		binary.BigEndian.PutUint64(numBytes[i:i+8], limb)
	}

	return new(big.Int).SetBytes(numBytes)
}

// convert a big-endian byte-slice to little-endian, ascending significance limbs
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

func LimbsToBytes(limbs []uint64) []byte {
	res := make([]byte, len(limbs)*8, len(limbs)*8)

	for i := 0; i < len(limbs); i++ {
		resOffset := (len(limbs) - (i + 1)) * 8
		binary.BigEndian.PutUint64(res[resOffset:resOffset+8], limbs[i])
	}
	return res
}

func PadBytes(val []byte, size uint64) []byte {
	pad_len := int(size) - len(val)
	if pad_len > 0 {
		padding := make([]byte, pad_len)
		return append(padding, val...)
	} else if pad_len < 0 {
		panic(fmt.Sprintf("pad_len < 0. val len = %d. pad len = %d\n", len(val), pad_len))
	} else {
		return val
	}
}
