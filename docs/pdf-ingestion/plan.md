# Plan: PDF Resume Ingestion (RAFT0003)

## Problem Statement
Implement Go text extraction from PDF files and structure it into section modules.

## Goals
- Add PDF extraction dependency.
- Define internal structures and map parsed text to them.

## Non-Goals
- Complex visual layout bounding-box extraction.

## Acceptance Criteria
- [ ] Sample PDF parser test extracts target text successfully.

## Related Tasks
- **Block by:** `RAFT0001`
- **Related:** `RAFT0004`

## Success Metrics
- Fully unit-tested PDF ingestion.

## Constraints & Dependencies
- Compatible Go modules, no external binary dependencies (like `pdftotext` system dependencies).

## Validation
- Validate via tests running extraction against mock/fixture PDFs.

## Subtasks & Implementation Plan
- [ ] `[CONFIRMED]` Add testing fixtures for standard resumes.
- [ ] `[CONFIRMED]` Write placeholder tests asserting PDF parsing mapping structures.
- [ ] `[CONFIRMED]` Implement parsing functions integrating the PDF extraction library.
