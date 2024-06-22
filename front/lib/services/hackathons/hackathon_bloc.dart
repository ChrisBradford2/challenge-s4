import 'package:bloc/bloc.dart';
import 'package:flutter/foundation.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import '../../models/hackathon_model.dart';
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

  Future<List<Hackathon>> _fetchHackathons(String token) async {
    final response = await http.get(
      Uri.parse('${Config.baseUrl}/hackathons'),
      headers: {'Authorization': 'Bearer $token'},
    );

    if (response.statusCode == 200) {
      List<dynamic> data = jsonDecode(response.body)['data'];
      return data.map((item) => Hackathon.fromJson(item)).toList();
    } else {
      throw Exception('Failed to load hackathons');
    }
  }

  void _onFetchSingleHackathons(FetchSingleHackathons event, Emitter<HackathonState> emit) async {
    emit(HackathonLoading());
    try {
      if (kDebugMode) {
        print('Fetching single hackathon for ID: ${event.id}');
      }
      final hackathon = await _fetchSingleHackathon(event.token, event.id);
      if (kDebugMode) {
        print('Fetched hackathon: ${hackathon.name}');
      }
      emit(HackathonLoaded([hackathon]));
      if (kDebugMode) {
        print('HackathonLoaded state emitted');
      }
    } catch (e) {
      emit(HackathonError(e.toString()));
    }
  }

  Future<Hackathon> _fetchSingleHackathon(String token, String id) async {
    final response = await http.get(
      Uri.parse('${Config.baseUrl}/hackathons/$id'),
      headers: {'Authorization': 'Bearer $token'},
    );

    if (kDebugMode) {
      print('HTTP response status: ${response.statusCode}');
      print('HTTP response body: ${response.body}');
    }

    if (response.statusCode == 200) {
      dynamic data = jsonDecode(response.body)['data'];
      return Hackathon.fromJson(data);
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

  Future<List<Hackathon>> _fetchHackathonForUser(String token) async {
    final response = await http.get(
      Uri.parse('${Config.baseUrl}/hackathons/user'),
      headers: {'Authorization': 'Bearer $token'},
    );

    if (response.statusCode == 200) {
      List<dynamic> data = jsonDecode(response.body)['data'];
      return data.map((item) => Hackathon.fromJson(item)).toList();
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

  Future<Hackathon> _addHackathon(String token, Map<String, dynamic> hackathonData) async {
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
      return Hackathon.fromJson(jsonDecode(response.body)['data']);
    } else if (response.statusCode == 400) {
      throw Exception('Bad request');
    } else {
      throw Exception('Failed to add hackathon');
    }
  }
}
