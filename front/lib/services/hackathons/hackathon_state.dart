import 'package:equatable/equatable.dart';
import '../../models/hackathon_model.dart';

abstract class HackathonState extends Equatable {
  const HackathonState();

  @override
  List<Object> get props => [];
}

class HackathonInitial extends HackathonState {}

class HackathonLoading extends HackathonState {}

class HackathonLoaded extends HackathonState {
  final List<Hackathon> hackathons;

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

class HackathonAdded extends HackathonState {
  final Hackathon hackathon;

  const HackathonAdded(this.hackathon);

  @override
  List<Object> get props => [hackathon];
}
