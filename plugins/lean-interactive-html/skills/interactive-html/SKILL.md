---
name: interactive-html
description: Generates a polished single-file interactive HTML artifact — report, stats dashboard, architecture diagram, migration or feature plan, decision questionnaire, process proposal, PR-review or audit summary. Use whenever the user asks for an "interactive html", an HTML report/doc/dashboard/plan, says "show me this in html" or "summarize this as html", or wants a clickable, browsable artifact to review or send to a PO, teammate, or auditor. Always self-contained — no CDN, no frameworks, no build step.
argument-hint: "[report|architecture|plan|questionnaire|proposal] <topic or data source>"
---

# Interactive HTML Artifacts

Produce a single self-contained HTML file the user can open, click through, and send to someone. The artifact is a deliverable, not scratch — people forward these to POs, teammates, and auditors.

## Workflow

1. Classify the request into an archetype (table below) from `$ARGUMENTS`; if no archetype is given, infer it from intent and proceed — don't ask.
2. Decide audience, language, and destination (rules below). If the file will be committed next to existing HTML docs in the target repo, read one of them first and match its style — new repo docs are expected to match the house style.
3. Gather REAL data: run the `git`/`gh`/`grep` commands that produce every number and claim you will display. Keep the command list — each stat must survive a "where did you take this from?" challenge.
4. Copy `${CLAUDE_PLUGIN_ROOT}/skills/interactive-html/references/skeleton.html` and adapt — never write the boilerplate from scratch. Read the matching section of [references/archetypes.md](references/archetypes.md) for the chosen archetype.
5. Embed data as an inline `<script type="application/json">` block or a JS `const` near the top of the script — one place to correct when numbers change.
6. Write the file to the destination, then ALWAYS open it in the browser (`open <file>` on macOS, `xdg-open` on Linux). The user should never have to ask.
7. Iterate on feedback (often screenshots). Corrections to facts beat corrections to style — fix the data source, not just the displayed text.

## Archetypes

| Archetype | Use for | Signature elements |
|---|---|---|
| `report` | audits, PR reviews, stats dashboards, post-mortems | header with meta chips, big-number stat cards, status badges (In place / Partial / Missing), severity colors, filter buttons, collapsible findings, copy-draft button, pure-CSS bar charts |
| `architecture` | system or feature architecture docs | one box per entity with per-entity accent colors, SVG connector lines computed at runtime, click-to-inspect detail panel, legend, layer grouping |
| `plan` | migration plans, feature plans, briefings | tabs with keyboard shortcuts 1–N, overview stat row, filterable table with click-to-expand rows, phase sections, checklist checkboxes persisted to localStorage, collapsible "evidence from code" blocks quoting real files |
| `questionnaire` | decisions the user answers and pastes back into the chat | one card per question with clickable options, sticky bottom bar with live answer summary and a Copy button producing plain text |
| `proposal` | process or convention proposals to a team | tabs plus collapsibles, current-vs-proposed toggle, step-through simulator (button-driven log of the proposed flow); ship as a pair: same-basename `.md` and `.html` with identical content |

## Audience, language, destination

- **Language**: English whenever ANY audience beyond the requester exists (PO, teammate, auditor, anything committed). The chat language must never leak into the artifact.
- **Tone**: informal-professional; jargon only where the audience expects it.
- **Destination**: committed team doc → the repo's `docs/`; personal or shareable deliverable → `~/Downloads/`; ephemeral review artifact → `/tmp/`.

## Rules

- ONE file, fully self-contained. No `<script src>`, no CDN, no frameworks, no fonts fetched from the network, no build step. Vanilla JS, CSS custom properties in `:root`, system font stack. Charts are pure-CSS bars or inline SVG — never Chart.js/mermaid/d3. Recipients open these offline, behind VPNs, from Slack; every external dependency is a way for the file to break on someone else's machine.
- Dark mode: `@media (prefers-color-scheme: dark)` for committed repo docs; FORCED always-dark for files sent to other people — their browser theme is unknown.
- Every number and claim must trace to a command you ran or a file you read. Never estimate, never pad with plausible-sounding stats. Wrong numbers are the #1 rejection reason for these artifacts.
- Long content needs in-page navigation: tabs or a TOC. Use collapsed-by-default `<details>` for anything secondary.
- State the user can change (checkboxes, filters, answers) persists to localStorage, keyed by artifact name.
- Anything the user produces inside the artifact (answers, draft text) gets a copy-to-clipboard button so it can be pasted back into the chat — the artifact is a node in the feedback loop, not a terminal output.
- Say each fact once. No duplicated stats across sections, no decorative boxes you can't explain, no product or codename branding the user didn't ask for.
- Never embed secrets, internal hostnames, credentials, or account identifiers — these files get forwarded beyond the original audience.

## Anti-patterns (every one of these caused a rework)

- Invented or unverified numbers. Derive every stat from a command you ran; list raw contributors only after filtering bots and one-off committers.
- The same information rendered twice in different sections.
- Boxes, labels, or branding nobody asked for. If you can't justify an element from the source data, delete it.
- Wrong framing or imprecise naming — calling a milestone "end of project", saying a web app has a "store". Match the domain's actual vocabulary.
- Chat-language leakage: artifact generated in the chat's language when the audience is international.
- Light-theme-only files sent to other people.
- External `<script src>`/CDN — every surviving artifact is dependency-free.
- Forgetting to open the file after writing it.
- Diagram connector lines crossing under boxes — compute line endpoints from rendered box positions and redraw on resize.
- Writing the HTML from scratch instead of copying `skeleton.html`.

## Reach for these references

- [references/skeleton.html](references/skeleton.html) — base page: design tokens, tab machinery, stat cards, badges, bar charts, collapsibles, localStorage + clipboard helpers. Copy and adapt; don't rewrite.
- [references/archetypes.md](references/archetypes.md) — per-archetype structure, which skeleton blocks to keep or drop, and the nontrivial JS recipes (SVG connectors, step-through simulator, sticky answer bar).
