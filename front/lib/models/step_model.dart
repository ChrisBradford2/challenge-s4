class Step {
  final int id;
  final String name;
  final String description;

  Step({
    required this.id,
    required this.name,
    required this.description,
  });

  factory Step.fromJson(Map<String, dynamic> json) {
    return Step(
      id: json['id'] is int ? json['id'] : int.tryParse(json['id'].toString()) ?? 0,
      name: json['name'] ?? '',
      description: json['description'] ?? '',
    );
  }
}
