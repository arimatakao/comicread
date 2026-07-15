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
	ReaderViewPageRange:        "сторінки %d-%d/%d",
	ReaderViewRendering:        "рендеринг",
	ReaderViewBookmarks:        "Закладки",
	ReaderViewNoBookmarks:      "(немає закладок)",
	ReaderViewBookmarksHelp:    "вгору/вниз рух  enter відкрити  esc закрити",
	ReaderViewHelp: `Клавіші

← →  попередня / наступна сторінка
↑ ↓  прокрутка збільшеної сторінки
+ -  збільшити / зменшити
b    додати / видалити закладку
v ← → попередня / наступна закладка
c v  закладки
q    вихід

?    закрити довідку`,

	FilepickerHeader:      "comicread — оберіть розділ\n%s\n\n",
	FilepickerNoEntries:   "  (немає підтримуваних елементів)\n",
	FilepickerHelp:        "\n↑/↓ рух  |  ← батьківська тека  |  → відкрити теку  |  enter відкрити файл  |  s обрати виділену теку  |  q вихід\n",
	FilepickerWindowTitle: "comicread — вибір файлу",

	FilepickerErrResolveDir: "не вдалося визначити теку %q: %w",
	FilepickerErrReadDir:    "не вдалося прочитати теку %q: %w",
	FilepickerErrRunPicker:  "помилка запуску вибору файлу: %w",

	LoadingViewOpening:     "відкриваю %s…",
	LoadingViewWindowTitle: "comicread — відкриття",

	CLIErrGetWorkingDir:             "не вдалося отримати робочу теку: %w",
	CLIErrPickFile:                  "не вдалося обрати файл: %w",
	CLIErrRunTUI:                    "помилка запуску TUI: %w",
	CLIErrParseArgs:                 "помилка розбору аргументів: %w",
	CLIErrOpenChapter:               "не вдалося відкрити розділ: %w",
	CLIErrOpenJournal:               "не вдалося відкрити журнал: %w",
	CLIErrClearJournal:              "не вдалося очистити журнал: %w",
	CLIErrClearJournalRequiresInput: "--clear-journal потребує файл або теку",
	CLIErrNoPages:                   "розділ не містить придатних для читання сторінок",
	CLIErrInspectInput:              "не вдалося перевірити вхідні дані %q: %w",
	CLIErrUnsupportedFile:           "непідтримуваний файл %q: підтримувані формати — CBZ, PDF, EPUB або тека із зображеннями",
	CLIFlagGraphicsUsage:            "рендерер: auto, ascii, dots, kitty, sixel або iterm2",
	CLIFlagVersionUsage:             "вивести версію та завершити роботу",
	CLIFlagEnvUsage:                 "вивести середовище comicread та завершити роботу",
	CLIFlagClearJournalUsage:        "видалити локальний журнал для файлу або теки та завершити роботу",
	CLIFlagBookViewUsage:            "показувати сторінки парами зліва направо",
	CLIFlagRightBookViewUsage:       "показувати сторінки парами справа наліво",
	CLIFlagCircleBookViewUsage:      "показувати перекривні пари сторінок зліва направо",
	CLIFlagRightCircleBookViewUsage: "показувати перекривні пари сторінок справа наліво",
	CLIErrMultipleBookViews:         "можна вказати лише один режим книжкового перегляду",
	CLIErrInvalidView:               "непідтримуване значення COMICREAD_VIEW %q (можливі: book-view, right-view, circle-view або right-circle-view)",
	CLIHelpHint:                     "виконайте 'comicread --help' для довідки",
	CLIUsage:                        "використання: comicread [опції] [файл]",
	CLIUsageFull: `comicread - мінімалістична манга-читалка для термінала

використання: comicread [опції] [файл]

опції:
  --graphics string   рендерер: auto, ascii, dots, kitty, sixel або iterm2 (за замовчуванням "auto")
  --book-view         показувати сторінки парами зліва направо
  --right-view        показувати сторінки парами справа наліво
  --circle-view       показувати перекривні пари сторінок зліва направо
  --right-circle-view
                      показувати перекривні пари сторінок справа наліво
  --clear-journal    видалити локальний журнал для файлу або теки та завершити роботу
  --env               вивести середовище comicread та завершити роботу
  -v, --version       вивести версію та завершити роботу
  -h, --help          показати цю довідку

Якщо файл або тека не вказані, відкриється інтерактивний вибір файлу в поточній теці.

змінні середовища:
  COMICREAD_GRAPHICS  рендерер за замовчуванням: auto, ascii, dots, kitty, sixel або iterm2
  COMICREAD_VIEW      режим за замовчуванням: book-view, right-view, circle-view або right-circle-view
  COMICREAD_LANG   мова повідомлень: en, uk, pl, de, fr, es, cs, ro, it, ko, ja, id, hi, el, tr, kk або ka (за замовчуванням "en")`,
}
