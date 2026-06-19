# Epic: PDF Resume Ingestion

## What
Parses uploaded PDF resumes and converts them into structured Go data types representing resume sections (Experience, Skills, Education).

## Why
Allows users to directly supply their standard PDF files rather than having to manually paste or write text files.

## How
- Integrate a Go PDF library (such as `github.com/ledongthuc/pdf`).
- Extract plain text and run lightweight parsing/regex to identify section boundaries.
- Build Go structs to represent the parsed sections.
