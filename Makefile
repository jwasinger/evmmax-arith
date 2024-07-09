.PHONY: build

build:
	cd generator && go build && cd ..  && ./generator/generator 64 
	gofmt -s -w mulmont-generated.go

test:
	go test -run=.

benchmark:
	go test -bench=.
