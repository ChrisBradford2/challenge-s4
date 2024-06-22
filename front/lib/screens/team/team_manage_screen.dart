import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../services/team/team_bloc.dart';
import '../../services/team/team_event.dart';
import '../../services/team/team_state.dart';
import '../../models/team_model.dart';

class TeamManagePage extends StatelessWidget {
  final Team team;
  final String token;

  const TeamManagePage({super.key, required this.team, required this.token});

  void _leaveTeam(BuildContext context) {
    context.read<TeamBloc>().add(LeaveTeam(team.id, token));
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Manage Team: ${team.name}'),
      ),
      body: BlocListener<TeamBloc, TeamState>(
        listener: (context, state) {
          if (state is TeamLeft) {
            ScaffoldMessenger.of(context).showSnackBar(
              SnackBar(content: Text(state.message)),
            );
            Navigator.pop(context, 'left'); // Return to previous screen
          } else if (state is TeamError) {
            ScaffoldMessenger.of(context).showSnackBar(
              SnackBar(content: Text(state.error)),
            );
          }
        },
        child: Padding(
          padding: const EdgeInsets.all(16.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              const Text('Members:', style: TextStyle(fontSize: 18)),
              const SizedBox(height: 8),
              Expanded(
                child: ListView.builder(
                  itemCount: team.users?.length ?? 0,
                  itemBuilder: (context, index) {
                    final user = team.users![index];
                    return ListTile(
                      title: Text(user.username),
                      subtitle: Text('${user.firstName} ${user.lastName}'),
                    );
                  },
                ),
              ),
              ElevatedButton(
                onPressed: () {
                  // Logique pour rendre le travail
                },
                child: const Text('Submit Work'),
              ),
              ElevatedButton(
                onPressed: () => _leaveTeam(context),
                child: const Text('Leave Team'),
                style: ElevatedButton.styleFrom(
                  backgroundColor: Colors.red, // Couleur du bouton
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
