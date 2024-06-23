import 'package:flutter/material.dart';
import '../../models/team_model.dart';
import 'package:url_launcher/url_launcher.dart';

class TeamDetailPage extends StatelessWidget {
  final Team team;
  final String token;

  const TeamDetailPage({super.key, required this.team, required this.token});

  void _downloadFile(Uri url) async {
    if (await canLaunchUrl(url)) {
      await launchUrl(url);
    } else {
      print('Could not launch $url');
      throw 'Could not launch $url';
    }
  }

  @override
  Widget build(BuildContext context) {
    print('Team: ${team.name}');
    print('Submission: ${team.submission}');
    if (team.submission != null) {
      print('Submission fileUrl: ${team.submission!.fileUrl}');
      print('Submission gitLink: ${team.submission!.gitLink}');
    }

    return Scaffold(
      appBar: AppBar(
        title: Text('Team: ${team.name}'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text('Team Members:', style: TextStyle(fontSize: 18)),
            SizedBox(height: 8),
            Expanded(
              child: team.users != null && team.users!.isNotEmpty
                  ? ListView.builder(
                itemCount: team.users!.length,
                itemBuilder: (context, index) {
                  final user = team.users![index];
                  return ListTile(
                    title: Text(user.username),
                    subtitle: Text('${user.firstName} ${user.lastName}'),
                  );
                },
              )
                  : Center(
                child: Text('None', style: TextStyle(fontSize: 18, fontStyle: FontStyle.italic)),
              ),
            ),
            SizedBox(height: 16),
            Text('Submission:', style: TextStyle(fontSize: 18)),
            SizedBox(height: 8),
            team.submission != null
                ? Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                if (team.submission!.gitLink != null && team.submission!.gitLink!.isNotEmpty)
                  Padding(
                    padding: const EdgeInsets.only(bottom: 8.0),
                    child: Text('GitHub Link: ${team.submission!.gitLink}', style: TextStyle(fontSize: 16)),
                  ),
                if (team.submission!.fileUrl != null && team.submission!.fileUrl!.isNotEmpty)
                  Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(
                        'File URL:',
                        style: TextStyle(fontSize: 16),
                      ),
                      SizedBox(height: 8),
                      ElevatedButton.icon(
                        onPressed: () {
                          print('Downloading file from: ${team.submission!.fileUrl}');
                          _downloadFile(Uri.parse(team.submission!.fileUrl!));
                        },
                        icon: Icon(Icons.download),
                        label: Text('Download'),
                        style: ElevatedButton.styleFrom(
                          minimumSize: Size(double.infinity, 36), // Button takes full width
                        ),
                      ),
                    ],
                  ),
              ],
            )
                : Text('Not ready yet', style: TextStyle(fontSize: 18, fontStyle: FontStyle.italic)),
          ],
        ),
      ),
    );
  }
}
