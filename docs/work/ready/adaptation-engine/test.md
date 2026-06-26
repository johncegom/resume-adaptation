# Test Specification - AI Adaptation and Alignment Engine

Verification matrix and quality criteria for the AI adaptation engine.

## Quality Criteria & Verification Matrix

- [ ] Verification of Gemini API client setup and API key validation. (See RAFT0006 in task.md)
- [ ] Verification of Job Description analyzer (checking correct extraction of keywords from sample JDs). (See RAFT0007 in task.md)
- [ ] Verification of resume rewriting (asserting that rewritten bullet points contain job description keywords and preserve candidate facts). (See RAFT0008 & RAFT0009 in task.md)
- [ ] Boundary test: Handle empty/very short job descriptions. (See RAFT0007 & RAFT0009 in task.md)
- [ ] Security test: Validate API key exposure safety (no logging of secrets/API keys). (See RAFT0006 in task.md)
