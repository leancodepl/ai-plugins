# Changelog

## 0.2.0

- The Marionette skills now come from the `marionette_flutter` package, through the new `flutter-marionette-skills` entry that this plugin depends on: `marionette-flutter-setup` (prepare an app), `marionette-flutter-drive-app` (drive it and write the session report) and `marionette-flutter-to-patrol` (turn a verified session into Patrol tests). Installing `flutter-marionette` installs them, and they follow Marionette's `main` branch.
- `references/marionette.md` and `references/marionette-widget-config.md` are removed; `marionette-flutter-setup` and `marionette-flutter-drive-app` cover the same ground for the current Marionette API.
- `flutter-marionette-usage` explains the plugin and routes each task to one of the three skills.
- The README is shorter: install, what you get, setup, the MCP launcher and the Patrol flow. Log collection, device config and custom design systems point to the Marionette documentation.

## 0.1.1

- Log collection: `LoggingLogCollector` and `LoggerLogCollector` come from the separate `marionette_logging` and `marionette_logger` packages, not from `marionette_flutter`. The README and `references/marionette.md` now say which package to add.
- `references/marionette.md`: the tool list matches Marionette MCP 0.6.0, with selectors for `tap` and `enter_text`, and the gesture, key, restart, device-config and custom-extension tools.
- Build mode: Marionette works in debug and profile builds, not only debug.
- `connect` fails when `marionette_mcp` and `marionette_flutter` versions differ; the README says how to align them.
- `flutter-marionette-usage` description names the runtime actions (tap, type, scroll, screenshots, smoke tests) so the skill loads for them.

## 0.1.0

- Initial `flutter-marionette` plugin in the public marketplace.
