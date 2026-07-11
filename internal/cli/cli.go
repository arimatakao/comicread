package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/arimatakao/comicfile"
	"github.com/arimatakao/comicread/internal/backend"
	"github.com/arimatakao/comicread/internal/tui"
)

func Run(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: comicread <file.cbz|file.pdf|file.epub|image-directory>")
	}

	path := args[0]
	if err := validateInput(path); err != nil {
		return err
	}

	chapter, err := openChapter(path)
	if err != nil {
		return err
	}

	model := tui.New(filepath.Base(path), chapter, backend.NewKitty())
	program := tea.NewProgram(model)
	if _, err := program.Run(); err != nil {
		return fmt.Errorf("run TUI: %w", err)
	}

	return nil
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
