import 'package:flutter/material.dart';

import '../screens/login/login_screen.dart';
import '../screens/profile/profile_screen.dart';
import '../screens/register/register_screen.dart';

Map<String, WidgetBuilder> getApplicationRoutes() {
  return {
    '/login': (BuildContext context) => const LoginPage(),
    '/register': (BuildContext context) => const RegisterPage(),
    '/profile': (BuildContext context) => ProfileScreen(),
  };
}

Route<dynamic> unknownRoute(RouteSettings settings) {
  return MaterialPageRoute(
    builder: (BuildContext context) => const Scaffold(
      body: Center(
        child: Text('Page not found :('),
      ),
    ),
  );
}
