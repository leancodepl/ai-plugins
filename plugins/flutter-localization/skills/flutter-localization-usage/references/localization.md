# Localization

All translations are stored in package specific `app_<language_code>.arb` files. `POEditor` with `poe2arb` tools are used to handle them.

- Never hardcode translations explicitly. Always add required translations to all `app_<language_code>.arb` files
- Always use translations in presentation layer (inside `Widgets`). Access translations as `final s = l10n(context);`
- Never use translations outside presentation layer. If needed map message models to localized value in presentation layer
- Always run `flutter gen-l10n` if any translation was added
- Always run `poe2arb poe` before editing ARB files so they hold the latest translations from `POEditor`
- Always run `poe2arb upload` if any translation was added, changed, or removed to sync them to `POEditor`. Run `poe2arb upload --dry-run` first and never pass `--force` without the user's confirmation
- Use `poe2arb seed` only to set up a fresh `POEditor` project
