.PHONY: help build build-release build-linux build-linux-amd64 build-linux-arm64 build-linux-386 build-macos build-macos-amd64 build-macos-arm64 build-windows build-windows-amd64 build-windows-arm64 build-windows-386 run

BINARY := bin/comicread
VERSION := "dev-version"
LDFLAGS := -X github.com/arimatakao/comicread/internal/cli.Version=$(VERSION)

help:
	@printf '%s\n' \
		'Available targets:' \
		'  help                  Show this help message.' \
		'  run FILE=<path>       Run the reader with a chapter path.' \
		'  build                 Build the current platform binary.' \
		'  build-release         Build with version metadata.' \
		'  build-linux           Build all Linux binaries.' \
		'  build-linux-amd64     Build the Linux amd64 binary.' \
		'  build-linux-arm64     Build the Linux arm64 binary.' \
		'  build-linux-386       Build the Linux 386 binary.' \
		'  build-macos           Build all macOS binaries.' \
		'  build-macos-amd64     Build the macOS amd64 binary.' \
		'  build-macos-arm64     Build the macOS arm64 binary.' \
		'  build-windows         Build all Windows binaries.' \
		'  build-windows-amd64   Build the Windows amd64 binary.' \
		'  build-windows-arm64   Build the Windows arm64 binary.' \
		'  build-windows-386     Build the Windows 386 binary.'

run:
	go run . $(FILE)

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
