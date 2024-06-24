import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:file_picker/file_picker.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:front/utils/config.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

import '../../services/team/team_bloc.dart';
import '../../services/team/team_event.dart';
import '../../services/team/team_state.dart';
import '../../models/team_model.dart';

class TeamManagePage extends StatelessWidget {
  final Team team;
  final String token;
  final int evaluationId;
  final int stepId;

  const TeamManagePage({
    super.key,
    required this.team,
    required this.token,
    required this.evaluationId,
    required this.stepId,
  });

  void _leaveTeam(BuildContext context) {
    context.read<TeamBloc>().add(LeaveTeam(team.id, token));
  }

  void _submitWork(BuildContext context, int teamId, int evaluationId, int stepId) {
    showDialog(
      context: context,
      builder: (BuildContext context) {
        String? githubLink;
        PlatformFile? pickedFile;

        return AlertDialog(
          title: Text('Submit Work'),
          content: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              TextField(
                decoration: InputDecoration(labelText: 'GitHub Link'),
                onChanged: (value) {
                  githubLink = value;
                },
              ),
              SizedBox(height: 10),
              ElevatedButton(
                onPressed: () async {
                  FilePickerResult? result = await FilePicker.platform.pickFiles();
                  if (result != null) {
                    pickedFile = result.files.first;
                  }
                },
                child: Text('Upload File'),
              ),
            ],
          ),
          actions: [
            ElevatedButton(
              onPressed: () {
                Navigator.of(context).pop();
              },
              child: Text('Cancel'),
            ),
            ElevatedButton(
              onPressed: () {
                if (githubLink != null && githubLink!.isNotEmpty) {
                  _sendGitHubLink(context, githubLink!, token, teamId, evaluationId, stepId);
                } else if (pickedFile != null) {
                  _uploadFile(context, pickedFile!, token, teamId, evaluationId, stepId);
                }
                Navigator.of(context).pop();
              },
              child: Text('Submit'),
            ),
          ],
        );
      },
    );
  }

  void _sendGitHubLink(BuildContext context, String link, String token, int teamId, int evaluationId, int stepId) async {
    final response = await http.post(
      Uri.parse('${Config.baseUrl}/submissions'),
      headers: {
        'Authorization': 'Bearer $token',
        'Content-Type': 'application/json',
      },
      body: jsonEncode({
        'team_id': teamId,
        'evaluation_id': evaluationId,
        'status': 'submitted',
        'git_link': link,
        'step_id': stepId,
      }),
    );

    if (response.statusCode == 201) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Submission successful')),
      );
    } else {
      if (kDebugMode) {
        print('Failed to submit GitHub link: ${response.body}');
        print('Status code: ${response.statusCode}');
        print('Team ID: $teamId');
      }
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Submission failed')),
      );
    }
  }

  void _uploadFile(BuildContext context, PlatformFile file, String token, int teamId, int evaluationId, int stepId) async {
    var request = http.MultipartRequest(
      'POST',
      Uri.parse('${Config.baseUrl}/submissions/upload'),
    );
    request.headers['Authorization'] = 'Bearer $token';
    request.fields['team_id'] = teamId.toString();
    request.fields['evaluation_id'] = evaluationId.toString();
    request.fields['status'] = 'submitted';
    request.fields['step_id'] = stepId.toString();
    request.files.add(await http.MultipartFile.fromPath('file', file.path!, filename: file.name));

    var response = await request.send();

    if (response.statusCode == 201) {
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('File uploaded successfully')),
      );
    } else {
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('File upload failed')),
      );
    }
  }


  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Manage Team: ${team.name}'),
      ),
      body: BlocListener<TeamBloc, TeamState>(
        listener: (context, state) {
          if (state is TeamLeft) {
            ScaffoldMessenger.of(context).showSnackBar(
              SnackBar(content: Text(state.message)),
            );
            Navigator.pop(context, 'left'); // Return to previous screen
          } else if (state is TeamError) {
            ScaffoldMessenger.of(context).showSnackBar(
              SnackBar(content: Text(state.error)),
            );
          }
        },
        child: Padding(
          padding: const EdgeInsets.all(16.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              const Text('Members:', style: TextStyle(fontSize: 18)),
              const SizedBox(height: 8),
              Expanded(
                child: ListView.builder(
                  itemCount: team.users?.length ?? 0,
                  itemBuilder: (context, index) {
                    final user = team.users![index];
                    return ListTile(
                      title: Text(user.username),
                      subtitle: Text('${user.firstName} ${user.lastName}'),
                    );
                  },
                ),
              ),
              ElevatedButton(
                onPressed: () => _submitWork(context, team.id, evaluationId, stepId),
                child: const Text('Submit Work'),
              ),
              ElevatedButton(
                onPressed: () => _leaveTeam(context),
                style: ElevatedButton.styleFrom(
                  backgroundColor: Colors.red, // Couleur du bouton
                ),
                child: const Text('Leave Team'),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
