import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import '../services/hackathons/hackathon_bloc.dart';
import '../services/hackathons/hackathon_event.dart';
import '../services/hackathons/hackathon_state.dart';
import '../services/logout/logout_bloc.dart';
import '../services/logout/logout_state.dart';
import '../services/user_service.dart';

class HomeScreen extends StatelessWidget {
  final String token;

  HomeScreen({super.key, required this.token});

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (context) => HackathonBloc()..add(FetchHackathons(token)),
      child: BlocListener<AuthenticationBloc, AuthenticationState>(
        listener: (context, state) {
          if (state is Unauthenticated) {
            Navigator.of(context).pushReplacementNamed('/login');
          }
        },
        child: Scaffold(
          appBar: AppBar(
            title: const Text("Accueil"),
          ),
          body: BlocBuilder<HackathonBloc, HackathonState>(
            builder: (context, state) {
              if (state is HackathonLoading) {
                return const Center(child: CircularProgressIndicator());
              } else if (state is HackathonLoaded) {
                if (state.hackathons.isEmpty) {
                  return const Center(child: Text("Aucun hackathon disponible."));
                }
                return Column(
                  mainAxisAlignment: MainAxisAlignment.start,
                  children: [
                    Padding(
                      padding: const EdgeInsets.all(16.0),
                      child: FutureBuilder<String>(
                        future: UserService().fetchFirstName(token),
                        builder: (context, userSnapshot) {
                          if (userSnapshot.connectionState == ConnectionState.done) {
                            if (userSnapshot.hasError) {
                              return Center(child: Text("Error: ${userSnapshot.error}"));
                            }
                            return Text(
                              "Salut ${userSnapshot.data}, content de te revoir !",
                              style: Theme.of(context).textTheme.headlineSmall,
                            );
                          } else {
                            return const CircularProgressIndicator();
                          }
                        },
                      ),
                    ),
                    Expanded(
                      child: ListView.builder(
                        itemCount: state.hackathons.length,
                        itemBuilder: (context, index) {
                          var hackathon = state.hackathons[index];
                          return ListTile(
                            title: Text(hackathon["name"]!),
                            subtitle: Text("${hackathon["date"]} - ${hackathon["location"]}"),
                            leading: const Icon(Icons.event),
                          );
                        },
                      ),
                    ),
                    FloatingActionButton(
                      onPressed: () {
                        Navigator.of(context).pushNamed('/profile', arguments: token);
                      },
                      child: const Icon(Icons.person),
                    ),
                  ],
                );
              } else if (state is HackathonError) {
                return Center(child: Text("Error: ${state.message}"));
              } else {
                return const Center(child: Text("Une erreur inconnue s'est produite."));
              }
            },
          ),
        ),
      ),
    );
  }
}