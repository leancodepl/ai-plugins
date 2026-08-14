# flutter-cubit-utils

LeanCode Flutter plugin for `leancode_cubit_utils` and `leancode_cubit_utils_cqrs`.

Contains the cubit-utils rules and examples:

- `skills/flutter-cubit-utils-usage/references/cubit-utils.md` — base classes and when to use them
- `skills/flutter-cubit-utils-usage/references/cubit-list.md` — paginated list cubit example
- `skills/flutter-cubit-utils-usage/references/cubit-details.md` — single-object/details cubit example
- `skills/flutter-cubit-utils-usage/references/cubit-action.md` — action/command cubit examples
- `skills/flutter-cubit-utils-usage/SKILL.md` — entry point for this plugin

## Agent

- `@flutter-cubit-utils` — agent for substantial, multi-step base-cubit work (choosing a base class, scaffolding list/details/action cubits, migrating ad-hoc cubits). It preloads the `flutter-cubit-utils-usage` skill and applies LeanCode conventions end to end. For quick inline questions, use `/flutter-cubit-utils-usage` instead.

## Example usage

- `/flutter-cubit-utils-usage` — get a short explanation of what this plugin does and which asset to use next

## What this plugin is NOT about

General BLoC/Cubit fundamentals and widget-side presentation patterns live in [`flutter-bloc`](../flutter-bloc/). CQRS contract design and data-access boundaries live in [`flutter-cqrs`](../flutter-cqrs/).

## Related plugins

- [`flutter-bloc`](../flutter-bloc/) — general BLoC/Cubit fundamentals
- [`flutter-cqrs`](../flutter-cqrs/) — CQRS contracts and data-access guidance
