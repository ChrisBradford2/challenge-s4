import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:http/http.dart' as http;

import '../services/logout/logout_bloc.dart';
import '../services/logout/logout_event.dart';
import '../services/logout/logout_state.dart';
import '../services/user_service.dart';
import '../utils/config.dart';

class HomeScreen extends StatelessWidget {
  final String token;

  const HomeScreen({super.key, required this.token});

  Future<List<Map<String, String>>> fetchHackathons() async {
    final response = await http.get(
      Uri.parse('${Config.baseUrl}/hackathons'),
      headers: {'Authorization': 'Bearer $token'},
    );

    if (response.statusCode == 200) {
      List<dynamic> data = jsonDecode(response.body)['data'];
      return data.map((item) {
        return {
          "name": item["name"].toString(),
          "date": item["start_date"].toString(),
          "location": item["location"].toString(),
        };
      }).toList();
    } else {
      throw Exception('Failed to load hackathons');
    }
  }

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
        body: FutureBuilder<List<Map<String, String>>>(
          future: fetchHackathons(),
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
                    child: FutureBuilder<String>(
                      future: userService.fetchFirstName(token),
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
                      itemCount: snapshot.data!.length,
                      itemBuilder: (context, index) {
                        var hackathon = snapshot.data![index];
                        return ListTile(
                          title: Text(hackathon["name"]!),
                          subtitle: Text("${hackathon["date"]} - ${hackathon["location"]}"),
                          leading: const Icon(Icons.event),
                        );
                      },
                    ),
                  ),
                  /*
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
                  ),*/
                  FloatingActionButton(
                    onPressed: () {
                      Navigator.of(context).pushNamed('/profile');
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
