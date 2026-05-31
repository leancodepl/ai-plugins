# flutter-localization

LeanCode Flutter localization plugin for Claude Code.

Contains the localization rule and workflow:

## Included assets

- `skills/flutter-localization-usage/references/localization.md` — localization conventions
- `skills/flutter-localization-usage/SKILL.md` — entry point for this plugin
- `skills/poe2arb-workflow/SKILL.md` — workflow for ARB updates, `flutter gen-l10n`, and `poe2arb seed`

## Agent

- `@flutter-localization` — agent for substantial, multi-step localization work (ARB translations, `l10n(context)` usage, POEditor/`poe2arb` sync). It preloads the `flutter-localization-usage` skill and applies LeanCode conventions end to end. For quick inline questions, use `/flutter-localization-usage` instead.

## Example usage

- `/flutter-localization-usage` — get a short explanation of what this plugin does and which asset to use next
- `/poe2arb-workflow`
