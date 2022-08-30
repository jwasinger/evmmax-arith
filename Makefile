.PHONY: build

clean:
	rm generated*.go

build:
	cd generator && go build && cd ..  && ./generator/generator 64 
	gofmt -s -w generated_addmod.go generated_mulmont.go generated_submod.go generated_presets.go

test:
	go test -run=.

benchmark:
	go test -bench=.
