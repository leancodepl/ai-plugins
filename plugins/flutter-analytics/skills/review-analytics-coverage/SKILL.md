---
name: review-analytics-coverage
description: Reviews a Flutter feature or page for missing LeanCode analytics IDs and inconsistent tracking conventions. Use when auditing analytics coverage or fixing missing page/button IDs.
argument-hint: "<feature-directory-or-file-path>"
---

# Review Analytics Coverage

Review the specified feature or file for analytics coverage gaps.

## Checklist

- [ ] Feature has a `<feature_name>_ids.dart` file in the feature root.
- [ ] The class is named `<FeatureName>AnalyticsId`.
- [ ] Every page has a page ID.
- [ ] Every button and clickable element has an `AnalyticsId`.
- [ ] Shared actions across multiple pages use a factory method when page context differs.
- [ ] No raw analytics strings are inlined in widget files.

## Output Format

For each finding, report:

```text
[ANA] Short title
Location: lib/features/booking/booking_page.dart:42
Violation: <what is missing or inconsistent>
Fix: <concrete change to make>
```

## Guardrails

- Focus on missing or inconsistent analytics coverage, not unrelated architecture issues.
- Treat buttons, cards, rows, icon taps, and similar user-triggered actions as clickable elements.
- If coverage already looks complete, say so explicitly.
