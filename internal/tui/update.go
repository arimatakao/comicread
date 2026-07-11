package tui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/arimatakao/comicread/internal/backend"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.area = imageArea(msg.Width, msg.Height, m.currentPageAspect())
		if m.area.Cols < 1 || m.area.Rows < 1 {
			m.status = "terminal window is too small"
			return m, tea.Raw(m.backend.Clear(m.displayedArea))
		}
		return m, m.renderPage()

	case tea.KeyPressMsg:
		key := msg.String()
		switch {
		case key == "q", key == "esc", key == "ctrl+c":
			return m, tea.Sequence(tea.Raw(m.backend.Clear(m.displayedArea)), tea.Quit)

		case isNextKey(key):
			if m.nextPage() {
				m.area = imageArea(m.width, m.height, m.currentPageAspect())
				return m, m.renderPage()
			}
			m.status = "last page"

		case isPreviousKey(key):
			if m.previousPage() {
				m.area = imageArea(m.width, m.height, m.currentPageAspect())
				return m, m.renderPage()
			}
			m.status = "first page"
		}

	case pageRenderedMsg:
		if msg.requestID != m.requestID || msg.page != m.page {
			return m, nil
		}

		m.rendering = false
		if msg.err != nil {
			m.renderError = msg.err
			m.status = fmt.Sprintf("render error: %v", msg.err)
			return m, nil
		}

		m.status = fmt.Sprintf("page %d/%d", m.page+1, m.chapter.TotalPages())
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
	return true
}

func (m *Model) previousPage() bool {
	if m.page == 0 {
		return false
	}
	m.page--
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

	availableRows := height - 2 // one header and one footer row
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
