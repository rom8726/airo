package tui

import (
	"fmt"
)

type dbItem struct {
	title    string
	code     string
	selected bool
}

func (d dbItem) Title() string {
	radio := "( )"
	if d.selected {
		radio = "(x)"
	}
	return fmt.Sprintf("%s %s", radio, d.title)
}
func (d dbItem) Description() string { return "" }
func (d dbItem) FilterValue() string { return d.title }
