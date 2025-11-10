package tui

import "fmt"

type wsItem struct {
	title    string
	selected bool
}

func (w wsItem) Title() string {
	radio := "[ ]"
	if w.selected {
		radio = "[✓]"
	}

	return fmt.Sprintf("%s %s", radio, w.title)
}

func (w wsItem) Description() string {
	return "WebSocket support with JWT token authentication"
}
func (w wsItem) FilterValue() string { return w.title }
