# Patrol Keys

## Core rules

- Assign a key ONLY to widgets involved in testing.
- ONLY add the `key` parameter to existing widgets — never modify widget signatures or refactor existing code structure.
- Never hardcode keys in the app — always use keys from the `keys.dart` file.
- Never create a key that is not assigned to a widget.
- Always make sure each key value is unique.
- Sort keys alphabetically.
- Add keys as the **first parameter** to the widget constructor.

## File structure

- Create `keys.dart` in the feature/widget directory — NOT in the main `lib/` directory.
- Maximum one `keys.dart` per feature directory (the file can contain multiple classes).
- The main `Keys` class in `lib/keys.dart` should only import and aggregate feature-specific keys.
- Create new `keys.dart` files as needed. Add imports for them in the main aggregator.

Example layout:

```
lib/features/home/keys.dart
lib/features/profile/keys.dart
lib/features/auth/keys.dart
lib/common/widgets/keys.dart
lib/keys.dart          ← aggregator only
```

## Key usage pattern

Always assign keys using this exact pattern:

```dart
key: keys.feature.widgetName
```

(except when using parameterized keys — see below)

## Grouping

Group related keys in a class named after the screen (e.g. `HomeKeys`). Use a private `ValueKey<String>` subclass to prefix all key values with the page or widget name.

```dart
// lib/features/home/keys.dart
import 'package:flutter/widgets.dart';

class _HomeKey extends ValueKey<String> {
  const _HomeKey(String value) : super('home_$value');
}

class HomeKeys {
  final menuIconButton = const _HomeKey('menuIconButton');
  _HomeKey navbarItem(String label) => _HomeKey('navbarItem_$label');
}
```

## Main aggregator

```dart
// lib/keys.dart
import 'features/home/keys.dart';
import 'features/profile/keys.dart';
import 'common/widgets/keys.dart';

final keys = Keys();

class Keys {
  final home = HomeKeys();
  final profile = ProfileKeys();
  final widgets = WidgetKeys();
}
```

## Grouping multiple screens per feature

```dart
// lib/features/product/keys.dart
import 'package:flutter/widgets.dart';

class _ProductPageKey extends ValueKey<String> {
  const _ProductPageKey(String value) : super('productPage_$value');
}

class ProductPageKeys {
  final menuIconButton = _ProductPageKey('menuIconButton');
  final productImage = _ProductPageKey('productImage');
  final productName = _ProductPageKey('productName');
}

class _ProductConnectingPageKey extends ValueKey<String> {
  const _ProductConnectingPageKey(String value)
    : super('productConnectingPage_$value');
}

class ProductConnectingPageKeys {
  final productName = _ProductConnectingPageKey('productName');
  final productImage = _ProductConnectingPageKey('productImage');
}
```

## Common widgets

For common widgets used across features, store keys in `WidgetKeys` placed in the directory of the common widget:

```dart
// lib/common/widgets/keys.dart
import 'package:flutter/widgets.dart';

class _WidgetKey extends ValueKey<String> {
  const _WidgetKey(String value) : super('widget_$value');
}

class WidgetKeys {
  final addButton = const _WidgetKey('addButton');
  final searchBar = const _WidgetKey('searchBar');
  final saveButton = const _WidgetKey('saveButton');
  final cancelButton = const _WidgetKey('cancelButton');
  final topBarHeaderMiddleText = const _WidgetKey('topBarHeaderMiddleText');
  final pickCurrencyButton = const _WidgetKey('pickCurrencyButton');
  _WidgetKey assetSlider(SelectAssetType type) =>
      _WidgetKey('assetSlider_$type');
  ValueKey<String> assetRow(String coinTitle) =>
      _WidgetKey('assetRow_${coinTitle}');
}
```

## External packages

For widgets in a separate package (e.g. widgetbook), create a `keys.dart` in that package:

```dart
// widgetbook/keys.dart
import 'package:flutter/widgets.dart';

final widgetBook = WidgetBookKeys();

class _WidgetBookKey extends ValueKey<String> {
  const _WidgetBookKey(String value) : super('widgetBook_$value');
}

class WidgetBookKeys {
  final tile = const _WidgetBookKey('tile');
}
```

Import to the main aggregator:

```dart
import 'package:common_ui/widgets/keys.dart' as ds;

final keys = Keys();

class Keys {
  final designSystem = ds.dsKeys;
  // ...
}
```

## Parameterized keys

Use parameterized keys when:

- Widgets are generated from dynamic data.
- Widgets are generated from a DTO or enum.
- Number of widgets is variable or large.
- Widgets are generated in loops or from lists.

Use individual keys when:

- Widgets are hardcoded and known at compile time.
- Widgets have distinct, meaningful names.

### Parameterized key rules

- ALWAYS prefer using enums or DTOs as the parameter if they already exist.
- Use existing widget properties for parameterized keys.
- NEVER assign a parameterized key in the app but use fixed values in the keys file (and vice versa).
- NEVER create helper methods — use parameterized keys instead.
- When widgets are generated from existing enums or DTOs, always use parameterized keys with those enum/DTO values.

### Example

```dart
// app widget
enum SizeDTO { small, medium, large }

Widget _sizeButton(SizeDTO size) {
  return _Button(
    key: keys.pickSize.sizeButton(size),
    value: size,
    currentValue: value,
    onPressed: onPressed,
  );
}

@override
Widget build(BuildContext context) {
  return Row(
    crossAxisAlignment: CrossAxisAlignment.start,
    children: [
      _sizeButton(SizeDTO.small),
      _sizeButton(SizeDTO.medium),
      _sizeButton(SizeDTO.large),
    ].spaced(24),
  );
}
```

```dart
// keys.dart
_PickSizeKey sizeButton(SizeDTO size) =>
    _PickSizeKey('sizeButton_${size.name}');
```

## Steps to assign a key to a widget

1. Identify the widget that is needed for testing.
2. Create the key assignment: `key: keys.feature.widgetName`.
3. Define the key in the feature's `keys.dart` file.
4. Verify the key is assigned to the widget.
