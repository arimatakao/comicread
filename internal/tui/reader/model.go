package reader

import (
	tea "charm.land/bubbletea/v2"
	"github.com/arimatakao/comicfile"
	"github.com/arimatakao/comicread/internal/backend"
	"github.com/arimatakao/comicread/internal/i18n"
	"github.com/arimatakao/comicread/internal/journal"
)

type Model struct {
	title              string
	chapter            comicfile.ContainerReader
	backend            backend.Renderer
	bookView           ViewMode
	page               int
	width              int
	height             int
	area               backend.Area
	pageAreas          [2]backend.Area
	displayedArea      backend.Area
	requestID          uint64
	layoutID           uint64
	layoutPending      bool
	zoom               int
	scroll             float64
	rendering          bool
	showingHelp        bool
	showingBookmarks   bool
	bookmarkPrefix     bool
	bookmarkListPrefix bool
	bookmarkIndex      int
	bookmarks          []int
	journal            *journal.Journal
	status             string
	renderError        error
	cache              *readerCache
}

func New(title string, chapter comicfile.ContainerReader, renderer backend.Renderer) Model {
	return NewWithBookView(title, chapter, renderer, SinglePageView)
}

// NewWithBookView creates a reader using the requested page layout.
func NewWithBookView(title string, chapter comicfile.ContainerReader, renderer backend.Renderer, bookView ViewMode) Model {
	return Model{
		title:    title,
		chapter:  chapter,
		backend:  renderer,
		bookView: bookView,
		zoom:     100,
		status:   i18n.T(i18n.ReaderStatusWaitingTerminalSize),
		cache:    newReaderCache(),
	}
}

// NewWithBookViewAndJournal creates a reader that restores and saves local
// reading progress and bookmarks.
func NewWithBookViewAndJournal(title string, chapter comicfile.ContainerReader, renderer backend.Renderer, bookView ViewMode, progress *journal.Journal) Model {
	m := NewWithBookView(title, chapter, renderer, bookView)
	m.journal = progress
	if progress == nil {
		return m
	}
	m.bookmarks = progress.Bookmarks()
	m.setPageNumber(progress.LastOpenedPage())
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}
