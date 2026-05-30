# Analytics Usage

## Pages

Every page must use a page ID from the feature's `<feature_name>_ids.dart` file. Do not inline page-name strings in page widgets.

Single-page feature:

```dart
BookingAnalyticsId.page
```

Multi-page feature:

```dart
BookingAnalyticsId.detailsPage
```

## Clickable Elements

Every button and other clickable element must use an `AnalyticsId` from the feature IDs file.

Good:

```dart
BookingAnalyticsId.confirmButton
BookingAnalyticsId.submitButton(BookingAnalyticsId.detailsPage)
```

Bad:

```dart
AnalyticsId('ConfirmButton', 'BookingPage')
```

## Reuse Across Pages

If the same logical action appears on multiple pages, keep one shared factory method in the IDs file and pass the current page ID into it.

```dart
static AnalyticsId submitButton(String page) => AnalyticsId('SubmitButton', page);
```

This keeps the action name stable while preserving the correct page context.
