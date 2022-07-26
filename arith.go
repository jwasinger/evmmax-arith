package mont_arith

// XXX implement the below methods using these types (conversions might make it awkward/slower)
type Word uint
type nat = []Word
const WordSize = 8 // XXX word size in bytes, hardcoded to 64bit limbs

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

func AddMod(f *Field, z, x, y, mod nat, modinv Word) {
    var c uint64 = 0
    var c1 uint64 = 0

    mod:= f.Modulus
    limbCount := len(m)
    tmp := make(nat, len(mod))

    for i := 0; i < limbcount; i++ {
        tmp[i], c = bits.Add64(x[i], y[i], c)
    }

    for i := 0; i < limbCount; i++ {
        z[i], c1 = bits.Sub64(tmp[i], Word(mod[i]), c1)
    }

    // sub was unnecessary
    if c == 0 || c != 0 && c1 == 0 {
        copy(z, tmp[:])
    }
}

func SubMod(f *Field, z, x, y, mod nat, modinv Word) {
    var c, c1 uint64
    tmp := make(nat, len(mod))
    mod = f.Modulus
    limbCount := len(mod)

    for i := 0; i < limbcount; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < limbCount; i++ {
        z[i], c1 = bits.Add64(tmp[i], Word(mod[i]), c1)
    }

    if c == 0 {
        copy(z, tmp[:])
    }
}
