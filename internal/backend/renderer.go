package backend

import (
	"fmt"
	"os"
	"strings"
)

// NewRenderer returns the requested graphics renderer. The auto setting keeps
// Kitty as the default and selects Sixel for terminals that identify themselves
// as Sixel-capable.
func NewRenderer(protocol string) (Renderer, error) {
	switch strings.ToLower(strings.TrimSpace(protocol)) {
	case "", "auto":
		if supportsSixel(os.Getenv("TERM"), os.Getenv("TERM_PROGRAM")) {
			return NewSixel(), nil
		}
		return NewKitty(), nil
	case "kitty":
		return NewKitty(), nil
	case "sixel":
		return NewSixel(), nil
	default:
		return nil, fmt.Errorf("unsupported graphics protocol %q (want auto, kitty, or sixel)", protocol)
	}
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
