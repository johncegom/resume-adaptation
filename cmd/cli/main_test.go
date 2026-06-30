package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEnvLoading(t *testing.T) {
	// Create a temporary .env file in the current test directory
	tempEnvPath := filepath.Join(".", ".env")
	envContent := "TEST_ENV_VAR=loaded_successfully"
	err := os.WriteFile(tempEnvPath, []byte(envContent), 0644)
	if err != nil {
		t.Fatalf("failed to create temp .env file: %v", err)
	}
	defer os.Remove(tempEnvPath)

	// Run the env loading setup (to be implemented in main.go)
	setupEnv()

	// Verify environment variable is loaded
	val := os.Getenv("TEST_ENV_VAR")
	if val != "loaded_successfully" {
		t.Errorf("expected TEST_ENV_VAR to be 'loaded_successfully', got '%s'", val)
	}
}
