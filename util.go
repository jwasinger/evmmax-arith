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
	numBytes := limbsToBytes(limbs)
	return new(big.Int).SetBytes(numBytes)
}

func reverseEndianess(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

// convert a big-endian byte-slice to little-endian, ascending significance limbs
func bytesToLimbs(b []byte) []uint64 {
	wordCount := (len(b) + 7) / 8

	// pad the bytes to be a multiple of 64bits
	paddedBytes := make([]byte, wordCount*8)
	copy(paddedBytes[wordCount*8-len(b):], b[:])

	limbs := make([]uint64, wordCount)
	for i := 0; i < wordCount; i++ {
		limbs[i] = binary.BigEndian.Uint64(paddedBytes[i*8 : (i+1)*8])
	}
	// reverse to little-endian limb ordering
	for i, j := 0, len(limbs)-1; i < j; i, j = i+1, j-1 {
		limbs[i], limbs[j] = limbs[j], limbs[i]
	}
	return limbs
}

// convert limbs format to big-endian bytes
func limbsToBytes(limbs []uint64) []byte {
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
