package tui

import (
	"image"
	"strings"
	"testing"

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
