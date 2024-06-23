import 'package:front/models/team_model.dart';

class Hackathon {
  final int id;
  final String name;
  final String description;
  final String location;
  final String date;
  final List<Team> teams;
  final double? distance;

  Hackathon({
    required this.id,
    required this.name,
    required this.description,
    required this.location,
    required this.date,
    required this.teams,
    this.distance,
  });

  factory Hackathon.fromJson(Map<String, dynamic> json) {
    var teamsJson = json['teams'] as List? ?? [];
    List<Team> teamsList = teamsJson.map((i) => Team.fromJson(i)).toList();

    return Hackathon(
      id: json['id'] is int ? json['id'] : int.tryParse(json['id'].toString()) ?? 0,
      name: json['name'] ?? '',
      description: json['description'] ?? '',
      location: json['location'] ?? '',
      date: json['start_date'] ?? '',
      teams: teamsList,
      distance: json['distance']?.toDouble(),
    );
  }
}
