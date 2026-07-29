package i18n

var roMessages = map[string]string{
	CLIFlagConfigUsage: "fișier de configurare de utilizat", CLIFlagResetConfigUsage: "resetează config.toml la valorile implicite și ieși", CLIFlagSetConfigUsage: "actualizează config.toml: cheie=valoare",
	ReaderStatusWaitingTerminalSize: "se așteaptă dimensiunea terminalului",
	ReaderStatusTerminalTooSmall:    "fereastra terminalului este prea mică",
	ReaderStatusLastPage:            "ultima pagină",
	ReaderStatusFirstPage:           "prima pagină",
	ReaderStatusRenderError:         "eroare de randare: %v",
	ReaderStatusMaximumZoom:         "zoom maxim",
	ReaderStatusMinimumZoom:         "zoom minim",
	ReaderStatusInvalidPage:         "număr de pagină nevalid",

	ReaderViewTerminalTooSmall: "comicread: fereastra terminalului este prea mică",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "pagini %d/%d",
	ReaderViewPageRange:        "pagini %d-%d/%d",
	ReaderViewRendering:        "se redă",
	ReaderViewBookmarks:        "Semne de carte",
	ReaderViewNoBookmarks:      "(fără semne de carte)",
	ReaderViewBookmarksHelp:    "sus/jos mută | enter deschide | esc închide",
	ReaderViewGoToPage:         "Mergi la pagina: %s",
	ReaderViewHelp: `Taste

← →  pagina anterioară / următoare
↑ ↓  derulează o pagină mărită
+ -  mărește / micșorează
b    adaugă / elimină semn de carte
v ← → semnul de carte anterior / următor
c v  semne de carte
g123 enter  mergi la pagina
q    ieșire

?    închide ajutorul`,

	FilepickerHeader:         "comicread — selectați un capitol\n%s\n\n",
	FilepickerNoEntries:      "  (nicio intrare acceptată)\n",
	FilepickerHelp:           "\n↑/↓ deplasare\n← director părinte\n→ intră în director\nenter deschide fișierul\ns selectează directorul evidențiat\nf adaugă / elimină directorul curent din favorite\nF adaugă director favorit\nb directoare favorite\no mergi la un director\nq ieșire\n",
	FilepickerWindowTitle:    "comicread — alegeți un fișier",
	FilepickerGoToPrompt:     "\nMergi la director: %s\n",
	FilepickerFavoritePrompt: "\nDirector favorit: %s\n",
	FilepickerFavorites:      "Directoare favorite\n\n",
	FilepickerNoFavorites:    "  (nu sunt configurate directoare favorite)\n",
	FilepickerFavoritesHelp:  "\n↑/↓ deplasare\nenter mergi la director\nd elimină favoritul\nesc înapoi\n",
	FilepickerFavoriteErr:    "  eroare la salvarea favoritelor: %s\n",
	FilepickerGoToErr:        "  eroare: %s\n",

	FilepickerErrResolveDir: "nu se poate determina directorul %q: %w",
	FilepickerErrReadDir:    "nu se poate citi directorul %q: %w",
	FilepickerErrRunPicker:  "eroare la pornirea selectorului de fișiere: %w",
	FilepickerErrEmptyPath:  "calea este goală",
	FilepickerErrNotDir:     "%q nu este un director",

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
	CLIFlagOpenUsage:                "director de deschis în selectorul de fișiere (implicit: COMICREAD_DIR sau directorul curent)",
	CLIErrOpenNotDir:                "deschide directorul %q: nu este un director",
	CLIFlagWebUsage:                 "pornește un cititor web local în locul interfeței de terminal",
	CLIErrWebArgs:                   "--web nu acceptă un argument de fișier sau director",
	WebServerStarted:                "cititorul web comicread rulează la %s (apăsați Ctrl+C pentru oprire)",
	WebErrListen:                    "pornește serverul web: %w",
	WebErrServe:                     "rulează serverul web: %w",
	CLIHelpHint:                     "rulați 'comicread --help' pentru ajutor",
	CLIUsage:                        "utilizare: comicread [opțiuni] [fișier]",
	CLIUsageFull: `comicread — un cititor de manga minimal pentru terminal

utilizare: comicread [opțiuni] [fișier]

opțiuni:
  --config string     fișier de configurare de utilizat
  --graphics string   renderer: auto, ascii, dots, kitty, sixel sau iterm2 (implicit "auto")
  --book-view         afișează perechi de pagini de la stânga la dreapta
  --right-view        afișează perechi de pagini de la dreapta la stânga
  --circle-view       afișează perechi de pagini suprapuse de la stânga la dreapta
  --right-circle-view
                      afișează perechi de pagini suprapuse de la dreapta la stânga
  --clear-journal    șterge jurnalul local pentru un fișier sau director și ieși
  --reset-config     resetează config.toml la valorile implicite și ieși
  --set-config value actualizează config.toml: cheie=valoare
  -o, --open string   director de deschis în selectorul de fișiere (implicit: COMICREAD_DIR sau directorul curent)
  --update            verifică actualizările și ieși
  --web               pornește un cititor web local în locul interfeței de terminal
  -v, --version       afișează versiunea și ieși
  -h, --help          afișează acest ajutor

Dacă nu este dat niciun fișier sau director, se deschide un selector interactiv de fișiere în COMICREAD_DIR
(dacă este setat la un director valid) sau, altfel, în directorul curent.`,
	ReaderViewMetadata: "Metadate",
}
