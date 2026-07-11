package backend

import "image"

// Area is a 1-based terminal rectangle measured in character cells.
type Area struct {
	X, Y       int
	Cols, Rows int
}

// Renderer converts an image into terminal protocol escape sequences.
// It never writes to stdout: Bubble Tea remains the sole output owner.
type Renderer interface {
	Name() string
	Render(image.Image, Area) (string, error)
	Clear(Area) string
}
