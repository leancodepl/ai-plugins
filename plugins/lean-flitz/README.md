# lean-flitz

> **Experimental.** New, and not yet run in every setup it documents.

Ask Claude for a build of your Flutter app, and get back a link a teammate opens on their phone.
In Slack with Claude Tag:

> @Claude build a Flitz bundle of the app

Claude runs `flitz publish` in the project and replies with the landing page: a QR code and an
install button for whoever asked.

## What Flitz is

[Flitz](https://docs.flitz.dev) is a tool from LeanCode for putting work-in-progress builds of a
Flutter app on a real device in seconds, without rebuilding, reinstalling, or going through a
store. It splits the app in two:

- **The host** is a normal, signed build of your app with the Flitz loader embedded. You build
  and install it once per device.
- **The bundle** is your Dart code and assets, compiled to portable bytecode. `flitz publish`
  builds one from the project and uploads it, as often as you like.

Opening the publish link on a device that has the host installed downloads the bundle and runs
it. Flitz is for your own team's development and test builds; production ships the normal way.
[How it works](https://docs.flitz.dev/concepts/how-it-works) has the long version.

This plugin covers only the publish step. Setting Flitz up in an app is a one-time job described
in the [Flitz getting-started guide](https://docs.flitz.dev/getting-started).

## Skills

| Skill | What it does |
| ----- | ------------ |
| [`/lean-flitz-usage`](skills/lean-flitz-usage/SKILL.md) | Explains the plugin and where it runs |
| [`/lean-flitz:publish-flitz-bundle`](skills/publish-flitz-bundle/SKILL.md) | Publishes a bundle of the project and returns the landing page |

## Before you start

- **A Flitz organization with an active license.** The publish is a licensed step; without one it
  fails with exit 3 before anything is compiled.
- **A project set up for Flitz:** a `flitz.yaml` pinning the SDK (written by `flitz sdk use`),
  and a host build installed on the device that will open the link. The skill checks for
  `flitz.yaml` and stops when there is none.
- **Linux x64 or an Apple Silicon Mac** to publish from. Windows, Intel Macs, and Linux arm64 are
  not supported by the CLI.

## How it works

1. **Project.** The directory holding `flitz.yaml`. The bundle is built from the working tree as
   it is; nothing is cloned, checked out, or stashed.
2. **CLI.** `flitz` from `PATH`, else the official installer, which verifies the binary's SHA-256.
3. **Credential**, in this order: `FLITZ_APIKEY` from the environment; the stored `flitz login`;
   otherwise a placeholder `FLITZ_APIKEY` for the Claude Tag Agent Proxy to replace. The order
   matters because inside the CLI a set `FLITZ_APIKEY` wins over a login, so the placeholder must
   never shadow one.
4. **Publish.** `flitz publish --yes --json` from the project directory. stdout is the JSON result,
   and the skill reports its `page_url`.

## Setup

### Claude Tag

[Claude Tag](https://claude.com/docs/claude-tag/overview) runs Claude in Slack. An Owner or a
Claude Tag admin sets this up once per bundle (the named set of connectors, allowed domains, and
plugins that covers a channel); the channel then needs nothing else. The steps follow
[Add connections](https://claude.com/docs/claude-tag/admins/add-connections).

1. **Attach the plugin.** Add `lean-flitz` to the bundle that covers the channel.
2. **Allow the domains.** Add `*.flitz.dev` (the CLI download, `api.flitz.dev`, and the bundle
   host `d.flitz.dev`) and `pub.dev`, for Flutter dependencies, as domains on that bundle.
3. **Add a connector** whose allowed websites include `api.flitz.dev`, holding your Flitz
   organization's API key. `flitz apikey status` shows whether a key exists, and
   `flitz apikey create` mints one from a signed-in session (see
   [flitz apikey](https://docs.flitz.dev/cli/apikey)). The Agent Proxy swaps the key into the
   CLI's requests after they leave the sandbox, so the key is never in the sandbox, the
   environment, or the conversation.

The organization has one API key, shared with any CI that publishes. Reuse it for the connector
rather than rotating it, since rotating breaks every pipeline that holds the old one.

Verify by asking in the channel for a bundle of any project with a `flitz.yaml`. A credential
problem surfaces as exit 3.

### Local (Claude Code)

Run `flitz login` once. The CLI is installed on first use, or by hand:

```bash
curl -fsSL https://download.flitz.dev/install.sh | sh
```

In a sandboxed Claude Code session the publish needs to write `~/.local/bin`, `~/.cache/flitz`,
and `~/.pub-cache`, and to reach `*.flitz.dev` and `pub.dev`; allow those or run it outside the
sandbox.

A first publish downloads the pinned SDK (several hundred MB), so it takes a few minutes. Later
ones take about a minute.

## Troubleshooting

| Symptom | Cause | Fix |
| ------- | ----- | --- |
| Skill stops with no `flitz.yaml` found | The app is not set up for Flitz | Follow the [getting-started guide](https://docs.flitz.dev/getting-started), then run `flitz sdk use` |
| Exit 3 with the placeholder key | No Flitz connector in the channel's bundle, or its key is wrong | Add or fix the connector above |
| Exit 3 locally | No login, an expired one, or an inactive license | `flitz login`; for a license, the organization's Owner restores it in the [portal](https://docs.flitz.dev/portal) |
| Exit 3 with `FLITZ_APIKEY` set in the environment | A stale key shadows your login | `unset FLITZ_APIKEY` |
| `curl: (22) … 403` while installing | `*.flitz.dev` is not an allowed domain | Add it to the bundle, or to the local sandbox's allowlist |
| Pub resolution fails during the build | `pub.dev` is blocked | Allow `pub.dev` the same way |
| Exit 4 ending in `no response within 30s` | The upload has a fixed 30-second limit | Re-run from a faster network, or trim large assets |

Every exit code is described in the [Flitz CLI docs](https://docs.flitz.dev/cli/exit-codes).

## What leaves the machine

The app's compiled code and assets go to Flitz, and the returned link downloads them. Links do not
expire and cannot be revoked, and anyone holding one can install the bundle, so the skill posts a
link only where it was asked for.

## Install

```
/plugin install lean-flitz@leancode-ai-plugins
```

See the [root README](../../README.md#install) for adding the marketplace itself.
