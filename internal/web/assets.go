package web

import _ "embed"

// Embedded frontend assets served as-is. Bundling them into the binary keeps
// the web reader self-contained: no CDN, no Node.js build step, and nothing
// to install alongside the comicread binary.
var (
	//go:embed assets/index.html
	indexHTML []byte

	//go:embed assets/app.js
	appJS []byte

	//go:embed assets/pico.min.css
	picoCSS []byte
)
