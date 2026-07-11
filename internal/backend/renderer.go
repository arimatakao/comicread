package backend

import (
	"fmt"
	"os"
	"strings"
)

// NewRenderer returns the requested renderer. The auto setting keeps Kitty as
// the default and selects iTerm2 or Sixel when the terminal identifies itself
// as supporting one of those protocols.
func NewRenderer(protocol string) (Renderer, error) {
	switch strings.ToLower(strings.TrimSpace(protocol)) {
	case "", "auto":
		if supportsIterm(os.Getenv("TERM"), os.Getenv("TERM_PROGRAM"), os.Getenv("LC_TERMINAL")) {
			return NewIterm(), nil
		}
		if supportsSixel(os.Getenv("TERM"), os.Getenv("TERM_PROGRAM")) {
			return NewSixel(), nil
		}
		return NewKitty(), nil
	case "kitty":
		return NewKitty(), nil
	case "sixel":
		return NewSixel(), nil
	case "iterm", "iterm2":
		return NewIterm(), nil
	case "ascii", "ansi":
		return NewASCII(), nil
	case "dots":
		return NewDots(), nil
	default:
		return nil, fmt.Errorf("unsupported graphics protocol %q (want auto, ascii, dots, kitty, sixel, or iterm2)", protocol)
	}
}

func supportsIterm(term, termProgram, lcTerminal string) bool {
	name := strings.ToLower(term + " " + termProgram + " " + lcTerminal)
	return strings.Contains(name, "iterm")
}

func supportsSixel(term, termProgram string) bool {
	name := strings.ToLower(term + " " + termProgram)
	for _, terminal := range []string{"mlterm", "yaft", "domterm", "contour", "mintty", "foot"} {
		if strings.Contains(name, terminal) {
			return true
		}
	}
	return false
}
