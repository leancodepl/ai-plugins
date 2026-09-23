# Changelog

## 0.1.1

- `references/dependency-injection.md` no longer tells Claude to subscribe with `BlocBuilder`, `BlocListener` or `BlocConsumer`, which contradicted `flutter-bloc` (`context.watch` instead of a top-level `BlocBuilder`, `bloc_presentation` for side effects). It now points to `flutter-bloc` for state reads.

## 0.1.0

- Initial `flutter-di` plugin in the public marketplace.
