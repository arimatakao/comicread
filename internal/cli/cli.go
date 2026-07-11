package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/arimatakao/comicfile"
	"github.com/arimatakao/comicread/internal/backend"
	"github.com/arimatakao/comicread/internal/i18n"
	"github.com/arimatakao/comicread/internal/tui/filepicker"
	"github.com/arimatakao/comicread/internal/tui/loading"
	"github.com/arimatakao/comicread/internal/tui/reader"
)

// ErrUsage marks errors caused by invalid command-line arguments, as opposed
// to runtime failures (missing files, decode errors, etc). Callers can check
// for it with errors.Is.
var ErrUsage = errors.New("usage error")

type usageError struct{ err error }

func (e *usageError) Error() string        { return e.err.Error() }
func (e *usageError) Unwrap() error        { return e.err }
func (e *usageError) Is(target error) bool { return target == ErrUsage }

func Run(args []string) error {
	graphics, path, err := parseArgs(args)
	if err != nil {
		return err
	}

	if path == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf(i18n.T(i18n.CLIErrGetWorkingDir), err)
		}
		path, err = filepicker.Pick(cwd)
		if err != nil {
			return fmt.Errorf(i18n.T(i18n.CLIErrPickFile), err)
		}
		if path == "" {
			return nil
		}
	}

	if err := validateInput(path); err != nil {
		return err
	}

	chapter, err := loading.Open(path, func() (comicfile.ContainerReader, error) {
		return openChapter(path)
	})
	if err != nil {
		return err
	}

	renderer, err := backend.NewRenderer(graphics)
	if err != nil {
		return err
	}

	model := reader.New(filepath.Base(path), chapter, renderer)
	program := tea.NewProgram(model)
	if _, err := program.Run(); err != nil {
		return fmt.Errorf(i18n.T(i18n.CLIErrRunTUI), err)
	}

	return nil
}

func parseArgs(args []string) (graphics, path string, err error) {
	flags := flag.NewFlagSet("comicread", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	graphicsFlag := flags.String("graphics", "auto", i18n.T(i18n.CLIFlagGraphicsUsage))
	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			fmt.Println(i18n.T(i18n.CLIUsageFull))
			return "", "", flag.ErrHelp
		}
		return "", "", &usageError{fmt.Errorf(i18n.T(i18n.CLIErrParseArgs), err)}
	}
	switch flags.NArg() {
	case 0:
		return *graphicsFlag, "", nil
	case 1:
		return *graphicsFlag, flags.Arg(0), nil
	default:
		return "", "", &usageError{errors.New(i18n.T(i18n.CLIUsage))}
	}
}

func openChapter(path string) (comicfile.ContainerReader, error) {
	chapter, err := comicfile.OpenContainer(path)
	if err != nil {
		return nil, fmt.Errorf(i18n.T(i18n.CLIErrOpenChapter), err)
	}
	if chapter.TotalPages() == 0 {
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
