package cli

import (
	"archive/zip"
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	"github.com/arimatakao/comicfile"
	"github.com/arimatakao/comicfile/metadata"
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

func TestOpenImageDirectory(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	for index := range 2 {
		path := filepath.Join(dir, fmt.Sprintf("%03d.png", index+1))
		file, err := os.Create(path)
		if err != nil {
			t.Fatal(err)
		}
		img := image.NewRGBA(image.Rect(0, 0, 2+index, 3+index))
		if err := png.Encode(file, img); err != nil {
			t.Fatal(err)
		}
		if err := file.Close(); err != nil {
			t.Fatal(err)
		}
	}

	if err := validateInput(dir); err != nil {
		t.Fatalf("validateInput() error = %v", err)
	}
	chapter, err := openChapter(dir)
	if err != nil {
		t.Fatalf("openChapter() error = %v", err)
	}
	if chapter.TotalPages() != 2 {
		t.Fatalf("TotalPages() = %d, want 2", chapter.TotalPages())
	}
}

func TestOpenChapterSupportedFormats(t *testing.T) {
	for _, format := range []string{comicfile.CBZ_EXT, comicfile.PDF_EXT, comicfile.EPUB_EXT} {
		t.Run(format, func(t *testing.T) {
			path := createChapter(t, format)
			chapter, err := openChapter(path)
			if err != nil {
				t.Fatalf("openChapter() error = %v", err)
			}
			if chapter.TotalPages() != 1 {
				t.Fatalf("TotalPages() = %d, want 1", chapter.TotalPages())
			}
		})
	}
}

func TestIsSupportedFile(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		path string
		want bool
	}{
		{"chapter.cbz", true},
		{"chapter.PDF", true},
		{"chapter.epub", true},
		{"chapter.zip", false},
		{"chapter", false},
	} {
		if got := isSupportedFile(test.path); got != test.want {
			t.Errorf("isSupportedFile(%q) = %v, want %v", test.path, got, test.want)
		}
	}
}

func TestParseArgs(t *testing.T) {
	t.Parallel()

	graphics, path, _, err := parseArgs([]string{"--graphics", "sixel", "chapter.cbz"})
	if err != nil {
		t.Fatalf("parseArgs() error = %v", err)
	}
	if graphics != "sixel" || path != "chapter.cbz" {
		t.Fatalf("parseArgs() = (%q, %q), want (sixel, chapter.cbz)", graphics, path)
	}
}

func TestValidateInputRejectsUnsupportedFile(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "chapter.zip")
	if err := os.WriteFile(path, nil, 0o600); err != nil {
		t.Fatal(err)
	}
	if err := validateInput(path); err == nil {
		t.Fatal("validateInput() error = nil, want unsupported file error")
	}
}

func createChapter(t *testing.T, format string) string {
	t.Helper()

	container, err := comicfile.NewContainer(format)
	if err != nil {
		t.Fatal(err)
	}
	page := image.NewRGBA(image.Rect(0, 0, 2, 3))
	page.Set(0, 0, color.RGBA{R: 255, A: 255})
	var encoded bytes.Buffer
	if err := png.Encode(&encoded, page); err != nil {
		t.Fatal(err)
	}
	if err := container.AddPage("png", encoded.Bytes()); err != nil {
		t.Fatal(err)
	}

	dir := t.TempDir()
	if err := container.WriteOnDiskAndClose(dir, "chapter", metadata.Metadata{}, ""); err != nil {
		t.Fatal(err)
	}
	return filepath.Join(dir, "chapter."+format)
}
