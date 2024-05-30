import 'dart:convert';

import 'package:bloc/bloc.dart';
import 'package:front/services/user/user_event.dart';
import 'package:front/services/user/user_state.dart';
import 'package:http/http.dart' as http;

import '../../models/user_model.dart';
import '../../utils/config.dart';

class UserBloc extends Bloc<UserEvent, UserState> {
  UserBloc() : super(UserInitial()) {
    on<FetchUser>(_onFetchUser);
  }

  void _onFetchUser(FetchUser event, Emitter<UserState> emit) async {
    emit(UserLoading());
    try {
      final user = await _fetchUser(event.token);
      emit(UserLoaded(user));
    } catch (e) {
      emit(UserError(e.toString()));
    }
  }

  Future<User> _fetchUser(String token) async {
    final response = await http.get(
      Uri.parse('${Config.baseUrl}/user/me'),
      headers: {'Authorization': 'Bearer $token'},
    );

    if (response.statusCode == 200) {
      return User.fromJson(jsonDecode(response.body));
    } else {
      throw Exception('Failed to load user');
    }
  }
}
