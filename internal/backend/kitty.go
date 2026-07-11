package backend

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"os"
	"strings"
)

const kittyChunkSize = 4096

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

	var pngData bytes.Buffer
	if err := png.Encode(&pngData, img); err != nil {
		return "", fmt.Errorf("encode PNG: %w", err)
	}

	encoded := base64.StdEncoding.EncodeToString(pngData.Bytes())
	var output strings.Builder
	output.Grow(len(encoded) + len(encoded)/kittyChunkSize*32 + 128)

	// Kitty places an image at the cursor position when the final chunk arrives.
	// Move there first and ask Kitty not to move the cursor after placement.
	fmt.Fprintf(&output, "\x1b[%d;%dH", area.Y, area.X)

	for offset := 0; offset < len(encoded); offset += kittyChunkSize {
		end := min(offset+kittyChunkSize, len(encoded))
		more := 0
		if end < len(encoded) {
			more = 1
		}

		if offset == 0 {
			fmt.Fprintf(
				&output,
				"\x1b_Ga=T,f=100,t=d,i=%d,p=%d,c=%d,r=%d,z=1,C=1,q=2,m=%d;%s\x1b\\",
				k.imageID,
				k.placementID,
				area.Cols,
				area.Rows,
				more,
				encoded[offset:end],
			)
			continue
		}

		fmt.Fprintf(&output, "\x1b_Gm=%d,q=2;%s\x1b\\", more, encoded[offset:end])
	}

	return output.String(), nil
}

func (k *Kitty) Clear() string {
	return fmt.Sprintf(
		"\x1b_Ga=d,d=I,i=%d,p=%d,q=2\x1b\\",
		k.imageID,
		k.placementID,
	)
}

var _ Renderer = (*Kitty)(nil)
