class UserService {
  Future<Map<String, String>> fetchUser() async {
    await Future.delayed(Duration(microseconds: 100));
    return {'name': "Alice", 'email': "alice@example.com"};
  }
}
