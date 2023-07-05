BINARY_NAME=mdbtn

build:
	go build -o bin/$(BINARY_NAME) -v

run:
	./bin/$(BINARY_NAME)

test: build run