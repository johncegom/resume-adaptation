package ui

import (
	"strings"
	"testing"

	"github.com/johncegom/resume-adaptation/internal/parser"
)

func TestRenderResumePreview(t *testing.T) {
	resume := &parser.Resume{
		Name:    "Alice Smith",
		Summary: "Talented developer.",
		Skills:  []string{"Go", "Docker"},
		Experience: []parser.WorkExperience{
			{
				Company:   "Tech Co",
				Title:     "Engineer",
				StartDate: "2020",
				EndDate:   "2023",
				Achievements: []string{
					"Wrote Go APIs",
				},
			},
		},
	}

	preview := RenderResumePreview(resume)
	if preview == "" {
		t.Fatal("expected non-empty preview rendering")
	}

	if !strings.Contains(preview, "Alice Smith") {
		t.Errorf("expected preview to contain candidate name, got %q", preview)
	}
	if !strings.Contains(preview, "Talented developer") {
		t.Errorf("expected preview to contain summary, got %q", preview)
	}
	if !strings.Contains(preview, "Go, Docker") {
		t.Errorf("expected preview to contain skills, got %q", preview)
	}
	if !strings.Contains(preview, "Tech Co") {
		t.Errorf("expected preview to contain experience, got %q", preview)
	}
}
