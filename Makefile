.PHONY: build build-release run

BINARY := bin/comicread
VERSION := "dev-version"
LDFLAGS := -X github.com/arimatakao/comicread/internal/cli.Version=$(VERSION)

build:
	go build -o $(BINARY) .

build-release:
	go build -ldflags "$(LDFLAGS)" -o $(BINARY) .

run:
	go run . $(FILE)
