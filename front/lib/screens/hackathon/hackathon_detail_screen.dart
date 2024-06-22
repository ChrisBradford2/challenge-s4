import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:jwt_decode/jwt_decode.dart';

import '../../models/hackathon_model.dart';
import '../../models/user_model.dart';
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
  final Set<int> joinedTeams = <int>{};
  late int userId;

  @override
  void initState() {
    super.initState();
    userId = _getUserIdFromToken(widget.token);
    context.read<HackathonBloc>().add(FetchSingleHackathons(widget.token, widget.id));
  }

  int _getUserIdFromToken(String token) {
    Map<String, dynamic> decodedToken = Jwt.parseJwt(token);
    return decodedToken['userId'];
  }

  void _initializeJoinedTeams(Hackathon hackathon) {
    joinedTeams.clear();
    for (var team in hackathon.teams) {
      for (var user in team.users!) {
        if (user.id == userId) {
          joinedTeams.add(team.id);
        }
      }
    }
  }

  void _updateTeamMembers(Hackathon hackathon, int teamId, bool isJoining) {
    setState(() {
      final team = hackathon.teams.firstWhere((t) => t.id == teamId);
      if (isJoining) {
        team.users!.add(User(id: userId, username: '', lastName: '', firstName: '', profilePicture: '', email: ''));
      } else {
        team.users!.removeWhere((user) => user.id == userId);
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Hackathon Details'),
      ),
      body: MultiBlocProvider(
        providers: [
          BlocProvider.value(
            value: context.read<HackathonBloc>(),
          ),
          BlocProvider(
            create: (context) => TeamBloc(),
          ),
        ],
        child: BlocListener<HackathonBloc, HackathonState>(
          listener: (context, hackathonState) {
            if (hackathonState is HackathonLoaded) {
              final hackathon = hackathonState.hackathons[0];
              _initializeJoinedTeams(hackathon);
              setState(() {}); // Met à jour l'UI après avoir initialisé joinedTeams
            }
          },
          child: BlocBuilder<HackathonBloc, HackathonState>(
            builder: (context, hackathonState) {
              if (hackathonState is HackathonLoading) {
                return const Center(child: CircularProgressIndicator());
              } else if (hackathonState is HackathonLoaded) {
                final hackathon = hackathonState.hackathons[0];
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
                                  const Text('Members:'),
                                  if (team.users != null)
                                    for (var user in team.users!) Text(user.username),
                                ],
                              ),
                              trailing: isJoined
                                  ? ElevatedButton(
                                onPressed: () {
                                  context.read<TeamBloc>().add(LeaveTeam(team.id, widget.token));
                                  _updateTeamMembers(hackathon, team.id, false);
                                },
                                child: const Text('Leave'),
                              )
                                  : ElevatedButton(
                                onPressed: () {
                                  context.read<TeamBloc>().add(JoinTeam(team.id, widget.token));
                                  _updateTeamMembers(hackathon, team.id, true);
                                },
                                child: const Text('Join'),
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
                          } else if (state is TeamLeft) {
                            ScaffoldMessenger.of(context).showSnackBar(
                              SnackBar(content: Text(state.message)),
                            );
                            setState(() {
                              joinedTeams.remove(state.teamId);
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
                if (kDebugMode) {
                  print('Error loading hackathon: ${hackathonState.message}');
                }
                return Center(child: Text('Error: ${hackathonState.message}'));
              } else {
                return const Center(child: Text('Something went wrong.'));
              }
            },
          ),
        ),
      ),
    );
  }
}
