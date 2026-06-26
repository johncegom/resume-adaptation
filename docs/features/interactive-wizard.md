# Feature: Interactive UI Adaptation Wizard

## What
A polished, keyboard-driven terminal user interface (TUI) that guides the user through the resume adaptation process. It facilitates selecting the original resume, pasting or providing the job description, entering configuration parameters, and reviewing/editing the generated resume.

## Why
A beautiful, responsive terminal wizard enhances the overall UX, making the tool feel premium and easy to use, supporting high user satisfaction.

## How
- Implement the wizard using Charmbracelet libraries: Lipgloss for styling, Huh for forms, and Bubble Tea for interactive navigation.
- Decouple TUI views from core adaptation and parsing logic.

## When
Implemented third, once the core engine is working, to wrap it in a user-facing terminal interface.
