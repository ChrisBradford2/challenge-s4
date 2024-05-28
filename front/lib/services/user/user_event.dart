import 'package:equatable/equatable.dart';

abstract class UserEvent extends Equatable {
  const UserEvent();

  @override
  List<Object> get props => [];
}

class FetchUser extends UserEvent {
  final String token;

  const FetchUser(this.token);

  @override
  List<Object> get props => [token];
}
