package tui

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDBItem_Title_Description_FilterValue(t *testing.T) {
	item := dbItem{title: "Postgres", code: "pg", selected: true}
	require.Equal(t, "[✓] Postgres", item.Title())
	require.Equal(t, "Database option", item.Description())
	require.Equal(t, "Postgres", item.FilterValue())

	item.selected = false
	require.Equal(t, "[ ] Postgres", item.Title())
}
