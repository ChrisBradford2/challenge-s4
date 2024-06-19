import 'package:flutter/material.dart';
import 'package:google_maps_webservice/places.dart';
import 'package:uuid/uuid.dart';

class AutocompleteTextField extends StatefulWidget {
  final String googleApiKey;
  final TextEditingController controller;

  AutocompleteTextField({required this.googleApiKey, required this.controller});

  @override
  _AutocompleteTextFieldState createState() => _AutocompleteTextFieldState();
}

class _AutocompleteTextFieldState extends State<AutocompleteTextField> {
  List<Prediction> _predictions = [];
  late GoogleMapsPlaces _places;

  @override
  void initState() {
    super.initState();
    _places = GoogleMapsPlaces(apiKey: widget.googleApiKey);
  }

  void _onChanged(String input) async {
    if (input.isEmpty) {
      setState(() {
        _predictions = [];
      });
      return;
    }

    final sessionToken = Uuid().v4();
    final response = await _places.autocomplete(
      input,
      sessionToken: sessionToken,
      components: [Component(Component.country, "fr")],
      language: "fr",
    );

    if (response.isOkay) {
      setState(() {
        _predictions = response.predictions;
      });
    } else {
      print('Error fetching predictions: ${response.errorMessage}');
    }
  }

  void _onPredictionTap(Prediction prediction) async {
    final detail = await _places.getDetailsByPlaceId(prediction.placeId!);
    if (detail.status == "OK") {
      setState(() {
        widget.controller.text = detail.result.formattedAddress ?? 'Address not found';
        _predictions = [];
      });
    } else {
      print('Error getting place details: ${detail.errorMessage}');
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Erreur lors de la récupération des détails du lieu: ${detail.errorMessage}')),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        TextField(
          controller: widget.controller,
          decoration: const InputDecoration(
            labelText: 'Lieu',
          ),
          onChanged: _onChanged,
        ),
        ..._predictions.map(
              (prediction) => ListTile(
            title: Text(prediction.description ?? ""),
            onTap: () => _onPredictionTap(prediction),
          ),
        ),
      ],
    );
  }
}
