import 'package:equatable/equatable.dart';

abstract class NotificationEvent extends Equatable {
  const NotificationEvent();

  @override
  List<Object> get props => [];
}

class ScheduleNotification extends NotificationEvent {
  final DateTime scheduledTime;
  final String title;
  final String body;

  const ScheduleNotification(this.scheduledTime, this.title, this.body);

  @override
  List<Object> get props => [scheduledTime, title, body];
}
