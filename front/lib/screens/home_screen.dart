import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../services/logout/logout_bloc.dart';
import '../services/logout/logout_event.dart';
import '../services/logout/logout_state.dart';
import '../services/user_service.dart';

class HomeScreen extends StatelessWidget {
  final String token;

  HomeScreen({super.key, required this.token});

  // @todo Delete this when the backend is ready
  final List<Map<String, String>> hackathons = [
    {
      "name": "Hackathon Innovate",
      "date": "2024-06-15",
      "location": "Paris, France"
    },
    {
      "name": "Code Challenge",
      "date": "2024-07-20",
      "location": "San Francisco, CA"
    },
    {
      "name": "Developers Fest",
      "date": "2024-08-12",
      "location": "Berlin, Germany"
    },
    {
      "name": "Tech Creators",
      "date": "2024-09-05",
      "location": "Tokyo, Japan"
    },
    {
      "name": "Global DevCon",
      "date": "2024-10-10",
      "location": "New York, NY"
    }
  ];

  @override
  Widget build(BuildContext context) {
    final userService = UserService();

    return BlocListener<AuthenticationBloc, AuthenticationState>(
      listener: (context, state) {
        if (state is Unauthenticated) {
          Navigator.of(context).pushReplacementNamed('/login');
        }
      },
      child: Scaffold(
        appBar: AppBar(
          title: const Text("Accueil"),
        ),
        body: FutureBuilder<String>(
          future: userService.fetchFirstName(token),
          builder: (context, snapshot) {
            if (snapshot.connectionState == ConnectionState.done) {
              if (snapshot.hasError) {
                return Center(child: Text("Error: ${snapshot.error}"));
              }
              return Column(
                mainAxisAlignment: MainAxisAlignment.start,
                children: [
                  Padding(
                    padding: const EdgeInsets.all(16.0),
                    child: Text(
                      "Salut ${snapshot.data}, content de te revoir !",
                      style: Theme.of(context).textTheme.headlineSmall,
                    ),
                  ),
                  Expanded(
                    child: ListView.builder(
                      itemCount: hackathons.length,
                      itemBuilder: (context, index) {
                        var hackathon = hackathons[index];
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
                      showDialog(
                        context: context,
                        builder: (context) {
                          return AlertDialog(
                            title: const Text("Déconnexion"),
                            content: const Text("Voulez-vous vraiment vous déconnecter ?"),
                            actions: [
                              TextButton(
                                onPressed: () {
                                  Navigator.of(context).pop();
                                },
                                child: const Text("Annuler"),
                              ),
                              TextButton(
                                onPressed: () {
                                  BlocProvider.of<AuthenticationBloc>(context).add(LogoutEvent());
                                  Navigator.of(context).pop();
                                },
                                child: const Text("Déconnexion"),
                              ),
                            ],
                          );
                        },
                      );
                    },
                    child: const Icon(Icons.person),
                  ),
                ],
              );
            } else {
              return const Center(child: CircularProgressIndicator());
            }
          },
        ),
      ),
    );
  }
}
