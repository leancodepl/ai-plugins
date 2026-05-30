# Dependency Injection

- Use `provider` package for dependency injection.
- Use `Provider` to tie lifetime to widget tree.
- Inject dependencies via feature entrypoint constructor.
- Scope dependencies to specific widget tree.
- Create provided values directly in the provider's `create` function.
- Provide page-scoped dependencies especially `Cubits` and `Blocs` inside its Page class using `BlocProvider`.
- Use `BlocBuilder`, `BlocListener` or `BlocConsumer` to subscribe to `Cubits` and `Blocs`.
- Use `GlobalProviders` widget in `main.dart` for providing global dependencies.
- Asynchronous initialization of global dependencies that is quick and doesn't require error handling should be handled in the main function before `runApp` is called.
