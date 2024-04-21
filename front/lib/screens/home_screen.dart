import 'package:flutter/material.dart';

import '../utils/translate.dart';

class HomeScreen extends StatelessWidget{
  const HomeScreen({Key? key}) : super(key: key);
  @override
  Widget build(BuildContext context){
    return Scaffold(
      appBar: AppBar(
        title: Text("Home"),
      ),
      body: Center(
        child: Text(t(context)!.helloWorld, style: Theme.of(context).textTheme.displayLarge,),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: (){},
        child: const Icon(Icons.language),
      ),
    );
  }
}