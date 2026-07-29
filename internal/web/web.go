// Package web implements comicread's local browser reader, started with the
// --web flag: a minimal HTTP server, bound only to 127.0.0.1, that serves a
// single-page picker/reader UI. There is no library — the browser picks a
// file with the native file input, the server opens it with comicfile
// (entirely in memory, no temporary files), and serves it back one page at a
// time as PNG images.
package web

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/arimatakao/comicread/internal/i18n"
)

// shutdownTimeout bounds how long Serve waits for in-flight requests (such
// as a page still being encoded) to finish once ctx is cancelled.
const shutdownTimeout = 5 * time.Second

// Serve starts the web reader and blocks until ctx is cancelled
// (SIGINT/SIGTERM), then shuts the server down gracefully.
//
// It listens on 127.0.0.1 — never on a non-loopback address — on port,
// or on an OS-assigned free port when port is 0. It prints the local URL,
// and best-effort opens it in the system default browser.
func Serve(ctx context.Context, port int) error {
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		return fmt.Errorf(i18n.T(i18n.WebErrListen), err)
	}

	url := "http://" + listener.Addr().String() + "/"
	fmt.Println(i18n.T(i18n.WebServerStarted, url))
	openBrowser(url)

	st := newStore()
	defer st.close()

	server := &http.Server{Handler: newMux(st)}

	serveErr := make(chan error, 1)
	go func() { serveErr <- server.Serve(listener) }()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()
		_ = server.Shutdown(shutdownCtx)
		<-serveErr
		return nil
	case err := <-serveErr:
		if err != nil && err != http.ErrServerClosed {
			return fmt.Errorf(i18n.T(i18n.WebErrServe), err)
		}
		return nil
	}
}
