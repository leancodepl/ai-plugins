# LeanCode AI Plugins

Claude Code plugins that make Claude write Flutter code the way an experienced Flutter team does.

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Works with Claude Code](https://img.shields.io/badge/works%20with-Claude%20Code-d97757.svg)](https://code.claude.com/docs/en/overview)

## What this is

A plugin in this marketplace is a set of Claude Code **skills**: focused instructions plus reference material that Claude loads only when they are relevant. Each skill is a `SKILL.md` whose short description stays in context, while its body — and any `references/*.md` alongside it — loads only once the skill actually fires, so deep convention documents cost nothing until they are needed. The result is that Claude picks up a team's architecture, state management, and testing conventions at the moment it is writing that kind of code, instead of being told everything up front.

## Who it's for

- **Flutter teams who want opinionated defaults.** These are LeanCode's conventions — our architecture, BLoC, CQRS, forms, and E2E choices, not a neutral survey of the Flutter ecosystem. You may well disagree with some of them. Each plugin installs independently, so take only the ones you agree with and leave the rest; nothing here assumes you installed the whole set.
- **Anyone who wants a worked example of a real plugin marketplace.** This repo is a complete, working Claude Code marketplace in public: a `marketplace.json` index, self-contained plugins, externally sourced entries that point at another repository, and CI that validates the whole thing. Read it as a reference for building your own.

Pick only the plugins that match your project — each is independently installable.

## A worked example

Install [`flutter-read-logs`](plugins/flutter-read-logs/), run your app from your editor, and then ask:

```
/read-logs why did the login screen flash before the dashboard
```

The skill resolves this project's log file (`/tmp/flutter-<repo>.log`, derived from the shared `.git` so every worktree maps to one file), tells you the exact path it is about to read, and checks how fresh the run is — if Dart files under `lib/` changed after the run was captured, it says so before drawing conclusions. It then normalizes the capture (a raw terminal transcript, or the JSON-framed Debug Adapter log VS Code and Cursor write) and reads it *with your question as the lens*: for an ordering question like this one it traces the relevant cubits and events in sequence rather than hunting for errors.

What comes back is an answer to your question, backed by lines from that run — not a summary of the log and not a list of every error in it. If no log exists yet, it walks you through the one-time editor setup instead of just reporting a missing file — with your go-ahead it writes your editor's local run/debug settings, and it never commits anything. Reading a run does send that run's contents to the model, which is why the skill names the file before it reads it; the plugin's `README.md` carries the full data-handling note.

## Install

These plugins are distributed as a Claude Code **plugin marketplace**. Add the marketplace once, then install the plugins you want — you do not need to clone this repo for day-to-day usage.

```
/plugin marketplace add leancodepl/ai-plugins
```

Then install any plugin from the list below by name, suffixed with this marketplace's name, `leancode-ai-plugins`:

```
/plugin install flutter-bloc@leancode-ai-plugins
```

After installing, run `/reload-plugins` to activate. Browse the full catalog with `/plugin` (the **Discover** tab).

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

**Claude Code** only, installed via the LeanCode plugin marketplace. Claude Code has no separate "rules" concept for plugins — `plugin.json` recognizes no `rules` key, and a `CLAUDE.md` at a plugin root is not loaded — so everything here ships as skills and their reference files.

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
- [`flutter-forms`](https://github.com/leancodepl/advanced_forms) - `advanced_forms` form and field controllers, validation (sync, async, cross-field), and submit handling, sourced directly from the package repository's `skills/` — the single source of truth for `advanced_forms` AI support

### UI and verification

- [`flutter-ui`](plugins/flutter-ui/) - design-system-driven UI, loading/error patterns, localized presentation text, and UI implementation checklists
- [`flutter-patrol`](https://github.com/leancodepl/patrol) - Patrol E2E test skills (write-test workflow, test architecture, key conventions, Patrol MCP), sourced directly from the Patrol repository's `skills/` — the single source of truth for Patrol AI support
- [`flutter-marionette`](plugins/flutter-marionette/) - runtime interaction with a live debug app through Marionette MCP for exploration, smoke checks, and UI debugging
- [`flutter-read-logs`](plugins/flutter-read-logs/) - read the running app's latest `flutter run` logs as on-demand context via `/read-logs`

## Per-plugin setup

Most plugins are pure skills and reference material with no setup. A few need one-time tooling or MCP setup — finish it from the plugin's `README.md`:

- [`flutter-patrol`](https://github.com/leancodepl/patrol) - Patrol CLI and Patrol MCP
- [`flutter-marionette`](plugins/flutter-marionette/) - Marionette MCP and app-side binding

## Most plugins have a `-usage` skill

Every plugin sourced from this repository exposes a `/<plugin-name>-usage` skill once installed — for example `/flutter-bloc-usage`, `/flutter-cqrs-usage`, `/flutter-ui-usage`. Run it to see what the plugin covers, its conventions, and example prompts to try next. It's the fastest way to learn a plugin without reading its full `README.md`. The externally sourced `flutter-forms` and `flutter-patrol` entries ship their own skills instead, named by their home repositories. If you're not sure where to begin, run `/lean-core-usage` for a tour of the whole marketplace.

> **Tip:** Install [`lean-core`](plugins/lean-core/) first and run `/lean-core-usage` for a guided tour of the whole marketplace.

## Contributing

Outside contributions are welcome — new skills, fixes and sharper wording in existing skills, and new Flutter plugins. You do not need to install anything to contribute:

1. Fork this repository.
2. Edit or add `plugins/<plugin-name>/skills/<skill>/SKILL.md` (plus any `skills/<skill>/references/*.md`), and register a new plugin in [`.claude-plugin/marketplace.json`](.claude-plugin/marketplace.json).
3. Run `go run ./cmd/validate-plugins` to check the structure locally.
4. Open a PR. CI re-runs that structure validation, Go formatting and lint, and the official Claude Code plugin-spec check (`claude plugin validate . --strict`, where warnings fail the build).

[`CONTRIBUTING.md`](CONTRIBUTING.md) has the same steps plus the guided path, and [`AGENTS.md`](AGENTS.md) holds the conventions.

## Repo layout

- `plugins/<plugin-name>/` - one self-contained plugin
- `plugins/<plugin-name>/skills/` - skills the plugin ships
- `plugins/<plugin-name>/skills/<skill>/references/` - supporting reference material a skill loads on demand
- `plugins/<plugin-name>/.claude-plugin/` - plugin manifest
- `.claude-plugin/marketplace.json` - marketplace index

## License

MIT — see [`LICENSE`](LICENSE).
