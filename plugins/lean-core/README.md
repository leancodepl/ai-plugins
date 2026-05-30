# lean-core

Entry point for the LeanCode AI plugins marketplace. Two manual skills:

- `/lean-core-usage` — explains what the marketplace covers and which plugins are available.
- `/lean-contribute` — walks a contributor through proposing a change and opening a PR.

Both skills are manual-only (`disable-model-invocation: true`). They don't auto-fire on inferred intent.

## Onboarding

Install the marketplace in your client (the supported list is in the [root README](../../README.md)), then run:

```
/lean-core-usage
```

It explains what's in the marketplace and points you at the right `/<plugin>-usage` for whatever you're doing.

## Contributing

```
/lean-contribute
```

Asks what kind of change you have in mind (tweak, new skill or rule, new plugin), shows the plugin shape, and prints the exact `git`/`gh` steps. CI validates every PR; `gh pr checks <pr-number>` reports the verdicts.
