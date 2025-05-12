package tui

import (
	"fmt"
)

type infraItem struct {
	title string
	used  bool
}

func (i infraItem) Title() string {
	checked := "[ ]"
	if i.used {
		checked = "[x]"
	}
	return fmt.Sprintf("%s %s", checked, i.title)
}
func (i infraItem) Description() string { return "" }
func (i infraItem) FilterValue() string { return i.title }
