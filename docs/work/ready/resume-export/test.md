# Test Specification - Structured Export to resume.lol Format

Verification matrix and quality criteria for the resume exporter.

## Quality Criteria & Verification Matrix

- [ ] Verification of Exporter template parsing (checks template compile errors). (See RAFT0017 in task.md)
- [ ] Unit test: Verification of `text/template` rendering correctness with a mock `Resume`. (See RAFT0018 in task.md)
- [ ] Integration test: Export to file path (confirming file is created and contains correct data). (See RAFT0019 in task.md)
- [ ] Security test: Validate path traversal protection (preventing writes or template loads using relative paths like `../../` outside allowed scopes using `filepath.IsLocal`/`filepath.Rel` verification). (See RAFT0019 in task.md)
- [ ] CLI integration: `export` subcommand functions correctly and handles pipes. (See RAFT0020 in task.md)
