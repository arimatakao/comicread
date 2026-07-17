package reader

import "github.com/arimatakao/comicread/internal/backend"

type pageRenderedMsg struct {
	requestID uint64
	page      int
	area      backend.Area
	output    string
	err       error
}

type renderAfterLayoutMsg struct {
	layoutID uint64
}

// pagePrefetchedMsg marks completion of background page preparation. The
// payload itself is held in readerCache, so the update loop need not handle it.
type pagePrefetchedMsg struct{}
