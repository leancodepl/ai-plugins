# Changelog

## 0.3.0

- Reading the TypeSafe docs is a precondition rather than a step inside the loop. `ask-jev/SKILL.md` opens with the pages to read before the first call: `concepts/state.md` every time, plus the page for each question type in use. It says plainly that the request contract is not a substitute for them, because a request can satisfy the contract, ignore the pages, and come back as a confident number about the wrong thing with nothing in the response to mark it.
- A page that will not load is now a reason to stop and tell the user. The clause it replaces allowed a call on the contract alone provided the agent said so in its report, which asked the agent that skipped the reads to be the one to disclose it.
- `jev.py` refuses a state over 400 characters passed as unnamed text, names `concepts/state.md`, and offers `--state-format text` for a genuine single blob such as a diff or a log. Flattening several items into one string is the framing mistake that costs the most, and it is the one an agent makes while believing the request is well formed.
- The examples in `SKILL.md` and `README.md` pass an object and point a question at one of its fields, so the shape a reader copies is the shape most states want.

## 0.2.0

First public release, carried over from LeanCode's internal marketplace at the same version.

- `skills/ask-jev/` decides when a judgment is worth delegating, decomposes it into atomic typed questions, and reports every number next to the question that produced it. References cover the API key and its failure modes, the recipe framework, and one worked recipe for judging a unit test.
- `skills/lean-jev-usage/` explains the plugin, ad-hoc questions versus recipes, cost, and when to keep the judgment.
- `scripts/jev.py` posts to `https://api.typesafe.ai/v1/systemone` from the standard library alone, reads the key from `~/.config/typesafe/env`, validates criteria shapes before spending a call, and retries throttling and transient failures.
- `tests/test_jev.py` covers key resolution, request validation, state handling, and that the key travels in the header rather than the payload.
