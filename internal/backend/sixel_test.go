package backend

import (
	"image"
	"image/color"
	"strings"
	"testing"
)

func TestSixelRender(t *testing.T) {
	t.Parallel()

	img := image.NewNRGBA(image.Rect(0, 0, 2, 3))
	img.SetNRGBA(0, 0, color.NRGBA{R: 255, A: 255})
	output, err := NewSixel().Render(img, Area{X: 4, Y: 2, Cols: 2, Rows: 1})
	if err != nil {
		t.Fatalf("Render() error = %v", err)
	}

	for _, want := range []string{
		"\x1b[2;4H",
		"\x1bP0;1q\"1;1;16;16",
		";2;100;0;0",
		"\x1b\\",
	} {
		if !strings.Contains(output, want) {
			t.Errorf("Render() output does not contain %q", want)
		}
	}
}

func TestSixelRenderRejectsInvalidArea(t *testing.T) {
	t.Parallel()

	_, err := NewSixel().Render(image.NewNRGBA(image.Rect(0, 0, 1, 1)), Area{})
	if err == nil {
		t.Fatal("Render() error = nil, want invalid area error")
	}
}

func TestSixelClearReplacesOnlyImageArea(t *testing.T) {
	t.Parallel()

	clear := NewSixel().Clear(Area{X: 3, Y: 2, Cols: 10, Rows: 2})
	for _, want := range []string{"\x1b[2;3H", "\x1bP0;0;0q\"1;1;80;32", "!80?", "\x1b\\"} {
		if !strings.Contains(clear, want) {
			t.Errorf("Clear() output does not contain %q", want)
		}
	}
	if strings.Contains(clear, "\x1b[2K") {
		t.Errorf("Clear() = %q, uses text-only erase sequence", clear)
	}
}
