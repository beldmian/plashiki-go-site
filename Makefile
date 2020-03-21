.PHONY=build
build:
	go build
.PHONY=run
run:
	set PORT=8080
	go run ./main.go
default: build
