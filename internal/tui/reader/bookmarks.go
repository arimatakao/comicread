package reader

import "slices"

func (m Model) currentPageNumber() int {
	page := -1
	for _, slot := range m.pageSlots() {
		if slot >= 0 && (page < 0 || slot < page) {
			page = slot
		}
	}
	return page + 1
}

func (m *Model) setPageNumber(page int) bool {
	if page < 1 || m.chapter.TotalPages() < 1 {
		return false
	}
	page = min(page, m.chapter.TotalPages())
	target := page - 1
	if !m.isBookView() {
		m.page = target
		return true
	}

	for cursor := 0; cursor < m.chapter.TotalPages(); cursor++ {
		m.page = cursor
		slots := m.pageSlots()
		if slices.Contains(slots[:], target) {
			return true
		}
		if !m.canNextPage() {
			break
		}
	}
	return false
}

func (m *Model) saveCurrentPage() {
	if m.journal == nil {
		return
	}
	if err := m.journal.SetLastOpenedPage(m.currentPageNumber()); err != nil {
		m.status = err.Error()
	}
}

func (m *Model) toggleBookmark() {
	if m.journal == nil {
		return
	}
	if _, err := m.journal.ToggleBookmark(m.currentPageNumber()); err != nil {
		m.status = err.Error()
		return
	}
	m.bookmarks = m.journal.Bookmarks()
	if m.bookmarkIndex >= len(m.bookmarks) {
		m.bookmarkIndex = max(0, len(m.bookmarks)-1)
	}
}

func (m Model) isCurrentPageBookmarked() bool {
	return slices.Contains(m.bookmarks, m.currentPageNumber())
}

func (m *Model) moveBookmark(forward bool) bool {
	if len(m.bookmarks) == 0 {
		return false
	}
	current := m.currentPageNumber()
	index := 0
	if forward {
		index = len(m.bookmarks) - 1
		for i, page := range m.bookmarks {
			if page > current {
				index = i
				break
			}
		}
	} else {
		for i := len(m.bookmarks) - 1; i >= 0; i-- {
			if m.bookmarks[i] < current {
				index = i
				break
			}
		}
	}
	if !m.setPageNumber(m.bookmarks[index]) {
		return false
	}
	m.bookmarkIndex = index
	m.scroll = 0
	m.saveCurrentPage()
	return true
}

func (m *Model) openBookmarks() {
	m.bookmarkPrefix = false
	m.bookmarkListPrefix = false
	m.showingBookmarks = true
	current := m.currentPageNumber()
	m.bookmarkIndex = 0
	for i, page := range m.bookmarks {
		if page >= current {
			m.bookmarkIndex = i
			break
		}
	}
}

func (m *Model) closeBookmarks() {
	m.showingBookmarks = false
	m.bookmarkPrefix = false
	m.bookmarkListPrefix = false
}

func (m *Model) selectBookmark() bool {
	if len(m.bookmarks) == 0 {
		return false
	}
	m.bookmarkIndex = min(max(0, m.bookmarkIndex), len(m.bookmarks)-1)
	if !m.setPageNumber(m.bookmarks[m.bookmarkIndex]) {
		return false
	}
	m.scroll = 0
	m.saveCurrentPage()
	return true
}
