package i18n

var enMessages = map[string]string{
	"reader.status.waiting_terminal_size": "waiting for terminal size",
	"reader.status.terminal_too_small":    "terminal window is too small",
	"reader.status.last_page":             "last page",
	"reader.status.first_page":            "first page",
	"reader.status.render_error":          "render error: %v",
	"reader.status.maximum_zoom":          "maximum zoom",
	"reader.status.minimum_zoom":          "minimum zoom",

	"reader.view.terminal_too_small": "comicread: terminal window is too small",
	"reader.view.window_title":       "comicread — %s",
	"reader.view.pages":              "pages %d/%d",
	"reader.view.rendering":          "rendering",

	"filepicker.header":       "comicread — select a chapter\n%s\n\n",
	"filepicker.no_entries":   "  (no supported entries)\n",
	"filepicker.help":         "\n↑/↓ move  ← parent dir  → enter dir  enter open file  s select highlighted directory  q quit\n",
	"filepicker.window_title": "comicread — pick a file",

	"filepicker.err.resolve_dir": "resolve directory %q: %w",
	"filepicker.err.read_dir":    "read directory %q: %w",
	"filepicker.err.run_picker":  "run file picker: %w",

	"cli.err.get_working_dir":  "get working directory: %w",
	"cli.err.pick_file":        "pick file: %w",
	"cli.err.run_tui":          "run TUI: %w",
	"cli.err.parse_args":       "parse arguments: %w",
	"cli.err.open_chapter":     "open chapter: %w",
	"cli.err.no_pages":         "chapter contains no readable image pages",
	"cli.err.inspect_input":    "inspect input %q: %w",
	"cli.err.unsupported_file": "unsupported file %q: supported formats are CBZ, PDF, EPUB, or an image directory",
	"cli.flag.graphics_usage":  "graphics protocol: auto, kitty, sixel, or iterm2",
	"cli.usage":                "usage: comicread [--graphics auto|kitty|sixel|iterm2] [file.cbz|file.pdf|file.epub|image-directory]",
}
