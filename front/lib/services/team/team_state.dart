import 'package:equatable/equatable.dart';

import '../../models/team_model.dart';

abstract class TeamState extends Equatable {
  const TeamState();

  @override
  List<Object> get props => [];
}

class TeamInitial extends TeamState {}

class TeamLoading extends TeamState {}

class TeamJoined extends TeamState {
  final String message;
  final int teamId;

  const TeamJoined(this.message, this.teamId);

  @override
  List<Object> get props => [message, teamId];
}

class TeamError extends TeamState {
  final String error;

  const TeamError(this.error);

  @override
  List<Object> get props => [error];
}

class TeamLeft extends TeamState {
  final String message;
  final int teamId;

  const TeamLeft(this.message, this.teamId);

  @override
  List<Object> get props => [message, teamId];
}

class TeamsLoaded extends TeamState {
  final List<Team> teams;

  const TeamsLoaded(this.teams);

  @override
  List<Object> get props => [teams];
}

