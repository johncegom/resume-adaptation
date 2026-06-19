# Task: CLI Core Foundation (RAFT0001)

## Problem Statement
The tool needs a clean, robust, and interactive CLI entry point so users can easily pass resumes and JDs, configure preferences, and initiate adaptations.

## Goals
- Setup the Cobra CLI application framework in Go.
- Establish subcommands `run` and `config`.
- Integrate Charmbracelet libraries (`huh`, `lipgloss`) for terminal UI styling.

## Non-Goals
- Performing PDF parsing or AI adaptation in this task.
- Implementing persistent configurations.

## Acceptance Criteria
- [ ] Cobra root command is bootable.
- [ ] Subcommand `run` is defined and parses flags.
- [ ] Subcommand `config` is defined and parses flags.
- [ ] Interactive prompts using `huh` display correctly if flags are omitted.

## Related Tasks
- **Block by:** None
- **Related:** `RAFT0002` (config), `RAFT0003` (pdf-ingestion)

## Success Metrics
- 100% of CLI flags are correctly parsed.
- Smooth terminal flow without crashing on standard inputs.

## Constraints & Dependencies
- Language: Go 1.22+
- Main libraries: `github.com/spf13/cobra`, `github.com/charmbracelet/huh`

## Validation
- Verify through automated CLI tests that commands print help output and parse arguments without breaking existing behavior.

## Subtasks
- [ ] `[CONFIRMED]` Initialize Go module and setup main command routing.
- [ ] `[CONFIRMED]` Integrate Cobra root command and basic commands layout under `cmd/cli/main.go`.
- [ ] `[CONFIRMED]` Define CLI interactive step-by-step logic using `huh` for user inputs.
