---
name: flutter-patrol-usage
description: Explain what the `flutter-patrol` plugin does and how to use it. Use when the user invokes `/flutter-patrol-usage`, asks what this plugin covers, or needs help with Patrol E2E tests, test keys, native automation, or Patrol MCP.
---

# Patrol Usage

## How to respond

- If the user invoked this skill without a concrete task, start by explaining what this plugin does and when Patrol is the right tool.
- Show how to use it through concrete testing tasks this plugin can handle.
- Point to the relevant rule for tests, keys, or Patrol MCP usage.
- If the user already gave a concrete Patrol task, briefly explain why this plugin fits and then do the work.
- Do not reply with filler like "skill loaded", "ready for the task", or "what would you like to do?" before explaining the plugin.

## What this plugin does

- Covers LeanCode Patrol conventions for E2E UI tests.
- Guides Patrol test structure, assertions, key naming, custom finders, and native automation boundaries.
- Helps with Patrol MCP usage for agent-driven test runs and debugging.

## How to use it

- Ask to create or refactor a Patrol E2E test.
- Ask to add or review test keys for a feature flow.
- Ask how to handle native dialogs, permissions, notifications, or other platform UI.
- Ask how to run or debug a test through Patrol MCP.

## Example requests

- "Create a Patrol test for the login flow."
- "Review this feature for missing Patrol keys."
- "Help me run this test through Patrol MCP."

## Reach for these assets

- `references/patrol-tests.md` - test structure and Patrol API conventions.
- `references/patrol-keys.md` - key naming and composition rules.
- Related plugin: `flutter-bloc` - state-management conventions for cubit-driven flows under test.
- Related plugin: `flutter-navigation` - route structure and navigation flows exercised by Patrol tests.
