package reader

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/arimatakao/comicread/internal/i18n"
)

func (m Model) View() tea.View {
	var content string
	if m.width < 1 || m.height < 2 {
		content = i18n.T("reader.view.terminal_too_small")
	} else {
		content = m.header() + strings.Repeat("\n", m.height-1)
	}

	view := tea.NewView(content)
	view.AltScreen = true
	view.WindowTitle = i18n.T("reader.view.window_title", m.title)
	return view
}

func (m Model) header() string {
	page := i18n.T("reader.view.pages", m.page+1, m.chapter.TotalPages())
	pageWidth := len([]rune(page))
	if pageWidth > m.width {
		return fitLine(page, m.width)
	}

	if !m.rendering {
		return strings.Repeat(" ", (m.width-pageWidth)/2) + page
	}

	status := i18n.T("reader.view.rendering")
	statusWidth := len([]rune(status))
	pageStart := (m.width - pageWidth) / 2
	statusStart := m.width - statusWidth
	if pageStart+pageWidth >= statusStart {
		return strings.Repeat(" ", pageStart) + page
	}

	return strings.Repeat(" ", pageStart) + page + strings.Repeat(" ", statusStart-pageStart-pageWidth) + status
}

func fitLine(value string, width int) string {
	if width <= 0 {
		return ""
	}
	runes := []rune(value)
	if len(runes) <= width {
		return value
	}
	if width == 1 {
		return string(runes[:1])
	}
	return string(runes[:width-1]) + "…"
}
