package i18n

var frMessages = map[string]string{
	ReaderStatusWaitingTerminalSize: "en attente de la taille du terminal",
	ReaderStatusTerminalTooSmall:    "la fenêtre du terminal est trop petite",
	ReaderStatusLastPage:            "dernière page",
	ReaderStatusFirstPage:           "première page",
	ReaderStatusRenderError:         "erreur de rendu : %v",
	ReaderStatusMaximumZoom:         "zoom maximal",
	ReaderStatusMinimumZoom:         "zoom minimal",

	ReaderViewTerminalTooSmall: "comicread : la fenêtre du terminal est trop petite",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "pages %d/%d",
	ReaderViewPageRange:        "pages %d-%d/%d",
	ReaderViewRendering:        "rendu en cours",
	ReaderViewBookmarks:        "Signets",
	ReaderViewNoBookmarks:      "(aucun signet)",
	ReaderViewBookmarksHelp:    "haut/bas déplacer | enter ouvrir | esc fermer",
	ReaderViewHelp: `Touches

← →  page précédente / suivante
↑ ↓  faire défiler une page zoomée
+ -  zoomer / dézoomer
b    ajouter / supprimer un signet
v ← → signet précédent / suivant
c v  signets
q    quitter

?    fermer l'aide`,

	FilepickerHeader:      "comicread — sélectionner un chapitre\n%s\n\n",
	FilepickerNoEntries:   "  (aucune entrée prise en charge)\n",
	FilepickerHelp:        "\n↑/↓ déplacer  |  ← dossier parent  |  → entrer dans le dossier  |  enter ouvrir le fichier  |  s sélectionner le dossier surligné  |  q quitter\n",
	FilepickerWindowTitle: "comicread — choisir un fichier",

	FilepickerErrResolveDir: "impossible de déterminer le dossier %q : %w",
	FilepickerErrReadDir:    "impossible de lire le dossier %q : %w",
	FilepickerErrRunPicker:  "erreur lors du lancement du sélecteur de fichiers : %w",

	LoadingViewOpening:     "ouverture de %s…",
	LoadingViewWindowTitle: "comicread — ouverture",

	CLIErrGetWorkingDir:             "impossible d'obtenir le répertoire de travail : %w",
	CLIErrPickFile:                  "impossible de choisir le fichier : %w",
	CLIErrRunTUI:                    "erreur lors du lancement de la TUI : %w",
	CLIErrParseArgs:                 "erreur d'analyse des arguments : %w",
	CLIErrOpenChapter:               "impossible d'ouvrir le chapitre : %w",
	CLIErrOpenJournal:               "impossible d'ouvrir le journal : %w",
	CLIErrClearJournal:              "impossible d'effacer le journal : %w",
	CLIErrClearJournalRequiresInput: "--clear-journal requiert un fichier ou un dossier",
	CLIErrNoPages:                   "le chapitre ne contient aucune page image lisible",
	CLIErrInspectInput:              "impossible d'inspecter l'entrée %q : %w",
	CLIErrUnsupportedFile:           "fichier non pris en charge %q : formats pris en charge : CBZ, PDF, EPUB ou dossier d'images",
	CLIFlagGraphicsUsage:            "moteur de rendu : auto, ascii, dots, kitty, sixel ou iterm2",
	CLIFlagVersionUsage:             "afficher la version et quitter",
	CLIFlagUpdateUsage:              "vérifier les mises à jour et quitter",
	CLIFlagEnvUsage:                 "afficher l'environnement comicread et quitter",
	CLIFlagClearJournalUsage:        "supprimer le journal local d'un fichier ou dossier et quitter",
	CLIFlagBookViewUsage:            "afficher les pages par paires de gauche à droite",
	CLIFlagRightBookViewUsage:       "afficher les pages par paires de droite à gauche",
	CLIFlagCircleBookViewUsage:      "afficher les paires de pages superposées de gauche à droite",
	CLIFlagRightCircleBookViewUsage: "afficher les paires de pages superposées de droite à gauche",
	CLIErrMultipleBookViews:         "une seule option d'affichage de livre peut être utilisée",
	CLIErrInvalidView:               "valeur COMICREAD_VIEW non prise en charge %q (attendu : book-view, right-view, circle-view ou right-circle-view)",
	CLIHelpHint:                     "exécutez 'comicread --help' pour l'aide",
	CLIUsage:                        "utilisation : comicread [options] [fichier]",
	CLIUsageFull: `comicread — un lecteur de mangas minimal pour le terminal

utilisation : comicread [options] [fichier]

options :
  --graphics string   moteur de rendu : auto, ascii, dots, kitty, sixel ou iterm2 (par défaut : "auto")
  --book-view         afficher les pages par paires de gauche à droite
  --right-view        afficher les pages par paires de droite à gauche
  --circle-view       afficher les paires de pages superposées de gauche à droite
  --right-circle-view
                      afficher les paires de pages superposées de droite à gauche
  --clear-journal    supprimer le journal local d'un fichier ou dossier et quitter
  --env               afficher l'environnement comicread et quitter
  --update            vérifier les mises à jour et quitter
  -v, --version       afficher la version et quitter
  -h, --help          afficher cette aide

Si aucun fichier ou dossier n'est indiqué, un sélecteur de fichiers interactif s'ouvre dans le dossier actuel.

environnement :
  COMICREAD_GRAPHICS  moteur de rendu par défaut : auto, ascii, dots, kitty, sixel ou iterm2
  COMICREAD_VIEW      affichage par défaut : book-view, right-view, circle-view ou right-circle-view
  COMICREAD_LANG      langue des messages : en, uk, pl, de, fr, es, cs, ro, it, ko, ja, id, hi, el, tr, kk ou ka (par défaut : "en")`,
}
