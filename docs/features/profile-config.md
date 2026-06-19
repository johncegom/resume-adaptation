# Epic: Profile & Setting Configurations

## What
Saves default paths, candidate details, and AI API credentials in a localized user configuration directory.

## Why
Improves usability so the user doesn't need to specify their Gemini API key or path to their default resume on every run.

## How
- Read and write configuration in a cross-platform directory (e.g. `~/.config/resume-adaptation/config.json`).
- Provide CLI command `resume-adapt config` to view and set variables.
