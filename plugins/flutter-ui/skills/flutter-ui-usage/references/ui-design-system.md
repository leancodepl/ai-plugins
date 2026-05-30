# UI Design System

## Design-system components

- Use design-system components in the app.
- Avoid using `material` or `cupertino` widgets directly.
- Common UI widgets should be defined in the app design system, not in `common`.

## Loading patterns

- Loading UI patterns should be defined globally and reused across all screens.
- Follow the project pattern for:
  - Full-screen loading
  - Loading part of the screen
  - Loading the next page of paginated lists
  - Loading on user actions

## Error patterns

- Error UI patterns should be defined globally and reused across all screens.
- Provide a retry option to the user.
- Follow the project pattern for:
  - Errors during screen loading
  - Errors on loading part of the screen
  - Errors on loading the next page of paginated lists
  - Errors on user actions

## Presentation layer

- Keep localization in the presentation layer.
- Resolve user-facing text through the project's localization setup.
- Use `flutter_hooks` for presentation state such as `TextEditingController`, `FocusNode`, `AnimationController`, and `GlobalKey`.
