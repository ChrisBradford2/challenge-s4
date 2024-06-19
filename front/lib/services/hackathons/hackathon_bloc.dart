import 'package:bloc/bloc.dart';
import 'package:flutter/foundation.dart';
import 'dart:convert';
import 'package:http/http.dart' as http;

import '../../utils/config.dart';
import 'hackathon_event.dart';
import 'hackathon_state.dart';

class HackathonBloc extends Bloc<HackathonEvent, HackathonState> {
  HackathonBloc() : super(HackathonInitial()) {
    on<FetchHackathons>(_onFetchHackathons);
    on<FetchSingleHackathons>(_onFetchSingleHackathons);
    on<AddHackathon>(_onAddHackathon);
    on<FetchHackathonForUser>(_onFetchHackathonForUser);
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

  void _onFetchHackathonForUser(FetchHackathonForUser event, Emitter<HackathonState> emit) async {
    emit(HackathonLoading());
    try {
      final hackathons = await _fetchHackathonForUser(event.token);
      emit(HackathonLoaded(hackathons));
    } catch (e) {
      emit(HackathonError(e.toString()));
    }
  }

  Future<List<Map<String, String>>> _fetchHackathonForUser(String token) async {
    final response = await http.get(
      Uri.parse('${Config.baseUrl}/hackathons/user'), // Update the URL to point to the user-specific endpoint
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

  void _onAddHackathon(AddHackathon event, Emitter<HackathonState> emit) async {
    emit(HackathonLoading());
    try {
      final hackathon = await _addHackathon(event.token, event.hackathonData);
      emit(HackathonAdded(hackathon));
      add(FetchHackathons(event.token));
    } catch (e) {
      if (kDebugMode) {
        print('Error adding hackathon: $e');
      }
      emit(HackathonError(e.toString()));
    }
  }

  Future<Map<String, dynamic>> _addHackathon(String token, Map<String, dynamic> hackathonData) async {
    final url = Uri.parse('${Config.baseUrl}/hackathons/');

    final response = await http.post(
      url,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $token',
      },
      body: jsonEncode(hackathonData),
    );

    if (kDebugMode) {
      print('Response status: ${response.statusCode}');
      print('Response headers: ${response.headers}');
      print('Response body: ${response.body}');
    }

    if (response.statusCode == 201) {
      return jsonDecode(response.body);
    } else if (response.statusCode == 400) {
      throw Exception('Bad request');
    } else {
      throw Exception('Failed to add hackathon');
    }
  }
}
