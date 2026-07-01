package ui

import (
	"testing"
)

func TestStylesInitialization(t *testing.T) {
	// Verify that style variables are initialized and non-empty
	if StyleTitle.String() == "" {
		t.Error("expected StyleTitle to be initialized")
	}
	if StyleBox.String() == "" {
		t.Error("expected StyleBox to be initialized")
	}
}
