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
	FilepickerHelp:        "\n↑/↓ déplacer\n← dossier parent\n→ entrer dans le dossier\nenter ouvrir le fichier\ns sélectionner le dossier surligné\no aller à un dossier\nq quitter\n",
	FilepickerWindowTitle: "comicread — choisir un fichier",
	FilepickerGoToPrompt:  "\nAller au dossier : %s\n",
	FilepickerGoToErr:     "  erreur : %s\n",

	FilepickerErrResolveDir: "impossible de déterminer le dossier %q : %w",
	FilepickerErrReadDir:    "impossible de lire le dossier %q : %w",
	FilepickerErrRunPicker:  "erreur lors du lancement du sélecteur de fichiers : %w",
	FilepickerErrEmptyPath:  "le chemin est vide",
	FilepickerErrNotDir:     "%q n'est pas un dossier",

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
	CLIFlagOpenUsage:                "dossier à ouvrir dans le sélecteur de fichiers (par défaut : COMICREAD_DIR ou le dossier actuel)",
	CLIErrOpenNotDir:                "ouvrir le dossier %q : ce n'est pas un dossier",
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
  -o, --open string   dossier à ouvrir dans le sélecteur de fichiers (par défaut : COMICREAD_DIR ou le dossier actuel)
  --env               afficher l'environnement comicread et quitter
  --update            vérifier les mises à jour et quitter
  -v, --version       afficher la version et quitter
  -h, --help          afficher cette aide

Si aucun fichier ou dossier n'est indiqué, un sélecteur de fichiers interactif s'ouvre dans COMICREAD_DIR
(s'il est défini sur un dossier valide) ou dans le dossier actuel sinon.

environnement :
  COMICREAD_GRAPHICS  moteur de rendu par défaut : auto, ascii, dots, kitty, sixel ou iterm2
  COMICREAD_PRERENDERED_NEXT      pages suivantes à pré-rendre (par défaut 1)
  COMICREAD_PRERENDERED_PREVIOUS  pages précédentes à pré-rendre (par défaut 1)
  COMICREAD_VIEW      affichage par défaut : book-view, right-view, circle-view ou right-circle-view
  COMICREAD_LANG      langue des messages : en, uk, pl, de, fr, es, cs, ro, it, ko, ja, id, hi, el, tr, kk ou ka (par défaut : "en")
  COMICREAD_DIR       dossier par défaut du sélecteur de fichiers lorsqu'aucun chemin n'est indiqué`,
}
