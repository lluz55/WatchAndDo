import 'package:flutter/material.dart';

class NewWatcher extends StatefulWidget {
  @override
  _NewWatcherState createState() => _NewWatcherState();
}

class _NewWatcherState extends State<NewWatcher>
    with SingleTickerProviderStateMixin {
  TextEditingController _controller;
  Animation<Offset> _offsetAnim;
  Animation<double> _fadeAnim;
  AnimationController _animationController;

  @override
  void initState() {
    _controller = TextEditingController(text: "");
    _animationController =
        AnimationController(duration: Duration(milliseconds: 750), vsync: this);

    _fadeAnim =
        Tween<double>(begin: 0.0, end: 1.0).animate(_animationController);
    _offsetAnim = Tween<Offset>(begin: Offset(0, -1), end: Offset.zero)
        .animate(_animationController);

    super.initState();
  }

  @override
  void dispose() {
    _animationController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Container(
        padding: EdgeInsets.only(left: 10.0, right: 10.0, top: 50.0),
        constraints: BoxConstraints(
          maxWidth: 380,
        ),
        child: Column(
          children: <Widget>[
            Text(
              "Adicionar arquivo/diretorio",
              style: _addWatcherStyle,
            ),
            SizedBox(height: 16.0),
            TextField(
              controller: _controller,
              onChanged: (val) {
                if (val.length > 2) {
                  _animationController.forward();
                } else {
                  _animationController.reverse();
                }
              },
              decoration: InputDecoration(
                contentPadding: EdgeInsets.all(14.0),
                border: OutlineInputBorder(),
                hintText: "Ex: C:/Go/bin/go.exe",
                suffixIcon: InkWell(
                  onTap: () {},
                  child: Icon(Icons.folder),
                ),
              ),
            ),
            SizedBox(height: 6.0),
            SlideTransition(
              position: _offsetAnim,
              child: FadeTransition(
                  opacity: _fadeAnim,
                  child: RaisedButton(
                    onPressed: () {
                      _animationController.reverse();
                    },
                    color: Theme.of(context).accentColor,
                    child: Row(
                      mainAxisSize: MainAxisSize.min,
                      mainAxisAlignment: MainAxisAlignment.spaceAround,
                      children: <Widget>[
                        Icon(
                          Icons.add_to_queue,
                          color: Colors.white,
                        ),
                        SizedBox(width: 6.0),
                        Text(
                          "Adicionar",
                          style: TextStyle(color: Colors.white),
                        )
                      ],
                    ),
                  )),
            ),
          ],
        ));
  }

  final _addWatcherStyle = TextStyle(
    fontSize: 16.0,
  );
}
