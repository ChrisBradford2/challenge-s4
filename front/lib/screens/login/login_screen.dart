import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:front/widgets/buttons/button_base.dart';

import '../../services/login/login_bloc.dart';
import '../../services/login/login_event.dart';
import '../../services/login/login_state.dart';
import '../home_screen.dart';

class LoginPage extends StatelessWidget {
  const LoginPage({super.key});

  @override
  Widget build(BuildContext context) {
    final loginBloc = BlocProvider.of<LoginBloc>(context);

    TextEditingController emailController = TextEditingController();
    TextEditingController passwordController = TextEditingController();

    return Scaffold(
      appBar: AppBar(
          title:  Center(
            child: Image.asset(
              'assets/logo.png', // Remplacez par le chemin de votre logo
              height: 100.0, // Vous pouvez ajuster la taille selon vos besoins
            ),
          ),
      ),
      body: BlocListener<LoginBloc, LoginState>(
        listener: (context, state) {
          if (state is LoginSuccess) {
            Navigator.of(context).pushReplacement(
                MaterialPageRoute(
                    builder: (context) => HomeScreen(token: state.token)
                )
            );
          } else if (state is LoginFailure) {
            ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text(state.error)));
          }
        },
        child: Padding(
          padding: const EdgeInsets.all(20.0),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.start,
            children: <Widget>[
              const SizedBox(height: 30),
              const Text(
                'Kiwi collective',
                style: TextStyle(
                  fontSize: 20.0,
                  fontWeight: FontWeight.w500,
                ),
              ),
              const SizedBox(height: 80),
              TextField(
                controller: emailController,
                decoration: const InputDecoration(labelText: 'Email'),
              ),
              TextField(
                controller: passwordController,
                decoration: const InputDecoration(labelText: 'Mot de passe'),
                obscureText: true,
              ),
              const SizedBox(height: 30),
              /*ElevatedButton(
                onPressed: () {
                  loginBloc.add(LoginButtonPressed(
                      email: emailController.text,
                      password: passwordController.text
                  ));
                },
                child: const Text('Login'),
              ),*/
              //SizedBox(height: 20),
              TextButton(
                onPressed: () {
                  Navigator.of(context).pushNamed('/register');
                },
                child: const Text('S\'inscrire'),
              ),
              SizedBox(
                width: MediaQuery.of(context).size.width,
                child: ButtonBase(
                  text: 'Se connecter',
                  onPressed: () {
                    loginBloc.add(LoginButtonPressed(
                      email: emailController.text,
                      password: passwordController.text,
                    ));
                  },
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
