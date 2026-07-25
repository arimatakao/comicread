package i18n

var huMessages = map[string]string{
	CLIFlagConfigUsage: "használandó konfigurációs fájl", CLIFlagResetConfigUsage: "a config.toml alaphelyzetbe állítása és kilépés", CLIFlagSetConfigUsage: "config.toml frissítése: kulcs=érték",
	ReaderStatusWaitingTerminalSize: "várakozás a terminál méretére",
	ReaderStatusTerminalTooSmall:    "a terminálablak túl kicsi",
	ReaderStatusLastPage:            "utolsó oldal",
	ReaderStatusFirstPage:           "első oldal",
	ReaderStatusRenderError:         "megjelenítési hiba: %v",
	ReaderStatusMaximumZoom:         "maximális nagyítás",
	ReaderStatusMinimumZoom:         "minimális nagyítás",
	ReaderStatusInvalidPage:         "érvénytelen oldalszám",

	ReaderViewTerminalTooSmall: "comicread: a terminálablak túl kicsi",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "oldalak %d/%d",
	ReaderViewPageRange:        "oldalak %d–%d/%d",
	ReaderViewRendering:        "megjelenítés",
	ReaderViewBookmarks:        "Könyvjelzők",
	ReaderViewNoBookmarks:      "(nincsenek könyvjelzők)",
	ReaderViewBookmarksHelp:    "fel/le lépés | enter megnyitás | esc bezárás",
	ReaderViewGoToPage:         "Ugrás oldalra: %s",
	ReaderViewHelp: `Billentyűk

← →  előző / következő oldal
↑ ↓  nagyított oldal görgetése
+ -  nagyítás / kicsinyítés
b    könyvjelző hozzáadása / eltávolítása
v ← → előző / következő könyvjelző
c v  könyvjelzők
g123 enter  ugrás oldalra
q    kilépés

?    súgó bezárása`,

	FilepickerHeader:         "comicread — fejezet kiválasztása\n%s\n\n",
	FilepickerNoEntries:      "  (nincsenek támogatott elemek)\n",
	FilepickerHelp:           "\n↑/↓ lépés\n← szülőkönyvtár\n→ belépés a könyvtárba\nenter fájl megnyitása\ns kijelölt könyvtár kiválasztása\nf jelenlegi könyvtár hozzáadása / eltávolítása a kedvencekből\nF kedvenc könyvtár hozzáadása\nb kedvenc könyvtárak\no ugrás könyvtárhoz\nq kilépés\n",
	FilepickerWindowTitle:    "comicread — fájl kiválasztása",
	FilepickerGoToPrompt:     "\nUgrás könyvtárhoz: %s\n",
	FilepickerFavoritePrompt: "\nKedvenc könyvtár: %s\n",
	FilepickerFavorites:      "Kedvenc könyvtárak\n\n",
	FilepickerNoFavorites:    "  (nincsenek beállított kedvenc könyvtárak)\n",
	FilepickerFavoritesHelp:  "\n↑/↓ lépés\nenter ugrás könyvtárhoz\nd kedvenc eltávolítása\nesc vissza\n",
	FilepickerFavoriteErr:    "  hiba a kedvencek mentésekor: %s\n",
	FilepickerGoToErr:        "  hiba: %s\n",

	FilepickerErrResolveDir: "könyvtár feloldása %q: %w",
	FilepickerErrReadDir:    "könyvtár olvasása %q: %w",
	FilepickerErrRunPicker:  "fájlválasztó futtatása: %w",
	FilepickerErrEmptyPath:  "az útvonal üres",
	FilepickerErrNotDir:     "%q nem könyvtár",

	LoadingViewOpening:     "%s megnyitása…",
	LoadingViewWindowTitle: "comicread — megnyitás",

	CLIErrGetWorkingDir:             "munkakönyvtár lekérése: %w",
	CLIErrPickFile:                  "fájl kiválasztása: %w",
	CLIErrRunTUI:                    "TUI futtatása: %w",
	CLIErrParseArgs:                 "argumentumok feldolgozása: %w",
	CLIErrOpenChapter:               "fejezet megnyitása: %w",
	CLIErrOpenJournal:               "napló megnyitása: %w",
	CLIErrClearJournal:              "napló törlése: %w",
	CLIErrClearJournalRequiresInput: "a --clear-journal fájlt vagy könyvtárat igényel",
	CLIErrNoPages:                   "a fejezet nem tartalmaz olvasható képes oldalakat",
	CLIErrInspectInput:              "bemenet vizsgálata %q: %w",
	CLIErrUnsupportedFile:           "nem támogatott fájl: %q; a támogatott formátumok: CBZ, PDF, EPUB vagy képkönyvtár",
	CLIFlagGraphicsUsage:            "megjelenítő: auto, ascii, dots, kitty, sixel vagy iterm2",
	CLIFlagVersionUsage:             "verzió kiírása és kilépés",
	CLIFlagUpdateUsage:              "frissítések keresése és kilépés",
	CLIFlagEnvUsage:                 "comicread környezet kiírása és kilépés",
	CLIFlagClearJournalUsage:        "fájl vagy könyvtár helyi naplójának törlése és kilépés",
	CLIFlagBookViewUsage:            "oldalpárok megjelenítése balról jobbra",
	CLIFlagRightBookViewUsage:       "oldalpárok megjelenítése jobbról balra",
	CLIFlagCircleBookViewUsage:      "átfedő oldalpárok megjelenítése balról jobbra",
	CLIFlagRightCircleBookViewUsage: "átfedő oldalpárok megjelenítése jobbról balra",
	CLIErrMultipleBookViews:         "csak egy könyvnézet-beállítás használható",
	CLIErrInvalidView:               "nem támogatott COMICREAD_VIEW %q (elvárt: book-view, right-view, circle-view vagy right-circle-view)",
	CLIFlagOpenUsage:                "a fájlválasztóban megnyitandó könyvtár (alapértelmezés: COMICREAD_DIR vagy a jelenlegi könyvtár)",
	CLIErrOpenNotDir:                "könyvtár megnyitása %q: nem könyvtár",
	CLIHelpHint:                     "használathoz futtasd: 'comicread --help'",
	CLIUsage:                        "használat: comicread [beállítások] [fájl]",
	CLIUsageFull: `comicread — minimalista terminálos mangaolvasó

használat: comicread [beállítások] [fájl]

beállítások:
  --graphics string   megjelenítő: auto, ascii, dots, kitty, sixel vagy iterm2 (alapértelmezés: "auto")
  --book-view         oldalpárok megjelenítése balról jobbra
  --right-view        oldalpárok megjelenítése jobbról balra
  --circle-view       átfedő oldalpárok megjelenítése balról jobbra
  --right-circle-view
                      átfedő oldalpárok megjelenítése jobbról balra
  --clear-journal     fájl vagy könyvtár helyi naplójának törlése és kilépés
  -o, --open string   a fájlválasztóban megnyitandó könyvtár (alapértelmezés: COMICREAD_DIR vagy a jelenlegi könyvtár)
  --env               comicread környezet kiírása és kilépés
  --update            frissítések keresése és kilépés
  -v, --version       verzió kiírása és kilépés
  -h, --help          súgó megjelenítése

Ha nincs megadva fájl vagy könyvtár, egy interaktív fájlválasztó nyílik meg a COMICREAD_DIR
(ha érvényes könyvtárra van beállítva) vagy egyébként a jelenlegi könyvtárban.

környezet:
  COMICREAD_GRAPHICS  alapértelmezett megjelenítő: auto, ascii, dots, kitty, sixel vagy iterm2
  COMICREAD_PRERENDERED_NEXT      előre megjelenítendő következő oldalak (alapértelmezés: 1)
  COMICREAD_PRERENDERED_PREVIOUS  előre megjelenítendő előző oldalak (alapértelmezés: 1)
  COMICREAD_VIEW      alapértelmezett nézet: book-view, right-view, circle-view vagy right-circle-view
  COMICREAD_LANG      üzenetek nyelve: https://github.com/arimatakao/comicread#environment-variables (alapértelmezés: "en")
  COMICREAD_DIR       alapértelmezett fájlválasztó-könyvtár, ha nincs útvonal megadva`,
}
