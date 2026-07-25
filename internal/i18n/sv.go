package i18n

var svMessages = map[string]string{
	CLIFlagConfigUsage: "konfigurationsfil att använda", CLIFlagResetConfigUsage: "återställ config.toml och avsluta", CLIFlagSetConfigUsage: "uppdatera config.toml: nyckel=värde",
	ReaderStatusWaitingTerminalSize: "väntar på terminalstorlek",
	ReaderStatusTerminalTooSmall:    "terminalfönstret är för litet",
	ReaderStatusLastPage:            "sista sidan",
	ReaderStatusFirstPage:           "första sidan",
	ReaderStatusRenderError:         "renderingsfel: %v",
	ReaderStatusMaximumZoom:         "maximal zoom",
	ReaderStatusMinimumZoom:         "minimal zoom",
	ReaderStatusInvalidPage:         "ogiltigt sidnummer",

	ReaderViewTerminalTooSmall: "comicread: terminalfönstret är för litet",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "sidor %d/%d",
	ReaderViewPageRange:        "sidor %d–%d/%d",
	ReaderViewRendering:        "renderar",
	ReaderViewBookmarks:        "Bokmärken",
	ReaderViewNoBookmarks:      "(inga bokmärken)",
	ReaderViewBookmarksHelp:    "upp/ned flytta | enter öppna | esc stäng",
	ReaderViewGoToPage:         "Gå till sida: %s",
	ReaderViewHelp: `Tangenter

← →  föregående / nästa sida
↑ ↓  rulla en inzoomad sida
+ -  zooma in / ut
b    lägg till / ta bort bokmärke
v ← → föregående / nästa bokmärke
c v  bokmärken
g123 enter  gå till sida
q    avsluta

?    stäng hjälpen`,

	FilepickerHeader:         "comicread — välj ett kapitel\n%s\n\n",
	FilepickerNoEntries:      "  (inga poster som stöds)\n",
	FilepickerHelp:           "\n↑/↓ flytta\n← överordnad katalog\n→ öppna katalog\nenter öppna fil\ns välj markerad katalog\nf lägg till / ta bort aktuell katalog från favoriter\nF lägg till favoritkatalog\nb favoritkataloger\no gå till katalog\nq avsluta\n",
	FilepickerWindowTitle:    "comicread — välj en fil",
	FilepickerGoToPrompt:     "\nGå till katalog: %s\n",
	FilepickerFavoritePrompt: "\nFavoritkatalog: %s\n",
	FilepickerFavorites:      "Favoritkataloger\n\n",
	FilepickerNoFavorites:    "  (inga favoritkataloger har konfigurerats)\n",
	FilepickerFavoritesHelp:  "\n↑/↓ flytta\nenter gå till katalog\nd ta bort favorit\nesc tillbaka\n",
	FilepickerFavoriteErr:    "  fel vid sparande av favoriter: %s\n",
	FilepickerGoToErr:        "  fel: %s\n",
	FilepickerErrResolveDir:  "lös katalog %q: %w", FilepickerErrReadDir: "läs katalog %q: %w", FilepickerErrRunPicker: "kör filväljare: %w", FilepickerErrEmptyPath: "sökvägen är tom", FilepickerErrNotDir: "%q är inte en katalog",
	LoadingViewOpening: "%s öppnas…", LoadingViewWindowTitle: "comicread — öppnar",
	CLIErrGetWorkingDir: "hämta arbetskatalog: %w", CLIErrPickFile: "välj fil: %w", CLIErrRunTUI: "kör TUI: %w", CLIErrParseArgs: "tolka argument: %w", CLIErrOpenChapter: "öppna kapitel: %w", CLIErrOpenJournal: "öppna journal: %w", CLIErrClearJournal: "rensa journal: %w", CLIErrClearJournalRequiresInput: "--clear-journal kräver en fil eller katalog", CLIErrNoPages: "kapitlet innehåller inga läsbara bildsidor", CLIErrInspectInput: "undersök indata %q: %w", CLIErrUnsupportedFile: "filen %q stöds inte: formaten CBZ, PDF, EPUB eller en bildkatalog stöds",
	CLIFlagGraphicsUsage: "renderare: auto, ascii, dots, kitty, sixel eller iterm2", CLIFlagVersionUsage: "skriv ut version och avsluta", CLIFlagUpdateUsage: "sök efter uppdateringar och avsluta", CLIFlagEnvUsage: "skriv ut comicreads miljö och avsluta", CLIFlagClearJournalUsage: "ta bort lokal journal för en fil eller katalog och avsluta", CLIFlagBookViewUsage: "visa sidpar från vänster till höger", CLIFlagRightBookViewUsage: "visa sidpar från höger till vänster", CLIFlagCircleBookViewUsage: "visa överlappande sidpar från vänster till höger", CLIFlagRightCircleBookViewUsage: "visa överlappande sidpar från höger till vänster", CLIErrMultipleBookViews: "endast ett bokvisningsalternativ får användas", CLIErrInvalidView: "COMICREAD_VIEW %q stöds inte (förväntar book-view, right-view, circle-view eller right-circle-view)", CLIFlagOpenUsage: "katalog att öppna i filväljaren (standard: COMICREAD_DIR eller aktuell katalog)", CLIErrOpenNotDir: "öppna katalog %q: inte en katalog", CLIHelpHint: "kör 'comicread --help' för användning", CLIUsage: "användning: comicread [alternativ] [fil]",
	CLIUsageFull: `comicread — en minimal mangaläsare för terminalen

användning: comicread [alternativ] [fil]

alternativ:
  --graphics string   renderare: auto, ascii, dots, kitty, sixel eller iterm2 (standard "auto")
  --book-view         visa sidpar från vänster till höger
  --right-view        visa sidpar från höger till vänster
  --circle-view       visa överlappande sidpar från vänster till höger
  --right-circle-view
                      visa överlappande sidpar från höger till vänster
  --clear-journal     ta bort lokal journal för en fil eller katalog och avsluta
  -o, --open string   katalog att öppna i filväljaren (standard: COMICREAD_DIR eller aktuell katalog)
  --env               skriv ut comicreads miljö och avsluta
  --update            sök efter uppdateringar och avsluta
  -v, --version       skriv ut version och avsluta
  -h, --help          visa detta hjälpmeddelande

Om ingen fil eller katalog anges öppnas en interaktiv filväljare i COMICREAD_DIR
(om den är en giltig katalog) eller annars i den aktuella katalogen.

miljö:
  COMICREAD_GRAPHICS  standardrenderare: auto, ascii, dots, kitty, sixel eller iterm2
  COMICREAD_PRERENDERED_NEXT      nästa sidor att förrendera (standard 1)
  COMICREAD_PRERENDERED_PREVIOUS  föregående sidor att förrendera (standard 1)
  COMICREAD_VIEW      standardvy: book-view, right-view, circle-view eller right-circle-view
  COMICREAD_LANG      meddelandespråk: https://github.com/arimatakao/comicread#environment-variables (standard "en")
  COMICREAD_DIR       standardkatalog för filväljaren när ingen sökväg anges`,
}
