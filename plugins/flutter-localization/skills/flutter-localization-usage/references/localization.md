# Localization

All translations are stored in package specific `app_<language_code>.arb` files. `POEditor` with `poe2arb` tools are used to handle them.

- Never hardcode translations explicitly. Always add required translations to all `app_<language_code>.arb` files
- Always use translations in presentation layer (inside `Widgets`). Access translations as `final s = l10n(context);`
- Never use translations outside presentation layer. If needed map message models to localized value in presentation layer
- Always run `flutter gen-l10n` if any translation was added
- Always run `poe2arb seed` if any translation was added to upload them to `POEditor`
