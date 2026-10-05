# lean-jev

> **Experimental.** This plugin sends the content you judge to an external API. Read
> [What leaves the machine](#what-leaves-the-machine) before pointing it at code you do not own.

Offload a judgment call to a model built for judgments, and ask it in a shape that model can
actually answer.

Jev is [TypeSafe](https://docs.typesafe.ai/)'s System One model. It answers typed questions with
numbers: a probability for a yes/no, a position on a rubric you wrote, a pick from a fixed set with
the distribution behind it. No prose, no reasoning, 70 to 500 ms and about $0.0004 a call, which is
cheap enough to judge every item in a list where you would otherwise judge three and generalise.

Asking well is the hard part, and it is the part this plugin is. Jev reads a question *literally*,
answers each one in isolation, and returns a number with no reasoning attached, so a badly framed
question comes back as a confident number that means nothing. The skill carries the rules that keep
a question answerable: one factor per question, the right type for the job, state filtered to what
the question needs.

## Two ways in

**Ad-hoc.** "Score these eleven candidates", "is this finding worth raising", "rate each of these
tests". The general guidelines cover a one-off: the skill picks the question type, splits anything
that bundles two factors, builds the state, batches everything into one call, and reports each
number next to the question that produced it.

**Recipes.** A recurring question deserves a written decomposition rather than a fresh guess every
time. A recipe takes one high-level question a person actually asks, breaks it into Jev-shaped
sub-questions, gives the criteria in full, and states plainly what it leaves unanswered. One ships
with the plugin: [is this test worth its lines?](skills/ask-jev/references/evaluating-a-test.md).
Write your own with [`writing-a-recipe.md`](skills/ask-jev/references/writing-a-recipe.md), which
covers the honesty rules: promise the real question, justify each sub-question as atomic, keep the
criteria unbiased, and tag every threshold as measured, documented, or guessed.

The rubric vocabulary is the durable part. Contextual questions beat generic ones, so what a recipe
preserves is your definition of "severe", "worth raising", "house pattern": templates with a slot
for today's specifics.

## Skills

| Skill | What it does |
| ----- | ------------ |
| [`/lean-jev-usage`](skills/lean-jev-usage/SKILL.md) | Explains the plugin and when to reach for it |
| [`/lean-jev:ask-jev`](skills/ask-jev/SKILL.md) | Turns a judgment call into typed questions, asks Jev, reports the numbers with their framing |

Alongside the skills the plugin ships `scripts/jev.py`, a stdlib-only client, and
`tests/test_jev.py` covering its key handling and request validation.

## What it is good for

Jev makes the judgment a knowledgeable person makes in a few seconds from what is in front of them.
Is this urgent? Which of these three buckets? How severe, on this scale?

Three properties the agent's own read lacks:

- **No drift down a list.** Every question in a call is evaluated independently and in parallel, so
  item 40 gets the same attention as item 1.
- **No stake in the work.** An agent judging code it wrote an hour ago is a compromised reviewer.
- **Consistency.** The same rubric means the same thing next week, which is what makes a threshold
  worth tuning at all.

The boundary: single-hop judgments are in, and anything needing several hops, simulated
consequences, or arithmetic is out. We ranked eleven candidate uses with Jev itself before writing
the skill. "Does this test assert behaviour or implementation" scored 0.86 on single-hop, while
"which architecture approach should we use" scored 0.19 and "which SDK should we adopt" 0.29 at
confidence 0.98, the most confident answer in the whole call. Recognising conformance works;
choosing does not.

## Setup

One file. The script needs no install of its own: Python 3 and a key are the whole story.

```bash
mkdir -p ~/.config/typesafe && chmod 700 ~/.config/typesafe
printf 'export TYPESAFE_API_KEY=%s\n' 'your-key' > ~/.config/typesafe/env && chmod 600 ~/.config/typesafe/env
```

Get the key at [console.typesafe.ai](https://console.typesafe.ai/), dedicated to this integration so
revoking it costs nothing else. The script reads the file itself, so the key stays out of client
configs, out of shell profiles, and off the command line. A `TYPESAFE_API_KEY` already in the
environment wins over the file.

To skip the approval prompt on every call, allow the script in your Claude Code settings:

```json
{ "permissions": { "allow": ["Bash(python3 *//lean-jev/scripts/jev.py:*)"] } }
```

### Verify

```bash
python3 <plugin-root>/scripts/jev.py key-status
python3 <plugin-root>/scripts/jev.py ask \
  --state '{"ticket": "The export button crashes the settings page in Safari.", "reported_by": "support"}' \
  --questions '{"is_bug": {"type": "noul", "instructions": "Does `ticket` report a software defect?"}}'
```

`key-status` reports whether a key was found and from where, never the key. A `noul` near 1 from the
second command means the whole path works.

State goes in a JSON object whenever it has more than one part, so every part carries a name a
question can point at with a backticked path. A bare string is for a single indivisible blob such
as a diff or a log, and the client asks you to declare that with `--state-format text` once such a
string runs long.

## Troubleshooting

| Symptom | Cause | Fix |
| ------- | ----- | --- |
| `python3: can't open file .../scripts/jev.py` | Wrong path, not a broken setup | The script sits at `<plugin-root>/scripts/jev.py`, what `${CLAUDE_PLUGIN_ROOT}` expands to. `ls` the directory |
| `State is N characters of unnamed text` | Several items were flattened into one string | Put each part in a JSON object under its own key and point questions at them with backticked paths, per [`concepts/state.md`](https://docs.typesafe.ai/concepts/state.md). For a genuine single blob, pass `--state-format text` |
| `{"key": "missing"}` from `key-status` | No key file | Write `~/.config/typesafe/env` as above; `ls -l` should show `-rw-------` |
| 401 from TypeSafe | A key was found and rejected | Check it is current at console.typesafe.ai, or a stale `TYPESAFE_API_KEY` in the environment is shadowing the file. `key-status` names the source |
| 422 from TypeSafe | Malformed request, and the API says how | Usually criteria shapes: `score` takes an ordered array, `choice` takes a map, and a `noul`'s optional criteria are a map too. The script checks those locally |
| Answers look like another account's | `TYPESAFE_API_KEY` set in the environment | Unset it to fall back to the file |

The same table lives in [`skills/ask-jev/references/setup.md`](skills/ask-jev/references/setup.md),
which the skill reads when a call fails mid-session.

## Cost

About $0.0004 per case: $42 per million input tokens, output free. A 33-question call over 4.3k
tokens of state costs well under a cent. Keys are per-seat, so usage is attributable.

## What leaves the machine

Whatever goes in `state`: the diff, the test, the ticket, the file. It goes to TypeSafe's API.

Their docs say customer requests are not used for training, and zero-data-retention is
enterprise-only. For code under an NDA that is a decision to make deliberately, and the skill
filtering state down to what each question needs only narrows the exposure. This plugin sends your
material to a third party by design.

## Why a script and not an MCP server

The API is one endpoint: `POST https://api.typesafe.ai/v1/systemone`, a Bearer key, a JSON body of
`{state, model, questions}`, JSON back. An earlier version of this plugin leaned on a community MCP
server that wrapped exactly that in 1290 lines of Go, a self-updater and a release-provenance chain.

Dropping it removed three risks: an unsigned binary with no build attestation, a self-updater that
pulls unreviewed code, and a fallback route that sent state to a second vendor when the TypeSafe key
was absent. It also means there is nothing to register, so the client is the same file wherever the
agent runs.

The key lives in a `600` file that `jev.py` reads itself, rather than in an agent's environment,
because an exported key is inherited by every command that agent spawns, including build scripts
from whatever repository happens to be open. And the script uses `urllib` rather than shelling out
to `curl`: a header handed to curl is an argv element that any `ps` on the machine can read while
the request is in flight, while `urllib` keeps it in the process. There is a comment in `jev.py`
saying so, because the curl version looks like a harmless simplification.

## When to keep the judgment

- Choosing between approaches, tools, or vendors: multi-hop, and it scored worst of everything we tried
- Counting, dates, thresholds, anything you can compute
- A single judgment you have no stake in. A tool call costs a turn, and a Jev number wrapped around
  the agent's own opinion is worse than that opinion stated plainly
- Code where the data-egress question is still open

## Install

```
/plugin install lean-jev@leancode-ai-plugins
```

Then write the key file as above. See the [root README](../../README.md#install) for adding the
marketplace itself.
