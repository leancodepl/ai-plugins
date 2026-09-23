# Judging a unit test with Jev

## What this measures

Whether a test is built well enough to be worth its lines: does it assert an outcome rather than a
mechanism, is the right thing replaced by a mock, does it cover what the code distinguishes, and
would a reader know what broke when it goes red.

It does not measure whether the test is *correct*, whether the behaviour it pins is the behaviour
you want, or whether it would catch any particular bug. For that last one, see
[Checking against wrong implementations](#checking-against-wrong-implementations) and read its
warning before you rely on it.

Everything you need in order to trust or distrust a number is in the calibration column of
[The checks](#the-checks).

## Terms

**Mock** — any stand-in a test substitutes for real code: a mocking-library mock, a stub, or a
hand-written fake. This document uses the everyday sense of the word, covering all three. Where the
difference matters — a fake you write by hand behaves differently from a mock you configure — the
text says which.

**Check** — one named judgement in this recipe, sent to Jev as one question. **Panel** — the set of
answers from a single call, read together.

## When not to use this

- **As a merge gate.** No check here is calibrated against a "good enough" bar.
- **On one test pulled out of its file**, for anything except the six test-scope checks below. A
  healthy test leans on its siblings; judged alone it looks weaker than it is.
- **On generated code you have not read.** Jev answers from the state you supply. Supply a
  faithful test file and the real code under test, or the numbers describe nothing.

## Scope: what goes into state

**Default to the whole test file plus the code it exercises.** Three checks are meaningless on a
single test, and one check misreads a single test systematically: a test that leans on its siblings
to exclude a degenerate implementation looks deficient alone. Judge a clean, well-covered test by
itself and the wrong-implementation check raises an alarm caused entirely by the siblings missing
from state.

Give state three fields:

| Field | Contents |
| --- | --- |
| `test_file` | Every test name in the file, the bodies of the ones being judged, and the shared setup, builders, and fakes they call. A faithful condensation is fine; an invented summary is not. |
| `code_under_test` | The methods the tests drive, plus any collaborator the test exercises for real. Without this, every question about what the test pins is unanswerable, and Jev answers anyway. |
| `fixture` and `expected` | The literal inputs and the asserted value. Add these only for the wrong-implementation check. |

**Asking a test-scope check across a file:** put the file in state once and ask one question per
test, pointing each at its own field with a backticked path — `Does the name of \`tests[3]\` describe
what its body asserts?`. Questions run in parallel and cannot see each other, so twenty of them cost
one round trip. Do not ask for a count or an average; ask per test and aggregate the numbers
yourself.

## The checks

Every check is phrased so that **a high number is the healthy answer**. The calibration column
reports measured answers and rates each check by how far apart they fell:

- **Sharp** — healthy and unhealthy inputs landed in bands that do not overlap.
- **Moderate** — they separated, by less than 0.3.
- **Weak** — no input has yet scored low, so the low end is unmeasured.
- **Coarse only** — separates gross cases and misses fine ones; the row says which.

Every threshold in this file is a guess. Every number is measured.

| Check | Type | Scope | Asks | Calibration |
| --- | --- | --- | --- | --- |
| `asserts_observable_result` | noul | test | Do the assertions read a value or state the unit exposes, rather than only which methods were called on a mock? | **Sharp.** 0.11 on an assertion-on-a-captured-argument test; 0.96–0.98 on three tests asserting state |
| `behaviour_not_implementation` | noul | file | Do the assertions refer only to what the unit exposes to callers, rather than internal steps, intermediate states, or call order? | **Sharp.** 0.33 on a file containing one call-verifying test; 0.92 on a file with none |
| `mock_target` | choice | test | What is replaced by a mock: nothing, an external boundary, a collaborator the project owns, or the unit under test itself? | **Sharp.** Correctly picked `external_boundary` (0.82), `owned_collaborator` (0.98), and `nothing` (0.97) |
| `no_logic_in_tests` | noul | test | Is the body free of conditionals, loops, and expected values computed at runtime? | Moderate. 0.63 on a body with a cast and `.single`; 0.94 on a body of literals |
| `better_as_property_test` | noul | file | Do several tests assert one general rule over hand-written inputs, so a property over generated inputs would express it more completely? | Moderate. 0.64 on a 5-test cubit file; 0.86 on a 57-test formatter file that restates one pattern rule |
| `name_matches_body` | noul | test | Does the name describe exactly what the body asserts? | **Sharp, and graded.** 0.07 where the name promises state the body never reads; 0.33 where the name says "empty" and the body passes null; 0.52 for a vague name over a sprawling body; 0.70–0.83 for honest names |
| `single_behavior` | noul | test | Does the body exercise one behaviour? | **Sharp.** 0.19 on a body asserting loading, mapping, the outgoing query, refresh, and close; 0.81–0.89 on single-behaviour bodies |
| `branch_coverage` | score | file | How completely do the tests reach the branches and input shapes the code distinguishes? | **Coarse only.** 0.02 (confidence 0.98) on a one-test file against five branches; 2.80 and 2.94 on two full files — including one where a branch was provably unreachable. See [Enumerate the branches yourself](#enumerate-the-branches-yourself) |
| `readability` | score | test | How much work is it to see which behaviour the test pins? | Weak. 1.87–2.65 with confidence 0.29–0.65 |

**Do** lead with the six checks marked sharp. They carry the signal, and three of them work at
file scope, so they cost one call for a whole file.

**Do** treat `better_as_property_test` as a prompt to widen the inputs, whether or not you adopt a
property-testing library. A high number says the file states one rule through examples; the cheap
response is more data points on the same assertion.

**Do** expect a middling `single_behavior` when a test's setup and its assertion are about different
things — a body that stubs two records and then asserts only that a query went out drives one
behaviour and checks another. Read it alongside `asserts_observable_result`, which falls on that
same shape. A middling `single_behavior` beside a low `asserts_observable_result` means the test is
mis-aimed; a low `single_behavior` beside a healthy one means it is overloaded.

**Do** read `name_matches_body` as a graded answer rather than a yes/no. It falls furthest when the
name claims an outcome the body never reads, settles mid-range when the name misdescribes only the
input, and dips just below the honest band when the name is merely vague. How far it falls tells you
how hard to push.

**Do not** threshold `readability`. No input has yet scored low on it, so its low end is unmeasured.

**Do not** read a high `branch_coverage` as "the branches are covered". It rates whether the test
list looks thorough, so a file with a full complement of named cases scores near the top even when
one branch cannot be reached by any fixture its helpers can build.

**Do not** ask Jev whether the mocking boundary is *right*. Rightness depends on a policy Jev has
no access to. Ask `mock_target` for the classification, which it does reliably, and apply your own
policy to the answer:

| `mock_target` returns | Reading |
| --- | --- |
| `unit_under_test` | Always a defect. The test replaced the very code it claims to verify. |
| `owned_collaborator` | Correct after a repository or service extraction; wrong during a characterization pass, where the point is stubbing the boundary that survives the refactor. |
| `external_boundary` | The default healthy answer. |
| `nothing` | Healthy for pure functions; suspicious for a unit that should have a collaborator. |

## Running it

1. **Assemble state** — the file, the code, nothing else. Accuracy falls as irrelevant material
   grows.
2. **Send one call** with every check you want, including speculative ones. Parallel execution makes
   the second question nearly free.
3. **Read the panel as a panel.** There is no weighted total here; the sample is too small to
   fit weights, and `readability` has never met an input that should score low.

```json
{
  "state": {
    "test_file": "<test names, bodies under judgement, shared setup and fakes>",
    "code_under_test": "<the methods driven, plus collaborators exercised for real>"
  },
  "questions": {
    "asserts_observable_result": {
      "type": "noul",
      "instructions": "Do the assertions in `test_file` check a returned value or the observable state of the unit, rather than only verifying which methods were called on a mock?",
      "criteria": {
        "true": "At least one assertion reads a value the unit produced or exposes",
        "false": "The assertions only verify calls, arguments, or call counts on a mock"
      }
    },
    "behaviour_not_implementation": {
      "type": "noul",
      "instructions": "Do the assertions in `test_file` refer only to values and states the unit exposes to its callers, rather than to internal steps, intermediate states, or the order in which collaborators are invoked?"
    },
    "mock_target": {
      "type": "choice",
      "instructions": "In `test_file`, what is replaced by a mock, stub, or fake rather than exercised for real?",
      "criteria": {
        "nothing": "No mocks, stubs, or fakes; the unit and its collaborators all run for real",
        "external_boundary": "Only a transport or IO boundary the project does not own — an HTTP/GraphQL client, the filesystem, a clock",
        "owned_collaborator": "A class the project owns and could have exercised for real, such as a repository or service",
        "unit_under_test": "The very code whose behaviour the test claims to verify is itself replaced by a mock",
        "other": "Something that fits none of these descriptions"
      }
    },
    "better_as_property_test": {
      "type": "noul",
      "instructions": "Do the tests in `test_file` assert the same relationship repeatedly over hand-written example inputs, such that one property over generated inputs would express that relationship more completely?",
      "criteria": {
        "true": "Several tests differ only in their input values while asserting one general rule",
        "false": "Each test pins a distinct behaviour that a single general rule would not express"
      }
    },
    "no_logic_in_tests": {
      "type": "noul",
      "instructions": "Is the body of `tests[0]` free of conditionals, loops, and expected values computed at runtime, with each expected result written as a literal?"
    },
    "readability": {
      "type": "score",
      "instructions": "How readily can a reader see which behaviour the body of `tests[0]` pins, given the shared setup in `test_file`?",
      "criteria": [
        "The reader must open helpers or shared setup elsewhere to know what inputs the test actually uses",
        "The inputs are visible but setup, action, and expectation are interleaved, so the behaviour must be reconstructed",
        "Inputs are visible and the phases are distinguishable; the behaviour is clear after a careful read",
        "The test reads as one statement of behaviour: inputs, the single action, and the expectation are each obvious at a glance"
      ]
    }
  }
}
```

## Reading a panel

A file whose cubit tests stub the GraphQL client and run the real repository:

```
asserts_observable_result   0.11   → the judged test verifies a captured argument and nothing else
behaviour_not_implementation 0.33  → the file as a whole leans on call-verification
mock_target                external_boundary, confidence 0.77 (owned_collaborator 0.16)
better_as_property_test     0.64
no_logic_in_tests           0.63   → a cast and a `.single` in the body
readability                 2.04, confidence 0.29
```

Act on the first three. The 0.11 names a specific defect in a specific test; the 0.33 says that
defect is the file's habit rather than one slip; the `mock_target` answer confirms the boundary is
the one you intended. The rest are context.

**Do** report the number next to the question that produced it. A confidently wrong number almost
always comes from a badly framed question, and the framing is the part you invented.

**Do** treat a low `confidence` on a score as the answer it is: the levels did not fit, or the file
is mixed. Split it and re-ask.

**Do not** compress the panel into a verdict such as "this file scores 0.7". Jev returns no
reasoning, so a summary number hides the one step a reader needs to check.

## Enumerate the branches yourself

`branch_coverage` answers whether a test list *looks* thorough. It is reliable at the coarse end —
a single-test file against five branches came back 0.02 with 0.98 confidence — and blind to one
missing branch in a file that otherwise looks complete.

**Do read the code, list the branches, and ask one noul per branch.** On the file the aggregate
score rated 2.80:

| Question | Answer |
| --- | --- |
| Does any test supply an input that makes the code take its error branch? | 0.89 |
| Does any test supply an empty or null collection? | 0.97 |
| Does any test supply a list that *contains* a null element, so `.nonNulls` removes something? | **0.18** |

The third is the branch the aggregate score missed.

**Do check the fixture helper's type when a branch looks uncovered.** A branch can be unreachable
because no helper in the file can build an input that gets there, which is invisible in a list of
test names. Put the helper's signature and the field's declared type in state and ask a literal
question about the type:

| Question | Answer |
| --- | --- |
| Can `helper_a`, typed `List<(DateTime, DateTime)>? slots`, produce a list containing a null element? | 0.13 |
| Can `helper_b`, typed `List<Foo?>? slots`, produce one? | 0.95 |
| Using only `helper_a`, can a test reach the case where `.nonNulls` removes an element? | 0.07 |

Widen the helper before adding cases; cases added to a helper that cannot express the input will
pass without executing the branch.

## Checking against wrong implementations

The rule every modern account of test quality agrees on: **a test earns its lines only if some
plausible wrong implementation would make it red.** This section is how to ask that, and why its
answer is worth less than the checks above.

**Do not ask Jev to invent the wrong implementations.** Enumerating candidates and simulating each
against the fixture is multi-hop reasoning, where Jev is weakest. Asked generically — *would a
plausible but wrong alternative implementation fail this test?* — it returned 0.84 on a fixture a
review had condemned as too weak and 0.93 on the fixture that fixed it. Nine points of separation,
pointing the wrong way.

**Do write the candidate implementations yourself**, then ask one question per candidate. Asked
that way, on the same pair of fixtures:

| Question | Weak fixture | Fixed fixture |
| --- | --- | --- |
| Would *take the first element* and *take the earliest by start time* both yield the expected value? | 0.90 | 0.25 |
| Is the first element in the fixture also the earliest by start time? | 0.97 | 0.03 |

**Do list four or five candidates, spanning likely and unlikely.** For a method that selects from a
collection: reads the last element, sorts before reading, returns a constant, ignores an argument,
never fires at all. The unlikely ones cost a few tokens each and occasionally catch something —
"never fires at all" is how a test asserting an empty result turns out to pass against a unit that
does nothing.

```json
{
  "state": {
    "code_under_test": "<the method>",
    "fixture": "<the literal inputs>",
    "expected": "<the asserted value>"
  },
  "questions": {
    "also_passes_if_reads_last": {
      "type": "noul",
      "instructions": "Would an implementation that reads the LAST element of `fixture` produce exactly `expected`?"
    },
    "also_passes_if_sorts_first": {
      "type": "noul",
      "instructions": "Would an implementation that sorts `fixture` by start time before reading the first element produce exactly `expected`?"
    },
    "also_passes_if_never_fires": {
      "type": "noul",
      "instructions": "Would an implementation that performs no work at all — emitting nothing and returning its initial value — produce exactly `expected`?"
    }
  }
}
```

Each answer near 1 names an implementation this fixture cannot rule out. Reshape the fixture until
the ones you care about come back low.

**Know what this is worth.** The checks in the previous section read properties of the test that are
present in the text, so their answers stand on their own. This one answers only the question you
thought to ask, and its phrasing is calibrated on one axis — first versus earliest, one matched
pair of fixtures. A fixture that clears five candidates is protected against those five, and nothing
follows about the sixth you did not think of. **Treat a clean result as the absence of a specific
alarm, never as evidence that the test pins the behaviour.** Where the two conflict, the general
checks are the more trustworthy signal.
