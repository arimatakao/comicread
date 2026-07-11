.PHONY: build run

BINARY := bin/comicread

build:
	go build -o $(BINARY) .

run:
	go run . $(FILE)
