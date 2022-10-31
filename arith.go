package mont_arith

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"math/bits"
	"reflect"
	"unsafe"
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

func madd3(a, b, c, d, e uint64) (uint64, uint64) {
	var carry uint64
	var c_uint uint64
	hi, lo := bits.Mul64(a, b)
	c_uint, carry = bits.Add64(c, d, 0)
	hi, _ = bits.Add64(hi, 0, carry)
	lo, carry = bits.Add64(lo, c_uint, 0)
	hi, _ = bits.Add64(hi, e, carry)
	return hi, lo
}

/*
 * begin mulmont implementations
 */

func mulMont64(f *Field, outBytes, xBytes, yBytes []byte) error {
	var product [2]uint64
	var c uint64
	mod := f.Modulus
	modinv := f.MontParamInterleaved

	x := (*[1]uint64)(unsafe.Pointer(&xBytes[0]))[:]
	y := (*[1]uint64)(unsafe.Pointer(&yBytes[0]))[:]
	out := (*[1]uint64)(unsafe.Pointer(&outBytes[0]))[:]

	if x[0] >= mod[0] || y[0] >= mod[0] {
		return errors.New(fmt.Sprintf("x/y gte modulus"))
	}

	product[1], product[0] = bits.Mul64(x[0], y[0])
	m := product[0] * modinv
	c, _ = madd1(m, mod[0], product[0])
	out[0] = c + product[1]

	if out[0] > mod[0] {
		out[0] = c - mod[0]
	}
	return nil
}

type arithFunc func(f *Field, out, x, y []byte) error

// TODO: compute y-m,x-m and compute GTE from that (like the template version)
func GTE(x, y []uint64) bool {
	for i := len(x) - 1; i > 0; i-- {
		if x[i] < y[i] {
			return false
		}
	}

	if x[0] >= y[0] {
		return true
	}

	return false
}

func Eq(n, other []uint64) bool {
	if len(n) != len(other) {
		panic("unequal lengths")
	}

	for i := 0; i < len(n); i++ {
		if n[i] != other[i] {
			return false
		}
	}
	return true
}

func leBytesToLimbs(b []byte) []uint64 {
	if len(b)%8 != 0 {
		panic("length of b must be divisible by 8")
	}

	result := make([]uint64, len(b)/8)
	for i := 0; i < len(result); i++ {
		result[i] = binary.LittleEndian.Uint64(b[i*8 : (i+1)*8])
	}

	return result
}

// https://groups.google.com/g/golang-nuts/c/aPjvemV4F0U?pli=1
// touint64 assumes len(x)%8 == 0
func toUint64(x []byte) []uint64 {
	xx := make([]uint64, 0, 0)
	hdrp := (*reflect.SliceHeader)(unsafe.Pointer(&xx))
	hdrp.Data = (*reflect.SliceHeader)(unsafe.Pointer(&x)).Data
	hdrp.Len = len(x) / 8
	hdrp.Cap = len(x) / 8
	return xx
}

func AddModGeneric(f *Field, zBytes, xBytes, yBytes []byte) error {
	var c uint64 = 0
	var c1 uint64 = 0

	mod := f.Modulus
	limbCount := len(mod)
	tmp := make([]uint64, len(mod))
	x := toUint64(xBytes)
	y := toUint64(yBytes)
	z := toUint64(zBytes)

	if GTE(x, mod) || GTE(y, mod) {
		return errors.New("x/y was gte modulus")
	}

	for i := 0; i < limbCount; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < limbCount; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	// final sub was unnecessary
	if c == 0 && c1 != 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModGeneric(f *Field, zBytes, xBytes, yBytes []byte) error {
	var c uint64 = 0
	var c1 uint64 = 0

	mod := f.Modulus
	limbCount := len(mod)
	tmp := make([]uint64, len(mod))
	x := toUint64(xBytes)
	y := toUint64(yBytes)
	z := toUint64(zBytes)

	if GTE(x, mod) || GTE(y, mod) {
		return errors.New("x/y was gte modulus")
	}

	for i := 0; i < limbCount; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < limbCount; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	// final add was unecessary
	if c == 0 {
		copy(z, tmp[:])
	}

	for i := 0; i < limbCount; i++ {
		binary.LittleEndian.PutUint64(zBytes[i*8:(i+1)*8], z[i])
	}
	return nil
}

func MulMontNonInterleaved(m *Field, zBytes, xBytes, yBytes []byte) error {
	product := new(big.Int)
	t := new(big.Int)
	x := LEBytesToInt(xBytes)
	y := LEBytesToInt(yBytes)

	if x.Cmp(m.ModulusNonInterleaved) > 0 || y.Cmp(m.ModulusNonInterleaved) > 0 {
		return errors.New("x/y >= modulus")
	}

	// TODO: replace .And(mask) with bit-shifts

	// T <- x * y
	product.Mul(x, y)

	// m <- ((T mod R)N`) mod R (using the same variable for t and m)
	t.And(product, m.mask)
	t.Mul(t, m.MontParamNonInterleaved)
	t.And(t, m.mask)

	// t <- (T + mN) / R
	t.Mul(t, m.ModulusNonInterleaved)
	t.Add(t, product)
	t.Rsh(t, m.NumLimbs*64)

	if t.Cmp(m.ModulusNonInterleaved) >= 0 {
		t.Sub(t, m.ModulusNonInterleaved)
	}
	copy(zBytes, LimbsToLEBytes(IntToLimbs(t, m.NumLimbs)))

	return nil
}
