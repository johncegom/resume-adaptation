package parser_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/johncegom/resume-adaptation/internal/parser"
)

func TestReadPlaintextFileTxt(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "resume.txt")
	content := "Jane Doe\nSenior Software Engineer\n\nExperience:\n- Built distributed systems"
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	resume, err := parser.ReadPlaintextFile(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resume.RawContent == "" {
		t.Fatal("expected RawContent to be populated")
	}
	if resume.RawContent != content {
		t.Errorf("RawContent: got %q, want %q", resume.RawContent, content)
	}
}

func TestReadPlaintextFileMd(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "resume.md")
	content := "# Jane Doe\n\n## Experience\n- Built distributed systems"
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	resume, err := parser.ReadPlaintextFile(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resume.RawContent == "" {
		t.Fatal("expected RawContent to be populated")
	}
	if resume.RawContent != content {
		t.Errorf("RawContent: got %q, want %q", resume.RawContent, content)
	}
}

func TestReadPlaintextFileNotFound(t *testing.T) {
	_, err := parser.ReadPlaintextFile("non_existent_file.txt")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestReadPlaintextFileUnsupportedExtension(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "resume.docx")
	if err := os.WriteFile(path, []byte("some content"), 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	_, err := parser.ReadPlaintextFile(path)
	if err == nil {
		t.Fatal("expected error for unsupported extension, got nil")
	}
}

func TestReadPlaintextFileEmpty(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "resume.txt")
	if err := os.WriteFile(path, []byte("   \n  "), 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	_, err := parser.ReadPlaintextFile(path)
	if err == nil {
		t.Fatal("expected error for empty file, got nil")
	}
}

func TestReadPlaintextFilePathTraversal(t *testing.T) {
	_, err := parser.ReadPlaintextFile("some/path/../../secret.txt")
	if err == nil {
		t.Fatal("expected error for path traversal, got nil")
	}
}
