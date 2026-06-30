package adaptation

import (
	"context"
	"testing"
	"time"

	"google.golang.org/genai"
)

// fakeClientFactory returns a no-op genai.Client for unit tests that
// need NewClient to succeed without making network calls.
func fakeClientFactory(_ context.Context, _ *genai.ClientConfig) (*genai.Client, error) {
	return &genai.Client{}, nil
}

func TestNewClient_missing_api_key(t *testing.T) {
	t.Setenv("GEMINI_API_KEY", "")

	_, err := NewClient(context.Background())
	if err == nil {
		t.Fatal("expected error when GEMINI_API_KEY is not set, got nil")
	}

	want := "environment variable GEMINI_API_KEY is not set"
	if err.Error() != want {
		t.Fatalf("unexpected error message: got %q, want %q", err.Error(), want)
	}
}

func TestNewClient_default_model(t *testing.T) {
	t.Setenv("GEMINI_API_KEY", "test-key-for-unit-test")

	c, err := NewClient(context.Background(), withFactory(fakeClientFactory))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if c.Model() != "gemini-2.5-flash" {
		t.Fatalf("default model: got %q, want %q", c.Model(), "gemini-2.5-flash")
	}
}

func TestNewClient_custom_model(t *testing.T) {
	t.Setenv("GEMINI_API_KEY", "test-key-for-unit-test")

	c, err := NewClient(context.Background(), withFactory(fakeClientFactory), WithModel("gemini-2.5-pro"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if c.Model() != "gemini-2.5-pro" {
		t.Fatalf("custom model: got %q, want %q", c.Model(), "gemini-2.5-pro")
	}
}

func TestNewClient_custom_timeout(t *testing.T) {
	t.Setenv("GEMINI_API_KEY", "test-key-for-unit-test")

	c, err := NewClient(context.Background(), withFactory(fakeClientFactory), WithTimeout(10*time.Second))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if c.Timeout() != 10*time.Second {
		t.Fatalf("custom timeout: got %v, want %v", c.Timeout(), 10*time.Second)
	}
}

func TestNewClient_default_timeout(t *testing.T) {
	t.Setenv("GEMINI_API_KEY", "test-key-for-unit-test")

	c, err := NewClient(context.Background(), withFactory(fakeClientFactory))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if c.Timeout() != 30*time.Second {
		t.Fatalf("default timeout: got %v, want %v", c.Timeout(), 30*time.Second)
	}
}
