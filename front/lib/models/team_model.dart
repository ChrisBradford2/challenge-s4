import 'package:front/models/user_model.dart';
import 'package:front/models/hackathon_model.dart';
import 'package:front/models/submission_model.dart';

class Team {
  final int id;
  final String name;
  final List<User>? users;
  final Hackathon? hackathon;
  final int? hackathonId;
  final int? nbOfMembers;
  final Submission? submission;
  final int? evaluationId;
  final int? stepId;

  Team({
    required this.id,
    required this.name,
    this.users,
    this.hackathon,
    this.hackathonId,
    this.nbOfMembers,
    this.submission,
    this.evaluationId,
    this.stepId,
  });

  factory Team.fromJson(Map<String, dynamic> json) {
    var usersList = json['users'] as List?;
    List<User>? users = usersList?.map((userJson) => User.fromJson(userJson)).toList();

    return Team(
      id: json['id'] is int ? json['id'] : int.tryParse(json['id'].toString()) ?? 0,
      name: json['name'] ?? '',
      users: users,
      hackathon: json['hackathon'] != null ? Hackathon.fromJson(json['hackathon']) : null,
      hackathonId: json['hackathon_id'],
      nbOfMembers: json['nbOfMembers'],
      submission: json['submission'] != null ? Submission.fromJson(json['submission']) : null,
      evaluationId: json['evaluation_id'],
      stepId: json['step_id'],
    );
  }
}
