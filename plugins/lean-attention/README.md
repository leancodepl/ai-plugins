# lean-attention

Your agent pauses on a permission prompt, a question, or an idle session, and the built-in signal
is easy to miss when it runs in a background terminal or you juggle several sessions. This plugin
shows a desktop notification with the project name and the agent's message, and speaks a short
line.

## Setup

Install it like any other plugin from this marketplace and restart the agent so the hook loads.
There is nothing to configure for the defaults.

What the hook needs on the machine:

| | Notification | Voice | Message text |
| --- | --- | --- | --- |
| Linux | `notify-send` (`libnotify-bin`) | `spd-say` (`speech-dispatcher`) | `jq` |
| macOS | `osascript` (built in) | `say` (built in) | `jq` |

Each piece is optional: without `jq` the notification shows a generic message, and a missing
notifier or speech tool is skipped. On macOS, allow notifications for Script Editor the first time
one appears.

## Configuration

Set these in the agent's environment (the `env` block of its settings file):

| Variable | Default | Effect |
| --- | --- | --- |
| `LEAN_ATTENTION_SPEAK` | `1` | `0` turns the voice off and keeps the notification |
| `LEAN_ATTENTION_PHRASE` | `Your agent is waiting for you` | The spoken line |
| `LEAN_ATTENTION_TITLE` | `Agent` | Notification title, shown as `<title> · <project>` |
| `LEAN_ATTENTION_LANG` | speech-dispatcher default | Linux voice language, e.g. `pl` for a Polish phrase |

```json
{
  "env": {
    "LEAN_ATTENTION_PHRASE": "Agent czeka na ciebie",
    "LEAN_ATTENTION_LANG": "pl"
  }
}
```

## Other agents

`scripts/notify.sh` is agent-agnostic and takes its input in three shapes:

- a JSON payload on stdin with `message` (and optionally `cwd`), as hook-style agents send it;
- a JSON payload as the last argument, as Codex's `notify` sends it;
- plain text arguments, used as the message: `notify.sh "Build finished"`.

Codex, in `~/.codex/config.toml` (its payload has no `message`, so the notification says "Waiting for you"):

```toml
notify = ["bash", "/path/to/lean-attention/scripts/notify.sh"]
```

For any other agent, point its "run a command when the agent stops or needs input" hook at the
script and send one of the shapes above. Only the bundled `Notification` hook is tested so far.

## Assets

- `hooks/hooks.json`: the `Notification` hook that runs the script.
- `scripts/notify.sh`: picks the notifier and voice for the platform; always exits 0, so it never blocks the agent.
- `/lean-attention-usage`: what the plugin does, how to fix a silent setup, and how to wire it to another agent.
