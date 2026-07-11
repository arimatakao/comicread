.PHONY: build build-release build-linux build-linux-amd64 build-linux-arm64 build-linux-386 build-macos build-macos-amd64 build-macos-arm64 build-windows build-windows-amd64 build-windows-arm64 build-windows-386 run

BINARY := bin/comicread
VERSION := "dev-version"
LDFLAGS := -X github.com/arimatakao/comicread/internal/cli.Version=$(VERSION)

build:
	go build -o $(BINARY) .

build-release:
	go build -ldflags "$(LDFLAGS)" -o $(BINARY) .

build-linux: build-linux-amd64 build-linux-arm64 build-linux-386

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o bin/comicread-linux-amd64 .

build-linux-arm64:
	GOOS=linux GOARCH=arm64 go build -o bin/comicread-linux-arm64 .

build-linux-386:
	GOOS=linux GOARCH=386 go build -o bin/comicread-linux-386 .

build-macos: build-macos-amd64 build-macos-arm64

build-macos-amd64:
	GOOS=darwin GOARCH=amd64 go build -o bin/comicread-darwin-amd64 .

build-macos-arm64:
	GOOS=darwin GOARCH=arm64 go build -o bin/comicread-darwin-arm64 .

build-windows: build-windows-amd64 build-windows-arm64 build-windows-386

build-windows-amd64:
	GOOS=windows GOARCH=amd64 go build -o bin/comicread-windows-amd64.exe .

build-windows-arm64:
	GOOS=windows GOARCH=arm64 go build -o bin/comicread-windows-arm64.exe .

build-windows-386:
	GOOS=windows GOARCH=386 go build -o bin/comicread-windows-386.exe .

run:
	go run . $(FILE)
