import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:geolocator/geolocator.dart';
import 'package:google_maps_webservice/geocoding.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';

import '../../services/hackathons/hackathon_bloc.dart';
import '../services/hackathons/hackathon_state.dart';

class GeoLocationButton extends StatefulWidget {
  final String token;
  final Function(List<dynamic>) onLocationSortedHackathons;
  final Function(bool) onLoadingStateChanged;

  const GeoLocationButton({
    super.key,
    required this.token,
    required this.onLocationSortedHackathons,
    required this.onLoadingStateChanged,
  });

  @override
  GeoLocationButtonState createState() => GeoLocationButtonState();
}

class GeoLocationButtonState extends State<GeoLocationButton> {
  Position? _currentPosition;

  Future<void> _getCurrentLocation() async {
    widget.onLoadingStateChanged(true);
    setState(() {
    });

    try {
      Position position = await Geolocator.getCurrentPosition(
          desiredAccuracy: LocationAccuracy.high);
      setState(() {
        _currentPosition = position;
      });
      await _sortHackathonsByDistance();
    } catch (e) {
      if (kDebugMode) {
        print('Error getting location: $e');
      }
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('Erreur lors de l\'obtention de la localisation')),
      );
      setState(() {
      });
    } finally {
      widget.onLoadingStateChanged(false);
    }
  }

  Future<void> _sortHackathonsByDistance() async {
    if (_currentPosition == null) return;

    final geocoding = GoogleMapsGeocoding(apiKey: dotenv.env['GOOGLE_PLACES_API_KEY']);
    final hackathonBloc = context.read<HackathonBloc>();

    if (hackathonBloc.state is HackathonLoaded) {
      final hackathons = (hackathonBloc.state as HackathonLoaded).hackathons;
      List<dynamic> hackathonsWithDistance = [];

      for (var hackathon in hackathons) {
        final location = hackathon.location;
        try {
          final response = await geocoding.searchByAddress(location);
          if (response.isOkay && response.results.isNotEmpty) {
            final hackathonLocation = response.results.first.geometry.location;
            final distanceInMeters = Geolocator.distanceBetween(
              _currentPosition!.latitude,
              _currentPosition!.longitude,
              hackathonLocation.lat,
              hackathonLocation.lng,
            );
            final distanceInKm = distanceInMeters / 1000;
            hackathonsWithDistance.add({
              'id': hackathon.id,
              'name': hackathon.name,
              'description': hackathon.description,
              'location': hackathon.location,
              'date': hackathon.date,
              'teams': hackathon.teams,
              'distance': distanceInKm,
            });
          } else {
            if (kDebugMode) {
              print('Error fetching geocoding data: ${response.errorMessage}');
            }
            ScaffoldMessenger.of(context).showSnackBar(
              SnackBar(content: Text('Erreur lors de la récupération des coordonnées pour l\'adresse : ${hackathon.location}')),
            );
            hackathonsWithDistance.add({
              'id': hackathon.id,
              'name': hackathon.name,
              'description': hackathon.description,
              'location': hackathon.location,
              'date': hackathon.date,
              'teams': hackathon.teams,
              'distance': '?',
            });
          }
        } catch (e) {
          if (kDebugMode) {
            print('Exception during geocoding: $e');
          }
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(content: Text('Exception lors de la géocodage pour l\'adresse : ${hackathon.location}')),
          );
          hackathonsWithDistance.add({
            'id': hackathon.id,
            'name': hackathon.name,
            'description': hackathon.description,
            'location': hackathon.location,
            'date': hackathon.date,
            'teams': hackathon.teams,
            'distance': '?',
          });
        }
      }

      hackathonsWithDistance.sort((a, b) {
        if (a['distance'] == '?' && b['distance'] != '?') {
          return 1;
        } else if (a['distance'] != '?' && b['distance'] == '?') {
          return -1;
        } else if (a['distance'] == '?' && b['distance'] == '?') {
          return 0;
        } else {
          return a['distance'].compareTo(b['distance']);
        }
      });

      widget.onLocationSortedHackathons(hackathonsWithDistance);
    }
  }

  @override
  Widget build(BuildContext context) {
    return IconButton(
      icon: const Icon(Icons.location_on),
      onPressed: _getCurrentLocation,
      tooltip: 'Me géolocaliser',
    );
  }
}
