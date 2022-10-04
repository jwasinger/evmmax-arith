# evmmax-arith

Library for performing modular addition, subtraction and Montgomery multiplication

## Usage

Build the code generator and generate the arithmetic code:
```
(cd generator && go build)
make build
```

Run benchmarks:
```
go test -bench=.
```

Run tests:
```
go test -run=.
```
