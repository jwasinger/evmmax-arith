.PHONY: build

clean:
	rm generated*.go

build:
	cd arith_generator && go build && cd ..  && ./arith_generator/arith_generator

test:
	go test -run=.

benchmark:
	go test -bench=.
