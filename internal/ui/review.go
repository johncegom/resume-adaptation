package ui

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/johncegom/resume-adaptation/internal/parser"
)

// RenderResumePreview returns a premium styled string preview of the Resume.
func RenderResumePreview(resume *parser.Resume) string {
	var b strings.Builder

	b.WriteString(StyleTitle.Render("ADAPTED RESUME REVIEW") + "\n\n")
	b.WriteString(fmt.Sprintf("%s %s\n\n", StyleHighlight.Render("Candidate Name:"), resume.Name))

	if resume.Summary != "" {
		b.WriteString(StyleSubTitle.Render("Professional Summary:") + "\n")
		b.WriteString(StyleText.Render(resume.Summary) + "\n\n")
	}

	if len(resume.Skills) > 0 {
		b.WriteString(StyleSubTitle.Render("Skills:") + "\n")
		b.WriteString(StyleText.Render(strings.Join(resume.Skills, ", ")) + "\n\n")
	}

	if len(resume.Experience) > 0 {
		b.WriteString(StyleSubTitle.Render("Professional Experience:") + "\n")
		for _, exp := range resume.Experience {
			b.WriteString(fmt.Sprintf("%s %s (%s - %s)\n", StyleHighlight.Render("• "+exp.Company), exp.Title, exp.StartDate, exp.EndDate))
			for _, ach := range exp.Achievements {
				b.WriteString(fmt.Sprintf("  - %s\n", ach))
			}
		}
		b.WriteString("\n")
	}

	return StyleBox.Render(b.String())
}

// RunReviewWizard displays the preview and allows the candidate to edit fields interactively.
// It loops until the user confirms they want to save and finish.
func RunReviewWizard(resume *parser.Resume) (*parser.Resume, error) {
	for {
		// 1. Show the current preview
		previewStr := RenderResumePreview(resume)
		fmt.Fprint(os.Stderr, previewStr)

		// 2. Ask if they want to edit anything
		var editChoice string
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Review Options").
					Description("Select a section to edit, or finish and export the resume").
					Options(
						huh.NewOption("Edit Summary", "summary"),
						huh.NewOption("Edit Skills", "skills"),
						huh.NewOption("Edit Experience Achievements", "experience"),
						huh.NewOption("Finish and Export", "finish"),
					).
					Value(&editChoice),
			),
		).WithProgramOptions(tea.WithOutput(os.Stderr))

		if err := form.Run(); err != nil {
			return nil, err
		}

		if editChoice == "finish" {
			break
		}

		switch editChoice {
		case "summary":
			var newSummary string
			formSummary := huh.NewForm(
				huh.NewGroup(
					huh.NewText().
						Title("Edit Summary").
						Value(&newSummary).
						Placeholder(resume.Summary),
				),
			).WithProgramOptions(tea.WithOutput(os.Stderr))
			// Pre-fill
			newSummary = resume.Summary
			if err := formSummary.Run(); err == nil && newSummary != "" {
				resume.Summary = newSummary
			}

		case "skills":
			var skillsInput string
			formSkills := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Title("Edit Skills (comma separated)").
						Value(&skillsInput).
						Placeholder(strings.Join(resume.Skills, ", ")),
				),
			).WithProgramOptions(tea.WithOutput(os.Stderr))
			// Pre-fill
			skillsInput = strings.Join(resume.Skills, ", ")
			if err := formSkills.Run(); err == nil {
				// Parse skills input
				parts := strings.Split(skillsInput, ",")
				var newSkills []string
				for _, part := range parts {
					trimmed := strings.TrimSpace(part)
					if trimmed != "" {
						newSkills = append(newSkills, trimmed)
					}
				}
				if len(newSkills) > 0 {
					resume.Skills = newSkills
				}
			}

		case "experience":
			if len(resume.Experience) == 0 {
				fmt.Fprintln(os.Stderr, StyleError.Render("No experiences to edit!"))
				continue
			}

			// Let user select which company's achievements to edit
			var compOptions []huh.Option[int]
			for idx, exp := range resume.Experience {
				compOptions = append(compOptions, huh.NewOption(exp.Company+" - "+exp.Title, idx))
			}

			var selectedIdx int
			formComp := huh.NewForm(
				huh.NewGroup(
					huh.NewSelect[int]().
						Title("Select Role to Edit").
						Options(compOptions...).
						Value(&selectedIdx),
				),
			).WithProgramOptions(tea.WithOutput(os.Stderr))

			if err := formComp.Run(); err != nil {
				continue
			}

			exp := &resume.Experience[selectedIdx]

			// Edit achievements as a single text block (newline-separated)
			achievementsText := strings.Join(exp.Achievements, "\n")
			var newAchText string
			formAch := huh.NewForm(
				huh.NewGroup(
					huh.NewText().
						Title(fmt.Sprintf("Edit Achievements for %s", exp.Company)).
						Description("One achievement bullet per line").
						Value(&newAchText),
				),
			).WithProgramOptions(tea.WithOutput(os.Stderr))

			newAchText = achievementsText
			if err := formAch.Run(); err == nil {
				lines := strings.Split(newAchText, "\n")
				var newAchs []string
				for _, line := range lines {
					trimmed := strings.TrimSpace(line)
					if trimmed != "" {
						newAchs = append(newAchs, trimmed)
					}
				}
				exp.Achievements = newAchs
			}
		}
	}

	return resume, nil
}
