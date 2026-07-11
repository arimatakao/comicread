# comicread

Minimal terminal manga reader written in Go.

The current MVP opens CBZ, PDF, and EPUB files, renders pages with the Kitty graphics
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

Supported formats: CBZ, image-based PDF files, and image-based EPUB files.

PDF pages must contain one embedded raster image per page. EPUB pages must
reference their page images through the EPUB spine.
