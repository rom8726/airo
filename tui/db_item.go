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
	radio := "[ ]"
	if d.selected {
		radio = "[✓]"
	}
	return fmt.Sprintf("%s %s", radio, d.title)
}
func (d dbItem) Description() string {
	switch d.code {
	case "postgres":
		return "PostgreSQL - SQL database"
	case "mysql":
		return "MySQL - SQL database"
	case "mongodb":
		return "MongoDB - NoSQL database"
	default:
		return "Database option"
	}
}
func (d dbItem) FilterValue() string { return d.title }
