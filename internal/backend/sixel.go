package backend

import (
	"bytes"
	"fmt"
	"image"
	"strings"

	"github.com/charmbracelet/x/ansi/sixel"
	"golang.org/x/image/draw"
)

const (
	sixelCellWidth  = 8
	sixelCellHeight = 16
)

// Sixel renders images with the DEC Sixel graphics protocol. It assumes the
// common 8 by 16 pixel terminal cell, which matches the layout calculation in
// the TUI.
type Sixel struct{}

func NewSixel() *Sixel { return &Sixel{} }

func (*Sixel) Name() string { return "sixel" }

func (*Sixel) Render(img image.Image, area Area) (string, error) {
	if img == nil {
		return "", fmt.Errorf("render nil image")
	}
	if area.X < 1 || area.Y < 1 || area.Cols < 1 || area.Rows < 1 {
		return "", fmt.Errorf("invalid terminal area: %+v", area)
	}

	width, height := area.Cols*sixelCellWidth, area.Rows*sixelCellHeight
	resized := image.NewNRGBA(image.Rect(0, 0, width, height))
	draw.CatmullRom.Scale(resized, resized.Bounds(), img, img.Bounds(), draw.Over, nil)

	var payload bytes.Buffer
	if err := (&sixel.Encoder{}).Encode(&payload, resized); err != nil {
		return "", fmt.Errorf("encode sixel image: %w", err)
	}

	var output strings.Builder
	output.Grow(payload.Len() + 32)
	// Save and restore the text cursor: a Sixel image remains at its placement
	// while Bubble Tea continues to own the text UI cursor.
	output.WriteString("\x1b7")
	fmt.Fprintf(&output, "\x1b[%d;%dH\x1bP0;1q", area.Y, area.X)
	output.Write(payload.Bytes())
	output.WriteString("\x1b\\\x1b8")
	return output.String(), nil
}

// Clear replaces the prior image rectangle with the terminal background,
// preserving the header and footer Bubble Tea draws in the alternate screen.
func (*Sixel) Clear(area Area) string {
	if area.X < 1 || area.Y < 1 || area.Cols < 1 || area.Rows < 1 {
		return ""
	}

	width, height := area.Cols*sixelCellWidth, area.Rows*sixelCellHeight
	// CSI erase commands only affect text cells. A Sixel raster with P2=0 and
	// zero-valued pixels paints the entire raster with the terminal background
	// color. Some terminals ignore a raster that contains attributes only, so
	// emit a zero-filled Sixel band for every six pixel rows.
	var output strings.Builder
	output.Grow(height/6*12 + 48)
	output.WriteString("\x1b7")
	fmt.Fprintf(&output, "\x1b[%d;%dH\x1bP0;0;0q\"1;1;%d;%d", area.Y, area.X, width, height)
	for top := 0; top < height; top += 6 {
		writeSixelRepeat(&output, width, '?')
		if top+6 < height {
			output.WriteByte('-')
		}
	}
	output.WriteString("\x1b\\\x1b8")
	return output.String()
}

func writeSixelRepeat(output *strings.Builder, count int, value byte) {
	for count > 255 {
		fmt.Fprintf(output, "!255%c", value)
		count -= 255
	}
	if count > 3 {
		fmt.Fprintf(output, "!%d%c", count, value)
		return
	}
	for range count {
		output.WriteByte(value)
	}
}

var _ Renderer = (*Sixel)(nil)
