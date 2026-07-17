package backend

import (
	"image"

	"golang.org/x/image/draw"
)

const (
	// Encode at 2x the common 8x16 cell size. This keeps Kitty/iTerm2 sharp on
	// HiDPI displays without returning to full-resolution page payloads.
	terminalCellWidth  = 16
	terminalCellHeight = 32
)

// scaleForTerminal bounds protocol encoding work to the pixels the terminal
// can display. Kitty and iTerm2 otherwise encode the full source image even
// when the protocol asks the terminal to show it in a small cell rectangle.
func scaleForTerminal(img image.Image, area Area) image.Image {
	width := area.Cols * terminalCellWidth
	height := area.Rows * terminalCellHeight
	bounds := img.Bounds()
	if width < 1 || height < 1 || (bounds.Dx() <= width && bounds.Dy() <= height) {
		return img
	}

	resized := image.NewNRGBA(image.Rect(0, 0, width, height))
	draw.CatmullRom.Scale(resized, resized.Bounds(), img, bounds, draw.Over, nil)
	return resized
}
