package tui

import (
	"image"
	"image/color"
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/arimatakao/comicfile/metadata"
	"github.com/arimatakao/comicread/internal/backend"
)

type fakeChapter struct {
	pages []image.Image
}

func TestFooterKeepsPageNumberWhileRendering(t *testing.T) {
	t.Parallel()

	m := New("test.cbz", fakeChapter{pages: []image.Image{
		image.NewRGBA(image.Rect(0, 0, 10, 20)),
		image.NewRGBA(image.Rect(0, 0, 10, 20)),
	}}, fakeBackend{})
	m.width = 80
	m.page = 1
	m.rendering = true

	footer := m.footer()
	if !strings.Contains(footer, "page 2/2") {
		t.Fatalf("footer %q does not contain page number", footer)
	}
	if strings.Contains(footer, "previous") || strings.Contains(footer, "next") || strings.Contains(footer, "quit") {
		t.Fatalf("footer %q contains button hints", footer)
	}
	if !strings.HasSuffix(footer, "rendering") {
		t.Fatalf("footer %q does not end with rendering status", footer)
	}
	if !strings.Contains(footer, "zoom 100%") {
		t.Fatalf("footer %q does not contain zoom", footer)
	}
}

func (f fakeChapter) TotalPages() int                     { return len(f.pages) }
func (fakeChapter) ErrPages() int                         { return 0 }
func (fakeChapter) Metadata() *metadata.Metadata          { return nil }
func (f fakeChapter) Page(index int) (image.Image, error) { return f.pages[index], nil }

type fakeBackend struct{}

func (fakeBackend) Name() string { return "fake" }
func (fakeBackend) Render(image.Image, backend.Area) (string, error) {
	return "rendered", nil
}
func (fakeBackend) Clear(backend.Area) string { return "clear" }

func TestNavigationBounds(t *testing.T) {
	t.Parallel()

	chapter := fakeChapter{pages: []image.Image{
		image.NewRGBA(image.Rect(0, 0, 10, 20)),
		image.NewRGBA(image.Rect(0, 0, 10, 20)),
	}}
	m := New("test.cbz", chapter, fakeBackend{})
	m.width, m.height = 80, 24
	m.area = imageArea(m.width, m.height, m.currentPageAspect())

	m.previousPage()
	if m.page != 0 {
		t.Fatalf("previous on first page moved to %d", m.page)
	}

	m.nextPage()
	if m.page != 1 {
		t.Fatalf("next page = %d, want 1", m.page)
	}

	m.nextPage()
	if m.page != 1 {
		t.Fatalf("next on last page moved to %d", m.page)
	}
}

func TestZoomAndScrollKeys(t *testing.T) {
	t.Parallel()

	chapter := fakeChapter{pages: []image.Image{
		image.NewRGBA(image.Rect(0, 0, 10, 20)),
		image.NewRGBA(image.Rect(0, 0, 10, 20)),
	}}
	m := New("test.cbz", chapter, fakeBackend{})
	m.width, m.height = 80, 24
	m.area = imageArea(m.width, m.height, m.currentPageAspect())

	updated, _ := m.Update(tea.KeyPressMsg{Code: '+'})
	m = updated.(Model)
	if m.zoom != 125 {
		t.Fatalf("zoom after + = %d, want 125", m.zoom)
	}

	updated, _ = m.Update(tea.KeyPressMsg{Code: tea.KeyDown})
	m = updated.(Model)
	if m.scroll <= 0 {
		t.Fatalf("scroll after down = %v, want positive", m.scroll)
	}

	updated, _ = m.Update(tea.KeyPressMsg{Code: tea.KeyRight})
	m = updated.(Model)
	if m.page != 1 {
		t.Fatalf("page after right = %d, want 1", m.page)
	}
	if m.scroll != 0 {
		t.Fatalf("scroll after changing page = %v, want 0", m.scroll)
	}
	if m.zoom != 125 {
		t.Fatalf("zoom after changing page = %d, want 125", m.zoom)
	}

	updated, _ = m.Update(tea.KeyPressMsg{Code: '-'})
	m = updated.(Model)
	if m.zoom != 100 {
		t.Fatalf("zoom after - = %d, want 100", m.zoom)
	}
}

func TestZoomedImageScrollsVertically(t *testing.T) {
	t.Parallel()

	img := image.NewRGBA(image.Rect(0, 0, 4, 10))
	for y := 0; y < 10; y++ {
		for x := 0; x < 4; x++ {
			img.SetRGBA(x, y, color.RGBA{R: uint8(y), A: 255})
		}
	}

	top := zoomedImage(img, 200, 0)
	if got := top.At(0, 0); got != (color.RGBA{A: 255}) {
		t.Fatalf("top crop starts with %v, want row 0", got)
	}
	if top.Bounds().Dx() != 4 || top.Bounds().Dy() != 5 {
		t.Fatalf("top crop bounds = %v, want 4x5", top.Bounds())
	}

	bottom := zoomedImage(img, 200, 1)
	if got := bottom.At(0, 0); got != (color.RGBA{R: 5, A: 255}) {
		t.Fatalf("bottom crop starts with %v, want row 5", got)
	}
}

func TestZoomedAreaKeepsImageProportions(t *testing.T) {
	t.Parallel()

	m := New("test.cbz", fakeChapter{pages: []image.Image{
		image.NewRGBA(image.Rect(0, 0, 10, 20)),
	}}, fakeBackend{})
	m.width, m.height = 80, 24
	m.zoom = 200

	area := m.zoomedArea()
	if area.Cols != 44 || area.Rows != 22 {
		t.Fatalf("zoomed area = %+v, want 44 columns by 22 rows", area)
	}
}
