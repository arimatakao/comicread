package backend

import (
	"fmt"
	"image"
	"strings"

	"github.com/imjasonh/dots"
)

// Dots renders images as ANSI-colored Unicode Braille characters. Each
// terminal cell represents a 2 by 4 pixel block, providing more detail than
// ASCII art while requiring no graphics protocol.
type Dots struct{}

func NewDots() *Dots { return &Dots{} }

func (*Dots) Name() string { return "dots" }

func (*Dots) Render(img image.Image, area Area) (string, error) {
	if img == nil {
		return "", fmt.Errorf("render nil image")
	}
	if area.X < 1 || area.Y < 1 || area.Cols < 1 || area.Rows < 1 {
		return "", fmt.Errorf("invalid terminal area: %+v", area)
	}

	lines := dots.Convert(img, dots.Options{
		Width:     area.Cols,
		Height:    area.Rows,
		Threshold: 115,
	})

	var output strings.Builder
	output.Grow(area.Cols * area.Rows * 12)
	output.WriteString("\x1b7")
	for row, line := range lines {
		fmt.Fprintf(&output, "\x1b[%d;%dH%s\x1b[0m", area.Y+row, area.X, line)
	}
	output.WriteString("\x1b8")
	return output.String(), nil
}

// Clear erases the terminal cells occupied by the prior Braille rendering.
func (*Dots) Clear(area Area) string {
	if area.X < 1 || area.Y < 1 || area.Cols < 1 || area.Rows < 1 {
		return ""
	}

	var output strings.Builder
	output.Grow(area.Rows * 16)
	output.WriteString("\x1b7\x1b[0m")
	for y := range area.Rows {
		fmt.Fprintf(&output, "\x1b[%d;%dH\x1b[%dX", area.Y+y, area.X, area.Cols)
	}
	output.WriteString("\x1b8")
	return output.String()
}

var _ Renderer = (*Dots)(nil)
