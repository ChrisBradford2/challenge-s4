// config_stub.dart
import 'package:flutter/cupertino.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import '../services/authentication_service.dart';
import '../services/login/login_bloc.dart';
import '../services/logout/logout_bloc.dart';
import '../services/register/register_bloc.dart';
import '../services/hackathons/hackathon_bloc.dart';

class Config {
  static String baseUrl = "http://localhost:8080"; // Adjust as needed for your web environment

  static List<LocalizationsDelegate> get localizationsDelegates => [
    AppLocalizations.delegate,
    GlobalMaterialLocalizations.delegate,
    GlobalWidgetsLocalizations.delegate,
    GlobalCupertinoLocalizations.delegate,
  ];

  void configureFirebaseEmulators() {
    // Firebase emulators are not typically used in web environments.
  }

  static List<BlocProvider> get blocProviders => [
    BlocProvider<AuthenticationBloc>(
      create: (context) => AuthenticationBloc(AuthenticationService()),
    ),
    BlocProvider<LoginBloc>(
      create: (context) => LoginBloc(AuthenticationService()),
    ),
    BlocProvider<RegistrationBloc>(
      create: (context) => RegistrationBloc(AuthenticationService()),
    ),
    BlocProvider<HackathonBloc>(
      create: (context) => HackathonBloc(),
    ),
  ];
}
