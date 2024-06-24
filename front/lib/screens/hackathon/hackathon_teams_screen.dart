import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:front/repositories/team_repository.dart';
import '../../models/team_model.dart';
import '../../models/hackathon_model.dart';
import '../../services/team/team_bloc.dart';
import '../../services/team/team_event.dart';
import '../../services/team/team_state.dart';
import '../team/team_detail_screen.dart';

class HackathonTeamsPage extends StatelessWidget {
  final Hackathon hackathon;
  final String token;

  const HackathonTeamsPage({Key? key, required this.hackathon, required this.token}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (context) => TeamBloc(TeamRepository())..add(FetchTeamsForHackathon(hackathon.id, token)),
      child: Scaffold(
        appBar: AppBar(
          title: Text('Teams for ${hackathon.name}'),
        ),
        body: BlocBuilder<TeamBloc, TeamState>(
          builder: (context, state) {
            if (state is TeamLoading) {
              return const Center(child: CircularProgressIndicator());
            } else if (state is TeamsLoaded) {
              if (state.teams.isEmpty) {
                return const Center(child: Text('No teams found.'));
              }
              return ListView.builder(
                itemCount: state.teams.length,
                itemBuilder: (context, index) {
                  final team = state.teams[index];
                  return ListTile(
                    title: Text(team.name),
                    subtitle: Text('Members: ${team.users?.length ?? 0}'),
                    onTap: () {
                      Navigator.push(
                        context,
                        MaterialPageRoute(
                          builder: (context) => TeamDetailPage(
                            team: team,
                            token: token,
                          ),
                        ),
                      );
                    },
                  );
                },
              );
            } else if (state is TeamError) {
              return Center(child: Text('Error: ${state.error}'));
            } else {
              return const Center(child: Text('No teams found.'));
            }
          },
        ),
      ),
    );
  }
}
