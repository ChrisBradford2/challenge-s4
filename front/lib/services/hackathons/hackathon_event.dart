import 'package:equatable/equatable.dart';

abstract class HackathonEvent extends Equatable {
  const HackathonEvent();

  @override
  List<Object> get props => [];
}

class FetchHackathons extends HackathonEvent {
  final String token;

  const FetchHackathons(this.token);

  @override
  List<Object> get props => [token];
}

class FetchSingleHackathons extends HackathonEvent {
  final String token;
  final String id;

  const FetchSingleHackathons(this.token, this.id);

  @override
  List<Object> get props => [token, id];
}

