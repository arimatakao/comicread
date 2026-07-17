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

func (m Model) pageAspect(page int) float64 {
	img, err := m.cachedPage(page)
	if err != nil || img.Bounds().Dy() == 0 {
		return 1
	}
	return float64(img.Bounds().Dx()) / float64(img.Bounds().Dy())
}

func (m Model) cachedPage(page int) (image.Image, error) {
	if m.cache == nil {
		return m.chapter.Page(page)
	}
	return m.cache.page(m.chapter, page)
}

func (m *Model) updateLayout() {
	m.pageAreas = [2]backend.Area{}
	if !m.isBookView() {
		m.area = imageArea(m.width, m.height, m.currentPageAspect()*float64(m.zoom)/100)
		m.pageAreas[0] = m.area
		return
	}

	slots := m.pageSlots()
	for slot, page := range slots {
		if page < 0 {
			continue
		}
		m.pageAreas[slot] = bookImageArea(m.width, m.height, slot, m.pageAspect(page)*float64(m.zoom)/100)
	}
	m.area = unionAreas(m.pageAreas[0], m.pageAreas[1])
}

func bookImageArea(width, height, slot int, imageAspect float64) backend.Area {
	const gap = 1

	availableRows := height - 1
	if width < 2 || availableRows < 1 {
		return backend.Area{}
	}
	availableCols := width - gap
	leftCols := availableCols / 2
	rightCols := availableCols - leftCols
	panelCols, panelX := leftCols, 1
	if slot == 1 {
		panelCols = rightCols
		panelX += leftCols + gap
	}
	if panelCols < 1 {
		return backend.Area{}
	}
	if imageAspect <= 0 {
		imageAspect = 1
	}

	cols := panelCols
	rows := max(1, int(float64(cols)*0.5/imageAspect))
	if rows > availableRows {
		rows = availableRows
		cols = max(1, int(float64(rows)*imageAspect/0.5))
		cols = min(cols, panelCols)
	}
	return backend.Area{
		X:    panelX + (panelCols-cols)/2,
		Y:    2 + (availableRows-rows)/2,
		Cols: cols,
		Rows: rows,
	}
}

func unionAreas(first, second backend.Area) backend.Area {
	if first.Cols < 1 || first.Rows < 1 {
		return second
	}
	if second.Cols < 1 || second.Rows < 1 {
		return first
	}
	left := min(first.X, second.X)
	top := min(first.Y, second.Y)
	right := max(first.X+first.Cols, second.X+second.Cols)
	bottom := max(first.Y+first.Rows, second.Y+second.Rows)
	return backend.Area{X: left, Y: top, Cols: right - left, Rows: bottom - top}
}
