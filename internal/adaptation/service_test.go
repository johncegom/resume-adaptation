package adaptation

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/johncegom/resume-adaptation/internal/parser"
	"google.golang.org/genai"
)

func TestService_Adapt_Success(t *testing.T) {
	t.Setenv("GEMINI_API_KEY", "test-key")

	mockAnalysisJSON := `{"keywords": ["Go", "Docker"], "responsibilities": ["Build APIs"], "requirements": ["Go experience"]}`
	mockResponseJSON := `{
		"summary": "Adapted Summary.",
		"experience": [
			{
				"company": "Acme Corp",
				"title": "Software Engineer",
				"achievements": ["Developed Go microservices"]
			}
		],
		"skills": ["Go", "Docker"]
	}`

	requestCount := 0
	transport := &mockTransport{
		roundTrip: func(req *http.Request) (*http.Response, error) {
			requestCount++
			var payload string
			if requestCount == 1 {
				payload = mockAnalysisJSON
			} else {
				payload = mockResponseJSON
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
			}`, payload)

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
	rewriter := NewRewriter(c)
	service := NewService(analyzer, rewriter)

	resume := &parser.Resume{
		Name:    "Jane Doe",
		Summary: "Original Summary.",
		Experience: []parser.WorkExperience{
			{
				Company:      "Acme Corp",
				Title:        "Software Engineer",
				StartDate:    "2020-01",
				EndDate:      "Present",
				Achievements: []string{"Developed software"},
			},
		},
		Skills: []string{"Software Development"},
	}

	adapted, err := service.Adapt(context.Background(), resume, "Job Description")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if adapted.Name != resume.Name {
		t.Errorf("expected name to be preserved, got %q", adapted.Name)
	}
	if adapted.Summary != "Adapted Summary." {
		t.Errorf("unexpected adapted summary: %q", adapted.Summary)
	}
	if len(adapted.Experience) != 1 {
		t.Fatalf("expected 1 work experience, got %d", len(adapted.Experience))
	}
	exp := adapted.Experience[0]
	if exp.Company != "Acme Corp" || exp.Title != "Software Engineer" {
		t.Errorf("company/title altered: %s - %s", exp.Company, exp.Title)
	}
	if len(exp.Achievements) != 1 || exp.Achievements[0] != "Developed Go microservices" {
		t.Errorf("unexpected achievements: %v", exp.Achievements)
	}
	if len(adapted.Skills) != 2 || adapted.Skills[0] != "Go" {
		t.Errorf("unexpected skills: %v", adapted.Skills)
	}
	if requestCount != 2 {
		t.Errorf("expected exactly 2 API requests, got %d", requestCount)
	}
}

func TestService_Adapt_NilResume(t *testing.T) {
	t.Setenv("GEMINI_API_KEY", "test-key")
	c, err := NewClient(context.Background(), withFactory(fakeClientFactory))
	if err != nil {
		t.Fatalf("unexpected error creating client: %v", err)
	}

	analyzer := NewJobAnalyzer(c)
	rewriter := NewRewriter(c)
	service := NewService(analyzer, rewriter)

	_, err = service.Adapt(context.Background(), nil, "Job Description")
	if err == nil {
		t.Fatal("expected error when resume is nil, got nil")
	}

	want := "resume cannot be nil"
	if err.Error() != want {
		t.Fatalf("unexpected error: got %q, want %q", err.Error(), want)
	}
}

func TestService_Adapt_EmptyJD(t *testing.T) {
	t.Setenv("GEMINI_API_KEY", "test-key")
	c, err := NewClient(context.Background(), withFactory(fakeClientFactory))
	if err != nil {
		t.Fatalf("unexpected error creating client: %v", err)
	}

	analyzer := NewJobAnalyzer(c)
	rewriter := NewRewriter(c)
	service := NewService(analyzer, rewriter)

	resume := &parser.Resume{Name: "Jane Doe"}
	_, err = service.Adapt(context.Background(), resume, "")
	if err == nil {
		t.Fatal("expected error when job description is empty, got nil")
	}

	want := "job description cannot be empty"
	if err.Error() != want {
		t.Fatalf("unexpected error: got %q, want %q", err.Error(), want)
	}
}

func TestService_Adapt_AnalyzerError(t *testing.T) {
	t.Setenv("GEMINI_API_KEY", "test-key")

	transport := &mockTransport{
		roundTrip: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       io.NopCloser(strings.NewReader("analyzer failure")),
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
	rewriter := NewRewriter(c)
	service := NewService(analyzer, rewriter)

	resume := &parser.Resume{Name: "Jane Doe"}
	_, err = service.Adapt(context.Background(), resume, "Job Description")
	if err == nil {
		t.Fatal("expected error from analyzer failure, got nil")
	}

	if !strings.Contains(err.Error(), "failed to generate content") {
		t.Fatalf("unexpected error message: %v", err)
	}
}

func TestService_Adapt_RewriterError(t *testing.T) {
	t.Setenv("GEMINI_API_KEY", "test-key")

	requestCount := 0
	transport := &mockTransport{
		roundTrip: func(req *http.Request) (*http.Response, error) {
			requestCount++
			if requestCount == 1 {
				mockAnalysisJSON := `{"keywords": ["Go"], "responsibilities": [], "requirements": []}`
				geminiResp := fmt.Sprintf(`{"candidates": [{"content": {"parts": [{"text": %q}]}}]}`, mockAnalysisJSON)
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(geminiResp)),
					Header:     make(http.Header),
				}, nil
			}
			return &http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       io.NopCloser(strings.NewReader("rewriter failure")),
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
	rewriter := NewRewriter(c)
	service := NewService(analyzer, rewriter)

	resume := &parser.Resume{Name: "Jane Doe"}
	_, err = service.Adapt(context.Background(), resume, "Job Description")
	if err == nil {
		t.Fatal("expected error from rewriter failure, got nil")
	}

	if !strings.Contains(err.Error(), "failed to generate content") {
		t.Fatalf("unexpected error message: %v", err)
	}
}
