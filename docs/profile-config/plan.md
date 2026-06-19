# Plan: Profile & Setting Configurations (RAFT0002)

## Problem Statement
Add cross-platform settings serialization for CLI credentials and paths.

## Goals
- Support JSON serialization of configurations.
- Map Cobra commands to set/get config fields.

## Non-Goals
- Adding complex multi-profile database systems.

## Acceptance Criteria
- [ ] Read/Write tests pass for temporary JSON configuration files.

## Related Tasks
- **Block by:** `RAFT0001`
- **Related:** `RAFT0004`

## Success Metrics
- Successfully reads config values during initialization in CLI commands.

## Constraints & Dependencies
- Safe JSON encoding/decoding, checking directory existence before writing.

## Validation
- Unit tests proving config storage, loading, and field updates.

## Subtasks & Implementation Plan
- [ ] `[CONFIRMED]` Write unit test cases for config read/write operations with temp files.
- [ ] `[CONFIRMED]` Implement config package with `LoadConfig()` and `SaveConfig()`.
- [ ] `[CONFIRMED]` Hook up the config utility to the CLI Cobra commands in `cmd/cli/main.go`.
