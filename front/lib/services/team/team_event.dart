import 'package:equatable/equatable.dart';

abstract class TeamEvent extends Equatable {
  const TeamEvent();

  @override
  List<Object> get props => [];
}

class JoinTeam extends TeamEvent {
  final String teamId;
  final String token;

  const JoinTeam(this.teamId, this.token);

  @override
  List<Object> get props => [teamId, token];
}
