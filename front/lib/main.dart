import 'dart:io';

import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:cloud_functions/cloud_functions.dart';
import 'package:firebase_auth/firebase_auth.dart';
import 'package:firebase_storage/firebase_storage.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:flutter_native_splash/flutter_native_splash.dart';
import 'package:front/screens/home_screen.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:front/screens/login/login_screen.dart';
import 'package:front/services/authentication_service.dart';
import 'package:front/services/login/login_bloc.dart';
import 'package:front/services/register/register_bloc.dart';
import 'package:front/services/service_factory.dart';
import 'package:front/utils/routes.dart';
import 'package:json_theme/json_theme.dart';
import 'package:flutter/services.dart';
import 'dart:convert';
import 'package:firebase_core/firebase_core.dart';
import 'firebase_options.dart';


Future<void> main() async {
  WidgetsBinding widgetsBinding = WidgetsFlutterBinding.ensureInitialized();
  // Splash persistance jusqu'a l'ignitialisation
  FlutterNativeSplash.preserve(widgetsBinding: widgetsBinding);
  // gestion des taches
  await Future.delayed(const Duration(seconds: 2));
  // Fin de l'écran splash
  FlutterNativeSplash.remove();
  // Initialisation de Firebase
  await Firebase.initializeApp(
    options: DefaultFirebaseOptions.currentPlatform,
  );

  final themeStr = await rootBundle.loadString('assets/theme.json'); // fichier config
  final themeJson = jsonDecode(themeStr); // fichier en obj
  final theme = ThemeDecoder.decodeThemeData(themeJson)!; // theme flutter

  // Firebase
  if(kDebugMode){
    final String host = Platform.isAndroid ? "10.0.2.2" : "localhost";
    // Connexion aux émulateurs
    await FirebaseAuth.instance.useAuthEmulator(host, 9099);
    FirebaseFirestore.instance.useFirestoreEmulator(host, 8082);
    FirebaseStorage.instance.useStorageEmulator(host, 9199);
    FirebaseFunctions.instance.useFunctionsEmulator(host, 5002);
  }
  // Evite d'avoir du cash
  await FirebaseFirestore.instance.terminate();
  await FirebaseFirestore.instance.clearPersistence();

  runApp(
    MultiBlocProvider(
      providers: [
        BlocProvider(
          create: (context) => LoginBloc(AuthenticationService()),
        ),
        BlocProvider(
            create: (context) => RegistrationBloc(AuthenticationService())
        ),
      ],
      child: App(
        theme: theme,
        firebaseAuth: FirebaseAuth.instance,
        firebaseFirestore: FirebaseFirestore.instance,
        firebaseStorage: FirebaseStorage.instance,
        firebaseFunctions: FirebaseFunctions.instance,
      ),
    ),
  );
}

class App extends StatelessWidget {

  final ThemeData theme;
  final FirebaseAuth firebaseAuth;
  final FirebaseFirestore firebaseFirestore;
  final FirebaseStorage firebaseStorage;
  final FirebaseFunctions firebaseFunctions;

  const App(
      {
        super.key,
        required this.theme,
        required this.firebaseAuth,
        required this.firebaseFirestore,
        required this.firebaseStorage,
        required this.firebaseFunctions,
      });

  @override
  Widget build(BuildContext context){
    return ServiceFactory(
        child: MaterialApp(
          debugShowCheckedModeBanner: false,
          title: "Kiwi Corporation",
          localizationsDelegates: const [
            AppLocalizations.delegate,
            GlobalMaterialLocalizations.delegate,
            GlobalWidgetsLocalizations.delegate,
            GlobalCupertinoLocalizations.delegate,
          ],
          initialRoute: '/',
          routes: getApplicationRoutes(),
          onUnknownRoute: (settings) {
            return MaterialPageRoute(
              builder: (context) => Scaffold(
                body: Center(
                  child: Text('No route defined for ${settings.name}'),
                ),
              ),
            );
          },
          supportedLocales: const [
            Locale('en', ''),
            Locale('fr', ''),
          ],
          theme: theme,
          home: StreamBuilder<User?>(
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
          ),
        ),
    );
  }
}