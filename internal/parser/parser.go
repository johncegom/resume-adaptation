package parser

import (
	"context"
	"encoding/json"
	"fmt"

	"google.golang.org/genai"
)

// PackageName returns the name of this package.
// It serves as a build-verification sentinel confirming
// the package structure is correctly configured.
func PackageName() string {
	return "parser"
}

// PDFParser is a resume parser that uses the Gemini API.
type PDFParser struct {
	client *genai.Client
	model  string
}

// NewPDFParser creates a new PDFParser with the given client and model.
// If model is empty, "gemini-2.5-flash" is used as the default.
func NewPDFParser(client *genai.Client, model string) *PDFParser {
	if model == "" {
		model = "gemini-2.5-flash"
	}
	return &PDFParser{
		client: client,
		model:  model,
	}
}

// resumeSchema defines the structured output schema for a candidate's resume.
var resumeSchema = &genai.Schema{
	Type: genai.TypeObject,
	Properties: map[string]*genai.Schema{
		"name":    {Type: genai.TypeString, Description: "Candidate's full name"},
		"summary": {Type: genai.TypeString, Description: "Professional summary or objective statement"},
		"experience": {
			Type: genai.TypeArray,
			Items: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"company":    {Type: genai.TypeString, Description: "Company or organization name"},
					"title":      {Type: genai.TypeString, Description: "Job title or role"},
					"start_date": {Type: genai.TypeString, Description: "Start date of the experience, e.g. YYYY-MM"},
					"end_date":   {Type: genai.TypeString, Description: "End date or 'Present'"},
					"achievements": {
						Type: genai.TypeArray,
						Items: &genai.Schema{
							Type: genai.TypeString,
						},
						Description: "Key achievements, responsibilities, or impact in this role",
					},
				},
				Required: []string{"company", "title"},
			},
			Description: "List of candidate's professional work experiences",
		},
		"education": {
			Type: genai.TypeArray,
			Items: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"institution": {Type: genai.TypeString, Description: "Name of the school, university, or academy"},
					"degree":      {Type: genai.TypeString, Description: "Degree earned, e.g. Bachelor of Science"},
					"major":       {Type: genai.TypeString, Description: "Field of study"},
					"start_date":  {Type: genai.TypeString, Description: "Start date, e.g. YYYY-MM"},
					"end_date":    {Type: genai.TypeString, Description: "Graduation or end date, e.g. YYYY-MM"},
				},
				Required: []string{"institution"},
			},
			Description: "Candidate's academic education history",
		},
		"projects": {
			Type: genai.TypeArray,
			Items: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"name":        {Type: genai.TypeString, Description: "Project name"},
					"role":        {Type: genai.TypeString, Description: "Candidate's role in the project"},
					"tech_stack":  {Type: genai.TypeArray, Items: &genai.Schema{Type: genai.TypeString}, Description: "Technologies used in the project"},
					"description": {Type: genai.TypeString, Description: "Brief description of the project and results"},
				},
				Required: []string{"name"},
			},
			Description: "List of notable projects",
		},
		"skills": {
			Type:        genai.TypeArray,
			Items:       &genai.Schema{Type: genai.TypeString},
			Description: "Candidate's key skills, technical or professional capabilities",
		},
	},
	Required: []string{"name", "summary", "experience", "education", "projects", "skills"},
}

// Parse uploads/sends PDF bytes to the Gemini API and requests a structured response
// matching the Resume JSON schema.
func (p *PDFParser) Parse(ctx context.Context, pdfBytes []byte) (*Resume, error) {
	if p.client == nil {
		return nil, fmt.Errorf("gemini client is not initialized")
	}
	if len(pdfBytes) == 0 {
		return nil, fmt.Errorf("pdf content is empty")
	}

	parts := []*genai.Part{
		{
			InlineData: &genai.Blob{
				Data:     pdfBytes,
				MIMEType: "application/pdf",
			},
		},
		{
			Text: "Extract the candidate's resume information into the requested structured format.",
		},
	}

	config := &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
		ResponseSchema:   resumeSchema,
	}

	resp, err := p.client.Models.GenerateContent(
		ctx,
		p.model,
		[]*genai.Content{{Parts: parts}},
		config,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to generate content: %w", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("received empty response from Gemini API")
	}

	responseText := resp.Candidates[0].Content.Parts[0].Text
	var resume Resume
	if err := json.Unmarshal([]byte(responseText), &resume); err != nil {
		return nil, fmt.Errorf("failed to unmarshal structured JSON: %w (raw response: %s)", err, responseText)
	}

	return &resume, nil
}
