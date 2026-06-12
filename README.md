# LeanCode AI Plugins

LeanCode's shared AI plugins for Claude Code, focused on Flutter development.

## Supported clients

- **Claude Code** — installed via the LeanCode plugin marketplace in Claude Code.

## Install

These plugins are distributed as a Claude Code **plugin marketplace**. Installing is a two-step process — add the marketplace once, then install the plugins you want. You do not need to clone this repo for day-to-day usage. For the full reference, see Claude Code's [Discover and install plugins](https://code.claude.com/docs/en/discover-plugins) docs.

### 1. Add the marketplace

In Claude Code, add this repository as a marketplace by its GitHub `owner/repo`:

```
/plugin marketplace add leancodepl/ai-plugins
```

This registers the catalog so you can browse it — nothing is installed yet. (The full git URL works too: `/plugin marketplace add https://github.com/leancodepl/ai-plugins.git`.)

### 2. Install the plugins you need

Browse the catalog with `/plugin` (the **Discover** tab), or install directly by name. Plugins install from the `leancode-ai-plugins` marketplace:

```
/plugin install flutter-bloc@leancode-ai-plugins
```

Pick only the plugins that match your project — each is independently installable. After installing, run `/reload-plugins` to activate them.

> **Tip:** Install [`lean-core`](plugins/lean-core/) first and run `/lean-core-usage` for a guided tour of the whole marketplace.

### Team setup (optional)

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

### Per-plugin setup

Most plugins are pure rules and skills with no setup. A few need one-time tooling or MCP setup — finish it from the plugin's `README.md`:

- [`flutter-patrol`](plugins/flutter-patrol/) - Patrol CLI and Patrol MCP
- [`flutter-marionette`](plugins/flutter-marionette/) - Marionette MCP and app-side binding

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
- [`flutter-forms`](plugins/flutter-forms/) - `leancode_forms`, validation behavior, naming conventions, and form cubit patterns

### UI and verification

- [`flutter-ui`](plugins/flutter-ui/) - design-system-driven UI, loading/error patterns, localized presentation text, and UI implementation checklists
- [`flutter-patrol`](plugins/flutter-patrol/) - Patrol test architecture, key conventions, and Patrol MCP workflow for AI-assisted E2E work
- [`flutter-marionette`](plugins/flutter-marionette/) - runtime interaction with a live debug app through Marionette MCP for exploration, smoke checks, and UI debugging
- [`flutter-read-logs`](plugins/flutter-read-logs/) - read the running app's latest `flutter run` logs as on-demand context via `/read-logs`

### Docs and reporting

- [`lean-interactive-html`](plugins/lean-interactive-html/) - single-file interactive HTML artifacts (reports, dashboards, plans, architecture diagrams, decision forms) via `/interactive-html` — not Flutter-specific

### Every plugin has a `-usage` skill

Once a plugin is installed, it exposes a `/<plugin-name>-usage` skill — for example `/flutter-bloc-usage`, `/flutter-cqrs-usage`, `/flutter-ui-usage`. Run it to see what the plugin covers, its conventions, and example prompts to try next. It's the fastest way to learn a plugin without reading its full `README.md`. If you're not sure where to begin, run `/lean-core-usage` for a tour of the whole marketplace.

## Repo layout

- `plugins/<plugin-name>/` - one self-contained plugin
- `plugins/<plugin-name>/skills/` - skills the plugin ships
- `plugins/<plugin-name>/skills/<skill>/references/` - supporting reference material a skill loads on demand
- `plugins/<plugin-name>/.claude-plugin/` - plugin manifest
- `.claude-plugin/marketplace.json` - marketplace index

## Contributing

If you are interested in contributing, please see [`CONTRIBUTING.md`](CONTRIBUTING.md) 
