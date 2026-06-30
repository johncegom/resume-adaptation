package adaptation

import (
	"context"
	"fmt"

	"github.com/johncegom/resume-adaptation/internal/parser"
)

// Service coordinates the job description analysis and resume rewrite steps.
type Service struct {
	analyzer *JobAnalyzer
	rewriter *Rewriter
}

// NewService creates a new adaptation Service.
func NewService(analyzer *JobAnalyzer, rewriter *Rewriter) *Service {
	return &Service{
		analyzer: analyzer,
		rewriter: rewriter,
	}
}

// Adapt runs the full job description analysis and resume rewriting workflow.
func (s *Service) Adapt(ctx context.Context, resume *parser.Resume, jobDescription string) (*parser.Resume, error) {
	if resume == nil {
		return nil, fmt.Errorf("resume cannot be nil")
	}
	if jobDescription == "" {
		return nil, fmt.Errorf("job description cannot be empty")
	}

	analysis, err := s.analyzer.Analyze(ctx, jobDescription)
	if err != nil {
		return nil, fmt.Errorf("failed to analyze job description: %w", err)
	}

	adapted, err := s.rewriter.Rewrite(ctx, resume, analysis)
	if err != nil {
		return nil, fmt.Errorf("failed to rewrite resume: %w", err)
	}

	return adapted, nil
}
