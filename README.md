# comicread

Minimal terminal manga reader written in Go.

The current MVP opens CBZ, PDF, EPUB, and image directories, renders pages with
the Kitty, Sixel, or iTerm2 graphics protocol, or as text art, and supports
forward/backward navigation.

## Installation

### Linux and macOS

```sh
curl -fsSL https://raw.githubusercontent.com/arimatakao/comicread/main/install.sh | bash
```

<details>
<summary>Manual installation or running</summary>

Download the archive for your operating system and architecture from the [latest release](https://github.com/arimatakao/comicread/releases/latest), then extract it.

On Linux or macOS:

```sh
tar -xzf comicread_*_linux_*.tar.gz # use *_darwin_* on macOS
./comicread /path/to/file.pdf # or file.cbz / file.epub
```

</details>

### Windows

```powershell
powershell -ExecutionPolicy Bypass -Command "iwr -useb https://raw.githubusercontent.com/arimatakao/comicread/main/install.ps1 | iex"
```


<details>
<summary>Manual installation or running</summary>

On Windows:

```powershell
Expand-Archive .\comicread_*_windows_*.zip -DestinationPath .\comicread
.\comicread\comicread.exe C:\path\to\file.pdf # or file.cbz / file.epub
```

Move the extracted executable to a directory in your `PATH` to install it manually.

</details>


## Terminal compatibility

`kitty`, `sixel`, and `iterm2` render raster images. `ascii` and `dots` are
text-art fallbacks: they need ANSI cursor control and colour (true colour is
recommended); `dots` also needs Unicode Braille support. The latter two
therefore work in virtually every current UTF-8 terminal, including terminals
without an image protocol.

| Terminal | `kitty` | `sixel` | `iterm2` | `ascii` | `dots` | Notes |
| --- | :---: | :---: | :---: | :---: | :---: | --- |
| Any current ANSI/UTF-8 terminal | - | - | - | X | X | Fallback modes; use `ascii` if Braille glyphs are unavailable. |
| [GNOME Terminal](https://gitlab.gnome.org/GNOME/gnome-terminal) | - | - | - | X | X | VTE-based; its optional SIXEL support is disabled by default. |
| [GNOME Console](https://apps.gnome.org/Console/) | - | - | - | X | X | VTE-based; its optional SIXEL support is disabled by default. |
| [Xfce Terminal](https://docs.xfce.org/apps/xfce4-terminal/start) | - | - | - | X | X | VTE-based; its optional SIXEL support is disabled by default. |
| [Tilix](https://github.com/gnunn1/tilix) | - | - | - | X | X | VTE-based; its optional SIXEL support is disabled by default. |
| [Terminator](https://gnome-terminator.readthedocs.io/) | - | - | - | X | X | VTE-based; its optional SIXEL support is disabled by default. |
| [MATE Terminal](https://wiki.mate-desktop.org/mate-desktop/applications/mate-terminal/) | - | - | - | X | X | Standard MATE terminal; no image protocol support. |
| [LXTerminal](https://github.com/lxde/lxterminal) | - | - | - | X | X | VTE-based; its optional SIXEL support is disabled by default. |
| [Guake](https://github.com/Guake/guake) | - | - | - | X | X | VTE-based drop-down terminal; SIXEL is disabled by default. |
| [Alacritty](https://github.com/alacritty/alacritty) | - | - | - | X | X | No support for these image protocols in the standard build. |
| [Kitty](https://sw.kovidgoyal.net/kitty/) | X | - | - | X | X | Native implementation of the Kitty graphics protocol. |
| [Ghostty](https://ghostty.org/docs/features) | X | - | - | X | X | Kitty graphics protocol. |
| [Konsole](https://konsole.kde.org/) | X | X | X | X | X | iTerm2 support is available since 22.04; animated images are limited. |
| [WezTerm](https://wezterm.org/features.html) | X | X | X | X | X | Enable `enable_kitty_graphics=true` for `kitty`; Sixel is experimental. |
| [iTerm2](https://iterm2.com/3.5/documentation-images.html) | X | X | X | X | X | Supports all three; its [release notes](https://iterm2.com/downloads.html?cve=title) document SIXEL and Kitty support. |
| [Warp](https://www.warp.dev/) | X | - | - | X | X | Implements the Kitty graphics protocol. |
| [wayst](https://github.com/91861/wayst) | X | - | - | X | X | Implements the Kitty graphics protocol. |
| [st](https://st.suckless.org/) with a graphics patch | X | - | - | X | X | Requires a Kitty-graphics implementation patch. |
| [xterm.js](https://xtermjs.org/) host | X | X | X | X | X | Requires the image add-on; its Kitty implementation is partial. |
| [xterm](https://invisible-island.net/xterm/ctlseqs/ctlseqs.html) | - | X | - | X | X | Must be built with SIXEL and configured as a DEC graphics terminal. |
| [mlterm](https://mlterm.sourceforge.net/) | - | X | X | X | X | `iterm2` requires a build with `SUPPORT_ITERM2_OSC1337`. |
| [foot](https://codeberg.org/dnkl/foot) | - | X | - | X | X | |
| [mintty](https://mintty.github.io/) / wsltty | - | X | X | X | X | Windows/Cygwin terminal and its WSL variant. |
| [Contour](https://contour-terminal.org/configuration/) | - | X | - | X | X | |
| [DomTerm](https://domterm.org/Features.html) | - | X | - | X | X | |
| [yaft](https://github.com/uobikiemukot/yaft) | - | X | - | X | X | Linux framebuffer terminal. |
| [SyncTERM](https://www.syncterm.net/) | - | X | - | X | X | |

## Run

```sh
go run . path/to/chapter.cbz
go run . path/to/image-directory
go run . --graphics sixel path/to/chapter.cbz
go run . --graphics iterm2 path/to/chapter.cbz
go run . --graphics dots path/to/chapter.cbz
```

`--graphics` accepts `auto` (the default), `kitty`, `sixel`, `iterm2`, `ascii`,
or `dots`. Auto-detection selects iTerm2 and the known Sixel terminals `mlterm`,
`yaft`, `DomTerm`, `Contour`, `mintty`, and `foot`; otherwise it defaults to
Kitty. Use an explicit backend for a terminal not detected this way.

Keys:

- `right`, `l`, `space`, `j`, `PageDown`: next page
- `left`, `h`, `backspace`, `k`, `PageUp`: previous page
- `+`, `-`: zoom in/out for every page in the open file
- `up`, `down`: scroll the zoomed page vertically
- `q`, `Esc`, `Ctrl+C`: quit

Supported formats: CBZ, image-based PDF files, image-based EPUB files, and directories containing image files.

PDF pages must contain one embedded raster image per page. EPUB pages must
reference their page images through the EPUB spine.

## Dependency licenses

The main direct dependencies are distributed under permissive licenses:

- [Bubble Tea v2](https://github.com/charmbracelet/bubbletea) — MIT
- [rasterm](https://github.com/BourgeoisBear/rasterm) — MIT
- [comicfile](https://github.com/arimatakao/comicfile) — MIT
- [ASCIIimage v2](https://github.com/fandasy/ASCIIimage) — MIT
- [dots](https://github.com/imjasonh/dots) — Apache License 2.0
- [golang.org/x/image](https://pkg.go.dev/golang.org/x/image) — BSD-style (Go Authors)

See the `LICENSE` file bundled with each dependency version in the Go module
cache for its full license text. Transitive dependencies may have their own
license terms.
