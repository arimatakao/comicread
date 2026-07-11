package backend

import (
	"bytes"
	"fmt"
	"image"
	"image/color/palette"
	stdraw "image/draw"
	"strings"

	"github.com/BourgeoisBear/rasterm"
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

// RenderSpread draws a book spread as one raster. Some Sixel terminals do not
// preserve horizontal cursor placement between consecutive sixel commands, so
// emitting a single raster prevents the second page from covering the first.
func (s *Sixel) RenderSpread(images []image.Image, areas []Area) (string, error) {
	if len(images) != len(areas) || len(images) < 1 {
		return "", fmt.Errorf("invalid sixel spread")
	}

	spread := unionArea(areas)
	if spread.Cols < 1 || spread.Rows < 1 {
		return "", fmt.Errorf("invalid sixel spread area: %+v", spread)
	}

	canvas := image.NewNRGBA(image.Rect(0, 0, spread.Cols*sixelCellWidth, spread.Rows*sixelCellHeight))
	for index, img := range images {
		area := areas[index]
		if img == nil || area.Cols < 1 || area.Rows < 1 {
			return "", fmt.Errorf("invalid sixel spread page")
		}
		rect := image.Rect(
			(area.X-spread.X)*sixelCellWidth,
			(area.Y-spread.Y)*sixelCellHeight,
			(area.X-spread.X+area.Cols)*sixelCellWidth,
			(area.Y-spread.Y+area.Rows)*sixelCellHeight,
		)
		draw.CatmullRom.Scale(canvas, rect, img, img.Bounds(), draw.Over, nil)
	}

	return s.Render(canvas, spread)
}

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

	paletted := image.NewPaletted(resized.Bounds(), palette.Plan9)
	stdraw.FloydSteinberg.Draw(paletted, paletted.Bounds(), resized, image.Point{})

	var payload bytes.Buffer
	if err := rasterm.SixelWriteImage(&payload, paletted); err != nil {
		return "", fmt.Errorf("encode sixel image: %w", err)
	}

	var output strings.Builder
	output.Grow(payload.Len() + 32)
	// Save and restore the text cursor: a Sixel image remains at its placement
	// while Bubble Tea continues to own the text UI cursor.
	output.WriteString("\x1b7")
	fmt.Fprintf(&output, "\x1b[%d;%dH", area.Y, area.X)
	output.Write(payload.Bytes())
	output.WriteString("\x1b8")
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

func unionArea(areas []Area) Area {
	var result Area
	for _, area := range areas {
		if area.Cols < 1 || area.Rows < 1 {
			continue
		}
		if result.Cols < 1 || result.Rows < 1 {
			result = area
			continue
		}
		left := min(result.X, area.X)
		top := min(result.Y, area.Y)
		right := max(result.X+result.Cols, area.X+area.Cols)
		bottom := max(result.Y+result.Rows, area.Y+area.Rows)
		result = Area{X: left, Y: top, Cols: right - left, Rows: bottom - top}
	}
	return result
}

var _ Renderer = (*Sixel)(nil)
var _ SpreadRenderer = (*Sixel)(nil)
