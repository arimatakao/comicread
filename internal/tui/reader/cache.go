package reader

import (
	"sync"
	"sync/atomic"

	"github.com/arimatakao/comicread/internal/backend"
)

// readerCache holds only the prefetched next render. comicfile decodes pages
// lazily, while Model keeps at most the images needed for its current layout.
type readerCache struct {
	mu            sync.Mutex
	renderMu      sync.Mutex
	latestRequest atomic.Uint64
	key           renderKey
	output        string
	ready         bool
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

func newReaderCache() *readerCache {
	return &readerCache{}
}

func (c *readerCache) render(key renderKey) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if !c.ready || c.key != key {
		return "", false
	}
	return c.output, true
}

func (c *readerCache) storeRender(key renderKey, output string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.key = key
	c.output = output
	c.ready = true
}
