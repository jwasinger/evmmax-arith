package evmmax_arith

import (
	"encoding/binary"
	"math/big"
	"math/bits"
)

// madd0 hi = a*b + c (discards lo bits)
func madd0(a, b, c uint64) uint64 {
	var carry, lo uint64
	hi, lo := bits.Mul64(a, b)
	_, carry = bits.Add64(lo, c, 0)
	hi, _ = bits.Add64(hi, 0, carry)
	return hi
}

// madd1 hi, lo = a*b + c
func madd1(a, b, c uint64) (uint64, uint64) {
	var carry uint64
	hi, lo := bits.Mul64(a, b)
	lo, carry = bits.Add64(lo, c, 0)
	hi, _ = bits.Add64(hi, 0, carry)
	return hi, lo
}

// madd2 hi, lo = a*b + c + d
func madd2(a, b, c, d uint64) (uint64, uint64) {
	var carry uint64
	hi, lo := bits.Mul64(a, b)
	c, carry = bits.Add64(c, d, 0)
	hi, _ = bits.Add64(hi, 0, carry)
	lo, carry = bits.Add64(lo, c, 0)
	hi, _ = bits.Add64(hi, 0, carry)
	return hi, lo
}

type mulFunc func(out, x, y, mod []uint64, modInv uint64)
type addOrSubFunc func(out, x, y, mod []uint64)

func lt(x, y []uint64) bool {
	for i := len(x); i > 0; i-- {
		if x[i-1] < y[i-1] {
			return true
		}
	}
	return false
}

// place a big-endian byte slice to a descending-significance ordered array of limbs
func placeBEBytesInOutput(out []uint64, b []byte) {
	// pad the bytes to be the same size as the output
	padded := make([]byte, len(out)*8)
	copy(padded[len(out)*8-len(b):], b)

	// place the bytes into the output limbs
	resultLimbs := len(out)
	for i := 0; i < resultLimbs; i++ {
		out[(resultLimbs-1)-i] = binary.BigEndian.Uint64(padded[i*8 : (i+1)*8])
	}
}

func MulModBinary(z, x, y, modulus []uint64, modInv uint64) {
	result := new(big.Int)
	result = result.Mul(limbsToInt(x), limbsToInt(y))
	result = result.Mod(result, limbsToInt(modulus))
	placeBEBytesInOutput(z, result.Bytes())
}

func AddModBinary(z, x, y, modulus []uint64) {
	result := new(big.Int)
	result = result.Add(limbsToInt(x), limbsToInt(y))
	result = result.Mod(result, limbsToInt(modulus))
	placeBEBytesInOutput(z, result.Bytes())
}

func SubModBinary(z, x, y, modulus []uint64) {
	result := new(big.Int)
	result = result.Sub(limbsToInt(x), limbsToInt(y))
	result = result.Mod(result, limbsToInt(modulus))
	placeBEBytesInOutput(z, result.Bytes())
}
