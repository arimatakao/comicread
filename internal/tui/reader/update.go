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
		m.area = imageArea(msg.Width, msg.Height, m.currentPageAspect())
		if m.area.Cols < 1 || m.area.Rows < 1 {
			m.status = i18n.T(i18n.ReaderStatusTerminalTooSmall)
			return m, tea.Raw(m.backend.Clear(m.displayedArea))
		}
		return m, m.renderPage()

	case tea.KeyPressMsg:
		key := msg.String()
		switch {
		case key == "q", key == "esc", key == "ctrl+c":
			return m, tea.Sequence(tea.Raw(m.backend.Clear(m.displayedArea)), tea.Quit)

		case isZoomInKey(msg):
			if m.zoomIn() {
				return m, m.renderPage()
			}

		case isZoomOutKey(msg):
			if m.zoomOut() {
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
				m.area = imageArea(m.width, m.height, m.currentPageAspect())
				return m, m.renderPage()
			}
			m.status = i18n.T(i18n.ReaderStatusLastPage)

		case isPreviousKey(key):
			if m.previousPage() {
				m.area = imageArea(m.width, m.height, m.currentPageAspect())
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
		return m, tea.Sequence(
			tea.Raw(m.backend.Clear(oldArea)),
			tea.Raw(msg.output),
		)
	}

	return m, nil
}

func (m *Model) nextPage() bool {
	if m.page+1 >= m.chapter.TotalPages() {
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
	img, err := m.chapter.Page(m.page)
	if err != nil || img.Bounds().Dy() == 0 {
		return 1
	}
	return float64(img.Bounds().Dx()) / float64(img.Bounds().Dy())
}

func imageArea(width, height int, imageAspect float64) backend.Area {
	const cellAspect = 0.5 // typical terminal cell width / height in pixels

	availableRows := height - 1 // one header row
	if width < 1 || availableRows < 1 {
		return backend.Area{}
	}
	if imageAspect <= 0 {
		imageAspect = 1
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
