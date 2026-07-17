//go:build !aix && !darwin && !dragonfly && !freebsd && !linux && !netbsd && !openbsd && !solaris

package reader

// Some platforms do not expose terminal pixel dimensions through the standard
// window-size API. The Sixel renderer keeps its conservative default there.
func terminalCellSize(_, _ int) (width, height int) { return 0, 0 }
