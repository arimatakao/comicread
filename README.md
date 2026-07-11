# comicread

Minimal terminal manga reader written in Go.

The current MVP opens CBZ files, renders pages with the Kitty graphics
protocol, and supports forward/backward navigation.

## Requirements

- Go 1.26.5 or newer
- a terminal implementing the Kitty graphics protocol, such as Kitty or Ghostty

## Run

```sh
go run . path/to/chapter.cbz
```

Keys:

- `right`, `l`, `space`, `j`, `PageDown`: next page
- `left`, `h`, `backspace`, `k`, `PageUp`: previous page
- `q`, `Esc`, `Ctrl+C`: quit

Only CBZ is supported by this first version.
