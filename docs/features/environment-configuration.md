# Feature: Environment Configuration Support

## What
Support loading environment variables from a local `.env` file during development and provide a `.env.example` template for configuration settings.

## Why
Currently, the application relies on variables like GEMINI_API_KEY being pre-set in the environment. Adding support for `.env` files allows developers to easily manage credentials and configurations locally without manual exports.

## How
- Install github.com/joho/godotenv dependency.
- Create a .env.example template in the project root containing keys like GEMINI_API_KEY, LOG_LEVEL, and LOG_FORMAT.
- Initialize godotenv.Load() inside cmd/cli/main.go entry point.
- Update README.md and related docs to describe the configuration setup.

## When
This will be completed immediately as a prerequisite for smooth local developer onboarding and testing.
