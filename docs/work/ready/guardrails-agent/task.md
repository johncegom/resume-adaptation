# Tasks - Validation and Guardrails Agent

This document tracks all implementation tasks for the validation and guardrails agent feature.

## Task List

- [ ] RAFT0021: Define Guardrail interfaces and validation result types
  - **Acceptance Criteria**: Interface definition for `Validator` and data structures representing validation checks (`ValidationResult` containing checks status) are written in `internal/adaptation/guardrails.go`.
  - **Dependencies**: None

- [ ] RAFT0022: Implement Fact Consistency Validator
  - **Acceptance Criteria**: Prompts Gemini API to compare original and adapted resumes to detect hallucinations, exaggerations, or inaccuracies, returning structured details of violations.
  - **Dependencies**: RAFT0021

- [ ] RAFT0023: Implement formatting and completeness validators
  - **Acceptance Criteria**: Validates structural fields (non-empty critical data) and basic markdown list formats.
  - **Dependencies**: RAFT0022

- [ ] RAFT0024: Implement automated Repair/Retry loop
  - **Acceptance Criteria**: Coordinates correction retries up to a maximum limit of 3, outputting a wrapped boundary error if limit is exceeded.
  - **Dependencies**: RAFT0023

- [ ] RAFT0025: Integrate Guardrails validation into CLI flow
  - **Acceptance Criteria**: Integrates the validation and repair loops into the `adapt` command, printing validation warnings/errors to Stderr, blocking export on critical failure, and returning clean exit codes.
  - **Dependencies**: RAFT0024
