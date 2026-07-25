// Package i18n provides minimal message translation for comicread's
// user-facing strings (status messages, view labels, CLI errors).
package i18n

import (
	"fmt"
)

// Lang identifies a supported message language.
type Lang string

const (
	English    Lang = "en"
	Ukrainian  Lang = "uk"
	Polish     Lang = "pl"
	German     Lang = "de"
	French     Lang = "fr"
	Spanish    Lang = "es"
	Czech      Lang = "cs"
	Romanian   Lang = "ro"
	Italian    Lang = "it"
	Korean     Lang = "ko"
	Japanese   Lang = "ja"
	Indonesian Lang = "id"
	Hindi      Lang = "hi"
	Greek      Lang = "el"
	Turkish    Lang = "tr"
	Kazakh     Lang = "kk"
	Georgian   Lang = "ka"
	Hungarian  Lang = "hu"
	Swedish    Lang = "sv"
	Norwegian  Lang = "no"
	Danish     Lang = "da"
	Finnish    Lang = "fi"
)

// Message keys.
const (
	ReaderStatusWaitingTerminalSize = "reader.status.waiting_terminal_size"
	ReaderStatusTerminalTooSmall    = "reader.status.terminal_too_small"
	ReaderStatusLastPage            = "reader.status.last_page"
	ReaderStatusFirstPage           = "reader.status.first_page"
	ReaderStatusRenderError         = "reader.status.render_error"
	ReaderStatusMaximumZoom         = "reader.status.maximum_zoom"
	ReaderStatusMinimumZoom         = "reader.status.minimum_zoom"
	ReaderStatusInvalidPage         = "reader.status.invalid_page"

	ReaderViewTerminalTooSmall = "reader.view.terminal_too_small"
	ReaderViewWindowTitle      = "reader.view.window_title"
	ReaderViewPages            = "reader.view.pages"
	ReaderViewPageRange        = "reader.view.page_range"
	ReaderViewRendering        = "reader.view.rendering"
	ReaderViewHelp             = "reader.view.help"
	ReaderViewBookmarks        = "reader.view.bookmarks"
	ReaderViewNoBookmarks      = "reader.view.no_bookmarks"
	ReaderViewBookmarksHelp    = "reader.view.bookmarks_help"
	ReaderViewGoToPage         = "reader.view.go_to_page"

	FilepickerHeader         = "filepicker.header"
	FilepickerNoEntries      = "filepicker.no_entries"
	FilepickerHelp           = "filepicker.help"
	FilepickerWindowTitle    = "filepicker.window_title"
	FilepickerGoToPrompt     = "filepicker.go_to_prompt"
	FilepickerFavoritePrompt = "filepicker.favorite_prompt"
	FilepickerGoToErr        = "filepicker.go_to_err"
	FilepickerFavorites      = "filepicker.favorites"
	FilepickerNoFavorites    = "filepicker.no_favorites"
	FilepickerFavoritesHelp  = "filepicker.favorites_help"
	FilepickerFavoriteErr    = "filepicker.favorite_err"
	FilepickerErrResolveDir  = "filepicker.err.resolve_dir"
	FilepickerErrReadDir     = "filepicker.err.read_dir"
	FilepickerErrRunPicker   = "filepicker.err.run_picker"
	FilepickerErrEmptyPath   = "filepicker.err.empty_path"
	FilepickerErrNotDir      = "filepicker.err.not_dir"

	LoadingViewOpening     = "loading.view.opening"
	LoadingViewWindowTitle = "loading.view.window_title"

	CLIErrGetWorkingDir             = "cli.err.get_working_dir"
	CLIErrPickFile                  = "cli.err.pick_file"
	CLIErrRunTUI                    = "cli.err.run_tui"
	CLIErrParseArgs                 = "cli.err.parse_args"
	CLIErrOpenChapter               = "cli.err.open_chapter"
	CLIErrOpenJournal               = "cli.err.open_journal"
	CLIErrClearJournal              = "cli.err.clear_journal"
	CLIErrClearJournalRequiresInput = "cli.err.clear_journal_requires_input"
	CLIErrNoPages                   = "cli.err.no_pages"
	CLIErrInspectInput              = "cli.err.inspect_input"
	CLIErrUnsupportedFile           = "cli.err.unsupported_file"
	CLIFlagGraphicsUsage            = "cli.flag.graphics_usage"
	CLIFlagConfigUsage              = "cli.flag.config_usage"
	CLIFlagResetConfigUsage         = "cli.flag.reset_config_usage"
	CLIFlagSetConfigUsage           = "cli.flag.set_config_usage"
	CLIFlagVersionUsage             = "cli.flag.version_usage"
	CLIFlagUpdateUsage              = "cli.flag.update_usage"
	CLIFlagEnvUsage                 = "cli.flag.env_usage"
	CLIFlagClearJournalUsage        = "cli.flag.clear_journal_usage"
	CLIFlagBookViewUsage            = "cli.flag.book_view_usage"
	CLIFlagRightBookViewUsage       = "cli.flag.right_book_view_usage"
	CLIFlagCircleBookViewUsage      = "cli.flag.circle_book_view_usage"
	CLIFlagRightCircleBookViewUsage = "cli.flag.right_circle_book_view_usage"
	CLIErrMultipleBookViews         = "cli.err.multiple_book_views"
	CLIErrInvalidView               = "cli.err.invalid_view"
	CLIFlagOpenUsage                = "cli.flag.open_usage"
	CLIErrOpenNotDir                = "cli.err.open_not_dir"
	CLIHelpHint                     = "cli.help_hint"
	CLIUsage                        = "cli.usage"
	CLIUsageFull                    = "cli.usage_full"
)

var currentMessages = messagesFor(English)

// SetLang overrides the active language selected from config.toml.
func SetLang(lang Lang) {
	currentMessages = messagesFor(lang)
}

// T returns the message for key in the active language, formatting it with
// args via fmt.Sprintf when args are given. It falls back to English, then
// to the key itself, if a translation is missing.
func T(key string, args ...any) string {
	msg, ok := currentMessages[key]
	if !ok {
		msg, ok = enMessages[key]
	}
	if !ok {
		msg = key
	}
	if len(args) == 0 {
		return msg
	}
	return fmt.Sprintf(msg, args...)
}

// messagesFor returns the translation catalog for lang. A switch keeps the
// catalogs explicit without an additional language-to-catalog map.
func messagesFor(lang Lang) map[string]string {
	var messages map[string]string
	switch lang {
	case Ukrainian:
		messages = ukMessages
	case Polish:
		messages = plMessages
	case German:
		messages = deMessages
	case French:
		messages = frMessages
	case Spanish:
		messages = esMessages
	case Czech:
		messages = csMessages
	case Romanian:
		messages = roMessages
	case Italian:
		messages = itMessages
	case Korean:
		messages = koMessages
	case Japanese:
		messages = jaMessages
	case Indonesian:
		messages = idMessages
	case Hindi:
		messages = hiMessages
	case Greek:
		messages = elMessages
	case Turkish:
		messages = trMessages
	case Kazakh:
		messages = kkMessages
	case Georgian:
		messages = kaMessages
	case Hungarian:
		messages = huMessages
	case Swedish:
		messages = svMessages
	case Norwegian:
		messages = noMessages
	case Danish:
		messages = daMessages
	case Finnish:
		messages = fiMessages
	default:
		messages = enMessages
	}
	return messages
}
