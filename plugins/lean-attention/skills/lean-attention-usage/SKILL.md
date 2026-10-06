---
name: lean-attention-usage
description: Explain the lean-attention plugin (desktop notification and spoken line when a coding agent waits for input) and troubleshoot it or wire it to another agent. Use when the user invokes `/lean-attention-usage`, asks why they got no notification or sound, wants to change or mute the spoken phrase, or wants the same alert in another agent. Manual-only entry point — does not auto-fire.
disable-model-invocation: true
---

# Lean Attention Usage

## What it does

A `Notification` hook runs `scripts/notify.sh` whenever the agent waits for the user (permission
prompt, question, idle session). The script shows a desktop notification titled with the project
name and speaks a short line; it is agent-agnostic and other agents can call it too. Setup, dependencies and the `LEAN_ATTENTION_*` variables are in the
plugin `README.md`; read it rather than restating it.

## How to respond

1. To explain the plugin, summarize the paragraph above and point at the plugin `README.md`.
2. To change or mute the voice, show the `env` example from the README with the variable the user needs.
3. To get pinged somewhere else (phone, chat) when they don't react, point at the README "Escalation" section; the user's own `LEAN_ATTENTION_ESCALATE_COMMAND` is the extension point, there are no built-in channels beyond desktop and voice.
4. To use it from another agent, point at the "Other agents" section of the README; do not invent hook config for an agent the README does not cover.
5. To troubleshoot silence, check in this order and stop at the first failure:
   - The hook is loaded: ask the user to open `/hooks` and look for a `Notification` entry from `lean-attention`. If it is missing, restart the agent.
   - The script runs: `echo '{"message":"test","cwd":"'"$PWD"'"}' | "<plugin root>/scripts/notify.sh"`.
   - The tools exist: Linux `command -v notify-send spd-say jq`, macOS `command -v osascript say jq`.
   - On Linux with no voice: `spd-say -w "test"`; if that is silent, speech-dispatcher has no output module (install `speech-dispatcher-espeak-ng`).
   - On macOS with no banner: System Settings → Notifications → Script Editor.
   - Escalation never fires: run the command from `LEAN_ATTENTION_ESCALATE_COMMAND` by hand with two arguments; then check that a pending escalation file appears under `${XDG_RUNTIME_DIR:-${TMPDIR:-/tmp}}/lean-attention/` after a notification.
