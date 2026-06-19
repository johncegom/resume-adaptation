# Plan: CLI Core Foundation (RAFT0001)

## Problem Statement
Establish a modular Go command-line interface that allows flags and interactive prompts.

## Goals
- Add Cobra commands scaffolding.
- Integrate Charmbracelet CLI helper styling.

## Non-Goals
- Adding AI API client connections.

## Acceptance Criteria
- [ ] Cobra framework successfully runs and responds to flags.
- [ ] Charmbracelet UI elements display as designed.

## Related Tasks
- **Block by:** None
- **Related:** `RAFT0002`

## Success Metrics
- Successfully builds and passes CLI smoke tests.

## Constraints & Dependencies
- Must run cleanly inside standard Windows/macOS/Linux terminal shells.

## Validation
- Validate execution by running `go test` on CLI arguments parsing logic.

## Subtasks & Implementation Plan
- [ ] `[CONFIRMED]` Propose and write placeholder unit tests for CLI argument routing.
- [ ] `[CONFIRMED]` Setup package `cmd/cli/` structure with Cobra commands.
- [ ] `[CONFIRMED]` Implement Charmbracelet `huh` prompt flow for JD and Resume path collection.
