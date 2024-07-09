package evmmax_arith

import (
	"encoding/binary"
	"fmt"
	"math"
)

func MaxModulus(limbCount int) []uint64 {
	mod := make([]uint64, limbCount, limbCount)

	for i := 0; i < limbCount; i++ {
		mod[i] = math.MaxUint64
	}
	return mod
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
