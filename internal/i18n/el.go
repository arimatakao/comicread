package i18n

var elMessages = map[string]string{
	ReaderStatusWaitingTerminalSize: "αναμονή για το μέγεθος του τερματικού",
	ReaderStatusTerminalTooSmall:    "το παράθυρο του τερματικού είναι πολύ μικρό",
	ReaderStatusLastPage:            "τελευταία σελίδα",
	ReaderStatusFirstPage:           "πρώτη σελίδα",
	ReaderStatusRenderError:         "σφάλμα απόδοσης: %v",
	ReaderStatusMaximumZoom:         "μέγιστη μεγέθυνση",
	ReaderStatusMinimumZoom:         "ελάχιστη μεγέθυνση",

	ReaderViewTerminalTooSmall: "comicread: το παράθυρο του τερματικού είναι πολύ μικρό",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "σελίδες %d/%d",
	ReaderViewPageRange:        "σελίδες %d-%d/%d",
	ReaderViewRendering:        "απόδοση",
	ReaderViewBookmarks:        "Σελιδοδείκτες",
	ReaderViewNoBookmarks:      "(δεν υπάρχουν σελιδοδείκτες)",
	ReaderViewBookmarksHelp:    "πάνω/κάτω μετακίνηση | enter άνοιγμα | esc κλείσιμο",
	ReaderViewHelp: `Πλήκτρα

← →  προηγούμενη / επόμενη σελίδα
↑ ↓  κύλιση σε μεγεθυμένη σελίδα
+ -  μεγέθυνση / σμίκρυνση
b    προσθήκη / αφαίρεση σελιδοδείκτη
v ← → προηγούμενος / επόμενος σελιδοδείκτης
c v  σελιδοδείκτες
q    έξοδος

?    κλείσιμο βοήθειας`,

	FilepickerHeader:      "comicread — επιλέξτε κεφάλαιο\n%s\n\n",
	FilepickerNoEntries:   "  (δεν υπάρχουν υποστηριζόμενες καταχωρήσεις)\n",
	FilepickerHelp:        "\n↑/↓ μετακίνηση\n← γονικός κατάλογος\n→ είσοδος στον κατάλογο\nenter άνοιγμα αρχείου\ns επιλογή επισημασμένου καταλόγου\no μετάβαση σε κατάλογο\nq έξοδος\n",
	FilepickerWindowTitle: "comicread — επιλογή αρχείου",
	FilepickerGoToPrompt:  "\nΜετάβαση σε κατάλογο: %s\n",
	FilepickerGoToErr:     "  σφάλμα: %s\n",

	FilepickerErrResolveDir: "δεν είναι δυνατός ο προσδιορισμός του καταλόγου %q: %w",
	FilepickerErrReadDir:    "δεν είναι δυνατή η ανάγνωση του καταλόγου %q: %w",
	FilepickerErrRunPicker:  "σφάλμα κατά την εκκίνηση της επιλογής αρχείου: %w",
	FilepickerErrEmptyPath:  "η διαδρομή είναι κενή",
	FilepickerErrNotDir:     "το %q δεν είναι κατάλογος",

	LoadingViewOpening:     "άνοιγμα του %s…",
	LoadingViewWindowTitle: "comicread — άνοιγμα",

	CLIErrGetWorkingDir:             "δεν είναι δυνατή η λήψη του καταλόγου εργασίας: %w",
	CLIErrPickFile:                  "δεν είναι δυνατή η επιλογή του αρχείου: %w",
	CLIErrRunTUI:                    "σφάλμα κατά την εκκίνηση του TUI: %w",
	CLIErrParseArgs:                 "σφάλμα κατά την ανάλυση των ορισμάτων: %w",
	CLIErrOpenChapter:               "δεν είναι δυνατό το άνοιγμα του κεφαλαίου: %w",
	CLIErrOpenJournal:               "δεν είναι δυνατό το άνοιγμα του ημερολογίου: %w",
	CLIErrClearJournal:              "δεν είναι δυνατή η εκκαθάριση του ημερολογίου: %w",
	CLIErrClearJournalRequiresInput: "το --clear-journal απαιτεί αρχείο ή κατάλογο",
	CLIErrNoPages:                   "το κεφάλαιο δεν περιέχει αναγνώσιμες σελίδες εικόνων",
	CLIErrInspectInput:              "δεν είναι δυνατός ο έλεγχος της εισόδου %q: %w",
	CLIErrUnsupportedFile:           "μη υποστηριζόμενο αρχείο %q: υποστηρίζονται CBZ, PDF, EPUB ή κατάλογος εικόνων",
	CLIFlagGraphicsUsage:            "αποδότης: auto, ascii, dots, kitty, sixel ή iterm2",
	CLIFlagVersionUsage:             "εμφάνιση έκδοσης και έξοδος",
	CLIFlagUpdateUsage:              "έλεγχος ενημερώσεων και έξοδος",
	CLIFlagEnvUsage:                 "εμφάνιση περιβάλλοντος comicread και έξοδος",
	CLIFlagClearJournalUsage:        "διαγραφή τοπικού ημερολογίου για αρχείο ή κατάλογο και έξοδος",
	CLIFlagBookViewUsage:            "εμφάνιση ζευγών σελίδων από αριστερά προς τα δεξιά",
	CLIFlagRightBookViewUsage:       "εμφάνιση ζευγών σελίδων από δεξιά προς τα αριστερά",
	CLIFlagCircleBookViewUsage:      "εμφάνιση επικαλυπτόμενων ζευγών σελίδων από αριστερά προς τα δεξιά",
	CLIFlagRightCircleBookViewUsage: "εμφάνιση επικαλυπτόμενων ζευγών σελίδων από δεξιά προς τα αριστερά",
	CLIErrMultipleBookViews:         "μπορεί να χρησιμοποιηθεί μόνο μία επιλογή προβολής βιβλίου",
	CLIErrInvalidView:               "μη υποστηριζόμενη τιμή COMICREAD_VIEW %q (αναμένονται: book-view, right-view, circle-view ή right-circle-view)",
	CLIFlagOpenUsage:                "κατάλογος για άνοιγμα στον επιλογέα αρχείων (προεπιλογή: COMICREAD_DIR ή ο τρέχων κατάλογος)",
	CLIErrOpenNotDir:                "άνοιγμα καταλόγου %q: δεν είναι κατάλογος",
	CLIHelpHint:                     "εκτελέστε 'comicread --help' για βοήθεια",
	CLIUsage:                        "χρήση: comicread [επιλογές] [αρχείο]",
	CLIUsageFull: `comicread — ένας μινιμαλιστικός αναγνώστης manga για τερματικό

χρήση: comicread [επιλογές] [αρχείο]

επιλογές:
  --graphics string   αποδότης: auto, ascii, dots, kitty, sixel ή iterm2 (προεπιλογή "auto")
  --book-view         εμφάνιση ζευγών σελίδων από αριστερά προς τα δεξιά
  --right-view        εμφάνιση ζευγών σελίδων από δεξιά προς τα αριστερά
  --circle-view       εμφάνιση επικαλυπτόμενων ζευγών σελίδων από αριστερά προς τα δεξιά
  --right-circle-view
                      εμφάνιση επικαλυπτόμενων ζευγών σελίδων από δεξιά προς τα αριστερά
  --clear-journal    διαγραφή τοπικού ημερολογίου για αρχείο ή κατάλογο και έξοδος
  -o, --open string   κατάλογος για άνοιγμα στον επιλογέα αρχείων (προεπιλογή: COMICREAD_DIR ή ο τρέχων κατάλογος)
  --env               εμφάνιση περιβάλλοντος comicread και έξοδος
  --update            έλεγχος ενημερώσεων και έξοδος
  -v, --version       εμφάνιση έκδοσης και έξοδος
  -h, --help          εμφάνιση αυτής της βοήθειας

Αν δεν δοθεί αρχείο ή κατάλογος, ανοίγει διαδραστική επιλογή αρχείου στο COMICREAD_DIR
(αν έχει οριστεί σε έγκυρο κατάλογο) ή αλλιώς στον τρέχοντα κατάλογο.

περιβάλλον:
  COMICREAD_GRAPHICS  προεπιλεγμένος αποδότης: auto, ascii, dots, kitty, sixel ή iterm2
  COMICREAD_PRERENDERED_NEXT      επόμενες σελίδες για προαπόδοση (προεπιλογή 1)
  COMICREAD_PRERENDERED_PREVIOUS  προηγούμενες σελίδες για προαπόδοση (προεπιλογή 1)
  COMICREAD_VIEW      προεπιλεγμένη προβολή: book-view, right-view, circle-view ή right-circle-view
  COMICREAD_LANG      γλώσσα μηνυμάτων: en, uk, pl, de, fr, es, cs, ro, it, ko, ja, id, hi, el, tr, kk ή ka (προεπιλογή "en")
  COMICREAD_DIR       προεπιλεγμένος κατάλογος για τον επιλογέα αρχείων όταν δεν δίνεται διαδρομή`,
}
