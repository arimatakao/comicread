// Package loading displays progress while a comic chapter is opened.
package loading

import (
	"path/filepath"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/arimatakao/comicfile"
	"github.com/arimatakao/comicread/internal/i18n"
)

const tickInterval = 100 * time.Millisecond

var frames = []string{"|", "/", "-", "\\"}

// Open displays an animated loading screen while opener opens a chapter.
func Open(path string, opener func() (comicfile.ContainerReader, error)) (comicfile.ContainerReader, error) {
	finalModel, err := tea.NewProgram(newModel(path, opener)).Run()
	if err != nil {
		return nil, err
	}

	m := finalModel.(model)
	return m.chapter, m.err
}

type openFinishedMsg struct {
	chapter comicfile.ContainerReader
	err     error
}

type tickMsg struct{}

type model struct {
	path    string
	opener  func() (comicfile.ContainerReader, error)
	chapter comicfile.ContainerReader
	err     error
	frame   int
}

func newModel(path string, opener func() (comicfile.ContainerReader, error)) model {
	return model{path: path, opener: opener}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.open(), nextTick())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		m.frame = (m.frame + 1) % len(frames)
		return m, nextTick()

	case openFinishedMsg:
		m.chapter = msg.chapter
		m.err = msg.err
		return m, tea.Quit
	}

	return m, nil
}

func (m model) View() tea.View {
	view := tea.NewView(frames[m.frame] + " " + i18n.T(i18n.LoadingViewOpening, filepath.Base(m.path)) + "\n")
	view.AltScreen = true
	view.WindowTitle = i18n.T(i18n.LoadingViewWindowTitle)
	return view
}

func (m model) open() tea.Cmd {
	return func() tea.Msg {
		chapter, err := m.opener()
		return openFinishedMsg{chapter: chapter, err: err}
	}
}

func nextTick() tea.Cmd {
	return tea.Tick(tickInterval, func(time.Time) tea.Msg { return tickMsg{} })
}
