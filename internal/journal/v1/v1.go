// Package v1 defines version 1 of the comicread journal file format.
package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"slices"
)

// Version is the journal format version implemented by this package.
const Version = 1

// Journal is the on-disk journal format.
type Journal struct {
	Version   int                 `json:"version"`
	Documents map[string]Document `json:"documents"`
}

// Document is the persisted state for one chapter.
type Document struct {
	LastOpenedPage int   `json:"last_opened_page"`
	Bookmarks      []int `json:"bookmarks"`
}

// New creates an empty version 1 journal.
func New() Journal {
	return Journal{
		Version:   Version,
		Documents: make(map[string]Document),
	}
}

// Decode parses and normalizes a version 1 journal.
func Decode(contents []byte) (Journal, error) {
	var journal Journal
	if err := json.Unmarshal(contents, &journal); err != nil {
		return Journal{}, err
	}
	if journal.Version != Version {
		return Journal{}, fmt.Errorf("unsupported journal version %d", journal.Version)
	}
	if journal.Documents == nil {
		journal.Documents = make(map[string]Document)
	}
	for path, document := range journal.Documents {
		document.Bookmarks = uniquePositive(document.Bookmarks)
		journal.Documents[path] = document
	}
	return journal, nil
}

// Encode returns the formatted version 1 journal contents.
func (j Journal) Encode() ([]byte, error) {
	contents, err := json.MarshalIndent(j, "", "  ")
	if err != nil {
		return nil, err
	}
	return append(contents, '\n'), nil
}

// LastOpenedPage returns the one-based saved page for path, or zero when no
// position has been recorded.
func (j Journal) LastOpenedPage(path string) int {
	return j.Documents[path].LastOpenedPage
}

// Bookmarks returns sorted one-based bookmarked page numbers for path.
func (j Journal) Bookmarks(path string) []int {
	return slices.Clone(j.Documents[path].Bookmarks)
}

// SetLastOpenedPage records a one-based page number for path.
func (j *Journal) SetLastOpenedPage(path string, page int) bool {
	if page < 1 {
		return false
	}
	document := j.Documents[path]
	if document.LastOpenedPage == page {
		return false
	}
	document.LastOpenedPage = page
	j.Documents[path] = document
	return true
}

// ToggleBookmark adds page when it is absent and removes it otherwise. It
// returns whether the page is bookmarked after the operation.
func (j *Journal) ToggleBookmark(path string, page int) (bool, error) {
	if page < 1 {
		return false, errors.New("bookmark page must be positive")
	}
	document := j.Documents[path]
	index, found := slices.BinarySearch(document.Bookmarks, page)
	if found {
		document.Bookmarks = append(document.Bookmarks[:index], document.Bookmarks[index+1:]...)
	} else {
		document.Bookmarks = append(document.Bookmarks, 0)
		copy(document.Bookmarks[index+1:], document.Bookmarks[index:])
		document.Bookmarks[index] = page
	}
	j.Documents[path] = document
	return !found, nil
}

func uniquePositive(pages []int) []int {
	pages = slices.Clone(pages)
	slices.Sort(pages)
	result := pages[:0]
	for _, page := range pages {
		if page > 0 && (len(result) == 0 || result[len(result)-1] != page) {
			result = append(result, page)
		}
	}
	return result
}
