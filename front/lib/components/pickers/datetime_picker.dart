import 'package:flutter/material.dart';

Future<void> selectDateTime(BuildContext context, TextEditingController controller) async {
  DateTime? pickedDate = await showDatePicker(
    context: context,
    initialDate: DateTime.now(),
    firstDate: DateTime(2000),
    lastDate: DateTime(2101),
  );

  if (pickedDate != null) {
    if (!context.mounted) return; // Check if the context is still valid
    TimeOfDay? pickedTime = await showTimePicker(
      context: context,
      initialTime: TimeOfDay.now(),
    );

    if (pickedTime != null && context.mounted) { // Check again after the async call
      final pickedDateTime = DateTime(
        pickedDate.year,
        pickedDate.month,
        pickedDate.day,
        pickedTime.hour,
        pickedTime.minute,
      );
      controller.text = pickedDateTime.toIso8601String().replaceFirst('T', ' ').split('.').first;
    }
  }
}
