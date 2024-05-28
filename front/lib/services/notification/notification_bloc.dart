import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_local_notifications/flutter_local_notifications.dart';
import 'package:timezone/data/latest.dart' as tz;
import 'package:timezone/timezone.dart' as tz;

import 'notification_event.dart';
import 'notification_state.dart';

class NotificationBloc extends Bloc<NotificationEvent, NotificationState> {
  final FlutterLocalNotificationsPlugin flutterLocalNotificationsPlugin;

  NotificationBloc(this.flutterLocalNotificationsPlugin)
      : super(NotificationInitial()) {
    tz.initializeTimeZones();
    on<ScheduleNotification>(_onScheduleNotification);
  }

  Future<void> _onScheduleNotification(ScheduleNotification event,
      Emitter<NotificationState> emit) async {
    final scheduledTime = tz.TZDateTime.from(event.scheduledTime, tz.local);

    const androidDetails = AndroidNotificationDetails(
      'your_channel_id',
      'your_channel_name',
      channelDescription: 'your channel description',
      importance: Importance.max,
      priority: Priority.high,
    );

    const platformDetails = NotificationDetails(
      android: androidDetails,
      iOS: DarwinNotificationDetails(),
    );

    await flutterLocalNotificationsPlugin.zonedSchedule(
      0,
      event.title,
      event.body,
      scheduledTime,
      platformDetails,
      androidAllowWhileIdle: true,
      uiLocalNotificationDateInterpretation:
      UILocalNotificationDateInterpretation.absoluteTime,
      matchDateTimeComponents: DateTimeComponents.dateAndTime,
    );

    emit(NotificationScheduled(event.scheduledTime));
  }
}