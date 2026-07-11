package i18n

var ukMessages = map[string]string{
	"reader.status.waiting_terminal_size": "очікування розміру термінала",
	"reader.status.terminal_too_small":    "вікно термінала замале",
	"reader.status.last_page":             "остання сторінка",
	"reader.status.first_page":            "перша сторінка",
	"reader.status.render_error":          "помилка рендерингу: %v",
	"reader.status.maximum_zoom":          "максимальне збільшення",
	"reader.status.minimum_zoom":          "мінімальне збільшення",

	"reader.view.terminal_too_small": "comicread: вікно термінала замале",
	"reader.view.window_title":       "comicread — %s",
	"reader.view.pages":              "сторінки %d/%d",
	"reader.view.rendering":          "рендеринг",

	"filepicker.header":       "comicread — оберіть розділ\n%s\n\n",
	"filepicker.no_entries":   "  (немає підтримуваних елементів)\n",
	"filepicker.help":         "\n↑/↓ рух  ← батьківська тека  → відкрити теку  enter відкрити файл  s обрати виділену теку  q вихід\n",
	"filepicker.window_title": "comicread — вибір файлу",

	"filepicker.err.resolve_dir": "не вдалося визначити теку %q: %w",
	"filepicker.err.read_dir":    "не вдалося прочитати теку %q: %w",
	"filepicker.err.run_picker":  "помилка запуску вибору файлу: %w",

	"cli.err.get_working_dir":  "не вдалося отримати робочу теку: %w",
	"cli.err.pick_file":        "не вдалося обрати файл: %w",
	"cli.err.run_tui":          "помилка запуску TUI: %w",
	"cli.err.parse_args":       "помилка розбору аргументів: %w",
	"cli.err.open_chapter":     "не вдалося відкрити розділ: %w",
	"cli.err.no_pages":         "розділ не містить придатних для читання сторінок",
	"cli.err.inspect_input":    "не вдалося перевірити вхідні дані %q: %w",
	"cli.err.unsupported_file": "непідтримуваний файл %q: підтримувані формати — CBZ, PDF, EPUB або тека із зображеннями",
	"cli.flag.graphics_usage":  "рендерер: auto, ascii, dots, kitty, sixel або iterm2",
	"cli.help_hint":            "виконайте 'comicread --help' для довідки",
	"cli.usage":                "використання: comicread [--graphics auto|ascii|dots|kitty|sixel|iterm2] [файл.cbz|файл.pdf|файл.epub|тека-із-зображеннями]",
	"cli.usage_full": `comicread - мінімалістична манга-читалка для термінала

використання: comicread [--graphics auto|ascii|dots|kitty|sixel|iterm2] [файл.cbz|файл.pdf|файл.epub|тека-із-зображеннями]

опції:
  --graphics string   рендерер: auto, ascii, dots, kitty, sixel або iterm2 (за замовчуванням "auto")
  -h, --help          показати цю довідку

Якщо файл або тека не вказані, відкриється інтерактивний вибір файлу в поточній теці.

змінні середовища:
  COMICREAD_LANG   мова повідомлень: en або uk (за замовчуванням "en")`,
}
