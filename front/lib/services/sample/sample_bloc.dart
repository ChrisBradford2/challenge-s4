import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';
import 'package:flutter/foundation.dart';

part 'sample_event.dart';
part 'sample_state.dart';

class SampleBloc extends Bloc<SampleEvent, SampleState>{
  SampleBloc(): super(SampleInitial()) {
    on<OnSampleBloc>((event, emit) {
      if (kDebugMode) {
        print("Coucou c'est un test ! ");
      }
    });
  }
}
