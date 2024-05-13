import 'dart:convert';
import 'package:flutter/foundation.dart';
import 'package:http/http.dart' as http;

import '../utils/config.dart';

class AuthenticationService {
  late String _token;
  String get token => _token;

  Future<bool> login(String email, String password) async {
    try {
      final response = await http.post(
        Uri.parse('${Config.baseUrl}/user/login'),
        headers: <String, String>{
          'Content-Type': 'application/json; charset=UTF-8',
        },
        body: jsonEncode(<String, String>{
          'email': email,
          'password': password,
        }),
      );

      if (response.statusCode == 200) {
        if (kDebugMode) {
          print('Login successfull, token: ${jsonDecode(response.body)['token']}');
        }
        _token = jsonDecode(response.body)['token'];
        return true;
      } else {
        if (kDebugMode) {
          print('Login failed: ${response.body}');
        }
        return false;
      }
    } catch (e) {
      if (kDebugMode) {
        print('Login error: $e');
      }
      return false;
    }
  }

  // Destroy the token
  Future<bool> logout() async {
    _token = '';
    // Move to the login page
    return true;
  }
}
