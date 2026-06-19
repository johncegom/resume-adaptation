# AGENTS

This file is the **Rules and guidelines** for humans and LLMs working in this Context Storehouse; with **detailed process, TDD, review, and completion behaviour.**.

## Getting Started (Agents)

1. Load any needed skills. If you identify a missing skill that you can suggest from memory/context without needing a web search, propose it to the human first and ask for approval to install it using `npx skill`.
2. Make the smallest reversible change. Follow TDD for code edits.
3. When closing out, report: file changed, commands/tests run, checks skipped (with reasons), and links to affected code/docs/specs/tasks.

## Initiate Context Storehouse

**RUN ONCE ONLY** If this section is not in `AGENTS.md`:
1. Ask the user for a task ID prefix (e.g., `CSHT`).
2. Suggest a task ID format (e.g., `0001`).
3. Append the prefix and format to `AGENTS.md`. Always use this task ID format.

### Project Context Refinement Flow
- **Agent Command `/project-init`**: When the user requests `/project-init` (or `project init`) in chat:
  1. Guide the user step-by-step to collect details for the project context.
  2. First, ask the user for the **Purpose** of the project. Refine it into an action-oriented, 1-2 sentence description detailing the core utility.
  3. Next, ask the user for the **Tech Stacks** used. Refine them into bulleted lists specifying language versions, main libraries, and runtimes.
  4. Then, ask the user for the **Entry Points** of the project. Refine them into markdown links using absolute `file:///` URIs accompanied by brief descriptions of their roles.
  5. Present the fully refined project context as a markdown preview block and ask the user for confirmation: *"Would you like me to write this refined context to AGENTS.md?"*
  6. Upon approval, execute the updates using the `csh context` CLI tool:
     - `csh context "Purpose" "<refined_purpose>"`
     - `csh context "Tech Stacks" "<refined_tech_stacks>"`
     - `csh context "Entry Points" "<refined_entry_points>"`

## General Guidelines (Stack-Agnostic)
 
### Feature Implementation Planning
- **Feature Overview File**: Each feature must have its own markdown reference file located under the `<root directory>/docs/features/` directory (e.g., `<root directory>/docs/features/<feature-name>.md`) describing the overview of the what, how, when, and why.
- **Feature Directory & Documents**: Each feature must have its own dedicated directory under `<root directory>/docs/` named after the feature (e.g., `<root directory>/docs/<feature-name>/`), containing:
  - `plan.md`: It tells what code to add, what boundaries to preserve , and what tests must prove before the task can be considered done.
  - `task.md`: It tells what the task is suppose to accomplish, where it came from, what it depends on, and how to know when it is done.
  - `test.md`: It tells what tests must be covered to ensure the feature runs smoothly and securely. This file serves as the quality checklist to guarantee the program is production-grade and ready for deployment.
- **Task Specification & Tracking**: The `plan.md` and `task.md` must:
  - List all subtasks using checkboxes (e.g., `- [ ]` for pending, `- [x]` for completed).
  - Categorize each task using either `[CONFIRMED]` or `[ASSUMPTION]`.
  - Explicitly map out task dependencies (**Block by**, **Related**).
  - `plan.md` and `task.md` must has: problem statement, goals, non-goals, acceptance criteria (using checkboxes (e.g., `- [ ]` for pending, `- [x]` for completed)), related tasks, success metrics, constraints & dependencies and validation (the task / feature must not break other stuff) sections.
  - `test.md` must has: feature scope & impact, functional testing (happy path), boundary & negative testing, security & vulnerability checks, regression & integration testing (must not break existing features), performance & scalability criteria and sign-off & verification checklist sections. For each item under these sections, it must use checkboxes (e.g., `- [ ]` for pending, `- [x]` for completed), explicitly show which test case or mechanism proves that it has been verified, and identify the type of testing used (e.g., unit test, integration test, fuzz test, manual verification, static check, etc.).

### Feature Implementating

- Clarity over cleverness.
- Small, reversible steps.
- **Strict TDD Flow for Code Edits:** Apply this as a loop for each subtask first (or per-test if a subtask is too large):
  1. Propose and write the test cases first as placeholders. Each placeholder must include a comment explaining the purpose of the test case, and a deliberate failing assertion stating that it is "not implemented yet".
  2. Iterate through those test cases one-by-one:
     a. Write the actual test case implementation (replacing the placeholder assertion with the real test logic).
     b. Run the tests to verify that the newly implemented test case fails (Red phase).
     c. Only after confirming the failure, propose and write the implementation code to make it pass (Green phase).
     d. Run the tests again to verify that this specific test case passes (while other unimplemented test cases continue to fail on their placeholder assertions).
  3. When a task is done (close a task), agent must stop, propose, and offer to create commit follows conventional commit.
  4. Never propose test changes and implementation changes in the same turn.
  5. When a feature is done, agent must stop and do validation and make sure it must all passed.

