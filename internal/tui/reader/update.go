package reader

import (
	tea "charm.land/bubbletea/v2"
	"github.com/arimatakao/comicread/internal/backend"
	"github.com/arimatakao/comicread/internal/i18n"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		return m.handleResize(msg)

	case renderAfterLayoutMsg:
		if msg.layoutID != m.layoutID || !m.layoutPending {
			return m, nil
		}
		return m, m.renderPage()

	case itermImageReadyMsg:
		if msg.requestID != m.requestID || msg.page != m.page || m.showingHelp || m.showingBookmarks {
			return m, nil
		}
		return m, tea.Raw(msg.output)

	case tea.KeyPressMsg:
		return m.handleKey(msg)

	case pageRenderedMsg:
		return m.handlePageRendered(msg)
	}

	return m, nil
}

func (m Model) handleResize(msg tea.WindowSizeMsg) (tea.Model, tea.Cmd) {
	m.width = msg.Width
	m.height = msg.Height
	m.updateCellSize()
	m.updateLayout()
	if m.showingHelp {
		return m, nil
	}
	if m.area.Cols < 1 || m.area.Rows < 1 {
		m.layoutPending = false
		m.status = i18n.T(i18n.ReaderStatusTerminalTooSmall)
		return m, m.clearAndRepaint(m.displayedArea)
	}
	m.layoutID++
	m.layoutPending = true
	return m, renderAfterLayout(m.layoutID)
}

func (m Model) handleKey(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	if m.showingBookmarks {
		return m.handleBookmarksKey(msg.String())
	}

	act := keyAction(msg)
	switch {
	case act == actionQuit:
		m.saveCurrentPage()
		return m, tea.Sequence(tea.Raw(m.backend.Clear(m.displayedArea)), tea.Quit)

	case act == actionHelp:
		return m.toggleHelp()

	case m.showingHelp:
		return m, nil

	case m.bookmarkListPrefix:
		m.bookmarkListPrefix = false
		if act == actionBookmarkPrefix { // "c v" opens the bookmark list
			m.openBookmarks()
			m.invalidateRender()
			return m, m.clearAndRepaint(m.displayedArea)
		}
		return m, nil

	case m.bookmarkPrefix:
		m.bookmarkPrefix = false
		switch msg.String() {
		case "left":
			if m.moveBookmark(false) {
				return m, m.rerender()
			}
		case "right":
			if m.moveBookmark(true) {
				return m, m.rerender()
			}
		}
		return m, nil
	}

	return m.handleReaderAction(act)
}

func (m Model) handleReaderAction(act action) (tea.Model, tea.Cmd) {
	switch act {
	case actionBookmarkListPrefix:
		m.bookmarkListPrefix = true

	case actionBookmarkPrefix:
		m.bookmarkPrefix = true

	case actionBookmark:
		m.toggleBookmark()

	case actionZoomIn:
		if m.zoomIn() {
			return m, m.rerender()
		}

	case actionZoomOut:
		if m.zoomOut() {
			return m, m.rerender()
		}

	case actionScrollDown:
		if m.scrollDown() {
			return m, m.renderPage()
		}

	case actionScrollUp:
		if m.scrollUp() {
			return m, m.renderPage()
		}

	case actionNext:
		if m.nextPage() {
			m.saveCurrentPage()
			return m, m.rerender()
		}
		m.status = i18n.T(i18n.ReaderStatusLastPage)

	case actionPrevious:
		if m.previousPage() {
			m.saveCurrentPage()
			return m, m.rerender()
		}
		m.status = i18n.T(i18n.ReaderStatusFirstPage)
	}

	return m, nil
}

func (m Model) handleBookmarksKey(key string) (tea.Model, tea.Cmd) {
	switch key {
	case "q", "esc":
		m.closeBookmarks()
		return m, m.renderPage()
	case "up":
		m.bookmarkIndex = max(0, m.bookmarkIndex-1)
	case "down":
		if len(m.bookmarks) > 0 {
			m.bookmarkIndex = min(len(m.bookmarks)-1, m.bookmarkIndex+1)
		}
	case "enter":
		if m.selectBookmark() {
			m.closeBookmarks()
			return m, m.rerender()
		}
	}
	return m, nil
}

