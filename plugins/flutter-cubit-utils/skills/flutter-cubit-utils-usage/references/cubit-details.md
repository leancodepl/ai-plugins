Generate Cubit to get a single object (or a list without pagination), e.g. details.

# Examples

## A cubit that gets details of a single service provider

It uses `ServiceProviderDetails` query from contracts.

service_details_cubit.dart
```dart
// service_details_cubit.dart
import 'package:app/infrastructure/cqrs/contracts/contracts.dart';
import 'package:app/infrastructure/cqrs/notifications/cqrs_notification.dart';
import 'package:app/infrastructure/cqrs/notifications/reactive_query_cubits.dart';
import 'package:leancode_contracts/leancode_contracts.dart';

class LeanBookingServiceDetailsCubit
    extends
        ArgsReactiveQueryCubit<
          DateTime,
          ServiceProviderDetailsDTO?,
          ServiceProviderDetailsDTO?
        > {
  LeanBookingServiceDetailsCubit({required super.cqrs, required this.serviceId})
    : super('LeanBookingServiceDetailsCubit');

  final String serviceId;

  @override
  List<Stream<CqrsNotification>> get notificationStreams => [
    cqrs.getNotifications<ServiceProviderDetails>(),
  ];

  @override
  ServiceProviderDetailsDTO? map(ServiceProviderDetailsDTO? data) =>
      data == null ? null : _copyWithAvailableTimeSlots(data);

  @override
  Future<QueryResult<ServiceProviderDetailsDTO?>> request(DateTime args) =>
      cqrs.get(
        ServiceProviderDetails(
          serviceProviderId: serviceId,
          calendarDate: CalendarDate.fromDateTime(args),
        ),
      );

  ServiceProviderDetailsDTO _copyWithAvailableTimeSlots(
    ServiceProviderDetailsDTO data,
  ) => ServiceProviderDetailsDTO(
    id: data.id,
    name: data.name,
    description: data.description,
    type: data.type,
    address: data.address,
    location: data.location,
    isPromotionActive: data.isPromotionActive,
    ratings: data.ratings,
    coverPhoto: data.coverPhoto,
    thumbnail: data.thumbnail,
    timeslots: data.timeslots
        .where((timeSlot) => !timeSlot.isReserved)
        .toList(),
  );
}
```
