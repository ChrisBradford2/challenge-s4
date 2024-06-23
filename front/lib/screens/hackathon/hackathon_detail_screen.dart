import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:front/utils/config.dart';
import 'package:jwt_decode/jwt_decode.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

import '../../models/hackathon_model.dart';
import '../../models/team_model.dart';
import '../../services/hackathons/hackathon_bloc.dart';
import '../../services/hackathons/hackathon_event.dart';
import '../../services/hackathons/hackathon_state.dart';
import '../../services/team/team_bloc.dart';
import '../../services/team/team_event.dart';
import '../../services/team/team_state.dart';
import '../../models/user_model.dart';
import '../team/team_manage_screen.dart';

class HackathonDetailPage extends StatefulWidget {
  final String id;
  final String token;

  const HackathonDetailPage({super.key, required this.id, required this.token});

  @override
  HackathonDetailPageState createState() => HackathonDetailPageState();
}

class HackathonDetailPageState extends State<HackathonDetailPage> {
  final Set<int> joinedTeams = <int>{};
  late int userId;
  User? currentUser; // Changed to nullable
  Hackathon? currentHackathon;
  late Future<void> fetchCurrentUserFuture;

  @override
  void initState() {
    super.initState();
    userId = _getUserIdFromToken(widget.token);
    fetchCurrentUserFuture = _fetchCurrentUser();
    context.read<HackathonBloc>().add(FetchSingleHackathons(widget.token, widget.id));
  }

  int _getUserIdFromToken(String token) {
    Map<String, dynamic> decodedToken = Jwt.parseJwt(token);
    return decodedToken['userId'];
  }

  Future<void> _fetchCurrentUser() async {
    final url = '${Config.baseUrl}/user/me'; // Remplacez YOUR_API_URL par l'URL de votre API
    final response = await http.get(
      Uri.parse(url),
      headers: {'Authorization': 'Bearer ${widget.token}'},
    );

    if (kDebugMode) {
      print('Fetching current user from $url');
      print('Response status: ${response.statusCode}');
      print('Response body: ${response.body}');
    }

    if (response.statusCode == 200) {
      setState(() {
        currentUser = User.fromJson(jsonDecode(response.body));
      });
    } else {
      throw Exception('Failed to load current user');
    }
  }

  void _initializeJoinedTeams(Hackathon hackathon) {
    setState(() {
      joinedTeams.clear();
      for (var team in hackathon.teams) {
        for (var user in team.users!) {
          if (user.id == userId) {
            joinedTeams.add(team.id);
          }
        }
      }
    });
  }

  void _updateTeamMembers(int teamId, bool isJoining) {
    if (currentUser == null) return; // Ensure currentUser is initialized
    setState(() {
      final team = currentHackathon!.teams.firstWhere((t) => t.id == teamId);
      if (isJoining) {
        team.users!.add(currentUser!);
        joinedTeams.add(teamId);
      } else {
        team.users!.removeWhere((u) => u.id == userId);
        joinedTeams.remove(teamId);
      }
    });
  }

  void _navigateToManagePage(BuildContext context, Team team) async {
    final result = await Navigator.push(
      context,
      MaterialPageRoute(
        builder: (context) => BlocProvider(
          create: (context) => TeamBloc(),
          child: TeamManagePage(team: team, token: widget.token),
        ),
      ),
    );

    if (result == 'left') {
      context.read<HackathonBloc>().add(FetchSingleHackathons(widget.token, widget.id));
    }
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
        child: FutureBuilder<void>(
          future: fetchCurrentUserFuture,
          builder: (context, snapshot) {
            if (snapshot.connectionState == ConnectionState.waiting) {
              return const Center(child: CircularProgressIndicator());
            } else if (snapshot.hasError) {
              return Center(child: Text('Error: ${snapshot.error}'));
            } else {
              return MultiBlocListener(
                listeners: [
                  BlocListener<HackathonBloc, HackathonState>(
                    listener: (context, hackathonState) {
                      if (hackathonState is HackathonLoaded) {
                        final hackathon = hackathonState.hackathons[0];
                        currentHackathon = hackathon;
                        _initializeJoinedTeams(hackathon);
                        setState(() {});
                      }
                    },
                  ),
                  BlocListener<TeamBloc, TeamState>(
                    listener: (context, state) {
                      if (state is TeamJoined) {
                        ScaffoldMessenger.of(context).showSnackBar(
                          SnackBar(content: Text(state.message)),
                        );
                        _updateTeamMembers(state.teamId, true);
                      } else if (state is TeamLeft) {
                        ScaffoldMessenger.of(context).showSnackBar(
                          SnackBar(content: Text(state.message)),
                        );
                        _updateTeamMembers(state.teamId, false);
                      } else if (state is TeamError) {
                        ScaffoldMessenger.of(context).showSnackBar(
                          SnackBar(content: Text(state.error)),
                        );
                      }
                    },
                  ),
                ],
                child: BlocBuilder<HackathonBloc, HackathonState>(
                  builder: (context, hackathonState) {
                    if (hackathonState is HackathonLoading) {
                      return const Center(child: CircularProgressIndicator());
                    } else if (hackathonState is HackathonLoaded) {
                      final hackathon = hackathonState.hackathons[0];
                      currentHackathon = hackathon;

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
                                      onPressed: () => _navigateToManagePage(context, team),
                                      child: const Text('Manage'),
                                    )
                                        : ElevatedButton(
                                      onPressed: currentUser == null
                                          ? null
                                          : () {
                                        context.read<TeamBloc>().add(JoinTeam(team.id, widget.token));
                                      },
                                      child: const Text('Join'),
                                    ),
                                  );
                                },
                              ),
                            ),
                          ],
                        ),
                      );
                    } else if (hackathonState is HackathonError) {
                      return Center(child: Text('Error: ${hackathonState.message}'));
                    } else {
                      return const Center(child: Text('Something went wrong.'));
                    }
                  },
                ),
              );
            }
          },
        ),
      ),
    );
  }
}
