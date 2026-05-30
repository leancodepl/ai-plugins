Generate Cubit that performs some user action (e.g. run a command) that can fail.

# Examples

## A cubit that just executes a command, without

cancel_visit_cubit.dart
```dart
// cancel_visit_cubit.dart
import 'package:app/infrastructure/cqrs/contracts/contracts.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:leancode_contracts/leancode_contracts.dart';

class LeanBookingCancelVisitCubit extends Cubit<LeanBookingCancelVisitState> {
  LeanBookingCancelVisitCubit(this._cqrs)
    : super(LeanBookingCancelVisitState.initial);

  final Cqrs _cqrs;

  Future<void> cancelVisit({required String visitId}) async {
    if (state == LeanBookingCancelVisitState.loading) {
      return;
    }
    emit(LeanBookingCancelVisitState.loading);

    final result = await _cqrs.run(CancelReservation(reservationId: visitId));

    if (isClosed) {
      return;
    }

    switch (result) {
      case CommandSuccess():
        emit(LeanBookingCancelVisitState.success);
      case CommandFailure():
        emit(LeanBookingCancelVisitState.error);
    }
  }
}

enum LeanBookingCancelVisitState { initial, loading, success, error }
```

### A cubit that executes a command, but then gathers some data

reserve_timeslot_cubit.dart
```dart
// reserve_timeslot_cubit.dart
import 'package:app/infrastructure/cqrs/contracts/contracts.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:leancode_contracts/leancode_contracts.dart';

class LeanBookingReserveTimeslotCubit
    extends Cubit<LeanBookingReserveTimeslotState> {
  LeanBookingReserveTimeslotCubit(this._cqrs)
    : super(const LeanBookingReserveTimeslotStateInitial());

  final Cqrs _cqrs;

  Future<void> reserveTimeSlot({
    required String timeslotId,
    required String calendarDayId,
  }) async {
    if (state is LeanBookingReserveTimeslotStateLoading) {
      return;
    }
    emit(const LeanBookingReserveTimeslotStateLoading());

    final result = await _cqrs.run(
      ReserveTimeslot(timeslotId: timeslotId, calendarDayId: calendarDayId),
    );

    if (isClosed) {
      return;
    }

    await _emitState(timeslotId, result);
  }

  Future<void> _emitState(String timeslotId, CommandResult result) async {
    switch (result) {
      case CommandSuccess():
        final visitResult = await _cqrs.get(
          MyReservationByTimeslotId(timeslotId: timeslotId),
        );

        if (isClosed) {
          return;
        }

        switch (visitResult) {
          case QueryFailure():
            emit(const LeanBookingReserveTimeslotStateError());
            return;
          case QuerySuccess(data: null):
            emit(const LeanBookingReserveTimeslotStateError());
            return;
          case QuerySuccess(:final data?):
            emit(LeanBookingReserveTimeslotStateSuccess(visitId: data.id));
        }
      case CommandFailure():
        emit(const LeanBookingReserveTimeslotStateError());
    }
  }
}

sealed class LeanBookingReserveTimeslotState with EquatableMixin {
  const LeanBookingReserveTimeslotState();

  @override
  List<Object?> get props => const [];
}

final class LeanBookingReserveTimeslotStateInitial
    extends LeanBookingReserveTimeslotState {
  const LeanBookingReserveTimeslotStateInitial();
}

final class LeanBookingReserveTimeslotStateLoading
    extends LeanBookingReserveTimeslotState {
  const LeanBookingReserveTimeslotStateLoading();
}

final class LeanBookingReserveTimeslotStateSuccess
    extends LeanBookingReserveTimeslotState {
  const LeanBookingReserveTimeslotStateSuccess({required this.visitId});

  final String visitId;

  @override
  List<Object?> get props => [visitId];
}

final class LeanBookingReserveTimeslotStateError
    extends LeanBookingReserveTimeslotState {
  const LeanBookingReserveTimeslotStateError();
}
```

