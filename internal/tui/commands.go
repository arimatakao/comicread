package tui

import tea "charm.land/bubbletea/v2"

func (m *Model) renderPage() tea.Cmd {
	if m.area.Cols < 1 || m.area.Rows < 1 {
		return nil
	}

	m.requestID++
	requestID := m.requestID
	page := m.page
	chapter := m.chapter
	renderer := m.backend
	area := m.area
	m.rendering = true
	m.renderError = nil

	return func() tea.Msg {
		img, err := chapter.Page(page)
		if err != nil {
			return pageRenderedMsg{requestID: requestID, page: page, err: err}
		}

		output, err := renderer.Render(img, area)
		return pageRenderedMsg{
			requestID: requestID,
			page:      page,
			area:      area,
			output:    output,
			err:       err,
		}
	}
}
