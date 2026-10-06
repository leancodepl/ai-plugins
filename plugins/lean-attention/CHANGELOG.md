# Changelog

## 0.1.0

- Initial `lean-attention` plugin.
- `hooks/hooks.json`: a `Notification` hook that runs `scripts/notify.sh`.
- `scripts/notify.sh`: desktop notification (`notify-send` / `osascript`) and a spoken line (`spd-say` / `say`), configurable through `LEAN_ATTENTION_*` env vars. Agent-agnostic input: JSON on stdin, JSON as the last argument (Codex `notify`), or plain text.
- `skills/lean-attention-usage/SKILL.md`: explains the plugin and troubleshoots a silent setup.
