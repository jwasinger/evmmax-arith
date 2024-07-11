package evmmax_arith

import (
	"fmt"
	"math/big"
	"math/rand"
	"testing"
)

func benchmarkOp(b *testing.B, op string, mod *big.Int) {
	fieldCtx, err := NewFieldContext(mod.Bytes(), 256)
	if err != nil {
		panic(err)
	}
	xIdxs := make([]int, 256)
	yIdxs := make([]int, 256)
	outIdxs := make([]int, 256)
	for i := 0; i < 256; i++ {
		outIdxs[i] = rand.Intn(255)
		xIdxs[i] = rand.Intn(255)
		yIdxs[i] = rand.Intn(255)
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

func BenchmarkOps(b *testing.B) {
	for i := 1; i <= 12; i++ {
		limbs := MaxModulus(i)
		mod := limbsToInt(limbs)

		b.Run(fmt.Sprintf("add-%d-bit", i*64), func(b *testing.B) {
			benchmarkOp(b, "add", mod)
		})
		b.Run(fmt.Sprintf("sub-%d-bit", i*64), func(b *testing.B) {
			benchmarkOp(b, "sub", mod)
		})
		b.Run(fmt.Sprintf("mul-%d-bit", i*64), func(b *testing.B) {
			benchmarkOp(b, "mul", mod)
		})
	}
}
