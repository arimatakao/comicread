// Package filepicker implements a small interactive Bubble Tea program for
// choosing a comic chapter (a CBZ, PDF, EPUB file, or an image directory).
package filepicker

import (
	"errors"
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
	return PickWithFavorites(root, nil)
}

// PickWithFavorites runs an interactive file picker rooted at root. Favorites
// are directory paths shown when the user presses "f".
func PickWithFavorites(root string, favorites []string) (string, error) {
	return pick(root, favorites, nil)
}

// PickWithFavoriteSaver runs an interactive file picker that persists
// favorites through saveFavorites whenever the user presses "b".
func PickWithFavoriteSaver(root string, favorites []string, saveFavorites func([]string) error) (string, error) {
	return pick(root, favorites, saveFavorites)
}

func pick(root string, favorites []string, saveFavorites func([]string) error) (string, error) {
	abs, err := filepath.Abs(root)
	if err != nil {
		return "", fmt.Errorf(i18n.T(i18n.FilepickerErrResolveDir), root, err)
	}

	model, err := newModelWithFavoriteSaver(abs, favorites, saveFavorites)
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

const (
	yellow = "\x1b[33m"
	reset  = "\x1b[0m"
)

type pickerModel struct {
	dir      string
	entries  []entry
	cursor   int
	offset   int
	selected string
	err      error
	width    int
	height   int

	goToInput        bool
	goToPath         string
	goToErr          string
	favoriteInput    bool
	favoritePath     string
	favoriteInputErr string
	commandInput     bool
	command          string
	commandErr       string

	favorites     []string
	favoriteIndex int
	favoriteList  bool
	favoriteErr   string
	saveFavorites func([]string) error
}

func newModel(dir string) (pickerModel, error) {
	return newModelWithFavorites(dir, nil)
}

func newModelWithFavorites(dir string, favorites []string) (pickerModel, error) {
	return newModelWithFavoriteSaver(dir, favorites, nil)
}

func newModelWithFavoriteSaver(dir string, favorites []string, saveFavorites func([]string) error) (pickerModel, error) {
	m := pickerModel{dir: dir, favorites: validFavorites(favorites), saveFavorites: saveFavorites}
	if err := m.readDir(); err != nil {
		return pickerModel{}, err
	}
	return m, nil
}

func validFavorites(favorites []string) []string {
	seen := make(map[string]struct{}, len(favorites))
	valid := make([]string, 0, len(favorites))
	for _, favorite := range favorites {
		if favorite == "" {
			continue
		}
		path, err := filepath.Abs(favorite)
		if err != nil {
			continue
		}
		info, err := os.Stat(path)
		if err != nil || !info.IsDir() {
			continue
		}
		path = filepath.Clean(path)
		if _, ok := seen[path]; ok {
			continue
		}
		seen[path] = struct{}{}
		valid = append(valid, path)
	}
	return valid
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
		if m.commandInput {
			return m.updateCommandInput(msg)
		}
		if m.goToInput {
			return m.updateGoToInput(msg)
		}
		if m.favoriteInput {
			return m.updateFavoriteInput(msg)
		}
		if m.favoriteList {
			return m.updateFavorites(msg)
		}

		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit

		case "o":
			m.goToInput = true
			m.goToPath = ""
			m.goToErr = ""
			return m, nil

		case ":":
			m.commandInput = true
			m.command = ""
			m.commandErr = ""
			return m, nil

		case "f":
			m.toggleFavorite()
			return m, nil

		case "F":
			m.favoriteInput = true
			m.favoritePath = ""
			m.favoriteInputErr = ""
			return m, nil

		case "b":
			m.favoriteList = true
			m.favoriteIndex = 0
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

func (m pickerModel) updateFavoriteInput(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		m.favoriteInput = false
		m.favoritePath = ""
		m.favoriteInputErr = ""
	case "enter":
		path := strings.TrimSpace(m.favoritePath)
		if path == "" {
			m.favoriteInputErr = i18n.T(i18n.FilepickerErrEmptyPath)
			return m, nil
		}
		abs, err := directoryPath(path)
		if err != nil {
			m.favoriteInputErr = err.Error()
			return m, nil
		}
		if err := m.addFavorite(abs); err != nil {
			m.favoriteInputErr = err.Error()
			return m, nil
		}
		m.favoriteInput = false
		m.favoritePath = ""
		m.favoriteInputErr = ""
	case "backspace":
		if len(m.favoritePath) > 0 {
			_, size := utf8.DecodeLastRuneInString(m.favoritePath)
			m.favoritePath = m.favoritePath[:len(m.favoritePath)-size]
		}
	default:
		if msg.Text != "" {
			m.favoritePath += msg.Text
		}
	}
	return m, nil
}

func (m pickerModel) updateCommandInput(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		m.commandInput = false
		m.command = ""
		m.commandErr = ""
	case "enter":
		command, path, ok := strings.Cut(strings.TrimSpace(m.command), " ")
		path = strings.TrimSpace(path)
		if !ok || path == "" {
			m.commandErr = "command requires a directory path"
			return m, nil
		}
		switch command {
		case "g":
			abs, err := directoryPath(path)
			if err != nil {
				m.commandErr = err.Error()
				return m, nil
			}
			m.commandInput = false
			m.command = ""
			m.commandErr = ""
			if cmd := m.enterDir(abs); cmd != nil {
				return m, cmd
			}
		case "f":
			abs, err := directoryPath(path)
			if err != nil {
				m.commandErr = err.Error()
				return m, nil
			}
			if err := m.addFavorite(abs); err != nil {
				m.commandErr = err.Error()
				return m, nil
			}
			m.commandInput = false
			m.command = ""
			m.commandErr = ""
		default:
			m.commandErr = fmt.Sprintf("unknown command %q", command)
		}
	case "backspace":
		if len(m.command) > 0 {
			_, size := utf8.DecodeLastRuneInString(m.command)
			m.command = m.command[:len(m.command)-size]
		}
	default:
		if msg.Text != "" {
			m.command += msg.Text
		}
	}
	return m, nil
}

