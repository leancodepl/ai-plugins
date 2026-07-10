# a11y-vo-log

Capture a timestamped **text transcript of everything VoiceOver speaks** during a
manual screen-reader session on macOS — and, optionally, each keyboard nav gesture
— then have Claude diagnose it into an accessibility bug report. A text alternative
to recording and re-watching video.

The developer drives VoiceOver by hand; Claude opens the logging Terminal, reads the
resulting file, and turns the announcement timeline into `phrase → widget →
file:line → fix` chains. Project-agnostic: it logs whatever macOS app is frontmost
(web, native, Flutter).

## Included assets

- `skills/a11y-vo-log/SKILL.md` — the workflow (spawn logger → dev runs flow → Claude
  diagnoses) plus the announcement-diagnosis table
- `skills/a11y-vo-log-usage/SKILL.md` — routing skill: what the plugin does and when to
  reach for it vs. `a11y-audit`
- `skills/a11y-vo-log/scripts/vo_log.sh` — the poller: reads VoiceOver's "last phrase"
  ~6×/s (deduped) and appends `VO` lines
- `skills/a11y-vo-log/scripts/start_log_terminal.sh` — opens the dedicated logging
  Terminal window running `vo_log.sh`
- `skills/a11y-vo-log/scripts/vo_gesture.sh` — appends a `GESTURE` line; called by
  Karabiner on each VoiceOver nav combo
- `skills/a11y-vo-log/assets/karabiner_vo_gesture_logger.json` — Karabiner rules
  template (the `__VO_GESTURE_SH__` placeholder is stamped during setup)

## How it works

`vo_log.sh` polls VoiceOver over AppleScript and writes one line per spoken phrase.
Keyboard gestures cannot be captured by a self-contained tap (VoiceOver disables a
`CGEventTap` after the first event), so gesture lines come from **Karabiner-Elements**:
each VoiceOver nav combo (VO+Right/Left/Up/Down, VO+Space, VO+Shift+Down…) runs
`vo_gesture.sh`, which appends a `GESTURE` line to the same log — including for
**silent focus moves** (a gesture with no `VO` line after it = focus moved but
nothing was spoken, the #1 screen-reader smell).

```
HH:MM:SS | VO      | <phrase VoiceOver spoke>
HH:MM:SS | GESTURE | vo-right          (auto, from Karabiner)
HH:MM:SS | ACTION  | <freeform note the dev typed in the log window>
```

## Setup (one-time, per machine)

### 1. Let VoiceOver be driven by AppleScript

VoiceOver Utility (**VO+F8**) → General → ☑ *"Allow VoiceOver to be controlled with
AppleScript"*.

### 2. Grant Automation permission

The process running Claude Code needs to control **VoiceOver** and **Terminal**. The
first run may prompt *"… wants to control Terminal / VoiceOver"* → **Allow**. If it
errors instead, add them under System Settings → Privacy & Security → Automation.

That is enough for **speech logging**. `/a11y-vo-log` now works — you just won't get
`GESTURE` lines until step 3.

### 3. Keyboard-gesture logging (Karabiner-Elements)

This is what adds the `GESTURE` lines (and lets Claude spot silent focus moves).
Install [Karabiner-Elements](https://karabiner-elements.pqrs.org/) first — that
install is the only manual step. The rest below (copy the logger to a stable spot,
stamp the rules) is plain shell, so **if you trust Claude, just ask it to do this
setup for you** — paste this section or say *"do the a11y-vo-log Karabiner setup"*
and it runs the commands. It only writes to `~/.local/bin/` and
`~/.config/karabiner/`; you still enable the rule yourself in the Karabiner UI at
the end. Prefer to do it by hand? Wire the bundled rules to a **stable** copy of the
gesture logger — stable so plugin updates (which land under a versioned path) don't
break Karabiner:

```bash
# Resolve the installed plugin's files (latest version wins)
GEST_SRC="$(find ~/.claude/plugins/cache -path '*/a11y-vo-log/skills/a11y-vo-log/scripts/vo_gesture.sh' | sort | tail -1)"
RULES_SRC="$(find ~/.claude/plugins/cache -path '*/a11y-vo-log/skills/a11y-vo-log/assets/karabiner_vo_gesture_logger.json' | sort | tail -1)"

# a) Copy the gesture logger to a stable location that survives plugin updates
mkdir -p ~/.local/bin
cp "$GEST_SRC" ~/.local/bin/vo_gesture.sh
chmod +x ~/.local/bin/vo_gesture.sh

# b) Stamp that absolute path into the rules and drop them where Karabiner imports from
mkdir -p ~/.config/karabiner/assets/complex_modifications
sed "s#__VO_GESTURE_SH__#$HOME/.local/bin/vo_gesture.sh#g" "$RULES_SRC" \
  > ~/.config/karabiner/assets/complex_modifications/vo_gesture_logger.json
```

Then in **Karabiner-Elements → Settings → Complex Modifications → Add rule**, enable
the *"VoiceOver gesture logger"* rules.

> Re-run step 3 only if you later change `vo_gesture.sh` itself; a normal plugin
> version bump does not require it, because Karabiner points at the stable copy.

## Usage

```
/a11y-vo-log
```

Claude opens the logging Terminal (starting VoiceOver for you if it's off), you run
the buggy flow, then say **"read the log"** and Claude writes up the bugs. Type a
short note + Enter in the log window to record a visual observation as an `ACTION`
line. To turn VoiceOver off again when done, press **Cmd+F5** — the logger stays
quiet while it's off. (Set `VO_NO_AUTOSTART=1` if you'd rather enable VoiceOver
yourself before starting.)
