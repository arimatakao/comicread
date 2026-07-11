package tui

func isNextKey(key string) bool {
	switch key {
	case "right", "l", "space", "pgdown", "j":
		return true
	default:
		return false
	}
}

func isPreviousKey(key string) bool {
	switch key {
	case "left", "h", "backspace", "pgup", "k":
		return true
	default:
		return false
	}
}
