.PHONY=build
build:
	go build
.PHONY=run
run:
	export PORT=8080
	go run main.go
default: build
