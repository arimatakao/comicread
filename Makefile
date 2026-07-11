.PHONY: build run

BINARY := bin/comicread

build:
	go build -o $(BINARY) ./cmd/comicread

run:
	go run ./cmd/comicread $(FILE)
