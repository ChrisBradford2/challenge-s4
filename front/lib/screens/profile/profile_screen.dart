import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../services/logout/logout_bloc.dart';
import '../../services/logout/logout_event.dart';
import '../../services/user/user_bloc.dart';
import '../../services/user/user_event.dart';
import '../../services/user/user_state.dart';

class ProfileScreen extends StatelessWidget {
  final String token;

  const ProfileScreen({super.key, required this.token});

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (context) => UserBloc()..add(FetchUser(token)),
      child: Scaffold(
        appBar: AppBar(
          title: const Text("Profil"),
        ),
        body: BlocBuilder<UserBloc, UserState>(
          builder: (context, state) {
            if (state is UserLoading) {
              return const Center(child: CircularProgressIndicator());
            } else if (state is UserLoaded) {
              return SingleChildScrollView(
                child: Center(
                  child: Padding(
                    padding: const EdgeInsets.all(16.0),
                    child: Column(
                      mainAxisAlignment: MainAxisAlignment.start,
                      crossAxisAlignment: CrossAxisAlignment.center,
                      children: [
                        SizedBox(height: 20),
                        CircleAvatar(
                          radius: 50,
                          backgroundImage: NetworkImage(state.user.profilePicture),
                        ),
                        SizedBox(height: 20),
                        Text(
                          state.user.firstName,
                          style: Theme.of(context).textTheme.headlineSmall,
                        ),
                        Text(
                          state.user.lastName,
                          style: Theme.of(context).textTheme.headlineSmall,
                        ),
                        Text(
                          state.user.email,
                          style: Theme.of(context).textTheme.bodyLarge,
                        ),
                        SizedBox(height: 20),
                        ElevatedButton(
                          onPressed: () {
                            BlocProvider.of<AuthenticationBloc>(context).add(LogoutEvent());
                          },
                          child: const Text("Déconnexion"),
                        ),
                      ],
                    ),
                  ),
                ),
              );
            } else if (state is UserError) {
              return Center(child: Text("Error: ${state.message}"));
            } else {
              return const Center(child: Text("Une erreur inconnue s'est produite."));
            }
          },
        ),
      ),
    );
  }
}
