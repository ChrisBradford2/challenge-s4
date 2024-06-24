import 'package:http/http.dart' as http;
import 'dart:convert';
import '../../models/team_model.dart';
import '../../utils/config.dart';

class TeamRepository {
  Future<List<Team>> fetchTeamsForHackathon(int hackathonId, String token) async {
    final response = await http.get(
      Uri.parse('${Config.baseUrl}/hackathons/$hackathonId/teams'),
      headers: {'Authorization': 'Bearer $token'},
    );

    if (response.statusCode == 200) {
      final Map<String, dynamic> jsonResponse = jsonDecode(response.body);
      final List<dynamic> teamJson = jsonResponse['data']; // Accéder à la clé 'data'
      return teamJson.map((json) => Team.fromJson(json)).toList();
    } else if (response.statusCode == 403) {
      throw Exception('You are not allowed to access this resource');
    } else {
      throw Exception('Failed to load teams');
    }
  }
}
