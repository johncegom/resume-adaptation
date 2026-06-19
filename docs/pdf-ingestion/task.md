# Task: PDF Resume Ingestion (RAFT0003)

## Problem Statement
The tool must read candidate details from a local PDF resume and structure it in-memory to be ready for AI adaptation.

## Goals
- Parse text contents from a PDF file in Go.
- Separate and identify key blocks (e.g., Contact Info, Experience, Education, Skills).
- Model resume contents using clean Go structs.

## Non-Goals
- Generating a PDF output.
- Complex OCR parsing for scanned images of resumes.

## Acceptance Criteria
- [ ] Ingests a candidate's resume PDF from a local filepath.
- [ ] Correctly extracts raw text strings.
- [ ] Structures text into Go model sections: `Resume`, `ExperienceBlock`, `SkillBlock`, `EducationBlock`.

## Related Tasks
- **Block by:** `RAFT0001`
- **Related:** `RAFT0004` (processes parsed models)

## Success Metrics
- Extraction speed: Under 500ms for a standard 2-page PDF resume.
- Section mapping accuracy > 90% for typical resume layouts.

## Constraints & Dependencies
- Dependency on a stable, dependency-free Go PDF parser library (e.g., `github.com/ledongthuc/pdf` or standard-compatible alternative).

## Validation
- Validate by parsing sample test resumes and asserting that extracted text contains expected keywords and section blocks.

## Subtasks
- [ ] `[CONFIRMED]` Research and select the most reliable, pure-Go PDF text extractor.
- [ ] `[CONFIRMED]` Setup resume Go structures/models.
- [ ] `[ASSUMPTION]` Build simple heuristics/regex-based parsers to split text into Sections.
