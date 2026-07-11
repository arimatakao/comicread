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

```sh
make build # write the executable to bin/comicread
```

## Coding Style & Naming Conventions

Follow standard Go style and run `gofmt` on every changed `.go` file. Use tabs for indentation as emitted by `gofmt`; do not hand-align code. Keep package names short and lowercase (`cli`, `tui`, `backend`). Export only cross-package APIs and use Go-style names: `New`, `OpenChapter`, `TotalPages`; use concise unexported names such as `imageArea` or `validateInput` within a package.

Keep terminal escape-sequence handling in `internal/backend`, and keep navigation and UI state in `internal/tui`. Return contextual errors rather than panicking for bad files or invalid rendering input.
