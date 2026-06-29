package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// supportedPlaintextExts lists the file extensions accepted by ReadPlaintextFile.
var supportedPlaintextExts = map[string]bool{
	".txt": true,
	".md":  true,
}

// ReadPlaintextFile reads a local plaintext file and returns a Resume
// with RawContent populated. It validates the file path, extension,
// and content before returning.
func ReadPlaintextFile(path string) (*Resume, error) {
	if err := validatePath(path); err != nil {
		return nil, err
	}

	ext := strings.ToLower(filepath.Ext(path))
	if !supportedPlaintextExts[ext] {
		return nil, fmt.Errorf("unsupported file extension %q: expected .txt or .md", ext)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	content := string(data)
	if strings.TrimSpace(content) == "" {
		return nil, fmt.Errorf("file is empty: %s", path)
	}

	return &Resume{RawContent: content}, nil
}

// validatePath checks that the path is safe and does not contain
// directory traversal patterns.
func validatePath(path string) error {
	if path == "" {
		return fmt.Errorf("file path is empty")
	}

	cleaned := filepath.Clean(path)
	if strings.Contains(cleaned, "..") {
		return fmt.Errorf("path traversal detected: %s", path)
	}

	return nil
}
