# AGENTS

This file contains the **Rules and guidelines** for humans and LLMs working in this Context Storehouse; detailing processes for planning, TDD, review, and completion behavior.

## Getting Started (Agents)

1. Load any needed skills. If you identify a missing skill that you can suggest from memory/context without needing a web search, propose it to the human first and ask for approval to install it using `npx skill`.
2. Make the smallest reversible change. Follow TDD for code edits.
3. When closing out, report: file changed, commands/tests run, checks skipped (with reasons), and links to affected code/docs/specs/tasks.
4. MUST not use emojis / icons in markdown file.

## Initiate Context Storehouse

**RUN ONCE ONLY** If this section is not in `AGENTS.md`:
1. Ask the user for a task ID prefix (e.g., `CSHT`).
2. Suggest a task ID format (e.g., `0001`).
3. Append the prefix and format to `AGENTS.md`. Always use this task ID format.

### Project Context Refinement Flow
- **Command `/project-init` / `project init`**:
  1. **Interactive Intake**: Collect and refine the following sequentially:
     - **Purpose**: Action-oriented 1-2 sentence core utility description.
     - **Tech Stacks**: Bulleted list (versions, main libraries, runtimes).
     - **Entry Points**: Absolute `file:///` URIs + brief role descriptions.
  2. **Confirm**: Show markdown preview block. Ask: *"Would you like me to write this refined context to AGENTS.md?"*
  3. **Execute CLI**: On approval, run `csh context` tool updates:
     - `csh context "Purpose" "<refined_purpose>"`
     - `csh context "Tech Stacks" "<refined_tech_stacks>"`
     - `csh context "Entry Points" "<refined_entry_points>"`
  4. **Generate Routing**: Map refined Tech Stacks into `## Required Language Skills` using its blueprint layout. Ensure triggers target intent/path scope to support the Planning Phase.

## Required Language Skills (Dynamic Loading)

### Active Routing Blueprint
1. **Scope Detection**: At start of every task/planning phase, explicitly declare target languages, paths, or sub-systems intended to be created, modified, or affected.
2. **Dynamic Load**: If scope matches a rule below, immediately load the specified orchestrator skill. Use it to lazily load sub-skills (e.g., loading testing skills only during test execution). Do not load statically.

#### Go / Golang CLI
- **Trigger**: Scope includes path `cmd/` or extension `*.go` or `go.mod` (regardless of current file existence).
- **Action**: Invoke `golang-how-to`.

## General Guidelines (Stack-Agnostic)

### Feature Implementation Planning
- **Feature Overview File**: Each feature must have its own markdown reference file located under the `<root directory>/docs/features/<feature-name>.md` describing the overview of what, how, when, and why.
- **Feature Directory & Documents**: Each feature occupies a dedicated directory under `<root directory>/docs/work/<status>/<feature-name>/`. Agents must physically migrate the feature directory sequentially through the lifecycle states below.
  
  ### Directory Status Lifecycle Flow
  1. `pending`: Raw backlog. Feature is proposed, but technical design has not started.
  2. `planning`: Active design phase. Agent writes `plan.md`, `task.md`, and `test.md` inside this directory.
  3. `ready`: Approved backlog. Transitioned **automatically** by the agent the moment all planning documents are completed and written. Signifies the design is locked and the feature is ready to be picked up.
  4. `inprogress`: Active development. Moved here **only** when the human gives an explicit command to start implementation or assigns a specific task ID to execute. Triggers the TDD loop.
  5. `completed`: Production-ready. Moved here only after all tasks are completed (`- [x]`) and final validation passes.

  ### Core Artifacts (Per Feature Folder)
  - `plan.md`: Technical design document (architecture, mermaid graphs, decisions). Prose-only; **no task status checkboxes**. References task IDs inline.
  - `task.md`: **Single Source of Truth (SSOT)** for tracking. Contains all task checkboxes (`- [ ]` / `- [x]`), dependencies, and acceptance criteria. This is the only file where task status checkboxes live.
  - `test.md`: Quality criteria & verification matrix (boundaries, security, regression). Uses checkboxes *only* for quality sign-off status, referencing task IDs via "See task.md". Does not contain task status checkboxes.
- **Task Specification**: Must use `csh task add` to create a task, and `csh task complete` to complete a task.

### Development Phases
Work is strictly divided into two distinct, human-controlled operational phases:
- **Planning Phase**: Active design generation. Ends **automatically** when `plan.md`, `task.md`, and `test.md` are written and the directory auto-migrates to `ready`. No code or workspace changes are allowed during this phase.
- **Implementation Phase**: Triggered **only** by explicit human task assignment or execution instruction on a `ready` feature. Follows the strict TDD loop. Never mix test and code changes in a single turn.

### Feature Implementing
- Clarity over cleverness.
- Small, reversible steps.
- Follow TDD for code edits.
- **Minimum Code Philosophy:** Write the minimum code necessary to make the current test case pass. Do not write code or features that are not yet covered by an active failing test case.
- **Strict TDD Flow for Code Edits:** Apply this as a loop for each subtask first (or per-test if a subtask is too large):
  1. Propose and write test cases first as placeholders with an explanatory comment and a deliberate failing assertion stating it is "not implemented yet".
  2. **MUST** iterate through those test cases one-by-one:
     a. Write actual test implementation (replacing placeholder assertion with real logic).
     b. Run tests to verify the newly implemented test case fails (Red phase).
     c. Propose and write implementation code to make it pass (Green phase).
     d. Run tests to verify this specific test case passes while remaining placeholders continue to fail.
  3. When a task is complete, stop, propose, and offer to create a commit following conventional commits.
  4. Never propose test changes and implementation changes in the same turn.
  5. When a feature is done, stop and run universal validations to ensure all checks pass.

