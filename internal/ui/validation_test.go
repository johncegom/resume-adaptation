package ui

import (
	"os"
	"path/filepath"
	"testing"
)

func TestValidateInputPath(t *testing.T) {
	// Create a temporary file
	tempFile, err := os.CreateTemp("", "test_file_*.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "test_dir_*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "valid file path",
			path:    tempFile.Name(),
			wantErr: false,
		},
		{
			name:    "directory path",
			path:    tempDir,
			wantErr: true,
		},
		{
			name:    "non-existent path",
			path:    filepath.Join(tempDir, "does-not-exist.txt"),
			wantErr: true,
		},
		{
			name:    "empty path",
			path:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateInputPath(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateInputPath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
