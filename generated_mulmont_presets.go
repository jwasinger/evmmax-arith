package mont_arith






func NewMulMontImpls() []mulMontFunc {
    return []mulMontFunc {
        MulMont64,
            MulModMont64,
            MulModMont128,
            MulModMont192,
            MulModMont256,
            MulModMont320,
            MulModMont384,
            MulModMont448,
            MulModMont512,
            MulModMont576,
            MulModMont640,
            MulModMont704,
            MulModMont768,
    }
}
