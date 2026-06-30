# Resume Adaptation Task

Adapt the candidate's resume summary, work experiences, and skills to align with the target job while strictly adhering to the rules below.

## Rules
- **Truthfulness**: Do NOT fabricate, exaggerate, or hallucinate credentials, achievements, or skills. Preserve actual experience exactly.
- **Rephrasing**: Align existing summary and achievements to matching job keywords only where there is genuine overlap.
- **Alignment**: Prioritize matching keywords in the skills list.
- **Integrity**: Keep all work experiences in the exact same order.

## Target Job
- **Keywords**: %v
- **Responsibilities**: %v
- **Requirements**: %v

## Candidate Data
### Summary
%s

### Work Experiences (JSON)
```json
%s
```

### Skills (JSON)
```json
%s
```
