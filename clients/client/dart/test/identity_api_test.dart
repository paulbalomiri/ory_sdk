import 'package:test/test.dart';
import 'package:ory_client/ory_client.dart';


/// tests for IdentityApi
void main() {
  final instance = OryClient().getIdentityApi();

  group(IdentityApi, () {
    // Deactivate a Session
    //
    // Calling this endpoint deactivates the specified session. Session data is not deleted.
    //
    //Future disableSession(String id) async
    test('test disableSession', () async {
      // TODO
    });

  });
}
