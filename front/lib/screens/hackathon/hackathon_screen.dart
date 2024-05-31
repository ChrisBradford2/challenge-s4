import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import '../../components/forms/add_hackathon_form.dart';

class HackathonScreen extends StatelessWidget {
  final String token;
  final String googleApiKey = dotenv.env['GOOGLE_PLACES_API_KEY'] ?? 'YOUR_FALLBACK_API_KEY'; // Load API key from .env

  HackathonScreen({super.key, required this.token});

  void _showAddHackathonDialog(BuildContext context, String token) {
    showDialog(
      context: context,
      builder: (context) {
        return AlertDialog(
          title: const Text('Ajouter un Hackathon'),
          content: AddHackathonForm(
            token: token,
            googleApiKey: googleApiKey,
          ),
        );
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Mes hackathons'),
      ),
      body: const Center(
        child: Text('Hello World'),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          _showAddHackathonDialog(context, token);
        },
        tooltip: 'Ajouter un hackathon',
        child: const Icon(Icons.add),
      ),
    );
  }
}
