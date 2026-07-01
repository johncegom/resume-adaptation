# Test Specification - Interactive UI Adaptation Wizard

Verification matrix and quality criteria for the interactive terminal user interface.

## Quality Criteria & Verification Matrix

- [ ] Verification of styling and layout rendering (compilation and error-free execution of CLI screens). (See RAFT0011 in task.md)
- [ ] Unit test: Validation helper for input paths (checking directories vs valid files). (See RAFT0012 in task.md)
- [ ] Manual verification: Spinner responsiveness during mocked network calls. (See RAFT0013 in task.md)
- [ ] Manual verification: Review editor screen scrolling and input changes. (See RAFT0014 in task.md)
- [ ] CLI integration: `adapt` command initializes forms correctly. Commands return errors to `main` without invoking `os.Exit` inside subcommands. (See RAFT0015 in task.md)
- [ ] I/O separation verification: Confirm that redirecting stdout captures only the clean adapted resume, while all prompts and spinners output to stderr. (See RAFT0015 in task.md)
- [ ] Isolation test: CLI command unit tests must instantiate a fresh Cobra command tree per execution to avoid flag state accumulation. (See RAFT0015 in task.md)
