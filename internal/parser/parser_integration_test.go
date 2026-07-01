//go:build integration

package parser_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/johncegom/resume-adaptation/internal/parser"
	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

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
