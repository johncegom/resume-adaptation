package main

import (
	"testing"
)

func TestAdaptCommandRegistration(t *testing.T) {
	rootCmd := NewRootCommand()
	
	// Check root usage
	if rootCmd.Use != "resume-adaptation" {
		t.Errorf("expected root command use 'resume-adaptation', got %q", rootCmd.Use)
	}

	// Verify adapt subcommand is registered
	adaptCmd, _, err := rootCmd.Find([]string{"adapt"})
	if err != nil {
		t.Fatalf("failed to find 'adapt' subcommand: %v", err)
	}
	if adaptCmd == nil {
		t.Fatal("expected 'adapt' subcommand to be non-nil")
	}

	if adaptCmd.SilenceUsage != true {
		t.Error("expected SilenceUsage to be true on adapt command")
	}
	if adaptCmd.SilenceErrors != true {
		t.Error("expected SilenceErrors to be true on adapt command")
	}
}
