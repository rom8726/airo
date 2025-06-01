package tui

import "fmt"

type testyItem struct {
	title    string
	selected bool
}

func (t testyItem) Title() string {
	radio := "[ ]"
	if t.selected {
		radio = "[✓]"
	}

	return fmt.Sprintf("%s %s", radio, t.title)
}

func (t testyItem) Description() string {
	return "Includes test containers and utilities for integration testing"
}
func (t testyItem) FilterValue() string { return t.title }
