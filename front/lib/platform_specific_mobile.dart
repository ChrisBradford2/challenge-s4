import 'dart:io';

import 'package:flutter/foundation.dart';

class MyHttpOverrides extends HttpOverrides {
  @override
  HttpClient createHttpClient(SecurityContext? context) {
    return super.createHttpClient(context)
      ..badCertificateCallback = (X509Certificate cert, String host, int port) => true;
  }
}

void setupPlatformSpecific() {
  if (kDebugMode) {
    HttpOverrides.global = MyHttpOverrides();
  }
}
