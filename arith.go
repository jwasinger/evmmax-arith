package mont_arith

import (
    "math/bits"
)

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
func MulModMontNonInterleaved(f *Field, zLimbs, xLimbs, yLimbs []uint64) error {
    // length x == y assumed

    product := new(big.Int)
    x := LimbsToInt(xLimbs)
    y := LimbsToInt(yLimbs)

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

    copy(zLimbs, IntToLimbs(x, m.NumLimbs))
    return nil
}
