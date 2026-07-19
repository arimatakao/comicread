package i18n

var fiMessages = map[string]string{
	ReaderStatusWaitingTerminalSize: "odotetaan päätteen kokoa", ReaderStatusTerminalTooSmall: "pääteikkuna on liian pieni", ReaderStatusLastPage: "viimeinen sivu", ReaderStatusFirstPage: "ensimmäinen sivu", ReaderStatusRenderError: "renderöintivirhe: %v", ReaderStatusMaximumZoom: "suurin zoomaus", ReaderStatusMinimumZoom: "pienin zoomaus", ReaderStatusInvalidPage: "virheellinen sivunumero",
	ReaderViewTerminalTooSmall: "comicread: pääteikkuna on liian pieni", ReaderViewWindowTitle: "comicread — %s", ReaderViewPages: "sivut %d/%d", ReaderViewPageRange: "sivut %d–%d/%d", ReaderViewRendering: "renderöidään", ReaderViewBookmarks: "Kirjanmerkit", ReaderViewNoBookmarks: "(ei kirjanmerkkejä)", ReaderViewBookmarksHelp: "ylös/alas siirry | enter avaa | esc sulje", ReaderViewGoToPage: "Siirry sivulle: %s",
	ReaderViewHelp: `Näppäimet

← →  edellinen / seuraava sivu
↑ ↓  vieritä zoomattua sivua
+ -  suurenna / pienennä
b    lisää / poista kirjanmerkki
v ← → edellinen / seuraava kirjanmerkki
c v  kirjanmerkit
g123 enter  siirry sivulle
q    lopeta

?    sulje ohje`,
	FilepickerHeader: "comicread — valitse luku\n%s\n\n", FilepickerNoEntries: "  (ei tuettuja kohteita)\n", FilepickerHelp: "\n↑/↓ siirry\n← ylähakemisto\n→ siirry hakemistoon\nenter avaa tiedosto\ns valitse korostettu hakemisto\no siirry hakemistoon\nq lopeta\n", FilepickerWindowTitle: "comicread — valitse tiedosto", FilepickerGoToPrompt: "\nSiirry hakemistoon: %s\n", FilepickerGoToErr: "  virhe: %s\n", FilepickerErrResolveDir: "selvitä hakemisto %q: %w", FilepickerErrReadDir: "lue hakemisto %q: %w", FilepickerErrRunPicker: "suorita tiedostovalitsin: %w", FilepickerErrEmptyPath: "polku on tyhjä", FilepickerErrNotDir: "%q ei ole hakemisto",
	LoadingViewOpening: "avataan %s…", LoadingViewWindowTitle: "comicread — avataan",
	CLIErrGetWorkingDir: "hae työhakemisto: %w", CLIErrPickFile: "valitse tiedosto: %w", CLIErrRunTUI: "suorita TUI: %w", CLIErrParseArgs: "jäsennä argumentit: %w", CLIErrOpenChapter: "avaa luku: %w", CLIErrOpenJournal: "avaa päiväkirja: %w", CLIErrClearJournal: "tyhjennä päiväkirja: %w", CLIErrClearJournalRequiresInput: "--clear-journal vaatii tiedoston tai hakemiston", CLIErrNoPages: "luku ei sisällä luettavia kuvasivuja", CLIErrInspectInput: "tutki syöte %q: %w", CLIErrUnsupportedFile: "tiedostoa %q ei tueta: tuetut muodot ovat CBZ, PDF, EPUB tai kuvahakemisto",
	CLIFlagGraphicsUsage: "renderöijä: auto, ascii, dots, kitty, sixel tai iterm2", CLIFlagVersionUsage: "tulosta versio ja lopeta", CLIFlagUpdateUsage: "tarkista päivitykset ja lopeta", CLIFlagEnvUsage: "tulosta comicread-ympäristö ja lopeta", CLIFlagClearJournalUsage: "poista tiedoston tai hakemiston paikallinen päiväkirja ja lopeta", CLIFlagBookViewUsage: "näytä sivuparit vasemmalta oikealle", CLIFlagRightBookViewUsage: "näytä sivuparit oikealta vasemmalle", CLIFlagCircleBookViewUsage: "näytä limittäiset sivuparit vasemmalta oikealle", CLIFlagRightCircleBookViewUsage: "näytä limittäiset sivuparit oikealta vasemmalle", CLIErrMultipleBookViews: "vain yhtä kirjanäkymäasetusta voidaan käyttää", CLIErrInvalidView: "COMICREAD_VIEW %q ei ole tuettu (odotetaan book-view, right-view, circle-view tai right-circle-view)", CLIFlagOpenUsage: "tiedostovalitsimessa avattava hakemisto (oletus: COMICREAD_DIR tai nykyinen hakemisto)", CLIErrOpenNotDir: "avaa hakemisto %q: ei ole hakemisto", CLIHelpHint: "suorita 'comicread --help' saadaksesi ohjeen", CLIUsage: "käyttö: comicread [asetukset] [tiedosto]",
	CLIUsageFull: `comicread — minimalistinen mangalukija päätteelle

käyttö: comicread [asetukset] [tiedosto]

asetukset:
  --graphics string   renderöijä: auto, ascii, dots, kitty, sixel tai iterm2 (oletus "auto")
  --book-view         näytä sivuparit vasemmalta oikealle
  --right-view        näytä sivuparit oikealta vasemmalle
  --circle-view       näytä limittäiset sivuparit vasemmalta oikealle
  --right-circle-view
                      näytä limittäiset sivuparit oikealta vasemmalle
  --clear-journal     poista tiedoston tai hakemiston paikallinen päiväkirja ja lopeta
  -o, --open string   tiedostovalitsimessa avattava hakemisto (oletus: COMICREAD_DIR tai nykyinen hakemisto)
  --env               tulosta comicread-ympäristö ja lopeta
  --update            tarkista päivitykset ja lopeta
  -v, --version       tulosta versio ja lopeta
  -h, --help          näytä tämä ohje

Jos tiedostoa tai hakemistoa ei anneta, interaktiivinen tiedostovalitsin avautuu COMICREAD_DIR-hakemistossa
(jos se on asetettu kelvolliseksi hakemistoksi) tai muuten nykyisessä hakemistossa.

ympäristö:
  COMICREAD_GRAPHICS  oletusrenderöijä: auto, ascii, dots, kitty, sixel tai iterm2
  COMICREAD_PRERENDERED_NEXT      esirenderöitävät seuraavat sivut (oletus 1)
  COMICREAD_PRERENDERED_PREVIOUS  esirenderöitävät edelliset sivut (oletus 1)
  COMICREAD_VIEW      oletusnäkymä: book-view, right-view, circle-view tai right-circle-view
  COMICREAD_LANG      viestien kieli: https://github.com/arimatakao/comicread#environment-variables (oletus "en")
  COMICREAD_DIR       tiedostovalitsimen oletushakemisto, kun polkua ei anneta`,
}
