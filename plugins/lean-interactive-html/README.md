# lean-interactive-html

LeanCode plugin that turns analysis into a single-file interactive HTML artifact you can
open, click through, and send to someone. Instead of a wall of markdown, you ask:

```
/interactive-html report summarize this security audit
```

Claude gathers the real data, builds a self-contained HTML file in the house style, and
opens it in your browser. Not Flutter-specific — it works in any repo.

## Archetypes

| Archetype | What you get |
|---|---|
| `report` | audit / PR-review / stats dashboard: stat cards, status badges, filterable findings, copy-draft button |
| `architecture` | clickable system diagram: entity boxes, runtime-drawn SVG connectors, detail side panel |
| `plan` | migration/feature briefing: tabs, phase checklists persisted across reloads, evidence-from-code collapsibles |
| `questionnaire` | decision form: clickable options, live progress bar, copy-answers button to paste back into the chat |
| `proposal` | process proposal: current-vs-proposed toggle plus a step-through simulator of the new flow |

The archetype is optional — Claude infers it from the request when omitted.

## Included assets

- `skills/interactive-html/SKILL.md` — the `/interactive-html` workhorse: classifies the
  request, gathers real data, adapts the skeleton, and opens the result.
- `skills/interactive-html/references/skeleton.html` — the copy-first base page: design
  tokens, dark mode, tabs with keyboard shortcuts, stat cards, badges, pure-CSS bar charts,
  filterable tables, localStorage persistence, and clipboard helpers.
- `skills/interactive-html/references/archetypes.md` — per-archetype structure and the
  nontrivial JS recipes (SVG connectors, step-through simulator, sticky answer bar).
- `skills/lean-interactive-html-usage/SKILL.md` — explains the plugin and routes here.

## Design rules the skill enforces

- **One file, zero dependencies.** No CDN, no frameworks, no fonts from the network, no
  build step. Recipients open these offline, behind VPNs, from Slack — every external
  dependency is a way for the file to break on someone else's machine.
- **Real numbers only.** Every stat must trace to a command Claude actually ran. Invented
  numbers are the #1 rejection reason for these artifacts.
- **Audience-aware.** English for anything with an audience beyond the requester; forced
  dark theme for files sent to other people; no secrets or internal hostnames ever — these
  files get forwarded.
