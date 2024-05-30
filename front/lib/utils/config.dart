import 'dart:io';

import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:cloud_functions/cloud_functions.dart';
import 'package:firebase_auth/firebase_auth.dart';
import 'package:firebase_storage/firebase_storage.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:flutter_localizations/flutter_localizations.dart';

import '../services/authentication_service.dart';
import '../services/login/login_bloc.dart';
import '../services/logout/logout_bloc.dart';
import '../services/register/register_bloc.dart';
import '../services/hackathons/hackathon_bloc.dart';

class Config {
  static String baseUrl =
      Platform.isAndroid ? "https://10.0.2.2" : "https://localhost";
  static List<LocalizationsDelegate> localizationsDelegates = [
    AppLocalizations.delegate,
    GlobalMaterialLocalizations.delegate,
    GlobalWidgetsLocalizations.delegate,
    GlobalCupertinoLocalizations.delegate,
  ];

  void configureFirebaseEmulators() {
    final host = Platform.isAndroid ? "10.0.2.2" : "localhost";
    FirebaseAuth.instance.useAuthEmulator(host, 9099);
    FirebaseFirestore.instance.useFirestoreEmulator(host, 8082);
    FirebaseStorage.instance.useStorageEmulator(host, 9199);
    FirebaseFunctions.instance.useFunctionsEmulator(host, 5002);
  }

  static List<BlocProvider> blocProviders = [
    BlocProvider<AuthenticationBloc>(
      create: (context) => AuthenticationBloc(AuthenticationService()),
    ),
    BlocProvider<LoginBloc>(
      create: (context) => LoginBloc(AuthenticationService()),
    ),
    BlocProvider<RegistrationBloc>(
      create: (context) => RegistrationBloc(AuthenticationService()),
    ),
    BlocProvider<HackathonBloc>( // Ajout du HackathonBloc
      create: (context) => HackathonBloc(),
    ),
  ];
}
