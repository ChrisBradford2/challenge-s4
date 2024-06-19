import 'package:bloc/bloc.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import '../../utils/config.dart';
import 'team_event.dart';
import 'team_state.dart';

class TeamBloc extends Bloc<TeamEvent, TeamState> {
  TeamBloc() : super(TeamInitial()) {
    on<JoinTeam>(_onJoinTeam);
  }

  Future<void> _onJoinTeam(JoinTeam event, Emitter<TeamState> emit) async {
    emit(TeamLoading());
    try {
      final response = await http.post(
        Uri.parse('${Config.baseUrl}/teams/${event.teamId}/register'),
        headers: {
          'Authorization': 'Bearer ${event.token}',
        },
      );

      if (response.statusCode == 200) {
        final message = jsonDecode(response.body)['message'];
        emit(TeamJoined(message));
      } else {
        final error = jsonDecode(response.body)['error'];
        emit(TeamError('Failed to join team: $error'));
      }
    } catch (e) {
      emit(TeamError('Failed to join team: $e'));
    }
  }
}
