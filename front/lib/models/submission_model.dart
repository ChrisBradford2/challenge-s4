class Submission {
  final int id;
  final String? gitLink;
  final String? fileUrl;
  final String status;
  final int? stepId;

  Submission({
    required this.id,
    this.gitLink,
    this.fileUrl,
    required this.status,
    this.stepId,
  });

  factory Submission.fromJson(Map<String, dynamic> json) {
    return Submission(
      id: json['id'] is int ? json['id'] : int.tryParse(json['id'].toString()) ?? 0,
      gitLink: json['git_link'],
      fileUrl: json['file_url'],
      status: json['status'] ?? '',
      stepId: json['step_id'] is int ? json['step_id'] : int.tryParse(json['step_id'].toString()),
    );
  }
}
