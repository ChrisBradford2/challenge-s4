import 'package:equatable/equatable.dart';
import 'dart:io';

abstract class SubmissionEvent extends Equatable {
  const SubmissionEvent();

  @override
  List<Object> get props => [];
}

class SubmitGitLink extends SubmissionEvent {
  final int teamId;
  final int stepId;
  final String gitLink;

  const SubmitGitLink({
    required this.teamId,
    required this.stepId,
    required this.gitLink,
  });

  @override
  List<Object> get props => [teamId, stepId, gitLink];
}

class SubmitFile extends SubmissionEvent {
  final int teamId;
  final int stepId;
  final File file;

  const SubmitFile({
    required this.teamId,
    required this.stepId,
    required this.file,
  });

  @override
  List<Object> get props => [teamId, stepId, file];
}
