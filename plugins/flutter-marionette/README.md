# flutter-marionette

LeanCode Flutter plugin for [Marionette MCP](https://github.com/leancodepl/marionette_mcp): runtime interaction with a live Flutter app via an MCP server, designed for AI-driven exploration, smoke verification, and UI debugging.

## Install

```text
/plugin install flutter-marionette@leancode-ai-plugins
```

This also installs `flutter-marionette-skills`: Marionette's Agent Skills, read straight from the [`marionette_flutter` package](https://github.com/leancodepl/marionette_mcp/tree/main/packages/marionette_flutter/skills). They follow Marionette's `main` branch.

## What you get

| Part | Comes from | Does |
| --- | --- | --- |
| `/flutter-marionette-usage` | this plugin | Explains the plugin and picks the skill for a task |
| `marionette` MCP server | this plugin (`.mcp.json`) | Runs `marionette_mcp`, so you do not edit MCP config |
| `marionette-flutter-setup` | `flutter-marionette-skills` | Prepares an app: binding, custom design system, logs, device config, versions |
| `marionette-flutter-drive-app` | `flutter-marionette-skills` | Drives the app with the MCP tools or the CLI, and writes the session report |
| `marionette-flutter-to-patrol` | `flutter-marionette-skills` | Turns a verified session into Patrol tests |

## Marionette vs Patrol

| | Marionette (this plugin) | Patrol ([`flutter-patrol`](https://github.com/leancodepl/patrol)) |
|---|---|---|
| Purpose | Runtime exploration, smoke verification | Deterministic E2E test suites |
| Runs against | Live `flutter run` debug session | `patrol develop` / `patrol test` |
| Test files | None; the agent drives the app | Dart test files in `patrol_test/` |
| Best for | Iterating on a feature, smoke after refactor | Regression-proof suites in CI |
| Build mode | Debug and profile | Debug + release |

Both plugins can coexist; they solve different parts of the AI-assisted testing workflow.

## Setup

1. Install the MCP server that the bundled launcher runs:

   ```bash
   dart pub global activate marionette_mcp
   ```

2. Ask Claude to "add Marionette to this app". `marionette-flutter-setup` adds `marionette_flutter`, the debug-only `MarionetteBinding`, and the configuration a custom design system needs. To do it by hand, follow [Start](https://marionette.leancode.co/docs/).

3. Run the app and let the agent connect:

   ```bash
   flutter run
   ```

   Paste the VM service URI from the output (`ws://127.0.0.1:PORT/TOKEN=/ws`) into the prompt, or ask the agent to run the app itself.

`connect` fails when `marionette_mcp` and the app's `marionette_flutter` are different versions. After upgrading `marionette_flutter`, activate the matching server: `dart pub global activate marionette_mcp <version>`.

### Other MCP clients

The plugin's `.mcp.json` launches `marionette_mcp` in Claude Code. If your MCP client does not load plugin-bundled servers, copy the same entry into its MCP config:

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

## From a Marionette run to a Patrol test

`marionette-flutter-to-patrol` turns a run that Marionette verified into Patrol tests, so the flow stays checked in CI. It builds on the `patrol-write-test` and `patrol-test-architecture` skills, so install `flutter-patrol` too:

```text
/plugin install flutter-patrol@leancode-ai-plugins
```

Turn session reports on in the app's binding (`MarionetteConfiguration(enableSessionReports: true)`; see [Reports](https://marionette.leancode.co/docs/reports/)), so each run leaves `.marionette/sessions/<name>/`. After the run, ask:

```text
/marionette-flutter-to-patrol Write Patrol tests from .marionette/sessions/checkout-20260922T1801.
Spec: PROJ-123. Device: emulator-5554. Seed reset: backend/README.md.
```

Have a free device and, when the flow reads backend data, a local backend with a resettable seed.

## Example usage

- `/flutter-marionette-usage`: what this plugin does and which skill fits a task
- "Add Marionette to this app."
- "Use Marionette to smoke-test the onboarding flow."
- `/marionette-flutter-to-patrol`: after a Marionette run with session reports on, turn the session into Patrol tests

The [Marionette documentation](https://marionette.leancode.co/docs/) covers custom design systems, log collection, device config and the CLI.

## Related plugins

- [`flutter-patrol`](https://github.com/leancodepl/patrol): deterministic E2E testing with Patrol MCP
- [`flutter-ui`](../flutter-ui/): design-system guidance (relevant when teaching Marionette your custom widgets)
