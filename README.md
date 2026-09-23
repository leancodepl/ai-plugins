# LeanCode AI Plugins

Claude Code plugins for Flutter: the conventions we use in production, and tools that let Claude see your running app. From the team behind [Patrol](https://github.com/leancodepl/patrol), [Marionette MCP](https://github.com/leancodepl/marionette_mcp) and [advanced_forms](https://github.com/leancodepl/advanced_forms).

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Works with Claude Code](https://img.shields.io/badge/works%20with-Claude%20Code-d97757.svg)](https://code.claude.com/docs/en/overview)

## Before and after

[`flutter-read-logs`](plugins/flutter-read-logs/) gives Claude the logs of your last `flutter run`. Same code, same model (Claude Sonnet 5), same question:

> Tapping Cancel on the second booking does nothing. Why?

|                | Without the plugin                                          | With `/read-logs`                                                                                      |
| -------------- | ----------------------------------------------------------- | ------------------------------------------------------------------------------------------------------ |
| **Cause**      | Guessed from the code: "404/409/500, or a network hiccup"   | Read from the run: `409 cancellation_window_closed`, because the booking starts in under 24 hours      |
| **Fix**        | A generic "Failed to cancel booking." snackbar              | Maps the 409 to the screen's error state, and suggests the server's reason: "Bookings can't be cancelled less than 24 hours before they start." |
| **Tool calls** | 17                                                          | 6                                                                                                      |
| **Cost**       | $0.50                                                       | $0.17                                                                                                  |

<sub>Quotes are from real runs on 2026-09-23: a small Flutter app on macOS, and a local API outside the app's repository. An integration test tapped Cancel, and `/read-logs` read that run's log file. Writing the log file takes one editor setting, and `/read-logs` walks you through it.</sub>

## Install

```
/plugin marketplace add leancodepl/ai-plugins
/plugin install flutter-read-logs@leancode-ai-plugins
```

Then run `/reload-plugins`. Each plugin installs on its own as `<plugin-name>@leancode-ai-plugins`, so take only the ones you agree with. The **Discover** tab in `/plugin` lists them all.

New here? Install [`lean-core`](plugins/lean-core/) and run `/lean-core-usage`. Every plugin in this repository has a `/<plugin-name>-usage` skill, such as `/flutter-bloc-usage`, that explains what it covers.

<details>
<summary>Team setup and the git URL</summary>

To have Claude Code offer the marketplace to everyone on a project, add it to the project's `.claude/settings.json`:

```json
{
  "extraKnownMarketplaces": {
    "leancode-ai-plugins": {
      "source": {
        "source": "github",
        "repo": "leancodepl/ai-plugins"
      }
    }
  }
}
```

The git URL works as a source too:

```
/plugin marketplace add https://github.com/leancodepl/ai-plugins.git
```

See [Discover and install plugins](https://code.claude.com/docs/en/discover-plugins) in the Claude Code docs.

</details>

## Supported clients

**Claude Code** only.

## Available plugins

### Getting started

| Plugin | Covers |
| ------ | ------ |
| [`lean-core`](plugins/lean-core/) | `/lean-core-usage` explains the marketplace; `/lean-contribute` walks you through a PR |

### Project foundations

| Plugin | Covers |
| ------ | ------ |
| [`flutter-leancode-architecture`](plugins/flutter-leancode-architecture/) | Project structure, error handling and logging; architecture review and feature scaffolding |
| [`flutter-di`](plugins/flutter-di/) | `provider`-based dependency injection, page-root providers, `GlobalProviders`, async initialization |
| [`flutter-navigation`](plugins/flutter-navigation/) | `auto_route` and `go_router`: typed routes, guards, route trees, deep links |
| [`flutter-analytics`](plugins/flutter-analytics/) | Analytics IDs and page and tap tracking; scaffolds IDs and reviews coverage |
| [`flutter-localization`](plugins/flutter-localization/) | ARB files, POEditor, `poe2arb` and `l10n(context)` |

### State and data

| Plugin | Covers |
| ------ | ------ |
| [`flutter-bloc`](plugins/flutter-bloc/) | BLoC and Cubit conventions, state modeling, side effects with `bloc_presentation`, `flutter_hooks` |
| [`flutter-cubit-utils`](plugins/flutter-cubit-utils/) | `QueryCubit`, `PaginatedQueryCubit` and `RequestCubit`, with recipes for lists, details and actions |
| [`flutter-cqrs`](plugins/flutter-cqrs/) | CQRS contracts, repositories, `cqrs.run` and `cqrs.get`, CQRS-backed cubits |
| [`flutter-forms`](https://github.com/leancodepl/advanced_forms) | `advanced_forms` controllers, sync, async and cross-field validation, submit handling |

### UI and verification

| Plugin | Covers | Setup |
| ------ | ------ | ----- |
| [`flutter-ui`](plugins/flutter-ui/) | Design-system-driven UI, loading and error states, localized text, implementation checklists | None |
| [`flutter-patrol`](https://github.com/leancodepl/patrol) | Setting up Patrol, writing E2E tests, test architecture and key conventions | Patrol CLI and Patrol MCP |
| [`flutter-marionette`](plugins/flutter-marionette/) | Driving a live debug app through Marionette MCP to explore, smoke-check and debug UI | Marionette MCP and an app-side binding |
| [`flutter-read-logs`](plugins/flutter-read-logs/) | `/read-logs` gives Claude the logs of your last `flutter run` | One editor setting |

`flutter-forms` and `flutter-patrol` install straight from the repositories of the packages they cover, so they stay in step with each release. Each plugin's `README.md` covers its setup.

### Judgment and review

| Plugin | Covers | Setup |
| ------ | ------ | ----- |
| [`lean-jev`](plugins/lean-jev/) | Hands a judgment call to Jev, TypeSafe's System One model, and frames it the way Jev needs it framed: typed questions, numbers instead of prose, cheap enough to judge every item in a list | A TypeSafe API key |

## Contributing

Fixes and new Flutter plugins are welcome. CI runs every check, so you need nothing installed locally. [`CONTRIBUTING.md`](CONTRIBUTING.md) has the steps, and [`AGENTS.md`](AGENTS.md) has the plugin conventions.

## License

[MIT](LICENSE)
