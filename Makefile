.PHONY=build
build:
	go build
.PHONY=run
run:
	go run ./main.go
default: build