### Impact Analysis Before Edits
- **No Blind Changes**: Before modifying or deleting any symbol (type, interface, class, function, variable), run a codebase search (e.g., grep) to locate all call sites and dependencies to prevent breaking other modules.

### Complete Implementations
- **No Placeholders**: Avoid writing `// TODO` comments, empty stubs, or partial logic blocks unless explicitly requested. Always deliver fully functional, production-ready code.

### Safe Error Handling
- **Defensive Execution**: Always validate inputs, check error returns, and handle nil/null or empty states gracefully. Never silently suppress errors or exceptions.

### Configuration & Setup Sync
- **Keep Setup Up-to-Date**: If a change adds, modifies, or deletes configuration flags, environment variables, or schema definitions, update both the code and the corresponding documentation or template files (e.g., `.env.example`, setup guides) in the same step.

## Boundaries

- **Consent First**: Do not create, write, modify, or delete any files in the workspace before presenting proposed changes in the chat and obtaining explicit approval from the user.
- **Directory Isolation**: Restrict all file operations and command executions to the project workspace. Do not read, write, or execute files outside the project folder unless explicitly instructed.
- **Minimal privileges & bypasses**: Only request sandbox bypasses (`BypassSandbox: true`) if a sandboxed command fails due to path visibility or necessary local configurations.
- **Resource Discipline**: Do not trigger long-running background tasks, infinite loops, or recursive commands that might hang the terminal or consume system resources indefinitely.
- **Ask first before**:
  - Adding new runtime dependencies.
  - Introducing a new workflow stage, skill, or phase agent.

## Documentation Standards

- Keep docs index updated.
- Update component READMEs when behavior changes.
- Prefer link-rich docs for LLM traversal.
- When close a task, report: summarize changed files, run commands/tests, skipped checks (with reasons), and links to affected code/docs/specs/tasks.

## Project Context
Refinement Flow
- **Agent Command `/project-init`**: When the user requests `/project-init` (or `project init`) in chat:
  1. Guide the user step-by-step to collect details for the project context.
  2. First, ask the user for the **Purpose** of the project. Refine it into an action-oriented, 1-2 sentence description detailing the core utility.
  3. Next, ask the user for the **Tech Stacks** used. Refine them into bulleted lists specifying language versions, main libraries, and runtimes.
  4. Then, ask the user for the **Entry Points** of the project. Refine them into markdown links using absolute `file:///` URIs accompanied by brief descriptions of their roles.
  5. Present the fully refined project context as a markdown preview block and ask the user for confirmation: *"Would you like me to write this refined context to AGENTS.md?"*
  6. Upon approval, execute the updates using the `csh context` CLI tool:
     - `csh context "Purpose" "<refined_purpose>"`
     - `csh context "Tech Stacks" "<refined_tech_stacks>"`
     - `csh context "Entry Points" "<refined_entry_points>"`

## General Guidelines (Stack-Agnostic)
 
### Feature Implementation Planning
- **Feature Overview File**: Each feature must have its own markdown reference file located under the `<root directory>/docs/features/` directory (e.g., `<root directory>/docs/features/<feature-name>.md`) describing the overview of the what, how, when, and why.
- **Feature Directory & Documents**: Each feature must have its own dedicated directory under `<root directory>/docs/` named after the feature (e.g., `<root directory>/docs/<feature-name>/`), containing:
  - `plan.md`: It tells what code to add, what boundaries to preserve , and what tests must prove before the task can be considered done.
  - `task.md`: It tells what the task is suppose to accomplish, where it came from, what it depends on, and how to know when it is done.
  - `test.md`: It tells what tests must be covered to ensure the feature runs smoothly and securely. This file serves as the quality checklist to guarantee the program is production-grade and ready for deployment.
- **Task Specification & Tracking**: The `plan.md` and `task.md` must:
  - List all subtasks using checkboxes (e.g., `- [ ]` for pending, `- [x]` for completed).
  - Categorize each task using either `[CONFIRMED]` or `[ASSUMPTION]`.
  - Explicitly map out task dependencies (**Block by**, **Related**).
  - `plan.md` and `task.md` must has: problem statement, goals, non-goals, acceptance criteria (using checkboxes (e.g., `- [ ]` for pending, `- [x]` for completed)), related tasks, success metrics, constraints & dependencies and validation (the task / feature must not break other stuff) sections.
  - `test.md` must has: feature scope & impact, functional testing (happy path), boundary & negative testing, security & vulnerability checks, regression & integration testing (must not break existing features), performance & scalability criteria and sign-off & verification checklist sections. For each item under these sections, it must use checkboxes (e.g., `- [ ]` for pending, `- [x]` for completed), explicitly show which test case or mechanism proves that it has been verified, and identify the type of testing used (e.g., unit test, integration test, fuzz test, manual verification, static check, etc.).

