import 'dart:async';
import 'package:bloc/bloc.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'submission_event.dart';
import 'submission_state.dart';
import '../../utils/config.dart';

class SubmissionBloc extends Bloc<SubmissionEvent, SubmissionState> {
  final String token;

  SubmissionBloc({required this.token}) : super(SubmissionInitial());

  Stream<SubmissionState> mapEventToState(SubmissionEvent event) async* {
    if (event is SubmitGitLink) {
      yield* _mapSubmitGitLinkToState(event);
    } else if (event is SubmitFile) {
      yield* _mapSubmitFileToState(event);
    }
  }

  Stream<SubmissionState> _mapSubmitGitLinkToState(SubmitGitLink event) async* {
    yield SubmissionLoading();
    try {
      final response = await http.post(
        Uri.parse('${Config.baseUrl}/submissions'),
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
        body: jsonEncode({
          'team_id': event.teamId,
          'step_id': event.stepId,
          'git_link': event.gitLink,
        }),
      );

      if (response.statusCode == 201) {
        yield const SubmissionSuccess('Git link submitted successfully.');
      } else {
        final error = jsonDecode(response.body)['error'];
        yield SubmissionFailure('Failed to submit git link: $error');
      }
    } catch (e) {
      yield SubmissionFailure('Failed to submit git link: $e');
    }
  }

  Stream<SubmissionState> _mapSubmitFileToState(SubmitFile event) async* {
    yield SubmissionLoading();
    try {
      var request = http.MultipartRequest(
        'POST',
        Uri.parse('${Config.baseUrl}/submissions/upload'),
      );
      request.headers['Authorization'] = 'Bearer $token';
      request.files.add(await http.MultipartFile.fromPath('file', event.file.path));
      request.fields['team_id'] = event.teamId.toString();
      request.fields['step_id'] = event.stepId.toString();

      final response = await request.send();
      final responseBody = await response.stream.bytesToString();

      if (response.statusCode == 201) {
        yield const SubmissionSuccess('File submitted successfully.');
      } else {
        final error = jsonDecode(responseBody)['error'];
        yield SubmissionFailure('Failed to submit file: $error');
      }
    } catch (e) {
      yield SubmissionFailure('Failed to submit file: $e');
    }
  }
}
