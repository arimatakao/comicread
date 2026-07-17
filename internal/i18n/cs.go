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
	ReaderViewBookmarks:        "Záložky",
	ReaderViewNoBookmarks:      "(žádné záložky)",
	ReaderViewBookmarksHelp:    "↑/↓ pohyb | enter otevřít | esc zavřít",
	ReaderViewHelp: `Klávesy

← →  předchozí / další strana
↑ ↓  posun přiblížené strany
+ -  přiblížit / oddálit
b    přidat / odstranit záložku
v ← → předchozí / další záložka
c v  záložky
q    ukončit

?    zavřít nápovědu`,

	FilepickerHeader:      "comicread — vyberte kapitolu\n%s\n\n",
	FilepickerNoEntries:   "  (žádné podporované položky)\n",
	FilepickerHelp:        "\n↑/↓ pohyb\n← nadřazená složka\n→ vstoupit do složky\nenter otevřít soubor\ns vybrat zvýrazněnou složku\no přejít do adresáře\nq ukončit\n",
	FilepickerWindowTitle: "comicread — vybrat soubor",
	FilepickerGoToPrompt:  "\nPřejít do adresáře: %s\n",
	FilepickerGoToErr:     "  chyba: %s\n",

	FilepickerErrResolveDir: "nelze zjistit složku %q: %w",
	FilepickerErrReadDir:    "nelze přečíst složku %q: %w",
	FilepickerErrRunPicker:  "chyba při spuštění výběru souboru: %w",
	FilepickerErrEmptyPath:  "cesta je prázdná",
	FilepickerErrNotDir:     "%q není adresář",

	LoadingViewOpening:     "otevírání %s…",
	LoadingViewWindowTitle: "comicread — otevírání",

	CLIErrGetWorkingDir:             "nelze získat pracovní složku: %w",
	CLIErrPickFile:                  "nelze vybrat soubor: %w",
	CLIErrRunTUI:                    "chyba při spuštění TUI: %w",
	CLIErrParseArgs:                 "chyba při analýze argumentů: %w",
	CLIErrOpenChapter:               "nelze otevřít kapitolu: %w",
	CLIErrOpenJournal:               "nelze otevřít deník: %w",
	CLIErrClearJournal:              "nelze vymazat deník: %w",
	CLIErrClearJournalRequiresInput: "--clear-journal vyžaduje soubor nebo složku",
	CLIErrNoPages:                   "kapitola neobsahuje žádné čitelné obrazové stránky",
	CLIErrInspectInput:              "nelze zkontrolovat vstup %q: %w",
	CLIErrUnsupportedFile:           "nepodporovaný soubor %q: podporované formáty jsou CBZ, PDF, EPUB nebo složka s obrázky",
	CLIFlagGraphicsUsage:            "vykreslovač: auto, ascii, dots, kitty, sixel nebo iterm2",
	CLIFlagVersionUsage:             "vypsat verzi a ukončit",
	CLIFlagUpdateUsage:              "zkontrolovat aktualizace a ukončit",
	CLIFlagEnvUsage:                 "vypsat prostředí comicread a ukončit",
	CLIFlagClearJournalUsage:        "odstranit místní deník pro soubor nebo složku a ukončit",
	CLIFlagBookViewUsage:            "zobrazit dvojice stránek zleva doprava",
	CLIFlagRightBookViewUsage:       "zobrazit dvojice stránek zprava doleva",
	CLIFlagCircleBookViewUsage:      "zobrazit překrývající se dvojice stránek zleva doprava",
	CLIFlagRightCircleBookViewUsage: "zobrazit překrývající se dvojice stránek zprava doleva",
	CLIErrMultipleBookViews:         "lze použít pouze jednu možnost knižního zobrazení",
	CLIErrInvalidView:               "nepodporovaná hodnota COMICREAD_VIEW %q (očekává se: book-view, right-view, circle-view nebo right-circle-view)",
	CLIFlagOpenUsage:                "adresář k otevření ve výběru souborů (výchozí: COMICREAD_DIR nebo aktuální adresář)",
	CLIErrOpenNotDir:                "otevřít adresář %q: není adresář",
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
  --clear-journal    odstranit místní deník pro soubor nebo složku a ukončit
  -o, --open string   adresář k otevření ve výběru souborů (výchozí: COMICREAD_DIR nebo aktuální adresář)
  --env               vypsat prostředí comicread a ukončit
  --update            zkontrolovat aktualizace a ukončit
  -v, --version       vypsat verzi a ukončit
  -h, --help          zobrazit tuto nápovědu

Pokud není zadán soubor ani složka, otevře se interaktivní výběr souboru v COMICREAD_DIR
(je-li nastavena na platný adresář), jinak v aktuálním adresáři.

prostředí:
  COMICREAD_GRAPHICS  výchozí vykreslovač: auto, ascii, dots, kitty, sixel nebo iterm2
  COMICREAD_PRERENDERED_NEXT      počet dalších stránek k předvykreslení (výchozí 1)
  COMICREAD_PRERENDERED_PREVIOUS  počet předchozích stránek k předvykreslení (výchozí 1)
  COMICREAD_VIEW      výchozí zobrazení: book-view, right-view, circle-view nebo right-circle-view
  COMICREAD_LANG      jazyk zpráv: en, uk, pl, de, fr, es, cs, ro, it, ko, ja, id, hi, el, tr, kk nebo ka (výchozí "en")
  COMICREAD_DIR       výchozí adresář pro výběr souborů, když není zadána cesta`,
}
