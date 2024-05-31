import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:geolocator/geolocator.dart';
import 'package:google_maps_webservice/geocoding.dart';

import '../../services/hackathons/hackathon_bloc.dart';
import '../../services/hackathons/hackathon_event.dart';
import '../../services/hackathons/hackathon_state.dart';
import 'hackathon/hackathon_detail_screen.dart';

class HomeScreen extends StatefulWidget {
  final String token;

  const HomeScreen({super.key, required this.token});

  @override
  HomeScreenState createState() => HomeScreenState();
}

class HomeScreenState extends State<HomeScreen> {
  Position? _currentPosition;
  List<dynamic> _sortedHackathons = [];
  bool _isLoading = false;

  @override
  void initState() {
    super.initState();
    // Fetch hackathons when the screen is initialized
    context.read<HackathonBloc>().add(FetchHackathons(widget.token));
  }

  Future<void> _getCurrentLocation() async {
    setState(() {
      _isLoading = true;
    });

    try {
      Position position = await Geolocator.getCurrentPosition(
          desiredAccuracy: LocationAccuracy.high);
      setState(() {
        _currentPosition = position;
      });
      await _sortHackathonsByDistance();
    } catch (e) {
      print('Error getting location: $e');
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Erreur lors de l\'obtention de la localisation')),
      );
    } finally {
      setState(() {
        _isLoading = false;
      });
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
        final location = hackathon['location'];
        final response = await geocoding.searchByAddress(location!);
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
            ...hackathon,
            'distance': distanceInKm,
          });
        } else if (!response.isOkay || response.errorMessage != null) {
          print('Error getting location for $location: ${response.errorMessage}');
        }
      }

      hackathonsWithDistance.sort((a, b) => a['distance'].compareTo(b['distance']));
      setState(() {
        _sortedHackathons = hackathonsWithDistance;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Accueil"),
        actions: [
          IconButton(
            icon: const Icon(Icons.location_on),
            onPressed: _getCurrentLocation,
          ),
        ],
      ),
      body: BlocListener<HackathonBloc, HackathonState>(
        listener: (context, state) {
          if (state is HackathonAdded) {
            // Fetch the updated list of hackathons
            context.read<HackathonBloc>().add(FetchHackathons(widget.token));
          }
        },
        child: _isLoading
            ? const Center(child: CircularProgressIndicator())
            : BlocBuilder<HackathonBloc, HackathonState>(
          builder: (context, state) {
            if (state is HackathonLoading) {
              return const Center(child: CircularProgressIndicator());
            } else if (state is HackathonLoaded) {
              final hackathons = _currentPosition == null ? state.hackathons : _sortedHackathons;

              return RefreshIndicator(
                onRefresh: () async {
                  context.read<HackathonBloc>().add(FetchHackathons(widget.token));
                },
                child: ListView.builder(
                  itemCount: hackathons.length,
                  itemBuilder: (context, index) {
                    final hackathon = hackathons[index];
                    final id = hackathon['id']!;
                    final distance = hackathon['distance'] != null
                        ? '(${hackathon['distance'].toStringAsFixed(1)} km)'
                        : '';

                    return ListTile(
                      leading: const Icon(Icons.event),
                      title: Text(hackathon['name']!),
                      subtitle: Row(
                        children: [
                          const SizedBox(width: 8),
                          Expanded(
                            child: Text(
                              '${hackathon['date']} - ${hackathon['location']} $distance',
                              overflow: TextOverflow.visible,
                            ),
                          ),
                        ],
                      ),
                      onTap: () {
                        Navigator.push(
                          context,
                          MaterialPageRoute(
                            builder: (context) => HackathonDetailPage(
                              id: id,
                              token: widget.token,
                            ),
                          ),
                        );
                      },
                    );
                  },
                ),
              );
            } else if (state is HackathonError) {
              return Center(child: Text('Error: ${state.message}'));
            } else {
              return Container();
            }
          },
        ),
      ),
    );
  }
}
