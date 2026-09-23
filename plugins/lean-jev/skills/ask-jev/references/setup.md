# The API key, and what each failure means

`scripts/jev.py` talks to TypeSafe's `POST /v1/systemone` directly over stdlib `urllib` — no
server to register, no binary to install, nothing to keep updated. Setup is one file.

This file duplicates the setup section of the plugin `README.md` on purpose: its whole job is to
be printed in the session where something went wrong, when the reader is not on GitHub. The README
is canonical and carries the fuller account of cost, data egress, and why the script reads the key
itself; keep the two in step.

## Write the key file

Get a key at [console.typesafe.ai](https://console.typesafe.ai/). Create one **dedicated to this
integration**, so revoking it costs nothing else — rotation is the real safety net, more reliable
than trying to hide a key from a process that has a shell.

```bash
mkdir -p ~/.config/typesafe && chmod 700 ~/.config/typesafe
printf 'export TYPESAFE_API_KEY=%s\n' 'your-key' > ~/.config/typesafe/env && chmod 600 ~/.config/typesafe/env
```

That is the whole setup. The script reads the file itself, so the key stays out of client configs,
out of shell profiles, and out of the command line. A `TYPESAFE_API_KEY` already in the
environment wins over the file, which is handy for a one-off and worth remembering when an answer
looks like it came from the wrong account.

Never put the key in a command you run in the transcript, and never echo it.

## Verify

```bash
python3 <plugin-root>/scripts/jev.py key-status
python3 <plugin-root>/scripts/jev.py ask \
  --state 'The export button crashes the settings page in Safari.' \
  --questions '{"is_bug": {"type": "noul", "instructions": "Does this report a software defect?"}}'
```

`key-status` reports whether a key was found and which source it came from, never the key itself.
A `noul` near 1 from the second command means the whole path works.

## Troubleshooting

Find the symptom, not the step you think you skipped.

### `python3: can't open file .../scripts/jev.py`

The path is wrong, not the setup. In a plugin install the script sits at
`<plugin-root>/scripts/jev.py`, where `<plugin-root>` is what `${CLAUDE_PLUGIN_ROOT}` expands to;
standalone, it sits next to the skill. `ls` the directory before guessing again.

### `{"key": "missing"}` from `key-status`

`~/.config/typesafe/env` is absent or unreadable. Write it as above, and check `ls -l` shows
`-rw-------` and a non-zero size.

### 401 from TypeSafe

A key was found and the API rejected it. Either it is not a current key — check
[console.typesafe.ai](https://console.typesafe.ai/) — or a stale `TYPESAFE_API_KEY` in the
environment is shadowing the file. `key-status` says which source was used.

### 422 from TypeSafe

The request was malformed and the API says how. Most of these are criteria shapes: a `score` takes
an ordered array of levels, a `choice` takes a map of option id to description, and a `noul` takes
an optional map describing its true and false sides. The script checks those locally, so a 422 that
reaches you is something it does not yet know about — read the `detail` in the message.

### The answers look like they came from a different account

Something set `TYPESAFE_API_KEY` in the environment, and it wins over the file. `key-status`
reports `"source": "environment"` when that happens; unset it to fall back to the file.

## What leaves the machine

Whatever goes in `state`. For source code under a client NDA, that is a decision to make
deliberately rather than discover later. TypeSafe states it does not train on customer requests;
zero-data-retention is enterprise-only.
