package backend

import (
	"image"
	"strings"
	"testing"
)

func TestKittyRender(t *testing.T) {
	t.Parallel()

	img := image.NewRGBA(image.Rect(0, 0, 2, 3))
	renderer := &Kitty{imageID: 7, placementID: 3}
	output, err := renderer.Render(img, Area{X: 4, Y: 2, Cols: 20, Rows: 10})
	if err != nil {
		t.Fatalf("Render() error = %v", err)
	}

	for _, want := range []string{
		"\x1b[2;4H",
		"\x1b_Ga=T,f=100,t=d,i=7,p=3,c=20,r=10",
		"m=0;",
		"\x1b\\",
	} {
		if !strings.Contains(output, want) {
			t.Errorf("Render() output does not contain %q", want)
		}
	}
}

func TestKittyRenderRejectsInvalidArea(t *testing.T) {
	t.Parallel()

	_, err := (&Kitty{imageID: 7, placementID: 3}).Render(image.NewRGBA(image.Rect(0, 0, 1, 1)), Area{})
	if err == nil {
		t.Fatal("Render() error = nil, want invalid area error")
	}
}

func TestKittyClearTargetsOwnPlacement(t *testing.T) {
	t.Parallel()

	clear := (&Kitty{imageID: 7, placementID: 3}).Clear(Area{})
	if !strings.Contains(clear, "a=d,d=I,i=7,p=3") {
		t.Fatalf("Clear() = %q", clear)
	}
}
