// Package filepicker implements a small interactive Bubble Tea program for
// choosing a comic chapter (a CBZ, PDF, EPUB file, or an image directory).
package filepicker

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/arimatakao/comicread/internal/i18n"
)

// Pick runs an interactive file picker rooted at root and returns the chosen
// chapter path. It returns an empty path if the user cancels.
func Pick(root string) (string, error) {
	abs, err := filepath.Abs(root)
	if err != nil {
		return "", fmt.Errorf(i18n.T(i18n.FilepickerErrResolveDir), root, err)
	}

	model, err := newModel(abs)
	if err != nil {
		return "", err
	}

	program := tea.NewProgram(model)
	finalModel, err := program.Run()
	if err != nil {
		return "", fmt.Errorf(i18n.T(i18n.FilepickerErrRunPicker), err)
	}

	result := finalModel.(pickerModel)
	return result.selected, result.err
}

type entry struct {
	name  string
	isDir bool
}

type pickerModel struct {
	dir      string
	entries  []entry
	cursor   int
	offset   int
	selected string
	err      error
	width    int
	height   int
}

func newModel(dir string) (pickerModel, error) {
	m := pickerModel{dir: dir}
	if err := m.readDir(); err != nil {
		return pickerModel{}, err
	}
	return m, nil
}

func (m *pickerModel) readDir() error {
	dirEntries, err := os.ReadDir(m.dir)
	if err != nil {
		return fmt.Errorf(i18n.T(i18n.FilepickerErrReadDir), m.dir, err)
	}

	items := make([]entry, 0, len(dirEntries)+1)
	if parent := filepath.Dir(m.dir); parent != m.dir {
		items = append(items, entry{name: "..", isDir: true})
	}
	for _, e := range dirEntries {
		if e.IsDir() {
			items = append(items, entry{name: e.Name(), isDir: true})
			continue
		}
		if isSupportedChapterFile(e.Name()) {
			items = append(items, entry{name: e.Name()})
		}
	}

	m.entries = items
	if m.cursor >= len(items) {
		m.cursor = 0
	}
	m.keepCursorVisible()
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

func (m pickerModel) Init() tea.Cmd {
	return nil
}

func (m pickerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.keepCursorVisible()

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
				e := m.entries[m.cursor]
				if e.name == ".." {
					m.selected = filepath.Dir(m.dir)
				} else {
					m.selected = filepath.Join(m.dir, e.name)
				}
			} else {
				m.selected = m.dir
			}
			return m, tea.Quit

		case "left", "h":
			m.dir = filepath.Dir(m.dir)
			if err := m.readDir(); err != nil {
				m.err = err
				return m, tea.Quit
			}

		case "right", "l":
			if len(m.entries) == 0 || !m.entries[m.cursor].isDir {
				return m, nil
			}
			e := m.entries[m.cursor]
			if e.name == ".." {
				m.dir = filepath.Dir(m.dir)
			} else {
				m.dir = filepath.Join(m.dir, e.name)
			}
			if err := m.readDir(); err != nil {
				m.err = err
				return m, tea.Quit
			}

		case "enter":
			if len(m.entries) == 0 {
				return m, nil
			}
			e := m.entries[m.cursor]
			if e.isDir {
				return m, nil
			}
			m.selected = filepath.Join(m.dir, e.name)
			return m, tea.Quit
		}
	}

	m.keepCursorVisible()
	return m, nil
}

func (m *pickerModel) keepCursorVisible() {
	rows := m.entryRows()
	if m.cursor < m.offset {
		m.offset = m.cursor
	}
	if m.cursor >= m.offset+rows {
		m.offset = m.cursor - rows + 1
	}

	maxOffset := max(0, len(m.entries)-rows)
	m.offset = min(max(0, m.offset), maxOffset)
}

func (m pickerModel) entryRows() int {
	const reservedRows = 5 // title, path, spacing, and help text
	if m.height <= 0 {
		return max(1, len(m.entries))
	}
	if m.height <= reservedRows {
		return 1
	}
	return m.height - reservedRows
}

func (m pickerModel) View() tea.View {
	var b strings.Builder
	fmt.Fprintf(&b, i18n.T(i18n.FilepickerHeader), m.dir)

	if len(m.entries) == 0 {
		b.WriteString(i18n.T(i18n.FilepickerNoEntries))
	}
	end := min(len(m.entries), m.offset+m.entryRows())
	for i := m.offset; i < end; i++ {
		e := m.entries[i]
		cursor := "  "
		if i == m.cursor {
			cursor = "> "
		}
		name := e.name
		if e.isDir {
			name += "/"
		}
		fmt.Fprintf(&b, "%s%s\n", cursor, name)
	}

	b.WriteString(i18n.T(i18n.FilepickerHelp))

	view := tea.NewView(b.String())
	view.AltScreen = true
	view.WindowTitle = i18n.T(i18n.FilepickerWindowTitle)
	return view
}
