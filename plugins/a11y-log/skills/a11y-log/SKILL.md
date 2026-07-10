---
name: a11y-log
description: Capture a text transcript of everything VoiceOver speaks (and each keyboard nav gesture) during a manual screen-reader session on macOS, then diagnose it into an accessibility bug report — a text alternative to screen recording. Claude spawns a dedicated Terminal running the logger, the dev enables VoiceOver and does the flow, then Claude reads the log and turns the announcement timeline into a bug report. Project-agnostic (any macOS app: web, native, Flutter). Trigger on "/a11y-log", "log VoiceOver", "capture VO speech", "record what the screen reader says".
---

# a11y-log — VoiceOver speech logging (developer-in-the-loop)

Turns a manual VoiceOver session into a timestamped transcript, then diagnoses it.
The dev drives VoiceOver by hand; every spoken phrase (and every nav gesture, if
Karabiner logging is set up) lands in one file; Claude reads it and writes up the
bugs. Project-agnostic — it captures whatever app is frontmost.

**One-time machine setup lives in this plugin's `README.md`, not here** (VoiceOver
AppleScript control, Automation permission, and the Karabiner gesture rules). The
logger script prechecks that setup and prints what is missing, so you rarely need
to explain it — just point the dev at the README's *Setup* section if a check fails.

## Workflow

### 1. Start the logger (Claude does this)

```bash
bash "${CLAUDE_PLUGIN_ROOT}/skills/a11y-log/scripts/start_log_terminal.sh" ~/vo_log.txt
```

A new Terminal window opens showing the live transcript. Do **not** activate/steal
focus to it — that moves the VoiceOver cursor off the app under test. If the
script's precheck fails, tell the dev to press **Cmd+F5** first (Claude must not
toggle VoiceOver — the dev controls it so focus lands where they expect) and see
the README *Setup* section.

### 2. Dev runs the session (human)

Tell the developer:
> Logger's live in the new window. Switch to the app, do the flow that's buggy.
> Gestures auto-log if you set up the Karabiner rules; otherwise type a short note
> + Enter in that window when you do something notable (e.g. "opened market
> dropdown"). When you're done, come back and say "read the log".

### 3. Read & diagnose (Claude does this)

Extract just this session (from the last `=== requested … ===` marker) and build
the timeline:

```bash
awk '/=== requested/{buf=""} {buf=buf $0 "\n"} END{printf "%s", buf}' ~/vo_log.txt
```

Then analyze the announcement sequence with the table below. For each bad
announcement, state the chain *phrase → focused widget → file:line → property →
fix*. Fixing and verifying continues via the `a11y-audit` skill (the AI re-drives
VoiceOver to prove the fix) or the dev re-runs this logger after the change.

## Diagnosing the transcript

| Phrase / symptom | Means | Likely cause & fix |
|------------------|-------|--------------------|
| "…, group" + "press Control-Option-Shift-Down" | interactable group, not a transparent region — user must drill in | `Semantics` with `container: true` **and** a `label` wrapping interactive children. Drop them; a landmark `role` alone is transparent |
| "end of, <name>" in the wrong place | hard group edge around a landmark | remove `container: true` |
| GESTURE with no VO line (silent stop) | unlabeled node / covered pane | add `Semantics(label:)`, or `ExcludeSemantics` if decorative; if a whole pane, check an `ExcludeSemantics(excluding:)` gating it |
| container announced with a child's name | name-steal — container unnamed | name the route/region (`namesRoute`) or remove the merging label |
| element read twice per swipe | duplicate focus stop | collapse nested focusables to one |
| title read as plain text | missing heading role | `Semantics(header: true)` / heading-level widget |
| visible label ≠ spoken name | label overrides visible text (2.5.3) | align `Semantics(label:)` with the visible text |

Useful greps (Flutter): `grep -rn "container: true" lib --include="*.dart"` (the
usual landmark culprit) and `grep -rn "Semantics(" lib/<feature>`.

## Reading signals

- **GESTURE with no following VO line** → silent focus move. Focus went somewhere
  unlabeled/unannounced. Prime suspect for "I don't know where I am".
- **Same phrase repeated** → focus bounced back, or a live-region is chattering.
- **A long stream after one action** → over-verbose node (e.g. a dialog reading its
  whole subtree). Note it, but it's often acceptable.

## Log format

```
HH:MM:SS | VO      | <phrase VoiceOver spoke>
HH:MM:SS | GESTURE | vo-right          (auto, from Karabiner)
HH:MM:SS | ACTION  | <freeform note the dev typed>
```

## Video fallback

For TalkBack (no macOS scripting) or inherently visual bugs (focus-ring position),
record with the caption panel on (VO+F10 / TalkBack "Display speech output"), then
slice and read it frame by frame:
`ffmpeg -v error -i <video> -vf "fps=2" ai_tmp/frames/f_%03d.png` — treat each
caption exactly like a `VO |` log line. Delete `ai_tmp/` afterwards.

## Companion skill

- `a11y-audit` — the AI *operates* VoiceOver itself (autonomous audit, bug repro,
  fix verification), no human swiping. Use whenever Claude should drive the screen
  reader instead of the dev.
