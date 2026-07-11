# comicread

The project is a minimal terminal manga reader written in Go. It opens CBZ, image-based PDF and EPUB files, or image directories, renders pages through the Kitty or Sixel graphics protocol, and supports keyboard navigation.

## Project Structure & Module Organization

```
.
├── AGENTS.md
├── CLAUDE.md
├── cmd
├── .gitignore
├── go.mod
├── go.sum
├── internal
│   ├── backend
│   ├── cli
│   └── tui
├── main.go
├── Makefile
└── README.md
```

Key directories:

- `cmd/` — reserved for standalone executable commands; currently empty.
- `internal/backend/` — terminal-rendering interfaces and terminal image processing protocol implementations.
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

# context-mode

context-mode MCP tools available. Rules protect context window from flooding. One unrouted command dumps 56 KB into context.

## When to use it

- `go build ./...`, `go vet ./...`, `go test ./...` output, or a wide `grep -r` across `internal/` can exceed a useful context size when something goes wrong (long compiler/test dumps). Route those through `ctx_batch_execute(commands, queries)` so only the relevant matches come back, instead of pasting the raw log into the conversation.
- Reading a `*_test.go` file or `.go` source file purely to **understand/summarize** it (not to `Edit` it) can go through `ctx_execute_file(path, language, code)`.
- Reading a file you're about to `Edit` should still use the normal `Read` tool — `Edit` needs the exact bytes to match against.

## Not needed here

- `git`, `mkdir`, `rm`, `mv`, `cd`, `ls`, `go build`, `go vet`, `go test`, `make build` — plain Bash is fine; output is small and there's no shared state to worry about (no ports, no lock files beyond `go.sum`).
- There's no web/API fetching or multi-URL work in this codebase, so `ctx_fetch_and_index` / concurrency batching don't apply — don't reach for them here.

## Memory

On resume, `ctx_search(sort: "timeline")` can surface prior decisions before asking the user again.

## ctx commands

| Command | Action |
|---------|--------|
| `ctx stats` | Call `ctx_stats` MCP tool, display full output verbatim |
| `ctx doctor` | Call `ctx_doctor` MCP tool, run returned shell command, display as checklist |
| `ctx upgrade` | Call `ctx_upgrade` MCP tool, run returned shell command, display as checklist |
| `ctx purge` | Call `ctx_purge` MCP tool with confirm: true. Warns before wiping knowledge base. |
