# flutter-forms

LeanCode Flutter forms plugin for Claude Code.

Contains the forms rule for `leancode_forms`, naming, validation behavior, and form cubit provisioning.

## Included assets

- `skills/flutter-forms-usage/references/forms.md` — forms conventions
- `skills/flutter-forms-usage/SKILL.md` — entry point for this plugin

## Agent

- `@flutter-forms` — agent for substantial, multi-step forms work (`leancode_forms` flows, validation behavior, form-cubit provisioning). It preloads the `flutter-forms-usage` skill and applies LeanCode conventions end to end. For quick inline questions, use `/flutter-forms-usage` instead.

## Example usage

- `/flutter-forms-usage` — get a short explanation of what this plugin does and which asset to use next

## What this plugin is NOT about

Page-root provider wiring lives in [`flutter-di`](../flutter-di/). General cubit/bloc logic outside form-specific rules lives in [`flutter-bloc`](../flutter-bloc/).

## Related plugins

- [`flutter-di`](../flutter-di/) — page-root dependency provisioning and scope decisions
- [`flutter-bloc`](../flutter-bloc/) — cubit and bloc state-management patterns
