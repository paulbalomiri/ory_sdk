import 'package:test/test.dart';
import 'package:ory_client/ory_client.dart';

// tests for SubmitSelfServiceRecoveryFlowWithCodeMethodBody
void main() {
  final instance = SubmitSelfServiceRecoveryFlowWithCodeMethodBodyBuilder();
  // TODO add properties to the builder and call build()

  group(SubmitSelfServiceRecoveryFlowWithCodeMethodBody, () {
    // Code from recovery email  Sent to the user once a recovery has been initiated and is used to prove that the user is in possession of the email
    // String code
    test('to test the property `code`', () async {
      // TODO
    });

    // Sending the anti-csrf token is only required for browser login flows.
    // String csrfToken
    test('to test the property `csrfToken`', () async {
      // TODO
    });

    // Email to Recover  Needs to be set when initiating the flow. If the email is a registered recovery email, a recovery link will be sent. If the email is not known, a email with details on what happened will be sent instead.  format: email
    // String email
    test('to test the property `email`', () async {
      // TODO
    });

    // Method supports `link` and `code` only right now.
    // String method
    test('to test the property `method`', () async {
      // TODO
    });

  });
}
