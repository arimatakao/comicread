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
	ReaderViewBookmarks:        "Bookmarks",
	ReaderViewNoBookmarks:      "(no bookmarks)",
	ReaderViewBookmarksHelp:    "up/down move | enter open | esc close",
	ReaderViewHelp: `Keys

← →  previous / next page
↑ ↓  scroll a zoomed page
+ -  zoom in / out
b    add / remove bookmark
v ← → previous / next bookmark
c v  bookmarks
q    quit

?    close help`,

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
	CLIErrOpenJournal:               "open journal: %w",
	CLIErrClearJournal:              "clear journal: %w",
	CLIErrClearJournalRequiresInput: "--clear-journal requires a file or directory",
	CLIErrNoPages:                   "chapter contains no readable image pages",
	CLIErrInspectInput:              "inspect input %q: %w",
	CLIErrUnsupportedFile:           "unsupported file %q: supported formats are CBZ, PDF, EPUB, or an image directory",
	CLIFlagGraphicsUsage:            "renderer: auto, ascii, dots, kitty, sixel, or iterm2",
	CLIFlagVersionUsage:             "print version and exit",
	CLIFlagEnvUsage:                 "print comicread environment and exit",
	CLIFlagClearJournalUsage:        "remove the local journal for a file or directory and exit",
	CLIFlagBookViewUsage:            "show pages left to right in pairs",
	CLIFlagRightBookViewUsage:       "show pages right to left in pairs",
	CLIFlagCircleBookViewUsage:      "show overlapping page pairs left to right",
	CLIFlagRightCircleBookViewUsage: "show overlapping page pairs right to left",
	CLIErrMultipleBookViews:         "only one book view option may be used",
	CLIErrInvalidView:               "unsupported COMICREAD_VIEW %q (want book-view, right-view, circle-view, or right-circle-view)",
	CLIHelpHint:                     "run 'comicread --help' for usage",
	CLIUsage:                        "usage: comicread [options] [file]",
	CLIUsageFull: `comicread - a minimal terminal manga reader

usage: comicread [options] [file]

options:
  --graphics string   renderer: auto, ascii, dots, kitty, sixel, or iterm2 (default "auto")
  --book-view         show pages left to right in pairs
  --right-view        show pages right to left in pairs
  --circle-view       show overlapping page pairs left to right
  --right-circle-view
                      show overlapping page pairs right to left
  --clear-journal    remove the local journal for a file or directory and exit
  --env               print comicread environment and exit
  -v, --version       print version and exit
  -h, --help          show this help message

If no file or directory is given, an interactive file picker opens in the current directory.

environment:
  COMICREAD_GRAPHICS  renderer default: auto, ascii, dots, kitty, sixel, or iterm2
  COMICREAD_VIEW      default view: book-view, right-view, circle-view, or right-circle-view
  COMICREAD_LANG   message language: en, uk, pl, de, fr, es, cs, ro, it, ko, ja, id, hi, el, tr, kk, or ka (default "en")`,
}
