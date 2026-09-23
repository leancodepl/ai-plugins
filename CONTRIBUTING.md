# Contributing

Fixes, sharper wording in existing skills and new Flutter plugins are all welcome. You do not need to install anything to contribute:

1. Fork the repo and make your change under `plugins/<plugin-name>/`: edit or add `skills/<skill>/SKILL.md` (and any `skills/<skill>/references/*.md`). Register a new plugin in [`.claude-plugin/marketplace.json`](.claude-plugin/marketplace.json).
2. Check the structure locally: `go run ./cmd/validate-plugins`.
3. Open a PR. CI runs the same structure validation, the validator's tests, Go formatting and lint, and the official Claude Code plugin-spec check (`claude plugin validate . --strict`, where warnings fail the build).

Plugin shape, skill and reference conventions, and the repository's sources of truth are in [`AGENTS.md`](AGENTS.md). Read it before adding a plugin.

If you would rather be walked through it, install [`lean-core`](plugins/lean-core/) and run **`/lean-contribute`**. It asks what kind of change you have in mind, shows the plugin shape, and prints the exact `git`/`gh` steps.

## Repo layout

- `.claude-plugin/marketplace.json` - marketplace index
- `plugins/<plugin-name>/` - one self-contained plugin
- `plugins/<plugin-name>/.claude-plugin/` - plugin manifest
- `plugins/<plugin-name>/skills/` - skills the plugin ships
- `plugins/<plugin-name>/skills/<skill>/references/` - supporting reference material a skill loads on demand
- `cmd/validate-plugins/`, `internal/pluginvalidation/` - the structure validator CI runs
