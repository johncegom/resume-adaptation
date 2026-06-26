# Feature: AI Adaptation and Alignment Engine

## What
An engine that matches the candidate's structured resume against a target job description. It identifies keyword gaps, extracts requirements, and generates tailored professional descriptions that emphasize the candidate's actual matching skills and achievements.

## Why
This is the core value proposition. Aligning the candidate's resume keywords to the job description maximizes ATS pass rates and shows the candidate's exact fit for the role.

## How
- Interface with the Google Gemini API using the Go SDK.
- Create prompt templates for analyzing job description requirements.
- Create prompt templates for rewriting and aligning professional achievements.
- Handle API client setup with timeouts and cancelable context propagation.

## When
This will be implemented second, immediately after parsing is established.
