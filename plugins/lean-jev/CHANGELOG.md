# Changelog

## 0.2.0

First public release, carried over from LeanCode's internal marketplace at the same version.

- `skills/ask-jev/` decides when a judgment is worth delegating, decomposes it into atomic typed questions, and reports every number next to the question that produced it. References cover the API key and its failure modes, the recipe framework, and one worked recipe for judging a unit test.
- `skills/lean-jev-usage/` explains the plugin, ad-hoc questions versus recipes, cost, and when to keep the judgment.
- `scripts/jev.py` posts to `https://api.typesafe.ai/v1/systemone` from the standard library alone, reads the key from `~/.config/typesafe/env`, validates criteria shapes before spending a call, and retries throttling and transient failures.
- `tests/test_jev.py` covers key resolution, request validation, state handling, and that the key travels in the header rather than the payload.
