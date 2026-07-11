package tui

import (
	"image"
	"testing"

	"github.com/arimatakao/comicfile/metadata"
	"github.com/arimatakao/comicread/internal/backend"
)

type fakeChapter struct {
	pages []image.Image
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
func (fakeBackend) Clear() string { return "clear" }

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
