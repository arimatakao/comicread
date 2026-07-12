// Command demomanga creates small sample chapters for testing comicread.
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/arimatakao/comicfile"
	"github.com/arimatakao/comicfile/metadata"
)

const (
	imageDir   = "demomanga"
	outputDir  = "demomanga/generated"
	outputName = "demo-manga"
)

func main() {
	pages, err := sourcePages()
	if err != nil {
		log.Fatal(err)
	}

	for _, format := range []string{
		comicfile.PDF_EXT,
		comicfile.EPUB_EXT,
		comicfile.CBZ_EXT,
		comicfile.DIR_EXT,
	} {
		if err := writeChapter(format, pages); err != nil {
			log.Fatalf("write %s: %v", format, err)
		}
	}

	fmt.Printf("Created sample chapters in %s:\n", outputDir)
	fmt.Printf("  %s\n", filepath.Join(outputDir, outputName+".pdf"))
	fmt.Printf("  %s\n", filepath.Join(outputDir, outputName+".epub"))
	fmt.Printf("  %s\n", filepath.Join(outputDir, outputName+".cbz"))
	fmt.Printf("  %s\n", filepath.Join(outputDir, outputName))
}

type page struct {
	extension string
	contents  []byte
}

func writeChapter(format string, pages []page) error {
	container, err := comicfile.NewContainer(format)
	if err != nil {
		return err
	}

	for _, page := range pages {
		if err := container.AddPage(page.extension, page.contents); err != nil {
			return err
		}
	}

	return container.WriteOnDiskAndClose(outputDir, outputName, sampleMetadata(len(pages)), "")
}

func sampleMetadata(pageCount int) metadata.Metadata {
	return metadata.Metadata{
		CBI: metadata.ComicBookMetadata{
			AppID: "demomanga",
			ComicBookInfoData: metadata.ComicBookInfo{
				Series:   "Demo Manga",
				Title:    "Demo Manga",
				Issue:    "1",
				Volume:   "1",
				Language: "en",
				Credits:  []metadata.Credit{{Person: "comicread", Role: "Writer"}},
				Tags:     []string{"demo", "test"},
			},
		},
		CI: metadata.ComicInfoMetadata{
			Title:       "Demo Manga",
			Number:      "1",
			Volume:      "1",
			Year:        2026,
			Writer:      "comicread",
			Penciller:   "comicread",
			PageCount:   pageCount,
			LanguageISO: "en",
			Manga:       "Yes",
			Summary:     "Generated sample pages for testing comicread.",
		},
		P: metadata.PlainMetadata{
			Authors: "comicread",
			Artists: "comicread",
			Tags:    "demo, test",
		},
	}
}

func sourcePages() ([]page, error) {
	entries, err := os.ReadDir(imageDir)
	if err != nil {
		return nil, err
	}

	pages := make([]page, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		extension := strings.TrimPrefix(strings.ToLower(filepath.Ext(entry.Name())), ".")
		if extension != "jpg" && extension != "jpeg" && extension != "png" {
			continue
		}

		contents, err := os.ReadFile(filepath.Join(imageDir, entry.Name()))
		if err != nil {
			return nil, err
		}
		pages = append(pages, page{extension: extension, contents: contents})
	}

	if len(pages) == 0 {
		return nil, fmt.Errorf("no JPG or PNG files in %s", imageDir)
	}

	return pages, nil
}
