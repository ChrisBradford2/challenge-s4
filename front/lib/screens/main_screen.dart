import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:front/services/hackathons/hackathon_bloc.dart';
import 'package:front/screens/profile/profile_screen.dart';
import '../services/hackathons/hackathon_event.dart';
import 'hackathon/hackathon_screen.dart';
import 'home_screen.dart';

class MainScreen extends StatefulWidget {
  final String token;

  const MainScreen({super.key, required this.token});

  @override
  MainScreenState createState() => MainScreenState();
}

class MainScreenState extends State<MainScreen> {
  int _selectedIndex = 0;
  late List<Widget> _widgetOptions;

  @override
  void initState() {
    super.initState();
    _widgetOptions = [
      HomeScreen(token: widget.token),
      HackathonScreen(token: widget.token),
      ProfileScreen(token: widget.token),
    ];
  }

  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });

    // Rafraîchit les données lorsque l'utilisateur navigue vers l'écran d'accueil
    if (index == 0) {
      context.read<HackathonBloc>().add(FetchHackathons(widget.token));
    }
  }

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (context) => HackathonBloc(),
      child: Scaffold(
        body: IndexedStack(
          index: _selectedIndex,
          children: _widgetOptions,
        ),
        bottomNavigationBar: BottomNavigationBar(
          items: const <BottomNavigationBarItem>[
            BottomNavigationBarItem(
              icon: Icon(Icons.home),
              label: 'Home',
            ),
            BottomNavigationBarItem(
              icon: Icon(Icons.event),
              label: 'Mes hackathons',
            ),
            BottomNavigationBarItem(
              icon: Icon(Icons.person),
              label: 'Profil',
            ),
          ],
          currentIndex: _selectedIndex,
          selectedItemColor: Colors.amber[800],
          onTap: _onItemTapped,
        ),
      ),
    );
  }
}
