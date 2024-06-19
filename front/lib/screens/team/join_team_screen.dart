import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import '../../services/team/team_bloc.dart';
import '../../services/team/team_event.dart';
import '../../services/team/team_state.dart';

class JoinTeamPage extends StatelessWidget {
  final String teamId;
  final String token;

  const JoinTeamPage({super.key, required this.teamId, required this.token});

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (context) => TeamBloc(),
      child: Scaffold(
        appBar: AppBar(
          title: const Text('Join Team'),
        ),
        body: Center(
          child: BlocConsumer<TeamBloc, TeamState>(
            listener: (context, state) {
              if (state is TeamJoined) {
                ScaffoldMessenger.of(context).showSnackBar(
                  SnackBar(content: Text(state.message)),
                );
                Navigator.pop(context);
              } else if (state is TeamError) {
                ScaffoldMessenger.of(context).showSnackBar(
                  SnackBar(content: Text(state.error)),
                );
              }
            },
            builder: (context, state) {
              if (state is TeamLoading) {
                return const CircularProgressIndicator();
              }
              return ElevatedButton(
                onPressed: () {
                  context.read<TeamBloc>().add(JoinTeam(teamId, token));
                },
                child: const Text('Join this Team'),
              );
            },
          ),
        ),
      ),
    );
  }
}
