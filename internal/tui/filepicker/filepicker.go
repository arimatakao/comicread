// Package filepicker implements a small interactive Bubble Tea program for
// choosing a comic chapter (a CBZ, PDF, EPUB file, or an image directory).
package filepicker

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

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

	goToInput bool
	goToPath  string
	goToErr   string
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
		if m.goToInput {
			return m.updateGoToInput(msg)
		}

		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit

		case "o":
			m.goToInput = true
			m.goToPath = ""
			m.goToErr = ""
			return m, nil

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.entries)-1 {
				m.cursor++
			}

		case "s":
			m.selected = m.dir
			if len(m.entries) > 0 && m.entries[m.cursor].isDir {
				m.selected = m.entryPath(m.entries[m.cursor])
			}
			return m, tea.Quit

		case "left", "h":
			if cmd := m.enterDir(filepath.Dir(m.dir)); cmd != nil {
				return m, cmd
			}

		case "right", "l":
			if len(m.entries) == 0 || !m.entries[m.cursor].isDir {
				return m, nil
			}
			if cmd := m.enterDir(m.entryPath(m.entries[m.cursor])); cmd != nil {
				return m, cmd
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

// updateGoToInput handles key presses while the user is typing a directory
// path to jump to (triggered by "o").
func (m pickerModel) updateGoToInput(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		m.goToInput = false
		m.goToPath = ""
		m.goToErr = ""

	case "enter":
		path := strings.TrimSpace(m.goToPath)
		if path == "" {
			m.goToErr = i18n.T(i18n.FilepickerErrEmptyPath)
			return m, nil
		}
		abs, err := filepath.Abs(path)
		if err != nil {
			m.goToErr = fmt.Sprintf(i18n.T(i18n.FilepickerErrResolveDir), path, err)
			return m, nil
		}
		info, err := os.Stat(abs)
		if err != nil || !info.IsDir() {
			m.goToErr = fmt.Sprintf(i18n.T(i18n.FilepickerErrNotDir), abs)
			return m, nil
		}
		m.goToInput = false
		m.goToPath = ""
		m.goToErr = ""
		if cmd := m.enterDir(abs); cmd != nil {
			return m, cmd
		}

	case "backspace":
		if len(m.goToPath) > 0 {
			_, size := utf8.DecodeLastRuneInString(m.goToPath)
			m.goToPath = m.goToPath[:len(m.goToPath)-size]
		}

	default:
		if msg.Text != "" {
			m.goToPath += msg.Text
		}
	}

	return m, nil
}

// entryPath resolves an entry to an absolute path; ".." points at the parent
// of the current directory.
func (m pickerModel) entryPath(e entry) string {
	if e.name == ".." {
		return filepath.Dir(m.dir)
	}
	return filepath.Join(m.dir, e.name)
}

// enterDir switches the picker to path. It returns a quit command when the
// directory cannot be read.
func (m *pickerModel) enterDir(path string) tea.Cmd {
	m.dir = path
	m.cursor = 0
	m.offset = 0
	if err := m.readDir(); err != nil {
		m.err = err
		return tea.Quit
	}
	return nil
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

	if m.goToInput {
		fmt.Fprintf(&b, i18n.T(i18n.FilepickerGoToPrompt), m.goToPath)
		if m.goToErr != "" {
			fmt.Fprintf(&b, i18n.T(i18n.FilepickerGoToErr), m.goToErr)
		}
	}

	b.WriteString(i18n.T(i18n.FilepickerHelp))

	view := tea.NewView(b.String())
	view.AltScreen = true
	view.WindowTitle = i18n.T(i18n.FilepickerWindowTitle)
	return view
}
