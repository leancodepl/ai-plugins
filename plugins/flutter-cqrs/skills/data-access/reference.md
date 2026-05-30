# CQRS data-access reference

## CQRS consists of

- Commands - actions that change the state of the system but do not return any data.
- Queries - actions that return some data from the domain but do not change it.
- Operations - actions that both change the system and return data. Use sparingly.

## Naming

- Query examples: `AllUsers`, `UserById`, `OrdersWithProduct`
- Command examples: `CreateOrder`, `UpdatePersonalDetails`, `PayForOrder`

## Calling the CQRS API

- `run` to execute commands: `final result = await cqrs.run(DeleteOwnAccount());`
- `get` to execute queries: `final result = await cqrs.get(AllProjects(sortByNameDescending: args.isDescending));`
- If you use base classes, they probably already call it.

## Base classes

- `QueryCubit<TApiResult, TItem>` for non-paginated queries.
- `PaginatedQueryCubit<TAdditionalData, TApiResult, TItem>` for paginated data.

## Repository boundary

- Avoid implementing repository unless additional logic is required like caching, merging multiple data sources, complex mapping of responses/requests.

## Contract generation

- Contracts are generated via:

```bash
dart run leancode_contracts_generator
dart run build_runner build
```

Run generation if any relevant contract is missing or if requested by the user. Consider no output from the command as a success.

## Operations

Do not create operations unless explicitly requested.