### Feature Implementating

- Clarity over cleverness.
- Small, reversible steps.
- **Strict TDD Flow for Code Edits:** Apply this as a loop for each subtask first (or per-test if a subtask is too large):
  1. Propose and write the test cases first as placeholders. Each placeholder must include a comment explaining the purpose of the test case, and a deliberate failing assertion stating that it is "not implemented yet".
  2. Iterate through those test cases one-by-one:
     a. Write the actual test case implementation (replacing the placeholder assertion with the real test logic).
     b. Run the tests to verify that the newly implemented test case fails (Red phase).
     c. Only after confirming the failure, propose and write the implementation code to make it pass (Green phase).
     d. Run the tests again to verify that this specific test case passes (while other unimplemented test cases continue to fail on their placeholder assertions).
  3. When a task is done (close a task), agent must stop, propose, and offer to create commit follows conventional commit.
  4. Never propose test changes and implementation changes in the same turn.
  5. When a feature is done, agent must stop and do validation and make sure it must all passed.

### Impact Analysis Before Edits
- **No Blind Changes**: Before modifying or deleting any symbol (type, interface, class, function, variable), run a codebase search (e.g., grep) to locate all call sites and dependencies to prevent breaking other modules.

### Complete Implementations
- **No Placeholders**: Avoid writing `// TODO` comments, empty stubs, or partial logic blocks unless explicitly requested. Always deliver fully functional, production-ready code.

### Safe Error Handling
- **Defensive Execution**: Always validate inputs, check error returns, and handle nil/null or empty states gracefully. Never silently suppress errors or exceptions.

### Configuration & Setup Sync
- **Keep Setup Up-to-Date**: If a change adds, modifies, or deletes configuration flags, environment variables, or schema definitions, update both the code and the corresponding documentation or template files (e.g., `.env.example`, setup guides) in the same step.

## Boundaries

- **Consent First**: Do not create, write, modify, or delete any files in the workspace before presenting proposed changes in the chat and obtaining explicit approval from the user.
- **Directory Isolation**: Restrict all file operations and command executions to the project workspace. Do not read, write, or execute files outside the project folder unless explicitly instructed.
- **Minimal privileges & bypasses**: Only request sandbox bypasses (`BypassSandbox: true`) if a sandboxed command fails due to path visibility or necessary local configurations.
- **Resource Discipline**: Do not trigger long-running background tasks, infinite loops, or recursive commands that might hang the terminal or consume system resources indefinitely.
- **Ask first before**:
  - Adding new runtime dependencies.
  - Introducing a new workflow stage, skill, or phase agent.

## Documentation Standards

- Keep docs index updated.
- Update component READMEs when behavior changes.
- Prefer link-rich docs for LLM traversal.
- When close a task, report: summarize changed files, run commands/tests, skipped checks (with reasons), and links to affected code/docs/specs/tasks.

## Project Context
Refinement Flow
- **Agent Command `/project-init`**: When the user requests `/project-init` (or `project init`) in chat:
  1. Guide the user step-by-step to collect details for the project context.
  2. First, ask the user for the **Purpose** of the project. Refine it into an action-oriented, 1-2 sentence description detailing the core utility.
  3. Next, ask the user for the **Tech Stacks** used. Refine them into bulleted lists specifying language versions, main libraries, and runtimes.
  4. Then, ask the user for the **Entry Points** of the project. Refine them into markdown links using absolute `file:///` URIs accompanied by brief descriptions of their roles.
  5. Present the fully refined project context as a markdown preview block and ask the user for confirmation: *"Would you like me to write this refined context to AGENTS.md?"*
  6. Upon approval, execute the updates using the `csh context` CLI tool:
     - `csh context "Purpose" "<refined_purpose>"`
     - `csh context "Tech Stacks" "<refined_tech_stacks>"`
     - `csh context "Entry Points" "<refined_entry_points>"`

## General Guidelines (Stack-Agnostic)
 
### Feature Implementation Planning
- **Feature Overview File**: Each feature must have its own markdown reference file located under the `<root directory>/docs/features/` directory (e.g., `<root directory>/docs/features/<feature-name>.md`) describing the overview of the what, how, when, and why.
- **Feature Directory & Documents**: Each feature must have its own dedicated directory under `<root directory>/docs/` named after the feature (e.g., `<root directory>/docs/<feature-name>/`), containing:
  - `plan.md`: It tells what code to add, what boundaries to preserve , and what tests must prove before the task can be considered done.
  - `task.md`: It says what the task is suppose to accomplish, where it came from, what it depends on, and how to know when it is done.
