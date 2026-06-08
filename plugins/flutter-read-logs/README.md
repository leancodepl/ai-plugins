# flutter-read-logs

LeanCode Flutter plugin that lets Claude read the running app's `flutter run` output as
on-demand context. Instead of pasting terminal logs, you run the app, then ask:

```
/read-logs why is the login screen stuck
```

Claude reads the most recent run's logs and works your task against them — it's a context
loader, not an auto-analyzer.

## Included assets

- `skills/read-logs/SKILL.md` — the `/read-logs` workhorse: resolves the per-project log,
  auto-detects format, reads task-led, flags stale runs, and guides first-run setup.
- `skills/flutter-read-logs-usage/SKILL.md` — explains the plugin and routes to setup.

## How it works

Each project logs to its own file, `/tmp/flutter-<repo>.log`, where `<repo>` is derived
from the shared `.git` (`git rev-parse --git-common-dir`) — so the main checkout and every
git worktree resolve to the **same** file. The file is overwritten on every launch (always
the latest run) and lives in `/tmp` — outside the repo, nothing to gitignore, cleared on
reboot. `/read-logs` derives the path at runtime and auto-detects the format.

The only per-developer step is making your editor write that file — **do it once.** If you
run `/read-logs` before setting it up, the skill walks you through it (and offers to apply
the change). It only edits **local** config; it never commits anything.

## ⚠️ Security / data handling — read this first

**Reading a run sends its contents to the model. That *is* the leak — it's how the tool
works, and it can't be prevented.** Run logs routinely contain auth/refresh/push tokens,
emails, account details, and other customer data. The file stays local in `/tmp`, but the
moment `/read-logs` reads it, that slice goes to the model.

Treat this as a **conscious decision**:

- **Enabling capture** (the one-time setup below) is your opt-in — the skill confirms it
  with you before wiring anything up.
- **Each `/read-logs`** announces the file it's about to read before reading it.
- **Prefer test/staging data** when you'll `/read-logs`; avoid production or real-customer
  runs unless you've accepted the exposure.

If you need to read production-shaped logs, the durable mitigation is a **redaction pass**
(mask tokens/emails/keys before the model sees them) — not yet built; tracked as a possible
fast-follow. Raise it if your team wants it before adopting this widely.

In the snippets below, replace `myapp` with your repo's folder name (it must match the path
`/read-logs` derives — i.e. `/tmp/flutter-<repo>.log`).

## VS Code / Cursor (one-time) — keeps F5

Add to your `.vscode/settings.json` (gitignored in most projects, so it stays local —
correct, since the path is machine-specific):

```json
{
  "dart.dapLogFile": "/tmp/flutter-myapp.log",
  "dart.maxLogLineLength": 1000000
}
```

- Keeps your normal debug workflow — **F5** and **Ctrl+F5** ("Run Without Debugging") both
  capture, because both run through the debug adapter. Breakpoints intact, no terminal hacks.
- `maxLogLineLength` defaults to 2000 and truncates long lines; bump it so output and stack
  traces aren't cut.
- Reload the window after adding (`Developer: Reload Window`) so Dart-Code picks it up.
- **Cursor** is a VS Code fork using the same Dart-Code extension and `.vscode/settings.json`
  — identical setup.
- Do **not** use `dapLogFile`'s `${workspaceName}` variable — it expands to the workspace
  folder (differs per worktree) and won't match how `/read-logs` derives the path.

**Caveats:**

- `dapLogFile` captures **debug-adapter sessions** (F5 / Ctrl+F5). If you launch via a plain
  terminal Run Task instead, use the Zed/`script` approach below.
- App logs only reach the debug console (and the file) on devices that forward them: **iOS
  simulator / Android / macOS / the `chrome` device**. A **`-d web-server`** run sends app
  logs to the *browser's* DevTools console, not the editor — so the file will have
  build/launch output but no app logs.
- Historical note: the old `dart.flutterRunLogFile` was **removed** from Dart-Code with the
  legacy debug adapters; `dapLogFile` is its replacement.

## Zed / terminal run task (one-time) — clean transcript

Wrap your run task's command in `script` so output is teed to the file while keeping the
interactive console (hot reload). **`script`'s syntax differs by OS.**

**macOS (BSD `script`)** — logfile, then the command as trailing args, in `.zed/tasks.json`:

```json
{
  "label": "Run app",
  "command": "script",
  "args": ["-q", "/tmp/flutter-myapp.log", "flutter", "run", "-t", "lib/main.dart"],
  "use_new_terminal": true
}
```

**Linux (GNU `script`)** — the command goes through `-c` as a single string, logfile **last**:

```json
{
  "label": "Run app",
  "command": "script",
  "args": ["-q", "-c", "flutter run -t lib/main.dart", "/tmp/flutter-myapp.log"],
  "use_new_terminal": true
}
```

(Wrap whatever your existing run command is — including `fvm flutter …`, flavors, and
`--dart-define`s — don't replace it.) **Windows** has no `script`; use the VS Code/Cursor
`dapLogFile` path above, which is cross-platform.

`script` runs Flutter under a pseudo-TTY, so hot reload (`r`) / hot restart (`R`) keep
working — unlike a plain `… | tee` pipe, which makes Flutter drop to non-interactive mode.
This produces a clean text transcript (no JSON parsing needed). Web logs are captured here
too, as long as the task uses `-d chrome` (not `-d web-server`).

## Note on duplicated setup

These setup steps also appear **inline** in `skills/read-logs/SKILL.md`, on purpose: that
skill applies them at runtime, when the developer is in their editor and not reading this
README. This `README.md` is the canonical copy — keep the two in sync when changing setup.

## Example usage

- `/read-logs <task>` — read the latest run's logs as context for `<task>`.
- `/flutter-read-logs-usage` — short explanation of what this plugin does and how to set it up.

## Related plugins

- [`flutter-marionette`](../flutter-marionette/) — the **proactive** counterpart: it drives
  a *live* app (taps, navigation, hot reload, query state) via MCP, needing instrumentation
  and a running connection. `flutter-read-logs` is **reactive** — it reads what a run
  *already produced*, including build failures and crashed/exited runs, with no
  instrumentation or live app. Rule of thumb: **marionette to *make* things happen,
  read-logs to *see what happened*.** They're complementary — each catches what the other
  can't (marionette: live state; read-logs: build/crash output before any connection).
