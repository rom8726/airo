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
		return "PostgreSQL database for robust relational data storage"
	case "mysql":
		return "MySQL database for web applications"
	case "mongodb":
		return "MongoDB for document-oriented NoSQL storage"
	default:
		return "Database option"
	}
}
func (d dbItem) FilterValue() string { return d.title }
