hello:
	echo "Hello"

build:
	GOFLAGS=-mod=mod go build -o bin/dka main.go 

install: build
	sudo mv dkargo /usr/local/bin/dkargo
run:
	go run main.go

all: build install