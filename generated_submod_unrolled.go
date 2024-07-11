


package evmmax_arith

import (
	"math/bits"
)




func SubMod64(out, x, y, mod []uint64) {
    var c uint64 = 0
    var c1 uint64 = 0
	tmp := [1]uint64 { 0}

    for i := 0; i < 1; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 1; i++ {
        out[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

    var src []uint64
    // final sub was unnecessary, but do the copy anyways to make the impl constant time
    if c == 0 {
        src = tmp[:]
    } else {
        src = out[:]
    }

    copy(out[:], src[:])
}





func SubMod128(out, x, y, mod []uint64) {
    var c uint64 = 0
    var c1 uint64 = 0
	tmp := [2]uint64 { 0, 0}

    for i := 0; i < 2; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 2; i++ {
        out[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

    var src []uint64
    // final sub was unnecessary, but do the copy anyways to make the impl constant time
    if c == 0 {
        src = tmp[:]
    } else {
        src = out[:]
    }

    copy(out[:], src[:])
}





func SubMod192(out, x, y, mod []uint64) {
    var c uint64 = 0
    var c1 uint64 = 0
	tmp := [3]uint64 { 0, 0, 0}

    for i := 0; i < 3; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 3; i++ {
        out[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

    var src []uint64
    // final sub was unnecessary, but do the copy anyways to make the impl constant time
    if c == 0 {
        src = tmp[:]
    } else {
        src = out[:]
    }

    copy(out[:], src[:])
}





func SubMod256(out, x, y, mod []uint64) {
    var c uint64 = 0
    var c1 uint64 = 0
	tmp := [4]uint64 { 0, 0, 0, 0}

    for i := 0; i < 4; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 4; i++ {
        out[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

    var src []uint64
    // final sub was unnecessary, but do the copy anyways to make the impl constant time
    if c == 0 {
        src = tmp[:]
    } else {
        src = out[:]
    }

    copy(out[:], src[:])
}





func SubMod320(out, x, y, mod []uint64) {
    var c uint64 = 0
    var c1 uint64 = 0
	tmp := [5]uint64 { 0, 0, 0, 0, 0}

    for i := 0; i < 5; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 5; i++ {
        out[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

    var src []uint64
    // final sub was unnecessary, but do the copy anyways to make the impl constant time
    if c == 0 {
        src = tmp[:]
    } else {
        src = out[:]
    }

    copy(out[:], src[:])
}





func SubMod384(out, x, y, mod []uint64) {
    var c uint64 = 0
    var c1 uint64 = 0
	tmp := [6]uint64 { 0, 0, 0, 0, 0, 0}

    for i := 0; i < 6; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 6; i++ {
        out[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

    var src []uint64
    // final sub was unnecessary, but do the copy anyways to make the impl constant time
    if c == 0 {
        src = tmp[:]
    } else {
        src = out[:]
    }

    copy(out[:], src[:])
}





func SubMod448(out, x, y, mod []uint64) {
    var c uint64 = 0
    var c1 uint64 = 0
	tmp := [7]uint64 { 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < 7; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 7; i++ {
        out[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

    var src []uint64
    // final sub was unnecessary, but do the copy anyways to make the impl constant time
    if c == 0 {
        src = tmp[:]
    } else {
        src = out[:]
    }

    copy(out[:], src[:])
}





func SubMod512(out, x, y, mod []uint64) {
    var c uint64 = 0
    var c1 uint64 = 0
	tmp := [8]uint64 { 0, 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < 8; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 8; i++ {
        out[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

    var src []uint64
    // final sub was unnecessary, but do the copy anyways to make the impl constant time
    if c == 0 {
        src = tmp[:]
    } else {
        src = out[:]
    }

    copy(out[:], src[:])
}





func SubMod576(out, x, y, mod []uint64) {
    var c uint64 = 0
    var c1 uint64 = 0
	tmp := [9]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < 9; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 9; i++ {
        out[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

    var src []uint64
    // final sub was unnecessary, but do the copy anyways to make the impl constant time
    if c == 0 {
        src = tmp[:]
    } else {
        src = out[:]
    }

    copy(out[:], src[:])
}





func SubMod640(out, x, y, mod []uint64) {
    var c uint64 = 0
    var c1 uint64 = 0
	tmp := [10]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < 10; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 10; i++ {
        out[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

    var src []uint64
    // final sub was unnecessary, but do the copy anyways to make the impl constant time
    if c == 0 {
        src = tmp[:]
    } else {
        src = out[:]
    }

    copy(out[:], src[:])
}





func SubMod704(out, x, y, mod []uint64) {
    var c uint64 = 0
    var c1 uint64 = 0
	tmp := [11]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < 11; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 11; i++ {
        out[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

    var src []uint64
    // final sub was unnecessary, but do the copy anyways to make the impl constant time
    if c == 0 {
        src = tmp[:]
    } else {
        src = out[:]
    }

    copy(out[:], src[:])
}





func SubMod768(out, x, y, mod []uint64) {
    var c uint64 = 0
    var c1 uint64 = 0
	tmp := [12]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < 12; i++ {
        tmp[i], c = bits.Sub64(x[i], y[i], c)
    }

    for i := 0; i < 12; i++ {
        out[i], c1 = bits.Add64(tmp[i], mod[i], c1)
    }

    var src []uint64
    // final sub was unnecessary, but do the copy anyways to make the impl constant time
    if c == 0 {
        src = tmp[:]
    } else {
        src = out[:]
    }

    copy(out[:], src[:])
}
