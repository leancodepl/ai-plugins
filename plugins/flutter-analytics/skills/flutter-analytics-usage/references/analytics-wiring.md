# Analytics Wiring

How the IDs from `<feature_name>_ids.dart` reach an analytics backend. The conventions in `analytics.md` and `analytics-usage.md` assume this wiring exists.

## Check the project first

`AnalyticsId` is not a type from a package. It is a small class that lives in the app. Before adding or reviewing analytics code:

- Find the project's `AnalyticsId` class and the widgets that accept an `analyticsId`, such as the app's button and tappable-card widgets.
- If they exist, follow them exactly, even where they differ from the examples below.
- If they don't exist, set them up as described here before adding feature IDs. Do not invent parameters on Flutter or third-party widgets: `ElevatedButton`, `InkWell` and `ListTile` have no `analyticsId` parameter.

## Packages

The events and interfaces come from [`leancode_analytics_base`](https://pub.dev/packages/leancode_analytics_base): `LeanAnalytics`, `AnalyticsEvent`, `TapAnalyticsEvent` and `LeanAnalyticsRoute`. The backend is a separate package, for example [`leancode_analytics_firebase`](https://pub.dev/packages/leancode_analytics_firebase) or [`leancode_analytics_posthog`](https://pub.dev/packages/leancode_analytics_posthog).

Provide one `LeanAnalytics` instance app-wide, typed as the base interface:

```dart
Provider<LeanAnalytics>(create: (context) => FirebaseLeanAnalytics()),
```

## The `AnalyticsId` class

Define it once, for example in `lib/common/analytics/analytics_id.dart`:

```dart
class AnalyticsId {
  const AnalyticsId(this.name, this.page);

  final String name;
  final String page;

  String toKey() => '$page/$name';
}
```

`page` takes a page ID from a feature's IDs class, for example `BookingAnalyticsId.page`.

## Tap events

A tappable widget takes the ID as a parameter and registers the tap itself, so feature code only passes an ID:

```dart
class AppButton extends StatelessWidget {
  const AppButton({
    super.key,
    required this.analyticsId,
    required this.label,
    required this.onPressed,
    this.analyticsParams = const {},
  });

  final AnalyticsId analyticsId;
  final String label;
  final VoidCallback onPressed;
  final Map<String, Object> analyticsParams;

  @override
  Widget build(BuildContext context) {
    return FilledButton(
      onPressed: () {
        context.read<LeanAnalytics>().register(
          TapAnalyticsEvent(
            key: analyticsId.toKey(),
            label: label,
            params: analyticsParams,
          ),
        );
        onPressed();
      },
      child: Text(label),
    );
  }
}
```

Feature code then reads:

```dart
AppButton(
  analyticsId: BookingAnalyticsId.confirmButton,
  label: l10n(context).booking_confirm,
  onPressed: cubit.confirm,
)
```

Wrap every tappable widget the design system offers, such as buttons, cards and list tiles, the same way. A raw `InkWell` or `GestureDetector` in a feature should get its tap registered through one of these wrappers.

## Page views

Page views are tracked through navigation, not inside page widgets. Make the app's route type implement `LeanAnalyticsRoute` and return the page ID as `id`. Then add `context.read<LeanAnalytics>().navigatorObservers` to the router's navigator observers. With `auto_route`, a custom route builder does this in one place; with `go_router`, pass the observers to `GoRouter(observers: ...)`.

A page is covered when its route reports the page ID from the feature's IDs class. A page widget has nothing to wrap.
