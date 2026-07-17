package i18n

var deMessages = map[string]string{
	ReaderStatusWaitingTerminalSize: "warte auf Terminalgröße",
	ReaderStatusTerminalTooSmall:    "Terminalfenster ist zu klein",
	ReaderStatusLastPage:            "letzte Seite",
	ReaderStatusFirstPage:           "erste Seite",
	ReaderStatusRenderError:         "Darstellungsfehler: %v",
	ReaderStatusMaximumZoom:         "maximale Vergrößerung",
	ReaderStatusMinimumZoom:         "minimale Vergrößerung",

	ReaderViewTerminalTooSmall: "comicread: Terminalfenster ist zu klein",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "Seiten %d/%d",
	ReaderViewPageRange:        "Seiten %d-%d/%d",
	ReaderViewRendering:        "wird dargestellt",
	ReaderViewBookmarks:        "Lesezeichen",
	ReaderViewNoBookmarks:      "(keine Lesezeichen)",
	ReaderViewBookmarksHelp:    "hoch/runter bewegen | enter öffnen | esc schließen",
	ReaderViewHelp: `Tasten

← →  vorherige / nächste Seite
↑ ↓  vergrößerte Seite scrollen
+ -  vergrößern / verkleinern
b    Lesezeichen hinzufügen / entfernen
v ← → vorheriges / nächstes Lesezeichen
c v  Lesezeichen
q    beenden

?    Hilfe schließen`,

	FilepickerHeader:      "comicread — Kapitel auswählen\n%s\n\n",
	FilepickerNoEntries:   "  (keine unterstützten Einträge)\n",
	FilepickerHelp:        "\n↑/↓ bewegen  |  ← übergeordneter Ordner  |  → Ordner öffnen  |  enter Datei öffnen  |  s markierten Ordner auswählen  |  q beenden\n",
	FilepickerWindowTitle: "comicread — Datei auswählen",

	FilepickerErrResolveDir: "Ordner %q kann nicht bestimmt werden: %w",
	FilepickerErrReadDir:    "Ordner %q kann nicht gelesen werden: %w",
	FilepickerErrRunPicker:  "Fehler beim Starten der Dateiauswahl: %w",

	LoadingViewOpening:     "%s wird geöffnet…",
	LoadingViewWindowTitle: "comicread — wird geöffnet",

	CLIErrGetWorkingDir:             "Arbeitsordner kann nicht ermittelt werden: %w",
	CLIErrPickFile:                  "Datei kann nicht ausgewählt werden: %w",
	CLIErrRunTUI:                    "Fehler beim Starten der TUI: %w",
	CLIErrParseArgs:                 "Fehler beim Verarbeiten der Argumente: %w",
	CLIErrOpenChapter:               "Kapitel kann nicht geöffnet werden: %w",
	CLIErrOpenJournal:               "Journal kann nicht geöffnet werden: %w",
	CLIErrClearJournal:              "Journal kann nicht gelöscht werden: %w",
	CLIErrClearJournalRequiresInput: "--clear-journal benötigt eine Datei oder einen Ordner",
	CLIErrNoPages:                   "Kapitel enthält keine lesbaren Bildseiten",
	CLIErrInspectInput:              "Eingabe %q kann nicht geprüft werden: %w",
	CLIErrUnsupportedFile:           "nicht unterstützte Datei %q: unterstützte Formate sind CBZ, PDF, EPUB oder ein Bildordner",
	CLIFlagGraphicsUsage:            "Renderer: auto, ascii, dots, kitty, sixel oder iterm2",
	CLIFlagVersionUsage:             "Version ausgeben und beenden",
	CLIFlagUpdateUsage:              "nach Updates suchen und beenden",
	CLIFlagEnvUsage:                 "comicread-Umgebung ausgeben und beenden",
	CLIFlagClearJournalUsage:        "lokales Journal für eine Datei oder einen Ordner löschen und beenden",
	CLIFlagBookViewUsage:            "Seiten paarweise von links nach rechts anzeigen",
	CLIFlagRightBookViewUsage:       "Seiten paarweise von rechts nach links anzeigen",
	CLIFlagCircleBookViewUsage:      "überlappende Seitenpaare von links nach rechts anzeigen",
	CLIFlagRightCircleBookViewUsage: "überlappende Seitenpaare von rechts nach links anzeigen",
	CLIErrMultipleBookViews:         "nur eine Buchansichtsoption darf verwendet werden",
	CLIErrInvalidView:               "nicht unterstützter COMICREAD_VIEW-Wert %q (erwartet: book-view, right-view, circle-view oder right-circle-view)",
	CLIHelpHint:                     "für Hilfe 'comicread --help' ausführen",
	CLIUsage:                        "Aufruf: comicread [Optionen] [Datei]",
	CLIUsageFull: `comicread — ein minimaler Manga-Reader für das Terminal

Aufruf: comicread [Optionen] [Datei]

Optionen:
  --graphics string   Renderer: auto, ascii, dots, kitty, sixel oder iterm2 (Standard: "auto")
  --book-view         Seiten paarweise von links nach rechts anzeigen
  --right-view        Seiten paarweise von rechts nach links anzeigen
  --circle-view       überlappende Seitenpaare von links nach rechts anzeigen
  --right-circle-view
                      überlappende Seitenpaare von rechts nach links anzeigen
  --clear-journal    lokales Journal für eine Datei oder einen Ordner löschen und beenden
  --env               comicread-Umgebung ausgeben und beenden
  --update            nach Updates suchen und beenden
  -v, --version       Version ausgeben und beenden
  -h, --help          diese Hilfe anzeigen

Wenn keine Datei und kein Ordner angegeben ist, öffnet sich im aktuellen Ordner eine interaktive Dateiauswahl.

Umgebung:
  COMICREAD_GRAPHICS  Standard-Renderer: auto, ascii, dots, kitty, sixel oder iterm2
  COMICREAD_PRERENDERED_NEXT      nächste Seiten vorab rendern (Standard 1)
  COMICREAD_PRERENDERED_PREVIOUS  vorherige Seiten vorab rendern (Standard 1)
  COMICREAD_VIEW      Standardansicht: book-view, right-view, circle-view oder right-circle-view
  COMICREAD_LANG      Sprache der Meldungen: en, uk, pl, de, fr, es, cs, ro, it, ko, ja, id, hi, el, tr, kk oder ka (Standard: "en")`,
}
