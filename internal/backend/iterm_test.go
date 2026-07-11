package backend

import (
	"image"
	"strings"
	"testing"
)

func TestItermRender(t *testing.T) {
	t.Parallel()

	output, err := NewIterm().Render(image.NewRGBA(image.Rect(0, 0, 2, 3)), Area{X: 4, Y: 2, Cols: 20, Rows: 10})
	if err != nil {
		t.Fatalf("Render() error = %v", err)
	}

	for _, want := range []string{
		"\x1b7",
		"\x1b[2;4H",
		"\x1b]1337;File=width=20;height=10;size=",
		"inline=1:",
		"\a",
		"\x1b8",
	} {
		if !strings.Contains(output, want) {
			t.Errorf("Render() output does not contain %q", want)
		}
	}
}

func TestItermRenderRejectsInvalidArea(t *testing.T) {
	t.Parallel()

	_, err := NewIterm().Render(image.NewRGBA(image.Rect(0, 0, 1, 1)), Area{})
	if err == nil {
		t.Fatal("Render() error = nil, want invalid area error")
	}
}

func TestItermClearErasesOnlyImageArea(t *testing.T) {
	t.Parallel()

	clear := NewIterm().Clear(Area{X: 3, Y: 2, Cols: 10, Rows: 2})
	for _, want := range []string{"\x1b7", "\x1b[2;3H\x1b[10X", "\x1b[3;3H\x1b[10X", "\x1b8"} {
		if !strings.Contains(clear, want) {
			t.Errorf("Clear() output does not contain %q", want)
		}
	}
}
