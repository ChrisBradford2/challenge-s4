import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:front/services/sample/sample_bloc.dart';

class ServiceFactory extends StatelessWidget {
  final Widget child;
  const ServiceFactory({Key? key, required this.child}): super(key: key);

  @override
  Widget build(BuildContext context) {
    return MultiBlocProvider(
        providers: [
          BlocProvider<SampleBloc>(create: (context) => SampleBloc()),
        ],
        child: child
    );
  }
}
