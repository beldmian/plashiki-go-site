.PHONY=build
build:
	go build -v ./cmd/server

default: build