import 'dart:convert';
import 'package:flutter/foundation.dart';
import 'package:http/http.dart' as http;
import 'package:flutter/material.dart';

import '../utils/config.dart';

class UserService {

  Future<String> fetchFirstName(String token) async {
    if (kDebugMode) {
      print('Token: $token');
    }
    try {
      final response = await http.get(
        Uri.parse('${Config.baseUrl}/user/me'),
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
      );
      if (kDebugMode) {
        print('Response status: ${response.statusCode}');
        print('Response body: ${response.body}');
      }
      if (response.statusCode == 401) {
        throw Exception('Unauthorized');
      }
      if (response.statusCode == 200) {
        final data = jsonDecode(response.body);
        if (kDebugMode) {
          print('User data: $data');
        }
        return data['first_name'];
      } else {
        throw Exception('Failed to load user data');
      }
    } catch (e) {
      throw Exception('Error fetching user data: $e');
    }
  }
}
