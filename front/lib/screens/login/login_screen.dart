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
      appBar: AppBar(title: const Text("Login")),
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
            mainAxisAlignment: MainAxisAlignment.center,
            children: <Widget>[
              TextField(
                controller: emailController,
                decoration: const InputDecoration(labelText: 'Email'),
              ),
              TextField(
                controller: passwordController,
                decoration: const InputDecoration(labelText: 'Password'),
                obscureText: true,
              ),
              const SizedBox(height: 20),
              ElevatedButton(
                onPressed: () {
                  loginBloc.add(LoginButtonPressed(
                      email: emailController.text,
                      password: passwordController.text
                  ));
                },
                child: const Text('Login'),
              ),
              SizedBox(height: 20),
              SizedBox(
                width: MediaQuery.of(context).size.width,
                child: ButtonBase(
                  text: 'Login',
                  onPressed: () {
                    loginBloc.add(LoginButtonPressed(
                      email: emailController.text,
                      password: passwordController.text,
                    ));
                  },
                ),
              ),
              TextButton(
                onPressed: () {
                  Navigator.of(context).pushNamed('/register');
                },
                child: const Text('Register'),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
