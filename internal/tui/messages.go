package tui

type pageRenderedMsg struct {
	requestID uint64
	page      int
	output    string
	err       error
}
