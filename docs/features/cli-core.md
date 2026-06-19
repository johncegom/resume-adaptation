# Epic: CLI Core & Interactive Shell

## What
Bootstraps the CLI command structure using Cobra and sets up interactive user prompts using Charmbracelet Huh/Lipgloss libraries.

## Why
Ensures a premium developer and candidate user experience with self-documenting commands and clean prompt inputs.

## How
- Set up a main command `resume-adapt` with subcommands like `run`, `config`.
- Integrate Charmbracelet Huh for step-by-step inputs (if flags aren't passed directly).
- Define default layout structures under the `cmd/` package.
