# Writing a recipe reference

A **recipe** takes one recurring high-level question — the kind someone actually asks out loud,
like "is this change any good?" — and shows how to break it into questions Jev is reliable at.
The recipe is where the thinking lives. The primitive references say what a Score *is*; a recipe
says which four Scores answer this particular question and why those four.

Add one when you have worked out a decomposition that held up on real inputs. Then add a row to
the Recipes table in `SKILL.md`.

## Why the honesty rules exist

A recipe is read later by someone — possibly you, in a fresh session — who was not there when it
was worked out, and who has no way to tell a threshold you measured from a threshold you invented.
If the recipe does not mark the difference, they cannot either, and a guessed 0.7 gets treated as
a finding. Every rule here exists to keep that from happening.

## Promise the real question

Open with the question in the words someone would actually use, not a pre-sanitised version of
it. If the honest question is "is this PR ready to merge", say that — then show the
decomposition, and be explicit that the decomposition does not fully answer it.

Do not quietly narrow the promise to fit what Jev can do. Narrowing is the recipe's whole job,
but it has to be visible: *"this measures four things that correlate with readiness; it does not
tell you whether the change is correct."*

## Decompose, and justify each part

For each sub-question, say what it is and why it is atomic — why it does not smuggle a second
judgment inside. If you cannot articulate that, it probably is not atomic.

Then say how the parts combine, and whose call that is. Weights and thresholds are policy, not
measurement; they belong in the open where someone can disagree with them.

## Keep the framing unbiased

The criteria are where bias enters, usually without anyone meaning it. Watch for:

- **Options that telegraph the wanted answer.** `{"clean": "well-structured, idiomatic", "messy": "bad"}`
  tells the model which answer you want. Describe each option in the terms its own advocate
  would use.
- **Ladders with a thumb on the scale.** A 5-level rubric where four levels are flavours of bad
  will report bad. Space the levels over the range you actually expect to see.
- **State that argues for a conclusion.** Passing your own summary ("this function is overly
  complex") as state and then asking whether it is complex measures nothing. Pass the source,
  not your read of it.
- **Missing escape hatches.** No `other`, no `none of the above`, means a forced answer. A
  Choice always picks the best of what you offered, including when nothing fits.

A useful test: write the criteria, then ask whether someone arguing the opposite side would
accept them as fair. If not, rewrite before running anything.

## Be honest about what was actually tried

Label every number and claim with where it came from. Three tags are enough:

- **Measured** — you ran it, on this many real inputs, and here is what came back.
- **From the docs** — TypeSafe documents this; link the page.
- **Guess** — it seemed sensible and has not been tested.

A guessed threshold is perfectly fine to ship. A guessed threshold presented as measured is not.
If the recipe was validated on six hand-picked examples, say six and say hand-picked; that is
useful information, and inflating it destroys the reader's ability to calibrate.

Include a worked example with real numbers when you have them. If you have none, mark the recipe
**untested** at the top and say so plainly — an untested decomposition is still worth writing
down, as long as nobody mistakes it for a validated one.

## Say what you do not know

Close with the open questions. The failure modes you suspect but have not confirmed. The inputs
you never tried it on. The places where you are unsure whether the decomposition holds. This
section is the most valuable part of the file for the next reader, and the easiest to skip
because it is the least satisfying to write.

Also say when **not** to use the recipe. A recipe that claims to work everywhere has not been
used enough to know better.

## Template

```markdown
# <High-level question, in plain words>

**Status:** untested | validated on N real inputs (<what they were>)

## What this answers, and what it does not
<The promise, and the honest limits of it.>

## When not to use this
<Inputs or situations where the decomposition breaks down.>

## The decomposition
| Sub-question | Type | Why it is atomic |
| --- | --- | --- |

## The call
<The full evaluate call, copy-pasteable.>

## Reading the result
<What each number means. Thresholds, each tagged measured / from the docs / guess.>

## Worked example
<Real state, real numbers, real interpretation — or "none yet".>

## Open questions
<What is unverified, untried, or suspected.>
```

## Keep the rubric, rewrite the question

Keep the rubric reusable and the state contextual. What is worth making durable is the
vocabulary — the severity ladder, the definition of "worth raising" — not the exact question
text. The instruction and the state should be rewritten per use; the criteria are what earns the
right to be copied.
