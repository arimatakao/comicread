package reader

import tea "charm.land/bubbletea/v2"

// action names a reader command that one or more keys are bound to.
type action uint8

const (
	actionNone action = iota
	actionQuit
	actionHelp
	actionBookmark
	actionBookmarkPrefix
	actionBookmarkListPrefix
	actionNext
	actionPrevious
	actionScrollDown
	actionScrollUp
	actionZoomIn
	actionZoomOut
	actionGoToPagePrefix
)

// keyAction maps a key press to the reader action it is bound to.
func keyAction(msg tea.KeyPressMsg) action {
	switch msg.String() {
	case "q", "esc", "ctrl+c":
		return actionQuit
	case "?":
		return actionHelp
	case "b":
		return actionBookmark
	case "v":
		return actionBookmarkPrefix
	case "c":
		return actionBookmarkListPrefix
	case "g":
		return actionGoToPagePrefix
	case "right", "l", "space", "pgdown", "j":
		return actionNext
	case "left", "h", "backspace", "pgup", "k":
		return actionPrevious
	case "down":
		return actionScrollDown
	case "up":
		return actionScrollUp
	}

	// Zoom keys are matched on the key code so keypad variants work too.
	switch {
	case msg.Text == "+" || msg.Code == '+' || msg.ShiftedCode == '+' || msg.Code == tea.KeyKpPlus:
		return actionZoomIn
	case msg.Text == "-" || msg.Code == '-' || msg.ShiftedCode == '-' || msg.Code == tea.KeyKpMinus:
		return actionZoomOut
	}
	return actionNone
}
