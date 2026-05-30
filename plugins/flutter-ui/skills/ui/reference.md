# UI implementation reference

## Practical checklist

- [ ] Use design-system components where they already exist.
- [ ] Avoid direct `material` or `cupertino` widgets for app UI.
- [ ] Put shared UI widgets in the app design system, not in `common`.
- [ ] Use the project's shared loading pattern for the relevant scenario.
- [ ] Use the project's shared error pattern and provide a retry path.
- [ ] Keep user-facing text in the localization setup and resolve it in the presentation layer.
- [ ] Use hooks for local presentation state when controllers, focus, animation, or keys are needed.
