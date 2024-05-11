import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:front/services/sample/sample_bloc.dart';

import '../utils/translate.dart';

class HomeScreen extends StatelessWidget{
  const HomeScreen({super.key});
  @override
  Widget build(BuildContext context){
    return Scaffold(
      appBar: AppBar(
        title: const Text("Home"),
      ),
      body: BlocBuilder<SampleBloc, SampleState>(
        bloc: context.read<SampleBloc>()..add(const OnSampleBloc()),
        builder: (context, state) {
          return Center(
            child: Text(
              t(context)!.helloWorld,
              style: Theme.of(context).textTheme.displayLarge,
            ),
          );
      },)
          ,
      floatingActionButton: FloatingActionButton(
        onPressed: (){},
        child: const Icon(Icons.language),
      ),
    );
  }
}