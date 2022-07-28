package mont_arith






type ArithPreset struct {
	AddModImpls []arithFunc
	SubModImpls []arithFunc
	MulMontImpls []arithFunc
}

// Preset same as default except it uses blst's go-asm impl of the arithmetic at 384bit widths
func Asm384Preset() *ArithPreset {
	addModImpls := []arithFunc {
			AddModNonUnrolled128,
			AddModNonUnrolled192,
			AddModNonUnrolled256,
			AddModNonUnrolled320,
			AddMod384_asm,
			AddModNonUnrolled448,
			AddModNonUnrolled512,
			AddModNonUnrolled576,
			AddModNonUnrolled640,
			AddModNonUnrolled704,
			AddModNonUnrolled768,
	}

	subModImpls := []arithFunc {
			SubModNonUnrolled64,
			SubModNonUnrolled128,
			SubModNonUnrolled192,
			SubModNonUnrolled256,
			SubModNonUnrolled320,
			SubMod384_asm,
			SubModNonUnrolled448,
			SubModNonUnrolled512,
			SubModNonUnrolled576,
			SubModNonUnrolled640,
			SubModNonUnrolled704,
			SubModNonUnrolled768,
	}
	mulMontImpls := []arithFunc {

			
				mulMont64,

			
				mulMont128,

			
				mulMont192,

			
				mulMont256,

			
				mulMont320,
				MulMont384_asm,

			
				mulMont448,

			
				mulMont512,

			
				mulMont576,

			
				mulMont640,

			
				MulMontNonInterleaved,

			
				MulMontNonInterleaved,
	}

	return &ArithPreset{addModImpls, subModImpls, mulMontImpls}
}

func DefaultPreset() *ArithPreset {
	addModImpls := []arithFunc {
			AddModNonUnrolled64,
			AddModNonUnrolled128,
			AddModNonUnrolled192,
			AddModNonUnrolled256,
			AddModNonUnrolled320,
			AddModNonUnrolled384,
			AddModNonUnrolled448,
			AddModNonUnrolled512,
			AddModNonUnrolled576,
			AddModNonUnrolled640,
			AddModNonUnrolled704,
			AddModNonUnrolled768,
	}

	subModImpls := []arithFunc {
			SubModNonUnrolled64,
			SubModNonUnrolled128,
			SubModNonUnrolled192,
			SubModNonUnrolled256,
			SubModNonUnrolled320,
			SubModNonUnrolled384,
			SubModNonUnrolled448,
			SubModNonUnrolled512,
			SubModNonUnrolled576,
			SubModNonUnrolled640,
			SubModNonUnrolled704,
			SubModNonUnrolled768,
	}
	mulMontImpls := []arithFunc {
			mulMont64,
			mulMont128,
			mulMont192,
			mulMont256,
			mulMont320,
			mulMont384,
			mulMont448,
			mulMont512,
			mulMont576,
			mulMont640,
			mulMont704,
			mulMont768,
	}

	return &ArithPreset{addModImpls, subModImpls, mulMontImpls}
}
