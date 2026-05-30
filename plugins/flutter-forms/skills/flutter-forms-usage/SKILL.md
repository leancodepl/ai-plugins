---
name: flutter-forms-usage
description: Explain what the `flutter-forms` plugin does and how to use it. Use when the user invokes `/flutter-forms-usage`, asks what this plugin covers, or needs help with `leancode_forms`, form naming, validation behavior, or provisioning.
---

# Forms Usage

## How to respond

- If the user invoked this skill without a concrete task, start with a short explanation of what this plugin does and when it should be used.
- Show how to use it through concrete form-related tasks this plugin can handle.
- Point to the main forms rule here, and route to related plugins when the task expands into DI or general cubit architecture.
- If the user already gave a concrete forms task, briefly explain why this plugin fits and then do the work.
- Do not reply with filler like "skill loaded", "ready for the task", or "what would you like to do?" before explaining the plugin.

## What this plugin does

- Covers LeanCode conventions built around `leancode_forms`.
- Keeps form class naming and validation behavior consistent.
- Guides how form cubits should be provisioned with `BlocProvider`.

## How to use it

- Ask to implement a new form feature.
- Ask to refactor validation behavior.
- Ask to review whether a form follows LeanCode naming and provisioning rules.
- Ask when form-specific work should move into DI or broader BLoC guidance.

## Example requests

- "Create a `leancode_forms` flow for the profile edit screen."
- "Review this form cubit setup and validation behavior."
- "Does this form provisioning belong here or in `flutter-di`?"

## Reach for these assets

- `references/forms.md` - core forms conventions.
- Related plugin: `flutter-di` - page-root provisioning and provider scope decisions.
- Related plugin: `flutter-bloc` - cubit and bloc state-management patterns outside the form-specific rules here.
