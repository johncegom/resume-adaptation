# Tasks - Resume Ingestion and Parsing

This document tracks all implementation tasks for the resume ingestion and parsing feature.

## Task List

- [ ] RAFT0001: Initialize Go module and setup project dependencies
  - **Acceptance Criteria**: `go.mod` is created, `google.golang.org/genai` is added as a dependency, and the environment is ready.
  - **Dependencies**: None

- [ ] RAFT0002: Define Resume models and structures
  - **Acceptance Criteria**: Go structs for Resume, WorkExperience, Education, and Projects are defined in `internal/parser/models.go` with JSON/YAML tags.
  - **Dependencies**: RAFT0001

- [ ] RAFT0003: Implement local plaintext reader
  - **Acceptance Criteria**: Service reads local `.txt` and `.md` files, providing path validation and basic formatting.
  - **Dependencies**: RAFT0002

- [ ] RAFT0004: Implement Gemini API PDF parsing with Structured Outputs
  - **Acceptance Criteria**: Client function uploads/sends PDF bytes to Gemini API, requesting a structured response matching the `Resume` JSON schema.
  - **Dependencies**: RAFT0003

- [ ] RAFT0005: Create parser verification tests
  - **Acceptance Criteria**: Integration tests verify that sample TXT and PDF resumes are correctly parsed and unmarshaled.
  - **Dependencies**: RAFT0004
