import 'package:bloc/bloc.dart';

import '../authentication_service.dart';
import 'login_event.dart';
import 'login_state.dart';


class LoginBloc extends Bloc<LoginEvent, LoginState> {
  final AuthenticationService _authService;

  LoginBloc(this._authService) : super(LoginInitial()) {
    on<LoginButtonPressed>(_onLoginButtonPressed);
  }

  void _onLoginButtonPressed(LoginButtonPressed event, Emitter<LoginState> emit) async {
    emit(LoginLoading());
    try {
      bool isLoggedIn = await _authService.login(event.email, event.password);
      if (isLoggedIn) {
        emit(LoginSuccess(token: _authService.token));
      } else {
        emit(LoginFailure(error: "Invalid email or password"));
      }
    } catch (error) {
      emit(LoginFailure(error: error.toString()));
    }
  }
}
