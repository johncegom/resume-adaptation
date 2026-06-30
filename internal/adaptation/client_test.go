package adaptation

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"google.golang.org/genai"
)

type mockTransport struct {
	roundTrip func(*http.Request) (*http.Response, error)
}

func (m *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.roundTrip(req)
}

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

func TestClient_GenerateContent_Success(t *testing.T) {
	t.Setenv("GEMINI_API_KEY", "test-key")

	mockResponse := `{"candidates": [{"content": {"parts": [{"text": "Hello World"}]}}]}`

	transport := &mockTransport{
		roundTrip: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResponse)),
				Header:     make(http.Header),
			}, nil
		},
	}

	c, err := NewClient(context.Background(), withFactory(func(ctx context.Context, config *genai.ClientConfig) (*genai.Client, error) {
		config.HTTPClient = &http.Client{Transport: transport}
		return genai.NewClient(ctx, config)
	}))
	if err != nil {
		t.Fatalf("unexpected error creating client: %v", err)
	}

	resp, err := c.GenerateContent(context.Background(), []*genai.Content{
		{Parts: []*genai.Part{{Text: "Test Prompt"}}},
	}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(resp.Candidates) == 0 || resp.Candidates[0].Content.Parts[0].Text != "Hello World" {
		t.Fatalf("unexpected response: %+v", resp)
	}
}
