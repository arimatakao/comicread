package web

import (
	"encoding/json"
	"errors"
	"image/png"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/arimatakao/comicfile"
)

// maxUploadSize caps how large an uploaded chapter may be. Chapters are read
// fully into memory (comicread never writes temporary files to open a book),
// so this bound protects against a malformed or oversized upload exhausting
// memory.
const maxUploadSize = 1 << 30 // 1 GiB

func newMux(st *store) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handleIndex)
	mux.HandleFunc("GET /pico.min.css", handleAsset(picoCSS, "text/css; charset=utf-8"))
	mux.HandleFunc("GET /app.js", handleAsset(appJS, "text/javascript; charset=utf-8"))
	mux.HandleFunc("POST /api/open", handleOpen(st))
	mux.HandleFunc("GET /api/books/{token}", handleBookInfo(st))
	mux.HandleFunc("DELETE /api/books/{token}", handleCloseBook(st))
	mux.HandleFunc("GET /api/books/{token}/pages/{index}", handlePage(st))
	return mux
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write(indexHTML)
}

func handleAsset(data []byte, contentType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		w.Header().Set("Cache-Control", "no-cache")
		_, _ = w.Write(data)
	}
}

// handleOpen accepts a multipart/form-data upload (field "file") holding a
// whole CBZ, PDF, or EPUB chapter, and opens it with comicfile.OpenBytes.
//
// It reads the upload through multipart.Reader directly rather than
// r.ParseMultipartForm, because ParseMultipartForm spools large uploads to
// temporary files on disk; comicread's web reader must never touch disk to
// open a book.
func handleOpen(st *store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)

		multipartReader, err := r.MultipartReader()
		if err != nil {
			writeError(w, http.StatusBadRequest, "expected a multipart/form-data upload")
			return
		}

		var (
			data     []byte
			filename string
		)
		for {
			part, err := multipartReader.NextPart()
			if errors.Is(err, io.EOF) {
				break
			}
			if err != nil {
				writeError(w, http.StatusBadRequest, "read upload: "+err.Error())
				return
			}
			if part.FormName() == "file" && filename == "" {
				filename = part.FileName()
				data, err = io.ReadAll(part)
				part.Close()
				if err != nil {
					writeError(w, http.StatusRequestEntityTooLarge, "upload too large or interrupted")
					return
				}
				continue
			}
			part.Close()
		}

		if filename == "" || len(data) == 0 {
			writeError(w, http.StatusBadRequest, "no file uploaded")
			return
		}

		extension := strings.TrimPrefix(strings.ToLower(filepath.Ext(filename)), ".")
		if extension != comicfile.CBZ_EXT && extension != comicfile.PDF_EXT && extension != comicfile.EPUB_EXT {
			writeError(w, http.StatusBadRequest, "unsupported file type: only CBZ, PDF, and EPUB are supported")
			return
		}

		title := strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))
		info, err := st.open(extension, data, title)
		if err != nil {
			writeError(w, http.StatusUnprocessableEntity, "open chapter: "+err.Error())
			return
		}

		writeJSON(w, http.StatusOK, info)
	}
}

func handleBookInfo(st *store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info, ok := st.info(r.PathValue("token"))
		if !ok {
			writeError(w, http.StatusNotFound, "book session not found")
			return
		}
		writeJSON(w, http.StatusOK, info)
	}
}

// handleCloseBook releases the book identified by token, freeing the
// in-memory container. The browser calls this once it has fetched every
// page and cached them itself, trading away the ability to resume this
// token later (e.g. after a reload) for not holding the chapter in server
// memory for longer than needed.
func handleCloseBook(st *store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		st.closeIfActive(r.PathValue("token"))
		w.WriteHeader(http.StatusNoContent)
	}
}

// handlePage renders one page as a PNG image on demand. Pages are streamed
// individually; the browser never receives more than the page it asked for.
func handlePage(st *store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		index, err := strconv.Atoi(r.PathValue("index"))
		if err != nil || index < 0 {
			writeError(w, http.StatusBadRequest, "invalid page index")
			return
		}

		img, err := st.page(r.PathValue("token"), index)
		if err != nil {
			if errors.Is(err, errStaleToken) {
				writeError(w, http.StatusNotFound, "book session not found")
				return
			}
			writeError(w, http.StatusNotFound, "page not found: "+err.Error())
			return
		}

		// The same token+index pair always maps to the same page, so
		// browsers and preloading may cache it indefinitely.
		w.Header().Set("Content-Type", "image/png")
		w.Header().Set("Cache-Control", "private, max-age=31536000, immutable")
		_ = png.Encode(w, img)
	}
}

type errorResponse struct {
	Error string `json:"error"`
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, errorResponse{Error: message})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
