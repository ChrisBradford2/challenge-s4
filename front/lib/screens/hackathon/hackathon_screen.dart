import 'package:flutter/material.dart';

class HackathonScreen extends StatelessWidget {
  final String token;

  const HackathonScreen({super.key, required this.token});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Mes hackathons'),
      ),
      body: const Center(
        child: Text('Hello World'),
      ),
    );
  }
}
