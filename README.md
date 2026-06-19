# Resume Adaptation CLI

A powerful command-line tool designed to dynamically adapt a candidate's resume to align with specific job descriptions, maximizing hiring potential while strictly preserving core values, achievements, and truthfulness.

## Features

- **Flexible Ingestion:** Ingests the Job Description (JD) and the Resume either as direct command-line arguments/prompts or from preconfigured defaults (e.g., a default candidate profile resume).
- **PDF Ingestion:** Accepts and parses PDF resumes to extract raw candidate text.
- **Context-Aware Adaptation:** Automatically analyzes and aligns resume text with the keywords and requirements of a target job description.
- **AI API Integration:** Communicates with AI APIs (such as Gemini API) to perform smart contextual rewriting and validation.
- **Core Value Preservation:** Safeguards candidate achievements, ensuring no exaggeration or fabrication occurs.
- **Agent Validation:** Utilizes an AI agent step to verify truthfulness, coherence, and professional relevance of the adapted resume.
- **Rich Terminal UI:** A highly interactive, beautiful terminal interface powered by Cobra and Charmbracelet libraries.
- **resume.lol Output Format:** Exports the adapted resume directly as a structured Markdown file formatted specifically to be drop-in ready for pasting into the [resume.lol](https://resume.lol/) editor.

## Tech Stack

- **Language & Runtime:** Go 1.22+
- **CLI Framework:** [Cobra](https://github.com/spf13/cobra)
- **Terminal UI & Styling:** [Charmbracelet](https://github.com/charmbracelet) (e.g., Lipgloss, Huh, Bubble Tea)
- **APIs:** AI Client APIs (e.g., Gemini SDK for Go / HTTP API clients)

## Getting Started

### Entry Point
- The primary entry point is located at [cmd/cli/main.go](file:///D:/Developer/PersonalProject/resume-adaptation/cmd/cli/main.go).

### Directory Layout
```text
.
├── cmd/
│   └── cli/
│       └── main.go       # CLI Bootstrap & Commands
├── AGENTS.md             # Developer Rules and Project Context
└── README.md             # Project Description and Usage
```
