class User {
  final int id;
  final String username;
  final String lastName;
  final String firstName;
  final String profilePicture;
  final String email;

  User({
    required this.id,
    required this.username,
    required this.lastName,
    required this.firstName,
    required this.profilePicture,
    required this.email,
  });

  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      id: json['id'],
      username: json['username'],
      lastName: json['last_name'],
      firstName: json['first_name'],
      profilePicture: json['profile_picture'],
      email: json['email'],
    );
  }
}
