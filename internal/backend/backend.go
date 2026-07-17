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

// SpreadRenderer can draw several images in one protocol command. Backends
// that do not implement it receive each page through Renderer.Render instead.
type SpreadRenderer interface {
	RenderSpread([]image.Image, []Area) (string, error)
}

// CellSizeRenderer uses the terminal's pixel dimensions for a character cell.
// Sixel rasters are sized in pixels, so they need this information to match a
// layout expressed in character cells.
type CellSizeRenderer interface {
	SetCellSize(width, height int)
	CellAspect() float64
}
