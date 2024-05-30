import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:image_picker/image_picker.dart';
import 'dart:io';

import '../../services/register/register_bloc.dart';
import '../../services/register/register_event.dart';
import '../../services/register/register_state.dart';
import '../../widgets/buttons/button_base.dart';

class RegisterPage extends StatefulWidget {
  const RegisterPage({super.key});

  @override
  RegisterScreenState createState() => RegisterScreenState();
}

class RegisterScreenState extends State<RegisterPage> {
  final GlobalKey<FormState> _formKey = GlobalKey<FormState>();
  final TextEditingController _usernameController = TextEditingController();
  final TextEditingController _lastNameController = TextEditingController();
  final TextEditingController _firstNameController = TextEditingController();
  final TextEditingController _emailController = TextEditingController();
  final TextEditingController _passwordController = TextEditingController();
  File? _profileImage;
  final ImagePicker _picker = ImagePicker();

  get pickedWhat => null;

  Future<void> _pickImage() async {
    final XFile? pickedFile = await _picker.pickImage(source: ImageSource.gallery);
    if (pickedFile != null) {
      setState(() {
        _profileImage = File(pickedFile.path);
      });
    }
  }

  void _onRegisterButtonPressed() {
    if (_formKey.currentState!.validate()) {
      BlocProvider.of<RegistrationBloc>(context).add(
        SignUpButtonPressed(
          username: _usernameController.text,
          lastName: _lastNameController.text,
          firstName: _firstNameController.text,
          email: _emailController.text,
          password: _passwordController.text,
          profilePicturePath: _profileImage?.path ?? '',
        ),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
          title: Padding(
            padding: const EdgeInsets.all(90.0),
            child: Image.asset(
              'assets/logo.png', // Remplacez par le chemin de votre logo
              height: 100.0, // Vous pouvez ajuster la taille selon vos besoins
            ),
          ),
      ),
      body: BlocListener<RegistrationBloc, RegistrationState>(
        listener: (context, state) {
          if (state is RegistrationSuccess) {
            ScaffoldMessenger.of(context)
              ..removeCurrentSnackBar()
              ..showSnackBar(const SnackBar(content: Text('Inscription réussie! Vous pouvez vous connecter maintenant.')));
          }
          if (state is RegistrationFailure) {
            ScaffoldMessenger.of(context)
              ..removeCurrentSnackBar()
              ..showSnackBar(SnackBar(content: Text('Erreur d\'inscription: ${state.error}')));
          }
        },
        child: Form(
          key: _formKey,
          child: SingleChildScrollView(
            padding: const EdgeInsets.all(20),
            child: Column(
              children: [
                const SizedBox(height: 30),
                const Text(
                  'Bienvenue sur Kiwi Collective',
                  style: TextStyle(
                    fontSize: 20.0,
                    fontWeight: FontWeight.w500,
                  ),
                ),
                const SizedBox(height: 10),
                const Center(
                  child: Text(
                    'Créer un compte pour pour commencer à participer à un hackathon',
                    style: TextStyle(
                      fontSize: 28.0,
                      fontWeight: FontWeight.bold,
                    ),
                    textAlign: TextAlign.center,
                  ),
                ),
                const SizedBox(height: 30),
                TextFormField(
                  controller: _usernameController,
                  decoration: const InputDecoration(labelText: 'Nom d\'utilisateur'),
                  validator: (value) {
                    if (value!.isEmpty) {
                      return 'Veuillez entrer un nom d\'utilisateur';
                    }
                    return null;
                  },
                ),
                TextFormField(
                  controller: _lastNameController,
                  decoration: const InputDecoration(labelText: 'Nom'),
                  validator: (value) {
                    if (value!.isEmpty) {
                      return 'Veuillez entrer un nom';
                    }
                    return null;
                  },
                ),
                TextFormField(
                  controller: _firstNameController,
                  decoration: const InputDecoration(labelText: 'Prénom'),
                  validator: (value) {
                    if (value!.isEmpty) {
                      return 'Veuillez entrer un prénom';
                    }
                    return null;
                  },
                ),
                TextFormField(
                  controller: _emailController,
                  decoration: const InputDecoration(labelText: 'Email'),
                  validator: (value) {
                    if (value!.isEmpty) {
                      return 'Veuillez entrer un email';
                    }
                    return null;
                  },
                ),
                TextFormField(
                  controller: _passwordController,
                  decoration: const InputDecoration(labelText: 'Mot de passe'),
                  obscureText: true,
                  validator: (value) {
                    if (value!.isEmpty) {
                      return 'Veuillez entrer un mot de passe';
                    }
                    return null;
                  },
                ),
                const SizedBox(height: 10),
                if (_profileImage != null)
                  SizedBox(
                    width: 100,
                    height: 100,
                    child: Image.file(_profileImage!, fit: BoxFit.cover),
                  ),
                Center(
                  child: TextButton(
                    onPressed: _pickImage,
                    child: const Row(
                      mainAxisSize: MainAxisSize.min, // Pour que la ligne ne prenne que la largeur nécessaire
                      children: [
                        Icon(Icons.file_upload_outlined, color: Colors.green), // Icône à gauche du texte
                        SizedBox(width: 8), // Espace entre l'icône et le texte
                        Text(
                          'Choisir une photo de profil',
                          style: TextStyle(color: Colors.green),
                        ),
                      ],
                    ),
                  ),
                ),
                /*ElevatedButton(
                  onPressed: _pickImage,
                  child: const Text('Choisir une image de profil'),
                ),*/
                const SizedBox(height: 30),
                SizedBox(
                  width: MediaQuery.of(context).size.width,
                  child: ButtonBase(
                    text: 'S\'inscrire',
                    onPressed: _onRegisterButtonPressed,
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
