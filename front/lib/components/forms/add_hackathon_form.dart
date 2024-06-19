import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_google_places/flutter_google_places.dart';
import 'package:google_maps_webservice/places.dart';
import 'package:uuid/uuid.dart';

import '../../services/hackathons/hackathon_bloc.dart';
import '../../services/hackathons/hackathon_event.dart';
import '../pickers/datetime_picker.dart';

class AddHackathonForm extends StatelessWidget {
  final String token;
  final String googleApiKey;

  const AddHackathonForm({super.key, required this.token, required this.googleApiKey});

  @override
  Widget build(BuildContext context) {
    final TextEditingController nameController = TextEditingController();
    final TextEditingController descriptionController = TextEditingController();
    final TextEditingController startDateTimeController = TextEditingController();
    final TextEditingController endDateTimeController = TextEditingController();
    final TextEditingController locationController = TextEditingController();
    final TextEditingController maxParticipantsController = TextEditingController();
    final TextEditingController maxParticipantsPerTeamController = TextEditingController();

    return SingleChildScrollView(
      child: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          TextField(
            controller: nameController,
            decoration: const InputDecoration(
              labelText: 'Nom du hackathon',
            ),
          ),
          TextField(
            controller: descriptionController,
            decoration: const InputDecoration(
              labelText: 'Description',
            ),
          ),
          GestureDetector(
            onTap: () => selectDateTime(context, startDateTimeController),
            child: AbsorbPointer(
              child: TextField(
                controller: startDateTimeController,
                decoration: const InputDecoration(
                  labelText: 'Date et heure de début',
                ),
              ),
            ),
          ),
          GestureDetector(
            onTap: () => selectDateTime(context, endDateTimeController),
            child: AbsorbPointer(
              child: TextField(
                controller: endDateTimeController,
                decoration: const InputDecoration(
                  labelText: 'Date et heure de fin',
                ),
              ),
            ),
          ),
          GestureDetector(
            onTap: () async {
              try {
                var sessionToken = Uuid().v4();
                Prediction? p = await PlacesAutocomplete.show(
                  context: context,
                  apiKey: googleApiKey,
                  mode: Mode.overlay,
                  language: "fr",
                  sessionToken: sessionToken,
                  components: [Component(Component.country, "fr")],
                  types: [],
                  radius: 1000,
                  strictbounds: false,
                  region: "fr",
                  hint: "Rechercher une ville",
                );
                if (p != null && p.placeId != null) {
                  GoogleMapsPlaces places = GoogleMapsPlaces(apiKey: googleApiKey);
                  PlacesDetailsResponse detail = await places.getDetailsByPlaceId(p.placeId!);
                  if (detail.status == "OK") {
                    locationController.text = detail.result.formattedAddress ?? 'Address not found';
                  } else {
                    print('Error getting place details: ${detail.errorMessage}');
                    ScaffoldMessenger.of(context).showSnackBar(
                      SnackBar(content: Text('Erreur lors de la récupération des détails du lieu: ${detail.errorMessage}')),
                    );
                  }
                } else {
                  print('Prediction is null or has no placeId');
                }
              } catch (e) {
                print('Error: $e');
                ScaffoldMessenger.of(context).showSnackBar(
                  SnackBar(content: Text('Erreur lors de la récupération des détails du lieu')),
                );
              }
            },
            child: AbsorbPointer(
              child: TextField(
                controller: locationController,
                decoration: const InputDecoration(
                  labelText: 'Lieu',
                ),
              ),
            ),
          ),
          TextField(
            controller: maxParticipantsController,
            decoration: const InputDecoration(
              labelText: 'Participants maximum',
            ),
            keyboardType: TextInputType.number,
          ),
          TextField(
            controller: maxParticipantsPerTeamController,
            decoration: const InputDecoration(
              labelText: 'Participants maximum / équipe',
            ),
            keyboardType: TextInputType.number,
          ),
          Row(
            mainAxisAlignment: MainAxisAlignment.end,
            children: [
              TextButton(
                onPressed: () {
                  Navigator.of(context).pop();
                },
                child: const Text('Annuler'),
              ),
              ElevatedButton(
                onPressed: () {
                  final hackathonData = {
                    'address': locationController.text,
                    'description': descriptionController.text,
                    'end_date': endDateTimeController.text,
                    'latitude': 0, // You need to fetch latitude from details
                    'location': locationController.text,
                    'longitude': 0, // You need to fetch longitude from details
                    'max_participants': int.tryParse(maxParticipantsController.text) ?? 0,
                    'name': nameController.text,
                    'max_participants_per_team': int.tryParse(maxParticipantsPerTeamController.text) ?? 0,
                    'start_date': startDateTimeController.text,
                  };

                  final hackathonBloc = BlocProvider.of<HackathonBloc>(context);
                  hackathonBloc.add(AddHackathon(token, hackathonData));
                  Navigator.of(context).pop();
                },
                child: const Text('Ajouter'),
              ),
            ],
          ),
        ],
      ),
    );
  }
}
