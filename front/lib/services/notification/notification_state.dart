import 'package:equatable/equatable.dart';

abstract class NotificationState extends Equatable {
  const NotificationState();

  @override
  List<Object> get props => [];
}

class NotificationInitial extends NotificationState {}

class NotificationScheduled extends NotificationState {
  final DateTime scheduledTime;

  const NotificationScheduled(this.scheduledTime);

  @override
  List<Object> get props => [scheduledTime];
}
