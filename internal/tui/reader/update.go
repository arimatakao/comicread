package reader

import (
	tea "charm.land/bubbletea/v2"
	"github.com/arimatakao/comicread/internal/backend"
	"github.com/arimatakao/comicread/internal/i18n"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
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
		key := msg.String()
		switch {
		case m.showingBookmarks:
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
					m.updateLayout()
					return m, m.renderPage()
				}
			}
			return m, nil

		case key == "q", key == "esc", key == "ctrl+c":
			m.saveCurrentPage()
			return m, tea.Sequence(tea.Raw(m.backend.Clear(m.displayedArea)), tea.Quit)

		case isHelpKey(key):
			m.showingHelp = !m.showingHelp
			m.status = ""
			if m.showingHelp {
				m.requestID++
				if m.cache != nil {
					m.cache.latestRequest.Store(m.requestID)
				}
				m.layoutPending = false
				m.rendering = false
				return m, m.clearAndRepaint(m.displayedArea)
			}
			return m, m.renderPage()

		case m.showingHelp:
			return m, nil

		case m.bookmarkListPrefix:
			m.bookmarkListPrefix = false
			if key == "v" {
				m.openBookmarks()
				m.requestID++
				if m.cache != nil {
					m.cache.latestRequest.Store(m.requestID)
				}
				m.layoutPending = false
				m.rendering = false
				return m, m.clearAndRepaint(m.displayedArea)
			}

		case m.bookmarkPrefix:
			m.bookmarkPrefix = false
			switch key {
			case "left":
				if m.moveBookmark(false) {
					m.updateLayout()
					return m, m.renderPage()
				}
			case "right":
				if m.moveBookmark(true) {
					m.updateLayout()
					return m, m.renderPage()
				}
			}

		case isBookmarkListPrefixKey(key):
			m.bookmarkListPrefix = true

		case isBookmarkPrefixKey(key):
			m.bookmarkPrefix = true

		case isBookmarkKey(key):
			m.toggleBookmark()

		case isZoomInKey(msg):
			if m.zoomIn() {
				m.updateLayout()
				return m, m.renderPage()
			}

		case isZoomOutKey(msg):
			if m.zoomOut() {
				m.updateLayout()
				return m, m.renderPage()
			}

		case isScrollDownKey(key):
			if m.scrollDown() {
				return m, m.renderPage()
			}

		case isScrollUpKey(key):
			if m.scrollUp() {
				return m, m.renderPage()
			}

		case isNextKey(key):
			if m.nextPage() {
				m.saveCurrentPage()
				m.updateLayout()
				return m, m.renderPage()
			}
			m.status = i18n.T(i18n.ReaderStatusLastPage)

		case isPreviousKey(key):
			if m.previousPage() {
				m.saveCurrentPage()
				m.updateLayout()
				return m, m.renderPage()
			}
			m.status = i18n.T(i18n.ReaderStatusFirstPage)
		}

	case pageRenderedMsg:
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

	return m, nil
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

func imageArea(width, height int, imageAspect float64) backend.Area {
	return imageAreaWithCellAspect(width, height, imageAspect, 0.5)
}

func (m Model) imageArea(imageAspect float64) backend.Area {
	return imageAreaWithCellAspect(m.width, m.height, imageAspect, m.cellAspect())
}

func imageAreaWithCellAspect(width, height int, imageAspect, cellAspect float64) backend.Area {
	availableRows := height - 1 // one header row
	if width < 1 || availableRows < 1 {
		return backend.Area{}
	}
	if imageAspect <= 0 {
		imageAspect = 1
	}
	if cellAspect <= 0 {
		cellAspect = 0.5
	}

	cols := width
	rows := max(1, int(float64(cols)*cellAspect/imageAspect))
	if rows > availableRows {
		rows = availableRows
		cols = max(1, int(float64(rows)*imageAspect/cellAspect))
	}
	cols = min(cols, width)

	return backend.Area{
		X:    1 + (width-cols)/2,
		Y:    2 + (availableRows-rows)/2,
		Cols: cols,
		Rows: rows,
	}
}
