Generate Cubit to get a paginated list of objects, with or without filtering & sorting.

# Examples

## A cubit that gets paginated data

It uses `AllServiceProviders` query from contracts.

most_popular_cubit.dart
```dart
// most_popular_cubit.dart
import 'package:app/infrastructure/cqrs/contracts/contracts.dart';
import 'package:app/infrastructure/cqrs/notifications/cqrs_notification.dart';
import 'package:app/infrastructure/cqrs/notifications/reactive_query_cubits.dart';
import 'package:leancode_contracts/leancode_contracts.dart';
import 'package:leancode_cubit_utils_cqrs/leancode_cubit_utils_cqrs.dart';

class LeanBookingHomeMostPopularCubit
    extends
        PaginatedReactiveQueryCubit<
          void,
          PaginatedResult<ServiceProviderSummaryDTO>,
          ServiceProviderSummaryDTO
        > {
  LeanBookingHomeMostPopularCubit({required super.cqrs})
    : super(loggerTag: 'LeanBookingHomeMostPopularCubit');

  @override
  List<Stream<CqrsNotification>> get notificationStreams => [
    cqrs.getNotifications<AllServiceProviders>(),
  ];

  @override
  Future<QueryResult<PaginatedResult<ServiceProviderSummaryDTO>>> requestPage(
    PaginatedArgs args,
  ) => cqrs.get(
    AllServiceProviders(
      pageNumber: args.pageNumber,
      pageSize: args.pageSize,
      sortBy: ServiceProviderSortFieldsDTO.name,
      sortByDescending: false,
      nameFilter: null,
      typeFilter: null,
      promotedOnly: true,
      favoriteOnly: false,
    ),
  );

  @override
  PaginatedResponse<Set<ServiceProviderSummaryDTO>, ServiceProviderSummaryDTO>
  onPageResult(PaginatedResult<ServiceProviderSummaryDTO> page) =>
      PaginatedResponse.append(
        items: page.items,
        hasNextPage: calculateHasNextPage(
          pageNumber: state.args.pageNumber,
          totalCount: page.totalCount,
        ),
      );
}
```

## A cubit that gets data, supports pagination, filtering and sorting

It uses `AllServiceProviders` query from contracts.

home_cubit.dart
```dart
// home_cubit.dart
import 'dart:async';

import 'package:app/features/home/service_type_filter.dart';
import 'package:app/infrastructure/cqrs/contracts/contracts.dart';
import 'package:app/infrastructure/cqrs/notifications/cqrs_notification.dart';
import 'package:app/infrastructure/cqrs/notifications/reactive_query_cubits.dart';
import 'package:leancode_contracts/leancode_contracts.dart';
import 'package:leancode_cubit_utils_cqrs/leancode_cubit_utils_cqrs.dart';

class LeanBookingHomeParams {
  const LeanBookingHomeParams({required this.typeFilter, required this.sort});

  final ServiceTypeFilter typeFilter;
  final ServiceProviderSortFieldsDTO sort;

  LeanBookingHomeParams copyWith({
    ServiceTypeFilter? typeFilter,
    ServiceProviderSortFieldsDTO? sort,
  }) => LeanBookingHomeParams(
    typeFilter: typeFilter ?? this.typeFilter,
    sort: sort ?? this.sort,
  );
}

class LeanBookingHomeCubit
    extends
        PaginatedReactiveQueryCubit<
          LeanBookingHomeParams,
          PaginatedResult<ServiceProviderSummaryDTO>,
          ServiceProviderSummaryDTO
        > {
  LeanBookingHomeCubit({required super.cqrs})
    : super(
        config: const PaginatedConfig(pageSize: 8),
        initialData: const LeanBookingHomeParams(
          typeFilter: ServiceTypeFilterAll(),
          sort: ServiceProviderSortFieldsDTO.name,
        ),
        loggerTag: 'LeanBookingHomeSearchCubit',
      );

  @override
  List<Stream<CqrsNotification>> get notificationStreams => [
    cqrs.getNotifications<AllServiceProviders>(),
  ];

  @override
  PaginatedResponse<LeanBookingHomeParams, ServiceProviderSummaryDTO>
  onPageResult(PaginatedResult<ServiceProviderSummaryDTO> page) =>
      PaginatedResponse.append(
        items: page.items,
        hasNextPage: calculateHasNextPage(
          pageNumber: state.args.pageNumber,
          totalCount: page.totalCount,
        ),
        data: state.data,
      );

  @override
  Future<QueryResult<PaginatedResult<ServiceProviderSummaryDTO>>> requestPage(
    PaginatedArgs args,
  ) => cqrs.get(
    AllServiceProviders(
      pageNumber: args.pageNumber,
      pageSize: args.pageSize,
      sortBy: state.data.sort,
      sortByDescending: false,
      nameFilter: args.searchQuery,
      typeFilter: state.data.typeFilter.toDto(),
      promotedOnly: false,
      favoriteOnly: false,
    ),
  );

  Future<void> sortBy(ServiceProviderSortFieldsDTO sortField) {
    emit(state.copyWith(data: state.data.copyWith(sort: sortField)));
    return run();
  }

  Future<void> filterBy(ServiceTypeFilter filter) {
    emit(state.copyWith(data: state.data.copyWith(typeFilter: filter)));
    return run();
  }
}
```
