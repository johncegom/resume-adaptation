package adaptation

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"google.golang.org/genai"
)

func TestJobAnalyzer_Analyze_Success(t *testing.T) {
	t.Setenv("GEMINI_API_KEY", "test-key")

	mockResponseJSON := `{
		"keywords": ["Go", "Docker", "REST API"],
		"responsibilities": ["Develop microservices", "Optimize database queries"],
		"requirements": ["3+ years of experience", "Degree in Computer Science"]
	}`

	transport := &mockTransport{
		roundTrip: func(req *http.Request) (*http.Response, error) {
			if req.Method != http.MethodPost {
				t.Errorf("expected POST request, got %q", req.Method)
			}
			if !strings.Contains(req.URL.Host, "generativelanguage.googleapis.com") {
				t.Errorf("unexpected host: %q", req.URL.Host)
			}

			geminiResp := fmt.Sprintf(`{
				"candidates": [
					{
						"content": {
							"parts": [
								{
									"text": %q
								}
							]
						}
					}
				]
			}`, mockResponseJSON)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(geminiResp)),
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

	analyzer := NewJobAnalyzer(c)
	analysis, err := analyzer.Analyze(context.Background(), "Target Job Description")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(analysis.Keywords) != 3 || analysis.Keywords[0] != "Go" {
		t.Errorf("unexpected keywords: %v", analysis.Keywords)
	}
	if len(analysis.Responsibilities) != 2 || analysis.Responsibilities[0] != "Develop microservices" {
		t.Errorf("unexpected responsibilities: %v", analysis.Responsibilities)
	}
	if len(analysis.Requirements) != 2 || analysis.Requirements[0] != "3+ years of experience" {
		t.Errorf("unexpected requirements: %v", analysis.Requirements)
	}
}

func TestJobAnalyzer_Analyze_EmptyJD(t *testing.T) {
	t.Setenv("GEMINI_API_KEY", "test-key")

	c, err := NewClient(context.Background(), withFactory(fakeClientFactory))
	if err != nil {
		t.Fatalf("unexpected error creating client: %v", err)
	}

	analyzer := NewJobAnalyzer(c)
	_, err = analyzer.Analyze(context.Background(), "")
	if err == nil {
		t.Fatal("expected error with empty job description, got nil")
	}

	want := "job description cannot be empty"
	if err.Error() != want {
		t.Fatalf("unexpected error: got %q, want %q", err.Error(), want)
	}
}

func TestJobAnalyzer_Analyze_APIError(t *testing.T) {
	t.Setenv("GEMINI_API_KEY", "test-key")

	transport := &mockTransport{
		roundTrip: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       io.NopCloser(strings.NewReader("internal server error")),
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

	analyzer := NewJobAnalyzer(c)
	_, err = analyzer.Analyze(context.Background(), "Target Job Description")
	if err == nil {
		t.Fatal("expected error from API failure, got nil")
	}

	if !strings.Contains(err.Error(), "failed to generate content") {
		t.Fatalf("unexpected error message: %v", err)
	}
}
