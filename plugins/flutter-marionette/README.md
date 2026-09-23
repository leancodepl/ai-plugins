# flutter-marionette

LeanCode Flutter plugin for [Marionette MCP](https://github.com/leancodepl/marionette_mcp): runtime interaction with a live Flutter app via an MCP server, designed for AI-driven exploration, smoke verification, and UI debugging.

## Included assets

- `skills/flutter-marionette-usage/references/marionette.md`: Marionette vs Patrol, setup, binding, MCP tools, and runtime workflow
- `skills/flutter-marionette-usage/references/marionette-widget-config.md`: `MarionetteConfiguration`, custom widgets, and screenshot sizing
- `skills/flutter-marionette-usage/SKILL.md`: explains what this plugin does, when to use it, and which rule or skill to reach for next
- `.mcp.json`: shipped `marionette` MCP launcher for plugin installs in Claude Code

## Marionette vs Patrol

| | Marionette (this plugin) | Patrol ([`flutter-patrol`](https://github.com/leancodepl/patrol)) |
|---|---|---|
| Purpose | Runtime exploration, smoke verification | Deterministic E2E test suites |
| Runs against | Live `flutter run` debug session | `patrol develop` / `patrol test` |
| Test files | None; the agent drives the app | Dart test files in `patrol_test/` |
| Best for | Iterating on a feature, smoke after refactor | Regression-proof suites in CI |
| Build mode | Debug and profile | Debug + release |

Both plugins can coexist; they solve different parts of the AI-assisted testing workflow.

## Setup highlights

### 1. Install the MCP server

For the bundled plugin launcher, the lowest-friction setup is a global install:

```bash
dart pub global activate marionette_mcp
```

Alternative: add it as a dev dependency and run it with `dart run marionette_mcp` from your own client-specific MCP launcher.

### 2. Add the Flutter package to your app

```bash
flutter pub add marionette_flutter
```

### 3. Initialize the binding in `main.dart`

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

For apps with a custom design system, pass a `MarionetteConfiguration`. See `skills/flutter-marionette-usage/references/marionette-widget-config.md`.

Important: `MarionetteBinding` must be the only binding initialized in the process. If tests call `main()` in debug mode, avoid initializing Marionette in tests by checking `FLUTTER_TEST` or by using a separate test entrypoint.

### 4. Use the bundled MCP launcher

When this plugin is installed in Claude Code, it ships the `marionette` MCP server entry for you. You should not need to hand-edit MCP config just to enable Marionette.

The bundled launcher runs:

```json
{
  "mcpServers": {
    "marionette": {
      "command": "marionette_mcp",
      "args": []
    }
  }
}
```

That means the remaining manual setup is:

- install `marionette_mcp`
- add `marionette_flutter` to the app
- initialize `MarionetteBinding`

### 5. Manual fallback for other MCP clients

If your MCP client does not load plugin-bundled servers, copy the same launcher into that client's MCP config:

```json
{
  "mcpServers": {
    "marionette": {
      "command": "marionette_mcp",
      "args": []
    }
  }
}
```

### 6. Run and connect

```bash
flutter run
```

Copy the VM service URI from the run output (format: `ws://127.0.0.1:PORT/ws`). In your AI agent, call the `connect` tool with that URI.

`connect` checks that `marionette_mcp` and the app's `marionette_flutter` are the same version and fails if they differ. After upgrading `marionette_flutter`, activate the matching server too: `dart pub global activate marionette_mcp <version>`.

## Log collection

`get_logs` needs a `LogCollector`, passed as `MarionetteConfiguration(logCollector: ...)`. The collectors for the two common logging packages ship as separate packages:

| Your app logs with | Add to the app | Collector |
| --- | --- | --- |
| [`logging`](https://pub.dev/packages/logging) | `flutter pub add marionette_logging` | `LoggingLogCollector()` from `package:marionette_logging/marionette_logging.dart` |
| [`logger`](https://pub.dev/packages/logger) | `flutter pub add marionette_logger` | `LoggerLogCollector()` from `package:marionette_logger/marionette_logger.dart`. Also add it to your `Logger`'s outputs. |
| Anything else | Nothing, it ships in `marionette_flutter` | `PrintLogCollector()`. Call `collector.addLog(message)` wherever your logs flow. |

```dart
MarionetteBinding.ensureInitialized(
  MarionetteConfiguration(logCollector: LoggingLogCollector()),
);
```

`get_logs` returns the logs since app start or the last hot reload. Without a collector it explains how to enable one. Upstream docs: [Log Collection](https://github.com/leancodepl/marionette_mcp/blob/main/docs/logging.md).

## Build-mode constraint

Marionette relies on Flutter's VM Service, so it works in debug and profile builds, not in release builds. The `kDebugMode` guard above enables it in debug only; use `!kReleaseMode` to include profile builds. Use Patrol for release-mode automation.

## Example usage

- `/flutter-marionette-usage`: get a short explanation of what this plugin does, how to use it, and example next requests

## Related plugins

- [`flutter-patrol`](https://github.com/leancodepl/patrol): deterministic E2E testing with Patrol MCP
- [`flutter-ui`](../flutter-ui/): design-system guidance (relevant when configuring `isInteractiveWidget` for custom widgets)
