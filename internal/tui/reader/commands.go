package reader

import (
	"image"
	"image/draw"
	"math"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/arimatakao/comicread/internal/backend"
)

const layoutRenderDelay = time.Second / 50

// renderAfterLayout waits for Bubble Tea to flush a resized frame before
// emitting text-art escape sequences. Otherwise the first frame can erase the
// ASCII or Braille output that was written through tea.Raw.
func renderAfterLayout(layoutID uint64) tea.Cmd {
	return tea.Tick(layoutRenderDelay, func(time.Time) tea.Msg {
		return renderAfterLayoutMsg{layoutID: layoutID}
	})
}

func (m *Model) renderPage() tea.Cmd {
	if m.area.Cols < 1 || m.area.Rows < 1 {
		return nil
	}

	m.layoutPending = false
	m.requestID++
	requestID := m.requestID
	page := m.page
	pages := m.pageSlots()
	chapter := m.chapter
	renderer := m.backend
	zoom := m.zoom
	scroll := m.scroll
	areas := m.pageAreas
	area := m.area
	m.rendering = true
	m.renderError = nil

	return func() tea.Msg {
		var output strings.Builder
		for slot, pageIndex := range pages {
			if pageIndex < 0 {
				continue
			}
			img, err := chapter.Page(pageIndex)
			if err != nil {
				return pageRenderedMsg{requestID: requestID, page: page, err: err}
			}
			rendered, err := renderer.Render(zoomedImage(img, zoom, scroll), areas[slot])
			if err != nil {
				return pageRenderedMsg{requestID: requestID, page: page, err: err}
			}
			output.WriteString(rendered)
		}
		return pageRenderedMsg{
			requestID: requestID,
			page:      page,
			area:      area,
			output:    output.String(),
		}
	}
}

func (m Model) zoomedArea() backend.Area {
	aspect := m.currentPageAspect() * float64(m.zoom) / 100
	return imageArea(m.width, m.height, aspect)
}

// zoomedImage returns the visible vertical part of a page at the selected
// scale. It keeps the full source width, while scroll controls the crop's
// vertical position.
func zoomedImage(img image.Image, zoom int, scroll float64) image.Image {
	if zoom <= 100 {
		return img
	}

	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	if width < 1 || height < 1 {
		return img
	}

	cropHeight := max(1, height*100/zoom)
	maxY := height - cropHeight
	y := bounds.Min.Y + int(math.Round(float64(maxY)*min(1, max(0, scroll))))
	crop := image.Rect(bounds.Min.X, y, bounds.Max.X, y+cropHeight)

	visible := image.NewRGBA(image.Rect(0, 0, width, cropHeight))
	draw.Draw(visible, visible.Bounds(), img, crop.Min, draw.Src)
	return visible
}
