package tui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m Model) View() tea.View {
	var content string
	if m.width < 1 || m.height < 3 {
		content = "comicread: terminal window is too small"
	} else {
		header := fitLine(
			fmt.Sprintf(" %s  %d/%d  [%s]", m.title, m.page+1, m.chapter.TotalPages(), m.backend.Name()),
			m.width,
		)
		footer := fitLine(" ←/h previous  →/l/space next  q quit  ·  "+m.status, m.width)
		content = header + "\n" + strings.Repeat("\n", m.height-2) + footer
	}

	view := tea.NewView(content)
	view.AltScreen = true
	view.WindowTitle = "comicread — " + m.title
	return view
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
