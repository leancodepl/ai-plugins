# flutter-di

LeanCode Flutter dependency injection plugin for Claude Code.

Contains the dependency injection rule:

- `skills/flutter-di-usage/references/dependency-injection.md` — `provider`-based DI conventions
- `skills/flutter-di-usage/SKILL.md` — entry point for this plugin

## Agent

- `@flutter-di` — agent for substantial, multi-step dependency-injection work (provider wiring, scope decisions, async initialization). It preloads the `flutter-di-usage` skill and applies LeanCode conventions end to end. For quick inline questions, use `/flutter-di-usage` instead.

## Example usage

- `/flutter-di-usage` — get a short explanation of what this plugin does and which asset to use next

## What this plugin is NOT about

General cubit/bloc logic after dependencies are provided lives in [`flutter-bloc`](../flutter-bloc/).

## Related plugins

- [`flutter-bloc`](../flutter-bloc/) — BLoC/Cubit state management after dependencies are provided
