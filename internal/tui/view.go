package tui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m Model) View() tea.View {
	var content string
	if m.width < 1 || m.height < 2 {
		content = "comicread: terminal window is too small"
	} else {
		content = m.header() + strings.Repeat("\n", m.height-1)
	}

	view := tea.NewView(content)
	view.AltScreen = true
	view.WindowTitle = "comicread — " + m.title
	return view
}

func (m Model) header() string {
	page := fmt.Sprintf("pages %d/%d", m.page+1, m.chapter.TotalPages())
	pageWidth := len([]rune(page))
	if pageWidth > m.width {
		return fitLine(page, m.width)
	}

	if !m.rendering {
		return strings.Repeat(" ", (m.width-pageWidth)/2) + page
	}

	status := "rendering"
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
