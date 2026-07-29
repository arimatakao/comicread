package i18n

var ukMessages = map[string]string{
	CLIFlagConfigUsage: "файл конфігурації для використання", CLIFlagResetConfigUsage: "скинути config.toml до значень за замовчуванням і завершити роботу", CLIFlagSetConfigUsage: "оновити config.toml: ключ=значення",
	ReaderStatusWaitingTerminalSize: "очікування розміру термінала",
	ReaderStatusTerminalTooSmall:    "вікно термінала замале",
	ReaderStatusLastPage:            "остання сторінка",
	ReaderStatusFirstPage:           "перша сторінка",
	ReaderStatusRenderError:         "помилка рендерингу: %v",
	ReaderStatusMaximumZoom:         "максимальне збільшення",
	ReaderStatusMinimumZoom:         "мінімальне збільшення",
	ReaderStatusInvalidPage:         "неправильний номер сторінки",

	ReaderViewTerminalTooSmall: "comicread: вікно термінала замале",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "сторінки %d/%d",
	ReaderViewPageRange:        "сторінки %d-%d/%d",
	ReaderViewRendering:        "рендеринг",
	ReaderViewBookmarks:        "Закладки",
	ReaderViewNoBookmarks:      "(немає закладок)",
	ReaderViewBookmarksHelp:    "вгору/вниз рух | enter відкрити | esc закрити",
	ReaderViewGoToPage:         "Перейти до сторінки: %s",
	ReaderViewHelp: `Клавіші

← →  попередня / наступна сторінка
↑ ↓  прокрутка збільшеної сторінки
+ -  збільшити / зменшити
b    додати / видалити закладку
v ← → попередня / наступна закладка
c v  закладки
g123 enter  перейти до сторінки
q    вихід

?    закрити довідку`,

	FilepickerHeader:         "comicread — оберіть розділ\n%s\n\n",
	FilepickerNoEntries:      "  (немає підтримуваних елементів)\n",
	FilepickerHelp:           "\n↑/↓ рух\n← батьківська тека\n→ відкрити теку\nenter відкрити файл\ns обрати виділену теку\nf додати / прибрати поточну теку з улюблених\nF додати улюблену теку\nb улюблені теки\no перейти до теки\nq вихід\n",
	FilepickerWindowTitle:    "comicread — вибір файлу",
	FilepickerGoToPrompt:     "\nПерейти до теки: %s\n",
	FilepickerFavoritePrompt: "\nУлюблена тека: %s\n",
	FilepickerGoToErr:        "  помилка: %s\n",
	FilepickerFavorites:      "Улюблені теки\n\n",
	FilepickerNoFavorites:    "  (немає налаштованих улюблених тек)\n",
	FilepickerFavoritesHelp:  "\n↑/↓ рух\nenter перейти до теки\nd прибрати з улюблених\nesc повернутися\n",
	FilepickerFavoriteErr:    "  помилка збереження улюблених: %s\n",

	FilepickerErrResolveDir: "не вдалося визначити теку %q: %w",
	FilepickerErrReadDir:    "не вдалося прочитати теку %q: %w",
	FilepickerErrRunPicker:  "помилка запуску вибору файлу: %w",
	FilepickerErrEmptyPath:  "шлях порожній",
	FilepickerErrNotDir:     "%q не є текою",

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
	CLIFlagUpdateUsage:              "перевірити оновлення та завершити роботу",
	CLIFlagEnvUsage:                 "вивести середовище comicread та завершити роботу",
	CLIFlagClearJournalUsage:        "видалити локальний журнал для файлу або теки та завершити роботу",
	CLIFlagBookViewUsage:            "показувати сторінки парами зліва направо",
	CLIFlagRightBookViewUsage:       "показувати сторінки парами справа наліво",
	CLIFlagCircleBookViewUsage:      "показувати перекривні пари сторінок зліва направо",
	CLIFlagRightCircleBookViewUsage: "показувати перекривні пари сторінок справа наліво",
	CLIErrMultipleBookViews:         "можна вказати лише один режим книжкового перегляду",
	CLIErrInvalidView:               "непідтримуване значення COMICREAD_VIEW %q (можливі: book-view, right-view, circle-view або right-circle-view)",
	CLIFlagOpenUsage:                "тека для відкриття у виборі файлів (за замовчуванням: COMICREAD_DIR або поточна тека)",
	CLIErrOpenNotDir:                "відкрити теку %q: не є текою",
	CLIFlagWebUsage:                 "запустити локальну вебчиталку замість інтерфейсу термінала",
	CLIErrWebArgs:                   "--web не приймає аргумент файлу або теки",
	WebServerStarted:                "вебчиталка comicread працює за адресою %s (натисніть Ctrl+C, щоб зупинити)",
	WebErrListen:                    "запуск вебсервера: %w",
	WebErrServe:                     "робота вебсервера: %w",
	CLIHelpHint:                     "виконайте 'comicread --help' для довідки",
	CLIUsage:                        "використання: comicread [опції] [файл]",
	CLIUsageFull: `comicread - мінімалістична манга-читалка для термінала

використання: comicread [опції] [файл]

опції:
  --config string     файл конфігурації для використання
  --graphics string   рендерер: auto, ascii, dots, kitty, sixel або iterm2 (за замовчуванням "auto")
  --book-view         показувати сторінки парами зліва направо
  --right-view        показувати сторінки парами справа наліво
  --circle-view       показувати перекривні пари сторінок зліва направо
  --right-circle-view
                      показувати перекривні пари сторінок справа наліво
  --clear-journal    видалити локальний журнал для файлу або теки та завершити роботу
  --reset-config     скинути config.toml до значень за замовчуванням і завершити роботу
  --set-config value оновити config.toml: ключ=значення
  -o, --open string   тека для відкриття у виборі файлів (за замовчуванням: COMICREAD_DIR або поточна тека)
  --update            перевірити оновлення та завершити роботу
  --web               запустити локальну вебчиталку замість інтерфейсу термінала
  -v, --version       вивести версію та завершити роботу
  -h, --help          показати цю довідку

Якщо файл або тека не вказані, відкриється інтерактивний вибір файлу в COMICREAD_DIR
(якщо задано коректну теку) або поточній теці.`,
	ReaderViewMetadata: "Метадані",
}
