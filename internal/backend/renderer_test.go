package backend

import "testing"

func TestNewRenderer(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		protocol string
		name     string
	}{
		{"kitty", "kitty"},
		{"sixel", "sixel"},
	} {
		t.Run(test.protocol, func(t *testing.T) {
			renderer, err := NewRenderer(test.protocol)
			if err != nil {
				t.Fatalf("NewRenderer() error = %v", err)
			}
			if renderer.Name() != test.name {
				t.Errorf("NewRenderer().Name() = %q, want %q", renderer.Name(), test.name)
			}
		})
	}
}

func TestNewRendererRejectsUnknownProtocol(t *testing.T) {
	t.Parallel()

	if _, err := NewRenderer("iterm2"); err == nil {
		t.Fatal("NewRenderer() error = nil, want unsupported protocol error")
	}
}

func TestSupportsSixel(t *testing.T) {
	t.Parallel()

	if !supportsSixel("xterm-256color", "mlterm") {
		t.Fatal("supportsSixel() = false, want true for mlterm")
	}
	if supportsSixel("xterm-256color", "kitty") {
		t.Fatal("supportsSixel() = true, want false for Kitty")
	}
}
