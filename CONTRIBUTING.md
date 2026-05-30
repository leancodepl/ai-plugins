# Contributing

Contributing to the LeanCode AI plugins is intentionally light.

The guided path is the [`lean-core`](plugins/lean-core/) plugin: run **`/lean-contribute`** and it asks what kind of change you have in mind (tweak, new skill or rule, new plugin), shows the plugin shape, and prints the exact `git`/`gh` steps. It is the canonical, always-current contributor reference — prefer it over this file.

If you'd rather work unguided:

1. Clone the repo and make your change under `plugins/<plugin-name>/` — edit or add `skills/<skill>/SKILL.md` (and any `skills/<skill>/references/*.md`), and register the plugin in [`.claude-plugin/marketplace.json`](.claude-plugin/marketplace.json).
2. Validate structure locally: `go run ./cmd/validate-plugins`.
3. Open a PR. CI validates structure on every PR.

Plugin shape, skill and rule conventions, and other requirements live in [`AGENTS.md`](AGENTS.md).

This repo currently ships LeanCode's Flutter plugins; Flutter contributions are welcome.
