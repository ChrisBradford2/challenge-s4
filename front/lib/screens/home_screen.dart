import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../services/hackathons/hackathon_bloc.dart';
import '../../services/hackathons/hackathon_event.dart';
import '../../services/hackathons/hackathon_state.dart';
import '../models/hackathon_model.dart';
import '../widgets/geolocation_button.dart';
import 'hackathon/hackathon_detail_screen.dart';

class HomeScreen extends StatefulWidget {
  final String token;

  const HomeScreen({super.key, required this.token});

  @override
  HomeScreenState createState() => HomeScreenState();
}

class HomeScreenState extends State<HomeScreen> {
  List<Hackathon> _sortedHackathons = [];
  bool _isLoading = false;
  final bool _errorOccurred = false;

  @override
  void initState() {
    super.initState();
    // Fetch hackathons when the screen is initialized
    context.read<HackathonBloc>().add(FetchHackathons(widget.token));
  }

  void _handleLocationSortedHackathons(List<Hackathon> sortedHackathons) {
    setState(() {
      _sortedHackathons = sortedHackathons;
    });
  }

  void _handleLoadingStateChanged(bool isLoading) {
    setState(() {
      _isLoading = isLoading;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Accueil"),
        actions: [
          GeoLocationButton(
            token: widget.token,
            onLocationSortedHackathons: (List<dynamic> hackathons) => _handleLocationSortedHackathons(hackathons.cast<Hackathon>()),
            onLoadingStateChanged: _handleLoadingStateChanged,
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
              final hackathons = _errorOccurred || _sortedHackathons.isEmpty
                  ? state.hackathons
                  : _sortedHackathons;

              return RefreshIndicator(
                onRefresh: () async {
                  context.read<HackathonBloc>().add(FetchHackathons(widget.token));
                },
                child: ListView.builder(
                  itemCount: hackathons.length,
                  itemBuilder: (context, index) {
                    final hackathon = hackathons[index];
                    final distance = hackathon.distance != null
                        ? '(${hackathon.distance!.toStringAsFixed(1)} km)'
                        : '(?)';

                    return ListTile(
                      leading: const Icon(Icons.event),
                      title: Text(hackathon.name),
                      subtitle: Row(
                        children: [
                          const SizedBox(width: 8),
                          Expanded(
                            child: Text(
                              '${hackathon.date} - ${hackathon.location} $distance',
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
                              id: hackathon.id.toString(), // Convert int to String
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
