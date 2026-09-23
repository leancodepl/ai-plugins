# Changelog

## 0.1.1

- `/read-logs` keeps DAP output lines that Dart-Code cut at `dart.maxLogLineLength`. They used to be dropped without a trace; now the surviving part is kept and marked as truncated.
- Note that with VS Code or Cursor, a test run started from the editor also overwrites the log file.
- The skill no longer points to redaction options in the README, since redaction is not built yet, and its description says first-run setup edits local editor config.

## 0.1.0

- Initial `flutter-read-logs` plugin.
- `skills/read-logs/SKILL.md`: capture and read the latest `flutter run` output as context, via `/read-logs`. Auto-detects Zed `script` transcripts and VS Code/Cursor `dapLogFile` (DAP) logs, reads task-led, flags stale runs, and guides first-run capture setup.
- `skills/flutter-read-logs-usage/SKILL.md`: explain what the plugin does and route to setup.
