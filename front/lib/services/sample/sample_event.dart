part of 'sample_bloc.dart';

abstract class SampleEvent extends Equatable {
  const SampleEvent();

  @override
  List<Object> get props => [];
}

class OnSampleBloc extends SampleEvent {
  const OnSampleBloc();
}
