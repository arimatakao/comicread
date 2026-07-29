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
	"github.com/arimatakao/comicread/internal/web"
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
	// Parse once with the default config only to locate a possible custom
	// config file. Excluding help lets us load the selected language before
	// the final parse renders its localized output.
	initial, err := parseOptionsWithConfig(withoutHelp(args), config.Default())
	if err != nil {
		return err
	}
	if initial.resetConfig {
		if initial.configPathSet {
			return config.ResetFile(initial.configPath)
		}
		return config.Reset()
	}

	settings, err := loadConfig(initial.configPath, initial.configPathSet)
	if err != nil {
		return err
	}
	if initial.setConfigSet {
		if err := config.SetOption(&settings, initial.setConfig); err != nil {
			return err
		}
		if initial.configPathSet {
			return config.SaveFile(initial.configPath, settings)
		}
		return config.Save(settings)
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

	if options.web {
		return web.Serve(ctx, settings.Web.Port)
	}

	path := options.path
	if path == "" {
		dir, err := pickerDir(options.open, options.openSet, settings.Directory)
		if err != nil {
			return err
		}
		saveFavorites := func(favorites []string) error {
			settings.Favorites = favorites
			if options.configPathSet {
				return config.SaveFile(options.configPath, settings)
			}
			return config.Save(settings)
		}
		path, err = filepicker.PickWithFavoriteSaver(dir, settings.Favorites, saveFavorites)
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
	graphics      string
	path          string
	open          string
	openSet       bool
	version       bool
	update        bool
	resetConfig   bool
	setConfig     string
	setConfigSet  bool
	configPath    string
	configPathSet bool
	clearJournal  bool
	bookView      reader.ViewMode
	web           bool
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
	initial, err := parseOptionsWithConfig(withoutHelp(args), config.Default())
	if err != nil {
		return options{}, err
	}
	settings, err := loadConfig(initial.configPath, initial.configPathSet)
	if err != nil {
		return options{}, err
	}
	i18n.SetLang(i18n.Lang(settings.Language))
	return parseOptionsWithConfig(args, settings)
}

// withoutHelp removes help flags that would otherwise stop the preliminary
// parse before the configured language can be loaded. Flags following -- are
// positional arguments and must remain untouched.
func withoutHelp(args []string) []string {
	filtered := make([]string, 0, len(args))
	for i, arg := range args {
		if arg == "--" {
			filtered = append(filtered, args[i:]...)
			break
		}
		if arg == "-h" || arg == "--help" {
			continue
		}
		filtered = append(filtered, arg)
	}
	return filtered
}

func loadConfig(path string, custom bool) (config.Config, error) {
	if custom {
		return config.LoadFile(path)
	}
	return config.Load()
}

func parseOptionsWithConfig(args []string, settings config.Config) (options, error) {
	args = normalizeOpenFlag(args)

	flags := flag.NewFlagSet("comicread", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	graphicsFlag := flags.String("graphics", settings.Graphics, i18n.T(i18n.CLIFlagGraphicsUsage))
	versionFlag := flags.Bool("version", false, i18n.T(i18n.CLIFlagVersionUsage))
	updateFlag := flags.Bool("update", false, i18n.T(i18n.CLIFlagUpdateUsage))
	resetConfigFlag := flags.Bool("reset-config", false, i18n.T(i18n.CLIFlagResetConfigUsage))
	setConfigFlag := flags.String("set-config", "", i18n.T(i18n.CLIFlagSetConfigUsage))
	configPathFlag := flags.String("config", "", i18n.T(i18n.CLIFlagConfigUsage))
	clearJournalFlag := flags.Bool("clear-journal", false, i18n.T(i18n.CLIFlagClearJournalUsage))
	bookViewFlag := flags.Bool("book-view", false, i18n.T(i18n.CLIFlagBookViewUsage))
	rightBookViewFlag := flags.Bool("right-view", false, i18n.T(i18n.CLIFlagRightBookViewUsage))
	circleBookViewFlag := flags.Bool("circle-view", false, i18n.T(i18n.CLIFlagCircleBookViewUsage))
	rightCircleBookViewFlag := flags.Bool("right-circle-view", false, i18n.T(i18n.CLIFlagRightCircleBookViewUsage))
	openFlag := flags.String("open", "", i18n.T(i18n.CLIFlagOpenUsage))
	webFlag := flags.Bool("web", false, i18n.T(i18n.CLIFlagWebUsage))
	flags.BoolVar(versionFlag, "v", false, i18n.T(i18n.CLIFlagVersionUsage))
	flags.StringVar(openFlag, "o", "", i18n.T(i18n.CLIFlagOpenUsage))
	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			fmt.Println(i18n.T(i18n.CLIUsage))
			fmt.Println()
			flags.SetOutput(os.Stdout)
			flags.PrintDefaults()
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
	setConfigSet := false
	configPathSet := false
	flags.Visit(func(f *flag.Flag) {
		if f.Name == "set-config" {
			setConfigSet = true
		}
		if f.Name == "config" {
			configPathSet = true
		}
	})
	if configPathSet && strings.TrimSpace(*configPathFlag) == "" {
		return options{}, &usageError{errors.New("--config requires a path")}
	}
	if *resetConfigFlag && setConfigSet {
		return options{}, &usageError{errors.New("--reset-config and --set-config cannot be used together")}
	}
	if (*resetConfigFlag || setConfigSet) && flags.NArg() != 0 {
		return options{}, &usageError{errors.New("config commands do not accept a file or directory")}
	}
	if *resetConfigFlag {
		return options{resetConfig: true, configPath: *configPathFlag, configPathSet: configPathSet}, nil
	}
	if setConfigSet {
		return options{setConfig: *setConfigFlag, setConfigSet: true, configPath: *configPathFlag, configPathSet: configPathSet}, nil
	}
	if *updateFlag {
		return options{update: true, configPath: *configPathFlag, configPathSet: configPathSet}, nil
	}
	if *webFlag {
		if flags.NArg() != 0 {
			return options{}, &usageError{errors.New(i18n.T(i18n.CLIErrWebArgs))}
		}
		return options{web: true, configPath: *configPathFlag, configPathSet: configPathSet}, nil
	}
	bookView, err := selectedBookView(settings.View, *bookViewFlag, *rightBookViewFlag, *circleBookViewFlag, *rightCircleBookViewFlag)
	if err != nil {
		return options{}, &usageError{err}
	}
	if *versionFlag {
		return options{version: true, configPath: *configPathFlag, configPathSet: configPathSet}, nil
	}
	if *clearJournalFlag && flags.NArg() != 1 {
		return options{}, &usageError{errors.New(i18n.T(i18n.CLIErrClearJournalRequiresInput))}
	}
	switch flags.NArg() {
	case 0:
		return options{graphics: *graphicsFlag, open: *openFlag, openSet: openSet, bookView: bookView, configPath: *configPathFlag, configPathSet: configPathSet}, nil
	case 1:
		return options{graphics: *graphicsFlag, path: flags.Arg(0), clearJournal: *clearJournalFlag, bookView: bookView, configPath: *configPathFlag, configPathSet: configPathSet}, nil
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
