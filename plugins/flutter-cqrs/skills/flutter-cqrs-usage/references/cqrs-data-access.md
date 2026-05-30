# CQRS and Data

CQRS is an approach for structuring domain API, which LeanCode uses to structure API for the application. CQRS APIs are exposed to end-clients. They can call it freely.

CQRS consists of:

- Commands - actions that change the state of the system but do not return any data.
- Queries - actions that return some data from the domain but do not change it.
- Operations - actions that both change the system and return data. Use sparingly.

Each command/query/operation has a corresponding handler that translates the intent to real results.

## Queries

- Queries are named like:
  - `AllUsers`,
  - `UserById`,
  - `OrdersWithProduct`.
- Queries intention is to retrieve data. They don't modify anything.
- A query defines the payload user sends and the response it will get back.
- Queries always succeed. If not every request is valid, `null` can be returned to notify clients.

## Commands

- Commands are named like imperatives, e.g.
  - `CreateOrder`,
  - `UpdatePersonalDetails`,
  - `PayForOrder`.
- Commands are validated and numeric error codes for validation errors are always defined. Commands cannot fail for no (predictable) reason, only if "unexpected thing happened" (e.g. database is down).
- Commands modify data, they don't return anything, even IDs. Thus, if the application flow requires knowing ID after command, it need to be passed to the command (i.e. be client-generated).

## Operations

- Operation naming is more like commands, although exceptions might exist.
- Operations don't have explicit validation, but it might be provided on a operation-by-operation basis.
- They need explicit developer permission to be created. DO NOT create operations if not explicitly requested.

## DTOs

As part of CQRS contracts definitions, there are DTOs, which are classes or enums that are not commands/queries/operations directly.

## Calling the CQRS API

To call the API, you need `Cqrs` object. It has the following methods:

- `run` to execute commands: `final result = await cqrs.run(DeleteOwnAccount());`. `result` is of type `CommandResult`.
- `get` to execute queries: `final result = await cqrs.get(AllProjects(sortByNameDescending: args.isDescending));`. `result` is of type `QueryResult<T>` where `T` is the result of a query.
- If you use base classes, they probably already call it.

## Base classes

There are base classes for CQRS API calls from `leancode_cubit_utils_cqrs`:

- `QueryCubit<TApiResult, TItem>` for use with non-paginated CQRS queries, most of the time `TApiResult = TItem`
- `PaginatedQueryCubit<TAdditionalData, TApiResult, TItem>` for use with paginated CQRS queries

Try to use them instead, if these apply to the request. Examples of the use are specified later.

## Cubit utils
To import cubit utils, `import 'package:leancode_cubit_utils_cqrs/leancode_cubit_utils_cqrs.dart';`

## Contracts
- Contracts are generated via `dart run leancode_contracts_generator` followed by `dart run build_runner build`. Run generation if any relevant contract is missing or if requested by the user. Consider no output from the command as a success.

## Repository
- Avoid implementing repository unless additional logic is required like caching, merging multiple data sources, complex mapping of responses/requests
