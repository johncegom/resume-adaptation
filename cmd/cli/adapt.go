package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/johncegom/resume-adaptation/internal/adaptation"
	"github.com/johncegom/resume-adaptation/internal/parser"
	"github.com/johncegom/resume-adaptation/internal/ui"
	"github.com/spf13/cobra"
	"google.golang.org/genai"
)

// NewAdaptCommand creates a new adapt command subcommand.
func NewAdaptCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "adapt",
		Short:         "Interactively adapt your resume to a job description",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			if ctx == nil {
				ctx = context.Background()
			}

			// Gather environment API Key
			apiKeyEnv := os.Getenv("GEMINI_API_KEY")

			// Run interactive input wizard (renders strictly to Stderr)
			inputs, err := ui.RunInputWizard(apiKeyEnv)
			if err != nil {
				return err
			}

			// If key was entered via TUI, set it temporarily
			apiKey := apiKeyEnv
			if apiKey == "" {
				apiKey = inputs.GeminiAPIKey
			}

			if apiKey == "" {
				return fmt.Errorf("GEMINI_API_KEY is required but not set")
			}

			// Parse Resume
			var originalResume *parser.Resume
			parseTask := func() error {
				ext := strings.ToLower(filepath.Ext(inputs.ResumePath))
				if ext == ".pdf" {
					genaiClient, err := genai.NewClient(ctx, &genai.ClientConfig{
						APIKey: apiKey,
					})
					if err != nil {
						return fmt.Errorf("failed to initialize Gemini Client: %w", err)
					}
					pdfBytes, err := os.ReadFile(inputs.ResumePath)
					if err != nil {
						return fmt.Errorf("failed to read PDF file: %w", err)
					}
					pdfParser := parser.NewPDFParser(genaiClient, "")
					parsed, err := pdfParser.Parse(ctx, pdfBytes)
					if err != nil {
						return fmt.Errorf("failed to parse PDF: %w", err)
					}
					originalResume = parsed
				} else if ext == ".txt" || ext == ".md" {
					parsed, err := parser.ReadPlaintextFile(inputs.ResumePath)
					if err != nil {
						return fmt.Errorf("failed to read plaintext resume: %w", err)
					}
					originalResume = parsed
				} else {
					return fmt.Errorf("unsupported resume file extension: %s (expected .pdf, .txt, .md)", ext)
				}
				return nil
			}

			err = ui.RunLoadingSpinner(ctx, "Parsing candidate resume", parseTask)
			if err != nil {
				return err
			}

			// Read Job Description
			var jdText string
			if inputs.JobDescType == "file" {
				content, err := os.ReadFile(inputs.JobDescPath)
				if err != nil {
					return fmt.Errorf("failed to read job description file: %w", err)
				}
				jdText = string(content)
			} else {
				jdText = inputs.JobDescText
			}

			if strings.TrimSpace(jdText) == "" {
				return fmt.Errorf("job description cannot be empty")
			}

			// Run Adaptation
			var adaptedResume *parser.Resume
			adaptTask := func() error {
				// Temporary inject the loaded key if it wasn't pre-set in the environment
				if os.Getenv("GEMINI_API_KEY") == "" {
					os.Setenv("GEMINI_API_KEY", apiKey)
					defer os.Unsetenv("GEMINI_API_KEY")
				}

				client, err := adaptation.NewClient(ctx)
				if err != nil {
					return fmt.Errorf("failed to create client: %w", err)
				}

				analyzer := adaptation.NewJobAnalyzer(client)
				rewriter := adaptation.NewRewriter(client)
				service := adaptation.NewService(analyzer, rewriter)

				adapted, err := service.Adapt(ctx, originalResume, jdText)
				if err != nil {
					return err
				}
				adaptedResume = adapted
				return nil
			}

			err = ui.RunLoadingSpinner(ctx, "Analyzing job description and adapting resume", adaptTask)
			if err != nil {
				return err
			}

			// Run Review Wizard
			reviewedResume, err := ui.RunReviewWizard(adaptedResume)
			if err != nil {
				return err
			}

			// Serialize to Markdown
			var markdownBuilder strings.Builder
			markdownBuilder.WriteString(fmt.Sprintf("# %s\n\n", reviewedResume.Name))
			if reviewedResume.Summary != "" {
				markdownBuilder.WriteString(fmt.Sprintf("## Professional Summary\n%s\n\n", reviewedResume.Summary))
			}
			if len(reviewedResume.Skills) > 0 {
				markdownBuilder.WriteString(fmt.Sprintf("## Skills\n%s\n\n", strings.Join(reviewedResume.Skills, ", ")))
			}
			if len(reviewedResume.Experience) > 0 {
				markdownBuilder.WriteString("## Professional Experience\n")
				for _, exp := range reviewedResume.Experience {
					markdownBuilder.WriteString(fmt.Sprintf("### %s — %s\n", exp.Company, exp.Title))
					markdownBuilder.WriteString(fmt.Sprintf("*%s - %s*\n", exp.StartDate, exp.EndDate))
					for _, ach := range exp.Achievements {
						markdownBuilder.WriteString(fmt.Sprintf("- %s\n", ach))
					}
					markdownBuilder.WriteString("\n")
				}
			}

			finalMarkdown := markdownBuilder.String()

			// Write to file (if output path specified)
			if inputs.OutputPath != "" {
				err = os.WriteFile(inputs.OutputPath, []byte(finalMarkdown), 0644)
				if err != nil {
					return fmt.Errorf("failed to write adapted resume to output file: %w", err)
				}
				fmt.Fprintln(os.Stderr, ui.StyleSuccess.Render(fmt.Sprintf("Adapted resume saved successfully to %s", inputs.OutputPath)))
			}

			// Output clean pipeable markdown to Stdout
			fmt.Fprint(cmd.OutOrStdout(), finalMarkdown)

			return nil
		},
	}
	return cmd
}
