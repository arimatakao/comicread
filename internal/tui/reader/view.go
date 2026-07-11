package reader

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/arimatakao/comicread/internal/i18n"
)

func (m Model) View() tea.View {
	view := tea.NewView(m.content())
	view.AltScreen = true
	view.WindowTitle = i18n.T(i18n.ReaderViewWindowTitle, m.title)
	return view
}

func (m Model) content() string {
	if m.width < 1 || m.height < 2 {
		return i18n.T(i18n.ReaderViewTerminalTooSmall)
	}
	return m.header() + strings.Repeat("\n", m.height-1)
}

// repaintText restores Bubble Tea's text layer after iTerm2 erases the whole
// display to remove a previous inline image.
func (m Model) repaintText() string {
	return "\x1b[H" + m.content()
}

func (m Model) header() string {
	page := m.pageLabel()
	pageWidth := len([]rune(page))
	if pageWidth > m.width {
		return fitLine(page, m.width)
	}

	if !m.rendering {
		return strings.Repeat(" ", (m.width-pageWidth)/2) + page
	}

	status := i18n.T(i18n.ReaderViewRendering)
	statusWidth := len([]rune(status))
	pageStart := (m.width - pageWidth) / 2
	statusStart := m.width - statusWidth
	if pageStart+pageWidth >= statusStart {
		return strings.Repeat(" ", pageStart) + page
	}

	return strings.Repeat(" ", pageStart) + page + strings.Repeat(" ", statusStart-pageStart-pageWidth) + status
}

func (m Model) pageLabel() string {
	slots := m.pageSlots()
	first, last := -1, -1
	for _, page := range slots {
		if page < 0 {
			continue
		}
		if first < 0 || page < first {
			first = page
		}
		if last < 0 || page > last {
			last = page
		}
	}
	if first == last {
		return i18n.T(i18n.ReaderViewPages, first+1, m.chapter.TotalPages())
	}
	return i18n.T(i18n.ReaderViewPageRange, first+1, last+1, m.chapter.TotalPages())
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
