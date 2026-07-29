package i18n

var deMessages = map[string]string{
	CLIFlagConfigUsage: "zu verwendende Konfigurationsdatei", CLIFlagResetConfigUsage: "config.toml auf Standardwerte zurücksetzen und beenden", CLIFlagSetConfigUsage: "config.toml aktualisieren: schlüssel=wert",
	ReaderStatusWaitingTerminalSize: "warte auf Terminalgröße",
	ReaderStatusTerminalTooSmall:    "Terminalfenster ist zu klein",
	ReaderStatusLastPage:            "letzte Seite",
	ReaderStatusFirstPage:           "erste Seite",
	ReaderStatusRenderError:         "Darstellungsfehler: %v",
	ReaderStatusMaximumZoom:         "maximale Vergrößerung",
	ReaderStatusMinimumZoom:         "minimale Vergrößerung",
	ReaderStatusInvalidPage:         "ungültige Seitenzahl",

	ReaderViewTerminalTooSmall: "comicread: Terminalfenster ist zu klein",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "Seiten %d/%d",
	ReaderViewPageRange:        "Seiten %d-%d/%d",
	ReaderViewRendering:        "wird dargestellt",
	ReaderViewBookmarks:        "Lesezeichen",
	ReaderViewNoBookmarks:      "(keine Lesezeichen)",
	ReaderViewBookmarksHelp:    "hoch/runter bewegen | enter öffnen | esc schließen",
	ReaderViewGoToPage:         "Zu Seite springen: %s",
	ReaderViewHelp: `Tasten

← →  vorherige / nächste Seite
↑ ↓  vergrößerte Seite scrollen
+ -  vergrößern / verkleinern
b    Lesezeichen hinzufügen / entfernen
v ← → vorheriges / nächstes Lesezeichen
c v  Lesezeichen
g123 enter  zu Seite springen
q    beenden

?    Hilfe schließen`,

	FilepickerHeader:         "comicread — Kapitel auswählen\n%s\n\n",
	FilepickerNoEntries:      "  (keine unterstützten Einträge)\n",
	FilepickerHelp:           "\n↑/↓ bewegen\n← übergeordneter Ordner\n→ Ordner öffnen\nenter Datei öffnen\ns markierten Ordner auswählen\nf aktuellen Ordner zu Favoriten hinzufügen / daraus entfernen\nF Favoritenordner hinzufügen\nb Favoritenordner\no zu Ordner wechseln\nq beenden\n",
	FilepickerWindowTitle:    "comicread — Datei auswählen",
	FilepickerGoToPrompt:     "\nZu Ordner wechseln: %s\n",
	FilepickerFavoritePrompt: "\nFavoritenordner: %s\n",
	FilepickerFavorites:      "Favoritenordner\n\n",
	FilepickerNoFavorites:    "  (keine Favoritenordner konfiguriert)\n",
	FilepickerFavoritesHelp:  "\n↑/↓ bewegen\nenter zu Ordner wechseln\nd Favorit entfernen\nesc zurück\n",
	FilepickerFavoriteErr:    "  Fehler beim Speichern der Favoriten: %s\n",
	FilepickerGoToErr:        "  Fehler: %s\n",

	FilepickerErrResolveDir: "Ordner %q kann nicht bestimmt werden: %w",
	FilepickerErrReadDir:    "Ordner %q kann nicht gelesen werden: %w",
	FilepickerErrRunPicker:  "Fehler beim Starten der Dateiauswahl: %w",
	FilepickerErrEmptyPath:  "Pfad ist leer",
	FilepickerErrNotDir:     "%q ist kein Ordner",

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
	CLIFlagOpenUsage:                "im Dateiauswahldialog zu öffnender Ordner (Standard: COMICREAD_DIR oder aktueller Ordner)",
	CLIErrOpenNotDir:                "Ordner öffnen %q: kein Ordner",
	CLIFlagWebUsage:                 "lokalen Browser-Reader statt der Terminaloberfläche starten",
	CLIErrWebArgs:                   "--web akzeptiert kein Datei- oder Ordnerargument",
	WebServerStarted:                "comicread-Web-Reader läuft unter %s (zum Beenden Ctrl+C drücken)",
	WebErrListen:                    "Webserver starten: %w",
	WebErrServe:                     "Webserver ausführen: %w",
	CLIHelpHint:                     "für Hilfe 'comicread --help' ausführen",
	CLIUsage:                        "Aufruf: comicread [Optionen] [Datei]",
	CLIUsageFull: `comicread — ein minimaler Manga-Reader für das Terminal

Aufruf: comicread [Optionen] [Datei]

Optionen:
  --config string     zu verwendende Konfigurationsdatei
  --graphics string   Renderer: auto, ascii, dots, kitty, sixel oder iterm2 (Standard: "auto")
  --book-view         Seiten paarweise von links nach rechts anzeigen
  --right-view        Seiten paarweise von rechts nach links anzeigen
  --circle-view       überlappende Seitenpaare von links nach rechts anzeigen
  --right-circle-view
                      überlappende Seitenpaare von rechts nach links anzeigen
  --clear-journal    lokales Journal für eine Datei oder einen Ordner löschen und beenden
  --reset-config     config.toml auf Standardwerte zurücksetzen und beenden
  --set-config value config.toml aktualisieren: schlüssel=wert
  -o, --open string   im Dateiauswahldialog zu öffnender Ordner (Standard: COMICREAD_DIR oder aktueller Ordner)
  --update            nach Updates suchen und beenden
  --web               lokalen Browser-Reader statt der Terminaloberfläche starten
  -v, --version       Version ausgeben und beenden
  -h, --help          diese Hilfe anzeigen

Wenn keine Datei und kein Ordner angegeben ist, öffnet sich eine interaktive Dateiauswahl in COMICREAD_DIR
(falls auf einen gültigen Ordner gesetzt), sonst im aktuellen Ordner.`,
	ReaderViewMetadata: "Metadaten",
}
