# Feature: Validation and Guardrails Agent

## What
An AI-driven verification layer that compares the adapted resume against the original resume. It checks for discrepancies, exaggerations, or fabricated facts, flagging them or re-running the adaptation to ensure absolute accuracy.

## Why
To ensure the resume's absolute integrity. If the AI fabricates experience, the candidate will fail the interview. Protecting the truth is a non-negotiable core value of this product.

## How
- Design validation prompt templates that test specific criteria (truthfulness, consistency, keyword relevance).
- Build automated retry/repair loops if validation fails.

## When
Implemented fifth, providing the quality assurance and safety net for the completed application.
