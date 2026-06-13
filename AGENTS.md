# AGENTS.md

## Purpose

LeanCode's shared AI plugins for Claude Code, focused on Flutter development.

## Repository layout

- `.claude-plugin/` holds the Claude Code marketplace metadata.
- `plugins/<plugin-name>/` contains one self-contained plugin.
- Each plugin contains `skills/`, and optionally `.mcp.json` if it drives an MCP server.
- Each plugin should have its own `README.md` describing scope and assets.

Claude Code also documents support for per-plugin `agents/`, `commands/`, and `hooks/` directories. No plugin in this repo currently uses them — treat them as not-yet-exercised.

## Working conventions

- Keep plugins self-contained and independently installable.
- Prefer minimal, documented manifest fields over convenience metadata unless Claude Code clearly requires more.
- When adding a skill, include the required frontmatter (`name`, `description`) so marketplace parsing keeps working.
- Update the plugin `README.md` when adding a new capability.
- Avoid coupling shared content to local-only files or directories.

## Sources of truth

- **Plugin list** — `.claude-plugin/marketplace.json`. Read it to enumerate plugins; do not maintain a separate list here or in skill bodies.
- **Supported clients** — the "Supported clients" section of the root `README.md`. Read it to enumerate clients.

When a plugin is added, removed, or renamed, update `.claude-plugin/marketplace.json` and the grouped overview in the root README.

## Skills and reference material

Everything a plugin ships is a skill. A skill is a directory `skills/<skill-name>/` with a `SKILL.md` entrypoint and optional supporting files alongside it.

- **`SKILL.md`** — required. YAML frontmatter (`name`, `description`) plus a concise body. The `description` is always in context and is what makes Claude load the skill; the body loads only when the skill fires. Keep the body under ~500 lines.
- **`references/*.md`** — optional supporting files holding detailed conventions, patterns, and domain knowledge. They load only when `SKILL.md` points at them, so long reference material costs nothing until needed. This is where the LeanCode coding conventions live (formerly the per-plugin rule files).
- Every plugin ships a `<plugin-name>-usage/SKILL.md` routing skill. Convention-style guidance for the plugin lives in `skills/<plugin-name>-usage/references/`, and `SKILL.md` lists those files under a "Reach for these references" section so Claude knows when to load each.

Claude Code has no separate "rules" concept for plugins — `plugin.json` recognizes no `rules` key, and a `CLAUDE.md` at the plugin root is not loaded. Ship guidance as skills and their reference files.

## Local checks

Validate plugin structure locally with Go:

```
go run ./cmd/validate-plugins
```

CI runs the same structure validation, Go formatting/lint, and the official Claude Code plugin-spec check (`claude plugin validate . --strict`, warnings fail the build) on every PR. After pushing, watch CI with `gh pr checks <pr-number>`.

## Platform notes

- Skills live inside each plugin directory (`./skills/`), each with `SKILL.md` and optional `references/`.
- Plugin manifests are thin: `.claude-plugin/plugin.json` points `skills` at `./skills/`. It carries no `rules` key — Claude Code does not recognize one.
