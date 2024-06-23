import 'package:bloc/bloc.dart';

import '../../models/user_model.dart';
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
      if (event.email.isEmpty || event.password.isEmpty) {
        emit(LoginFailure(error: "Email and password cannot be empty"));
        return;
      }

      bool isLoggedIn = await _authService.login(event.email, event.password);
      if (isLoggedIn) {
        emit(LoginSuccess(token: _authService.token));
      } else {
        emit(LoginFailure(error: "Invalid email or password"));
      }
    } catch (error) {
      emit(LoginFailure(error: "An error occurred: $error"));
    }
  }

  Future<User?> getUserDetails(String token) async {
    try {
      final user = await _authService.getUserDetails(token);
      return user;
    } catch (error) {
      print('Error fetching user details: $error');
      return null;
    }
  }
}
