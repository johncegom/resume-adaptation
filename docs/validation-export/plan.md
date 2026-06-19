# Plan: Agent Validation, Feedback & resume.lol Export (RAFT0005)

## Problem Statement
Implement validation heuristics, terminal diff display, and resume.lol Markdown generator.

## Goals
- Develop auditing agent interface.
- Implement side-by-side terminal diff rendering.
- Write resume.lol Markdown serializer.

## Non-Goals
- Generating non-standard PDF formats.

## Acceptance Criteria
- [ ] Export tests successfully write `.md` file structures compatible with standard Markdown.

## Related Tasks
- **Block by:** `RAFT0004`
- **Related:** `RAFT0001`

## Success Metrics
- Fully unit-tested Markdown generator.

## Constraints & Dependencies
- Safe file write routines.

## Validation
- Validate via unit tests asserting structure and sections of the exported Markdown resume.

## Subtasks & Implementation Plan
- [ ] `[CONFIRMED]` Write unit test cases for the Markdown generator.
- [ ] `[CONFIRMED]` Implement the Auditor Agent checks (Match & Preservation checks).
- [ ] `[CONFIRMED]` Implement terminal side-by-side rendering using Lipgloss.
- [ ] `[CONFIRMED]` Write exporter module that saves `.md` file to target directory.
