//go:build integration

package adaptation

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/johncegom/resume-adaptation/internal/parser"
)

func TestService_Adapt_LiveIntegration(t *testing.T) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		t.Skip("skipping live integration test: GEMINI_API_KEY environment variable not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := NewClient(ctx)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	analyzer := NewJobAnalyzer(client)
	rewriter := NewRewriter(client)
	service := NewService(analyzer, rewriter)

	resume := &parser.Resume{
		Name:    "Jane Doe",
		Summary: "Experienced Software Engineer with a focus on building web applications using Python and Django.",
		Experience: []parser.WorkExperience{
			{
				Company:      "Acme Corp",
				Title:        "Backend Developer",
				StartDate:    "2022-01",
				EndDate:      "Present",
				Achievements: []string{"Designed and maintained Django-based backend APIs", "Optimized database queries"},
			},
		},
		Skills: []string{"Python", "Django", "PostgreSQL"},
	}

	jobDescription := `
	We are looking for a Go Backend Engineer.
	Key responsibilities include building microservices in Go, containerizing applications using Docker, and deploying to Kubernetes.
	Requirements:
	- Strong experience with Go or Python with interest to learn Go.
	- Experience with containerization.
	`

	adapted, err := service.Adapt(ctx, resume, jobDescription)
	if err != nil {
		t.Fatalf("failed to adapt resume: %v", err)
	}

	if adapted == nil {
		t.Fatal("expected non-nil adapted resume")
	}

	// Verify that key structure is intact
	if adapted.Name != resume.Name {
		t.Errorf("expected name to be preserved, got %q", adapted.Name)
	}
	if len(adapted.Experience) != 1 {
		t.Fatalf("expected 1 experience, got %d", len(adapted.Experience))
	}
	if adapted.Summary == "" {
		t.Error("expected non-empty adapted summary")
	}
	if len(adapted.Skills) == 0 {
		t.Error("expected non-empty adapted skills list")
	}
}
