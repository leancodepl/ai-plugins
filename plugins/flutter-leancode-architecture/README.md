# flutter-leancode-architecture

LeanCode Flutter architecture standards plugin for Claude Code.

Contains the architecture rules and skills for project structure, error handling, and logging.

- `skills/flutter-leancode-architecture-usage/references/project-structure.md` — project structure conventions
- `skills/flutter-leancode-architecture-usage/references/error-handling.md` — error handling conventions
- `skills/flutter-leancode-architecture-usage/references/logging.md` — logging conventions
- `skills/flutter-leancode-architecture-usage/SKILL.md` — entry point for this plugin
- `skills/review-leancode-arch/SKILL.md` — architecture review workflow
- `skills/scaffold-feature/SKILL.md` — feature scaffolding workflow

## What this plugin is NOT about

Dependency injection, navigation, analytics, localization, and state management live in dedicated sibling plugins.

## Agent

- `@flutter-leancode-architecture` — agent for substantial, multi-step architecture work (feature scaffolding, architecture review, project-structure/error-handling/logging conventions). It preloads the `flutter-leancode-architecture-usage` skill and applies LeanCode conventions end to end. For quick inline questions, use `/flutter-leancode-architecture-usage` instead.

## Example usage

- `/flutter-leancode-architecture-usage` — get a short explanation of what this plugin does and which asset to use next
- `/review-leancode-arch lib/features/booking/`
- `/scaffold-feature booking`
