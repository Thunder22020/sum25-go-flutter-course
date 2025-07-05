import 'package:flutter/material.dart';
import 'package:lab02_chat/user_service.dart';

// UserProfile displays and updates user info
class UserProfile extends StatefulWidget {
  final UserService userService;

  const UserProfile({Key? key, required this.userService}) : super(key: key);

  @override
  State<UserProfile> createState() => _UserProfileState();
}

class _UserProfileState extends State<UserProfile> {
  Map<String, String>? _user;
  String? _error;
  bool _loading = false;

  @override
  void initState() {
    super.initState();
    _loadUser();
  }

  void _loadUser() async {
    setState(() {
      _loading = true;
      _error = null;
    });

    try {
      final user = await widget.userService.fetchUser();
      setState(() {
        _user = user;
      });
    } catch (e) {
      setState(() {
        _error = 'error: $e';
      });
    } finally {
      setState(() {
        _loading = false;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('User Profile')),
      body: _loading
          ? const Center(child: CircularProgressIndicator())
          : _error != null
              ? Center(
                  child:
                      Text(_error!, style: const TextStyle(color: Colors.red)),
                )
              : _user != null
                  ? Padding(
                      padding: const EdgeInsets.all(16.0),
                      child: Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          Row(
                            children: [
                              Text(
                                "Name: ",
                                style: const TextStyle(fontSize: 20),
                              ),
                              Text(
                                "${_user!['name']}",
                                style: const TextStyle(fontSize: 20),
                              ),
                            ],
                          ),
                          const SizedBox(height: 8),
                          Row(
                            children: [
                              Text(
                                "Email: ",
                                style: const TextStyle(fontSize: 20),
                              ),
                              Text(
                                "${_user!['email']}",
                                style: const TextStyle(fontSize: 20),
                              ),
                            ],
                          )
                        ],
                      ),
                    )
                  : const Center(child: Text("No user data")),
    );
  }
}
