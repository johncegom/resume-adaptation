# Task: AI Resume Adaptation Engine (RAFT0004)

## Problem Statement
The tool must rewrite the resume sections using the Gemini API to align with the keywords and demands of the Job Description without inventing experiences.

## Goals
- Connect to Gemini API using Go client SDK.
- Formulate prompt templates with inputs: `parsed resume sections` and `job description text`.
- Implement bullet point rewriting logic matching key requirements.

## Non-Goals
- Formatting the output as markdown or PDF in this task.
- Validating the rewrite for truthfulness (handled by `RAFT0005`).

## Acceptance Criteria
- [ ] Connects to Gemini API using API key from configurations.
- [ ] Sends structured prompts and handles model responses.
- [ ] Receives rewritten bullet points that map to job description keywords.

## Related Tasks
- **Block by:** `RAFT0002` (config), `RAFT0003` (pdf-ingestion)
- **Related:** `RAFT0005` (validation)

## Success Metrics
- Average latency per section adaptation: Under 3 seconds.
- Matches relevant job keywords in the output sections.

## Constraints & Dependencies
- Dependency on the official Google GenAI Go library (`google.golang.org/genai` or similar stable client).

## Validation
- Validate using mock Gemini client tests that assert correct prompt generation.

## Subtasks
- [ ] `[CONFIRMED]` Integrate the official Google GenAI SDK.
- [ ] `[CONFIRMED]` Develop structured prompts instructing the AI model to rephrase experience bullet points.
- [ ] `[ASSUMPTION]` Add simple mock/stub implementations for offline development and testing.