func (m Model) toggleHelp() (tea.Model, tea.Cmd) {
	m.showingHelp = !m.showingHelp
	m.status = ""
	if m.showingHelp {
		m.invalidateRender()
		return m, m.clearAndRepaint(m.displayedArea)
	}
	return m, m.renderPage()
}

func (m Model) handlePageRendered(msg pageRenderedMsg) (tea.Model, tea.Cmd) {
	if msg.requestID != m.requestID || msg.page != m.page {
		return m, nil
	}

	m.rendering = false
	if msg.err != nil {
		m.renderError = msg.err
		m.status = i18n.T(i18n.ReaderStatusRenderError, msg.err)
		return m, nil
	}

	m.status = ""
	oldArea := m.displayedArea
	m.displayedArea = msg.area
	return m, tea.Batch(m.clearAndRender(oldArea, msg.output), m.preRenderAround())
}

// invalidateRender abandons any in-flight render so its output is dropped.
func (m *Model) invalidateRender() {
	m.requestID++
	if m.cache != nil {
		m.cache.latestRequest.Store(m.requestID)
	}
	m.layoutPending = false
	m.rendering = false
}

// rerender recomputes the page layout and renders the current position.
func (m *Model) rerender() tea.Cmd {
	m.updateLayout()
	return m.renderPage()
}

func (m Model) clearAndRepaint(area backend.Area) tea.Cmd {
	clear := tea.Raw(m.backend.Clear(area))
	if m.backend.Name() != "iterm2" {
		return clear
	}
	// iTerm2 removes inline images only by clearing the whole display. Tell
	// Bubble Tea about that reset so it redraws its text layer itself; writing
	// the text through tea.Raw would leave its cursor state out of sync.
	return tea.Sequence(clear, tea.ClearScreen)
}

func (m Model) clearAndRender(area backend.Area, output string) tea.Cmd {
	clear := tea.Raw(m.backend.Clear(area))
	image := tea.Raw(output)
	if m.backend.Name() != "iterm2" || area.Cols < 1 || area.Rows < 1 {
		return tea.Sequence(clear, image)
	}
	return tea.Sequence(
		clear,
		tea.ClearScreen,
		renderItermImageAfterClear(m.requestID, m.page, output),
	)
}

func (m *Model) nextPage() bool {
	if !m.canNextPage() {
		return false
	}
	m.page++
	m.scroll = 0
	return true
}

func (m *Model) previousPage() bool {
	if m.page == 0 {
		return false
	}
	m.page--
	m.scroll = 0
	return true
}

const (
	zoomStep   = 25
	maxZoom    = 400
	scrollStep = 0.1
)

func (m *Model) zoomIn() bool {
	if m.zoom >= maxZoom {
		m.status = i18n.T(i18n.ReaderStatusMaximumZoom)
		return false
	}
	m.zoom += zoomStep
	m.status = ""
	return true
}

func (m *Model) zoomOut() bool {
	if m.zoom <= 100 {
		m.status = i18n.T(i18n.ReaderStatusMinimumZoom)
		return false
	}
	m.zoom -= zoomStep
	if m.zoom == 100 {
		m.scroll = 0
	}
	m.status = ""
	return true
}

func (m *Model) scrollDown() bool {
	if m.zoom <= 100 || m.scroll >= 1 {
		return false
	}
	m.scroll = min(1, m.scroll+scrollStep)
	m.status = ""
	return true
}

func (m *Model) scrollUp() bool {
	if m.zoom <= 100 || m.scroll <= 0 {
		return false
	}
	m.scroll = max(0, m.scroll-scrollStep)
	m.status = ""
	return true
}

func (m Model) currentPageAspect() float64 {
	slots := m.pageSlots()
	if slots[0] == m.layoutPages[0] && m.layoutImages[0] != nil {
		return imageAspect(m.layoutImages[0])
	}
	img, err := m.chapter.Page(slots[0])
	if err != nil {
		return 1
	}
	return imageAspect(img)
}
