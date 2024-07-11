


package evmmax_arith

import (
	"math/bits"
)





func AddMod64(out, x, y, mod []uint64) {
    _ = mod[0]
    _ = x[0]
    _ = y[0]
    _ = out[0]

    var c uint64 = 0
    var c1 uint64 = 0
	tmp := [1]uint64 { 0}

    for i := 0; i < 1; i++ {
        tmp[i], c = bits.Add64(x[i], y[i], c)
    }

    for i := 0; i < 1; i++ {
        out[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
    }

    var src []uint64
    // final sub was unnecessary, but do the copy anyways to make the impl constant time
    if c == 0 && c1 != 0 {
        src = tmp[:]
    } else {
        src = out[:]
    }

    copy(out[:], src[:])
}






func AddMod128(out, x, y, mod []uint64) {
    _ = mod[1]
    _ = x[1]
    _ = y[1]
    _ = out[1]

    var c uint64 = 0
    var c1 uint64 = 0
	tmp := [2]uint64 { 0, 0}

    for i := 0; i < 2; i++ {
        tmp[i], c = bits.Add64(x[i], y[i], c)
    }

    for i := 0; i < 2; i++ {
        out[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
    }

    var src []uint64
    // final sub was unnecessary, but do the copy anyways to make the impl constant time
    if c == 0 && c1 != 0 {
        src = tmp[:]
    } else {
        src = out[:]
    }

    copy(out[:], src[:])
}






func AddMod192(out, x, y, mod []uint64) {
    _ = mod[2]
    _ = x[2]
    _ = y[2]
    _ = out[2]

    var c uint64 = 0
    var c1 uint64 = 0
	tmp := [3]uint64 { 0, 0, 0}

    for i := 0; i < 3; i++ {
        tmp[i], c = bits.Add64(x[i], y[i], c)
    }

    for i := 0; i < 3; i++ {
        out[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
    }

    var src []uint64
    // final sub was unnecessary, but do the copy anyways to make the impl constant time
    if c == 0 && c1 != 0 {
        src = tmp[:]
    } else {
        src = out[:]
    }

    copy(out[:], src[:])
}






func AddMod256(out, x, y, mod []uint64) {
    _ = mod[3]
    _ = x[3]
    _ = y[3]
    _ = out[3]

    var c uint64 = 0
    var c1 uint64 = 0
	tmp := [4]uint64 { 0, 0, 0, 0}

    for i := 0; i < 4; i++ {
        tmp[i], c = bits.Add64(x[i], y[i], c)
    }

    for i := 0; i < 4; i++ {
        out[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
    }

    var src []uint64
    // final sub was unnecessary, but do the copy anyways to make the impl constant time
    if c == 0 && c1 != 0 {
        src = tmp[:]
    } else {
        src = out[:]
    }

    copy(out[:], src[:])
}






func AddMod320(out, x, y, mod []uint64) {
    _ = mod[4]
    _ = x[4]
    _ = y[4]
    _ = out[4]

    var c uint64 = 0
    var c1 uint64 = 0
	tmp := [5]uint64 { 0, 0, 0, 0, 0}

    for i := 0; i < 5; i++ {
        tmp[i], c = bits.Add64(x[i], y[i], c)
    }

    for i := 0; i < 5; i++ {
        out[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
    }

    var src []uint64
    // final sub was unnecessary, but do the copy anyways to make the impl constant time
    if c == 0 && c1 != 0 {
        src = tmp[:]
    } else {
        src = out[:]
    }

    copy(out[:], src[:])
}






func AddMod384(out, x, y, mod []uint64) {
    _ = mod[5]
    _ = x[5]
    _ = y[5]
    _ = out[5]

    var c uint64 = 0
    var c1 uint64 = 0
	tmp := [6]uint64 { 0, 0, 0, 0, 0, 0}

    for i := 0; i < 6; i++ {
        tmp[i], c = bits.Add64(x[i], y[i], c)
    }

    for i := 0; i < 6; i++ {
        out[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
    }

    var src []uint64
    // final sub was unnecessary, but do the copy anyways to make the impl constant time
    if c == 0 && c1 != 0 {
        src = tmp[:]
    } else {
        src = out[:]
    }

    copy(out[:], src[:])
}






func AddMod448(out, x, y, mod []uint64) {
    _ = mod[6]
    _ = x[6]
    _ = y[6]
    _ = out[6]

    var c uint64 = 0
    var c1 uint64 = 0
	tmp := [7]uint64 { 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < 7; i++ {
        tmp[i], c = bits.Add64(x[i], y[i], c)
    }

    for i := 0; i < 7; i++ {
        out[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
    }

    var src []uint64
    // final sub was unnecessary, but do the copy anyways to make the impl constant time
    if c == 0 && c1 != 0 {
        src = tmp[:]
    } else {
        src = out[:]
    }

    copy(out[:], src[:])
}






func AddMod512(out, x, y, mod []uint64) {
    _ = mod[7]
    _ = x[7]
    _ = y[7]
    _ = out[7]

    var c uint64 = 0
    var c1 uint64 = 0
	tmp := [8]uint64 { 0, 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < 8; i++ {
        tmp[i], c = bits.Add64(x[i], y[i], c)
    }

    for i := 0; i < 8; i++ {
        out[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
    }

    var src []uint64
    // final sub was unnecessary, but do the copy anyways to make the impl constant time
    if c == 0 && c1 != 0 {
        src = tmp[:]
    } else {
        src = out[:]
    }

    copy(out[:], src[:])
}






func AddMod576(out, x, y, mod []uint64) {
    _ = mod[8]
    _ = x[8]
    _ = y[8]
    _ = out[8]

    var c uint64 = 0
    var c1 uint64 = 0
	tmp := [9]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < 9; i++ {
        tmp[i], c = bits.Add64(x[i], y[i], c)
    }

    for i := 0; i < 9; i++ {
        out[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
    }

    var src []uint64
    // final sub was unnecessary, but do the copy anyways to make the impl constant time
    if c == 0 && c1 != 0 {
        src = tmp[:]
    } else {
        src = out[:]
    }

    copy(out[:], src[:])
}






func AddMod640(out, x, y, mod []uint64) {
    _ = mod[9]
    _ = x[9]
    _ = y[9]
    _ = out[9]

    var c uint64 = 0
    var c1 uint64 = 0
	tmp := [10]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < 10; i++ {
        tmp[i], c = bits.Add64(x[i], y[i], c)
    }

    for i := 0; i < 10; i++ {
        out[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
    }

    var src []uint64
    // final sub was unnecessary, but do the copy anyways to make the impl constant time
    if c == 0 && c1 != 0 {
        src = tmp[:]
    } else {
        src = out[:]
    }

    copy(out[:], src[:])
}






func AddMod704(out, x, y, mod []uint64) {
    _ = mod[10]
    _ = x[10]
    _ = y[10]
    _ = out[10]

    var c uint64 = 0
    var c1 uint64 = 0
	tmp := [11]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < 11; i++ {
        tmp[i], c = bits.Add64(x[i], y[i], c)
    }

    for i := 0; i < 11; i++ {
        out[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
    }

    var src []uint64
    // final sub was unnecessary, but do the copy anyways to make the impl constant time
    if c == 0 && c1 != 0 {
        src = tmp[:]
    } else {
        src = out[:]
    }

    copy(out[:], src[:])
}






func AddMod768(out, x, y, mod []uint64) {
    _ = mod[11]
    _ = x[11]
    _ = y[11]
    _ = out[11]

    var c uint64 = 0
    var c1 uint64 = 0
	tmp := [12]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < 12; i++ {
        tmp[i], c = bits.Add64(x[i], y[i], c)
    }

    for i := 0; i < 12; i++ {
        out[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
    }

    var src []uint64
    // final sub was unnecessary, but do the copy anyways to make the impl constant time
    if c == 0 && c1 != 0 {
        src = tmp[:]
    } else {
        src = out[:]
    }

    copy(out[:], src[:])
}
