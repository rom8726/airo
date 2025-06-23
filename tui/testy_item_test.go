package tui

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTestyItem_Title_Description_FilterValue(t *testing.T) {
	item := testyItem{title: "Testy", selected: true}
	require.Equal(t, "[✓] Testy", item.Title())
	require.Equal(t, "Integration testing framework", item.Description())
	require.Equal(t, "Testy", item.FilterValue())

	item.selected = false
	require.Equal(t, "[ ] Testy", item.Title())
}
