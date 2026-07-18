package reader

import (
	"strconv"
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
	if m.showingHelp {
		return "\n" + i18n.T(i18n.ReaderViewHelp)
	}
	if m.showingBookmarks {
		return m.bookmarkList()
	}
	return m.header() + strings.Repeat("\n", m.height-1)
}

func (m Model) header() string {
	if m.gotoPagePrefix {
		return fitLine(i18n.T(i18n.ReaderViewGoToPage, m.gotoPageInput), m.width)
	}

	page := m.progressBar()
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

func (m Model) bookmarkList() string {
	var content strings.Builder
	content.WriteString(i18n.T(i18n.ReaderViewBookmarks))
	content.WriteString("\n\n")
	if len(m.bookmarks) == 0 {
		content.WriteString(i18n.T(i18n.ReaderViewNoBookmarks))
		content.WriteString("\n\n")
	} else {
		for index, page := range m.bookmarks {
			prefix := "  "
			if index == m.bookmarkIndex {
				prefix = "> "
			}
			content.WriteString(prefix)
			content.WriteString(strconv.Itoa(page))
			content.WriteByte('\n')
		}
		content.WriteByte('\n')
	}
	content.WriteString(i18n.T(i18n.ReaderViewBookmarksHelp))
	return content.String()
}

func (m Model) progressBar() string {
	total := m.chapter.TotalPages()
	current := m.currentPageNumber()
	if total < 1 || current < 1 {
		return ""
	}

	const maxBarWidth = 14
	left, right := strconv.Itoa(current), strconv.Itoa(total)
	barWidth := min(maxBarWidth, max(1, m.width-len([]rune(left))-len([]rune(right))-5))
	var bar strings.Builder
	bar.Grow(barWidth + 2)
	bar.WriteRune('╟')
	for cell := range barWidth {
		firstPage := cell*total/barWidth + 1
		lastPage := min(total, (cell+1)*total/barWidth)
		switch {
		case m.hasBookmarkInRange(firstPage, lastPage):
			bar.WriteRune('^')
		case firstPage <= current:
			bar.WriteRune('─')
		default:
			bar.WriteRune('·')
		}
	}
	bar.WriteRune('╢')
	return left + " " + bar.String() + " " + right
}

func (m Model) hasBookmarkInRange(first, last int) bool {
	for _, bookmark := range m.bookmarks {
		if bookmark >= first && bookmark <= last {
			return true
		}
	}
	return false
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
