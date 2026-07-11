package tui

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tea "charm.land/bubbletea/v2"
)

// PickFile runs an interactive file picker rooted at root and returns the
// chosen chapter path. It returns an empty path if the user cancels.
func PickFile(root string) (string, error) {
	abs, err := filepath.Abs(root)
	if err != nil {
		return "", fmt.Errorf("resolve directory %q: %w", root, err)
	}

	model, err := newFilePickerModel(abs)
	if err != nil {
		return "", err
	}

	program := tea.NewProgram(model)
	finalModel, err := program.Run()
	if err != nil {
		return "", fmt.Errorf("run file picker: %w", err)
	}

	result := finalModel.(filePickerModel)
	return result.selected, result.err
}

type pickerEntry struct {
	name  string
	isDir bool
}

type filePickerModel struct {
	dir      string
	entries  []pickerEntry
	cursor   int
	selected string
	err      error
	width    int
	height   int
}

func newFilePickerModel(dir string) (filePickerModel, error) {
	m := filePickerModel{dir: dir}
	if err := m.readDir(); err != nil {
		return filePickerModel{}, err
	}
	return m, nil
}

func (m *filePickerModel) readDir() error {
	entries, err := os.ReadDir(m.dir)
	if err != nil {
		return fmt.Errorf("read directory %q: %w", m.dir, err)
	}

	items := make([]pickerEntry, 0, len(entries)+1)
	if parent := filepath.Dir(m.dir); parent != m.dir {
		items = append(items, pickerEntry{name: "..", isDir: true})
	}
	for _, entry := range entries {
		if entry.IsDir() {
			items = append(items, pickerEntry{name: entry.Name(), isDir: true})
			continue
		}
		if isSupportedChapterFile(entry.Name()) {
			items = append(items, pickerEntry{name: entry.Name()})
		}
	}

	m.entries = items
	if m.cursor >= len(items) {
		m.cursor = 0
	}
	return nil
}

func isSupportedChapterFile(name string) bool {
	switch strings.ToLower(filepath.Ext(name)) {
	case ".cbz", ".pdf", ".epub":
		return true
	default:
		return false
	}
}

func (m filePickerModel) Init() tea.Cmd {
	return nil
}

func (m filePickerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.entries)-1 {
				m.cursor++
			}

		case "s":
			if len(m.entries) > 0 && m.entries[m.cursor].isDir {
				entry := m.entries[m.cursor]
				if entry.name == ".." {
					m.selected = filepath.Dir(m.dir)
				} else {
					m.selected = filepath.Join(m.dir, entry.name)
				}
			} else {
				m.selected = m.dir
			}
			return m, tea.Quit

		case "enter", " ":
			if len(m.entries) == 0 {
				return m, nil
			}
			entry := m.entries[m.cursor]
			if !entry.isDir {
				m.selected = filepath.Join(m.dir, entry.name)
				return m, tea.Quit
			}

			if entry.name == ".." {
				m.dir = filepath.Dir(m.dir)
			} else {
				m.dir = filepath.Join(m.dir, entry.name)
			}
			if err := m.readDir(); err != nil {
				m.err = err
				return m, tea.Quit
			}
		}
	}

	return m, nil
}

func (m filePickerModel) View() tea.View {
	var b strings.Builder
	fmt.Fprintf(&b, "comicread — select a chapter\n%s\n\n", m.dir)

	if len(m.entries) == 0 {
		b.WriteString("  (no supported entries)\n")
	}
	for i, entry := range m.entries {
		cursor := "  "
		if i == m.cursor {
			cursor = "> "
		}
		name := entry.name
		if entry.isDir {
			name += "/"
		}
		fmt.Fprintf(&b, "%s%s\n", cursor, name)
	}

	b.WriteString("\n↑/↓ move  enter open/select  s select highlighted directory  q quit\n")

	view := tea.NewView(b.String())
	view.AltScreen = true
	view.WindowTitle = "comicread — pick a file"
	return view
}
