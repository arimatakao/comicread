package backend

import (
	"fmt"
	"image"
	"strings"

	"golang.org/x/image/draw"
)

const (
	sixelCellWidth  = 8
	sixelCellHeight = 16
	sixelColors     = 256
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
	indexed, palette, used := sixelImage(img, width, height)

	var output strings.Builder
	output.Grow(width * height / 2)
	// Save and restore the text cursor: a Sixel image remains at its placement
	// while Bubble Tea continues to own the text UI cursor.
	output.WriteString("\x1b7")
	fmt.Fprintf(&output, "\x1b[%d;%dH\x1bP0;0;0q\"1;1;%d;%d", area.Y, area.X, width, height)
	for index, color := range palette {
		if !used[index] {
			continue
		}
		fmt.Fprintf(&output, "#%d;2;%d;%d;%d", index, color.r, color.g, color.b)
	}

	for y := 0; y < height; y += 6 {
		for index := range palette {
			if !used[index] {
				continue
			}
			line := sixelLine(indexed, width, height, y, int16(index))
			if line == "" {
				continue
			}
			fmt.Fprintf(&output, "#%d%s$", index, line)
		}
		output.WriteByte('-')
	}
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

type sixelColor struct{ r, g, b int }

func sixelImage(source image.Image, width, height int) ([]int16, []sixelColor, []bool) {
	resized := image.NewNRGBA(image.Rect(0, 0, width, height))
	draw.CatmullRom.Scale(resized, resized.Bounds(), source, source.Bounds(), draw.Over, nil)

	indexed := make([]int16, width*height)
	used := make([]bool, sixelColors)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel := resized.NRGBAAt(x, y)
			if pixel.A == 0 {
				indexed[y*width+x] = -1
				continue
			}
			index := sixelIndex(pixel.R, pixel.G, pixel.B)
			indexed[y*width+x] = int16(index)
			used[index] = true
		}
	}

	palette := make([]sixelColor, sixelColors)
	for index := range palette {
		palette[index] = sixelPalette(uint8(index))
	}
	return indexed, palette, used
}

func sixelIndex(r, g, b uint8) uint8 {
	return (r>>5)<<5 | (g>>5)<<2 | b>>6
}

func sixelPalette(index uint8) sixelColor {
	r := int(index>>5) * 100 / 7
	g := int((index>>2)&7) * 100 / 7
	b := int(index&3) * 100 / 3
	return sixelColor{r: r, g: g, b: b}
}

func sixelLine(indexed []int16, width, height, top int, color int16) string {
	values := make([]byte, width)
	hasPixel := false
	for x := range values {
		for bit := 0; bit < 6 && top+bit < height; bit++ {
			if indexed[(top+bit)*width+x] == color {
				values[x] |= 1 << bit
				hasPixel = true
			}
		}
		values[x] += '?'
	}
	if !hasPixel {
		return ""
	}

	var output strings.Builder
	for start := 0; start < len(values); {
		end := start + 1
		for end < len(values) && values[end] == values[start] {
			end++
		}
		if count := end - start; count > 3 {
			writeSixelRepeat(&output, count, values[start])
		} else {
			output.WriteString(string(values[start:end]))
		}
		start = end
	}
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
