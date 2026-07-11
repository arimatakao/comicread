package backend

import (
	"bytes"
	"fmt"
	"image"
	"os"

	"github.com/BourgeoisBear/rasterm"
)

// Kitty renders images with the Kitty graphics protocol.
type Kitty struct {
	imageID     uint32
	placementID uint32
}

func NewKitty() *Kitty {
	imageID := uint32(os.Getpid())
	if imageID == 0 {
		imageID = 1
	}
	return &Kitty{imageID: imageID, placementID: 1}
}

func (*Kitty) Name() string { return "kitty" }

func (k *Kitty) Render(img image.Image, area Area) (string, error) {
	if img == nil {
		return "", fmt.Errorf("render nil image")
	}
	if area.X < 1 || area.Y < 1 || area.Cols < 1 || area.Rows < 1 {
		return "", fmt.Errorf("invalid terminal area: %+v", area)
	}

	var output bytes.Buffer
	// rasterm places images at the current cursor position. Save and restore it
	// so Bubble Tea keeps ownership of the text cursor.
	fmt.Fprintf(&output, "\x1b7\x1b[%d;%dH", area.Y, area.X)
	if err := rasterm.KittyWriteImage(&output, img, rasterm.KittyImgOpts{
		DstCols:     uint32(area.Cols),
		DstRows:     uint32(area.Rows),
		ZIndex:      1,
		ImageId:     k.imageID,
		PlacementId: k.placementID,
	}); err != nil {
		return "", fmt.Errorf("encode Kitty image: %w", err)
	}
	output.WriteString("\x1b8")

	return output.String(), nil
}

func (k *Kitty) Clear(Area) string {
	return rasterm.KittyImgOpts{
		ImageId:     k.imageID,
		PlacementId: k.placementID,
	}.ToHeader("a=d", "d=I", "q=2") + "\x1b\\"
}

var _ Renderer = (*Kitty)(nil)
