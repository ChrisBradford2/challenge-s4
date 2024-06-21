import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../services/hackathons/hackathon_bloc.dart';
import '../../services/hackathons/hackathon_event.dart';
import '../../services/hackathons/hackathon_state.dart';
import '../../services/team/team_bloc.dart';
import '../../services/team/team_event.dart';
import '../../services/team/team_state.dart';

class HackathonDetailPage extends StatefulWidget {
  final String id;
  final String token;

  const HackathonDetailPage({super.key, required this.id, required this.token});

  @override
  _HackathonDetailPageState createState() => _HackathonDetailPageState();
}

class _HackathonDetailPageState extends State<HackathonDetailPage> {
  final Set<int> joinedTeams = Set<int>();

  @override
  Widget build(BuildContext context) {
    if (kDebugMode) {
      print('HackathonDetailPage id: ${widget.id}, token: ${widget.token}');
    } // Debug print

    return Scaffold(
      appBar: AppBar(
        title: const Text('Hackathon Details'),
      ),
      body: MultiBlocProvider(
        providers: [
          BlocProvider(
            create: (context) => HackathonBloc()..add(FetchSingleHackathons(widget.token, widget.id)),
          ),
          BlocProvider(
            create: (context) => TeamBloc(),
          ),
        ],
        child: BlocBuilder<HackathonBloc, HackathonState>(
          builder: (context, hackathonState) {
            if (hackathonState is HackathonLoading) {
              return const Center(child: CircularProgressIndicator());
            } else if (hackathonState is HackathonLoaded) {
              final hackathon = hackathonState.hackathons[0];
              if (kDebugMode) {
                print('Loaded hackathon: ${hackathon.name}, id: ${hackathon.id}');
              } // Debug print

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
                    Expanded(
                      child: ListView.builder(
                        itemCount: hackathon.teams.length,
                        itemBuilder: (context, index) {
                          final team = hackathon.teams[index];
                          final bool isJoined = joinedTeams.contains(team.id);
                          return ListTile(
                            title: Text(team.name),
                            subtitle: Column(
                              crossAxisAlignment: CrossAxisAlignment.start,
                              children: [
                                Text('Members:'),
                                if (team.users != null)
                                  for (var user in team.users!) Text(user.username),
                              ],
                            ),
                            trailing: ElevatedButton(
                              onPressed: isJoined
                                  ? null
                                  : () {
                                context.read<TeamBloc>().add(JoinTeam(team.id, widget.token));
                              },
                              child: Text(isJoined ? 'Joined' : 'Join'),
                            ),
                          );
                        },
                      ),
                    ),
                    BlocListener<TeamBloc, TeamState>(
                      listener: (context, state) {
                        if (state is TeamJoined) {
                          ScaffoldMessenger.of(context).showSnackBar(
                            SnackBar(content: Text(state.message)),
                          );
                          setState(() {
                            joinedTeams.add(state.teamId);
                          });
                        } else if (state is TeamError) {
                          ScaffoldMessenger.of(context).showSnackBar(
                            SnackBar(content: Text(state.error)),
                          );
                        }
                      },
                      child: Container(),
                    ),
                  ],
                ),
              );
            } else if (hackathonState is HackathonError) {
              return Center(child: Text('Error: ${hackathonState.message}'));
            } else {
              return Container();
            }
          },
        ),
      ),
    );
  }
}
