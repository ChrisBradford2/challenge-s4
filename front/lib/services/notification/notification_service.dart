import 'package:flutter_local_notifications/flutter_local_notifications.dart';
import 'package:timezone/timezone.dart' as tz;
import 'package:timezone/data/latest.dart' as tz;
import 'package:flutter_local_notifications/flutter_local_notifications.dart';

class NotificationService {
  final FlutterLocalNotificationsPlugin _flutterLocalNotificationsPlugin =
  FlutterLocalNotificationsPlugin();

  Future<void> initialize() async {
    // Initialize the plugin
    const AndroidInitializationSettings initializationSettingsAndroid =
    AndroidInitializationSettings('@mipmap/ic_launcher');
    final InitializationSettings initializationSettings =
    InitializationSettings(android: initializationSettingsAndroid);
    await _flutterLocalNotificationsPlugin.initialize(
      initializationSettings,
      // No onSelectNotification callback needed in version 7.0.0+
    );

    // Initialize timezone database
    tz.initializeTimeZones();
  }

  Future<void> showNotificationNow({  required String title,
    required String body,}) async {

    const AndroidNotificationDetails androidNotificationDetails =
    AndroidNotificationDetails(
      'your channel id', // ID de votre canal de notification
      'your channel name', // Nom de votre canal de notification
      channelDescription: 'your channel description', // Description de votre canal de notification
      importance: Importance.max,
      priority: Priority.high,
    );
    const NotificationDetails notificationDetails =
    NotificationDetails(android: androidNotificationDetails);

    await _flutterLocalNotificationsPlugin.show(
      0, // ID de la notification (doit être unique pour chaque notification)
      title, // Titre de la notification
      body, // Corps de la notification
      notificationDetails,
      payload: 'item x', // Payload facultatif pour traiter la notification
    );
  }

  Future<void> scheduleNotification({
    required DateTime scheduledDate,
    required String title,
    required String body,
  }) async {
    final AndroidNotificationDetails androidPlatformChannelSpecifics =
    AndroidNotificationDetails(
      'your channel id', // ID de votre canal de notification
      'your channel name', // Nom de votre canal de notification
      channelDescription: 'your channel description', // Description de votre canal de notification
      importance: Importance.max,
      priority: Priority.high,
    );

    final NotificationDetails platformChannelSpecifics =
    NotificationDetails(android: androidPlatformChannelSpecifics);

    await _flutterLocalNotificationsPlugin.zonedSchedule(
      0, // ID de la notification (doit être unique pour chaque notification)
      title,
      body,
      tz.TZDateTime.from(scheduledDate, tz.local),
      platformChannelSpecifics,
      androidAllowWhileIdle: true,
      uiLocalNotificationDateInterpretation:
      UILocalNotificationDateInterpretation.absoluteTime,
      matchDateTimeComponents: DateTimeComponents.time,
    );
  }
}
