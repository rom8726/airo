package tui

import (
	"os"
	"path/filepath"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/require"
)

func TestNewFileBrowser_Basic(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "file1.yml"), []byte("test"), 0644)
	os.Mkdir(filepath.Join(dir, "subdir"), 0755)
	fb, err := NewFileBrowser(dir, 10, 5)
	require.NoError(t, err)
	require.NotNil(t, fb)
	items, err := fb.getItems()
	require.NoError(t, err)
	var foundFile, foundDir bool
	for _, it := range items {
		fi := it.(fileItem)
		if fi.name == "file1.yml" {
			foundFile = true
		}
		if fi.name == "subdir" {
			foundDir = true
		}
	}
	require.True(t, foundFile)
	require.True(t, foundDir)
}

func TestFileBrowser_SelectFile(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "file1.yml"), []byte("test"), 0644)
	fb, err := NewFileBrowser(dir, 10, 5)
	require.NoError(t, err)
	fb.list.Select(0)
	for i, it := range fb.list.Items() {
		fi := it.(fileItem)
		if fi.name == "file1.yml" {
			fb.list.Select(i)
			break
		}
	}
	msg := tea.KeyMsg{Type: tea.KeyEnter}
	fb2, _ := fb.Update(msg)
	require.NotNil(t, fb2)

	require.Contains(t, fb2.SelectedFile(), "file1.yml")
}
