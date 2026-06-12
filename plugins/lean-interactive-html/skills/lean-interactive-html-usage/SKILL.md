---
name: lean-interactive-html-usage
description: Explain what the `lean-interactive-html` plugin does and how to use it. Use when the user invokes `/lean-interactive-html-usage`, asks what this plugin covers, or needs help generating interactive HTML reports, dashboards, plans, or architecture docs.
---

# Interactive HTML Usage

## How to respond

- If the user invoked this skill without a concrete task, explain what the plugin does and when to reach for `/interactive-html`.
- If they already have something to visualize — a plan, an audit, an architecture, stats — briefly explain the fit, then run `/interactive-html`.
- Do not reply with filler like "skill loaded" or "ready for the task" before explaining the plugin.

## What this plugin does

- Generates polished, single-file interactive HTML artifacts: reports, stats dashboards, architecture diagrams, migration/feature plans, decision questionnaires, and process proposals — via `/interactive-html [archetype] <topic>`.
- Every artifact is fully self-contained (no CDN, no frameworks, no build step), so it opens offline, behind VPNs, and from Slack on anyone's machine.
- Ships a copy-first `skeleton.html` (design tokens, tabs, stat cards, badges, filterable tables, localStorage persistence, copy-to-clipboard) plus per-archetype recipes, so artifacts share a consistent house style instead of being reinvented each time.

## When to use it

- "Summarize this audit / PR review as an interactive HTML report."
- "Draw the architecture of this system as a clickable diagram."
- "Turn this migration plan into a browsable briefing with progress checkboxes."
- "Give me an HTML questionnaire for these decisions so I can click through and paste answers back."
- "Make an interactive proposal of this process change for the team."

## Heads-up

These files get forwarded beyond their original audience — the skill enforces English for any shared artifact and forbids embedding secrets, internal hostnames, or credentials. Every displayed number must trace to a command that was actually run.
