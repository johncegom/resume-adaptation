package main

import (
	"github.com/spf13/cobra"
)

// NewRootCommand creates a fresh Cobra command tree for execution.
func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "resume-adaptation",
		Short:         "A CLI tool to dynamically adapt resumes to job descriptions",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.AddCommand(NewAdaptCommand())
	return cmd
}
