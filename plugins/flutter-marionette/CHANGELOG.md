# Changelog

## 0.1.1

- Log collection: `LoggingLogCollector` and `LoggerLogCollector` come from the separate `marionette_logging` and `marionette_logger` packages, not from `marionette_flutter`. The README and `references/marionette.md` now say which package to add.
- `references/marionette.md`: the tool list matches Marionette MCP 0.6.0, with selectors for `tap` and `enter_text`, and the gesture, key, restart, device-config and custom-extension tools.
- Build mode: Marionette works in debug and profile builds, not only debug.
- `connect` fails when `marionette_mcp` and `marionette_flutter` versions differ; the README says how to align them.
- `flutter-marionette-usage` description names the runtime actions (tap, type, scroll, screenshots, smoke tests) so the skill loads for them.

## 0.1.0

- Initial `flutter-marionette` plugin in the public marketplace.
