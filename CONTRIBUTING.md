# Contributing

Contributing to the LeanCode AI plugins is intentionally light. Outside contributions are welcome — new skills, fixes and sharper wording in existing ones, and new Flutter plugins.

You do not need to install anything to contribute:

1. Fork the repo and make your change under `plugins/<plugin-name>/` — edit or add `skills/<skill>/SKILL.md` (and any `skills/<skill>/references/*.md`), and register a new plugin in [`.claude-plugin/marketplace.json`](.claude-plugin/marketplace.json).
2. Validate structure locally: `go run ./cmd/validate-plugins`.
3. Open a PR. CI runs the same structure validation, Go formatting and lint, and the official Claude Code plugin-spec check (`claude plugin validate . --strict`, where warnings fail the build).

Plugin shape, skill and reference conventions, and the repository's sources of truth live in [`AGENTS.md`](AGENTS.md) — read it before adding a plugin.

If you would rather be walked through it, the [`lean-core`](plugins/lean-core/) plugin offers a guided path: run **`/lean-contribute`** and it asks what kind of change you have in mind (tweak, new skill, new plugin), shows the plugin shape, and prints the exact `git`/`gh` steps. It is an option, not a requirement — the unguided steps above are fully supported.

This repo currently ships LeanCode's Flutter plugins; Flutter contributions are welcome.
