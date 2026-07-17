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

// itermImageReadyMsg is emitted after Bubble Tea has flushed a full redraw
// following iTerm2's display-wide image cleanup.
type itermImageReadyMsg struct {
	requestID uint64
	page      int
	output    string
}

// pagePrefetchedMsg marks completion of background page preparation. The
// payload itself is held in readerCache, so the update loop need not handle it.
type pagePrefetchedMsg struct{}

// ExternalQuitMsg requests the same graceful shutdown as the quit key. It is
// sent from outside the program when the process receives SIGINT or SIGTERM,
// so reading progress is saved and the displayed image is cleared before exit.
type ExternalQuitMsg struct{}
