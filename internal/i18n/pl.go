package i18n

var plMessages = map[string]string{
	ReaderStatusWaitingTerminalSize: "oczekiwanie na rozmiar terminala",
	ReaderStatusTerminalTooSmall:    "okno terminala jest za małe",
	ReaderStatusLastPage:            "ostatnia strona",
	ReaderStatusFirstPage:           "pierwsza strona",
	ReaderStatusRenderError:         "błąd renderowania: %v",
	ReaderStatusMaximumZoom:         "maksymalne powiększenie",
	ReaderStatusMinimumZoom:         "minimalne powiększenie",
	ReaderStatusInvalidPage:         "nieprawidłowy numer strony",

	ReaderViewTerminalTooSmall: "comicread: okno terminala jest za małe",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "strony %d/%d",
	ReaderViewPageRange:        "strony %d-%d/%d",
	ReaderViewRendering:        "renderowanie",
	ReaderViewBookmarks:        "Zakładki",
	ReaderViewNoBookmarks:      "(brak zakładek)",
	ReaderViewBookmarksHelp:    "góra/dół ruch | enter otwórz | esc zamknij",
	ReaderViewGoToPage:         "Przejdź do strony: %s",
	ReaderViewHelp: `Klawisze

← →  poprzednia / następna strona
↑ ↓  przewijanie powiększonej strony
+ -  powiększ / pomniejsz
b    dodaj / usuń zakładkę
v ← → poprzednia / następna zakładka
c v  zakładki
g123 enter  przejdź do strony
q    wyjście

?    zamknij pomoc`,

	FilepickerHeader:      "comicread — wybierz rozdział\n%s\n\n",
	FilepickerNoEntries:   "  (brak obsługiwanych elementów)\n",
	FilepickerHelp:        "\n↑/↓ ruch\n← katalog nadrzędny\n→ wejdź do katalogu\nenter otwórz plik\ns wybierz zaznaczony katalog\no przejdź do katalogu\nq wyjście\n",
	FilepickerWindowTitle: "comicread — wybór pliku",
	FilepickerGoToPrompt:  "\nPrzejdź do katalogu: %s\n",
	FilepickerGoToErr:     "  błąd: %s\n",

	FilepickerErrResolveDir: "nie można ustalić katalogu %q: %w",
	FilepickerErrReadDir:    "nie można odczytać katalogu %q: %w",
	FilepickerErrRunPicker:  "błąd uruchamiania wyboru pliku: %w",
	FilepickerErrEmptyPath:  "ścieżka jest pusta",
	FilepickerErrNotDir:     "%q nie jest katalogiem",

	LoadingViewOpening:     "otwieranie %s…",
	LoadingViewWindowTitle: "comicread — otwieranie",

	CLIErrGetWorkingDir:             "nie można uzyskać katalogu roboczego: %w",
	CLIErrPickFile:                  "nie można wybrać pliku: %w",
	CLIErrRunTUI:                    "błąd uruchamiania TUI: %w",
	CLIErrParseArgs:                 "błąd analizy argumentów: %w",
	CLIErrOpenChapter:               "nie można otworzyć rozdziału: %w",
	CLIErrOpenJournal:               "nie można otworzyć dziennika: %w",
	CLIErrClearJournal:              "nie można usunąć dziennika: %w",
	CLIErrClearJournalRequiresInput: "--clear-journal wymaga pliku lub katalogu",
	CLIErrNoPages:                   "rozdział nie zawiera czytelnych stron obrazów",
	CLIErrInspectInput:              "nie można sprawdzić danych wejściowych %q: %w",
	CLIErrUnsupportedFile:           "nieobsługiwany plik %q: obsługiwane formaty to CBZ, PDF, EPUB lub katalog obrazów",
	CLIFlagGraphicsUsage:            "renderer: auto, ascii, dots, kitty, sixel lub iterm2",
	CLIFlagVersionUsage:             "wypisz wersję i zakończ",
	CLIFlagUpdateUsage:              "sprawdź aktualizacje i zakończ",
	CLIFlagEnvUsage:                 "wypisz środowisko comicread i zakończ",
	CLIFlagClearJournalUsage:        "usuń lokalny dziennik dla pliku lub katalogu i zakończ",
	CLIFlagBookViewUsage:            "pokaż pary stron od lewej do prawej",
	CLIFlagRightBookViewUsage:       "pokaż pary stron od prawej do lewej",
	CLIFlagCircleBookViewUsage:      "pokaż nakładające się pary stron od lewej do prawej",
	CLIFlagRightCircleBookViewUsage: "pokaż nakładające się pary stron od prawej do lewej",
	CLIErrMultipleBookViews:         "można użyć tylko jednej opcji widoku książki",
	CLIErrInvalidView:               "nieobsługiwana wartość COMICREAD_VIEW %q (dozwolone: book-view, right-view, circle-view lub right-circle-view)",
	CLIFlagOpenUsage:                "katalog do otwarcia w wyborze plików (domyślnie: COMICREAD_DIR lub bieżący katalog)",
	CLIErrOpenNotDir:                "otwórz katalog %q: nie jest katalogiem",
	CLIHelpHint:                     "uruchom 'comicread --help', aby uzyskać pomoc",
	CLIUsage:                        "użycie: comicread [opcje] [plik]",
	CLIUsageFull: `comicread — minimalny czytnik mangi dla terminala

użycie: comicread [opcje] [plik]

opcje:
  --graphics string   renderer: auto, ascii, dots, kitty, sixel lub iterm2 (domyślnie "auto")
  --book-view         pokaż pary stron od lewej do prawej
  --right-view        pokaż pary stron od prawej do lewej
  --circle-view       pokaż nakładające się pary stron od lewej do prawej
  --right-circle-view
                      pokaż nakładające się pary stron od prawej do lewej
  --clear-journal    usuń lokalny dziennik dla pliku lub katalogu i zakończ
  -o, --open string   katalog do otwarcia w wyborze plików (domyślnie: COMICREAD_DIR lub bieżący katalog)
  --env               wypisz środowisko comicread i zakończ
  --update            sprawdź aktualizacje i zakończ
  -v, --version       wypisz wersję i zakończ
  -h, --help          pokaż tę pomoc

Jeśli nie podano pliku ani katalogu, interaktywny wybór pliku otworzy się w COMICREAD_DIR
(jeśli ustawiono prawidłowy katalog) lub w bieżącym katalogu.

zmienne środowiskowe:
  COMICREAD_GRAPHICS  domyślny renderer: auto, ascii, dots, kitty, sixel lub iterm2
  COMICREAD_PRERENDERED_NEXT      kolejne strony do wstępnego renderowania (domyślnie 1)
  COMICREAD_PRERENDERED_PREVIOUS  poprzednie strony do wstępnego renderowania (domyślnie 1)
  COMICREAD_VIEW      domyślny widok: book-view, right-view, circle-view lub right-circle-view
  COMICREAD_LANG      język komunikatów: en, uk, pl, de, fr, es, cs, ro, it, ko, ja, id, hi, el, tr, kk lub ka (domyślnie "en")
  COMICREAD_DIR       domyślny katalog dla wyboru plików, gdy nie podano ścieżki`,
}
