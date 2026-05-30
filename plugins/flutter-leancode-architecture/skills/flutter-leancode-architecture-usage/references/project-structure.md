# Project Structure

Consider only the `lib` folder inside the app. It is the main Flutter code. It has the following structure:

- `common` - a folder with shared, infrastructural code. You might find important code there, but try not to modify it.
- `data` - folder with API access code. `data/contracts/contracts.dart` is auto-generated file with the API that the app can use.
- `features` - folder with features that follow Feature-based Flutter App Architecture.
- `navigation` - navigation code.
- `resources` - generated resources (assets, translated strings).
- `widgets` - shared widgets, you can consider this a part of design system.

## Feature Structure

- Create dedicated directory for each feature.
- Follow pattern: `features/feature_name`.
- Try not to cross-reference features.
- Put each feature part in the same folder. If a feature consists of multiple pages, put each page in a separate subfolder.

## Common Code

- Try not to modify `common` unless really needed.
- Common code should not depend on features.
- Avoid extracting code to common speculatively.
- Shared UI should live in the design system rather than in `common`.

## Packages

- Custom plugins should be extracted into separate packages.
