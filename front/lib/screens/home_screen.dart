import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:front/services/sample/sample_bloc.dart';

import '../utils/translate.dart';

class HomeScreen extends StatelessWidget{
  const HomeScreen({Key? key}) : super(key: key);
  @override
  Widget build(BuildContext context){
    return Scaffold(
      appBar: AppBar(
        title: Text("Home"),
      ),
      body: BlocBuilder<SampleBloc, SampleState>(
        bloc: context.read<SampleBloc>()..add(OnSampleBloc()),
        builder: (context, state) {
          return Center(
            child: Text(
              t(context)!.helloWorld,
              style: Theme.of(context).textTheme.headline1,
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