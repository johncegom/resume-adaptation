package adaptation

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"

	"github.com/johncegom/resume-adaptation/internal/parser"
	"google.golang.org/genai"
)

//go:embed prompts/resume_rewrite_prompt.md
var resumeRewritePromptTemplate string

// Rewriter adapts resume experience segments to match the target job keywords.
type Rewriter struct {
	client *Client
}

var adaptedResumeSchema = &genai.Schema{
	Type: genai.TypeObject,
	Properties: map[string]*genai.Schema{
		"summary": {
			Type:        genai.TypeString,
			Description: "The adapted overall professional summary/objective statement, highlighting matching keywords and alignment to the role",
		},
		"experience": {
			Type: genai.TypeArray,
			Items: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"company": {Type: genai.TypeString, Description: "Must match the original company name exactly"},
					"title":   {Type: genai.TypeString, Description: "Must match the original job title exactly"},
					"achievements": {
						Type:        genai.TypeArray,
						Items:       &genai.Schema{Type: genai.TypeString},
						Description: "The adapted list of achievements and responsibilities, rephrased to align with job keywords while preserving exact truthfulness",
					},
				},
				Required: []string{"company", "title", "achievements"},
			},
			Description: "The list of adapted work experiences, in the exact same order as the original list",
		},
		"skills": {
			Type:        genai.TypeArray,
			Items:       &genai.Schema{Type: genai.TypeString},
			Description: "The adapted list of skills, reordered to prioritize matching keywords and filtered to only include those the candidate actually possesses",
		},
	},
	Required: []string{"summary", "experience", "skills"},
}

// NewRewriter creates a new Rewriter.
func NewRewriter(client *Client) *Rewriter {
	return &Rewriter{client: client}
}

// Rewrite adapts the summary, achievements and skills of the resume.
func (r *Rewriter) Rewrite(ctx context.Context, resume *parser.Resume, analysis *JobAnalysis) (*parser.Resume, error) {
	if resume == nil {
		return nil, fmt.Errorf("resume cannot be nil")
	}
	if analysis == nil {
		return nil, fmt.Errorf("job analysis cannot be nil")
	}

	origExpBytes, err := json.Marshal(resume.Experience)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal original experience: %w", err)
	}

	origSkillsBytes, err := json.Marshal(resume.Skills)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal original skills: %w", err)
	}

	prompt := fmt.Sprintf(resumeRewritePromptTemplate,
		analysis.Keywords, analysis.Responsibilities, analysis.Requirements,
		resume.Summary, string(origExpBytes), string(origSkillsBytes))

	contents := []*genai.Content{
		{Parts: []*genai.Part{{Text: prompt}}},
	}

	config := &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
		ResponseSchema:   adaptedResumeSchema,
	}

	resp, err := r.client.GenerateContent(ctx, contents, config)
	if err != nil {
		return nil, fmt.Errorf("failed to generate content: %w", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("received empty response from Gemini API")
	}

	responseText := resp.Candidates[0].Content.Parts[0].Text
	var respData struct {
		Summary    string `json:"summary"`
		Experience []struct {
			Company      string   `json:"company"`
			Title        string   `json:"title"`
			Achievements []string `json:"achievements"`
		} `json:"experience"`
		Skills []string `json:"skills"`
	}
	if err := json.Unmarshal([]byte(responseText), &respData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal adapted resume JSON: %w (raw response: %s)", err, responseText)
	}

	// Create a new Resume copy to avoid side-effects (regression guardrails)
	adapted := &parser.Resume{
		Name:       resume.Name,
		Summary:    respData.Summary,
		RawContent: resume.RawContent,
		Skills:     respData.Skills,
		Education:  resume.Education,
		Projects:   resume.Projects,
		Experience: make([]parser.WorkExperience, len(resume.Experience)),
	}

	for i, orig := range resume.Experience {
		adapted.Experience[i] = parser.WorkExperience{
			Company:      orig.Company,
			Title:        orig.Title,
			StartDate:    orig.StartDate,
			EndDate:      orig.EndDate,
			Achievements: orig.Achievements, // Default to original if not adapted
		}

		for _, adExp := range respData.Experience {
			if adExp.Company == orig.Company && adExp.Title == orig.Title {
				adapted.Experience[i].Achievements = adExp.Achievements
				break
			}
		}
	}

	return adapted, nil
}