- **Task Specification & Tracking**: The `plan.md` and `task.md` must:
  - List all subtasks using checkboxes (e.g., `- [ ]` for pending, `- [x]` for completed).
  - Categorize each task using either `[CONFIRMED]` or `[ASSUMPTION]`.
  - Explicitly map out task dependencies (**Block by**, **Related**).
  - `plan.md` and `task.md` must has: problem statement, goals, non-goals, acceptance criteria (using checkboxes (e.g., `- [ ]` for pending, `- [x]` for completed)), related tasks, success metrics, constraints & dependencies and validation (the task / feature must not break other stuff) sections.

### Feature Implementating

- Clarity over cleverness.
- Small, reversible steps.
- **Strict TDD Flow for Code Edits:** Apply this as a loop for each subtask first (or per-test if a subtask is too large):
  1. Propose and write the test cases first as placeholders. Each placeholder must include a comment explaining the purpose of the test case, and a deliberate failing assertion stating that it is "not implemented yet".
  2. Iterate through those test cases one-by-one:
     a. Write the actual test case implementation (replacing the placeholder assertion with the real test logic).
     b. Run the tests to verify that the newly implemented test case fails (Red phase).
     c. Only after confirming the failure, propose and write the implementation code to make it pass (Green phase).
     d. Run the tests again to verify that this specific test case passes (while other unimplemented test cases continue to fail on their placeholder assertions).
  3. When a task is done (close a task), agent must stop, propose, and offer to create commit follows conventional commit.
  4. Never propose test changes and implementation changes in the same turn.
  5. When a feature is done, agent must stop and do validation and make sure it must all passed.

### Impact Analysis Before Edits
- **No Blind Changes**: Before modifying or deleting any symbol (type, interface, class, function, variable), run a codebase search (e.g., grep) to locate all call sites and dependencies to prevent breaking other modules.

### Complete Implementations
- **No Placeholders**: Avoid writing `// TODO` comments, empty stubs, or partial logic blocks unless explicitly requested. Always deliver fully functional, production-ready code.

### Safe Error Handling
- **Defensive Execution**: Always validate inputs, check error returns, and handle nil/null or empty states gracefully. Never silently suppress errors or exceptions.

### Configuration & Setup Sync
- **Keep Setup Up-to-Date**: If a change adds, modifies, or deletes configuration flags, environment variables, or schema definitions, update both the code and the corresponding documentation or template files (e.g., `.env.example`, setup guides) in the same step.

## Boundaries

- **Consent First**: Do not create, write, modify, or delete any files in the workspace before presenting proposed changes in the chat and obtaining explicit approval from the user.
- **Directory Isolation**: Restrict all file operations and command executions to the project workspace. Do not read, write, or execute files outside the project folder unless explicitly instructed.
- **Minimal privileges & bypasses**: Only request sandbox bypasses (`BypassSandbox: true`) if a sandboxed command fails due to path visibility or necessary local configurations.
- **Resource Discipline**: Do not trigger long-running background tasks, infinite loops, or recursive commands that might hang the terminal or consume system resources indefinitely.
- **Ask first before**:
  - Adding new runtime dependencies.
  - Introducing a new workflow stage, skill, or phase agent.

## Documentation Standards

- Keep docs index updated.
- Update component READMEs when behavior changes.
- Prefer link-rich docs for LLM traversal.
- When close a task, report: summarize changed files, run commands/tests, skipped checks (with reasons), and links to affected code/docs/specs/tasks.

## Project Context
### Purpose
A tool that dynamically adapts a candidate's resume to align with a specific job description, maximizing hiring potential while strictly preserving the resume's core values and achievements. The adapted content is validated by an agent to ensure truthfulness, coherence, and professional relevance.

### Tech Stacks
*   **Language & Runtime:** Go 1.22+ (Latest stable release)
*   **CLI Framework:** Cobra (`github.com/spf13/cobra`) for command line structuring and arguments
*   **UI & Formatting:** Charmbracelet libraries (e.g. `lipgloss`, `huh`, or `bubbletea`) for beautiful styling and interactive prompts

### Entry Points
*   [cmd/cli/main.go](file:///D:/Developer/PersonalProject/resume-adaptation/cmd/cli/main.go): The main entry point to bootstrap and run the CLI application.

**Configured Workspace ID Format:**
- Prefix: `RAFT`
- Format: `0001` (e.g., `RAFT0001`)
