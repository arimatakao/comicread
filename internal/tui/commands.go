package tui

import (
	"image"
	"image/draw"
	"math"

	tea "charm.land/bubbletea/v2"
	"github.com/arimatakao/comicread/internal/backend"
)

func (m *Model) renderPage() tea.Cmd {
	if m.area.Cols < 1 || m.area.Rows < 1 {
		return nil
	}

	m.requestID++
	requestID := m.requestID
	page := m.page
	chapter := m.chapter
	renderer := m.backend
	zoom := m.zoom
	scroll := m.scroll
	area := m.zoomedArea()
	m.rendering = true
	m.renderError = nil

	return func() tea.Msg {
		img, err := chapter.Page(page)
		if err != nil {
			return pageRenderedMsg{requestID: requestID, page: page, err: err}
		}

		output, err := renderer.Render(zoomedImage(img, zoom, scroll), area)
		return pageRenderedMsg{
			requestID: requestID,
			page:      page,
			area:      area,
			output:    output,
			err:       err,
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
