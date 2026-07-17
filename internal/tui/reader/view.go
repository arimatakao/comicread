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
	page := m.pageLabel()
	pageWidth := len([]rune(page))
	if pageWidth > m.width {
		return m.withBookmarkMarker(fitLine(page, m.width))
	}

	if !m.rendering {
		return m.withBookmarkMarker(strings.Repeat(" ", (m.width-pageWidth)/2) + page)
	}

	status := i18n.T(i18n.ReaderViewRendering)
	statusWidth := len([]rune(status))
	pageStart := (m.width - pageWidth) / 2
	statusStart := m.width - statusWidth
	if pageStart+pageWidth >= statusStart {
		return m.withBookmarkMarker(strings.Repeat(" ", pageStart) + page)
	}

	return m.withBookmarkMarker(strings.Repeat(" ", pageStart) + page + strings.Repeat(" ", statusStart-pageStart-pageWidth) + status)
}

func (m Model) withBookmarkMarker(line string) string {
	if !m.isCurrentPageBookmarked() || m.width < 1 {
		return line
	}
	runes := []rune(line)
	if len(runes) == 0 {
		return "*"
	}
	runes[0] = '*'
	return string(runes)
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
