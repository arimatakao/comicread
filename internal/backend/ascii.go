package backend

import (
	"fmt"
	"image"
	"strings"

	"github.com/fandasy/ASCIIimage/v2/core"
	"golang.org/x/image/draw"
)

// ASCII renders an image as ANSI-styled ASCII characters. It works in terminals
// that do not implement a graphical image protocol.
type ASCII struct {
}

func NewASCII() *ASCII { return &ASCII{} }

func (*ASCII) Name() string { return "ascii" }

func (*ASCII) Render(img image.Image, area Area) (string, error) {
	if img == nil {
		return "", fmt.Errorf("render nil image")
	}
	if area.X < 1 || area.Y < 1 || area.Cols < 1 || area.Rows < 1 {
		return "", fmt.Errorf("invalid terminal area: %+v", area)
	}

	// Core supplies the luminance-to-character gradient. Its image-oriented
	// generator is not suitable for a text terminal, so we use that same public
	// gradient while emitting terminal cells directly.
	chars := core.DefaultOptions().Chars
	resized := image.NewNRGBA(image.Rect(0, 0, area.Cols, area.Rows))
	// Nearest-neighbour preserves inked edges. Smoother filters introduce many
	// intermediate shades, which look like visual noise in ASCII art.
	draw.NearestNeighbor.Scale(resized, resized.Bounds(), img, img.Bounds(), draw.Over, nil)
	chars = core.NewChars("@#*:. ")

	var output strings.Builder
	output.Grow(area.Cols * area.Rows * 20)
	output.WriteString("\x1b7")
	for y := range area.Rows {
		fmt.Fprintf(&output, "\x1b[%d;%dH", area.Y+y, area.X)
		for x := range area.Cols {
			r, g, b, alpha := resized.At(x, y).RGBA()
			r, g, b = blendOverBlack(r, g, b, alpha)
			brightness := asciiBrightness(r, g, b)
			r, g, b = contrastColor(r, g, b, brightness)
			face := uint32(255)
			if brightness > 127 {
				face = 0
			}
			fmt.Fprintf(&output, "\x1b[38;2;%d;%d;%d;48;2;%d;%d;%dm%c", face, face, face, r>>8, g>>8, b>>8, chars[brightness])
		}
		output.WriteString("\x1b[0m")
	}
	output.WriteString("\x1b8")
	return output.String(), nil
}

// Clear erases only the previous text-art rectangle, preserving the reader
// header and Bubble Tea's alternate-screen layout.
func (*ASCII) Clear(area Area) string {
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

// asciiBrightness converts a pixel into a high-contrast perceptual luminance.
// It suppresses nearly white paper and nearly black ink, making page contours
// readable at a terminal's limited character resolution.
func asciiBrightness(r, g, b uint32) uint32 {
	luma := (299*r + 587*g + 114*b) / (1000 * 257)
	switch {
	case luma <= 32:
		return 0
	case luma >= 224:
		return 255
	default:
		return (luma - 32) * 255 / 192
	}
}

// contrastColor preserves the original hue while applying the expanded
// luminance used by asciiBrightness.
func contrastColor(r, g, b, brightness uint32) (uint32, uint32, uint32) {
	luma := (299*r + 587*g + 114*b) / (1000 * 257)
	if luma == 0 {
		return 0, 0, 0
	}
	return min(0xffff, r*brightness/luma), min(0xffff, g*brightness/luma), min(0xffff, b*brightness/luma)
}

func blendOverBlack(r, g, b, a uint32) (uint32, uint32, uint32) {
	if a == 0xffff {
		return r, g, b
	}
	return r * a / 0xffff, g * a / 0xffff, b * a / 0xffff
}

var _ Renderer = (*ASCII)(nil)
