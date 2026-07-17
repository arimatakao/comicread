package reader

import (
	"image"
	"sync"

	"github.com/arimatakao/comicfile"
	"github.com/arimatakao/comicread/internal/backend"
)

const renderCacheSize = 4

// readerCache retains every source page opened during this reader session.
// Source images are independent from zoom and terminal size; rendered protocol
// payloads are not, so they remain a small LRU.
type readerCache struct {
	mu         sync.Mutex
	loadMu     sync.Mutex
	prefetchMu sync.Mutex
	images     map[int]image.Image
	renders    map[renderKey]string
	renderLRU  []renderKey
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
	return &readerCache{
		images:  make(map[int]image.Image),
		renders: make(map[renderKey]string),
	}
}

func (c *readerCache) page(chapter comicfile.ContainerReader, index int) (image.Image, error) {
	c.mu.Lock()
	if img, ok := c.images[index]; ok {
		c.mu.Unlock()
		return img, nil
	}
	c.mu.Unlock()

	// Container readers are not assumed to support concurrent reads. Keep the
	// potentially slow decode outside the map lock, but serialize it with other
	// page loads from foreground rendering and prefetching.
	c.loadMu.Lock()
	defer c.loadMu.Unlock()
	c.mu.Lock()
	if img, ok := c.images[index]; ok {
		c.mu.Unlock()
		return img, nil
	}
	c.mu.Unlock()

	img, err := chapter.Page(index)
	if err != nil {
		return nil, err
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	if cached, ok := c.images[index]; ok {
		return cached, nil
	}
	c.images[index] = img
	return img, nil
}

func (c *readerCache) render(key renderKey) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	output, ok := c.renders[key]
	if !ok {
		return "", false
	}
	for i, cached := range c.renderLRU {
		if cached == key {
			c.renderLRU = append(c.renderLRU[:i], c.renderLRU[i+1:]...)
			break
		}
	}
	c.renderLRU = append(c.renderLRU, key)
	return output, true
}

func (c *readerCache) storeRender(key renderKey, output string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, exists := c.renders[key]; exists {
		c.renders[key] = output
		return
	}
	c.renders[key] = output
	c.renderLRU = append(c.renderLRU, key)
	if len(c.renderLRU) > renderCacheSize {
		evicted := c.renderLRU[0]
		delete(c.renders, evicted)
		c.renderLRU = c.renderLRU[1:]
	}
}
