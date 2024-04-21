import 'package:flutter/material.dart';

import '../utils/translate.dart';

class HomeScreen extends StatelessWidget{
  const HomeScreen({Key? key}) : super(key: key);
  @override
  Widget build(BuildContext context){
    return Scaffold(
      body: Center(
        child: Text(t(context)!.helloWorld),
      ),
    );
  }
}