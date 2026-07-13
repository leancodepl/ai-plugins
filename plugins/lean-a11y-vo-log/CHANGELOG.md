# Changelog

## 0.1.0

- Initial `lean-a11y-vo-log` plugin.
- `skills/lean-a11y-vo-log/SKILL.md` — spawn a Terminal logger that captures a timestamped transcript of everything VoiceOver speaks (and each keyboard nav gesture, via Karabiner) during a manual screen-reader session, then diagnose the announcement timeline into an accessibility bug report, via `/lean-a11y-vo-log`. `vo_log.sh` auto-starts VoiceOver if it's off; `start_log_terminal.sh` reuses the launch window instead of opening a second Terminal.
- `skills/lean-a11y-vo-log-usage/SKILL.md` — explain what the plugin does, when to reach for it, and route to the README setup.
