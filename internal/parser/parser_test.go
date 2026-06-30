package parser_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/johncegom/resume-adaptation/internal/parser"
	"google.golang.org/genai"
)

type mockTransport struct {
	roundTrip func(*http.Request) (*http.Response, error)
}

func (m *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.roundTrip(req)
}

func TestParserPackageExists(t *testing.T) {
	// Verify the parser package is importable and the project compiles.
	// Importing parser and calling a sentinel function confirms linkage.
	got := parser.PackageName()
	if got != "parser" {
		t.Fatalf("expected package name %q, got %q", "parser", got)
	}
}

func TestGenaiDependencyAvailable(t *testing.T) {
	// Verify that the google.golang.org/genai dependency is resolvable.
	// Importing the genai package and referencing a type proves the
	// dependency is correctly added to go.mod and downloadable.
	var client *genai.Client
	if client != nil {
		t.Fatal("unexpected non-nil client")
	}
}

func TestPDFParser_Parse_NilClient(t *testing.T) {
	ctx := context.Background()
	p := parser.NewPDFParser(nil, "")
	_, err := p.Parse(ctx, []byte("some content"))
	if err == nil {
		t.Fatal("expected error with nil client, got nil")
	}
	if !strings.Contains(err.Error(), "client is not initialized") {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestPDFParser_Parse_EmptyPDF(t *testing.T) {
	ctx := context.Background()
	client, _ := genai.NewClient(ctx, &genai.ClientConfig{APIKey: "fake-key"})
	p := parser.NewPDFParser(client, "")
	_, err := p.Parse(ctx, nil)
	if err == nil {
		t.Fatal("expected error with empty PDF bytes, got nil")
	}
	if !strings.Contains(err.Error(), "pdf content is empty") {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestPDFParser_Parse_SuccessMock(t *testing.T) {
	mockResponseJSON := `{
		"name": "Jane Doe",
		"summary": "Experienced engineer.",
		"experience": [
			{
				"company": "Acme Corp",
				"title": "Software Engineer",
				"start_date": "2020-01",
				"end_date": "Present",
				"achievements": ["Developed features"]
			}
		],
		"education": [
			{
				"institution": "University of Go",
				"degree": "B.S.",
				"major": "CS",
				"start_date": "2015-09",
				"end_date": "2019-06"
			}
		],
		"projects": [
			{
				"name": "Resume Adaptation",
				"role": "Lead",
				"tech_stack": ["Go", "Gemini"],
				"description": "CLI tool"
			}
		],
		"skills": ["Go", "Gemini"]
	}`

	transport := &mockTransport{
		roundTrip: func(req *http.Request) (*http.Response, error) {
			// Mock the exact response payload structure returned by the Gemini GenerateContent API.
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

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: "fake-key",
		HTTPClient: &http.Client{
			Transport: transport,
		},
	})
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	p := parser.NewPDFParser(client, "")
	resume, err := p.Parse(ctx, []byte("fake-pdf-content"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resume.Name != "Jane Doe" {
		t.Errorf("expected candidate name 'Jane Doe', got %q", resume.Name)
	}
	if len(resume.Experience) == 0 || resume.Experience[0].Company != "Acme Corp" {
		t.Errorf("unexpected or missing experience")
	}
	if len(resume.Skills) != 2 || resume.Skills[0] != "Go" {
		t.Errorf("unexpected or missing skills: %v", resume.Skills)
	}
}

func TestPDFParser_Parse_Live(t *testing.T) {
	_ = godotenv.Load(filepath.Join("..", "..", ".env"))
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		t.Skip("skipping live integration test: GEMINI_API_KEY environment variable not set")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		t.Fatalf("failed to create live Gemini client: %v", err)
	}

	// Minimal valid PDF structure
	dummyPDFBytes := []byte("%PDF-1.4\n1 0 obj\n<< /Type /Catalog /Pages 2 0 R >>\nendobj\n2 0 obj\n<< /Type /Pages /Kids [3 0 R] /Count 1 >>\nendobj\n3 0 obj\n<< /Type /Page /Parent 2 0 R /MediaBox [0 0 612 792] /Resources << >> /Contents 4 0 R >>\nendobj\n4 0 obj\n<< /Length 40 >>\nstream\nBT /F1 12 Tf 70 700 Td (John Doe Resume) Tj ET\nendstream\nendobj\nxref\n0 5\n0000000000 65535 f \n0000000009 00000 n \n0000000056 00000 n \n0000000111 00000 n \n0000000212 00000 n \ntrailer\n<< /Size 5 /Root 1 0 R >>\nstartxref\n303\n%%EOF")

	p := parser.NewPDFParser(client, "gemini-2.5-flash")
	resume, err := p.Parse(ctx, dummyPDFBytes)
	if err != nil {
		t.Fatalf("failed to parse PDF: %v", err)
	}
	if resume.Name == "" {
		t.Error("expected candidate name to be extracted and non-empty")
	}
}
