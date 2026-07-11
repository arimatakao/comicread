package cli

import (
	"archive/zip"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"testing"
)

func TestOpenChapter(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "chapter.cbz")
	file, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	archive := zip.NewWriter(file)
	for index := range 2 {
		entry, err := archive.Create(fmt.Sprintf("%03d.png", index+1))
		if err != nil {
			t.Fatal(err)
		}
		img := image.NewRGBA(image.Rect(0, 0, 2+index, 3+index))
		img.Set(0, 0, color.RGBA{R: 255, A: 255})
		if err := png.Encode(entry, img); err != nil {
			t.Fatal(err)
		}
	}
	if err := archive.Close(); err != nil {
		t.Fatal(err)
	}
	if err := file.Close(); err != nil {
		t.Fatal(err)
	}

	chapter, err := openChapter(path)
	if err != nil {
		t.Fatalf("openChapter() error = %v", err)
	}
	if chapter.TotalPages() != 2 {
		t.Fatalf("TotalPages() = %d, want 2", chapter.TotalPages())
	}
	if _, err := chapter.Page(1); err != nil {
		t.Fatalf("Page(1) error = %v", err)
	}
}

func TestOpenChapterRejectsEmptyCBZ(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "empty.cbz")
	file, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	archive := zip.NewWriter(file)
	if err := archive.Close(); err != nil {
		t.Fatal(err)
	}
	if err := file.Close(); err != nil {
		t.Fatal(err)
	}

	if _, err := openChapter(path); err == nil {
		t.Fatal("openChapter() error = nil, want empty chapter error")
	}
}
