import 'package:bloc/bloc.dart';
import 'dart:convert';
import 'package:http/http.dart' as http;

import 'hackathon_event.dart';
import 'hackathon_state.dart';

// States

// BLoC
class HackathonBloc extends Bloc<HackathonEvent, HackathonState> {
  HackathonBloc() : super(HackathonInitial());

  Stream<HackathonState> mapEventToState(HackathonEvent event) async* {
    if (event is FetchHackathons) {
      yield HackathonLoading();
      try {
        final hackathons = await _fetchHackathons(event.token);
        yield HackathonLoaded(hackathons);
      } catch (e) {
        yield HackathonError(e.toString());
      }
    }
  }

  Future<List<Map<String, String>>> _fetchHackathons(String token) async {
    final response = await http.get(
      Uri.parse('https://localhost/hackathon'),
      headers: {'Authorization': 'Bearer $token'},
    );

    if (response.statusCode == 200) {
      List<dynamic> data = jsonDecode(response.body)['data'];
      return data.map((item) {
        return {
          "name": item["name"].toString(),
          "date": item["start_date"].toString(),
          "location": item["location"].toString(),
        };
      }).toList();
    } else {
      throw Exception('Failed to load hackathons');
    }
  }
}
