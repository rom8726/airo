package tui

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInfraItem_Title_Description_FilterValue(t *testing.T) {
	item := infraItem{title: "Redis", code: "redis", used: true}
	require.Equal(t, "[✓] Redis", item.Title())
	require.Equal(t, "In-memory cache", item.Description())
	require.Equal(t, "Redis", item.FilterValue())

	item.used = false
	require.Equal(t, "[ ] Redis", item.Title())
}
