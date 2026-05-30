# Analytics

## Source of Truth

Keep feature analytics IDs in one `<feature_name>_ids.dart` file. Do not introduce raw analytics strings in page, content, or widget files.

## Analytics ID Files

Define analytics IDs per feature in a `<feature_name>_ids.dart` file in the feature root directory.

```
features/
  booking/
    booking_ids.dart       ← analytics IDs here
    booking_page.dart
    booking_cubit.dart
```

The class is named `<FeatureName>AnalyticsId`:

```dart
class BookingAnalyticsId {
  // Page IDs — static strings, defined first
  static const page = 'BookingPage';
  static const detailsPage = 'BookingDetailsPage';

  // Button IDs — AnalyticsId instances, defined below page IDs
  static final confirmButton = AnalyticsId('ConfirmButton', page);
  static final cancelButton = AnalyticsId('CancelButton', page);
}
```

## Page IDs

Assign analytics IDs to every page in the feature. Naming convention:
- Single-page feature: `page = '<FeatureName>Page'`
- Multi-page feature: `<pageName>Page = '<FeatureName><PageName>Page'`

Page IDs are defined as `static const String` fields at the top of the class.

## Clickable Element IDs

Assign analytics IDs to all buttons and other clickable elements. This includes taps on cards, tiles, icons, rows, and other user-triggered actions.

Define clickable IDs as `AnalyticsId` instances below page IDs.

Single-page feature:
```dart
static final submitButton = AnalyticsId('SubmitButton', page);
```

Multi-page feature (when button names would conflict across pages):
```dart
static AnalyticsId submitButton(String page) => AnalyticsId('SubmitButton', page);
```

Call the factory method from the specific page: `BookingAnalyticsId.submitButton(BookingAnalyticsId.detailsPage)`.
