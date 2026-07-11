package reader

import "github.com/arimatakao/comicread/internal/backend"

type pageRenderedMsg struct {
	requestID uint64
	page      int
	area      backend.Area
	output    string
	err       error
}
