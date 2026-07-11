# comicread

Minimal terminal manga reader written in Go.

The current MVP opens CBZ, PDF, EPUB, and image directories, renders pages with
the Kitty or Sixel graphics protocol, and supports forward/backward navigation.

## Requirements

- Go 1.26.5 or newer
- a terminal implementing the Kitty graphics protocol (such as Kitty or Ghostty)
  or the Sixel protocol (such as mlterm, mintty, foot, or xterm built with Sixel)

## Run

```sh
go run . path/to/chapter.cbz
go run . path/to/image-directory
go run . --graphics sixel path/to/chapter.cbz
```

`--graphics` accepts `auto` (the default), `kitty`, or `sixel`. Auto-detection
uses Sixel for terminals that identify as known Sixel terminals; use
`--graphics sixel` for terminals such as xterm where Sixel support cannot be
reliably identified from environment variables.

Keys:

- `right`, `l`, `space`, `j`, `PageDown`: next page
- `left`, `h`, `backspace`, `k`, `PageUp`: previous page
- `+`, `-`: zoom in/out for every page in the open file
- `up`, `down`: scroll the zoomed page vertically
- `q`, `Esc`, `Ctrl+C`: quit

Supported formats: CBZ, image-based PDF files, image-based EPUB files, and directories containing image files.

PDF pages must contain one embedded raster image per page. EPUB pages must
reference their page images through the EPUB spine.
