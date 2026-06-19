# Task: Agent Validation, Feedback & resume.lol Export (RAFT0005)

## Problem Statement
The user needs to be certain that the adapted resume is truthful, understand the changes made (with strengths/gaps analysis), and export it to a format compatible with resume.lol.

## Goals
- Build a dual-agent validation loop (JD Match Score + Preservation/Truthfulness Score).
- Generate qualitative reports (Strengths, Gaps/Improvements, Rewriting Explanations).
- Render a side-by-side terminal diff showing modifications.
- Export output to resume.lol Markdown schema.

## Non-Goals
- Editing raw PDF structures for output.

## Acceptance Criteria
- [ ] Preservation Score check flags any changes to metrics or facts.
- [ ] Terminal UI displays: Scores, Strengths list, Missing Gaps list, and a side-by-side comparison of changes.
- [ ] Exports structured Markdown file matching resume.lol format.

## Related Tasks
- **Block by:** `RAFT0004` (requires adapted resume)
- **Related:** `RAFT0001` (terminal printing)

## Success Metrics
- 100% preservation check accuracy on test metric manipulations.
- Output file is directly pasteable into resume.lol without parser syntax errors.

## Constraints & Dependencies
- Standard Markdown and CSS integration rules for resume.lol.

## Validation
- Validate using test suites that verify markdown output structure.

## Subtasks
- [ ] `[CONFIRMED]` Build the prompt and response handling for the auditor agent.
- [ ] `[CONFIRMED]` Develop the console output formatter for scores and comparison diffs.
- [ ] `[CONFIRMED]` Create the resume.lol markdown exporter.
