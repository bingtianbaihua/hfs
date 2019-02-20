all: fmt build
	GOBIN="`pwd`/bin" go install -v ./cmd/...

build:
	go build -o ./bin/hfs ./cmd

clean:
	rm -rf bin

fmt:
	go fmt ./...
