package ui

import (
	"context"
	"errors"
	"testing"
)

func TestRunLoadingSpinner(t *testing.T) {
	ctx := context.Background()

	// Test success task
	err := RunLoadingSpinner(ctx, "Success Task", func() error {
		return nil
	})
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}

	// Test failing task
	expectedErr := errors.New("failing task error")
	err = RunLoadingSpinner(ctx, "Failing Task", func() error {
		return expectedErr
	})
	if err != expectedErr {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}
}
