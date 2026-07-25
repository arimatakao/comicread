package i18n

var itMessages = map[string]string{
	CLIFlagConfigUsage: "file di configurazione da usare", CLIFlagResetConfigUsage: "reimposta config.toml e termina", CLIFlagSetConfigUsage: "aggiorna config.toml: chiave=valore",
	ReaderStatusWaitingTerminalSize: "in attesa delle dimensioni del terminale",
	ReaderStatusTerminalTooSmall:    "la finestra del terminale è troppo piccola",
	ReaderStatusLastPage:            "ultima pagina",
	ReaderStatusFirstPage:           "prima pagina",
	ReaderStatusRenderError:         "errore di rendering: %v",
	ReaderStatusMaximumZoom:         "zoom massimo",
	ReaderStatusMinimumZoom:         "zoom minimo",
	ReaderStatusInvalidPage:         "numero di pagina non valido",

	ReaderViewTerminalTooSmall: "comicread: la finestra del terminale è troppo piccola",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "pagine %d/%d",
	ReaderViewPageRange:        "pagine %d-%d/%d",
	ReaderViewRendering:        "rendering",
	ReaderViewBookmarks:        "Segnalibri",
	ReaderViewNoBookmarks:      "(nessun segnalibro)",
	ReaderViewBookmarksHelp:    "su/giù sposta | enter apri | esc chiudi",
	ReaderViewGoToPage:         "Vai alla pagina: %s",
	ReaderViewHelp: `Tasti

← →  pagina precedente / successiva
↑ ↓  scorri una pagina ingrandita
+ -  ingrandisci / riduci
b    aggiungi / rimuovi segnalibro
v ← → segnalibro precedente / successivo
c v  segnalibri
g123 enter  vai alla pagina
q    esci

?    chiudi l'aiuto`,

	FilepickerHeader:         "comicread — seleziona un capitolo\n%s\n\n",
	FilepickerNoEntries:      "  (nessuna voce supportata)\n",
	FilepickerHelp:           "\n↑/↓ sposta\n← cartella superiore\n→ entra nella cartella\nenter apri file\ns seleziona la cartella evidenziata\nf aggiungi / rimuovi la cartella corrente dai preferiti\nF aggiungi cartella preferita\nb cartelle preferite\no vai a una cartella\nq esci\n",
	FilepickerWindowTitle:    "comicread — scegli un file",
	FilepickerGoToPrompt:     "\nVai alla cartella: %s\n",
	FilepickerFavoritePrompt: "\nCartella preferita: %s\n",
	FilepickerFavorites:      "Cartelle preferite\n\n",
	FilepickerNoFavorites:    "  (nessuna cartella preferita configurata)\n",
	FilepickerFavoritesHelp:  "\n↑/↓ sposta\nenter vai alla cartella\nd rimuovi preferito\nesc indietro\n",
	FilepickerFavoriteErr:    "  errore nel salvataggio dei preferiti: %s\n",
	FilepickerGoToErr:        "  errore: %s\n",

	FilepickerErrResolveDir: "impossibile risolvere la cartella %q: %w",
	FilepickerErrReadDir:    "impossibile leggere la cartella %q: %w",
	FilepickerErrRunPicker:  "errore nell'avvio del selettore file: %w",
	FilepickerErrEmptyPath:  "il percorso è vuoto",
	FilepickerErrNotDir:     "%q non è una cartella",

	LoadingViewOpening:     "apertura di %s…",
	LoadingViewWindowTitle: "comicread — apertura",

	CLIErrGetWorkingDir:             "impossibile ottenere la cartella di lavoro: %w",
	CLIErrPickFile:                  "impossibile scegliere il file: %w",
	CLIErrRunTUI:                    "errore nell'avvio della TUI: %w",
	CLIErrParseArgs:                 "errore nell'analisi degli argomenti: %w",
	CLIErrOpenChapter:               "impossibile aprire il capitolo: %w",
	CLIErrOpenJournal:               "impossibile aprire il registro: %w",
	CLIErrClearJournal:              "impossibile eliminare il registro: %w",
	CLIErrClearJournalRequiresInput: "--clear-journal richiede un file o una directory",
	CLIErrNoPages:                   "il capitolo non contiene pagine immagine leggibili",
	CLIErrInspectInput:              "impossibile controllare l'input %q: %w",
	CLIErrUnsupportedFile:           "file non supportato %q: i formati supportati sono CBZ, PDF, EPUB o una cartella di immagini",
	CLIFlagGraphicsUsage:            "renderer: auto, ascii, dots, kitty, sixel o iterm2",
	CLIFlagVersionUsage:             "mostra la versione ed esci",
	CLIFlagUpdateUsage:              "verifica gli aggiornamenti ed esci",
	CLIFlagEnvUsage:                 "mostra l'ambiente comicread ed esci",
	CLIFlagClearJournalUsage:        "rimuovi il registro locale per un file o una directory ed esci",
	CLIFlagBookViewUsage:            "mostra coppie di pagine da sinistra a destra",
	CLIFlagRightBookViewUsage:       "mostra coppie di pagine da destra a sinistra",
	CLIFlagCircleBookViewUsage:      "mostra coppie di pagine sovrapposte da sinistra a destra",
	CLIFlagRightCircleBookViewUsage: "mostra coppie di pagine sovrapposte da destra a sinistra",
	CLIErrMultipleBookViews:         "può essere usata una sola opzione di visualizzazione libro",
	CLIErrInvalidView:               "valore COMICREAD_VIEW non supportato %q (previsti: book-view, right-view, circle-view o right-circle-view)",
	CLIFlagOpenUsage:                "directory da aprire nel selettore di file (predefinita: COMICREAD_DIR o la directory corrente)",
	CLIErrOpenNotDir:                "apri directory %q: non è una directory",
	CLIHelpHint:                     "esegui 'comicread --help' per l'aiuto",
	CLIUsage:                        "uso: comicread [opzioni] [file]",
	CLIUsageFull: `comicread — un lettore di manga minimale per il terminale

uso: comicread [opzioni] [file]

opzioni:
  --graphics string   renderer: auto, ascii, dots, kitty, sixel o iterm2 (predefinito "auto")
  --book-view         mostra coppie di pagine da sinistra a destra
  --right-view        mostra coppie di pagine da destra a sinistra
  --circle-view       mostra coppie di pagine sovrapposte da sinistra a destra
  --right-circle-view
                      mostra coppie di pagine sovrapposte da destra a sinistra
  --clear-journal    rimuovi il registro locale per un file o una directory ed esci
  -o, --open string   directory da aprire nel selettore di file (predefinita: COMICREAD_DIR o la directory corrente)
  --env               mostra l'ambiente comicread ed esci
  --update            verifica gli aggiornamenti ed esci
  -v, --version       mostra la versione ed esci
  -h, --help          mostra questo aiuto

Se non viene fornito alcun file o cartella, si apre un selettore di file interattivo in COMICREAD_DIR
(se impostata su una directory valida) o altrimenti nella directory corrente.

ambiente:
  COMICREAD_GRAPHICS  renderer predefinito: auto, ascii, dots, kitty, sixel o iterm2
  COMICREAD_PRERENDERED_NEXT      pagine successive da prerenderizzare (predefinito 1)
  COMICREAD_PRERENDERED_PREVIOUS  pagine precedenti da prerenderizzare (predefinito 1)
  COMICREAD_VIEW      visualizzazione predefinita: book-view, right-view, circle-view o right-circle-view
  COMICREAD_LANG      lingua dei messaggi: https://github.com/arimatakao/comicread#environment-variables (predefinito "en")
  COMICREAD_DIR       directory predefinita per il selettore di file quando non è indicato alcun percorso`,
	ReaderViewMetadata: "Metadati",
}
