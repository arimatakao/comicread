package reader

import (
	"image"

	"github.com/arimatakao/comicread/internal/backend"
)

// ViewMode controls how pages are arranged in the reader.
type ViewMode uint8

const (
	SinglePageView ViewMode = iota
	BookView
	RightBookView
	CircleBookView
	RightCircleBookView
)

func (m Model) isBookView() bool {
	return m.bookView != SinglePageView
}

func (m Model) isRightToLeft() bool {
	return m.bookView == RightBookView || m.bookView == RightCircleBookView
}

func (m Model) isCircleBookView() bool {
	return m.bookView == CircleBookView || m.bookView == RightCircleBookView
}

// pageSlots returns the zero-based pages displayed in the left and right
// halves of the terminal. A negative value leaves that half empty.
func (m Model) pageSlots() [2]int {
	total := m.chapter.TotalPages()
	if total < 1 {
		return [2]int{-1, -1}
	}

	if !m.isBookView() {
		return [2]int{m.page, -1}
	}

	var pages [2]int
	if m.isCircleBookView() {
		right := m.circleRightPage(total)
		left := 0
		if m.page > 0 {
			left = m.circleRightPageForSpread(m.page-1, total)
		}
		pages = [2]int{left, right}
	} else {
		left := m.page * 2
		right := left + 1
		if right >= total {
			right = -1
		}
		pages = [2]int{left, right}
	}

	if m.isRightToLeft() {
		pages[0], pages[1] = pages[1], pages[0]
	}
	return pages
}

// circleRightPageForSpread follows the requested circular sequence:
// 1-2, 2-4, 4-6, with the final available page used when the next even page
// would exceed the chapter length.
func (m Model) circleRightPageForSpread(spread, total int) int {
	if spread == 0 {
		return min(1, total-1)
	}
	return min(spread*2+1, total-1)
}

func (m Model) circleRightPage(total int) int {
	return m.circleRightPageForSpread(m.page, total)
}

func (m Model) canNextPage() bool {
	total := m.chapter.TotalPages()
	if !m.isBookView() {
		return m.page+1 < total
	}
	if !m.isCircleBookView() {
		return (m.page+1)*2 < total
	}
	return m.circleRightPage(total) < total-1
}

func (m *Model) updateLayout() {
	slots := m.pageSlots()
	m.layoutPages = slots
	m.layoutImages = [2]image.Image{}
	m.pageAreas = [2]backend.Area{}
	for slot, page := range slots {
		if page < 0 {
			continue
		}
		img, err := m.chapter.Page(page)
		if err == nil {
			m.layoutImages[slot] = img
		}
	}
	if !m.isBookView() {
		m.area = m.imageArea(imageAspect(m.layoutImages[0]) * float64(m.zoom) / 100)
		m.pageAreas[0] = m.area
		return
	}

	for slot, page := range slots {
		if page < 0 {
			continue
		}
		m.pageAreas[slot] = m.bookImageArea(slot, imageAspect(m.layoutImages[slot])*float64(m.zoom)/100)
	}
	m.area = unionAreas(m.pageAreas[0], m.pageAreas[1])
}
