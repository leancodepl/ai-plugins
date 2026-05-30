# flutter-analytics

LeanCode Flutter analytics plugin for Claude Code.

Contains the analytics rules and helper skills for feature-level IDs, page IDs, and clickable-element IDs.

## Included assets

- `skills/flutter-analytics-usage/references/analytics.md` - conventions for `*_ids.dart`
- `skills/flutter-analytics-usage/references/analytics-usage.md` - page and clickable-element usage patterns
- `skills/flutter-analytics-usage/SKILL.md` - entry point for this plugin
- `skills/scaffold-analytics-ids/SKILL.md` - create or extend a feature IDs file
- `skills/review-analytics-coverage/SKILL.md` - audit a feature or page for missing analytics coverage

## Example usage

- `/flutter-analytics-usage` - get a short explanation of what this plugin does and which asset to use next
- `/scaffold-analytics-ids booking`
- `/scaffold-analytics-ids booking --page details --action submit`
- `/review-analytics-coverage lib/features/booking/`

## What this plugin is NOT about

Project structure, error handling, and logging live in [`flutter-leancode-architecture`](../flutter-leancode-architecture/).
