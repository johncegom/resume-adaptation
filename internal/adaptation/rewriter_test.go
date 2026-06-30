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

func TestRewriter_Rewrite_Success(t *testing.T) {
	t.Setenv("GEMINI_API_KEY", "test-key")

	mockResponseJSON := `{
		"summary": "Adapted Summary highlighting Go and microservices experience.",
		"experience": [
			{
				"company": "Acme Corp",
				"title": "Software Engineer",
				"achievements": ["Developed Go microservices", "Improved performance by 20%"]
			}
		],
		"skills": ["Go", "Docker", "REST API"]
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

	analysis := &JobAnalysis{
		Keywords:         []string{"Go", "Docker"},
		Responsibilities: []string{"Build backend APIs"},
		Requirements:     []string{"Experience with Go"},
	}

	rewriter := NewRewriter(c)
	adapted, err := rewriter.Rewrite(context.Background(), resume, analysis)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if adapted.Name != resume.Name {
		t.Errorf("expected name to be preserved, got %q", adapted.Name)
	}
	if adapted.Summary != "Adapted Summary highlighting Go and microservices experience." {
		t.Errorf("unexpected adapted summary: %q", adapted.Summary)
	}
	if len(adapted.Experience) != 1 {
		t.Fatalf("expected 1 work experience, got %d", len(adapted.Experience))
	}
	exp := adapted.Experience[0]
	if exp.Company != "Acme Corp" || exp.Title != "Software Engineer" {
		t.Errorf("company/title altered: %s - %s", exp.Company, exp.Title)
	}
	if len(exp.Achievements) != 2 || exp.Achievements[0] != "Developed Go microservices" {
		t.Errorf("unexpected achievements: %v", exp.Achievements)
	}
	if len(adapted.Skills) != 3 || adapted.Skills[0] != "Go" {
		t.Errorf("unexpected skills: %v", adapted.Skills)
	}
}

func TestRewriter_Rewrite_NilResume(t *testing.T) {
	t.Setenv("GEMINI_API_KEY", "test-key")
	c, err := NewClient(context.Background(), withFactory(fakeClientFactory))
	if err != nil {
		t.Fatalf("unexpected error creating client: %v", err)
	}

	analysis := &JobAnalysis{
		Keywords: []string{"Go"},
	}

	rewriter := NewRewriter(c)
	_, err = rewriter.Rewrite(context.Background(), nil, analysis)
	if err == nil {
		t.Fatal("expected error when resume is nil, got nil")
	}

	want := "resume cannot be nil"
	if err.Error() != want {
		t.Fatalf("unexpected error: got %q, want %q", err.Error(), want)
	}
}

func TestRewriter_Rewrite_NilAnalysis(t *testing.T) {
	t.Setenv("GEMINI_API_KEY", "test-key")
	c, err := NewClient(context.Background(), withFactory(fakeClientFactory))
	if err != nil {
		t.Fatalf("unexpected error creating client: %v", err)
	}

	resume := &parser.Resume{
		Name: "Jane Doe",
	}

	rewriter := NewRewriter(c)
	_, err = rewriter.Rewrite(context.Background(), resume, nil)
	if err == nil {
		t.Fatal("expected error when job analysis is nil, got nil")
	}

	want := "job analysis cannot be nil"
	if err.Error() != want {
		t.Fatalf("unexpected error: got %q, want %q", err.Error(), want)
	}
}

func TestRewriter_Rewrite_APIError(t *testing.T) {
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

	resume := &parser.Resume{
		Name: "Jane Doe",
	}
	analysis := &JobAnalysis{
		Keywords: []string{"Go"},
	}

	rewriter := NewRewriter(c)
	_, err = rewriter.Rewrite(context.Background(), resume, analysis)
	if err == nil {
		t.Fatal("expected error from API failure, got nil")
	}

	if !strings.Contains(err.Error(), "failed to generate content") {
		t.Fatalf("unexpected error message: %v", err)
	}
}
