import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import '../../services/hackathons/hackathon_bloc.dart';
import '../../services/hackathons/hackathon_event.dart';
import '../../services/hackathons/hackathon_state.dart';

class HackathonDetailPage extends StatelessWidget {
  final String id;
  final String token;

  HackathonDetailPage({required this.id, required this.token});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Hackathon Details'),
      ),
      body: BlocProvider(
        create: (context) => HackathonBloc()..add(FetchSingleHackathons(token, id)),
        child: BlocBuilder<HackathonBloc, HackathonState>(
          builder: (context, state) {
            if (state is HackathonLoading) {
              return Center(child: CircularProgressIndicator());
            } else if (state is HackathonLoaded) {
              final hackathon = state.hackathons[0];
              return Padding(
                padding: const EdgeInsets.all(16.0),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text('Name: ${hackathon["name"]}', style: TextStyle(fontSize: 24)),
                    SizedBox(height: 16),
                    Text('Date: ${hackathon["date"]}', style: TextStyle(fontSize: 18)),
                    SizedBox(height: 16),
                    Text('Location: ${hackathon["location"]}', style: TextStyle(fontSize: 18)),
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
