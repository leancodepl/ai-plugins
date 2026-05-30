# Navigation

LeanCode uses `auto_route` as the default for new projects. Many existing projects still use `go_router`, so before adding or changing navigation, detect which router the project already uses and follow only that variant's conventions.

## Selecting the variant

- Use `navigation-auto-route.md` when `pubspec.yaml` depends on `auto_route`, route files import `package:auto_route/auto_route.dart`, or the project has `AppRouter` annotated with `@AutoRouterConfig()`.
- Use `navigation-go-router.md` when `pubspec.yaml` depends on `go_router`, route files import `package:go_router/go_router.dart`, or the project has `route_tree.dart`, `routes.dart`, `<package>_routes.dart`, `TypedGoRoute`, or `TypedStatefulShellRoute`.
- If both routers are present during a migration, preserve the convention already used by the feature or route tree you are editing. Do not migrate between routers unless explicitly asked.
- If neither router is present in a new project, prefer `auto_route` unless the user or project requirements say otherwise.

## Shared conventions

- Never navigate using raw path strings when a typed route API exists.
- Always name path segments using kebab-case and path parameters using camelCase.
- Prefer path parameters for entity identity and other data that belongs in a deep link.
- Use material handlers such as `showDialog()` and `showBottomSheet()` for dialogs and sheets unless the surface must be addressable as a route or deep link.
