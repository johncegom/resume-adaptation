# Tasks - Structured Export to resume.lol Format

This document tracks all implementation tasks for the resume exporter feature.

## Task List

- [ ] RAFT0016: Define Exporter interface and project setup
  - **Acceptance Criteria**: Package structure is created under `internal/export/` with defined interface methods.
  - **Dependencies**: None

- [ ] RAFT0017: Create default embedded resume.lol markdown template
  - **Acceptance Criteria**: The default Markdown template is designed and embedded in the binary using Go's `embed` package.
  - **Dependencies**: RAFT0016

- [ ] RAFT0018: Implement template renderer
  - **Acceptance Criteria**: Template parser parses the embedded template and correctly translates the `Resume` struct fields into Markdown.
  - **Dependencies**: RAFT0017

- [ ] RAFT0019: Implement file writing and pipeline output with path validation
  - **Acceptance Criteria**: Service outputs the rendered string to a specified output file path with file path validation (using `filepath.IsLocal` / `filepath.Rel` traversal protections), or outputs to stdout.
  - **Dependencies**: RAFT0018

- [ ] RAFT0020: Add Cobra CLI `export` command
  - **Acceptance Criteria**: Registers the command `go run cmd/cli/main.go export` which reads adapted data and triggers the exporter. Implements stdout redirection capability.
  - **Dependencies**: RAFT0019
