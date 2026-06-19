# Task: Profile & Setting Configurations (RAFT0002)

## Problem Statement
The user needs to store defaults like candidate name, default resume PDF path, and AI credentials to avoid entering them on every CLI run.

## Goals
- Create a local configuration utility in Go.
- Support reading/writing settings to a cross-platform location (e.g. `$HOME/.config/resume-adaptation/config.json`).
- Implement the `config` CLI subcommand to view/modify these settings.

## Non-Goals
- Performing online credential verification with AI APIs.

## Acceptance Criteria
- [ ] Config file is generated if it doesn't exist.
- [ ] Config file supports fields: `gemini_api_key`, `default_resume_path`, and `default_output_dir`.
- [ ] Command `resume-adapt config set <key> <value>` writes to file.
- [ ] Command `resume-adapt config show` reads and prints settings.

## Related Tasks
- **Block by:** `RAFT0001`
- **Related:** `RAFT0004` (AI adaptation uses API key), `RAFT0003` (defaults to config resume path)

## Success Metrics
- Read/Write configuration operations execute under 50ms.
- 100% preservation of file permissions (readable only by the local user).

## Constraints & Dependencies
- Cross-platform file path resolution (`os.UserConfigDir()`).

## Validation
- Validate using tests that write temp configs and verify their retrieval values.

## Subtasks
- [ ] `[CONFIRMED]` Setup configuration schema structs and file handling logic.
- [ ] `[CONFIRMED]` Integrate the config handling with Cobra subcommand routing.
- [ ] `[ASSUMPTION]` Mask sensitive credentials (like API keys) when showing config in the terminal.
