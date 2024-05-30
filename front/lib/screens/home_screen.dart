import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../services/hackathons/hackathon_bloc.dart';
import '../services/hackathons/hackathon_event.dart';
import '../services/hackathons/hackathon_state.dart';
import 'hackathon/hackathon_detail_screen.dart';

class HomeScreen extends StatefulWidget {
  final String token;

  HomeScreen({super.key, required this.token});

  @override
  _HomeScreenState createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  @override
  void initState() {
    super.initState();
    // Fetch hackathons when the screen is initialized
    context.read<HackathonBloc>().add(FetchHackathons(widget.token));
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Accueil"),
      ),
      body: BlocListener<HackathonBloc, HackathonState>(
        listener: (context, state) {
          if (state is HackathonAdded) {
            // Fetch the updated list of hackathons
            context.read<HackathonBloc>().add(FetchHackathons(widget.token));
          }
        },
        child: BlocBuilder<HackathonBloc, HackathonState>(
          builder: (context, state) {
            if (state is HackathonLoading) {
              return const Center(child: CircularProgressIndicator());
            } else if (state is HackathonLoaded) {
              return RefreshIndicator(
                onRefresh: () async {
                  context.read<HackathonBloc>().add(FetchHackathons(widget.token));
                },
                child: ListView.builder(
                  itemCount: state.hackathons.length,
                  itemBuilder: (context, index) {
                    final hackathon = state.hackathons[index];
                    final id = hackathon['id']!;

                    return ListTile(
                      leading: const Icon(Icons.event),
                      title: Text(hackathon['name']!),
                      subtitle: Row(
                        children: [
                          const SizedBox(width: 8),
                          Text(hackathon['date']!),
                        ],
                      ),
                      onTap: () {
                        Navigator.push(
                          context,
                          MaterialPageRoute(
                            builder: (context) => HackathonDetailPage(
                              id: id,
                              token: widget.token,
                            ),
                          ),
                        );
                      },
                    );
                  },
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
