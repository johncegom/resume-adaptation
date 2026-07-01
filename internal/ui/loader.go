package ui

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type errMsg error
type quitMsg struct{}

type spinnerModel struct {
	spinner  spinner.Model
	title    string
	taskFn   func() error
	err      error
	quitting bool
}

func initialSpinnerModel(title string, taskFn func() error) spinnerModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(ColorAccent)
	return spinnerModel{
		spinner: s,
		title:   title,
		taskFn:  taskFn,
	}
}

func (m spinnerModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		func() tea.Msg {
			if err := m.taskFn(); err != nil {
				return errMsg(err)
			}
			return quitMsg{}
		},
	)
}

func (m spinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case quitMsg:
		m.quitting = true
		return m, tea.Quit

	case errMsg:
		m.err = msg
		m.quitting = true
		return m, tea.Quit

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	default:
		return m, nil
	}
}

func (m spinnerModel) View() string {
	if m.err != nil {
		return StyleError.Render(fmt.Sprintf("Error: %v\n", m.err))
	}
	if m.quitting {
		return ""
	}
	return fmt.Sprintf("\n %s %s...\n\n", m.spinner.View(), StyleText.Render(m.title))
}

// RunLoadingSpinner runs a Bubble Tea program displaying a spinner while executing the taskFn.
// It outputs strictly to Stderr to keep Stdout clear.
func RunLoadingSpinner(ctx context.Context, title string, taskFn func() error) error {
	m := initialSpinnerModel(title, taskFn)
	p := tea.NewProgram(m, tea.WithOutput(os.Stderr), tea.WithContext(ctx))
	runModel, err := p.Run()
	if err != nil {
		return err
	}
	if finalModel, ok := runModel.(spinnerModel); ok && finalModel.err != nil {
		return finalModel.err
	}
	return nil
}
