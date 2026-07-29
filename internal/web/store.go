package web

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"image"
	"sync"

	"github.com/arimatakao/comicfile"
)

// errStaleToken is returned when a request references a book token that is
// unknown or no longer the active one.
var errStaleToken = errors.New("book session is no longer active")

// errNoPages is returned when an opened container has no readable pages.
var errNoPages = errors.New("chapter contains no readable image pages")

// bookInfo describes the currently opened chapter, as returned to the
// browser.
type bookInfo struct {
	Token      string `json:"token"`
	Title      string `json:"title"`
	TotalPages int    `json:"totalPages"`
}

// store holds at most one open comic chapter at a time, mirroring
// comicread's single-reading-session model: opening a new book closes and
// replaces whatever was open before. token changes on every open, so
// requests still referencing a replaced or closed book are rejected instead
// of silently serving stale pages.
type store struct {
	mu    sync.RWMutex
	token string
	book  comicfile.ContainerReader
	title string
	total int
}

func newStore() *store {
	return &store{}
}

// open decodes data as a chapter of the given comicfile extension
// (comicfile.CBZ_EXT, PDF_EXT, or EPUB_EXT), entirely in memory, and makes it
// the active book. Any previously active book is closed.
func (s *store) open(extension string, data []byte, title string) (bookInfo, error) {
	reader, err := comicfile.OpenBytes(extension, data)
	if err != nil {
		return bookInfo{}, err
	}
	if reader.TotalPages() == 0 {
		_ = reader.Close()
		return bookInfo{}, errNoPages
	}

	token, err := newToken()
	if err != nil {
		_ = reader.Close()
		return bookInfo{}, err
	}

	s.mu.Lock()
	previous := s.book
	s.token = token
	s.book = reader
	s.title = title
	s.total = reader.TotalPages()
	s.mu.Unlock()

	if previous != nil {
		_ = previous.Close()
	}

	return bookInfo{Token: token, Title: title, TotalPages: reader.TotalPages()}, nil
}

// info returns the active book's metadata when token matches the currently
// active book.
func (s *store) info(token string) (bookInfo, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if token == "" || token != s.token || s.book == nil {
		return bookInfo{}, false
	}
	return bookInfo{Token: s.token, Title: s.title, TotalPages: s.total}, true
}

// page returns the decoded image for the zero-based page index of the book
// identified by token.
func (s *store) page(token string, index int) (image.Image, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if token == "" || token != s.token || s.book == nil {
		return nil, errStaleToken
	}
	return s.book.Page(index)
}

// close releases the active book, if any. Called once, on server shutdown.
func (s *store) close() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.book != nil {
		_ = s.book.Close()
		s.book = nil
		s.token = ""
	}
}

// closeIfActive releases the book identified by token, but only if it is
// still the active one. It is a no-op when token has already been replaced
// by a newer open() or already closed, so a late or duplicate close request
// can never release a book it no longer refers to.
func (s *store) closeIfActive(token string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if token == "" || token != s.token || s.book == nil {
		return
	}
	_ = s.book.Close()
	s.book = nil
	s.token = ""
}

func newToken() (string, error) {
	raw := make([]byte, 16)
	if _, err := rand.Read(raw); err != nil {
		return "", err
	}
	return hex.EncodeToString(raw), nil
}
