import 'package:bloc/bloc.dart';
import 'dart:convert';
import 'package:http/http.dart' as http;

import '../../utils/config.dart';
import 'hackathon_event.dart';
import 'hackathon_state.dart';

// BLoC
class HackathonBloc extends Bloc<HackathonEvent, HackathonState> {
  HackathonBloc() : super(HackathonInitial()) {
    on<FetchHackathons>(_onFetchHackathons);
    on<FetchSingleHackathons>(_onFetchSingleHackathons);
    on<AddHackathon>(_onAddHackathon);
  }

  void _onFetchHackathons(FetchHackathons event, Emitter<HackathonState> emit) async {
    emit(HackathonLoading());
    try {
      final hackathons = await _fetchHackathons(event.token);
      emit(HackathonLoaded(hackathons));
    } catch (e) {
      emit(HackathonError(e.toString()));
    }
  }

  Future<List<Map<String, String>>> _fetchHackathons(String token) async {
    final response = await http.get(
      Uri.parse('${Config.baseUrl}/hackathons'),
      headers: {'Authorization': 'Bearer $token'},
    );

    if (response.statusCode == 200) {
      List<dynamic> data = jsonDecode(response.body)['data'];
      return data.map((item) {
        return {
          "id": item["id"]?.toString() ?? '',
          "name": item["name"]?.toString() ?? 'Unknown',
          "date": item["start_date"]?.toString() ?? 'Unknown',
          "location": item["location"]?.toString() ?? 'Unknown',
        };
      }).toList();
    } else {
      throw Exception('Failed to load hackathons');
    }
  }

  void _onFetchSingleHackathons(FetchSingleHackathons event, Emitter<HackathonState> emit) async {
    emit(HackathonLoading());
    try {
      final hackathons = await _fetchSingleHackathon(event.token, event.id);
      emit(HackathonLoaded([hackathons]));
    } catch (e) {
      emit(HackathonError(e.toString()));
    }
  }

  Future<Map<String, String>> _fetchSingleHackathon(String token, String id) async {
    final response = await http.get(
      Uri.parse('${Config.baseUrl}/hackathons/$id'),
      headers: {'Authorization': 'Bearer $token'},
    );

    if (response.statusCode == 200) {
      dynamic data = jsonDecode(response.body)['data'];
      return {
        "name": data["name"].toString(),
        "date": data["start_date"].toString(),
        "location": data["location"].toString(),
      };
    } else {
      throw Exception('Failed to load hackathon');
    }
  }

  void _onAddHackathon(AddHackathon event, Emitter<HackathonState> emit) async {
    emit(HackathonLoading());
    try {
      print('Calling _addHackathon with data: ${event.hackathonData}');
      final hackathon = await _addHackathon(event.token, event.hackathonData);
      print('Hackathon added: $hackathon');
      emit(HackathonAdded(hackathon));
      add(FetchHackathons(event.token));
    } catch (e) {
      print('Error adding hackathon: $e');
      emit(HackathonError(e.toString()));
    }
  }

  Future<Map<String, dynamic>> _addHackathon(String token, Map<String, dynamic> hackathonData) async {
    final url = Uri.parse('${Config.baseUrl}/hackathons/');
    print('Sending POST request to $url with body: $hackathonData');

    final response = await http.post(
      url,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $token',
      },
      body: jsonEncode(hackathonData),
    );

    print('Response status: ${response.statusCode}');
    print('Response headers: ${response.headers}');
    print('Response body: ${response.body}');

    if (response.statusCode == 201) {
      return jsonDecode(response.body);
    } else if (response.statusCode == 400) {
      throw Exception('Bad request');
    } else {
      throw Exception('Failed to add hackathon');
    }
  }
}
