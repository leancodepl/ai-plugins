# Navigation with auto_route

- Use `auto_route` for navigation between screens. Avoid raw `Navigator` APIs for page-to-page navigation unless integrating with a framework API that explicitly requires a `NavigatorState`.
- Define the application's route tree in a single `AppRouter` class annotated with `@AutoRouterConfig()` and extending `RootStackRouter`.
- Keep the router in the app package under `lib/common/navigation/app_router.dart` with `part 'app_router.gr.dart';`. Treat `app_router.gr.dart` as generated output.
- Register every routable page widget with `@RoutePage()`. The generated route name is derived from the page class name; pass `name:` only when intentionally overriding that default.
- Add new pages to `AppRouter.routes` with `AutoRoute(page: SomeRoute.page, path: '/some-path')`.
- Always run `dart run build_runner build` after adding, removing, or renaming routes or route parameters.
- Never navigate using raw path strings. Use generated route classes such as `context.navigateTo(const SettingsRoute())`, `context.pushRoute(DetailsRoute(id: id))`, `context.replaceRoute(const LoginRoute())`, and `context.pop()`.
- Prefer `context.navigateTo(...)` for top-level screen transitions and tab/root destinations. Use `context.pushRoute(...)` for detail pages, overlays, or pages reachable from multiple places.
- Always name path segments using kebab-case and path parameters using camelCase.
- Prefer path parameters for entity identity and other data that belongs in a deep link. Every page should be able to fetch its data from IDs carried in the route.
- Use `@pathParam` on page constructor parameters that map to `:paramName` path segments.
- Use nullable `@queryParam` constructor parameters only for URL-level optional state such as verification codes, callback data, filters, or coordinates that must be deep-linkable.
- Do not use non-serializable constructor arguments for deep-link data. Passing callbacks or in-memory objects as regular constructor parameters is allowed only for flows that do not need URL restoration.
- Use material handlers such as `showDialog()` and `showBottomSheet()` for dialogs and sheets unless the surface must be addressable as a route or deep link.

## Adding a route

- Annotate the page with `@RoutePage()` and expose the page class globally.
- Import the page in `app_router.dart`.
- Add an `AutoRoute` entry to `AppRouter.routes`.
- Add or update the route's `PageId` mapping if page analytics are used.
- Update the `PageId` enum for the new page if needed.
- Regenerate `app_router.gr.dart` with build_runner.
- Update navigation calls to use the generated route class, not a string path.

## Navigation shell

- Use `AutoTabsRouter` for bottom-tab or shell navigation.
- Model the shell as a regular page annotated with `@RoutePage()` and add it to `AppRouter.routes` with child `AutoRoute`s for each tab.
- Keep the `AutoTabsRouter.routes` list in the same order as the shell's child routes in `AppRouter.routes`.
- Switch tabs through `AutoTabsRouter.of(context).setActiveIndex(index)` or an app navigation bar controller connected to the tabs router.
- Put shell tab paths under the shell route and use top-level absolute paths for authenticated detail pages that should sit outside the tab stacks.
