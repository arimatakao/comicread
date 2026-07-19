package cli

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	tea "charm.land/bubbletea/v2"
	"github.com/arimatakao/comicfile"
	"github.com/arimatakao/comicread/internal/backend"
	"github.com/arimatakao/comicread/internal/config"
	"github.com/arimatakao/comicread/internal/i18n"
	"github.com/arimatakao/comicread/internal/journal"
	"github.com/arimatakao/comicread/internal/tui/filepicker"
	"github.com/arimatakao/comicread/internal/tui/loading"
	"github.com/arimatakao/comicread/internal/tui/reader"
)

// ErrUsage marks errors caused by invalid command-line arguments, as opposed
// to runtime failures (missing files, decode errors, etc). Callers can check
// for it with errors.Is.
var ErrUsage = errors.New("usage error")

// Version is the program version, overridden at build time via
// -ldflags "-X github.com/arimatakao/comicread/internal/cli.Version=YOUR_VERSION".
var Version = "dev"

type usageError struct{ err error }

func (e *usageError) Error() string        { return e.err.Error() }
func (e *usageError) Unwrap() error        { return e.err }
func (e *usageError) Is(target error) bool { return target == ErrUsage }

func Run(args []string) error {
	settings, err := config.Load()
	if err != nil {
		return err
	}
	i18n.SetLang(i18n.Lang(settings.Language))

	options, err := parseOptionsWithConfig(args, settings)
	if err != nil {
		return err
	}

	if options.version {
		fmt.Println(Version)
		return nil
	}
	if options.update {
		return checkForUpdate(settings.Language)
	}

	// ctx is cancelled on SIGINT/SIGTERM so every stage below can shut down
	// gracefully: save reading progress, clear the image and close the chapter.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	path := options.path
	if path == "" {
		dir, err := pickerDir(options.open, options.openSet, settings.Directory)
		if err != nil {
			return err
		}
		path, err = filepicker.Pick(dir)
		if err != nil {
			if errors.Is(err, tea.ErrInterrupted) {
				return nil
			}
			return fmt.Errorf(i18n.T(i18n.CLIErrPickFile), err)
		}
		if path == "" {
			return nil
		}
	}

	if err := validateInput(path); err != nil {
		return err
	}
	if options.clearJournal {
		if err := journal.Clear(path); err != nil {
			return fmt.Errorf(i18n.T(i18n.CLIErrClearJournal), err)
		}
		return nil
	}

	chapter, err := loading.Open(path, func() (comicfile.ContainerReader, error) {
		return openChapter(path)
	})
	if err != nil {
		if errors.Is(err, tea.ErrInterrupted) {
			return nil
		}
		return err
	}
	if chapter == nil {
		// The loading screen was terminated by a signal before the chapter
		// finished opening; there is nothing to read or clean up.
		return nil
	}
	defer chapter.Close()

	renderer, err := backend.NewRenderer(options.graphics)
	if err != nil {
		return err
	}

	progress, err := journal.Open(path)
	if err != nil {
		return fmt.Errorf(i18n.T(i18n.CLIErrOpenJournal), err)
	}
	model := reader.NewWithBookViewAndJournalAndPreRender(filepath.Base(path), chapter, renderer, options.bookView, progress, settings.Prerender.Next, settings.Prerender.Previous)
	// Bubble Tea's own SIGINT/SIGTERM handler stops the event loop without
	// consulting the model, skipping progress saving and image cleanup. It is
	// disabled here and signals are routed through ctx instead, so an external
	// termination follows the same path as pressing the quit key.
	program := tea.NewProgram(model, tea.WithoutSignalHandler())
	go func() {
		<-ctx.Done()
		program.Send(reader.ExternalQuitMsg{})
	}()
	if _, err := program.Run(); err != nil {
		if errors.Is(err, tea.ErrInterrupted) {
			return nil
		}
		return fmt.Errorf(i18n.T(i18n.CLIErrRunTUI), err)
	}

	return nil
}

// pickerDir returns the directory the file picker should open initially.
// If -o/--open was given with a value, it is used and must be a valid,
// existing directory. If -o/--open was given with no value, the current
// working directory is used, bypassing the configured directory. Otherwise
// configuredDir is used when set to a valid, existing directory; failing that,
// the current working directory is used.
func pickerDir(open string, openSet bool, configuredDir string) (string, error) {
	if open != "" {
		info, err := os.Stat(open)
		if err != nil {
			return "", fmt.Errorf(i18n.T(i18n.CLIErrInspectInput), open, err)
		}
		if !info.IsDir() {
			return "", errors.New(i18n.T(i18n.CLIErrOpenNotDir, open))
		}
		return open, nil
	}
	if !openSet {
		if dir := configuredDir; dir != "" {
			if info, err := os.Stat(dir); err == nil && info.IsDir() {
				return dir, nil
			}
		}
	}
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf(i18n.T(i18n.CLIErrGetWorkingDir), err)
	}
	return cwd, nil
}

func parseArgs(args []string) (graphics, path string, version bool, err error) {
	options, err := parseOptions(args)
	return options.graphics, options.path, options.version, err
}

type options struct {
	graphics     string
	path         string
	open         string
	openSet      bool
	version      bool
	update       bool
	clearJournal bool
	bookView     reader.ViewMode
}

