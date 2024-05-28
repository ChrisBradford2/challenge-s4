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