func directoryPath(path string) (string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf(i18n.T(i18n.FilepickerErrResolveDir), path, err)
	}
	info, err := os.Stat(abs)
	if err != nil || !info.IsDir() {
		return "", errors.New(i18n.T(i18n.FilepickerErrNotDir, abs))
	}
	return filepath.Clean(abs), nil
}

func (m *pickerModel) toggleFavorite() {
	if m.saveFavorites == nil {
		return
	}
	path := filepath.Clean(m.dir)
	index := -1
	for i, favorite := range m.favorites {
		if favorite == path {
			index = i
			break
		}
	}
	updated := append([]string(nil), m.favorites...)
	if index >= 0 {
		updated = append(updated[:index], updated[index+1:]...)
	} else {
		updated = append(updated, path)
	}
	if err := m.saveFavorites(updated); err != nil {
		m.favoriteErr = err.Error()
		return
	}
	m.favorites = updated
	m.favoriteErr = ""
}

func (m *pickerModel) addFavorite(path string) error {
	if m.saveFavorites == nil {
		return errors.New("favorites cannot be saved")
	}
	path = filepath.Clean(path)
	for _, favorite := range m.favorites {
		if favorite == path {
			return nil
		}
	}
	updated := append(append([]string(nil), m.favorites...), path)
	if err := m.saveFavorites(updated); err != nil {
		return fmt.Errorf("save favorites: %w", err)
	}
	m.favorites = updated
	m.favoriteErr = ""
	return nil
}

func (m pickerModel) isFavorite(path string) bool {
	path = filepath.Clean(path)
	for _, favorite := range m.favorites {
		if favorite == path {
			return true
		}
	}
	return false
}

func (m pickerModel) updateFavorites(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit
	case "esc", "f":
		m.favoriteList = false
	case "up", "k":
		if m.favoriteIndex > 0 {
			m.favoriteIndex--
		}
	case "down", "j":
		if m.favoriteIndex < len(m.favorites)-1 {
			m.favoriteIndex++
		}
	case "d":
		if len(m.favorites) == 0 || m.saveFavorites == nil {
			return m, nil
		}
		updated := append([]string(nil), m.favorites...)
		updated = append(updated[:m.favoriteIndex], updated[m.favoriteIndex+1:]...)
		if err := m.saveFavorites(updated); err != nil {
			m.favoriteErr = err.Error()
			return m, nil
		}
		m.favorites = updated
		m.favoriteIndex = min(m.favoriteIndex, len(m.favorites)-1)
		m.favoriteErr = ""
	case "enter", "right", "l":
		if len(m.favorites) == 0 {
			return m, nil
		}
		m.favoriteList = false
		if cmd := m.enterDir(m.favorites[m.favoriteIndex]); cmd != nil {
			return m, cmd
		}
	}
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
	if m.favoriteList {
		b.WriteString(i18n.T(i18n.FilepickerFavorites))
		if m.favoriteErr != "" {
			fmt.Fprintf(&b, i18n.T(i18n.FilepickerFavoriteErr), m.favoriteErr)
		}
		if len(m.favorites) == 0 {
			b.WriteString(i18n.T(i18n.FilepickerNoFavorites))
		}
		for i, favorite := range m.favorites {
			cursor := "  "
			if i == m.favoriteIndex {
				cursor = "> "
			}
			fmt.Fprintf(&b, "%s%s/\n", cursor, favorite)
		}
		b.WriteString(i18n.T(i18n.FilepickerFavoritesHelp))
		view := tea.NewView(b.String())
		view.AltScreen = true
		view.WindowTitle = i18n.T(i18n.FilepickerWindowTitle)
		return view
	}
	fmt.Fprintf(&b, i18n.T(i18n.FilepickerHeader), m.dir)
	if m.favoriteErr != "" {
		fmt.Fprintf(&b, i18n.T(i18n.FilepickerFavoriteErr), m.favoriteErr)
	}

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
			if m.isFavorite(m.entryPath(e)) {
				name = yellow + name + reset
			}
		}
		fmt.Fprintf(&b, "%s%s\n", cursor, name)
	}

	if m.goToInput {
		fmt.Fprintf(&b, i18n.T(i18n.FilepickerGoToPrompt), m.goToPath)
		if m.goToErr != "" {
			fmt.Fprintf(&b, i18n.T(i18n.FilepickerGoToErr), m.goToErr)
		}
	}
	if m.favoriteInput {
		fmt.Fprintf(&b, i18n.T(i18n.FilepickerFavoritePrompt), m.favoritePath)
		if m.favoriteInputErr != "" {
			fmt.Fprintf(&b, i18n.T(i18n.FilepickerGoToErr), m.favoriteInputErr)
		}
	}

	b.WriteString(i18n.T(i18n.FilepickerHelp))

	view := tea.NewView(b.String())
	view.AltScreen = true
	view.WindowTitle = i18n.T(i18n.FilepickerWindowTitle)
	return view
}
