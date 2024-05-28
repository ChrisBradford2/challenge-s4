import 'package:flutter/material.dart';

class ProfileScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Profil"),
      ),
      body: Center(
        child: Text(
          "Bienvenue sur votre profil",
          style: Theme.of(context).textTheme.headlineSmall,
        ),
      ),
    );
  }
}
