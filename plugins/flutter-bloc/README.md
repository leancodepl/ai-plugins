# flutter-bloc

LeanCode Flutter BLoC plugin for Claude Code.

Covers LeanCode BLoC/Cubit fundamentals, including `flutter_bloc`, `bloc_presentation`, `flutter_hooks`, and general cubit/state conventions.

## Included assets

- `skills/flutter-bloc-usage/references/state-management.md` — main BLoC/Cubit fundamentals
- `skills/flutter-bloc-usage/SKILL.md` — entry point for this plugin
- `skills/flutter-context-watch-instead-of-top-builder/SKILL.md` — focused refactor workflow for replacing top-level `BlocBuilder` with `context.watch`

## Example usage

- `/flutter-bloc-usage` — get a short explanation of what this plugin does and which asset to use next

## What this plugin is NOT about

Base-class decisions and canonical cubit recipes (`QueryCubit`, `PaginatedQueryCubit`, `RequestCubit`, etc.) live in [`flutter-cubit-utils`](../flutter-cubit-utils/).

## Related plugins

- [`flutter-cubit-utils`](../flutter-cubit-utils/) — base classes and canonical cubit recipes
