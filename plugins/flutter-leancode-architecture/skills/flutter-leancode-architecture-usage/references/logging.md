# Logging

- Use `logging/logging.dart` package for logging messages to the console.
- Always catch and log exceptions that are likely to occur (e.g. from external libraries).
- Always log traces of user actions, especially navigation events and business process lifecycle.
- Never log sensitive data such as personal informations and secrets.
- Avoid logging frequent automatic actions, not to clutter the logs.
