# leancode_cubit_utils types

There are base classes for CQRS API calls from `leancode_cubit_utils_cqrs`:

- `QueryCubit<TApiResult, TItem>` for use with non-paginated CQRS queries, most of the time `TApiResult = TItem`
- `PaginatedQueryCubit<TAdditionalData, TApiResult, TItem>` for use with paginated CQRS queries

There are base classes from `leancode_cubit_utils`:

- `RequestCubit<TRes, TData, TOut, TError>`
- `ArgsRequestCubit<TArgs, TRes, TData, TOut, TError>`
- `PaginatedCubit<TData, TRes, TResData, TItem>`

To import cubit utils, `import 'package:leancode_cubit_utils_cqrs/leancode_cubit_utils_cqrs.dart';`

## Use base classes

- **USE BASE CLASSES** for Cubits. Implement cubits yourself only if the base classes don't fit the rules.
  - For data retrieval with queries, you MUST use `QueryCubit` or `PaginatedQueryCubit`
  - For multiple queries, you MUST use `QueryCubit`
  - If it is a mix of data retrieval and changing state, ONLY THEN YOU ARE ALLOWED to implement the cubit yourself
  - Only implement custom cubits if you can explicitly justify why none of the base classes work
  - Use `Args*` base classes only if requested explicitly or the arguments are evidently changing. Do not use it if you need to download something "by id" and the ID can be given.
- If the cubit pattern matches base class, **Do not reimplement the data yourself, adjust to the base class**
- If you need sorting and filtering for paginated requests, introduce `*Params` class that keeps the data and has `copyWith` method. Then, use the class as part of the Cubit state (or as part of additional data in the base classes)
- A paginated query will return `PaginatedResult` and will be based on `SortedQuery` or `PaginatedQuery`
- Use LeanCode-provided base classes if possible
- **Do not create model classes**, use DTOs directly
- Use `leancode_cubit_utils` for supported use cases such as simple data fetching from the backend

## Relevant rules

- `cubit-list.md`
- `cubit-details.md`
- `cubit-action.md`
