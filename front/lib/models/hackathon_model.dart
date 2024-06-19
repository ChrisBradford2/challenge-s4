import 'package:front/models/team_model.dart';

class Hackathon {
  final int id;
  final String name;
  final String description;
  final String location;
  final String date;
  final List<Team> teams;
  final double? distance; // Nullable distance field

  Hackathon({
    required this.id,
    required this.name,
    required this.description,
    required this.location,
    required this.date,
    required this.teams,
    this.distance, // Nullable distance parameter
  });

  factory Hackathon.fromJson(Map<String, dynamic> json) {
    var list = json['teams'] as List? ?? [];
    List<Team> teamsList = list.map((i) => Team.fromJson(i)).toList();

    return Hackathon(
      id: json['id'] ?? 0,
      name: json['name'] ?? '',
      description: json['description'] ?? '',
      location: json['location'] ?? '',
      date: json['start_date'] ?? '',
      teams: teamsList,
      distance: json['distance'] != null ? json['distance'].toDouble() : null,
    );
  }
}
