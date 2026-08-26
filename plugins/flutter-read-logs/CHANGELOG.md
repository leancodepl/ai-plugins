# Changelog

## Unreleased

- `/read-logs` now states upfront what can follow the command: a concrete `argument-hint`
  in the slash-command menu, an examples table in the `README.md`, a "What to put after the
  command" section in `flutter-read-logs-usage`, and a bare `/read-logs` that names the
  options instead of asking an open question.

## 0.1.0

- Initial `flutter-read-logs` plugin.
- `skills/read-logs/SKILL.md` — capture and read the latest `flutter run` output as context, via `/read-logs`. Auto-detects Zed `script` transcripts and VS Code/Cursor `dapLogFile` (DAP) logs, reads task-led, flags stale runs, and guides first-run capture setup.
- `skills/flutter-read-logs-usage/SKILL.md` — explain what the plugin does and route to setup.
