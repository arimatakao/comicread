package i18n

var daMessages = map[string]string{
	CLIFlagConfigUsage: "konfigurationsfil der skal bruges", CLIFlagResetConfigUsage: "nulstil config.toml og afslut", CLIFlagSetConfigUsage: "opdater config.toml: nøgle=værdi",
	ReaderStatusWaitingTerminalSize: "venter på terminalstørrelse", ReaderStatusTerminalTooSmall: "terminalvinduet er for lille", ReaderStatusLastPage: "sidste side", ReaderStatusFirstPage: "første side", ReaderStatusRenderError: "renderingsfejl: %v", ReaderStatusMaximumZoom: "maksimal zoom", ReaderStatusMinimumZoom: "minimal zoom", ReaderStatusInvalidPage: "ugyldigt sidetal",
	ReaderViewTerminalTooSmall: "comicread: terminalvinduet er for lille", ReaderViewWindowTitle: "comicread — %s", ReaderViewPages: "sider %d/%d", ReaderViewPageRange: "sider %d–%d/%d", ReaderViewRendering: "renderer", ReaderViewBookmarks: "Bogmærker", ReaderViewNoBookmarks: "(ingen bogmærker)", ReaderViewBookmarksHelp: "op/ned flyt | enter åbn | esc luk", ReaderViewGoToPage: "Gå til side: %s",
	ReaderViewHelp: `Taster

← →  forrige / næste side
↑ ↓  rul en zoomet side
+ -  zoom ind / ud
b    tilføj / fjern bogmærke
v ← → forrige / næste bogmærke
c v  bogmærker
g123 enter  gå til side
q    afslut

?    luk hjælpen`,
	FilepickerHeader: "comicread — vælg et kapitel\n%s\n\n", FilepickerNoEntries: "  (ingen understøttede poster)\n", FilepickerHelp: "\n↑/↓ flyt\n← overordnet mappe\n→ gå ind i mappe\nenter åbn fil\ns vælg fremhævet mappe\nf tilføj / fjern den aktuelle mappe fra favoritter\nF tilføj favoritmappe\nb favoritmapper\no gå til mappe\nq afslut\n", FilepickerWindowTitle: "comicread — vælg en fil", FilepickerGoToPrompt: "\nGå til mappe: %s\n", FilepickerFavoritePrompt: "\nFavoritmappe: %s\n", FilepickerGoToErr: "  fejl: %s\n", FilepickerFavorites: "Favoritmapper\n\n", FilepickerNoFavorites: "  (ingen favoritmapper er konfigureret)\n", FilepickerFavoritesHelp: "\n↑/↓ flyt\nenter gå til mappe\nd fjern favorit\nesc tilbage\n", FilepickerFavoriteErr: "  fejl ved lagring af favoritter: %s\n", FilepickerErrResolveDir: "find mappe %q: %w", FilepickerErrReadDir: "læs mappe %q: %w", FilepickerErrRunPicker: "kør filvælger: %w", FilepickerErrEmptyPath: "stien er tom", FilepickerErrNotDir: "%q er ikke en mappe",
	LoadingViewOpening: "åbner %s…", LoadingViewWindowTitle: "comicread — åbner",
	CLIErrGetWorkingDir: "hent arbejdsmappe: %w", CLIErrPickFile: "vælg fil: %w", CLIErrRunTUI: "kør TUI: %w", CLIErrParseArgs: "fortolk argumenter: %w", CLIErrOpenChapter: "åbn kapitel: %w", CLIErrOpenJournal: "åbn journal: %w", CLIErrClearJournal: "ryd journal: %w", CLIErrClearJournalRequiresInput: "--clear-journal kræver en fil eller mappe", CLIErrNoPages: "kapitlet indeholder ingen læsbare billedsider", CLIErrInspectInput: "undersøg input %q: %w", CLIErrUnsupportedFile: "filen %q understøttes ikke: understøttede formater er CBZ, PDF, EPUB eller en billedmappe",
	CLIFlagGraphicsUsage: "renderer: auto, ascii, dots, kitty, sixel eller iterm2", CLIFlagVersionUsage: "udskriv version og afslut", CLIFlagUpdateUsage: "søg efter opdateringer og afslut", CLIFlagEnvUsage: "udskriv comicread-miljøet og afslut", CLIFlagClearJournalUsage: "fjern lokal journal for en fil eller mappe og afslut", CLIFlagBookViewUsage: "vis sidepar fra venstre mod højre", CLIFlagRightBookViewUsage: "vis sidepar fra højre mod venstre", CLIFlagCircleBookViewUsage: "vis overlappende sidepar fra venstre mod højre", CLIFlagRightCircleBookViewUsage: "vis overlappende sidepar fra højre mod venstre", CLIErrMultipleBookViews: "kun én bogvisning kan bruges", CLIErrInvalidView: "COMICREAD_VIEW %q understøttes ikke (forventet book-view, right-view, circle-view eller right-circle-view)", CLIFlagOpenUsage: "mappe, der skal åbnes i filvælgeren (standard: COMICREAD_DIR eller den aktuelle mappe)", CLIErrOpenNotDir: "åbn mappe %q: ikke en mappe", CLIFlagWebUsage: "start en lokal web-læser i stedet for terminalgrænsefladen", CLIErrWebArgs: "--web accepterer ikke et fil- eller mappeargument", WebServerStarted: "comicread-web-læseren kører på %s (tryk Ctrl+C for at stoppe)", WebErrListen: "start webserver: %w", WebErrServe: "kør webserver: %w", CLIHelpHint: "kør 'comicread --help' for brug", CLIUsage: "brug: comicread [indstillinger] [fil]",
	CLIUsageFull: `comicread — en minimal mangalæser til terminalen

brug: comicread [indstillinger] [fil]

indstillinger:
  --config string     konfigurationsfil der skal bruges
  --graphics string   renderer: auto, ascii, dots, kitty, sixel eller iterm2 (standard "auto")
  --book-view         vis sidepar fra venstre mod højre
  --right-view        vis sidepar fra højre mod venstre
  --circle-view       vis overlappende sidepar fra venstre mod højre
  --right-circle-view
                      vis overlappende sidepar fra højre mod venstre
  --clear-journal     fjern lokal journal for en fil eller mappe og afslut
  --reset-config     nulstil config.toml og afslut
  --set-config value opdater config.toml: nøgle=værdi
  -o, --open string   mappe, der skal åbnes i filvælgeren (standard: COMICREAD_DIR eller den aktuelle mappe)
  --update            søg efter opdateringer og afslut
  --web               start en lokal web-læser i stedet for terminalgrænsefladen
  -v, --version       udskriv version og afslut
  -h, --help          vis denne hjælp

Hvis ingen fil eller mappe er angivet, åbnes en interaktiv filvælger i COMICREAD_DIR
(hvis den er sat til en gyldig mappe) eller ellers i den aktuelle mappe.`,
	ReaderViewMetadata: "Metadata",
}
