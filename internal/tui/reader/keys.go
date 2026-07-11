package reader

import tea "charm.land/bubbletea/v2"

func isHelpKey(key string) bool {
	return key == "?"
}

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

func isScrollDownKey(key string) bool {
	return key == "down"
}

func isScrollUpKey(key string) bool {
	return key == "up"
}

func isZoomInKey(key tea.KeyPressMsg) bool {
	return key.Text == "+" || key.Code == '+' || key.ShiftedCode == '+' || key.Code == tea.KeyKpPlus
}

func isZoomOutKey(key tea.KeyPressMsg) bool {
	return key.Text == "-" || key.Code == '-' || key.ShiftedCode == '-' || key.Code == tea.KeyKpMinus
}