### Impact Analysis
- **No Blind Changes**: Search the codebase (e.g., grep) for call sites and dependencies before modifying or deleting any public symbols to prevent breaking imports.

### Complete Implementations
- **No Placeholders**: Avoid writing `// TODO` comments, empty stubs, or partial logic blocks unless explicitly requested. Always deliver fully functional, production-ready code.

### Safe Error Handling
- **Defensive Execution**: Always validate inputs, check error returns, and handle nil/null or empty states gracefully. Never silently suppress errors or exceptions.

### Configuration & Setup Sync
- **Keep Setup Up-to-Date**: If a change adds, modifies, or deletes configuration flags, environment variables, or schema definitions, update both the code and the corresponding documentation or template files (e.g., `.env.example`, setup guides) in the same step.

### Validation Constraint
- **Universal Validation Rules**: To ensure the system is production-grade at both Feature (Epic) and Task (Subtask) levels:
  - Validate all configurations, incoming payloads, and input arguments at the boundary of execution.
  - Reject malformed or invalid inputs early with descriptive error messages and clean exit/status codes.
  - Design systems to be fail-safe; invalid or missing inputs must not trigger unhandled exceptions, panics, or resource leaks.
  - Write dedicated unit and integration tests covering invalid input ranges, boundary conditions, and error propagation pathways to verify validation coverage.

## Boundaries

- **Consent First**: Do not create, write, modify, or delete any files in the workspace before presenting proposed changes in the chat and obtaining explicit approval ("approved") from the user.
- **Directory Isolation**: Restrict all file operations and command executions to the project workspace. Do not read, write, or execute files outside the project folder unless explicitly instructed.
- **Minimal privileges & bypasses**: Only request sandbox bypasses (`BypassSandbox: true`) if a sandboxed command fails due to path visibility or necessary local configurations.
- **Resource Discipline**: Do not trigger long-running background tasks, infinite loops, or recursive commands that might hang the terminal or consume system resources indefinitely.
- **Ask first before**:
  - Adding new runtime dependencies.
  - Introducing a new workflow stage, skill, or phase agent.

## Tooling
- **Go Toolchain**: Use standard Go commands for building and testing:
  - Run all tests: `go test -v ./...`
  - Run package tests: `go test -v ./internal/...`
  - Format code: `go fmt ./...`
  - Check static analysis: `go vet ./...`
- **CLI Development**: Run the application directly using `go run cmd/cli/main.go`.
- **Task Management**: Use `csh task add` to register new tasks and `csh task complete <task-id>` to mark tasks as completed.

## Documentation Standards

- Keep docs index updated.
- Update component READMEs when behavior changes.
- Prefer link-rich docs for LLM traversal.
- Record significant decisions alongside the relevant phase artifacts (under `docs/`).
- When closing a task, report: summarize changed files, run commands/tests, skipped checks (with reasons), and links to affected code/docs/specs/tasks.

## Project Context
### Purpose
A command-line tool designed to dynamically adapt a candidate's resume to align with specific job descriptions, maximizing hiring potential while strictly preserving core values, achievements, and truthfulness.

### Tech Stacks
- Go 1.22+
- Cobra CLI framework (github.com/spf13/cobra)
- Charmbracelet libraries (Lipgloss, Huh, Bubble Tea)
- Google Gemini API (Go SDK)

### Entry Points
- [main.go](file:///C:/Users/Admin/Data/Developer/PersonalProject/resume-adaptation/cmd/cli/main.go) - Primary CLI bootstrap and command router.

### Repo Structure
- `cmd/cli/`: CLI entry point bootstrap and command declarations using Cobra.
- `internal/`: Core, non-importable application logic:
  - `internal/adaptation/`: Resume rewrite logic, prompt templates, and Google Gemini API client setup.
  - `internal/parser/`: PDF resume parsing utilities.
  - `internal/ui/`: Interactive terminal components using Charmbracelet's Lipgloss, Huh, and Bubble Tea.
- `docs/`: Documentation and workflow tracking:
  - `docs/features/`: Markdown references detailing feature requirements.
  - `docs/work/`: Directory-based workflow statuses (`pending`, `planning`, `ready`, `inprogress`, `completed`).

### Operating Principles
- **No Hallucinations**: The resume adapter must strictly align existing candidate experiences and achievements to the job description without fabricating, exaggerating, or hallucinating new qualifications.
- **Separation of Concerns**: Keep CLI commands and UI view models (`cmd/` and `internal/ui/`) completely decoupled from core adaptation and parsing logic (`internal/adaptation/` and `internal/parser/`).
- **Defensive API Calls**: Ensure all remote calls to Google Gemini API are protected with proper timeouts, retries, and context cancellation propagation.
- **Fail-Safe Processing**: Provide user-friendly CLI errors and clean exit codes if parsing or adaptation fails (e.g., malformed PDF, invalid API key, network issues) rather than allowing panics.

**Configured Workspace ID Format:**
- Prefix: `RAFT`
- Format: `0001` (e.g., `RAFT0001`)
