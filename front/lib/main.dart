import 'dart:convert';
import 'dart:io';

import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:firebase_auth/firebase_auth.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_native_splash/flutter_native_splash.dart';
import 'package:front/screens/home_screen.dart';
import 'package:front/screens/login/login_screen.dart';
import 'package:front/services/logout/logout_bloc.dart';
import 'package:front/services/logout/logout_state.dart';
import 'package:front/utils/config.dart';
import 'package:front/utils/routes.dart';
import 'package:json_theme/json_theme.dart';
import 'firebase_options.dart';

Future<void> main() async {
  WidgetsBinding widgetsBinding = WidgetsFlutterBinding.ensureInitialized();

  HttpOverrides.global = MyHttpOverrides();

  try {
    if (!kIsWeb) {
      if (kDebugMode) {
        print(Platform.operatingSystem);
      }
    }
  } catch (e) {
    if (kDebugMode) {
      print('Platform not available: $e');
    }
  }

  // Splash persistance jusqu'à l'initialisation
  FlutterNativeSplash.preserve(widgetsBinding: widgetsBinding);

  // Gestion des tâches
  await Future.delayed(const Duration(seconds: 2));

  // Fin de l'écran splash
  FlutterNativeSplash.remove();

  await Firebase.initializeApp(options: DefaultFirebaseOptions.currentPlatform);

  // Chargement du thème depuis un fichier JSON
  final themeStr = await rootBundle.loadString('assets/theme.json');
  final theme = ThemeDecoder.decodeThemeData(jsonDecode(themeStr))!;

  // Configuration Firebase pour les émulateurs en mode debug
  if (kDebugMode) {
    Config().configureFirebaseEmulators();
  }

  // Évite d'avoir du cache
  await FirebaseFirestore.instance.terminate();
  await FirebaseFirestore.instance.clearPersistence();

  runApp(MyApp(theme: theme));
}

class MyHttpOverrides extends HttpOverrides {
  @override
  HttpClient createHttpClient(SecurityContext? context) {
    return super.createHttpClient(context)
      ..badCertificateCallback = (X509Certificate cert, String host, int port) => true;
  }
}

class MyApp extends StatelessWidget {
  final ThemeData theme;

  const MyApp({super.key, required this.theme});

  @override
  Widget build(BuildContext context) {
    return MultiBlocProvider(
      providers: Config.blocProviders,
      child: MaterialApp(
        debugShowCheckedModeBanner: false,
        title: "Kiwi Corporation",
        localizationsDelegates: Config.localizationsDelegates,
        initialRoute: '/',
        routes: getApplicationRoutes(),
        onUnknownRoute: unknownRoute,
        supportedLocales: const [
          Locale('en', ''),
          Locale('fr', ''),
        ],
        theme: theme,
        builder: (context, child) {
          return BlocListener<AuthenticationBloc, AuthenticationState>(
            listener: (context, state) {
              if (state is Unauthenticated) {
                Navigator.of(context).pushReplacementNamed('/login');
              }
            },
            child: child,
          );
        },
        home: _buildHomeScreen(),
      ),
    );
  }

  Widget _buildHomeScreen() {
    return StreamBuilder<User?>(
      stream: FirebaseAuth.instance.authStateChanges(),
      builder: (context, snapshot) {
        if (snapshot.connectionState == ConnectionState.active) {
          if (snapshot.hasData) {
            return HomeScreen(token: snapshot.data!.uid);
          } else {
            return const LoginPage();
          }
        }
        return const CircularProgressIndicator();
      },
    );
  }
}
