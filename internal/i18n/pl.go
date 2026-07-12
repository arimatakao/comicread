package i18n

var plMessages = map[string]string{
	ReaderStatusWaitingTerminalSize: "oczekiwanie na rozmiar terminala",
	ReaderStatusTerminalTooSmall:    "okno terminala jest za małe",
	ReaderStatusLastPage:            "ostatnia strona",
	ReaderStatusFirstPage:           "pierwsza strona",
	ReaderStatusRenderError:         "błąd renderowania: %v",
	ReaderStatusMaximumZoom:         "maksymalne powiększenie",
	ReaderStatusMinimumZoom:         "minimalne powiększenie",

	ReaderViewTerminalTooSmall: "comicread: okno terminala jest za małe",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "strony %d/%d",
	ReaderViewPageRange:        "strony %d-%d/%d",
	ReaderViewRendering:        "renderowanie",
	ReaderViewHelp: `Klawisze

← →  poprzednia / następna strona
↑ ↓  przewijanie powiększonej strony
+ -  powiększ / pomniejsz
q    wyjście

?    zamknij pomoc`,

	FilepickerHeader:      "comicread — wybierz rozdział\n%s\n\n",
	FilepickerNoEntries:   "  (brak obsługiwanych elementów)\n",
	FilepickerHelp:        "\n↑/↓ ruch  |  ← katalog nadrzędny  |  → wejdź do katalogu  |  enter otwórz plik  |  s wybierz zaznaczony katalog  |  q wyjście\n",
	FilepickerWindowTitle: "comicread — wybór pliku",

	FilepickerErrResolveDir: "nie można ustalić katalogu %q: %w",
	FilepickerErrReadDir:    "nie można odczytać katalogu %q: %w",
	FilepickerErrRunPicker:  "błąd uruchamiania wyboru pliku: %w",

	LoadingViewOpening:     "otwieranie %s…",
	LoadingViewWindowTitle: "comicread — otwieranie",

	CLIErrGetWorkingDir:             "nie można uzyskać katalogu roboczego: %w",
	CLIErrPickFile:                  "nie można wybrać pliku: %w",
	CLIErrRunTUI:                    "błąd uruchamiania TUI: %w",
	CLIErrParseArgs:                 "błąd analizy argumentów: %w",
	CLIErrOpenChapter:               "nie można otworzyć rozdziału: %w",
	CLIErrNoPages:                   "rozdział nie zawiera czytelnych stron obrazów",
	CLIErrInspectInput:              "nie można sprawdzić danych wejściowych %q: %w",
	CLIErrUnsupportedFile:           "nieobsługiwany plik %q: obsługiwane formaty to CBZ, PDF, EPUB lub katalog obrazów",
	CLIFlagGraphicsUsage:            "renderer: auto, ascii, dots, kitty, sixel lub iterm2",
	CLIFlagVersionUsage:             "wypisz wersję i zakończ",
	CLIFlagEnvUsage:                 "wypisz środowisko comicread i zakończ",
	CLIFlagBookViewUsage:            "pokaż pary stron od lewej do prawej",
	CLIFlagRightBookViewUsage:       "pokaż pary stron od prawej do lewej",
	CLIFlagCircleBookViewUsage:      "pokaż nakładające się pary stron od lewej do prawej",
	CLIFlagRightCircleBookViewUsage: "pokaż nakładające się pary stron od prawej do lewej",
	CLIErrMultipleBookViews:         "można użyć tylko jednej opcji widoku książki",
	CLIErrInvalidView:               "nieobsługiwana wartość COMICREAD_VIEW %q (dozwolone: book-view, right-view, circle-view lub right-circle-view)",
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
  --env               wypisz środowisko comicread i zakończ
  -v, --version       wypisz wersję i zakończ
  -h, --help          pokaż tę pomoc

Jeśli nie podano pliku ani katalogu, w bieżącym katalogu otworzy się interaktywny wybór pliku.

zmienne środowiskowe:
  COMICREAD_GRAPHICS  domyślny renderer: auto, ascii, dots, kitty, sixel lub iterm2
  COMICREAD_VIEW      domyślny widok: book-view, right-view, circle-view lub right-circle-view
  COMICREAD_LANG      język komunikatów: en, uk, pl, de, fr, es, cs, ro, it, ko, ja, id, hi, el lub tr (domyślnie "en")`,
}
