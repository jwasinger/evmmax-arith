package mont_arith

import (
    "math/bits"
)

// XXX implement the below methods using these types (conversions might make it awkward/slower)

type nat []uint64

func Eq(n, other nat) bool {
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

func AddMod(f *Field, z, x, y nat) {
    var c uint64 = 0

    mod := f.Modulus
    limbCount := len(mod)
    tmp := make(nat, len(mod))

    for i := 0; i < limbCount; i++ {
        tmp[i], c = bits.Add64(x[i], y[i], c)
    }

    c = 0
    for i := 0; i < limbCount; i++ {
        z[i], c = bits.Sub64(tmp[i], mod[i], c)
    }

    // sub was unnecessary
    if c != 0 {
        copy(z, tmp[:])
    }
}

func SubMod(f *Field, z, x, y nat) {
    var c, c1 uint64
    mod := f.Modulus
    tmp := make(nat, len(mod))
    limbCount := len(mod)

    for i := 0; i < limbCount; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < limbCount; i++ {
        z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

    if c == 0 {
        copy(z, tmp[:])
    }
}
