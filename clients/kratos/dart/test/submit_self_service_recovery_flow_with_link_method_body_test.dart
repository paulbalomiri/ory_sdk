import 'package:test/test.dart';
import 'package:ory_kratos_client/ory_kratos_client.dart';

// tests for SubmitSelfServiceRecoveryFlowWithLinkMethodBody
void main() {
  final instance = SubmitSelfServiceRecoveryFlowWithLinkMethodBodyBuilder();
  // TODO add properties to the builder and call build()

  group(SubmitSelfServiceRecoveryFlowWithLinkMethodBody, () {
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

    // Method supports `link` only right now.
    // String method
    test('to test the property `method`', () async {
      // TODO
    });

  });
}
