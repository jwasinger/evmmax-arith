package evmmax_arith

import (
	"fmt"
	"math/big"
	"math/rand"
	"testing"
)

func benchmarkOp(b *testing.B, op string, mod *big.Int, preset384 int) {
	fieldCtx, err := NewFieldContext(mod.Bytes(), 256, preset384)
	if err != nil {
		panic(err)
	}
	xIdxs := make([]uint, 256)
	yIdxs := make([]uint, 256)
	outIdxs := make([]uint, 256)
	for i := 0; i < 256; i++ {
		outIdxs[i] = uint(rand.Intn(255))
		xIdxs[i] = uint(rand.Intn(255))
		yIdxs[i] = uint(rand.Intn(255))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		switch op {
		case "add":
			fieldCtx.AddMod(outIdxs[i%256], xIdxs[i%256], yIdxs[i%256])
		case "sub":
			fieldCtx.SubMod(outIdxs[i%256], xIdxs[i%256], yIdxs[i%256])
		case "mul":
			fieldCtx.MulMod(outIdxs[i%256], xIdxs[i%256], yIdxs[i%256])
		default:
			panic("invalid op")
		}
	}
}

func benchmarkSetmod(b *testing.B, mod *big.Int) {
	for i := 0; i < b.N; i++ {
		_, err := NewFieldContext(mod.Bytes(), 1)
		if err != nil {
			panic(err)
		}
	}		
}

func BenchmarkOps(b *testing.B) {
	for i := 1; i <= 12; i++ {
		limbs := MaxModulus(i)
		mod := limbsToInt(limbs)

		b.Run(fmt.Sprintf("add-%d-bit", i*64), func(b *testing.B) {
			benchmarkOp(b, "add", mod, FallBackOnly)
		})
		b.Run(fmt.Sprintf("sub-%d-bit", i*64), func(b *testing.B) {
			benchmarkOp(b, "sub", mod, FallBackOnly)
		})
		b.Run(fmt.Sprintf("mul-%d-bit", i*64), func(b *testing.B) {
			benchmarkOp(b, "mul", mod, FallBackOnly)
		})
		b.Run(fmt.Sprintf("setmod-%d-bit", i*64), func(b *testing.B) {
			benchmarkSetmod(b, mod)
		})
	}

	limbs := MaxModulus(6)
	mod := limbsToInt(limbs)

	b.Run(fmt.Sprintf("mul-%d-bit-asm", 384), func(b *testing.B) {
		benchmarkOp(b, "mul", mod, AllAsm)
	})
	b.Run(fmt.Sprintf("add-%d-bit-asm", 384), func(b *testing.B) {
		benchmarkOp(b, "add", mod, AllAsm)
	})
	b.Run(fmt.Sprintf("sub-%d-bit-asm", 384), func(b *testing.B) {
		benchmarkOp(b, "sub", mod, AllAsm)
	})
}
