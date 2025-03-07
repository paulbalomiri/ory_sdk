import 'package:test/test.dart';
import 'package:ory_client/ory_client.dart';

// tests for NormalizedProjectRevisionHook
void main() {
  final instance = NormalizedProjectRevisionHookBuilder();
  // TODO add properties to the builder and call build()

  group(NormalizedProjectRevisionHook, () {
    // The Hooks Config Key
    // String configKey
    test('to test the property `configKey`', () async {
      // TODO
    });

    // The Project's Revision Creation Date
    // DateTime createdAt
    test('to test the property `createdAt`', () async {
      // TODO
    });

    // The Hook Type
    // String hook
    test('to test the property `hook`', () async {
      // TODO
    });

    // ID of the entry
    // String id
    test('to test the property `id`', () async {
      // TODO
    });

    // The Revision's ID this schema belongs to
    // String projectRevisionId
    test('to test the property `projectRevisionId`', () async {
      // TODO
    });

    // Last Time Project's Revision was Updated
    // DateTime updatedAt
    test('to test the property `updatedAt`', () async {
      // TODO
    });

    // Whether to send the API Key in the HTTP Header or as a HTTP Cookie
    // String webHookConfigAuthApiKeyIn
    test('to test the property `webHookConfigAuthApiKeyIn`', () async {
      // TODO
    });

    // The name of the api key
    // String webHookConfigAuthApiKeyName
    test('to test the property `webHookConfigAuthApiKeyName`', () async {
      // TODO
    });

    // The value of the api key
    // String webHookConfigAuthApiKeyValue
    test('to test the property `webHookConfigAuthApiKeyValue`', () async {
      // TODO
    });

    // The password to be sent in the HTTP Basic Auth Header
    // String webHookConfigAuthBasicAuthPassword
    test('to test the property `webHookConfigAuthBasicAuthPassword`', () async {
      // TODO
    });

    // The username to be sent in the HTTP Basic Auth Header
    // String webHookConfigAuthBasicAuthUser
    test('to test the property `webHookConfigAuthBasicAuthUser`', () async {
      // TODO
    });

    // HTTP Auth Method to use for the Web-Hook
    // String webHookConfigAuthType
    test('to test the property `webHookConfigAuthType`', () async {
      // TODO
    });

    // URI pointing to the JsonNet template used for Web-Hook payload generation. Only used for those HTTP methods, which support HTTP body payloads.
    // String webHookConfigBody
    test('to test the property `webHookConfigBody`', () async {
      // TODO
    });

    // If enabled allows the web hook to interrupt / abort the self-service flow. It only applies to certain flows (registration/verification/login/settings) and requires a valid response format.
    // bool webHookConfigCanInterrupt
    test('to test the property `webHookConfigCanInterrupt`', () async {
      // TODO
    });

    // The HTTP method to use (GET, POST, etc) for the Web-Hook
    // String webHookConfigMethod
    test('to test the property `webHookConfigMethod`', () async {
      // TODO
    });

    // Whether to ignore the Web Hook response
    // bool webHookConfigResponseIgnore
    test('to test the property `webHookConfigResponseIgnore`', () async {
      // TODO
    });

    // The URL the Web-Hook should call
    // String webHookConfigUrl
    test('to test the property `webHookConfigUrl`', () async {
      // TODO
    });

  });
}
