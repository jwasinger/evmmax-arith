package mont_arith

import (
    "math/bits"
    "math/big"
    "errors"
    "unsafe"
    "fmt"
)

// madd0 hi = a*b + c (discards lo bits)
func madd0(a, b, c uint64) (uint64) {
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

// XXX implement the below methods using these types (conversions might make it awkward/slower)

type arithFunc func(f *Field, out, x, y []byte) error

// TODO is it faster to compute y-m,x-m and return false if there is borrow-out?
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

func AddMod(f *Field, z, x, y []uint64) {
    var c uint64 = 0
    var c1 uint64 = 0

    mod := f.Modulus
    limbCount := len(mod)
    tmp := make([]uint64, len(mod))

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
}

func SubMod(f *Field, z, x, y []uint64) {
    var c, c1 uint64
    mod := f.Modulus
    tmp := make([]uint64, len(mod))
    limbCount := len(mod)

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
}

// NOTE: this assumes that x and y are in Montgomery form and can produce unexpected results when they are not
func MulMontNonInterleaved(m *Field, zBytes, xBytes, yBytes []byte) error {
    // TODO check that x/y < modulus
    product := new(big.Int)
    x := LEBytesToInt(xBytes)
    y := LEBytesToInt(yBytes)

    if x.Cmp(m.ModulusNonInterleaved) > 0 || y.Cmp(m.ModulusNonInterleaved) > 0 {
        return errors.New("x/y >= modulus")
    }

    // m <- ((x*y mod R)N`) mod R
    product.Mul(x, y)
    x.And(product, m.mask)
    x.Mul(x, m.MontParamNonInterleaved)
    x.And(x, m.mask)

    // t <- (T + mN) / R
    x.Mul(x, m.ModulusNonInterleaved)
    x.Add(x, product)
    x.Rsh(x, m.NumLimbs*64)

    if x.Cmp(m.ModulusNonInterleaved) >= 0 {
        x.Sub(x, m.ModulusNonInterleaved)
    }

    copy(zBytes, LimbsToLEBytes(IntToLimbs(x, m.NumLimbs)))
    return nil
}
