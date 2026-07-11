// Package i18n provides minimal message translation for comicread's
// user-facing strings (status messages, view labels, CLI errors).
package i18n

import (
	"fmt"
	"os"
	"strings"
)

// Lang identifies a supported message language.
type Lang string

const (
	English   Lang = "en"
	Ukrainian Lang = "uk"
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

	ReaderViewTerminalTooSmall = "reader.view.terminal_too_small"
	ReaderViewWindowTitle      = "reader.view.window_title"
	ReaderViewPages            = "reader.view.pages"
	ReaderViewPageRange        = "reader.view.page_range"
	ReaderViewRendering        = "reader.view.rendering"

	FilepickerHeader        = "filepicker.header"
	FilepickerNoEntries     = "filepicker.no_entries"
	FilepickerHelp          = "filepicker.help"
	FilepickerWindowTitle   = "filepicker.window_title"
	FilepickerErrResolveDir = "filepicker.err.resolve_dir"
	FilepickerErrReadDir    = "filepicker.err.read_dir"
	FilepickerErrRunPicker  = "filepicker.err.run_picker"

	LoadingViewOpening     = "loading.view.opening"
	LoadingViewWindowTitle = "loading.view.window_title"

	CLIErrGetWorkingDir             = "cli.err.get_working_dir"
	CLIErrPickFile                  = "cli.err.pick_file"
	CLIErrRunTUI                    = "cli.err.run_tui"
	CLIErrParseArgs                 = "cli.err.parse_args"
	CLIErrOpenChapter               = "cli.err.open_chapter"
	CLIErrNoPages                   = "cli.err.no_pages"
	CLIErrInspectInput              = "cli.err.inspect_input"
	CLIErrUnsupportedFile           = "cli.err.unsupported_file"
	CLIFlagGraphicsUsage            = "cli.flag.graphics_usage"
	CLIFlagVersionUsage             = "cli.flag.version_usage"
	CLIFlagBookViewUsage            = "cli.flag.book_view_usage"
	CLIFlagRightBookViewUsage       = "cli.flag.right_book_view_usage"
	CLIFlagCircleBookViewUsage      = "cli.flag.circle_book_view_usage"
	CLIFlagRightCircleBookViewUsage = "cli.flag.right_circle_book_view_usage"
	CLIErrMultipleBookViews         = "cli.err.multiple_book_views"
	CLIHelpHint                     = "cli.help_hint"
	CLIUsage                        = "cli.usage"
	CLIUsageFull                    = "cli.usage_full"
)

var current = detect()

// detect picks a language from the COMICREAD_LANG environment variable,
// falling back to English.
func detect() Lang {
	if strings.HasPrefix(strings.ToLower(os.Getenv("COMICREAD_LANG")), "uk") {
		return Ukrainian
	}
	return English
}

// SetLang overrides the active language (used mainly for testing/tools).
func SetLang(lang Lang) {
	current = lang
}

// T returns the message for key in the active language, formatting it with
// args via fmt.Sprintf when args are given. It falls back to English, then
// to the key itself, if a translation is missing.
func T(key string, args ...any) string {
	msg, ok := messages[current][key]
	if !ok {
		msg, ok = messages[English][key]
	}
	if !ok {
		msg = key
	}
	if len(args) == 0 {
		return msg
	}
	return fmt.Sprintf(msg, args...)
}

var messages = map[Lang]map[string]string{
	English:   enMessages,
	Ukrainian: ukMessages,
}
