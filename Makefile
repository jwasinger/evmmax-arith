.PHONY: build

clean:
	rm generated*.go

build:
	cd arith_generator && go build && cd ..  && ./arith_generator/arith_generator
	gofmt -s -w generated_addmod.go generated_mulmont.go generated_submod.go generated_presets.go

test:
	go test -run=.

benchmark:
	go test -bench=.