### A cubit with presentation events for password change (without base class)

change_password_cubit.dart
```dart
// change_password_cubit.dart
import 'package:app/features/auth/menu/change_password/change_password_form_cubit.dart';
import 'package:equatable/equatable.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:leancode_kratos_client/leancode_kratos_client.dart';

class ChangePasswordCubit extends Cubit<ChangePasswordState> {
  ChangePasswordCubit({required this.kratosClient})
    : super(const ChangePasswordStateInitial());

  final KratosClient kratosClient;

  final formCubit = ChangePasswordFormCubit();

  Future<void> submit() async {
    if (!formCubit.validate()) {
      return;
    }

    emit(const ChangePasswordStateLoading());

    final password = formCubit.password.state.value;
    final result = await kratosClient.updatePassword(password: password);

    switch (result) {
      case UpdateSuccess():
        emit(const ChangePasswordStateSuccess());
      case UpdateRequiresReauthorization():
        emit(const ChangePasswordStateAuthorizationNeeded());
      case UpdateFailure():
        emit(ChangePasswordStateError(result.error));
    }
  }

  @override
  Future<void> close() async {
    await formCubit.close();
    return super.close();
  }
}

sealed class ChangePasswordState with EquatableMixin {
  const ChangePasswordState();

  @override
  List<Object?> get props => [];
}

final class ChangePasswordStateInitial extends ChangePasswordState {
  const ChangePasswordStateInitial();
}

final class ChangePasswordStateLoading extends ChangePasswordState {
  const ChangePasswordStateLoading();
}

final class ChangePasswordStateSuccess extends ChangePasswordState {
  const ChangePasswordStateSuccess();
}

final class ChangePasswordStateError extends ChangePasswordState {
  const ChangePasswordStateError(this.kratosError);

  final KratosMessage? kratosError;

  @override
  List<Object?> get props => [kratosError];
}

final class ChangePasswordStateAuthorizationNeeded extends ChangePasswordState {
  const ChangePasswordStateAuthorizationNeeded();
}
```

### A cubit with some more intricate logic (waiting between two consecutive actions)

resend_button_cubit.dart
```dart
// resend_button_cubit.dart
import 'dart:async';

import 'package:clock/clock.dart';
import 'package:equatable/equatable.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

/// Controls the visibility and clickable state of a resend button.
/// The button becomes visible after a delay and has a cooldown period between resends.
class ResendButtonCubit extends Cubit<ResendButtonState> {
  ResendButtonCubit({
    this.visibleAfterDuration = const Duration(seconds: 1),
    this.delayBetweenResend = const Duration(minutes: 1),
  }) : super(
          ResendButtonState(
            nextResendTime: clock.now(),
            isButtonVisible: false,
          ),
        );

  final Duration visibleAfterDuration;
  final Duration delayBetweenResend;

  late final Timer _timer;

  void init() {
    _timer = Timer(
      visibleAfterDuration,
      () => emit(state.copyWith(isButtonVisible: true)),
    );
  }

  void markSendAction() {
    _timer.cancel();
    emit(state.copyWith(nextResendTime: clock.fromNowBy(delayBetweenResend)));
  }

  @override
  Future<void> close() {
    _timer.cancel();
    return super.close();
  }
}

final class ResendButtonState with EquatableMixin {
  const ResendButtonState({
    required this.nextResendTime,
    required this.isButtonVisible,
  });

  final DateTime nextResendTime;
  final bool isButtonVisible;

  ResendButtonState copyWith({
    DateTime? nextResendTime,
    bool? isButtonVisible,
  }) => ResendButtonState(
    nextResendTime: nextResendTime ?? this.nextResendTime,
    isButtonVisible: isButtonVisible ?? this.isButtonVisible,
  );

  @override
  List<Object?> get props => [nextResendTime, isButtonVisible];
}
```
