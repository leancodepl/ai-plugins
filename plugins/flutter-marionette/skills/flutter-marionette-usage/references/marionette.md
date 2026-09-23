# Marionette MCP

Marionette is a LeanCode MCP server that lets an AI agent interact with a **running Flutter app at runtime** — inspect the widget tree, tap elements, enter text, scroll, take screenshots, read logs, and trigger hot reload.

Think of it as "Playwright for Flutter, for AI agents."

## Marionette vs Patrol

Both are LeanCode testing tools, but they solve different problems:

| | Marionette | Patrol |
|---|---|---|
| Purpose | Runtime exploration, smoke verification | Deterministic E2E test suites |
| Runs against | A live `flutter run` debug session | `patrol develop` or `patrol test` runner |
| Test files | None — agent drives the app interactively | Dart test files in `patrol_test/` |
| Best for | "Does my new feature work?", smoke after refactor, debugging unresponsive UI | Regression-proof test suites in CI |
| Build mode | Debug and profile, not release (requires VM Service) | Debug + release (via Patrol runner) |

Rules of thumb:
- Iterating on a new feature and want the agent to click around? **Marionette.**
- Writing a repeatable E2E test that runs in CI? **Patrol** (see `flutter-patrol` plugin).
- Want quick visual verification after a change without writing tests? **Marionette.**

## App preparation

Prepare the Flutter app before trying to drive it:

- Add `marionette_flutter` to the app.
- Install the MCP server either as a global tool (`dart pub global activate marionette_mcp`) or as a dev dependency (`dart pub add dev:marionette_mcp`) and run it with `dart run marionette_mcp`.

## Available MCP tools

When the `marionette` MCP server is configured and connected, use these tools:

- `connect` / `disconnect`: manage the VM service connection. Call `connect` first, passing the `ws://127.0.0.1:PORT/ws` URI printed by `flutter run`. It fails when the `marionette_mcp` server and the app's `marionette_flutter` are different versions; align them.
- `get_interactive_elements`: list the visible interactive elements with their type, text, key and Semantics identifier. **Always call this before acting** to see what is available and to get stable selectors.
- `tap`, `double_tap`, `long_press`, `secondary_tap` (desktop only), `pinch_zoom`: gestures on one element, matched by exactly one of `key`, `identifier`, `text`, `type` or `coordinates`. Prefer `key` (a `ValueKey<String>`), then `identifier` (the Semantics identifier). Tapping a text field focuses it.
- `swipe`: element-based (`key`, `identifier` or `text`, plus `direction`) or coordinate-based. Use it for `PageView`, `Dismissible`, `Drawer` and sliders.
- `scroll_to`: scroll until an element matching `key`, `identifier` or `text` is visible.
- `press_back_button`: the system back action (Android back, iOS swipe-back).
- `enter_text`: type into a field. Pass `input` and exactly one of `key`, `identifier`, or `focused_element: true` after tapping the field. There is no text matcher.
- `press_key`: a real key event on the focused element (`enter`, `tab`, `escape`, arrows, characters, optional `modifiers`). On iOS and Android, change a field's value with `enter_text` instead.
- `take_screenshots`: capture all active views as base64 PNGs. Use to verify visual state after an action or to debug why something is not found.
- `get_logs`: app logs since start or the last hot reload. Needs a `LogCollector` (see "Log collection"). Use when an action silently fails or state is unclear.
- `hot_reload`: apply code changes without losing app state. `hot_restart`: restart from `main()` and reset all state.
- `set_device_config`: override text scale, bold text and brightness. The app must mount `MarionetteDeviceConfig` first.
- `list_custom_extensions` / `call_custom_extension`: app-specific extensions registered with `registerMarionetteExtension`.

Parameters for every tool are in the upstream [MCP Tools](https://github.com/leancodepl/marionette_mcp/blob/main/docs/mcp-tools.md) doc.

## Binding initialization

Initialize `MarionetteBinding` in `main.dart` under `kDebugMode` so it compiles out of release builds:

```dart
import 'package:flutter/foundation.dart';
import 'package:flutter/widgets.dart';
import 'package:marionette_flutter/marionette_flutter.dart';

void main() {
  if (kDebugMode) {
    MarionetteBinding.ensureInitialized();
  } else {
    WidgetsFlutterBinding.ensureInitialized();
  }
  runApp(const MyApp());
}
```

For apps with a custom design system, pass a `MarionetteConfiguration` — see `marionette-widget-config.md`.

`MarionetteBinding` must be the only binding initialized in the process. If tests call `main()` while `kDebugMode` is true, avoid initializing Marionette in tests by checking `FLUTTER_TEST` or by using a separate test entrypoint.

## Log collection

`get_logs` works only when a `LogCollector` is passed as `MarionetteConfiguration(logCollector: ...)`. The `logging` and `logger` collectors are separate packages, not part of `marionette_flutter`:

- `logging` package: `flutter pub add marionette_logging`, then `LoggingLogCollector()` from `package:marionette_logging/marionette_logging.dart`.
- `logger` package: `flutter pub add marionette_logger`, then `LoggerLogCollector()` from `package:marionette_logger/marionette_logger.dart`. It is also a `LogOutput`, so add the same instance to the `Logger`'s outputs.
- Anything else: `PrintLogCollector()` from `marionette_flutter`, fed with `collector.addLog(message)`.
- If no collector is configured, `get_logs` explains how to enable it.

```dart
MarionetteBinding.ensureInitialized(
  MarionetteConfiguration(logCollector: LoggingLogCollector()),
);
```

## Workflow: driving the app as an agent

1. Start the app: `flutter run` (take note of the VM service URI).
2. Call `connect` with the URI.
3. Call `get_interactive_elements` to see what is on screen.
4. Act: `tap`, `enter_text`, `scroll_to`, `swipe`, `press_back_button` and the other gestures.
5. Call `take_screenshots` to verify visual state (and when the element list looks wrong).
6. Use `get_logs` if an action has no visible effect.
7. Call `hot_reload` after a code change to re-test without losing state, or `hot_restart` when the change needs a fresh start.
8. Call `disconnect` when finished.

## Best-effort caveat

Marionette simulates gestures best-effort. Results may vary with platform differences, overlays, and custom widget implementations. If interactions are flaky:

- Expose clearer widget keys.
- Simplify hit targets (avoid deeply nested `GestureDetector`s).
- Add `isInteractiveWidget` / `extractText` hooks for custom widgets — see `marionette-widget-config.md`.

## The agent does not understand your app flow

Marionette observes the UI but does not automatically know your product's flows, naming conventions, or edge cases. When prompting the AI:

- Describe expected screen names, widget keys, and labels.
- State preconditions (e.g. "assume user is already logged in").
- State the interaction goal in one sentence.

## Build-mode constraint

Marionette relies on Flutter's VM Service, so it works in debug and profile builds, not in release builds. The `kDebugMode` guard above limits it to debug; `!kReleaseMode` also covers profile. For release-mode automation use Patrol instead.
