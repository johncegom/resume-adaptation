package adaptation

import (
	"context"
	"encoding/json"
	"fmt"

	"google.golang.org/genai"
)

// JobAnalysis represents the structured output of the job description analysis.
type JobAnalysis struct {
	Keywords         []string `json:"keywords"`
	Responsibilities []string `json:"responsibilities"`
	Requirements     []string `json:"requirements"`
}

var jobAnalysisSchema = &genai.Schema{
	Type: genai.TypeObject,
	Properties: map[string]*genai.Schema{
		"keywords": {
			Type:        genai.TypeArray,
			Items:       &genai.Schema{Type: genai.TypeString},
			Description: "Core technical keywords, skills, tool names, libraries, framework names, and programming languages mentioned in the job description",
		},
		"responsibilities": {
			Type:        genai.TypeArray,
			Items:       &genai.Schema{Type: genai.TypeString},
			Description: "Key responsibilities, tasks, duties, and deliverables expected in this role",
		},
		"requirements": {
			Type:        genai.TypeArray,
			Items:       &genai.Schema{Type: genai.TypeString},
			Description: "Required qualifications, experience levels, certifications, degrees, and minimum requirements for the candidate",
		},
	},
	Required: []string{"keywords", "responsibilities", "requirements"},
}

// JobAnalyzer parses and extracts key metadata from a target job description.
type JobAnalyzer struct {
	client *Client
}

// NewJobAnalyzer creates a new JobAnalyzer.
func NewJobAnalyzer(client *Client) *JobAnalyzer {
	return &JobAnalyzer{client: client}
}

// Analyze parses the job description and extracts keywords, responsibilities, and requirements.
func (a *JobAnalyzer) Analyze(ctx context.Context, jobDescription string) (*JobAnalysis, error) {
	if jobDescription == "" {
		return nil, fmt.Errorf("job description cannot be empty")
	}

	prompt := fmt.Sprintf(`Analyze the following job description and extract:
1. Core technical keywords, programming languages, technologies, and skills.
2. Main responsibilities and daily duties of the role.
3. Minimum and preferred requirements, experience, and qualifications.

Job Description:
%s`, jobDescription)

	contents := []*genai.Content{
		{Parts: []*genai.Part{{Text: prompt}}},
	}

	config := &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
		ResponseSchema:   jobAnalysisSchema,
	}

	resp, err := a.client.GenerateContent(ctx, contents, config)
	if err != nil {
		return nil, fmt.Errorf("failed to generate content: %w", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("received empty response from Gemini API")
	}

	responseText := resp.Candidates[0].Content.Parts[0].Text
	var analysis JobAnalysis
	if err := json.Unmarshal([]byte(responseText), &analysis); err != nil {
		return nil, fmt.Errorf("failed to unmarshal job analysis JSON: %w (raw response: %s)", err, responseText)
	}

	return &analysis, nil
}
