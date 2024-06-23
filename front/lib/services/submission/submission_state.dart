import 'package:equatable/equatable.dart';

abstract class SubmissionState extends Equatable {
  const SubmissionState();

  @override
  List<Object> get props => [];
}

class SubmissionInitial extends SubmissionState {}

class SubmissionLoading extends SubmissionState {}

class SubmissionSuccess extends SubmissionState {
  final String message;

  const SubmissionSuccess(this.message);

  @override
  List<Object> get props => [message];
}

class SubmissionFailure extends SubmissionState {
  final String error;

  const SubmissionFailure(this.error);

  @override
  List<Object> get props => [error];
}
