package backend

import (
	"bytes"
	"fmt"
	"image"
	"strconv"

	"github.com/BourgeoisBear/rasterm"
)

// Iterm renders images with the iTerm2 inline image protocol.
type Iterm struct{}

func NewIterm() *Iterm { return &Iterm{} }

func (*Iterm) Name() string { return "iterm2" }

func (*Iterm) Render(img image.Image, area Area) (string, error) {
	if img == nil {
		return "", fmt.Errorf("render nil image")
	}
	if area.X < 1 || area.Y < 1 || area.Cols < 1 || area.Rows < 1 {
		return "", fmt.Errorf("invalid terminal area: %+v", area)
	}

	var output bytes.Buffer
	// The protocol inserts an inline image at the cursor. Save and restore it so
	// Bubble Tea remains the sole owner of the text cursor.
	fmt.Fprintf(&output, "\x1b7\x1b[%d;%dH", area.Y, area.X)
	if err := rasterm.ItermWriteImageWithOptions(&output, img, rasterm.ItermImgOpts{
		Width:         strconv.Itoa(area.Cols),
		Height:        strconv.Itoa(area.Rows),
		DisplayInline: true,
	}); err != nil {
		return "", fmt.Errorf("encode iTerm2 image: %w", err)
	}
	output.WriteString("\x1b8")

	return output.String(), nil
}

// Clear erases the character cells containing the prior inline image. iTerm2
// has no image-deletion escape sequence, so erasing only this rectangle leaves
// Bubble Tea's header and footer untouched.
func (*Iterm) Clear(area Area) string {
	if area.X < 1 || area.Y < 1 || area.Cols < 1 || area.Rows < 1 {
		return ""
	}

	var output bytes.Buffer
	output.Grow(area.Rows * 16)
	output.WriteString("\x1b7")
	for row := area.Y; row < area.Y+area.Rows; row++ {
		fmt.Fprintf(&output, "\x1b[%d;%dH\x1b[%dX", row, area.X, area.Cols)
	}
	output.WriteString("\x1b8")
	return output.String()
}

var _ Renderer = (*Iterm)(nil)
