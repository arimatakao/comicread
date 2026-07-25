package i18n

var enMessages = map[string]string{
	CLIFlagConfigUsage: "configuration file to use", CLIFlagResetConfigUsage: "reset config.toml to its defaults and exit", CLIFlagSetConfigUsage: "update config.toml: key=value",
	ReaderStatusWaitingTerminalSize: "waiting for terminal size",
	ReaderStatusTerminalTooSmall:    "terminal window is too small",
	ReaderStatusLastPage:            "last page",
	ReaderStatusFirstPage:           "first page",
	ReaderStatusRenderError:         "render error: %v",
	ReaderStatusMaximumZoom:         "maximum zoom",
	ReaderStatusMinimumZoom:         "minimum zoom",
	ReaderStatusInvalidPage:         "invalid page number",

	ReaderViewTerminalTooSmall: "comicread: terminal window is too small",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "pages %d/%d",
	ReaderViewPageRange:        "pages %d-%d/%d",
	ReaderViewRendering:        "rendering",
	ReaderViewBookmarks:        "Bookmarks",
	ReaderViewNoBookmarks:      "(no bookmarks)",
	ReaderViewBookmarksHelp:    "up/down move | enter open | esc close",
	ReaderViewGoToPage:         "Go to page: %s",
	ReaderViewHelp: `Keys

← →  previous / next page
↑ ↓  scroll a zoomed page
+ -  zoom in / out
b    add / remove bookmark
v ← → previous / next bookmark
c v  bookmarks
g123 enter  go to page
q    quit

?    close help`,

	FilepickerHeader:         "comicread — select a chapter\n%s\n\n",
	FilepickerNoEntries:      "  (no supported entries)\n",
	FilepickerHelp:           "\n↑/↓  move\n←    parent dir\n→    enter dir\nenter  open file\ns      select highlighted directory\nf      add current directory to favorites\nF      add favorite directory\nb      favorite directories\no      go to directory\nq      quit\n",
	FilepickerWindowTitle:    "comicread — pick a file",
	FilepickerGoToPrompt:     "\nGo to directory: %s\n",
	FilepickerFavoritePrompt: "\nFavorite directory: %s\n",
	FilepickerGoToErr:        "  error: %s\n",
	FilepickerFavorites:      "Favorite directories\n\n",
	FilepickerNoFavorites:    "  (no configured favorite directories)\n",
	FilepickerFavoritesHelp:  "\n↑/↓  move\nenter  go to directory\nesc    return\n",
	FilepickerFavoriteErr:    "  error: save favorites: %s\n",

	FilepickerErrResolveDir: "resolve directory %q: %w",
	FilepickerErrReadDir:    "read directory %q: %w",
	FilepickerErrRunPicker:  "run file picker: %w",
	FilepickerErrEmptyPath:  "path is empty",
	FilepickerErrNotDir:     "%q is not a directory",

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
	CLIFlagUpdateUsage:              "check for updates and exit",
	CLIFlagClearJournalUsage:        "remove the local journal for a file or directory and exit",
	CLIFlagBookViewUsage:            "show pages left to right in pairs",
	CLIFlagRightBookViewUsage:       "show pages right to left in pairs",
	CLIFlagCircleBookViewUsage:      "show overlapping page pairs left to right",
	CLIFlagRightCircleBookViewUsage: "show overlapping page pairs right to left",
	CLIErrMultipleBookViews:         "only one book view option may be used",
	CLIErrInvalidView:               "unsupported view %q (want single-page, book-view, right-view, circle-view, or right-circle-view)",
	CLIFlagOpenUsage:                "directory to open in the file picker (default: configured directory or current directory)",
	CLIErrOpenNotDir:                "open directory %q: not a directory",
	CLIHelpHint:                     "run 'comicread --help' for usage",
	CLIUsage:                        "usage: comicread [options] [file]",
	CLIUsageFull: `comicread - a minimal terminal manga reader

usage: comicread [options] [file]

options:
  --config string     configuration file to use
  --graphics string   renderer: auto, ascii, dots, kitty, sixel, or iterm2 (default "auto")
  --book-view         show pages left to right in pairs
  --right-view        show pages right to left in pairs
  --circle-view       show overlapping page pairs left to right
  --right-circle-view
                      show overlapping page pairs right to left
  --clear-journal    remove the local journal for a file or directory and exit
  --reset-config     reset config.toml to its defaults and exit
  --set-config value update config.toml: key=value
  -o, --open string   directory to open in the file picker (default: configured directory or the current directory)
  --update            check for updates and exit
  -v, --version       print version and exit
  -h, --help          show this help message

If no file or directory is given, an interactive file picker opens in the configured
directory (if valid) or the current directory otherwise.`,
}
