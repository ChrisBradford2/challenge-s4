import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:front/utils/roles.dart';
import '../../services/login/login_bloc.dart';
import '../../services/login/login_event.dart';
import '../../services/login/login_state.dart';
import '../../utils/config.dart';  // Importez les rôles
import 'admin_screen.dart';

class AdminLoginPage extends StatelessWidget {
  const AdminLoginPage({super.key});

  @override
  Widget build(BuildContext context) {
    final loginBloc = BlocProvider.of<LoginBloc>(context);

    TextEditingController emailController = TextEditingController();
    TextEditingController passwordController = TextEditingController();

    return Scaffold(
      appBar: AppBar(title: const Text("Admin Login")),
      body: BlocListener<LoginBloc, LoginState>(
        listener: (context, state) async {
          if (state is LoginSuccess) {
            final user = await loginBloc.getUserDetails(state.token);
            if (user != null && (user.roles & Roles.roleAdmin) != 0) {
              Navigator.of(context).pushReplacement(
                MaterialPageRoute(
                  builder: (context) => AdminScreen(),
                ),
              );
            } else {
              ScaffoldMessenger.of(context).showSnackBar(
                const SnackBar(content: Text('Access denied. Admins only.')),
              );
            }
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
                child: const Text('Caca'),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
