# State Management

This app uses BLoC/Cubit pattern for state management.

## Bloc and Cubits

- Use predefined base Cubits if possible.
- Use `flutter_bloc` package handlers.
- Never reference `BuildContext` inside `Blocs` and `Cubits`.
- Prefer using cubits for simple cases.
- Define all `Bloc`/`Cubit` related classes including bloc/cubit itself, events, states and presentation events in a single file.
- Use `enum` to define simple states and events handled by `Bloc`/`Cubit`, use sealed class for more complicated cases.
- All `Bloc`/`Cubit` states should have `EquatableMixin` mixed in and `props` field overridden.
- Use `copyWith` from `copy_with_extension` package for creating new, altered instances of the state.
- Use LeanCode-provided base classes if possible.
- Use `leancode_cubit_utils` for supported use cases such as simple data fetching from the backend.

## Bloc Presentation

- Use `bloc_presentation` package, especially `BlocPresentationMixin`, for presentation side-effects including navigation, dialogs, snackbars, toasts etc.
- Use `enum` to define presentation events.

## Internal state

- Use `flutter_hooks` package for managing presentation state.
- Avoid using `StatefulWidget` at all if possible.
