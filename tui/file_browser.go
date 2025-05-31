package tui

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// fileItem represents a file or directory in the file browser
type fileItem struct {
	path     string
	name     string
	isDir    bool
	selected bool
}

// FilterValue implements list.Item interface
func (i fileItem) FilterValue() string {
	return i.name
}

// Title returns the name of the file or directory
func (i fileItem) Title() string {
	if i.isDir {
		return fmt.Sprintf("📁 %s", i.name)
	}
	return fmt.Sprintf("📄 %s", i.name)
}

// Description returns additional information about the file
func (i fileItem) Description() string {
	if i.isDir {
		return "Directory"
	}
	return "File"
}

// FileBrowser is a component for browsing and selecting files
type FileBrowser struct {
	list         list.Model
	currentPath  string
	selectedFile string
	err          error
}

// NewFileBrowser creates a new file browser starting at the given path
func NewFileBrowser(startPath string, width, height int) (*FileBrowser, error) {
	// If startPath is empty, use the current directory
	if startPath == "" {
		var err error
		startPath, err = os.Getwd()
		if err != nil {
			return nil, err
		}
	}

	// Make sure the path is absolute
	absPath, err := filepath.Abs(startPath)
	if err != nil {
		return nil, err
	}

	// Create the file browser
	fb := &FileBrowser{
		currentPath: absPath,
	}

	// Initialize the list
	items, err := fb.getItems()
	if err != nil {
		return nil, err
	}

	// Create the list model
	fb.list = list.New(items, list.NewDefaultDelegate(), width, height)
	fb.list.Title = "Select a file (use [↑/↓] to navigate, [Enter] to select)"
	fb.list.SetShowStatusBar(false)
	fb.list.SetFilteringEnabled(false)
	fb.list.SetShowHelp(true)
	fb.list.SetWidth(width)
	fb.list.SetHeight(height)

	return fb, nil
}

// getItems returns a list of files and directories in the current path
func (fb *FileBrowser) getItems() ([]list.Item, error) {
	// Read the directory
	entries, err := os.ReadDir(fb.currentPath)
	if err != nil {
		return nil, err
	}

	// Create items for each entry
	items := make([]list.Item, 0, len(entries)+1)

	// Add parent directory if not at root
	if fb.currentPath != "/" {
		items = append(items, fileItem{
			path:  filepath.Dir(fb.currentPath),
			name:  "..",
			isDir: true,
		})
	}

	// Add files and directories
	for _, entry := range entries {
		// Skip hidden files
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		items = append(items, fileItem{
			path:  filepath.Join(fb.currentPath, entry.Name()),
			name:  entry.Name(),
			isDir: entry.IsDir(),
		})
	}

	// Sort items: directories first, then files, both alphabetically
	sort.Slice(items, func(i, j int) bool {
		itemI := items[i].(fileItem)
		itemJ := items[j].(fileItem)

		// Special case for parent directory
		if itemI.name == ".." {
			return true
		}
		if itemJ.name == ".." {
			return false
		}

		// Directories before files
		if itemI.isDir && !itemJ.isDir {
			return true
		}
		if !itemI.isDir && itemJ.isDir {
			return false
		}

		// Alphabetical order
		return itemI.name < itemJ.name
	})

	return items, nil
}

// Update handles user input and updates the file browser state
func (fb *FileBrowser) Update(msg tea.Msg) (*FileBrowser, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			// Get the selected item
			item, ok := fb.list.SelectedItem().(fileItem)
			if !ok {
				return fb, nil
			}

			// If it's a directory, navigate to it
			if item.isDir {
				fb.currentPath = item.path
				items, err := fb.getItems()
				if err != nil {
					fb.err = err
					return fb, nil
				}
				fb.list.SetItems(items)
				return fb, nil
			}

			// If it's a file, select it
			fb.selectedFile = item.path
			return fb, nil
		}
	}

	// Update the list
	fb.list, cmd = fb.list.Update(msg)
	return fb, cmd
}

// View renders the file browser
func (fb *FileBrowser) View() string {
	if fb.err != nil {
		return fmt.Sprintf("Error: %s", fb.err)
	}

	// Show the current path and the list
	return fmt.Sprintf("Current directory: %s\n\n%s", fb.currentPath, fb.list.View())
}

// SelectedFile returns the selected file path
func (fb *FileBrowser) SelectedFile() string {
	return fb.selectedFile
}

// SetSelectedFile sets the selected file path
func (fb *FileBrowser) SetSelectedFile(path string) {
	fb.selectedFile = path
}

// CurrentPath returns the current directory path
func (fb *FileBrowser) CurrentPath() string {
	return fb.currentPath
}
