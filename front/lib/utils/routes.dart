import 'package:flutter/cupertino.dart';

import '../screens/login/login_screen.dart';
import '../screens/register/register_screen.dart';

Map<String, WidgetBuilder> getApplicationRoutes() {
  return {
    '/': (BuildContext context) => const LoginPage(),
    '/register': (BuildContext context) => const RegisterPage(),
  };
}
