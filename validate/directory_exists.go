package validate

import (
	"errors"
	"os"
	"path/filepath"
)

// DirectoryExists checks if a directory with the given name exists in the current working directory.
// Returns an error if the directory exists, nil otherwise.
func DirectoryExists(name string) error {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return errors.New("failed to get current directory")
	}

	// Construct the full path to the potential directory
	dirPath := filepath.Join(cwd, name)

	// Check if the directory exists
	info, err := os.Stat(dirPath)
	if err == nil && info.IsDir() {
		return errors.New("directory already exists, please choose a different name or delete the existing directory")
	}

	return nil
}
