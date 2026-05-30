---
name: poe2arb-workflow
description: Handle Flutter localization workflows backed by POEditor and poe2arb. Use when adding translations, syncing ARB files with POEditor, or when the user mentions poe2arb, POEditor, seed, or flutter gen-l10n.
---

# POEditor and poe2arb Workflow

Keep the localization workflow aligned with the shared localization rules.

## What to keep true

- Keep translations in package-specific `app_<language_code>.arb` files.
- Never hardcode translations.
- Use translations only in the presentation layer, inside widgets.
- Access translations as `final s = l10n(context);`
- If other layers need localized text, map message models to localized values in the presentation layer.

## Workflow when adding translations

1. Add the required translation to all relevant `app_<language_code>.arb` files.
2. Run `flutter gen-l10n`.
3. Run `poe2arb seed` to upload the added translation to POEditor.

## Output expectations

- If asked to change translations, update the ARB files first.
- Clearly say whether `flutter gen-l10n` and `poe2arb seed` were run or still need to be run.
