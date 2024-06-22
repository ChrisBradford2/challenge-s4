import 'package:bloc/bloc.dart';
import 'package:front/services/team/team_event.dart';
import 'package:front/services/team/team_state.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import '../../models/user_model.dart';
import '../../utils/config.dart';

class TeamBloc extends Bloc<TeamEvent, TeamState> {
  TeamBloc() : super(TeamInitial()) {
    on<JoinTeam>(_onJoinTeam);
    on<LeaveTeam>(_onLeaveTeam);
  }

  Future<void> _onJoinTeam(JoinTeam event, Emitter<TeamState> emit) async {
    emit(TeamLoading());
    try {
      final response = await http.post(
        Uri.parse('${Config.baseUrl}/teams/${event.teamId}/register'),
        headers: {'Authorization': 'Bearer ${event.token}'},
      );

      if (response.statusCode == 200) {
        final message = jsonDecode(response.body)['message'];
        emit(TeamJoined(message, event.teamId));
      } else {
        emit(const TeamError('Failed to join team'));
      }
    } catch (e) {
      emit(TeamError('Failed to join team: $e'));
    }
  }

  Future<void> _onLeaveTeam(LeaveTeam event, Emitter<TeamState> emit) async {
    emit(TeamLoading());
    try {
      final response = await http.post(
        Uri.parse('${Config.baseUrl}/teams/${event.teamId}/leave'),
        headers: {'Authorization': 'Bearer ${event.token}'},
      );

      if (response.statusCode == 200) {
        final message = jsonDecode(response.body)['message'];
        emit(TeamLeft(message, event.teamId));
      } else {
        emit(const TeamError('Failed to leave team'));
      }
    } catch (e) {
      emit(TeamError('Failed to leave team: $e'));
    }
  }

  Future<User> getUserInfo(String token) async {
    final response = await http.get(
      Uri.parse('${Config.baseUrl}/users/me'),
      headers: {'Authorization': 'Bearer $token'},
    );

    if (response.statusCode == 200) {
      return User.fromJson(jsonDecode(response.body)['data']);
    } else {
      throw Exception('Failed to load user info');
    }
  }
}
