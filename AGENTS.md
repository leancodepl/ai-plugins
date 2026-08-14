# AGENTS.md

## Purpose

LeanCode's shared AI plugins for Claude Code, focused on Flutter development.

## Repository layout

- `.claude-plugin/` holds the Claude Code marketplace metadata.
- `plugins/<plugin-name>/` contains one self-contained plugin.
- Each plugin contains `skills/`, optionally `agents/` (most `flutter-*` plugins ship one), and optionally `.mcp.json` if it drives an MCP server.
- Each plugin should have its own `README.md` describing scope and assets.

Claude Code also documents support for per-plugin `commands/` and `hooks/` directories. No plugin in this repo currently uses them — treat them as not-yet-exercised.

## Externally sourced plugins

A plugin whose content is owned by another repository is registered in `.claude-plugin/marketplace.json` with an object source and `strict: false`:

```json
{
  "name": "flutter-patrol",
  "source": { "source": "github", "repo": "leancodepl/patrol" },
  "strict": false,
  "skills": ["./skills/patrol-write-test", "./skills/patrol-test-architecture"]
}
```

The external repo is the single source of truth — there is no `plugins/<name>/` directory and no copy of its content here; do not vendor one. The validator skips external entries (no local directory to check); the `claude plugin validate` CI step covers the entry's schema.

Currently external: `flutter-patrol` (from [leancodepl/patrol](https://github.com/leancodepl/patrol), which owns all Patrol AI support).

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

## Agents

Most `flutter-*` plugins ship a per-plugin agent at `agents/<plugin-name>.md`, auto-discovered by Claude Code (no `agents` key in `plugin.json`). The agent is a **thin wrapper over the plugin's existing skills** — it never duplicates convention content. Conventions stay in `skills/<plugin>-usage/references/*.md` as the single source of truth.

Conventions for these agents:

- **Name** = plugin name; invoked `@<plugin-name>` (e.g. `@flutter-navigation`). No collision with skills, which use `/<skill-name>`.
- **`skills:` frontmatter preloads only the `<plugin>-usage` router** (the small routing skill), not its references. The system-prompt body instructs the agent to inspect the project and lazy-load only the matching reference — preserving progressive disclosure.
- **`description`** scopes the agent to substantial, multi-step work; the `/<plugin>-usage` skill remains the path for quick inline questions.
- **Tools:** editing plugins use `Read, Glob, Grep, Edit, Write, Bash, Skill`.
- **`model`** is omitted everywhere (inherits the session model).
- No agent for `lean-core` (marketplace meta) or the MCP-integration plugins `flutter-patrol` and `flutter-marionette` — those are backed by their own separate tools/repos, so a wrapper agent there is premature.

These agents are Claude Code only and do not mirror to other targets; keeping them thin over `references/*.md` is what contains that cost.

## Local checks

Validate plugin structure locally with Go:

```
go run ./cmd/validate-plugins
```

CI runs the same structure validation, Go formatting/lint, and the official Claude Code plugin-spec check (`claude plugin validate . --strict`, warnings fail the build) on every PR. After pushing, watch CI with `gh pr checks <pr-number>`.

## Platform notes

- Skills live inside each plugin directory (`./skills/`), each with `SKILL.md` and optional `references/`.
- Plugin manifests are thin: `.claude-plugin/plugin.json` points `skills` at `./skills/`. It carries no `rules` key — Claude Code does not recognize one.
