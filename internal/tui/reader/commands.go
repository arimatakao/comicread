package reader

import (
	"image"
	"image/draw"
	"math"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/arimatakao/comicfile"
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
	cache := m.cache
	key := renderKey{
		pages: pages, areas: areas, width: m.width, height: m.height,
		zoom: zoom, scroll: scroll, view: m.bookView, renderer: renderer.Name(),
	}
	m.rendering = true
	m.renderError = nil

	return func() tea.Msg {
		// Kitty payloads carry image IDs allocated during encoding. Replaying one
		// would conflict with its placement lifecycle, so only cache source pages.
		if renderer.Name() != "kitty" && cache != nil {
			if output, ok := cache.render(key); ok {
				return pageRenderedMsg{requestID: requestID, page: page, area: area, output: output}
			}
		}

		output, err := renderOutput(chapter, renderer, cache, pages, areas, zoom, scroll)
		if err != nil {
			return pageRenderedMsg{requestID: requestID, page: page, err: err}
		}
		result := pageRenderedMsg{
			requestID: requestID,
			page:      page,
			area:      area,
			output:    output,
		}
		if renderer.Name() != "kitty" && cache != nil {
			cache.storeRender(key, result.output)
		}
		return result
	}
}

// preRenderNext prepares the next visible page or spread without emitting any
// terminal output. A later navigation can then use the render cache directly.
func (m Model) preRenderNext() tea.Cmd {
	if !m.canNextPage() || m.cache == nil {
		return nil
	}

	next := m
	next.page++
	next.scroll = 0

	return func() tea.Msg {
		if !next.cache.prefetchMu.TryLock() {
			return pagePrefetchedMsg{}
		}
		defer next.cache.prefetchMu.Unlock()

		// updateLayout obtains page aspect ratios through the decoded-image cache.
		// It runs here so opening the next page never delays the current display.
		next.updateLayout()
		pages := next.pageSlots()

		// Kitty output contains image IDs allocated while encoding. It cannot be
		// safely replayed later, but decoding now still avoids file I/O on paging.
		if next.backend.Name() == "kitty" {
			return pagePrefetchedMsg{}
		}

		key := renderKey{
			pages: pages, areas: next.pageAreas, width: next.width, height: next.height,
			zoom: next.zoom, scroll: next.scroll, view: next.bookView, renderer: next.backend.Name(),
		}
		if _, ok := next.cache.render(key); ok {
			return pagePrefetchedMsg{}
		}
		output, err := renderOutput(next.chapter, next.backend, next.cache, pages, next.pageAreas, next.zoom, next.scroll)
		if err == nil {
			next.cache.storeRender(key, output)
		}
		return pagePrefetchedMsg{}
	}
}

func renderOutput(chapter comicfile.ContainerReader, renderer backend.Renderer, cache *readerCache, pages [2]int, areas [2]backend.Area, zoom int, scroll float64) (string, error) {
	images := make([]image.Image, 0, len(pages))
	pageAreas := make([]backend.Area, 0, len(pages))
	for slot, pageIndex := range pages {
		if pageIndex < 0 {
			continue
		}
		var img image.Image
		var err error
		if cache == nil {
			img, err = chapter.Page(pageIndex)
		} else {
			img, err = cache.page(chapter, pageIndex)
		}
		if err != nil {
			return "", err
		}
		images = append(images, zoomedImage(img, zoom, scroll))
		pageAreas = append(pageAreas, areas[slot])
	}

	if spreadRenderer, ok := renderer.(backend.SpreadRenderer); ok && len(images) > 1 {
		return spreadRenderer.RenderSpread(images, pageAreas)
	}

	var output strings.Builder
	for index, img := range images {
		rendered, err := renderer.Render(img, pageAreas[index])
		if err != nil {
			return "", err
		}
		output.WriteString(rendered)
	}
	return output.String(), nil
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
