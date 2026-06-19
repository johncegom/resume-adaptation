# Epic: Agent Validation, Feedback & resume.lol Export

## What
Performs an independent verification loop on adapted resumes, generates a qualitative feedback report showing strengths and missing gaps, displays a terminal diff, and exports the final file in resume.lol markdown format.

## Why
Ensures the user knows exactly why the AI made specific changes, highlights critical skills they might need to add manually, and verifies that the output is honest and highly targeted.

## How
- **Feedback & Benchmarking Engine:**
  - **Strengths:** Identify matches between resume items and JD requirements.
  - **Gap Analysis:** Scan the JD for requirements not found in the resume, outputting actionable advice (e.g., *"The JD requires Docker experience, which is not in your resume. If you have Docker experience, consider adding it to your Skills section"*).
  - **Preservation Check:** Audit changes to ensure no metrics (e.g., team sizes, percentages) were fabricated or exaggerated.
- **Terminal UI Report:**
  - Render a clear text-based scorecard with Bulletins for "Strengths", "Missing Gaps/Improvement Areas", and "AI Rewriting Explanations".
  - Render a side-by-side comparative diff of original vs. adapted text.
- **resume.lol Markdown Export:** Export the adapted resume text to a `.md` file structured for direct copy-paste.
