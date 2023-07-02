BINARY_NAME=mdbtn

build:
	go build -o bin/$(BINARY_NAME) -v

run: build
	./bin/$(BINARY_NAME)