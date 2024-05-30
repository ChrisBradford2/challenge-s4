import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:front/services/hackathons/hackathon_bloc.dart';
import 'package:front/services/hackathons/hackathon_event.dart';

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
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          _showAddHackathonDialog(context, token);
        },
        child: const Icon(Icons.add),
        tooltip: 'Ajouter un hackathon',
      ),
    );
  }

  void _showAddHackathonDialog(BuildContext context, String token) {
    final TextEditingController nameController = TextEditingController();
    final TextEditingController descriptionController = TextEditingController();
    final TextEditingController startDateController = TextEditingController();
    final TextEditingController endDateController = TextEditingController();
    final TextEditingController locationController = TextEditingController();
    final TextEditingController maxParticipantsController = TextEditingController();
    final TextEditingController nbOfTeamsController = TextEditingController();

    showDialog(
      context: context,
      builder: (context) {
        return AlertDialog(
          title: const Text('Ajouter un Hackathon'),
          content: SingleChildScrollView(
            child: Column(
              mainAxisSize: MainAxisSize.min,
              children: [
                TextField(
                  controller: nameController,
                  decoration: const InputDecoration(
                    labelText: 'Nom du hackathon',
                  ),
                ),
                TextField(
                  controller: descriptionController,
                  decoration: const InputDecoration(
                    labelText: 'Description',
                  ),
                ),
                TextField(
                  controller: startDateController,
                  decoration: const InputDecoration(
                    labelText: 'Date de début (YYYY-MM-DD)',
                  ),
                ),
                TextField(
                  controller: endDateController,
                  decoration: const InputDecoration(
                    labelText: 'Date de fin (YYYY-MM-DD)',
                  ),
                ),
                TextField(
                  controller: locationController,
                  decoration: const InputDecoration(
                    labelText: 'Lieu',
                  ),
                ),
                TextField(
                  controller: maxParticipantsController,
                  decoration: const InputDecoration(
                    labelText: 'Nombre maximum de participants',
                  ),
                  keyboardType: TextInputType.number,
                ),
                TextField(
                  controller: nbOfTeamsController,
                  decoration: const InputDecoration(
                    labelText: 'Nombre d\'équipes',
                  ),
                  keyboardType: TextInputType.number,
                ),
              ],
            ),
          ),
          actions: [
            TextButton(
              onPressed: () {
                Navigator.of(context).pop();
              },
              child: const Text('Annuler'),
            ),
            ElevatedButton(
              onPressed: () {
                final hackathonData = {
                  'address': locationController.text,
                  'description': descriptionController.text,
                  'end_date': endDateController.text,
                  'latitude': 0,
                  'location': locationController.text,
                  'longitude': 0,
                  'max_participants': int.tryParse(maxParticipantsController.text) ?? 0,
                  'name': nameController.text,
                  'nb_of_teams': int.tryParse(nbOfTeamsController.text) ?? 0,
                  'start_date': startDateController.text,
                };

                final hackathonBloc = BlocProvider.of<HackathonBloc>(context);
                hackathonBloc.add(AddHackathon(token, hackathonData));
                Navigator.of(context).pop();
              },
              child: const Text('Ajouter'),
            ),
          ],
        );
      },
    );
  }
}
