import 'package:equatable/equatable.dart';

abstract class TeamEvent extends Equatable {
  const TeamEvent();

  @override
  List<Object> get props => [];
}

class JoinTeam extends TeamEvent {
  final int teamId;
  final String token;

  const JoinTeam(this.teamId, this.token);

  @override
  List<Object> get props => [teamId, token];
}

class LeaveTeam extends TeamEvent {
  final int teamId;
  final String token;

  const LeaveTeam(this.teamId, this.token);

  @override
  List<Object> get props => [teamId, token];
}
