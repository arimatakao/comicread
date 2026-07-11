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
	"github.com/arimatakao/comicread/internal/tui/reader"
)

func Run(args []string) error {
	graphics, path, err := parseArgs(args)
	if err != nil {
		return err
	}

	if path == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf(i18n.T("cli.err.get_working_dir"), err)
		}
		path, err = filepicker.Pick(cwd)
		if err != nil {
			return fmt.Errorf(i18n.T("cli.err.pick_file"), err)
		}
		if path == "" {
			return nil
		}
	}

	if err := validateInput(path); err != nil {
		return err
	}

	chapter, err := openChapter(path)
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
		return fmt.Errorf(i18n.T("cli.err.run_tui"), err)
	}

	return nil
}

func parseArgs(args []string) (graphics, path string, err error) {
	flags := flag.NewFlagSet("comicread", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	graphicsFlag := flags.String("graphics", "auto", i18n.T("cli.flag.graphics_usage"))
	if err := flags.Parse(args); err != nil {
		return "", "", fmt.Errorf(i18n.T("cli.err.parse_args"), err)
	}
	switch flags.NArg() {
	case 0:
		return *graphicsFlag, "", nil
	case 1:
		return *graphicsFlag, flags.Arg(0), nil
	default:
		return "", "", errors.New(i18n.T("cli.usage"))
	}
}

func openChapter(path string) (comicfile.ContainerReader, error) {
	chapter, err := comicfile.OpenContainer(path)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("cli.err.open_chapter"), err)
	}
	if chapter.TotalPages() == 0 {
		return nil, errors.New(i18n.T("cli.err.no_pages"))
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
		return fmt.Errorf(i18n.T("cli.err.inspect_input"), path, err)
	}
	if info.IsDir() || isSupportedFile(path) {
		return nil
	}
	return fmt.Errorf(i18n.T("cli.err.unsupported_file"), path)
}
