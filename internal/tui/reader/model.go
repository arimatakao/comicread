package reader

import (
	tea "charm.land/bubbletea/v2"
	"github.com/arimatakao/comicfile"
	"github.com/arimatakao/comicread/internal/backend"
	"github.com/arimatakao/comicread/internal/i18n"
)

type Model struct {
	title         string
	chapter       comicfile.ContainerReader
	backend       backend.Renderer
	page          int
	width         int
	height        int
	area          backend.Area
	displayedArea backend.Area
	requestID     uint64
	zoom          int
	scroll        float64
	rendering     bool
	status        string
	renderError   error
}

func New(title string, chapter comicfile.ContainerReader, renderer backend.Renderer) Model {
	return Model{
		title:   title,
		chapter: chapter,
		backend: renderer,
		zoom:    100,
		status:  i18n.T(i18n.ReaderStatusWaitingTerminalSize),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
