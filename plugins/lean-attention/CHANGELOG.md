# Changelog

## 0.2.0

- `LEAN_ATTENTION_ESCALATE_COMMAND`: your own command, run when nobody reacts within `LEAN_ATTENTION_ESCALATE_AFTER` seconds (default 120, `0` runs it right away). Gets the title and message as `$1` and `$2`.
- `hooks/hooks.json`: `UserPromptSubmit`, `PostToolUse`, `Stop` and `SessionEnd` hooks cancel a pending escalation through `notify.sh --cancel`.
- No new built-in channels; README shows a KDE Connect phone ping as an example.

## 0.1.0

- Initial `lean-attention` plugin.
- `hooks/hooks.json`: a `Notification` hook that runs `scripts/notify.sh`.
- `scripts/notify.sh`: desktop notification (`notify-send` / `osascript`) and a spoken line (`spd-say` / `say`), configurable through `LEAN_ATTENTION_*` env vars. Agent-agnostic input: JSON on stdin, JSON as the last argument (Codex `notify`), or plain text.
- `skills/lean-attention-usage/SKILL.md`: explains the plugin and troubleshoots a silent setup.
