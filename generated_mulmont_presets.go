package mont_arith






func NewMulMontImpls() []mulMontFunc {
    return []mulMontFunc {
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
}
