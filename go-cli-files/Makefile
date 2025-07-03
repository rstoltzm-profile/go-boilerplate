BINARY=cli-template

.PHONY: all build build-linux build-windows test clean

all: build

build: build-linux build-windows

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/$(BINARY) main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/$(BINARY).exe main.go

test:
	go test ./internal -v

integ-test:
	./test-cli-linux.sh

clean:
	rm -f bin/$(BINARY)-linux bin/$(BINARY)-windows.exe