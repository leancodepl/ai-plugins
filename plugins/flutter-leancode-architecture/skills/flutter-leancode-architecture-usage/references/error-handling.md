# Error Handling

- Return failure objects rather than throw exceptions.
- Always define failure objects as a dedicated class.
- Show precise error messages, hinting to the user what went wrong.
- Implement retry option for the operations that could go wrong.
- Network errors should be handled separately and provide clear feedback to the user.
- Use already defined patterns for displaying errors or define new reusable patterns.
