import 'package:equatable/equatable.dart';

abstract class TeamState extends Equatable {
  const TeamState();

  @override
  List<Object> get props => [];
}

class TeamInitial extends TeamState {}

class TeamLoading extends TeamState {}

class TeamJoined extends TeamState {
  final String message;

  const TeamJoined(this.message);

  @override
  List<Object> get props => [message];
}

class TeamError extends TeamState {
  final String error;

  const TeamError(this.error);

  @override
  List<Object> get props => [error];
}
