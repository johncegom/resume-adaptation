# Tasks - Interactive UI Adaptation Wizard

This document tracks all implementation tasks for the TUI wizard feature.

## Task List

- [x] RAFT0011: Create TUI layouts and styling variables
  - **Acceptance Criteria**: Lipgloss styling parameters, custom colors, and layouts are defined in `internal/ui/styles.go`.
  - **Dependencies**: None

- [x] RAFT0012: Implement the interactive Huh Form input sequence
  - **Acceptance Criteria**: Form captures resume path, job description inputs, and API configurations, leveraging Viper flag/environment variable bindings.
  - **Dependencies**: RAFT0011

- [x] RAFT0013: Implement process spinner and progress loader
  - **Acceptance Criteria**: A responsive Bubble Tea loader spinner displays while mock processes execute, rendering strictly to Stderr.
  - **Dependencies**: RAFT0012

- [x] RAFT0014: Implement review and manual editing terminal view
  - **Acceptance Criteria**: Renders the adapted resume structure, allowing keyboard scrolling and editing of fields.
  - **Dependencies**: RAFT0013

- [x] RAFT0015: Add Cobra CLI `adapt` command router
  - **Acceptance Criteria**: Registers the TUI runner under Cobra command `go run cmd/cli/main.go adapt`. Uses `RunE`, sets `SilenceUsage: true`/`SilenceErrors: true`, and propagates signal context cancellations.
  - **Dependencies**: RAFT0014