// normalizeOpenFlag rewrites a bare, value-less -o/--open into an explicit
// empty value (-o= / --open=) so the flag package accepts it instead of
// erroring with "flag needs an argument", and so it is still reported as
// set by flags.Visit.
func normalizeOpenFlag(args []string) []string {
	out := make([]string, 0, len(args))
	for i, arg := range args {
		if arg == "--" {
			out = append(out, args[i:]...)
			break
		}
		if (arg == "-o" || arg == "--open") && (i+1 >= len(args) || strings.HasPrefix(args[i+1], "-")) {
			out = append(out, arg+"=")
			continue
		}
		out = append(out, arg)
	}
	return out
}

func parseOptions(args []string) (options, error) {
	settings, err := config.Load()
	if err != nil {
		return options{}, err
	}
	return parseOptionsWithConfig(args, settings)
}

func parseOptionsWithConfig(args []string, settings config.Config) (options, error) {
	args = normalizeOpenFlag(args)

	flags := flag.NewFlagSet("comicread", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	graphicsFlag := flags.String("graphics", settings.Graphics, i18n.T(i18n.CLIFlagGraphicsUsage))
	versionFlag := flags.Bool("version", false, i18n.T(i18n.CLIFlagVersionUsage))
	updateFlag := flags.Bool("update", false, i18n.T(i18n.CLIFlagUpdateUsage))
	clearJournalFlag := flags.Bool("clear-journal", false, i18n.T(i18n.CLIFlagClearJournalUsage))
	bookViewFlag := flags.Bool("book-view", false, i18n.T(i18n.CLIFlagBookViewUsage))
	rightBookViewFlag := flags.Bool("right-view", false, i18n.T(i18n.CLIFlagRightBookViewUsage))
	circleBookViewFlag := flags.Bool("circle-view", false, i18n.T(i18n.CLIFlagCircleBookViewUsage))
	rightCircleBookViewFlag := flags.Bool("right-circle-view", false, i18n.T(i18n.CLIFlagRightCircleBookViewUsage))
	openFlag := flags.String("open", "", i18n.T(i18n.CLIFlagOpenUsage))
	flags.BoolVar(versionFlag, "v", false, i18n.T(i18n.CLIFlagVersionUsage))
	flags.StringVar(openFlag, "o", "", i18n.T(i18n.CLIFlagOpenUsage))
	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			fmt.Println(i18n.T(i18n.CLIUsageFull))
			return options{}, flag.ErrHelp
		}
		return options{}, &usageError{fmt.Errorf(i18n.T(i18n.CLIErrParseArgs), err)}
	}
	openSet := false
	flags.Visit(func(f *flag.Flag) {
		if f.Name == "open" || f.Name == "o" {
			openSet = true
		}
	})
	if *updateFlag {
		return options{update: true}, nil
	}
	bookView, err := selectedBookView(settings.View, *bookViewFlag, *rightBookViewFlag, *circleBookViewFlag, *rightCircleBookViewFlag)
	if err != nil {
		return options{}, &usageError{err}
	}
	if *versionFlag {
		return options{version: true}, nil
	}
	if *clearJournalFlag && flags.NArg() != 1 {
		return options{}, &usageError{errors.New(i18n.T(i18n.CLIErrClearJournalRequiresInput))}
	}
	switch flags.NArg() {
	case 0:
		return options{graphics: *graphicsFlag, open: *openFlag, openSet: openSet, bookView: bookView}, nil
	case 1:
		return options{graphics: *graphicsFlag, path: flags.Arg(0), clearJournal: *clearJournalFlag, bookView: bookView}, nil
	default:
		return options{}, &usageError{errors.New(i18n.T(i18n.CLIUsage))}
	}
}

func selectedBookView(view string, book, rightBook, circle, rightCircle bool) (reader.ViewMode, error) {
	selected := 0
	for _, enabled := range []bool{book, rightBook, circle, rightCircle} {
		if enabled {
			selected++
		}
	}
	if selected > 1 {
		return reader.SinglePageView, errors.New(i18n.T(i18n.CLIErrMultipleBookViews))
	}
	switch {
	case book:
		return reader.BookView, nil
	case rightBook:
		return reader.RightBookView, nil
	case circle:
		return reader.CircleBookView, nil
	case rightCircle:
		return reader.RightCircleBookView, nil
	}

	switch strings.ToLower(strings.TrimSpace(view)) {
	case "", "single-page":
		return reader.SinglePageView, nil
	case "book-view":
		return reader.BookView, nil
	case "right-view":
		return reader.RightBookView, nil
	case "circle-view":
		return reader.CircleBookView, nil
	case "right-circle-view":
		return reader.RightCircleBookView, nil
	default:
		return reader.SinglePageView, errors.New(i18n.T(i18n.CLIErrInvalidView, view))
	}
}

func openChapter(path string) (comicfile.ContainerReader, error) {
	chapter, err := comicfile.OpenContainer(path)
	if err != nil {
		return nil, fmt.Errorf(i18n.T(i18n.CLIErrOpenChapter), err)
	}
	if chapter.TotalPages() == 0 {
		_ = chapter.Close()
		return nil, errors.New(i18n.T(i18n.CLIErrNoPages))
	}
	return chapter, nil
}

func isSupportedFile(path string) bool {
	switch strings.ToLower(filepath.Ext(path)) {
	case ".cbz", ".pdf", ".epub":
		return true
	default:
		return false
	}
}

func validateInput(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf(i18n.T(i18n.CLIErrInspectInput), path, err)
	}
	if info.IsDir() || isSupportedFile(path) {
		return nil
	}
	return fmt.Errorf(i18n.T(i18n.CLIErrUnsupportedFile), path)
}
