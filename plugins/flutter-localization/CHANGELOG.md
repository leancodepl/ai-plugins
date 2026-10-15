# Changelog

## 0.2.0

- `poe2arb-workflow` and the localization conventions (`flutter-localization-usage/references/localization.md`) use `poe2arb upload` to sync ARB changes to POEditor, replacing `poe2arb seed`, which is now only for setting up a fresh project.
- The workflow pulls with `poe2arb poe` before editing ARBs, previews with `poe2arb upload --dry-run`, requires poe2arb 2.6.0 or newer, and passes `--force` only after the user confirms the listed term deletions.
