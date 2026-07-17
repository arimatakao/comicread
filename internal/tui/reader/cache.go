package reader

import (
	"sync"
	"sync/atomic"

	"github.com/arimatakao/comicread/internal/backend"
)

// readerCache holds a bounded set of prefetched renders. comicfile decodes
// pages lazily, while Model keeps at most the images needed for its layout.
type readerCache struct {
	mu            sync.Mutex
	renderMu      sync.Mutex
	latestRequest atomic.Uint64
	limit         int
	renders       map[renderKey]string
	order         []renderKey
}

type renderKey struct {
	pages    [2]int
	areas    [2]backend.Area
	width    int
	height   int
	zoom     int
	scroll   float64
	view     ViewMode
	renderer string
}

func newReaderCache(limit int) *readerCache {
	return &readerCache{limit: max(1, limit), renders: make(map[renderKey]string)}
}

func (m Model) currentRenderKey() renderKey {
	return renderKey{
		pages: m.pageSlots(), areas: m.pageAreas, width: m.width, height: m.height,
		zoom: m.zoom, scroll: m.scroll, view: m.bookView, renderer: m.backend.Name(),
	}
}

// cachedRender returns a stored render when replaying it is safe. Kitty
// payloads carry image IDs allocated during encoding; replaying one would
// conflict with its placement lifecycle, so kitty always renders freshly.
func (m Model) cachedRender(key renderKey) (string, bool) {
	if m.cache == nil || m.backend.Name() == "kitty" {
		return "", false
	}
	return m.cache.render(key)
}

func (m Model) storeCachedRender(key renderKey, output string) {
	if m.cache != nil && m.backend.Name() != "kitty" {
		m.cache.storeRender(key, output)
	}
}

func (c *readerCache) render(key renderKey) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	output, ok := c.renders[key]
	if !ok {
		return "", false
	}
	c.touch(key)
	return output, true
}

func (c *readerCache) storeRender(key renderKey, output string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, exists := c.renders[key]; exists {
		c.renders[key] = output
		c.touch(key)
		return
	}
	c.renders[key] = output
	c.order = append(c.order, key)
	if len(c.order) > c.limit {
		evicted := c.order[0]
		delete(c.renders, evicted)
		c.order = c.order[1:]
	}
}

func (c *readerCache) touch(key renderKey) {
	for index, cached := range c.order {
		if cached != key {
			continue
		}
		copy(c.order[index:], c.order[index+1:])
		c.order[len(c.order)-1] = key
		return
	}
}
