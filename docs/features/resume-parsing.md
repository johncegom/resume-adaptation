# Feature: Resume Ingestion and Parsing

## What
Provides the capability to ingest and parse candidate resumes from PDF documents and raw text files. It parses the document content, identifies core sections (profile summary, professional experience, education, skills, projects), and structures them into a clean internal Go representation.

## Why
Users want to quickly import their existing resumes without manual copying and pasting. Support for PDF is critical as it is the most common format for existing resumes, providing a seamless user experience.

## How
- Use a Go-native PDF parsing library to extract text content.
- Implement parsing models and schemas to represent the structured resume.
- Integrate section segmenters to partition the resume content.

## When
This is the foundational phase of the data pipeline and should be implemented first.
