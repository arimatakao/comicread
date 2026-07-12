package i18n

var kkMessages = map[string]string{
	ReaderStatusWaitingTerminalSize: "терминал өлшемі күтілуде",
	ReaderStatusTerminalTooSmall:    "терминал терезесі тым кіші",
	ReaderStatusLastPage:            "соңғы бет",
	ReaderStatusFirstPage:           "бірінші бет",
	ReaderStatusRenderError:         "рендеринг қатесі: %v",
	ReaderStatusMaximumZoom:         "ең үлкен масштаб",
	ReaderStatusMinimumZoom:         "ең кіші масштаб",

	ReaderViewTerminalTooSmall: "comicread: терминал терезесі тым кіші",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "беттер %d/%d",
	ReaderViewPageRange:        "беттер %d-%d/%d",
	ReaderViewRendering:        "рендеринг",
	ReaderViewHelp: `Пернелер

← →  алдыңғы / келесі бет
↑ ↓  ұлғайтылған бетті айналдыру
+ -  үлкейту / кішірейту
q    шығу

?    анықтаманы жабу`,

	FilepickerHeader:      "comicread — тарауды таңдаңыз\n%s\n\n",
	FilepickerNoEntries:   "  (қолдау көрсетілетін элементтер жоқ)\n",
	FilepickerHelp:        "\n↑/↓ жылжу  |  ← аталық қалта  |  → қалтаны ашу  |  enter файлды ашу  |  s таңдалған қалтаны таңдау  |  q шығу\n",
	FilepickerWindowTitle: "comicread — файл таңдау",

	FilepickerErrResolveDir: "%q қалтасын анықтау мүмкін болмады: %w",
	FilepickerErrReadDir:    "%q қалтасын оқу мүмкін болмады: %w",
	FilepickerErrRunPicker:  "файл таңдау қатесі: %w",

	LoadingViewOpening:     "%s ашылуда…",
	LoadingViewWindowTitle: "comicread — ашылуда",

	CLIErrGetWorkingDir:             "жұмыс қалтасын алу мүмкін болмады: %w",
	CLIErrPickFile:                  "файлды таңдау мүмкін болмады: %w",
	CLIErrRunTUI:                    "TUI іске қосу қатесі: %w",
	CLIErrParseArgs:                 "аргументтерді талдау қатесі: %w",
	CLIErrOpenChapter:               "тарауды ашу мүмкін болмады: %w",
	CLIErrNoPages:                   "тарауда оқуға жарамды беттер жоқ",
	CLIErrInspectInput:              "%q кіріс деректерін тексеру мүмкін болмады: %w",
	CLIErrUnsupportedFile:           "қолдау көрсетілмейтін файл %q: қолдау көрсетілетін пішімдер — CBZ, PDF, EPUB немесе сурет қалтасы",
	CLIFlagGraphicsUsage:            "рендерер: auto, ascii, dots, kitty, sixel немесе iterm2",
	CLIFlagVersionUsage:             "нұсқасын шығарып, жұмысты аяқтау",
	CLIFlagEnvUsage:                 "comicread ортасын шығарып, жұмысты аяқтау",
	CLIFlagBookViewUsage:            "беттерді солдан оңға жұп етіп көрсету",
	CLIFlagRightBookViewUsage:       "беттерді оңнан солға жұп етіп көрсету",
	CLIFlagCircleBookViewUsage:      "беттерді солдан оңға қабаттасқан жұп етіп көрсету",
	CLIFlagRightCircleBookViewUsage: "беттерді оңнан солға қабаттасқан жұп етіп көрсету",
	CLIErrMultipleBookViews:         "тек бір кітап көрінісі режимін көрсетуге болады",
	CLIErrInvalidView:               "COMICREAD_VIEW %q мәні қолдау көрсетілмейді (мүмкін мәндер: book-view, right-view, circle-view немесе right-circle-view)",
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
  --env               comicread ортасын шығарып, жұмысты аяқтау
  -v, --version       нұсқасын шығарып, жұмысты аяқтау
  -h, --help          осы анықтаманы көрсету

Файл немесе қалта көрсетілмесе, ағымдағы қалтада интерактивті файл таңдау ашылады.

орта айнымалылары:
  COMICREAD_GRAPHICS  әдепкі рендерер: auto, ascii, dots, kitty, sixel немесе iterm2
  COMICREAD_VIEW      әдепкі режим: book-view, right-view, circle-view немесе right-circle-view
  COMICREAD_LANG   хабарлама тілі: en, uk, pl, de, fr, es, cs, ro, it, ko, ja, id, hi, el, tr, kk немесе ka (әдепкі "en")`,
}
