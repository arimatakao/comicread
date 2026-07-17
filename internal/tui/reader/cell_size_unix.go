//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris

package reader

import (
	"os"

	"golang.org/x/sys/unix"
)

// terminalCellSize reads the terminal's actual pixel geometry. Bubble Tea
// supplies the cell count, while TIOCGWINSZ supplies the corresponding pixels.
func terminalCellSize(cols, rows int) (width, height int) {
	if cols < 1 || rows < 1 {
		return 0, 0
	}
	size, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil || size.Xpixel < uint16(cols) || size.Ypixel < uint16(rows) {
		return 0, 0
	}
	return int(size.Xpixel) / cols, int(size.Ypixel) / rows
}
