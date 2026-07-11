package cli

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/arimatakao/comicfile"
	"github.com/arimatakao/comicread/internal/backend"
	"github.com/arimatakao/comicread/internal/tui"
)

func Run(args []string) error {
	graphics, path, err := parseArgs(args)
	if err != nil {
		return err
	}

	if path == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("get working directory: %w", err)
		}
		path, err = tui.PickFile(cwd)
		if err != nil {
			return fmt.Errorf("pick file: %w", err)
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

	model := tui.New(filepath.Base(path), chapter, renderer)
	program := tea.NewProgram(model)
	if _, err := program.Run(); err != nil {
		return fmt.Errorf("run TUI: %w", err)
	}

	return nil
}

func parseArgs(args []string) (graphics, path string, err error) {
	flags := flag.NewFlagSet("comicread", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	graphicsFlag := flags.String("graphics", "auto", "graphics protocol: auto, kitty, sixel, or iterm2")
	if err := flags.Parse(args); err != nil {
		return "", "", fmt.Errorf("parse arguments: %w", err)
	}
	switch flags.NArg() {
	case 0:
		return *graphicsFlag, "", nil
	case 1:
		return *graphicsFlag, flags.Arg(0), nil
	default:
		return "", "", fmt.Errorf("usage: comicread [--graphics auto|kitty|sixel|iterm2] [file.cbz|file.pdf|file.epub|image-directory]")
	}
}

func openChapter(path string) (comicfile.ContainerReader, error) {
	chapter, err := comicfile.OpenContainer(path)
	if err != nil {
		return nil, fmt.Errorf("open chapter: %w", err)
	}
	if chapter.TotalPages() == 0 {
		return nil, fmt.Errorf("chapter contains no readable image pages")
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
		return fmt.Errorf("inspect input %q: %w", path, err)
	}
	if info.IsDir() || isSupportedFile(path) {
		return nil
	}
	return fmt.Errorf("unsupported file %q: supported formats are CBZ, PDF, EPUB, or an image directory", path)
}
