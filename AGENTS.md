# comicread

The project is a minimal terminal manga reader written in Go. It opens CBZ, image-based PDF and EPUB files, or image directories, renders pages through the Kitty or Sixel graphics protocol, and supports keyboard navigation.

## Project Structure & Module Organization

```
.
├── AGENTS.md
├── CLAUDE.md -> AGENTS.md
├── cmd
├── .gitignore
├── go.mod
├── go.sum
├── internal
│   ├── backend
│   │   ├── backend.go
│   │   ├── kitty.go
│   │   ├── kitty_test.go
│   │   ├── renderer.go
│   │   ├── renderer_test.go
│   │   ├── sixel.go
│   │   └── sixel_test.go
│   ├── cli
│   │   ├── cli.go
│   │   └── cli_test.go
│   └── tui
│       ├── commands.go
│       ├── keys.go
│       ├── messages.go
│       ├── model.go
│       ├── update.go
│       ├── update_test.go
│       └── view.go
├── main.go
├── Makefile
└── README.md
```

Key directories:

- `cmd/` — reserved for standalone executable commands; currently empty.
- `internal/backend/` — terminal-rendering interfaces and the Kitty/Sixel protocol implementations.
- `internal/cli/` — input-path validation and opening CBZ, PDF, EPUB, or image-directory chapters.
- `internal/tui/` — the Bubble Tea model, key handling, navigation, and view rendering.

`main.go` starts the application. Tests live beside the code they cover as `*_test.go`; `Makefile` writes the binary to `bin/`.

## Guidelines

DO NOT:
- Update README.md
- Add new external dependencies/packages
- Update/add tests (*_test.go)

Do these only if the user explicitly asks for it.

If any .go files were changed, run `go vet ./...` and `make build` to confirm the code compiles.

