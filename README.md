# comicread

Minimal terminal manga reader written in Go.

The current MVP opens CBZ, PDF, EPUB, and image directories, renders pages with
the Kitty, Sixel, or iTerm2 graphics protocol, and supports forward/backward navigation.

## Requirements

- Go 1.26.5 or newer
- a terminal implementing the Kitty graphics protocol (such as Kitty or Ghostty)
  or the Sixel protocol (such as mlterm, mintty, foot, or xterm built with Sixel)
  or iTerm2

## Run

```sh
go run . path/to/chapter.cbz
go run . path/to/image-directory
go run . --graphics sixel path/to/chapter.cbz
go run . --graphics iterm2 path/to/chapter.cbz
```

`--graphics` accepts `auto` (the default), `kitty`, `sixel`, or `iterm2`.
Auto-detection selects iTerm2 and known Sixel terminals from their environment
variables; use an explicit protocol for terminals that do not identify
themselves reliably.

Keys:

- `right`, `l`, `space`, `j`, `PageDown`: next page
- `left`, `h`, `backspace`, `k`, `PageUp`: previous page
- `+`, `-`: zoom in/out for every page in the open file
- `up`, `down`: scroll the zoomed page vertically
- `q`, `Esc`, `Ctrl+C`: quit

Supported formats: CBZ, image-based PDF files, image-based EPUB files, and directories containing image files.

PDF pages must contain one embedded raster image per page. EPUB pages must
reference their page images through the EPUB spine.
