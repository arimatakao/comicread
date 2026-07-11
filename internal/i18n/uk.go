package i18n

var ukMessages = map[string]string{
	ReaderStatusWaitingTerminalSize: "очікування розміру термінала",
	ReaderStatusTerminalTooSmall:    "вікно термінала замале",
	ReaderStatusLastPage:            "остання сторінка",
	ReaderStatusFirstPage:           "перша сторінка",
	ReaderStatusRenderError:         "помилка рендерингу: %v",
	ReaderStatusMaximumZoom:         "максимальне збільшення",
	ReaderStatusMinimumZoom:         "мінімальне збільшення",

	ReaderViewTerminalTooSmall: "comicread: вікно термінала замале",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "сторінки %d/%d",
	ReaderViewRendering:        "рендеринг",

	FilepickerHeader:      "comicread — оберіть розділ\n%s\n\n",
	FilepickerNoEntries:   "  (немає підтримуваних елементів)\n",
	FilepickerHelp:        "\n↑/↓ рух  |  ← батьківська тека  |  → відкрити теку  |  enter відкрити файл  |  s обрати виділену теку  |  q вихід\n",
	FilepickerWindowTitle: "comicread — вибір файлу",

	FilepickerErrResolveDir: "не вдалося визначити теку %q: %w",
	FilepickerErrReadDir:    "не вдалося прочитати теку %q: %w",
	FilepickerErrRunPicker:  "помилка запуску вибору файлу: %w",

	LoadingViewOpening:     "відкриваю %s…",
	LoadingViewWindowTitle: "comicread — відкриття",

	CLIErrGetWorkingDir:   "не вдалося отримати робочу теку: %w",
	CLIErrPickFile:        "не вдалося обрати файл: %w",
	CLIErrRunTUI:          "помилка запуску TUI: %w",
	CLIErrParseArgs:       "помилка розбору аргументів: %w",
	CLIErrOpenChapter:     "не вдалося відкрити розділ: %w",
	CLIErrNoPages:         "розділ не містить придатних для читання сторінок",
	CLIErrInspectInput:    "не вдалося перевірити вхідні дані %q: %w",
	CLIErrUnsupportedFile: "непідтримуваний файл %q: підтримувані формати — CBZ, PDF, EPUB або тека із зображеннями",
	CLIFlagGraphicsUsage:  "рендерер: auto, ascii, dots, kitty, sixel або iterm2",
	CLIHelpHint:           "виконайте 'comicread --help' для довідки",
	CLIUsage:              "використання: comicread [--graphics auto|ascii|dots|kitty|sixel|iterm2] [файл.cbz|файл.pdf|файл.epub|тека-із-зображеннями]",
	CLIUsageFull: `comicread - мінімалістична манга-читалка для термінала

використання: comicread [--graphics auto|ascii|dots|kitty|sixel|iterm2] [файл.cbz|файл.pdf|файл.epub|тека-із-зображеннями]

опції:
  --graphics string   рендерер: auto, ascii, dots, kitty, sixel або iterm2 (за замовчуванням "auto")
  -h, --help          показати цю довідку

Якщо файл або тека не вказані, відкриється інтерактивний вибір файлу в поточній теці.

змінні середовища:
  COMICREAD_LANG   мова повідомлень: en або uk (за замовчуванням "en")`,
}
