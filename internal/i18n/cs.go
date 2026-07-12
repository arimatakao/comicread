package i18n

var csMessages = map[string]string{
	ReaderStatusWaitingTerminalSize: "čekání na velikost terminálu",
	ReaderStatusTerminalTooSmall:    "okno terminálu je příliš malé",
	ReaderStatusLastPage:            "poslední stránka",
	ReaderStatusFirstPage:           "první stránka",
	ReaderStatusRenderError:         "chyba vykreslení: %v",
	ReaderStatusMaximumZoom:         "maximální přiblížení",
	ReaderStatusMinimumZoom:         "minimální přiblížení",

	ReaderViewTerminalTooSmall: "comicread: okno terminálu je příliš malé",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "strany %d/%d",
	ReaderViewPageRange:        "strany %d-%d/%d",
	ReaderViewRendering:        "vykreslování",
	ReaderViewHelp: `Klávesy

← →  předchozí / další strana
↑ ↓  posun přiblížené strany
+ -  přiblížit / oddálit
q    ukončit

?    zavřít nápovědu`,

	FilepickerHeader:      "comicread — vyberte kapitolu\n%s\n\n",
	FilepickerNoEntries:   "  (žádné podporované položky)\n",
	FilepickerHelp:        "\n↑/↓ pohyb  |  ← nadřazená složka  |  → vstoupit do složky  |  enter otevřít soubor  |  s vybrat zvýrazněnou složku  |  q ukončit\n",
	FilepickerWindowTitle: "comicread — vybrat soubor",

	FilepickerErrResolveDir: "nelze zjistit složku %q: %w",
	FilepickerErrReadDir:    "nelze přečíst složku %q: %w",
	FilepickerErrRunPicker:  "chyba při spuštění výběru souboru: %w",

	LoadingViewOpening:     "otevírání %s…",
	LoadingViewWindowTitle: "comicread — otevírání",

	CLIErrGetWorkingDir:             "nelze získat pracovní složku: %w",
	CLIErrPickFile:                  "nelze vybrat soubor: %w",
	CLIErrRunTUI:                    "chyba při spuštění TUI: %w",
	CLIErrParseArgs:                 "chyba při analýze argumentů: %w",
	CLIErrOpenChapter:               "nelze otevřít kapitolu: %w",
	CLIErrNoPages:                   "kapitola neobsahuje žádné čitelné obrazové stránky",
	CLIErrInspectInput:              "nelze zkontrolovat vstup %q: %w",
	CLIErrUnsupportedFile:           "nepodporovaný soubor %q: podporované formáty jsou CBZ, PDF, EPUB nebo složka s obrázky",
	CLIFlagGraphicsUsage:            "vykreslovač: auto, ascii, dots, kitty, sixel nebo iterm2",
	CLIFlagVersionUsage:             "vypsat verzi a ukončit",
	CLIFlagEnvUsage:                 "vypsat prostředí comicread a ukončit",
	CLIFlagBookViewUsage:            "zobrazit dvojice stránek zleva doprava",
	CLIFlagRightBookViewUsage:       "zobrazit dvojice stránek zprava doleva",
	CLIFlagCircleBookViewUsage:      "zobrazit překrývající se dvojice stránek zleva doprava",
	CLIFlagRightCircleBookViewUsage: "zobrazit překrývající se dvojice stránek zprava doleva",
	CLIErrMultipleBookViews:         "lze použít pouze jednu možnost knižního zobrazení",
	CLIErrInvalidView:               "nepodporovaná hodnota COMICREAD_VIEW %q (očekává se: book-view, right-view, circle-view nebo right-circle-view)",
	CLIHelpHint:                     "nápovědu zobrazíte příkazem 'comicread --help'",
	CLIUsage:                        "použití: comicread [volby] [soubor]",
	CLIUsageFull: `comicread — minimalistická čtečka mangy pro terminál

použití: comicread [volby] [soubor]

volby:
  --graphics string   vykreslovač: auto, ascii, dots, kitty, sixel nebo iterm2 (výchozí "auto")
  --book-view         zobrazit dvojice stránek zleva doprava
  --right-view        zobrazit dvojice stránek zprava doleva
  --circle-view       zobrazit překrývající se dvojice stránek zleva doprava
  --right-circle-view
                      zobrazit překrývající se dvojice stránek zprava doleva
  --env               vypsat prostředí comicread a ukončit
  -v, --version       vypsat verzi a ukončit
  -h, --help          zobrazit tuto nápovědu

Pokud není zadán soubor ani složka, otevře se v aktuální složce interaktivní výběr souboru.

prostředí:
  COMICREAD_GRAPHICS  výchozí vykreslovač: auto, ascii, dots, kitty, sixel nebo iterm2
  COMICREAD_VIEW      výchozí zobrazení: book-view, right-view, circle-view nebo right-circle-view
  COMICREAD_LANG      jazyk zpráv: en, uk, pl, de, fr, es, cs, ro, it, ko, ja, id, hi, el nebo tr (výchozí "en")`,
}
