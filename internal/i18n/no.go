package i18n

var noMessages = map[string]string{
	CLIFlagConfigUsage: "konfigurasjonsfil som skal brukes", CLIFlagResetConfigUsage: "tilbakestill config.toml og avslutt", CLIFlagSetConfigUsage: "oppdater config.toml: nøkkel=verdi",
	ReaderStatusWaitingTerminalSize: "venter på terminalstørrelse", ReaderStatusTerminalTooSmall: "terminalvinduet er for lite", ReaderStatusLastPage: "siste side", ReaderStatusFirstPage: "første side", ReaderStatusRenderError: "gjengivelsesfeil: %v", ReaderStatusMaximumZoom: "maksimal zoom", ReaderStatusMinimumZoom: "minimal zoom", ReaderStatusInvalidPage: "ugyldig sidetall",
	ReaderViewTerminalTooSmall: "comicread: terminalvinduet er for lite", ReaderViewWindowTitle: "comicread — %s", ReaderViewPages: "sider %d/%d", ReaderViewPageRange: "sider %d–%d/%d", ReaderViewRendering: "gjengir", ReaderViewBookmarks: "Bokmerker", ReaderViewNoBookmarks: "(ingen bokmerker)", ReaderViewBookmarksHelp: "opp/ned flytt | enter åpne | esc lukk", ReaderViewGoToPage: "Gå til side: %s",
	ReaderViewHelp: `Taster

← →  forrige / neste side
↑ ↓  rull en zoomet side
+ -  zoom inn / ut
b    legg til / fjern bokmerke
v ← → forrige / neste bokmerke
c v  bokmerker
g123 enter  gå til side
q    avslutt

?    lukk hjelpen`,
	FilepickerHeader: "comicread — velg et kapittel\n%s\n\n", FilepickerNoEntries: "  (ingen støttede oppføringer)\n", FilepickerHelp: "\n↑/↓ flytt\n← overordnet katalog\n→ gå inn i katalog\nenter åpne fil\ns velg markert katalog\nf legg til / fjern gjeldende katalog fra favoritter\nF legg til favorittkatalog\nb favorittkataloger\no gå til katalog\nq avslutt\n", FilepickerWindowTitle: "comicread — velg en fil", FilepickerGoToPrompt: "\nGå til katalog: %s\n", FilepickerFavoritePrompt: "\nFavorittkatalog: %s\n", FilepickerGoToErr: "  feil: %s\n", FilepickerFavorites: "Favorittkataloger\n\n", FilepickerNoFavorites: "  (ingen favorittkataloger er konfigurert)\n", FilepickerFavoritesHelp: "\n↑/↓ flytt\nenter gå til katalog\nd fjern favoritt\nesc tilbake\n", FilepickerFavoriteErr: "  feil ved lagring av favoritter: %s\n", FilepickerErrResolveDir: "løs katalog %q: %w", FilepickerErrReadDir: "les katalog %q: %w", FilepickerErrRunPicker: "kjør filvelger: %w", FilepickerErrEmptyPath: "stien er tom", FilepickerErrNotDir: "%q er ikke en katalog",
	LoadingViewOpening: "%s åpnes…", LoadingViewWindowTitle: "comicread — åpner",
	CLIErrGetWorkingDir: "hent arbeidskatalog: %w", CLIErrPickFile: "velg fil: %w", CLIErrRunTUI: "kjør TUI: %w", CLIErrParseArgs: "tolk argumenter: %w", CLIErrOpenChapter: "åpne kapittel: %w", CLIErrOpenJournal: "åpne journal: %w", CLIErrClearJournal: "tøm journal: %w", CLIErrClearJournalRequiresInput: "--clear-journal krever en fil eller katalog", CLIErrNoPages: "kapittelet inneholder ingen lesbare bildesider", CLIErrInspectInput: "undersøk inndata %q: %w", CLIErrUnsupportedFile: "filen %q støttes ikke: støttede formater er CBZ, PDF, EPUB eller en bildekatalog",
	CLIFlagGraphicsUsage: "gjengiver: auto, ascii, dots, kitty, sixel eller iterm2", CLIFlagVersionUsage: "skriv ut versjon og avslutt", CLIFlagUpdateUsage: "se etter oppdateringer og avslutt", CLIFlagEnvUsage: "skriv ut comicread-miljøet og avslutt", CLIFlagClearJournalUsage: "fjern lokal journal for en fil eller katalog og avslutt", CLIFlagBookViewUsage: "vis sidepar fra venstre til høyre", CLIFlagRightBookViewUsage: "vis sidepar fra høyre til venstre", CLIFlagCircleBookViewUsage: "vis overlappende sidepar fra venstre til høyre", CLIFlagRightCircleBookViewUsage: "vis overlappende sidepar fra høyre til venstre", CLIErrMultipleBookViews: "bare ett bokvisningsvalg kan brukes", CLIErrInvalidView: "COMICREAD_VIEW %q støttes ikke (forventet book-view, right-view, circle-view eller right-circle-view)", CLIFlagOpenUsage: "katalog som skal åpnes i filvelgeren (standard: COMICREAD_DIR eller gjeldende katalog)", CLIErrOpenNotDir: "åpne katalog %q: ikke en katalog", CLIHelpHint: "kjør 'comicread --help' for bruk", CLIUsage: "bruk: comicread [valg] [fil]",
	CLIFlagWebUsage: "start en lokal nettleserbasert leser i stedet for terminalgrensesnittet", CLIErrWebArgs: "--web godtar ikke et fil- eller katalogargument", WebServerStarted: "comicreads nettleserbaserte leser kjører på %s (trykk Ctrl+C for å stoppe)", WebErrListen: "start nettserver: %w", WebErrServe: "kjør nettserver: %w",
	CLIUsageFull: `comicread — en minimal mangaleser for terminalen

bruk: comicread [valg] [fil]

valg:
  --config string     konfigurasjonsfil som skal brukes
  --graphics string   gjengiver: auto, ascii, dots, kitty, sixel eller iterm2 (standard "auto")
  --book-view         vis sidepar fra venstre til høyre
  --right-view        vis sidepar fra høyre til venstre
  --circle-view       vis overlappende sidepar fra venstre til høyre
  --right-circle-view
                      vis overlappende sidepar fra høyre til venstre
  --clear-journal     fjern lokal journal for en fil eller katalog og avslutt
  --reset-config     tilbakestill config.toml og avslutt
  --set-config value oppdater config.toml: nøkkel=verdi
  -o, --open string   katalog som skal åpnes i filvelgeren (standard: COMICREAD_DIR eller gjeldende katalog)
  --update            se etter oppdateringer og avslutt
  --web               start en lokal nettleserbasert leser i stedet for terminalgrensesnittet
  -v, --version       skriv ut versjon og avslutt
  -h, --help          vis denne hjelpen

Hvis ingen fil eller katalog er oppgitt, åpnes en interaktiv filvelger i COMICREAD_DIR
(hvis den er satt til en gyldig katalog) eller ellers i gjeldende katalog.`,
	ReaderViewMetadata: "Metadata",
}
