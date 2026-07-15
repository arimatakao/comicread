// Package journal persists reader progress and bookmarks next to a chapter.
package journal

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/arimatakao/comicread/internal/journal/v1"
)

const (
	directoryName = ".comicread"
	fileName      = "journal.json"
)

type versionHeader struct {
	Version int `json:"version"`
}

// Journal stores the state of one input chapter in its nearest local journal.
type Journal struct {
	path     string
	document string
	data     v1.Journal
}

// Open loads the journal for inputPath. It does not create files until state
// changes need to be saved.
func Open(inputPath string) (*Journal, error) {
	documentPath, journalDir, err := paths(inputPath)
	if err != nil {
		return nil, err
	}

	j := &Journal{
		path:     filepath.Join(journalDir, directoryName, fileName),
		document: documentPath,
		data:     v1.New(),
	}
	if err := j.load(); err != nil {
		return nil, err
	}
	return j, nil
}

// Clear removes the local journal directory for inputPath.
func Clear(inputPath string) error {
	_, journalDir, err := paths(inputPath)
	if err != nil {
		return err
	}
	if err := os.RemoveAll(filepath.Join(journalDir, directoryName)); err != nil {
		return fmt.Errorf("remove journal: %w", err)
	}
	return nil
}

// LastOpenedPage returns the saved one-based page number, or zero when no
// position has been recorded yet.
func (j *Journal) LastOpenedPage() int {
	return j.data.LastOpenedPage(j.document)
}

// Bookmarks returns sorted one-based bookmarked page numbers.
func (j *Journal) Bookmarks() []int {
	return j.data.Bookmarks(j.document)
}

// SetLastOpenedPage records a one-based page number.
func (j *Journal) SetLastOpenedPage(page int) error {
	if !j.data.SetLastOpenedPage(j.document, page) {
		return nil
	}
	return j.save()
}

// ToggleBookmark adds page when it is absent and removes it otherwise. It
// returns whether the page is bookmarked after the operation.
func (j *Journal) ToggleBookmark(page int) (bool, error) {
	bookmarked, err := j.data.ToggleBookmark(j.document, page)
	if err != nil {
		return false, err
	}
	if err := j.save(); err != nil {
		return !bookmarked, err
	}
	return bookmarked, nil
}

func (j *Journal) load() error {
	contents, err := os.ReadFile(j.path)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("read %q: %w", j.path, err)
	}

	var header versionHeader
	if err := json.Unmarshal(contents, &header); err != nil {
		return fmt.Errorf("read journal version: %w", err)
	}
	switch header.Version {
	case v1.Version:
		data, err := v1.Decode(contents)
		if err != nil {
			return fmt.Errorf("parse %q: %w", j.path, err)
		}
		j.data = data
		return nil
	default:
		return fmt.Errorf("unsupported journal version %d", header.Version)
	}
}

func (j *Journal) save() error {
	if err := os.MkdirAll(filepath.Dir(j.path), 0o755); err != nil {
		return fmt.Errorf("create journal directory: %w", err)
	}
	contents, err := j.data.Encode()
	if err != nil {
		return fmt.Errorf("encode journal: %w", err)
	}

	temporary, err := os.CreateTemp(filepath.Dir(j.path), ".journal-*.tmp")
	if err != nil {
		return fmt.Errorf("create temporary journal: %w", err)
	}
	temporaryName := temporary.Name()
	defer os.Remove(temporaryName)
	if _, err := temporary.Write(contents); err != nil {
		temporary.Close()
		return fmt.Errorf("write temporary journal: %w", err)
	}
	if err := temporary.Chmod(0o600); err != nil {
		temporary.Close()
		return fmt.Errorf("set journal permissions: %w", err)
	}
	if err := temporary.Close(); err != nil {
		return fmt.Errorf("close temporary journal: %w", err)
	}
	if err := os.Rename(temporaryName, j.path); err != nil {
		return fmt.Errorf("replace journal: %w", err)
	}
	return nil
}

func paths(inputPath string) (documentPath, journalDir string, err error) {
	documentPath, err = filepath.Abs(inputPath)
	if err != nil {
		return "", "", fmt.Errorf("resolve input path: %w", err)
	}
	documentPath = filepath.Clean(documentPath)

	info, err := os.Stat(documentPath)
	if err != nil {
		return "", "", fmt.Errorf("inspect input path: %w", err)
	}
	journalDir = filepath.Dir(documentPath)
	if info.IsDir() {
		journalDir = documentPath
	}
	return documentPath, journalDir, nil
}
