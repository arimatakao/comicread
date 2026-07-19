package i18n

var noMessages = map[string]string{
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
	FilepickerHeader: "comicread — velg et kapittel\n%s\n\n", FilepickerNoEntries: "  (ingen støttede oppføringer)\n", FilepickerHelp: "\n↑/↓ flytt\n← overordnet katalog\n→ gå inn i katalog\nenter åpne fil\ns velg markert katalog\no gå til katalog\nq avslutt\n", FilepickerWindowTitle: "comicread — velg en fil", FilepickerGoToPrompt: "\nGå til katalog: %s\n", FilepickerGoToErr: "  feil: %s\n", FilepickerErrResolveDir: "løs katalog %q: %w", FilepickerErrReadDir: "les katalog %q: %w", FilepickerErrRunPicker: "kjør filvelger: %w", FilepickerErrEmptyPath: "stien er tom", FilepickerErrNotDir: "%q er ikke en katalog",
	LoadingViewOpening: "%s åpnes…", LoadingViewWindowTitle: "comicread — åpner",
	CLIErrGetWorkingDir: "hent arbeidskatalog: %w", CLIErrPickFile: "velg fil: %w", CLIErrRunTUI: "kjør TUI: %w", CLIErrParseArgs: "tolk argumenter: %w", CLIErrOpenChapter: "åpne kapittel: %w", CLIErrOpenJournal: "åpne journal: %w", CLIErrClearJournal: "tøm journal: %w", CLIErrClearJournalRequiresInput: "--clear-journal krever en fil eller katalog", CLIErrNoPages: "kapittelet inneholder ingen lesbare bildesider", CLIErrInspectInput: "undersøk inndata %q: %w", CLIErrUnsupportedFile: "filen %q støttes ikke: støttede formater er CBZ, PDF, EPUB eller en bildekatalog",
	CLIFlagGraphicsUsage: "gjengiver: auto, ascii, dots, kitty, sixel eller iterm2", CLIFlagVersionUsage: "skriv ut versjon og avslutt", CLIFlagUpdateUsage: "se etter oppdateringer og avslutt", CLIFlagEnvUsage: "skriv ut comicread-miljøet og avslutt", CLIFlagClearJournalUsage: "fjern lokal journal for en fil eller katalog og avslutt", CLIFlagBookViewUsage: "vis sidepar fra venstre til høyre", CLIFlagRightBookViewUsage: "vis sidepar fra høyre til venstre", CLIFlagCircleBookViewUsage: "vis overlappende sidepar fra venstre til høyre", CLIFlagRightCircleBookViewUsage: "vis overlappende sidepar fra høyre til venstre", CLIErrMultipleBookViews: "bare ett bokvisningsvalg kan brukes", CLIErrInvalidView: "COMICREAD_VIEW %q støttes ikke (forventet book-view, right-view, circle-view eller right-circle-view)", CLIFlagOpenUsage: "katalog som skal åpnes i filvelgeren (standard: COMICREAD_DIR eller gjeldende katalog)", CLIErrOpenNotDir: "åpne katalog %q: ikke en katalog", CLIHelpHint: "kjør 'comicread --help' for bruk", CLIUsage: "bruk: comicread [valg] [fil]",
	CLIUsageFull: `comicread — en minimal mangaleser for terminalen

bruk: comicread [valg] [fil]

valg:
  --graphics string   gjengiver: auto, ascii, dots, kitty, sixel eller iterm2 (standard "auto")
  --book-view         vis sidepar fra venstre til høyre
  --right-view        vis sidepar fra høyre til venstre
  --circle-view       vis overlappende sidepar fra venstre til høyre
  --right-circle-view
                      vis overlappende sidepar fra høyre til venstre
  --clear-journal     fjern lokal journal for en fil eller katalog og avslutt
  -o, --open string   katalog som skal åpnes i filvelgeren (standard: COMICREAD_DIR eller gjeldende katalog)
  --env               skriv ut comicread-miljøet og avslutt
  --update            se etter oppdateringer og avslutt
  -v, --version       skriv ut versjon og avslutt
  -h, --help          vis denne hjelpen

Hvis ingen fil eller katalog er oppgitt, åpnes en interaktiv filvelger i COMICREAD_DIR
(hvis den er satt til en gyldig katalog) eller ellers i gjeldende katalog.

miljø:
  COMICREAD_GRAPHICS  standardgjengiver: auto, ascii, dots, kitty, sixel eller iterm2
  COMICREAD_PRERENDERED_NEXT      neste sider som skal forhåndsgjengis (standard 1)
  COMICREAD_PRERENDERED_PREVIOUS  forrige sider som skal forhåndsgjengis (standard 1)
  COMICREAD_VIEW      standardvisning: book-view, right-view, circle-view eller right-circle-view
  COMICREAD_LANG      meldingsspråk: https://github.com/arimatakao/comicread#environment-variables (standard "en")
  COMICREAD_DIR       standardkatalog for filvelgeren når ingen sti er angitt`,
}
