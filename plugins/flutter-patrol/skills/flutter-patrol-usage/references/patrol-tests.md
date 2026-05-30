# Patrol Tests

Patrol is LeanCode's Flutter-first E2E UI testing framework. It extends `integration_test` with native automation (`$.platform`) and a custom finder system.

`patrol_mcp` (released 2026-03-31) exposes a Dart-based MCP server that wraps `patrol develop`, letting AI agents run tests, capture screenshots, inspect the native UI tree, and iterate on failing tests without manual handoff.

## Order of actions when writing new tests

1. Read provided test steps.
2. Inspect existing modules for functions that can be reused.
3. Consider whether an existing function can be adjusted to match both its existing usage and the new test.
4. Assign test keys to required elements if they are not assigned yet (see `patrol-keys`).
5. Start writing the test: reuse existing modules/functions and put new test steps in a new test file.
6. Write Patrol actions directly in the test file — do not create new methods in modules yet.
7. After the test passes, reorganize the new code into existing or new modules.
8. Recheck the test after refactoring.

## Patrol MCP usage

When the `patrol_mcp` MCP server is configured, use these tools instead of manual `patrol develop` runs:

- `patrol-run({ "testFile": "patrol_test/your_test.dart" })` — run a test and wait for completion. If no session is running, starts a new one; if a session is running, restarts the current test.
- `patrol-screenshot({ "platform": "android" })` or `patrol-screenshot({ "platform": "ios" })` — capture a screenshot for debugging test failures.
- `patrol-quit({})` — gracefully quit the session.
- `patrol-status({})` — check current session status and recent output.
- `patrol-native-tree({})` — fetch the current native UI tree hierarchy (used for writing native interactions or interacting with apps other than the app under test).

**Known limitation:** on iOS, `patrol develop` tends to time out after 360 seconds.

See the plugin README for how to install `patrol_mcp` and register the MCP launcher.

## Test architecture

### Structure

- Use the custom `testApp` wrapper for all tests.
- Each test file should contain **one** test.
- Organize tests using the `Modules`, `System`, and `ApiClients` pattern:
  - **Modules** — one per user-perspective feature (e.g. `Auth`, `Home`, `Downloads`, `Player`).
  - **System** — class for native interactions via `$.platform` that are not part of your app (e.g. enabling airplane mode to test offline mode). Do not put app-specific methods here.
  - **ApiClients** — aggregates API clients used during testing (e.g. backend test client, mail server client, third-party services).

### Method organization in modules

- Do not write comments in methods — use descriptive method names instead.
- Split long methods when they become hard to read and represent multiple nameable logical steps.
- If a series of steps is reused across many tests as a whole (e.g. a long onboarding flow), create a wrapper method in the module that calls private methods for each step.

## Patrol API rules

- Any file that directly uses Patrol APIs (`$()`, `.scrollTo()`, `.tap()`, `.enterText()`, `.waitUntilVisible()`, etc.) must `import 'package:patrol/patrol.dart';`.
- ALWAYS inspect the Patrol API before implementing a test action:
  - Search the codebase for existing Patrol API usage patterns.
  - Check `$.platform` APIs for the specific action.
  - If a method is not found in the codebase, check: https://patrol.leancode.co/.
  - Only implement after confirming the correct API method.
- ALWAYS inspect `$.platform` methods before implementing native test actions.
- Do not use the `flutter_test` package. Use only Patrol API.
- Only run tests with the Patrol MCP server (for a single test) or `patrol test` (for all tests). Never use the `flutter test` command.
- Do not write `patrolSetUp` or `patrolTearDown` methods yourself.

## Action rules

- Do not use `$.pump`, `waitUntilVisible`, or `waitUntilExists` (or other wait methods) before or after `tap`, `scrollTo`, or `enterText`. Patrol handles waiting automatically. Use wait methods only at the end of the test for assertions.
- To find widgets, use only keys.
- Do not write `try/catch` blocks unless absolutely necessary.
- After writing a test, check that it works by running it with the MCP server; fix it if it fails.
- If a test fails because an element was not found, check whether it needs to be scrolled to in the app code and adjust the test if needed.

## Assertion rules

- Do not write assertions after actions. Write them at the end of the test.
- Prefer `waitUntilVisible` as the final assertion.
- Use `expect()` for assertions only when `waitUntilVisible` is not enough.

## Native dialog handling

- ALWAYS handle native dialogs that appear during the flow:
  - Handle dialogs immediately after the action that triggers them.
  - For native permissions, prefer `$.platform.mobile.grantPermissionWhenInUse` over `$.platform.mobile.tap`.

## Code examples

### Module

```dart
// patrol_test/modules/home.dart
import 'package:patrol/patrol.dart';
import 'module.dart';

final class Home extends Module {
  Home(super.$);

  Future<void> navigateToSettings() async {
    await $(keys.home.settingsButton).scrollTo().tap();
  }

  Future<void> searchForItem(String searchPhrase) async {
    await $(keys.home.searchButton).scrollTo().tap();
    await $(keys.home.searchInput).enterText(searchPhrase);
    await $(keys.home.searchSubmitButton).tap();
  }
}
```

### Modules aggregator

```dart
// patrol_test/modules/modules.dart
final class Modules {
  Modules(this._$);
  final PatrolIntegrationTester _$;

  late final home = Home(_$);
  late final auth = Auth(_$);
}
```

### System class

```dart
// patrol_test/modules/system.dart
final class System extends PlatformAutomator {
  System({required super.config});

  Future<void> checkIfNativePlayerIsVisible() async {
    // Implementation
  }
}
```

### ApiClients class

```dart
// patrol_test/modules/api_clients.dart
final class ApiClients {
  final backend = BackendClient();
  final mailpitClient = MailpitClient();
}
```

### Complete test

```dart
testApp('Download a chapter and play it offline', ($, modules, system, apiClients) async {
    await modules.auth.getAuthToken();
    await apiClients.backend.addFavourites();
    await openApp($);
    await modules.home.goToOldTestament();
    await modules.testament.expandBook(bookName: 'Ksiega Rodzaju');
    await modules.testament.chooseChapterOfBook(
      bookName: 'Ksiega Rodzaju',
      chapterIndex: 0,
    );
    await modules.player.expandChapters();
    await modules.player.downloadChapter(chapterIndex: 9);
    await modules.player.waitUntilDownloaded();
    // equivalent of await $.platform.mobile.enableAirplaneMode();
    await system.enableAirplaneMode();
    await modules.player.rollDownChapters();
    await modules.player.closeChapterPlayer();
    await modules.testament.closeTestament();
    await modules.bottomNavigation.goToLibrary();
    await modules.library.goToDownloads();
    await modules.downloads.expandOldTestament();
    await modules.downloads.goToBook(bookName: 'Ksiega Rodzaju');
    await modules.player.checkIfChapterIsCorrect(chapterIndex: 0);
    await modules.player.playCurrentTrack();
    // uses a wrapper method that calls multiple $.platform methods
    await system.checkIfNativePlayerIsVisible();
  });
```
