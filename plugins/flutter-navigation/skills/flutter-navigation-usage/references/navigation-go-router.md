# Navigation with go_router

- Use `go_router` and its handlers by default for any navigation between screens.
- Always use `go_router` for pushing, popping, and replacing the current route. Never use material `Navigator` for between-screens or between-pages navigation.
- Always define type-safe routes. Where possible, extend `AppGoRoute` when defining new routes.
- Never use raw paths for navigation.
- Define routes in the app main package in `lib/navigation`.
- Define the navigation tree in a separate `route_tree.dart` file.
- Define all routes in a single `routes.dart` file for single-package projects. For multi-package projects, define routes per package in `<package-name>_routes.dart`.
- Page representing widgets should always be globally exposed.
- Always use path parameters in navigation. Never use query parameters for passing required page data.
- Never use `extra` for passing data between pages. Every page should be able to fetch the data it needs based on an ID it got as a path parameter.
- Prefer using `go` over `push` for navigating between pages. Use `push` only for pages that can be accessed from multiple places.
- Always name path segments using kebab-case.
- Always name path parameters using camelCase.
- Always use extensions from `go_router` for navigation, e.g. use `MyRoute(foo, bar).go(context)` over `context.go(MyRoute(foo, bar).location)`.
- Always use material handlers `showDialog()` and `showBottomSheet()` for dialogs and sheets unless deep linking is required.

## Navigation shell

- If a navigation shell is required, use `TypedStatefulShellRoute`. Use `TypedStatefulShellBranch` for shell branches.
- Each branch can have nested routes that maintain its state.
- Use `go_router` navigation methods to switch between branches and routes.
- Deep linking and URL handling should work automatically with the shell route structure.
