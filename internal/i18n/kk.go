package i18n

var kkMessages = map[string]string{
	CLIFlagConfigUsage: "пайдаланылатын конфигурация файлы", CLIFlagResetConfigUsage: "config.toml әдепкі мәндерін қалпына келтіріп, шығу", CLIFlagSetConfigUsage: "config.toml жаңарту: кілт=мән",
	ReaderStatusWaitingTerminalSize: "терминал өлшемі күтілуде",
	ReaderStatusTerminalTooSmall:    "терминал терезесі тым кіші",
	ReaderStatusLastPage:            "соңғы бет",
	ReaderStatusFirstPage:           "бірінші бет",
	ReaderStatusRenderError:         "рендеринг қатесі: %v",
	ReaderStatusMaximumZoom:         "ең үлкен масштаб",
	ReaderStatusMinimumZoom:         "ең кіші масштаб",
	ReaderStatusInvalidPage:         "бет нөмірі жарамсыз",

	ReaderViewTerminalTooSmall: "comicread: терминал терезесі тым кіші",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "беттер %d/%d",
	ReaderViewPageRange:        "беттер %d-%d/%d",
	ReaderViewRendering:        "рендеринг",
	ReaderViewBookmarks:        "Бетбелгілер",
	ReaderViewNoBookmarks:      "(бетбелгілер жоқ)",
	ReaderViewBookmarksHelp:    "↑/↓ жылжу | enter ашу | esc жабу",
	ReaderViewGoToPage:         "Бетке өту: %s",
	ReaderViewHelp: `Пернелер

← →  алдыңғы / келесі бет
↑ ↓  ұлғайтылған бетті айналдыру
+ -  үлкейту / кішірейту
b    бетбелгі қосу / жою
v ← → алдыңғы / келесі бетбелгі
c v  бетбелгілер
g123 enter  бетке өту
q    шығу

?    анықтаманы жабу`,

	FilepickerHeader:         "comicread — тарауды таңдаңыз\n%s\n\n",
	FilepickerNoEntries:      "  (қолдау көрсетілетін элементтер жоқ)\n",
	FilepickerHelp:           "\n↑/↓ жылжу\n← аталық қалта\n→ қалтаны ашу\nenter файлды ашу\ns таңдалған қалтаны таңдау\nf ағымдағы қалтаны таңдаулыларға қосу / жою\nF таңдаулы қалтаны қосу\nb таңдаулы қалталар\no қалтаға өту\nq шығу\n",
	FilepickerWindowTitle:    "comicread — файл таңдау",
	FilepickerGoToPrompt:     "\nҚалтаға өту: %s\n",
	FilepickerFavoritePrompt: "\nТаңдаулы қалта: %s\n",
	FilepickerFavorites:      "Таңдаулы қалталар\n\n",
	FilepickerNoFavorites:    "  (бапталған таңдаулы қалталар жоқ)\n",
	FilepickerFavoritesHelp:  "\n↑/↓ жылжу\nenter қалтаға өту\nd таңдаулыны жою\nesc оралу\n",
	FilepickerFavoriteErr:    "  таңдаулыларды сақтау қатесі: %s\n",
	FilepickerGoToErr:        "  қате: %s\n",

	FilepickerErrResolveDir: "%q қалтасын анықтау мүмкін болмады: %w",
	FilepickerErrReadDir:    "%q қалтасын оқу мүмкін болмады: %w",
	FilepickerErrRunPicker:  "файл таңдау қатесі: %w",
	FilepickerErrEmptyPath:  "жол бос",
	FilepickerErrNotDir:     "%q қалта емес",

	LoadingViewOpening:     "%s ашылуда…",
	LoadingViewWindowTitle: "comicread — ашылуда",

	CLIErrGetWorkingDir:             "жұмыс қалтасын алу мүмкін болмады: %w",
	CLIErrPickFile:                  "файлды таңдау мүмкін болмады: %w",
	CLIErrRunTUI:                    "TUI іске қосу қатесі: %w",
	CLIErrParseArgs:                 "аргументтерді талдау қатесі: %w",
	CLIErrOpenChapter:               "тарауды ашу мүмкін болмады: %w",
	CLIErrOpenJournal:               "журналды ашу мүмкін болмады: %w",
	CLIErrClearJournal:              "журналды жою мүмкін болмады: %w",
	CLIErrClearJournalRequiresInput: "--clear-journal файлды немесе қалтаны қажет етеді",
	CLIErrNoPages:                   "тарауда оқуға жарамды беттер жоқ",
	CLIErrInspectInput:              "%q кіріс деректерін тексеру мүмкін болмады: %w",
	CLIErrUnsupportedFile:           "қолдау көрсетілмейтін файл %q: қолдау көрсетілетін пішімдер — CBZ, PDF, EPUB немесе сурет қалтасы",
	CLIFlagGraphicsUsage:            "рендерер: auto, ascii, dots, kitty, sixel немесе iterm2",
	CLIFlagVersionUsage:             "нұсқасын шығарып, жұмысты аяқтау",
	CLIFlagUpdateUsage:              "жаңартуларды тексеріп, жұмысты аяқтау",
	CLIFlagEnvUsage:                 "comicread ортасын шығарып, жұмысты аяқтау",
	CLIFlagClearJournalUsage:        "файл не қалта үшін жергілікті журналды жойып, жұмысты аяқтау",
	CLIFlagBookViewUsage:            "беттерді солдан оңға жұп етіп көрсету",
	CLIFlagRightBookViewUsage:       "беттерді оңнан солға жұп етіп көрсету",
	CLIFlagCircleBookViewUsage:      "беттерді солдан оңға қабаттасқан жұп етіп көрсету",
	CLIFlagRightCircleBookViewUsage: "беттерді оңнан солға қабаттасқан жұп етіп көрсету",
	CLIErrMultipleBookViews:         "тек бір кітап көрінісі режимін көрсетуге болады",
	CLIErrInvalidView:               "COMICREAD_VIEW %q мәні қолдау көрсетілмейді (мүмкін мәндер: book-view, right-view, circle-view немесе right-circle-view)",
	CLIFlagOpenUsage:                "файл таңдағышта ашылатын қалта (әдепкі: COMICREAD_DIR немесе ағымдағы қалта)",
	CLIErrOpenNotDir:                "қалтаны ашу %q: қалта емес",
	CLIHelpHint:                     "анықтама алу үшін 'comicread --help' орындаңыз",
	CLIUsage:                        "қолданылуы: comicread [опциялар] [файл]",
	CLIUsageFull: `comicread - терминалға арналған минималистік манга оқырманы

қолданылуы: comicread [опциялар] [файл]

опциялар:
  --graphics string   рендерер: auto, ascii, dots, kitty, sixel немесе iterm2 (әдепкі "auto")
  --book-view         беттерді солдан оңға жұп етіп көрсету
  --right-view        беттерді оңнан солға жұп етіп көрсету
  --circle-view       беттерді солдан оңға қабаттасқан жұп етіп көрсету
  --right-circle-view
                      беттерді оңнан солға қабаттасқан жұп етіп көрсету
  --clear-journal    файл не қалта үшін жергілікті журналды жойып, жұмысты аяқтау
  -o, --open string   файл таңдағышта ашылатын қалта (әдепкі: COMICREAD_DIR немесе ағымдағы қалта)
  --env               comicread ортасын шығарып, жұмысты аяқтау
  --update            жаңартуларды тексеріп, жұмысты аяқтау
  -v, --version       нұсқасын шығарып, жұмысты аяқтау
  -h, --help          осы анықтаманы көрсету

Файл немесе қалта көрсетілмесе, интерактивті файл таңдау COMICREAD_DIR ішінде ашылады
(егер ол дұрыс қалтаға орнатылса) немесе ағымдағы қалтада ашылады.

орта айнымалылары:
  COMICREAD_GRAPHICS  әдепкі рендерер: auto, ascii, dots, kitty, sixel немесе iterm2
  COMICREAD_PRERENDERED_NEXT      алдын ала көрсетуге келесі беттер саны (әдепкі 1)
  COMICREAD_PRERENDERED_PREVIOUS  алдын ала көрсетуге алдыңғы беттер саны (әдепкі 1)
  COMICREAD_VIEW      әдепкі режим: book-view, right-view, circle-view немесе right-circle-view
  COMICREAD_LANG   хабарлама тілі: https://github.com/arimatakao/comicread#environment-variables (әдепкі "en")
  COMICREAD_DIR    жол көрсетілмегенде файл таңдағышқа арналған әдепкі қалта`,
	ReaderViewMetadata: "Метадеректер",
}
