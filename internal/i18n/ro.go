package i18n

var roMessages = map[string]string{
	ReaderStatusWaitingTerminalSize: "se așteaptă dimensiunea terminalului",
	ReaderStatusTerminalTooSmall:    "fereastra terminalului este prea mică",
	ReaderStatusLastPage:            "ultima pagină",
	ReaderStatusFirstPage:           "prima pagină",
	ReaderStatusRenderError:         "eroare de randare: %v",
	ReaderStatusMaximumZoom:         "zoom maxim",
	ReaderStatusMinimumZoom:         "zoom minim",

	ReaderViewTerminalTooSmall: "comicread: fereastra terminalului este prea mică",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "pagini %d/%d",
	ReaderViewPageRange:        "pagini %d-%d/%d",
	ReaderViewRendering:        "se redă",
	ReaderViewBookmarks:        "Semne de carte",
	ReaderViewNoBookmarks:      "(fără semne de carte)",
	ReaderViewBookmarksHelp:    "sus/jos mută | enter deschide | esc închide",
	ReaderViewHelp: `Taste

← →  pagina anterioară / următoare
↑ ↓  derulează o pagină mărită
+ -  mărește / micșorează
b    adaugă / elimină semn de carte
v ← → semnul de carte anterior / următor
c v  semne de carte
q    ieșire

?    închide ajutorul`,

	FilepickerHeader:      "comicread — selectați un capitol\n%s\n\n",
	FilepickerNoEntries:   "  (nicio intrare acceptată)\n",
	FilepickerHelp:        "\n↑/↓ deplasare  |  ← director părinte  |  → intră în director  |  enter deschide fișierul  |  s selectează directorul evidențiat  |  q ieșire\n",
	FilepickerWindowTitle: "comicread — alegeți un fișier",

	FilepickerErrResolveDir: "nu se poate determina directorul %q: %w",
	FilepickerErrReadDir:    "nu se poate citi directorul %q: %w",
	FilepickerErrRunPicker:  "eroare la pornirea selectorului de fișiere: %w",

	LoadingViewOpening:     "se deschide %s…",
	LoadingViewWindowTitle: "comicread — deschidere",

	CLIErrGetWorkingDir:             "nu se poate obține directorul de lucru: %w",
	CLIErrPickFile:                  "nu se poate alege fișierul: %w",
	CLIErrRunTUI:                    "eroare la pornirea TUI: %w",
	CLIErrParseArgs:                 "eroare la analizarea argumentelor: %w",
	CLIErrOpenChapter:               "nu se poate deschide capitolul: %w",
	CLIErrOpenJournal:               "nu se poate deschide jurnalul: %w",
	CLIErrClearJournal:              "nu se poate șterge jurnalul: %w",
	CLIErrClearJournalRequiresInput: "--clear-journal necesită un fișier sau director",
	CLIErrNoPages:                   "capitolul nu conține pagini imagine care pot fi citite",
	CLIErrInspectInput:              "nu se poate inspecta intrarea %q: %w",
	CLIErrUnsupportedFile:           "fișier neacceptat %q: formatele acceptate sunt CBZ, PDF, EPUB sau un director de imagini",
	CLIFlagGraphicsUsage:            "renderer: auto, ascii, dots, kitty, sixel sau iterm2",
	CLIFlagVersionUsage:             "afișează versiunea și ieși",
	CLIFlagUpdateUsage:              "verifică actualizările și ieși",
	CLIFlagEnvUsage:                 "afișează mediul comicread și ieși",
	CLIFlagClearJournalUsage:        "șterge jurnalul local pentru un fișier sau director și ieși",
	CLIFlagBookViewUsage:            "afișează perechi de pagini de la stânga la dreapta",
	CLIFlagRightBookViewUsage:       "afișează perechi de pagini de la dreapta la stânga",
	CLIFlagCircleBookViewUsage:      "afișează perechi de pagini suprapuse de la stânga la dreapta",
	CLIFlagRightCircleBookViewUsage: "afișează perechi de pagini suprapuse de la dreapta la stânga",
	CLIErrMultipleBookViews:         "poate fi utilizată o singură opțiune de vizualizare a cărții",
	CLIErrInvalidView:               "valoare COMICREAD_VIEW neacceptată %q (se așteaptă: book-view, right-view, circle-view sau right-circle-view)",
	CLIHelpHint:                     "rulați 'comicread --help' pentru ajutor",
	CLIUsage:                        "utilizare: comicread [opțiuni] [fișier]",
	CLIUsageFull: `comicread — un cititor de manga minimal pentru terminal

utilizare: comicread [opțiuni] [fișier]

opțiuni:
  --graphics string   renderer: auto, ascii, dots, kitty, sixel sau iterm2 (implicit "auto")
  --book-view         afișează perechi de pagini de la stânga la dreapta
  --right-view        afișează perechi de pagini de la dreapta la stânga
  --circle-view       afișează perechi de pagini suprapuse de la stânga la dreapta
  --right-circle-view
                      afișează perechi de pagini suprapuse de la dreapta la stânga
  --clear-journal    șterge jurnalul local pentru un fișier sau director și ieși
  --env               afișează mediul comicread și ieși
  --update            verifică actualizările și ieși
  -v, --version       afișează versiunea și ieși
  -h, --help          afișează acest ajutor

Dacă nu este dat niciun fișier sau director, se deschide un selector interactiv de fișiere în directorul curent.

mediu:
  COMICREAD_GRAPHICS  renderer implicit: auto, ascii, dots, kitty, sixel sau iterm2
  COMICREAD_PRERENDERED_NEXT      pagini următoare pentru prerenderizare (implicit 1)
  COMICREAD_PRERENDERED_PREVIOUS  pagini anterioare pentru prerenderizare (implicit 1)
  COMICREAD_VIEW      vizualizare implicită: book-view, right-view, circle-view sau right-circle-view
  COMICREAD_LANG      limba mesajelor: en, uk, pl, de, fr, es, cs, ro, it, ko, ja, id, hi, el, tr, kk sau ka (implicit "en")`,
}
