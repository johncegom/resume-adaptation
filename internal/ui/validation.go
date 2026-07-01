package ui

import (
	"errors"
	"fmt"
	"os"
)

// ValidateInputPath checks if a file path is non-empty, exists, and is a file (not a directory).
func ValidateInputPath(path string) error {
	if path == "" {
		return errors.New("path cannot be empty")
	}

	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file does not exist: %s", path)
		}
		return err
	}

	if info.IsDir() {
		return fmt.Errorf("path is a directory, expected a file: %s", path)
	}

	return nil
}
