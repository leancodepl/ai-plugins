---
name: a11y-vo-log-usage
description: Explain what the `a11y-vo-log` plugin does and how to use it. Use when the user invokes `/a11y-vo-log-usage`, asks what this plugin covers, or needs help with VoiceOver speech logging, gesture capture, or the one-time setup.
---

# a11y-vo-log Usage

## How to respond

- If the user invoked this skill without a concrete task, start by explaining what
  the plugin is for and when it beats a screen recording.
- Point to the one-time setup in the plugin `README.md` (VoiceOver AppleScript
  control, Automation permission, optional Karabiner gesture rules).
- If the user already has a concrete logging task, hand off to the `a11y-vo-log`
  skill and do the work.
- Do not reply with filler like "skill loaded" or "ready for the task" before
  explaining the plugin.

## What this plugin does

- Captures a **timestamped text transcript** of everything VoiceOver speaks during
  a manual screen-reader session on macOS — a text alternative to recording video.
- Optionally interleaves **keyboard nav gestures** (via Karabiner-Elements), so the
  transcript shows exactly what the dev did between announcements — including silent
  focus moves (a gesture with no VO line after it), the #1 screen-reader smell.
- Then **diagnoses** the announcement timeline into an accessibility bug report,
  chaining each bad announcement to the likely widget, file, and fix.
- Is **project-agnostic**: it logs whatever macOS app is frontmost (web, native,
  Flutter).

## When to reach for it

The *dev* drives VoiceOver by hand and Claude reads the resulting transcript. Best
for reproducing a bug the dev already feels, or capturing a real human navigation
flow — without recording and re-watching a video.

## Setup

One-time, per machine — see the plugin `README.md`:

1. VoiceOver → allow AppleScript control.
2. Grant Automation permission to the process running Claude Code (Terminal +
   VoiceOver).
3. (For gesture lines) import the bundled Karabiner rules pointing at a stable copy
   of `vo_gesture.sh`.

## Example usage

- `/a11y-vo-log` — Claude opens the logging Terminal, you run your flow, then say
  "read the log" and Claude writes up the bugs.
