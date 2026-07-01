package ui

import (
	"errors"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

// FormInputs captures the outputs from the interactive form wizard.
type FormInputs struct {
	ResumePath   string
	JobDescType  string // "file" or "text"
	JobDescPath  string
	JobDescText  string
	OutputPath   string
	GeminiAPIKey string
}

// RunInputWizard executes the Huh TUI input wizard, collecting parameters for the adaptation.
// It renders strictly to Stderr to preserve Stdout for pipeable output.
func RunInputWizard(apiKeyEnv string) (*FormInputs, error) {
	inputs := &FormInputs{}

	// Step 1: Gather Resume Path and Job Description type
	var jdType string
	form1 := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Candidate Resume Path").
				Description("Path to the plain text, markdown, or PDF resume").
				Value(&inputs.ResumePath).
				Validate(ValidateInputPath),

			huh.NewSelect[string]().
				Title("Job Description Source").
				Description("How would you like to provide the target job description?").
				Options(
					huh.NewOption("Provide path to a file", "file"),
					huh.NewOption("Paste description text directly", "text"),
				).
				Value(&jdType),
		),
	).WithProgramOptions(tea.WithOutput(os.Stderr))

	if err := form1.Run(); err != nil {
		return nil, err
	}

	inputs.JobDescType = jdType

	// Step 2: Gather Job Description content based on choice
	var form2 *huh.Form
	if jdType == "file" {
		form2 = huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Job Description File Path").
					Description("Path to the job description file").
					Value(&inputs.JobDescPath).
					Validate(ValidateInputPath),
			),
		).WithProgramOptions(tea.WithOutput(os.Stderr))
	} else {
		form2 = huh.NewForm(
			huh.NewGroup(
				huh.NewText().
					Title("Job Description Text").
					Description("Paste the raw text of the job description").
					Value(&inputs.JobDescText).
					Placeholder("Requirements:\n- 3+ years Go experience..."),
			),
		).WithProgramOptions(tea.WithOutput(os.Stderr))
	}

	if err := form2.Run(); err != nil {
		return nil, err
	}

	// Step 3: Gather Output Path and API Key (if not provided in environment)
	var fields []huh.Field
	fields = append(fields, huh.NewInput().
		Title("Output Path").
		Description("Path to save the adapted resume (markdown format)").
		Value(&inputs.OutputPath).
		Placeholder("adapted_resume.md"))

	if apiKeyEnv == "" {
		var confirmKey bool
		formConfirm := huh.NewForm(
			huh.NewGroup(
				huh.NewConfirm().
					Title("Gemini API Key Missing").
					Description("No GEMINI_API_KEY environment variable was found. Would you like to enter it now?").
					Value(&confirmKey),
			),
		).WithProgramOptions(tea.WithOutput(os.Stderr))

		if err := formConfirm.Run(); err != nil {
			return nil, err
		}

		if confirmKey {
			formKey := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Title("Gemini API Key").
						Description("Enter your Google Gemini API Key").
						Value(&inputs.GeminiAPIKey).
						Password(true).
						Validate(func(s string) error {
							if s == "" {
								return errors.New("API key cannot be empty")
							}
							return nil
						}),
				),
			).WithProgramOptions(tea.WithOutput(os.Stderr))

			if err := formKey.Run(); err != nil {
				return nil, err
			}
		}
	}

	// Apply default output path if not specified
	if inputs.OutputPath == "" {
		inputs.OutputPath = "adapted_resume.md"
	}

	return inputs, nil
}
