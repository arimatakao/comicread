package backend

import (
	"bytes"
	"fmt"
	"image"
	"os"
	"sync"

	"github.com/BourgeoisBear/rasterm"
)

// Kitty renders images with the Kitty graphics protocol.
type Kitty struct {
	mu            sync.Mutex
	imageID       uint32
	placementID   uint32
	nextImageID   uint32
	currentImages []uint32
	pendingImages []uint32
}

func NewKitty() *Kitty {
	imageID := uint32(os.Getpid())
	if imageID == 0 {
		imageID = 1
	}
	return &Kitty{imageID: imageID, placementID: 1, nextImageID: imageID}
}

func (*Kitty) Name() string { return "kitty" }

func (k *Kitty) Render(img image.Image, area Area) (string, error) {
	if img == nil {
		return "", fmt.Errorf("render nil image")
	}
	if area.X < 1 || area.Y < 1 || area.Cols < 1 || area.Rows < 1 {
		return "", fmt.Errorf("invalid terminal area: %+v", area)
	}
	imageID := k.reserveImageID()

	var output bytes.Buffer
	// rasterm places images at the current cursor position. Save and restore it
	// so Bubble Tea keeps ownership of the text cursor.
	fmt.Fprintf(&output, "\x1b7\x1b[%d;%dH", area.Y, area.X)
	if err := rasterm.KittyWriteImage(&output, img, rasterm.KittyImgOpts{
		DstCols:     uint32(area.Cols),
		DstRows:     uint32(area.Rows),
		ZIndex:      1,
		ImageId:     imageID,
		PlacementId: k.placement(),
	}); err != nil {
		return "", fmt.Errorf("encode Kitty image: %w", err)
	}
	output.WriteString("\x1b8")

	return output.String(), nil
}

func (k *Kitty) reserveImageID() uint32 {
	k.mu.Lock()
	defer k.mu.Unlock()

	if k.nextImageID == 0 {
		k.nextImageID = k.imageID
		if k.nextImageID == 0 {
			k.nextImageID = 1
		}
	}
	imageID := k.nextImageID
	k.nextImageID++
	if k.nextImageID == 0 {
		k.nextImageID = 1
	}
	k.pendingImages = append(k.pendingImages, imageID)
	return imageID
}

func (k *Kitty) placement() uint32 {
	if k.placementID == 0 {
		return 1
	}
	return k.placementID
}

func (k *Kitty) Clear(Area) string {
	k.mu.Lock()
	defer k.mu.Unlock()

	var output bytes.Buffer
	if len(k.currentImages) == 0 && k.imageID != 0 {
		output.WriteString(rasterm.KittyImgOpts{
			ImageId:     k.imageID,
			PlacementId: k.placement(),
		}.ToHeader("a=d", "d=I", "q=2"))
		output.WriteString("\x1b\\")
	}
	for _, imageID := range k.currentImages {
		output.WriteString(rasterm.KittyImgOpts{ImageId: imageID}.ToHeader("a=d", "d=I", "q=2"))
		output.WriteString("\x1b\\")
	}
	k.currentImages = k.pendingImages
	k.pendingImages = nil
	return output.String()
}

var _ Renderer = (*Kitty)(nil)
