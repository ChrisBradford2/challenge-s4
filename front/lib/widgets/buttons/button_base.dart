import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class ButtonBase extends StatelessWidget{
  final String text;
  final VoidCallback onPressed;

  const ButtonBase({
    Key? key,
    required this.text,
    required this.onPressed,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
      onPressed: onPressed,
      style: ElevatedButton.styleFrom(
        backgroundColor: Color.fromRGBO(41, 92, 66, 1), // Background color// Text color
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(20.0), // Border radius
        ),
      ),
      child: Padding(
        padding: const EdgeInsets.symmetric(vertical: 10.0, horizontal: 20.0),
        child: Text(
          text,
          style: const TextStyle(
            fontSize: 16.0,
            color: Colors.white
          ),
        ),
      ),
    );
  }
}