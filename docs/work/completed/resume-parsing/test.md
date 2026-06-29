# Test Specification - Resume Ingestion and Parsing

Verification matrix and quality criteria for the resume ingestion and parsing engine.

## Quality Criteria & Verification Matrix

- [x] Verification of Go module setup and library dependencies installation. (See RAFT0001 in task.md)
- [x] Verification of model serialization/deserialization correctness. (See RAFT0002 in task.md)
- [x] Integration test: Parse raw text file with valid format. (See RAFT0003 in task.md)
- [x] Integration test: Parse raw PDF file via Gemini API Structured Outputs. (See RAFT0004 in task.md)
- [x] Boundary test: Handling of empty or malformed files (empty PDF, malformed text). (See RAFT0003 & RAFT0004 in task.md)
- [x] Security test: Validate file path traversal prevention during local file ingestion. (See RAFT0003 in task.md)
