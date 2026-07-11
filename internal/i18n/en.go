package i18n

var enMessages = map[string]string{
	ReaderStatusWaitingTerminalSize: "waiting for terminal size",
	ReaderStatusTerminalTooSmall:    "terminal window is too small",
	ReaderStatusLastPage:            "last page",
	ReaderStatusFirstPage:           "first page",
	ReaderStatusRenderError:         "render error: %v",
	ReaderStatusMaximumZoom:         "maximum zoom",
	ReaderStatusMinimumZoom:         "minimum zoom",

	ReaderViewTerminalTooSmall: "comicread: terminal window is too small",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "pages %d/%d",
	ReaderViewPageRange:        "pages %d-%d/%d",
	ReaderViewRendering:        "rendering",

	FilepickerHeader:      "comicread — select a chapter\n%s\n\n",
	FilepickerNoEntries:   "  (no supported entries)\n",
	FilepickerHelp:        "\n↑/↓ move  |  ← parent dir  |  → enter dir  |  enter open file  |  s select highlighted directory  |  q quit\n",
	FilepickerWindowTitle: "comicread — pick a file",

	FilepickerErrResolveDir: "resolve directory %q: %w",
	FilepickerErrReadDir:    "read directory %q: %w",
	FilepickerErrRunPicker:  "run file picker: %w",

	LoadingViewOpening:     "opening %s…",
	LoadingViewWindowTitle: "comicread — opening",

	CLIErrGetWorkingDir:             "get working directory: %w",
	CLIErrPickFile:                  "pick file: %w",
	CLIErrRunTUI:                    "run TUI: %w",
	CLIErrParseArgs:                 "parse arguments: %w",
	CLIErrOpenChapter:               "open chapter: %w",
	CLIErrNoPages:                   "chapter contains no readable image pages",
	CLIErrInspectInput:              "inspect input %q: %w",
	CLIErrUnsupportedFile:           "unsupported file %q: supported formats are CBZ, PDF, EPUB, or an image directory",
	CLIFlagGraphicsUsage:            "renderer: auto, ascii, dots, kitty, sixel, or iterm2",
	CLIFlagVersionUsage:             "print version and exit",
	CLIFlagBookViewUsage:            "show pages left to right in pairs",
	CLIFlagRightBookViewUsage:       "show pages right to left in pairs",
	CLIFlagCircleBookViewUsage:      "show overlapping page pairs left to right",
	CLIFlagRightCircleBookViewUsage: "show overlapping page pairs right to left",
	CLIErrMultipleBookViews:         "only one book view option may be used",
	CLIHelpHint:                     "run 'comicread --help' for usage",
	CLIUsage:                        "usage: comicread [--graphics auto|ascii|dots|kitty|sixel|iterm2] [--book-view|--right-book-view|--circle-book-view|--right-circle-book-view] [file.cbz|file.pdf|file.epub|image-directory]",
	CLIUsageFull: `comicread - a minimal terminal manga reader

usage: comicread [--graphics auto|ascii|dots|kitty|sixel|iterm2] [--book-view|--right-book-view|--circle-book-view|--right-circle-book-view] [file.cbz|file.pdf|file.epub|image-directory]

options:
  --graphics string   renderer: auto, ascii, dots, kitty, sixel, or iterm2 (default "auto")
  --book-view         show pages left to right in pairs
  --right-book-view   show pages right to left in pairs
  --circle-book-view  show overlapping page pairs left to right
  --right-circle-book-view
                      show overlapping page pairs right to left
  -v, --version       print version and exit
  -h, --help          show this help message

If no file or directory is given, an interactive file picker opens in the current directory.

environment:
  COMICREAD_LANG   message language: en or uk (default "en")`,
}
