package mont_arith

import (
	"encoding/binary"
	"math/big"
)

/*
Methods for converting between Go big-int (64bit little-endian limbs, big-endian limb ordering)
and the bigint representation expected here: 64bit little-endian limbs, little-endian ordered
*/
func LEBytesToInt(v []byte) *big.Int {
	result := new(big.Int)

	if len(v)%8 != 0 {
		panic("invalid val length for modext bytes")
	}

	val := make([]byte, len(v), len(v))
	copy(val, v)

	// byteswap 8 bytes at a time
	for i := 0; i < len(val)/2; i++ {
		val[i], val[len(val)-1-i] = val[len(val)-1-i], val[i]
	}

	result.SetBytes(val)
	return result
}

func LimbsToLEBytes(val []uint64) []byte {
	result := make([]byte, len(val)*8)

	for i := 0; i < len(val); i++ {
		startIdx := i * 8

		for j := 0; j < 8; j++ {
			result[startIdx+j] = byte(val[i] >> (j * 8))
		}
	}

	return result
}

// convert big.Int (big-endian) to little-endian limbs
func IntToLimbs(val *big.Int, num_limbs uint) []uint64 {
	val_bytes := val.Bytes()

	// pad length to be a multiple of 64bits
	if len(val_bytes) < 8*int(num_limbs) {
		pad_len := 8*int(num_limbs) - len(val_bytes)
		pad := make([]byte, pad_len, pad_len)
		val_bytes = append(pad, val_bytes...)
	} else if len(val_bytes) > 8*int(num_limbs) {
		panic("val too big to fit in specified number of limbs")
	}

	result := make([]uint64, len(val_bytes)/8, len(val_bytes)/8)

	// place byteswapped (little-endian) val into result
	for i := 0; i < len(result); i++ {
		startIdx := (len(result) - (i + 1)) * 8
		endIdx := (len(result) - i) * 8

		// TODO: this assumes that the system is little-endian.  is that okay?
		// on a LE system, this swaps big-endian to little-endian
		result[i] = binary.BigEndian.Uint64(val_bytes[startIdx:endIdx])
	}

	return result
}

// convert little-endian limbs to big.Int
func LimbsToInt(limbs []uint64) *big.Int {
	limbs_bytes := make([]byte, 8*len(limbs), 8*len(limbs))
	for i := 0; i < len(limbs); i++ {
		startIdx := (len(limbs) - (i + 1)) * 8
		endIdx := (len(limbs) - i) * 8

		binary.BigEndian.PutUint64(limbs_bytes[startIdx:endIdx], limbs[i])
	}

	return new(big.Int).SetBytes(limbs_bytes)
}

// **NOTE** naming confusing.  actually the second-largest modulus (largest would have modinv as 1)
func MaxModulus(limbCount uint) []uint64 {
	mod := make([]uint64, limbCount, limbCount)

	mod[0] = 0xfffffffffffffffd
	for i := uint(1); i < limbCount; i++ {
		mod[i] = 0xffffffffffffffff
	}

	return mod
}

// utility for unit testing.  returns  (1 << (((limbCount - 1) * limbBits) + limbBits / 2)) - 1
func GenTestModulus(limbCount uint) []uint64 {
	mod_int := big.NewInt(1)
	mod_int.Lsh(mod_int, (((limbCount - 1) * 64) + 32))
	mod_int.Sub(mod_int, big.NewInt(1))

	return IntToLimbs(mod_int, limbCount)
}

func LimbsEq(x, y []uint64) bool {
	if len(x) != len(y) {
		panic("unequally-sized elements")
	}

	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}

// return a hex-string literal for given limbs
/*
func LimbsToString(limbs []uint64) string {
	result := ""
	for i := 0; i < len(limbs); i++ {
		result += fmt.Sprintf("%06x", limbs[i])
	}
}
*/

func One(limbCount uint) []uint64 {
	one := make([]uint64, limbCount, limbCount)
	one[0] = 1
	return one
}

func RSquared(modulus []uint64) []uint64 {
	mod := LimbsToInt(modulus[:])
	r := new(big.Int)
	r.Exp(big.NewInt(2), big.NewInt(int64(len(modulus))*64), mod)
	r.Mul(r, r)
	r.Mod(r, mod)

	result := IntToLimbs(r, uint(len(modulus)))
	return result
}

// does the Python equivalent of pow(-modulus, -1, 1<<64)
func MontConstant_Interleaved(modulus []uint64) uint64 {
	mod_int := LimbsToInt(modulus)

	// 1<<64
	aux_mod, _ := new(big.Int).SetString("18446744073709551616", 10)
	negative_one, _ := new(big.Int).SetString("-1", 10)

	mod_int.Mul(mod_int, negative_one)
	return mod_int.ModInverse(mod_int, aux_mod).Uint64()
}
