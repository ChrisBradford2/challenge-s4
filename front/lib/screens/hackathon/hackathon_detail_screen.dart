import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import '../../services/hackathons/hackathon_bloc.dart';
import '../../services/hackathons/hackathon_event.dart';
import '../../services/hackathons/hackathon_state.dart';
import '../team/join_team_screen.dart';

class HackathonDetailPage extends StatelessWidget {
  final String id;
  final String token;

  const HackathonDetailPage({super.key, required this.id, required this.token});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Hackathon Details'),
      ),
      body: BlocProvider(
        create: (context) => HackathonBloc()..add(FetchSingleHackathons(token, id)),
        child: BlocBuilder<HackathonBloc, HackathonState>(
          builder: (context, state) {
            if (state is HackathonLoading) {
              return const Center(child: CircularProgressIndicator());
            } else if (state is HackathonLoaded) {
              final hackathon = state.hackathons[0];
              return Padding(
                padding: const EdgeInsets.all(16.0),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text('Name: ${hackathon.name}', style: const TextStyle(fontSize: 24)),
                    const SizedBox(height: 16),
                    Text('Date: ${hackathon.date}', style: const TextStyle(fontSize: 18)),
                    const SizedBox(height: 16),
                    Text('Location: ${hackathon.location}', style: const TextStyle(fontSize: 18)),
                    const SizedBox(height: 16),
                    Text('Description: ${hackathon.description}', style: const TextStyle(fontSize: 18)),
                    const SizedBox(height: 16),
                    const Text('Teams:', style: TextStyle(fontSize: 18)),
                    ListView.builder(
                      shrinkWrap: true,
                      itemCount: hackathon.teams.length,
                      itemBuilder: (context, index) {
                        final team = hackathon.teams[index];
                        return ListTile(
                          title: Text(team.name),
                          trailing: ElevatedButton(
                            onPressed: () {
                              Navigator.push(
                                context,
                                MaterialPageRoute(
                                  builder: (context) => JoinTeamPage(
                                    teamId: team.id.toString(), // Ensure this is a String
                                    token: token,
                                  ),
                                ),
                              );
                            },
                            child: const Text('Join'),
                          ),
                        );
                      },
                    ),
                  ],
                ),
              );
            } else if (state is HackathonError) {
              return Center(child: Text('Error: ${state.message}'));
            } else {
              return Container();
            }
          },
        ),
      ),
    );
  }
}
