# LeanCode AI Plugins

Claude Code plugins for Flutter: the conventions LeanCode uses in production apps, and tools that let Claude look at your running app.

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Works with Claude Code](https://img.shields.io/badge/works%20with-Claude%20Code-d97757.svg)](https://code.claude.com/docs/en/overview)

We are the team behind [Patrol](https://github.com/leancodepl/patrol), [Marionette MCP](https://github.com/leancodepl/marionette_mcp) and [advanced_forms](https://github.com/leancodepl/advanced_forms). These plugins are the conventions we give our own agents. They are opinionated — our choices, not a survey of the ecosystem — and every plugin installs on its own, so take the ones you agree with.

## See it work

The app misbehaves at runtime, and the code looks fine. With [`flutter-read-logs`](plugins/flutter-read-logs/), Claude reads the run you just made before it answers.

> Tapping Cancel on the second booking does nothing. Why?

Same code and same question, asked in two Claude Code sessions (Claude Sonnet 5):

**Without the plugin**, Claude can only read the code. It finds that the cubit swallows every error except a 400, and has to guess what actually came back:

> So when cancelling the second booking fails for any reason other than a 400 (e.g. it returns 404/409/500, or a network hiccup), it falls into the generic `catch (e)` in the cubit, which only logs a warning [...]

Its fix is a generic "Failed to cancel booking." snackbar.

**With `/read-logs`**, Claude reads the run first:

> the API rejects the cancellation with `409 cancellation_window_closed` (b-2 — "Haircut, Studio 7" — starts 2026-09-24T02:43, less than 24h out), and the cubit logs `WARNING BookingsCubit: Cancel b-2 failed — ApiException(409)` and stops there — no `emit`.

Its fix maps the 409 to the error the screen already shows, and suggests surfacing the server's own reason: "Bookings can't be cancelled less than 24 hours before they start." It also got there in 6 tool calls instead of 17.

<sub>Both answers are quoted from real runs on 2026-09-23: a small Flutter app on macOS, talking to a local API outside the app's repository. An integration test tapped Cancel, and the run was recorded to the log file the plugin reads. Setting up that log file is a one-time editor setting, and `/read-logs` walks you through it.</sub>

## Install

Add the marketplace once, then install the plugins you want:

```
/plugin marketplace add leancodepl/ai-plugins
/plugin install flutter-read-logs@leancode-ai-plugins
```

Every plugin installs as `<plugin-name>@leancode-ai-plugins`. Run `/reload-plugins` to activate, or browse the catalog with `/plugin` (the **Discover** tab).

Not sure where to start? Install [`lean-core`](plugins/lean-core/) and run `/lean-core-usage` for a tour. Plugins stored in this repository also ship a `/<plugin-name>-usage` skill, such as `/flutter-bloc-usage`, that explains what the plugin covers and suggests prompts to try.

<details>
<summary>Team setup and other install variants</summary>

To have Claude Code prompt collaborators to install the marketplace automatically, add it to your project's `.claude/settings.json`:

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

The full git URL works as a marketplace source too:

```
/plugin marketplace add https://github.com/leancodepl/ai-plugins.git
```

For the full reference, see Claude Code's [Discover and install plugins](https://code.claude.com/docs/en/discover-plugins) docs.

</details>

## Supported clients

**Claude Code** only. Plugins ship as skills with reference material that loads when Claude needs it; a few also bundle an MCP server.

## Available plugins

### Getting started

- [`lean-core`](plugins/lean-core/) - marketplace entry point: `/lean-core-usage` explains what's available, `/lean-contribute` walks you through opening a PR

### Project foundations

- [`flutter-leancode-architecture`](plugins/flutter-leancode-architecture/) - project structure, error handling, logging, plus architecture review and feature scaffolding skills
- [`flutter-di`](plugins/flutter-di/) - `provider`-based dependency injection, page-root providers, `GlobalProviders`, and async initialization patterns
- [`flutter-navigation`](plugins/flutter-navigation/) - `auto_route` and `go_router`, typed routes, route guards, route tree organization, and deep links
- [`flutter-analytics`](plugins/flutter-analytics/) - analytics IDs, page/button tracking, plus skills to scaffold IDs and review coverage
- [`flutter-localization`](plugins/flutter-localization/) - ARB workflows, POEditor, `poe2arb`, and `l10n(context)`

### State and data

- [`flutter-bloc`](plugins/flutter-bloc/) - BLoC/Cubit conventions, state modeling, presentation side effects, `bloc_presentation`, and `flutter_hooks`
- [`flutter-cubit-utils`](plugins/flutter-cubit-utils/) - `QueryCubit`, `PaginatedQueryCubit`, `RequestCubit`, and recipes for lists, details, and actions
- [`flutter-cqrs`](plugins/flutter-cqrs/) - CQRS contracts, repositories, direct `cqrs.run` / `cqrs.get` usage, and CQRS-backed cubits
- [`flutter-forms`](https://github.com/leancodepl/advanced_forms) - `advanced_forms` form and field controllers, validation (sync, async, cross-field), and submit handling, sourced directly from the package repository's `skills/`

### UI and verification

- [`flutter-ui`](plugins/flutter-ui/) - design-system-driven UI, loading/error patterns, localized presentation text, and UI implementation checklists
- [`flutter-patrol`](https://github.com/leancodepl/patrol) - Patrol E2E test skills (a write-test workflow, and test architecture with key conventions), sourced directly from the Patrol repository's `skills/`. **Needs setup:** Patrol CLI and Patrol MCP
- [`flutter-marionette`](plugins/flutter-marionette/) - runtime interaction with a live debug app through Marionette MCP for exploration, smoke checks, and UI debugging. **Needs setup:** Marionette MCP and app-side binding
- [`flutter-read-logs`](plugins/flutter-read-logs/) - read the running app's latest `flutter run` logs as on-demand context via `/read-logs`. **Needs setup:** one editor setting so runs are written to a log file

`flutter-forms` and `flutter-patrol` are not copied into this repository: the marketplace installs them straight from the package that owns them, so they stay in step with the package itself. Where a plugin needs setup, its `README.md` walks through it.

## Contributing

Fixes, sharper wording and new Flutter plugins are welcome, and CI runs every check, so you do not need to install anything to send one. [`CONTRIBUTING.md`](CONTRIBUTING.md) has the steps; [`AGENTS.md`](AGENTS.md) has the plugin conventions.

## License

MIT — see [`LICENSE`](LICENSE).
