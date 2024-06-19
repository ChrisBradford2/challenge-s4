import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import '../../components/forms/add_hackathon_form.dart';
import '../../services/hackathons/hackathon_bloc.dart';
import '../../services/hackathons/hackathon_event.dart';
import '../../services/hackathons/hackathon_state.dart';

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
    return BlocProvider(
      create: (context) => HackathonBloc()..add(FetchHackathonForUser(token)),
      child: Scaffold(
        appBar: AppBar(
          title: const Text('Mes hackathons'),
        ),
        body: BlocBuilder<HackathonBloc, HackathonState>(
          builder: (context, state) {
            if (state is HackathonLoading) {
              return const Center(child: CircularProgressIndicator());
            } else if (state is HackathonLoaded) {
              if (state.hackathons.isEmpty) {
                return const Center(child: Text('Aucun hackathon trouvé.'));
              }
              return ListView.builder(
                itemCount: state.hackathons.length,
                itemBuilder: (context, index) {
                  final hackathon = state.hackathons[index];
                  return ListTile(
                    title: Text(hackathon['name'] ?? 'Unknown'),
                    subtitle: Text('${hackathon['date'] ?? 'Unknown'} - ${hackathon['location'] ?? 'Unknown'}'),
                  );
                },
              );
            } else if (state is HackathonError) {
              return Center(child: Text('Erreur: ${state.message}'));
            } else {
              return const Center(child: Text('Hello World'));
            }
          },
        ),
        floatingActionButton: FloatingActionButton(
          onPressed: () {
            _showAddHackathonDialog(context, token);
          },
          tooltip: 'Ajouter un hackathon',
          child: const Icon(Icons.add),
        ),
      ),
    );
  }
}
