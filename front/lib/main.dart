import 'dart:convert';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:flutter_native_splash/flutter_native_splash.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:firebase_auth/firebase_auth.dart';
import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:front/screens/admin/admin_login_screen.dart';
import 'package:front/screens/app.dart';
import 'package:front/screens/login/login_screen.dart';
import 'package:front/screens/profile/profile_screen.dart';
import 'package:front/services/logout/logout_bloc.dart';
import 'package:front/services/logout/logout_state.dart';
import 'package:front/utils/config.dart';
import 'package:front/utils/routes.dart';
import 'package:json_theme/json_theme.dart';

import 'firebase_options.dart';
import 'platform_overrides.dart';

Future<void> main() async {
  WidgetsBinding widgetsBinding = WidgetsFlutterBinding.ensureInitialized();

  // Setup platform-specific configurations
  setupPlatformSpecific();

  await dotenv.load(fileName: ".env");

  // Splash persistence until initialization
  FlutterNativeSplash.preserve(widgetsBinding: widgetsBinding);

  // Handle initialization tasks
  await Future.delayed(const Duration(seconds: 2));

  // End the splash screen
  FlutterNativeSplash.remove();

  await Firebase.initializeApp(options: DefaultFirebaseOptions.currentPlatform);

  // Load theme from JSON file
  final themeStr = await rootBundle.loadString('assets/theme.json');
  final theme = ThemeDecoder.decodeThemeData(jsonDecode(themeStr))!;

  // Configure Firebase for emulators in debug mode (only for mobile and desktop platforms)
  if (kDebugMode && !kIsWeb) {
    if (defaultTargetPlatform == TargetPlatform.android ||
        defaultTargetPlatform == TargetPlatform.iOS ||
        defaultTargetPlatform == TargetPlatform.macOS ||
        defaultTargetPlatform == TargetPlatform.windows) {
      Config().configureFirebaseEmulators();
    }
  }

  // Avoid cache
  await FirebaseFirestore.instance.terminate();
  await FirebaseFirestore.instance.clearPersistence();

  runApp(MyApp(theme: theme));
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
        onGenerateRoute: (settings) {
          if (settings.name == '/profile') {
            final String token = settings.arguments as String;
            return MaterialPageRoute(
              builder: (context) => ProfileScreen(token: token),
            );
          }
          return unknownRoute(settings);
        },
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
            return MainScreen(token: snapshot.data!.uid); // Use MainScreen with the token
          } else {
            // Redirect to AdminLoginPage if on web, else to LoginPage
            return kIsWeb ? const AdminLoginPage() : const LoginPage();
          }
        }
        return const CircularProgressIndicator();
      },
    );
  }
}
