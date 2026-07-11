package cli

import (
	"fmt"
	"path/filepath"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/arimatakao/comicfile"
	"github.com/arimatakao/comicread/internal/backend"
	"github.com/arimatakao/comicread/internal/tui"
)

func Run(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: comicread <file.cbz>")
	}

	path := args[0]
	if !strings.EqualFold(filepath.Ext(path), ".cbz") {
		return fmt.Errorf("unsupported file %q: only CBZ is available for now", path)
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
		return nil, fmt.Errorf("open CBZ: %w", err)
	}
	if chapter.TotalPages() == 0 {
		return nil, fmt.Errorf("CBZ contains no readable image pages")
	}
	return chapter, nil
}
