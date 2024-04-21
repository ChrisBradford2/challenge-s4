import 'package:flutter/material.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'Flutter Circular Menu',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: MyHomePage(),
    );
  }
}

class MyHomePage extends StatefulWidget {
  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> with SingleTickerProviderStateMixin {
  late AnimationController animationController;
  late Animation<double> degOneTranslationAnimation, degTwoTranslationAnimation, degThreeTranslationAnimation, rotationAnimation;

  double getRadiansFromDegree(double degree) {
    double unitRadian = 57.295779513;
    return degree / unitRadian;
  }

  @override
  void dispose() {
    animationController.dispose();
    super.dispose();
  }

  @override
  void initState() {
    super.initState();
    animationController = AnimationController(vsync: this, duration: Duration(milliseconds: 250));
    degOneTranslationAnimation = Tween<double>(begin: 0.0, end: 1.0).animate(animationController);
    degTwoTranslationAnimation = Tween<double>(begin: 0.0, end: 1.0).animate(animationController);
    degThreeTranslationAnimation = Tween<double>(begin: 0.0, end: 1.0).animate(animationController);
    rotationAnimation = Tween<double>(begin: 180.0, end: 0.0).animate(CurvedAnimation(
      parent: animationController,
      curve: Curves.easeOut,
    ));

    animationController.addListener(() {
      setState(() {});
    });
  }

  @override
  Widget build(BuildContext context) {
    Size size = MediaQuery.of(context).size;
    return Scaffold(
      body: Container(
        width: size.width,
        height: size.height,
        child: Stack(
          children: <Widget>[
            Positioned(
              right: 30,
              bottom: 30,
              child: Stack(
                alignment: Alignment.bottomRight,
                children: <Widget>[
                  IgnorePointer(
                    child: Container(
                      color: Colors.transparent,
                      height: 150.0,
                      width: 150.0,
                    ),
                  ),
                  Transform.translate(
                    offset: Offset.fromDirection(getRadiansFromDegree(270), (degOneTranslationAnimation.value ?? 0.0) * 100),
                    child: Transform(
                      transform: Matrix4.rotationZ(getRadiansFromDegree(rotationAnimation.value ?? 0.0))..scale(degOneTranslationAnimation.value ?? 0.0),
                      alignment: Alignment.center,
                      child: CircularButton(
                        color: Colors.blue,
                        width: 50,
                        height: 50,
                        icon: Icon(
                          Icons.add,
                          color: Colors.white,
                        ),
                        onClick: () {
                          print('First Button');
                        },
                      ),
                    ),
                  ),
                  Transform.translate(
                    offset: Offset.fromDirection(getRadiansFromDegree(225), (degTwoTranslationAnimation.value ?? 0.0) * 100),
                    child: Transform(
                      transform: Matrix4.rotationZ(getRadiansFromDegree(rotationAnimation.value ?? 0.0))..scale(degTwoTranslationAnimation.value ?? 0.0),
                      alignment: Alignment.center,
                      child: CircularButton(
                        color: Colors.black,
                        width: 50,
                        height: 50,
                        icon: Icon(
                          Icons.camera_alt,
                          color: Colors.white,
                        ),
                        onClick: () {
                          print('Second button');
                        },
                      ),
                    ),
                  ),
                  Transform.translate(
                    offset: Offset.fromDirection(getRadiansFromDegree(180), (degThreeTranslationAnimation.value ?? 0.0) * 100),
                    child: Transform(
                      transform: Matrix4.rotationZ(getRadiansFromDegree(rotationAnimation.value ?? 0.0))..scale(degThreeTranslationAnimation.value ?? 0.0),
                      alignment: Alignment.center,
                      child: CircularButton(
                        color: Colors.orangeAccent,
                        width: 50,
                        height: 50,
                        icon: Icon(
                          Icons.person,
                          color: Colors.white,
                        ),
                        onClick: () {
                          print('Third Button');
                        },
                      ),
                    ),
                  ),
                  Transform(
                    transform: Matrix4.rotationZ(getRadiansFromDegree(rotationAnimation.value ?? 0.0)),
                    alignment: Alignment.center,
                    child: CircularButton(
                      color: Colors.red,
                      width: 60,
                      height: 60,
                      icon: Icon(
                        Icons.menu,
                        color: Colors.white,
                      ),
                      onClick: () {
                        if (animationController.isCompleted) {
                          animationController.reverse();
                        } else {
                          animationController.forward();
                        }
                      },
                    ),
                  )
                ],
              ),
            )
          ],
        ),
      ),
    );
  }
}

class CircularButton extends StatelessWidget {
  final double width;
  final double height;
  final Color color;
  final Icon icon;
  final Function onClick;

  CircularButton({
    required this.color,
    required this.width,
    required this.height,
    required this.icon,
    required this.onClick,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: BoxDecoration(
        color: color,
        shape: BoxShape.circle,
      ),
      width: width,
      height: height,
      child: IconButton(
        icon: icon,
        enableFeedback: true,
        onPressed: () => onClick(),
      ),
    );
  }
}
