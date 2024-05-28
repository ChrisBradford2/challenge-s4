import 'package:equatable/equatable.dart';

abstract class HackathonState extends Equatable {
  const HackathonState();

  @override
  List<Object> get props => [];
}

class HackathonInitial extends HackathonState {}

class HackathonLoading extends HackathonState {}

class HackathonLoaded extends HackathonState {
  final List<Map<String, String>> hackathons;

  const HackathonLoaded(this.hackathons);

  @override
  List<Object> get props => [hackathons];
}

class HackathonError extends HackathonState {
  final String message;

  const HackathonError(this.message);

  @override
  List<Object> get props => [message];
}
