package web

import (
	"os/exec"
	"runtime"
)

// openBrowser best-effort launches the system default browser at url.
// Failures are ignored: the URL is always printed to the terminal by the
// caller as a fallback, so a missing browser or headless environment is not
// fatal.
func openBrowser(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	_ = cmd.Start()
}
