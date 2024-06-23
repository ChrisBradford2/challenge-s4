import 'package:front/models/team_model.dart';
import 'package:front/models/step_model.dart';

class Hackathon {
  final int id;
  final String name;
  final String description;
  final String location;
  final String date;
  final List<Team> teams;
  final List<Step> steps; // Ajouté
  final double? distance;

  Hackathon({
    required this.id,
    required this.name,
    required this.description,
    required this.location,
    required this.date,
    required this.teams,
    required this.steps, // Ajouté
    this.distance,
  });

  factory Hackathon.fromJson(Map<String, dynamic> json) {
    var teamsJson = json['teams'] as List? ?? [];
    List<Team> teamsList = teamsJson.map((i) => Team.fromJson(i)).toList();

    var stepsJson = json['steps'] as List? ?? [];
    List<Step> stepsList = stepsJson.map((i) => Step.fromJson(i)).toList();

    return Hackathon(
      id: json['id'] is int ? json['id'] : int.tryParse(json['id'].toString()) ?? 0,
      name: json['name'] ?? '',
      description: json['description'] ?? '',
      location: json['location'] ?? '',
      date: json['start_date'] ?? '',
      teams: teamsList,
      steps: stepsList,
      distance: json['distance']?.toDouble(),
    );
  }
}
