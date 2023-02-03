.PHONY: build

clean:
	rm -f generated*.go

build:
	cd generator && go build && cd ..  && ./generator/generator 64 
	gofmt -s -w generated_addmod_unrolled.go generated_addmod_nonunrolled.go generated_mulmont_unrolled.go generated_mulmont_nonunrolled.go generated_submod_unrolled.go generated_submod_nonunrolled.go generated_presets.go

test:
	go test -run=.

benchmark:
	go test -bench=.
