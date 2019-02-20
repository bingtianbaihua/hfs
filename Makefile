all: fmt build

build:
	go build -o ./bin/hfs ./cmd

clean:
	rm -rf bin

fmt:
	go fmt ./...
