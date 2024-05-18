import 'package:bloc/bloc.dart';
import 'package:front/services/register/register_event.dart';
import 'package:front/services/register/register_state.dart';

import '../authentication_service.dart';

class RegistrationBloc extends Bloc<RegistrationEvent, RegistrationState> {
  final AuthenticationService _authenticationService;

  RegistrationBloc(this._authenticationService) : super(RegistrationInitial()) {
    on<SignUpButtonPressed>(_onSignUpButtonPressed);
  }

  void _onSignUpButtonPressed(SignUpButtonPressed event, Emitter<RegistrationState> emit) async {
    emit(RegistrationLoading());
    try {
      await _authenticationService.register(
          username: event.username,
          lastName: event.lastName,
          firstName: event.firstName,
          email: event.email,
          password: event.password,
          profilePicturePath: event.profilePicturePath
      );
      emit(RegistrationSuccess());
    } catch (e) {
      emit(RegistrationFailure(error: e.toString()));
    }
  }
}
