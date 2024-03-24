package evmmax_arith

import (
	"math/rand"
	"testing"
)

func benchmarkMulMod(b *testing.B, limbCount int) {
	mod := MaxModulus(uint(limbCount))
	modState, _ := NewModulusState(mod, 256)

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
		modState.MulMod(modState.modInv, modState.Modulus,
			modState.scratchSpace[outIdxs[i%256]:outIdxs[i%256]+48],
			modState.scratchSpace[xIdxs[i%256]:xIdxs[i%256]+48],
			modState.scratchSpace[yIdxs[i%256]:yIdxs[i%256]+48])
	}
}

type opFn func(*testing.B, uint)

func BenchmarkOps(b *testing.B) {
	for i := 0; i < 12; i++ {
		benchmarkMulMod(b, i)
	}
}
