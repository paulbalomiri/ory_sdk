# ory_client (EXPERIMENTAL)
Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed
with a valid Personal Access Token. Public APIs are mostly used in browsers.


This Dart package is automatically generated by the [OpenAPI Generator](https://openapi-generator.tech) project:

- API version: v0.2.0-alpha.60
- Build package: org.openapitools.codegen.languages.DartDioClientCodegen

## Requirements

* Dart 2.12.0 or later OR Flutter 1.26.0 or later
* Dio 4.0.0+

## Installation & Usage

### pub.dev
To use the package from [pub.dev](https://pub.dev), please include the following in pubspec.yaml
```yaml
dependencies:
  ory_client: 0.2.0-alpha.60
```

### Github
If this Dart package is published to Github, please include the following in pubspec.yaml
```yaml
dependencies:
  ory_client:
    git:
      url: https://github.com/ory/sdk.git
      #ref: main
```

### Local development
To use the package from your local drive, please include the following in pubspec.yaml
```yaml
dependencies:
  ory_client:
    path: /path/to/ory_client
```

## Getting Started

Please follow the [installation procedure](#installation--usage) and then run the following:

```dart
import 'package:ory_client/ory_client.dart';


final api = OryClient().getAdminApi();
final String id = id_example; // String | The id of the OAuth 2.0 Client.
final UpdateOAuth2ClientLifespans updateOAuth2ClientLifespans = ; // UpdateOAuth2ClientLifespans | 

try {
    final response = await api.updateOAuth2ClientLifespans(id, updateOAuth2ClientLifespans);
    print(response);
} catch on DioError (e) {
    print("Exception when calling AdminApi->updateOAuth2ClientLifespans: $e\n");
}

```

## Documentation for API Endpoints

All URIs are relative to *https://playground.projects.oryapis.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
[*AdminApi*](doc/AdminApi.md) | [**updateOAuth2ClientLifespans**](doc/AdminApi.md#updateoauth2clientlifespans) | **PUT** /admin/clients/{id}/lifespans | 
[*MetadataApi*](doc/MetadataApi.md) | [**getVersion**](doc/MetadataApi.md#getversion) | **GET** /version | Return Running Software Version.
[*MetadataApi*](doc/MetadataApi.md) | [**isAlive**](doc/MetadataApi.md#isalive) | **GET** /health/alive | Check HTTP Server Status
[*MetadataApi*](doc/MetadataApi.md) | [**isReady**](doc/MetadataApi.md#isready) | **GET** /health/ready | Check HTTP Server and Database Status
[*NamespacesApi*](doc/NamespacesApi.md) | [**getNamespaces**](doc/NamespacesApi.md#getnamespaces) | **GET** /namespaces | Query namespaces
[*ReadApi*](doc/ReadApi.md) | [**getCheck**](doc/ReadApi.md#getcheck) | **GET** /relation-tuples/check/openapi | Check a relation tuple
[*ReadApi*](doc/ReadApi.md) | [**getCheckMirrorStatus**](doc/ReadApi.md#getcheckmirrorstatus) | **GET** /relation-tuples/check | Check a relation tuple
[*ReadApi*](doc/ReadApi.md) | [**getExpand**](doc/ReadApi.md#getexpand) | **GET** /relation-tuples/expand | Expand a Relation Tuple
[*ReadApi*](doc/ReadApi.md) | [**getRelationTuples**](doc/ReadApi.md#getrelationtuples) | **GET** /relation-tuples | Query relation tuples
[*ReadApi*](doc/ReadApi.md) | [**postCheck**](doc/ReadApi.md#postcheck) | **POST** /relation-tuples/check/openapi | Check a relation tuple
[*ReadApi*](doc/ReadApi.md) | [**postCheckMirrorStatus**](doc/ReadApi.md#postcheckmirrorstatus) | **POST** /relation-tuples/check | Check a relation tuple
[*SyntaxApi*](doc/SyntaxApi.md) | [**checkOplSyntax**](doc/SyntaxApi.md#checkoplsyntax) | **POST** /opl/syntax/check | Check the syntax of an OPL file
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminAcceptOAuth2ConsentRequest**](doc/V0alpha2Api.md#adminacceptoauth2consentrequest) | **PUT** /admin/oauth2/auth/requests/consent/accept | Accept an OAuth 2.0 Consent Request
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminAcceptOAuth2LoginRequest**](doc/V0alpha2Api.md#adminacceptoauth2loginrequest) | **PUT** /admin/oauth2/auth/requests/login/accept | Accept an OAuth 2.0 Login Request
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminAcceptOAuth2LogoutRequest**](doc/V0alpha2Api.md#adminacceptoauth2logoutrequest) | **PUT** /admin/oauth2/auth/requests/logout/accept | Accept an OAuth 2.0 Logout Request
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminCreateIdentity**](doc/V0alpha2Api.md#admincreateidentity) | **POST** /admin/identities | Create an Identity
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminCreateJsonWebKeySet**](doc/V0alpha2Api.md#admincreatejsonwebkeyset) | **POST** /admin/keys/{set} | Generate a New JSON Web Key
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminCreateOAuth2Client**](doc/V0alpha2Api.md#admincreateoauth2client) | **POST** /admin/clients | Create an OAuth 2.0 Client
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminCreateSelfServiceRecoveryCode**](doc/V0alpha2Api.md#admincreateselfservicerecoverycode) | **POST** /admin/recovery/code | Create a Recovery Link
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminCreateSelfServiceRecoveryLink**](doc/V0alpha2Api.md#admincreateselfservicerecoverylink) | **POST** /admin/recovery/link | Create a Recovery Link
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminDeleteIdentity**](doc/V0alpha2Api.md#admindeleteidentity) | **DELETE** /admin/identities/{id} | Delete an Identity
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminDeleteIdentitySessions**](doc/V0alpha2Api.md#admindeleteidentitysessions) | **DELETE** /admin/identities/{id}/sessions | Calling this endpoint irrecoverably and permanently deletes and invalidates all sessions that belong to the given Identity.
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminDeleteJsonWebKey**](doc/V0alpha2Api.md#admindeletejsonwebkey) | **DELETE** /admin/keys/{set}/{kid} | Delete a JSON Web Key
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminDeleteJsonWebKeySet**](doc/V0alpha2Api.md#admindeletejsonwebkeyset) | **DELETE** /admin/keys/{set} | Delete a JSON Web Key Set
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminDeleteOAuth2Client**](doc/V0alpha2Api.md#admindeleteoauth2client) | **DELETE** /admin/clients/{id} | Deletes an OAuth 2.0 Client
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminDeleteOAuth2Token**](doc/V0alpha2Api.md#admindeleteoauth2token) | **DELETE** /admin/oauth2/tokens | Delete OAuth2 Access Tokens from a Client
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminDeleteTrustedOAuth2JwtGrantIssuer**](doc/V0alpha2Api.md#admindeletetrustedoauth2jwtgrantissuer) | **DELETE** /admin/trust/grants/jwt-bearer/issuers/{id} | Delete a Trusted OAuth2 JWT Bearer Grant Type Issuer
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminExtendSession**](doc/V0alpha2Api.md#adminextendsession) | **PATCH** /admin/sessions/{id}/extend | Calling this endpoint extends the given session ID. If &#x60;session.earliest_possible_extend&#x60; is set it will only extend the session after the specified time has passed.
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminGetIdentity**](doc/V0alpha2Api.md#admingetidentity) | **GET** /admin/identities/{id} | Get an Identity
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminGetJsonWebKey**](doc/V0alpha2Api.md#admingetjsonwebkey) | **GET** /admin/keys/{set}/{kid} | Fetch a JSON Web Key
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminGetJsonWebKeySet**](doc/V0alpha2Api.md#admingetjsonwebkeyset) | **GET** /admin/keys/{set} | Retrieve a JSON Web Key Set
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminGetOAuth2Client**](doc/V0alpha2Api.md#admingetoauth2client) | **GET** /admin/clients/{id} | Get an OAuth 2.0 Client
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminGetOAuth2ConsentRequest**](doc/V0alpha2Api.md#admingetoauth2consentrequest) | **GET** /admin/oauth2/auth/requests/consent | Get OAuth 2.0 Consent Request Information
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminGetOAuth2LoginRequest**](doc/V0alpha2Api.md#admingetoauth2loginrequest) | **GET** /admin/oauth2/auth/requests/login | Get an OAuth 2.0 Login Request
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminGetOAuth2LogoutRequest**](doc/V0alpha2Api.md#admingetoauth2logoutrequest) | **GET** /admin/oauth2/auth/requests/logout | Get an OAuth 2.0 Logout Request
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminGetTrustedOAuth2JwtGrantIssuer**](doc/V0alpha2Api.md#admingettrustedoauth2jwtgrantissuer) | **GET** /admin/trust/grants/jwt-bearer/issuers/{id} | Get a Trusted OAuth2 JWT Bearer Grant Type Issuer
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminIntrospectOAuth2Token**](doc/V0alpha2Api.md#adminintrospectoauth2token) | **POST** /admin/oauth2/introspect | Introspect OAuth2 Access or Refresh Tokens
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminListCourierMessages**](doc/V0alpha2Api.md#adminlistcouriermessages) | **GET** /admin/courier/messages | List Messages
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminListIdentities**](doc/V0alpha2Api.md#adminlistidentities) | **GET** /admin/identities | List Identities
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminListIdentitySessions**](doc/V0alpha2Api.md#adminlistidentitysessions) | **GET** /admin/identities/{id}/sessions | This endpoint returns all sessions that belong to the given Identity.
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminListOAuth2Clients**](doc/V0alpha2Api.md#adminlistoauth2clients) | **GET** /admin/clients | List OAuth 2.0 Clients
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminListOAuth2SubjectConsentSessions**](doc/V0alpha2Api.md#adminlistoauth2subjectconsentsessions) | **GET** /admin/oauth2/auth/sessions/consent | List OAuth 2.0 Consent Sessions of a Subject
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminListTrustedOAuth2JwtGrantIssuers**](doc/V0alpha2Api.md#adminlisttrustedoauth2jwtgrantissuers) | **GET** /admin/trust/grants/jwt-bearer/issuers | List Trusted OAuth2 JWT Bearer Grant Type Issuers
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminPatchIdentity**](doc/V0alpha2Api.md#adminpatchidentity) | **PATCH** /admin/identities/{id} | Patch an Identity
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminPatchOAuth2Client**](doc/V0alpha2Api.md#adminpatchoauth2client) | **PATCH** /admin/clients/{id} | Patch an OAuth 2.0 Client
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminRejectOAuth2ConsentRequest**](doc/V0alpha2Api.md#adminrejectoauth2consentrequest) | **PUT** /admin/oauth2/auth/requests/consent/reject | Reject an OAuth 2.0 Consent Request
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminRejectOAuth2LoginRequest**](doc/V0alpha2Api.md#adminrejectoauth2loginrequest) | **PUT** /admin/oauth2/auth/requests/login/reject | Reject an OAuth 2.0 Login Request
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminRejectOAuth2LogoutRequest**](doc/V0alpha2Api.md#adminrejectoauth2logoutrequest) | **PUT** /admin/oauth2/auth/requests/logout/reject | Reject an OAuth 2.0 Logout Request
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminRevokeOAuth2ConsentSessions**](doc/V0alpha2Api.md#adminrevokeoauth2consentsessions) | **DELETE** /admin/oauth2/auth/sessions/consent | Revokes OAuth 2.0 Consent Sessions of a Subject for a Specific OAuth 2.0 Client
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminRevokeOAuth2LoginSessions**](doc/V0alpha2Api.md#adminrevokeoauth2loginsessions) | **DELETE** /admin/oauth2/auth/sessions/login | Invalidates All OAuth 2.0 Login Sessions of a Certain User
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminTrustOAuth2JwtGrantIssuer**](doc/V0alpha2Api.md#admintrustoauth2jwtgrantissuer) | **POST** /admin/trust/grants/jwt-bearer/issuers | Trust an OAuth2 JWT Bearer Grant Type Issuer
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminUpdateIdentity**](doc/V0alpha2Api.md#adminupdateidentity) | **PUT** /admin/identities/{id} | Update an Identity
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminUpdateJsonWebKey**](doc/V0alpha2Api.md#adminupdatejsonwebkey) | **PUT** /admin/keys/{set}/{kid} | Update a JSON Web Key
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminUpdateJsonWebKeySet**](doc/V0alpha2Api.md#adminupdatejsonwebkeyset) | **PUT** /admin/keys/{set} | Update a JSON Web Key Set
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**adminUpdateOAuth2Client**](doc/V0alpha2Api.md#adminupdateoauth2client) | **PUT** /admin/clients/{id} | Update an OAuth 2.0 Client
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**createProject**](doc/V0alpha2Api.md#createproject) | **POST** /projects | Create a Project
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**createProjectApiKey**](doc/V0alpha2Api.md#createprojectapikey) | **POST** /projects/{project}/tokens | Create API Token
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**createSelfServiceLogoutFlowUrlForBrowsers**](doc/V0alpha2Api.md#createselfservicelogoutflowurlforbrowsers) | **GET** /self-service/logout/browser | Create a Logout URL for Browsers
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**deleteProjectApiKey**](doc/V0alpha2Api.md#deleteprojectapikey) | **DELETE** /projects/{project}/tokens/{token_id} | Delete API Token
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**discoverJsonWebKeys**](doc/V0alpha2Api.md#discoverjsonwebkeys) | **GET** /.well-known/jwks.json | Discover JSON Web Keys
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**discoverOidcConfiguration**](doc/V0alpha2Api.md#discoveroidcconfiguration) | **GET** /.well-known/openid-configuration | OpenID Connect Discovery
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**dynamicClientRegistrationCreateOAuth2Client**](doc/V0alpha2Api.md#dynamicclientregistrationcreateoauth2client) | **POST** /oauth2/register | Register an OAuth 2.0 Client using the OpenID / OAuth2 Dynamic Client Registration Management Protocol
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**dynamicClientRegistrationDeleteOAuth2Client**](doc/V0alpha2Api.md#dynamicclientregistrationdeleteoauth2client) | **DELETE** /oauth2/register/{id} | Deletes an OAuth 2.0 Client using the OpenID / OAuth2 Dynamic Client Registration Management Protocol
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**dynamicClientRegistrationGetOAuth2Client**](doc/V0alpha2Api.md#dynamicclientregistrationgetoauth2client) | **GET** /oauth2/register/{id} | Get an OAuth 2.0 Client using the OpenID / OAuth2 Dynamic Client Registration Management Protocol
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**dynamicClientRegistrationUpdateOAuth2Client**](doc/V0alpha2Api.md#dynamicclientregistrationupdateoauth2client) | **PUT** /oauth2/register/{id} | Update an OAuth 2.0 Client using the OpenID / OAuth2 Dynamic Client Registration Management Protocol
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**getIdentitySchema**](doc/V0alpha2Api.md#getidentityschema) | **GET** /schemas/{id} | 
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**getOidcUserInfo**](doc/V0alpha2Api.md#getoidcuserinfo) | **GET** /userinfo | OpenID Connect Userinfo
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**getProject**](doc/V0alpha2Api.md#getproject) | **GET** /projects/{project_id} | Get a Project
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**getProjectMembers**](doc/V0alpha2Api.md#getprojectmembers) | **GET** /projects/{project_id}/members | Get all members associated with this project.
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**getSelfServiceError**](doc/V0alpha2Api.md#getselfserviceerror) | **GET** /self-service/errors | Get Self-Service Errors
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**getSelfServiceLoginFlow**](doc/V0alpha2Api.md#getselfserviceloginflow) | **GET** /self-service/login/flows | Get Login Flow
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**getSelfServiceRecoveryFlow**](doc/V0alpha2Api.md#getselfservicerecoveryflow) | **GET** /self-service/recovery/flows | Get Recovery Flow
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**getSelfServiceRegistrationFlow**](doc/V0alpha2Api.md#getselfserviceregistrationflow) | **GET** /self-service/registration/flows | Get Registration Flow
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**getSelfServiceSettingsFlow**](doc/V0alpha2Api.md#getselfservicesettingsflow) | **GET** /self-service/settings/flows | Get Settings Flow
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**getSelfServiceVerificationFlow**](doc/V0alpha2Api.md#getselfserviceverificationflow) | **GET** /self-service/verification/flows | Get Verification Flow
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**getWebAuthnJavaScript**](doc/V0alpha2Api.md#getwebauthnjavascript) | **GET** /.well-known/ory/webauthn.js | Get WebAuthn JavaScript
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**initializeSelfServiceLoginFlowForBrowsers**](doc/V0alpha2Api.md#initializeselfserviceloginflowforbrowsers) | **GET** /self-service/login/browser | Initialize Login Flow for Browsers
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**initializeSelfServiceLoginFlowWithoutBrowser**](doc/V0alpha2Api.md#initializeselfserviceloginflowwithoutbrowser) | **GET** /self-service/login/api | Initialize Login Flow for APIs, Services, Apps, ...
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**initializeSelfServiceRecoveryFlowForBrowsers**](doc/V0alpha2Api.md#initializeselfservicerecoveryflowforbrowsers) | **GET** /self-service/recovery/browser | Initialize Recovery Flow for Browsers
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**initializeSelfServiceRecoveryFlowWithoutBrowser**](doc/V0alpha2Api.md#initializeselfservicerecoveryflowwithoutbrowser) | **GET** /self-service/recovery/api | Initialize Recovery Flow for APIs, Services, Apps, ...
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**initializeSelfServiceRegistrationFlowForBrowsers**](doc/V0alpha2Api.md#initializeselfserviceregistrationflowforbrowsers) | **GET** /self-service/registration/browser | Initialize Registration Flow for Browsers
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**initializeSelfServiceRegistrationFlowWithoutBrowser**](doc/V0alpha2Api.md#initializeselfserviceregistrationflowwithoutbrowser) | **GET** /self-service/registration/api | Initialize Registration Flow for APIs, Services, Apps, ...
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**initializeSelfServiceSettingsFlowForBrowsers**](doc/V0alpha2Api.md#initializeselfservicesettingsflowforbrowsers) | **GET** /self-service/settings/browser | Initialize Settings Flow for Browsers
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**initializeSelfServiceSettingsFlowWithoutBrowser**](doc/V0alpha2Api.md#initializeselfservicesettingsflowwithoutbrowser) | **GET** /self-service/settings/api | Initialize Settings Flow for APIs, Services, Apps, ...
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**initializeSelfServiceVerificationFlowForBrowsers**](doc/V0alpha2Api.md#initializeselfserviceverificationflowforbrowsers) | **GET** /self-service/verification/browser | Initialize Verification Flow for Browser Clients
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**initializeSelfServiceVerificationFlowWithoutBrowser**](doc/V0alpha2Api.md#initializeselfserviceverificationflowwithoutbrowser) | **GET** /self-service/verification/api | Initialize Verification Flow for APIs, Services, Apps, ...
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**listIdentitySchemas**](doc/V0alpha2Api.md#listidentityschemas) | **GET** /schemas | 
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**listProjectApiKeys**](doc/V0alpha2Api.md#listprojectapikeys) | **GET** /projects/{project}/tokens | List a Project&#39;s API Tokens
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**listProjects**](doc/V0alpha2Api.md#listprojects) | **GET** /projects | List All Projects
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**listSessions**](doc/V0alpha2Api.md#listsessions) | **GET** /sessions | This endpoints returns all other active sessions that belong to the logged-in user. The current session can be retrieved by calling the &#x60;/sessions/whoami&#x60; endpoint.
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**patchProject**](doc/V0alpha2Api.md#patchproject) | **PATCH** /projects/{project_id} | Patch an Ory Cloud Project Configuration&#x60;
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**performOAuth2AuthorizationFlow**](doc/V0alpha2Api.md#performoauth2authorizationflow) | **GET** /oauth2/auth | The OAuth 2.0 Authorize Endpoint
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**performOAuth2TokenFlow**](doc/V0alpha2Api.md#performoauth2tokenflow) | **POST** /oauth2/token | The OAuth 2.0 Token Endpoint
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**performOidcFrontOrBackChannelLogout**](doc/V0alpha2Api.md#performoidcfrontorbackchannellogout) | **GET** /oauth2/sessions/logout | OpenID Connect Front- or Back-channel Enabled Logout
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**purgeProject**](doc/V0alpha2Api.md#purgeproject) | **DELETE** /projects/{project_id} | Irrecoverably Purge a Project
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**removeProjectMember**](doc/V0alpha2Api.md#removeprojectmember) | **DELETE** /projects/{project_id}/members/{member_id} | Remove a member associated with this project. This also sets their invite status to &#x60;REMOVED&#x60;.
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**revokeOAuth2Token**](doc/V0alpha2Api.md#revokeoauth2token) | **POST** /oauth2/revoke | Revoke an OAuth2 Access or Refresh Token
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**revokeSession**](doc/V0alpha2Api.md#revokesession) | **DELETE** /sessions/{id} | Calling this endpoint invalidates the specified session. The current session cannot be revoked. Session data are not deleted.
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**revokeSessions**](doc/V0alpha2Api.md#revokesessions) | **DELETE** /sessions | Calling this endpoint invalidates all except the current session that belong to the logged-in user. Session data are not deleted.
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**submitSelfServiceLoginFlow**](doc/V0alpha2Api.md#submitselfserviceloginflow) | **POST** /self-service/login | Submit a Login Flow
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**submitSelfServiceLogoutFlow**](doc/V0alpha2Api.md#submitselfservicelogoutflow) | **GET** /self-service/logout | Complete Self-Service Logout
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**submitSelfServiceLogoutFlowWithoutBrowser**](doc/V0alpha2Api.md#submitselfservicelogoutflowwithoutbrowser) | **DELETE** /self-service/logout/api | Perform Logout for APIs, Services, Apps, ...
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**submitSelfServiceRecoveryFlow**](doc/V0alpha2Api.md#submitselfservicerecoveryflow) | **POST** /self-service/recovery | Complete Recovery Flow
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**submitSelfServiceRegistrationFlow**](doc/V0alpha2Api.md#submitselfserviceregistrationflow) | **POST** /self-service/registration | Submit a Registration Flow
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**submitSelfServiceSettingsFlow**](doc/V0alpha2Api.md#submitselfservicesettingsflow) | **POST** /self-service/settings | Complete Settings Flow
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**submitSelfServiceVerificationFlow**](doc/V0alpha2Api.md#submitselfserviceverificationflow) | **POST** /self-service/verification | Complete Verification Flow
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**toSession**](doc/V0alpha2Api.md#tosession) | **GET** /sessions/whoami | Check Who the Current HTTP Session Belongs To
[*V0alpha2Api*](doc/V0alpha2Api.md) | [**updateProject**](doc/V0alpha2Api.md#updateproject) | **PUT** /projects/{project_id} | Update an Ory Cloud Project Configuration
[*WriteApi*](doc/WriteApi.md) | [**createRelationTuple**](doc/WriteApi.md#createrelationtuple) | **PUT** /admin/relation-tuples | Create a Relation Tuple
[*WriteApi*](doc/WriteApi.md) | [**deleteRelationTuples**](doc/WriteApi.md#deleterelationtuples) | **DELETE** /admin/relation-tuples | Delete Relation Tuples
[*WriteApi*](doc/WriteApi.md) | [**patchRelationTuples**](doc/WriteApi.md#patchrelationtuples) | **PATCH** /admin/relation-tuples | Patch Multiple Relation Tuples


## Documentation For Models

 - [AcceptOAuth2ConsentRequest](doc/AcceptOAuth2ConsentRequest.md)
 - [AcceptOAuth2ConsentRequestSession](doc/AcceptOAuth2ConsentRequestSession.md)
 - [AcceptOAuth2LoginRequest](doc/AcceptOAuth2LoginRequest.md)
 - [ActiveProject](doc/ActiveProject.md)
 - [AdminCreateIdentityBody](doc/AdminCreateIdentityBody.md)
 - [AdminCreateIdentityImportCredentialsOidc](doc/AdminCreateIdentityImportCredentialsOidc.md)
 - [AdminCreateIdentityImportCredentialsOidcConfig](doc/AdminCreateIdentityImportCredentialsOidcConfig.md)
 - [AdminCreateIdentityImportCredentialsOidcProvider](doc/AdminCreateIdentityImportCredentialsOidcProvider.md)
 - [AdminCreateIdentityImportCredentialsPassword](doc/AdminCreateIdentityImportCredentialsPassword.md)
 - [AdminCreateIdentityImportCredentialsPasswordConfig](doc/AdminCreateIdentityImportCredentialsPasswordConfig.md)
 - [AdminCreateJsonWebKeySetBody](doc/AdminCreateJsonWebKeySetBody.md)
 - [AdminCreateSelfServiceRecoveryCodeBody](doc/AdminCreateSelfServiceRecoveryCodeBody.md)
 - [AdminCreateSelfServiceRecoveryLinkBody](doc/AdminCreateSelfServiceRecoveryLinkBody.md)
 - [AdminIdentityImportCredentials](doc/AdminIdentityImportCredentials.md)
 - [AdminTrustOAuth2JwtGrantIssuerBody](doc/AdminTrustOAuth2JwtGrantIssuerBody.md)
 - [AdminUpdateIdentityBody](doc/AdminUpdateIdentityBody.md)
 - [AuthenticatorAssuranceLevel](doc/AuthenticatorAssuranceLevel.md)
 - [CloudAccount](doc/CloudAccount.md)
 - [CnameSettings](doc/CnameSettings.md)
 - [CourierMessageStatus](doc/CourierMessageStatus.md)
 - [CourierMessageType](doc/CourierMessageType.md)
 - [CreateCustomHostnameBody](doc/CreateCustomHostnameBody.md)
 - [CreateProjectApiKeyRequest](doc/CreateProjectApiKeyRequest.md)
 - [CreateProjectBody](doc/CreateProjectBody.md)
 - [CreateSubscriptionPayload](doc/CreateSubscriptionPayload.md)
 - [ErrorAuthenticatorAssuranceLevelNotSatisfied](doc/ErrorAuthenticatorAssuranceLevelNotSatisfied.md)
 - [ExpandTree](doc/ExpandTree.md)
 - [GenericError](doc/GenericError.md)
 - [GenericErrorContent](doc/GenericErrorContent.md)
 - [GetCheckResponse](doc/GetCheckResponse.md)
 - [GetManagedIdentitySchemaLocation](doc/GetManagedIdentitySchemaLocation.md)
 - [GetNamespacesResponse](doc/GetNamespacesResponse.md)
 - [GetRelationTuplesResponse](doc/GetRelationTuplesResponse.md)
 - [GetVersion200Response](doc/GetVersion200Response.md)
 - [HandledOAuth2ConsentRequest](doc/HandledOAuth2ConsentRequest.md)
 - [HandledOAuth2LoginRequest](doc/HandledOAuth2LoginRequest.md)
 - [HandledOAuth2LogoutRequest](doc/HandledOAuth2LogoutRequest.md)
 - [Headers](doc/Headers.md)
 - [HealthNotReadyStatus](doc/HealthNotReadyStatus.md)
 - [HealthStatus](doc/HealthStatus.md)
 - [IDTokenClaims](doc/IDTokenClaims.md)
 - [Identity](doc/Identity.md)
 - [IdentityCredentials](doc/IdentityCredentials.md)
 - [IdentityCredentialsOidc](doc/IdentityCredentialsOidc.md)
 - [IdentityCredentialsOidcProvider](doc/IdentityCredentialsOidcProvider.md)
 - [IdentityCredentialsPassword](doc/IdentityCredentialsPassword.md)
 - [IdentityCredentialsType](doc/IdentityCredentialsType.md)
 - [IdentitySchemaContainer](doc/IdentitySchemaContainer.md)
 - [IdentitySchemaPreset](doc/IdentitySchemaPreset.md)
 - [IdentityState](doc/IdentityState.md)
 - [IntrospectedOAuth2Token](doc/IntrospectedOAuth2Token.md)
 - [InvitePayload](doc/InvitePayload.md)
 - [IsOwnerForProjectBySlug](doc/IsOwnerForProjectBySlug.md)
 - [IsOwnerForProjectBySlugPayload](doc/IsOwnerForProjectBySlugPayload.md)
 - [IsReady200Response](doc/IsReady200Response.md)
 - [IsReady503Response](doc/IsReady503Response.md)
 - [JsonError](doc/JsonError.md)
 - [JsonPatch](doc/JsonPatch.md)
 - [JsonWebKey](doc/JsonWebKey.md)
 - [JsonWebKeySet](doc/JsonWebKeySet.md)
 - [KetoNamespace](doc/KetoNamespace.md)
 - [ManagedIdentitySchema](doc/ManagedIdentitySchema.md)
 - [ManagedIdentitySchemaValidationResult](doc/ManagedIdentitySchemaValidationResult.md)
 - [Message](doc/Message.md)
 - [Namespace](doc/Namespace.md)
 - [NeedsPrivilegedSessionError](doc/NeedsPrivilegedSessionError.md)
 - [NormalizedProject](doc/NormalizedProject.md)
 - [NormalizedProjectRevision](doc/NormalizedProjectRevision.md)
 - [NormalizedProjectRevisionHook](doc/NormalizedProjectRevisionHook.md)
 - [NormalizedProjectRevisionIdentitySchema](doc/NormalizedProjectRevisionIdentitySchema.md)
 - [NormalizedProjectRevisionThirdPartyProvider](doc/NormalizedProjectRevisionThirdPartyProvider.md)
 - [NullPlan](doc/NullPlan.md)
 - [OAuth2AccessRequest](doc/OAuth2AccessRequest.md)
 - [OAuth2ApiError](doc/OAuth2ApiError.md)
 - [OAuth2Client](doc/OAuth2Client.md)
 - [OAuth2ConsentRequest](doc/OAuth2ConsentRequest.md)
 - [OAuth2ConsentRequestOpenIDConnectContext](doc/OAuth2ConsentRequestOpenIDConnectContext.md)
 - [OAuth2ConsentSession](doc/OAuth2ConsentSession.md)
 - [OAuth2ConsentSessionExpiresAt](doc/OAuth2ConsentSessionExpiresAt.md)
 - [OAuth2LoginRequest](doc/OAuth2LoginRequest.md)
 - [OAuth2LogoutRequest](doc/OAuth2LogoutRequest.md)
 - [OAuth2TokenResponse](doc/OAuth2TokenResponse.md)
 - [OidcConfiguration](doc/OidcConfiguration.md)
 - [OidcUserInfo](doc/OidcUserInfo.md)
 - [Pagination](doc/Pagination.md)
 - [PaginationHeaders](doc/PaginationHeaders.md)
 - [ParseError](doc/ParseError.md)
 - [PatchDelta](doc/PatchDelta.md)
 - [PostCheckOplSyntaxResponse](doc/PostCheckOplSyntaxResponse.md)
 - [PreviousOAuth2ConsentSession](doc/PreviousOAuth2ConsentSession.md)
 - [Project](doc/Project.md)
 - [ProjectApiKey](doc/ProjectApiKey.md)
 - [ProjectHost](doc/ProjectHost.md)
 - [ProjectInvite](doc/ProjectInvite.md)
 - [ProjectMetadata](doc/ProjectMetadata.md)
 - [ProjectServiceIdentity](doc/ProjectServiceIdentity.md)
 - [ProjectServiceOAuth2](doc/ProjectServiceOAuth2.md)
 - [ProjectServicePermission](doc/ProjectServicePermission.md)
 - [ProjectServices](doc/ProjectServices.md)
 - [ProvisionMockSubscriptionPayload](doc/ProvisionMockSubscriptionPayload.md)
 - [QuotaCustomDomains](doc/QuotaCustomDomains.md)
 - [QuotaProjectMemberSeats](doc/QuotaProjectMemberSeats.md)
 - [RecoveryIdentityAddress](doc/RecoveryIdentityAddress.md)
 - [RefreshTokenHookRequest](doc/RefreshTokenHookRequest.md)
 - [RefreshTokenHookResponse](doc/RefreshTokenHookResponse.md)
 - [RejectOAuth2Request](doc/RejectOAuth2Request.md)
 - [RelationQuery](doc/RelationQuery.md)
 - [RelationTuple](doc/RelationTuple.md)
 - [RevokedSessions](doc/RevokedSessions.md)
 - [SchemaPatch](doc/SchemaPatch.md)
 - [SelfServiceBrowserLocationChangeRequiredError](doc/SelfServiceBrowserLocationChangeRequiredError.md)
 - [SelfServiceError](doc/SelfServiceError.md)
 - [SelfServiceFlowExpiredError](doc/SelfServiceFlowExpiredError.md)
 - [SelfServiceLoginFlow](doc/SelfServiceLoginFlow.md)
 - [SelfServiceLogoutUrl](doc/SelfServiceLogoutUrl.md)
 - [SelfServiceRecoveryCode](doc/SelfServiceRecoveryCode.md)
 - [SelfServiceRecoveryFlow](doc/SelfServiceRecoveryFlow.md)
 - [SelfServiceRecoveryFlowState](doc/SelfServiceRecoveryFlowState.md)
 - [SelfServiceRecoveryLink](doc/SelfServiceRecoveryLink.md)
 - [SelfServiceRegistrationFlow](doc/SelfServiceRegistrationFlow.md)
 - [SelfServiceSettingsFlow](doc/SelfServiceSettingsFlow.md)
 - [SelfServiceSettingsFlowState](doc/SelfServiceSettingsFlowState.md)
 - [SelfServiceVerificationFlow](doc/SelfServiceVerificationFlow.md)
 - [SelfServiceVerificationFlowState](doc/SelfServiceVerificationFlowState.md)
 - [Session](doc/Session.md)
 - [SessionAuthenticationMethod](doc/SessionAuthenticationMethod.md)
 - [SessionDevice](doc/SessionDevice.md)
 - [SettingsProfileFormConfig](doc/SettingsProfileFormConfig.md)
 - [SourcePosition](doc/SourcePosition.md)
 - [StripeCustomerResponse](doc/StripeCustomerResponse.md)
 - [SubjectSet](doc/SubjectSet.md)
 - [SubmitSelfServiceFlowWithWebAuthnRegistrationMethod](doc/SubmitSelfServiceFlowWithWebAuthnRegistrationMethod.md)
 - [SubmitSelfServiceLoginFlowBody](doc/SubmitSelfServiceLoginFlowBody.md)
 - [SubmitSelfServiceLoginFlowWithLookupSecretMethodBody](doc/SubmitSelfServiceLoginFlowWithLookupSecretMethodBody.md)
 - [SubmitSelfServiceLoginFlowWithOidcMethodBody](doc/SubmitSelfServiceLoginFlowWithOidcMethodBody.md)
 - [SubmitSelfServiceLoginFlowWithPasswordMethodBody](doc/SubmitSelfServiceLoginFlowWithPasswordMethodBody.md)
 - [SubmitSelfServiceLoginFlowWithTotpMethodBody](doc/SubmitSelfServiceLoginFlowWithTotpMethodBody.md)
 - [SubmitSelfServiceLoginFlowWithWebAuthnMethodBody](doc/SubmitSelfServiceLoginFlowWithWebAuthnMethodBody.md)
 - [SubmitSelfServiceLogoutFlowWithoutBrowserBody](doc/SubmitSelfServiceLogoutFlowWithoutBrowserBody.md)
 - [SubmitSelfServiceRecoveryFlowBody](doc/SubmitSelfServiceRecoveryFlowBody.md)
 - [SubmitSelfServiceRecoveryFlowWithCodeMethodBody](doc/SubmitSelfServiceRecoveryFlowWithCodeMethodBody.md)
 - [SubmitSelfServiceRecoveryFlowWithLinkMethodBody](doc/SubmitSelfServiceRecoveryFlowWithLinkMethodBody.md)
 - [SubmitSelfServiceRegistrationFlowBody](doc/SubmitSelfServiceRegistrationFlowBody.md)
 - [SubmitSelfServiceRegistrationFlowWithOidcMethodBody](doc/SubmitSelfServiceRegistrationFlowWithOidcMethodBody.md)
 - [SubmitSelfServiceRegistrationFlowWithPasswordMethodBody](doc/SubmitSelfServiceRegistrationFlowWithPasswordMethodBody.md)
 - [SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody](doc/SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody.md)
 - [SubmitSelfServiceSettingsFlowBody](doc/SubmitSelfServiceSettingsFlowBody.md)
 - [SubmitSelfServiceSettingsFlowWithLookupMethodBody](doc/SubmitSelfServiceSettingsFlowWithLookupMethodBody.md)
 - [SubmitSelfServiceSettingsFlowWithOidcMethodBody](doc/SubmitSelfServiceSettingsFlowWithOidcMethodBody.md)
 - [SubmitSelfServiceSettingsFlowWithPasswordMethodBody](doc/SubmitSelfServiceSettingsFlowWithPasswordMethodBody.md)
 - [SubmitSelfServiceSettingsFlowWithProfileMethodBody](doc/SubmitSelfServiceSettingsFlowWithProfileMethodBody.md)
 - [SubmitSelfServiceSettingsFlowWithTotpMethodBody](doc/SubmitSelfServiceSettingsFlowWithTotpMethodBody.md)
 - [SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody](doc/SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody.md)
 - [SubmitSelfServiceVerificationFlowBody](doc/SubmitSelfServiceVerificationFlowBody.md)
 - [SubmitSelfServiceVerificationFlowWithLinkMethodBody](doc/SubmitSelfServiceVerificationFlowWithLinkMethodBody.md)
 - [Subscription](doc/Subscription.md)
 - [SuccessfulOAuth2RequestResponse](doc/SuccessfulOAuth2RequestResponse.md)
 - [SuccessfulProjectUpdate](doc/SuccessfulProjectUpdate.md)
 - [SuccessfulSelfServiceLoginWithoutBrowser](doc/SuccessfulSelfServiceLoginWithoutBrowser.md)
 - [SuccessfulSelfServiceRegistrationWithoutBrowser](doc/SuccessfulSelfServiceRegistrationWithoutBrowser.md)
 - [TokenPagination](doc/TokenPagination.md)
 - [TokenPaginationHeaders](doc/TokenPaginationHeaders.md)
 - [TrustedOAuth2JwtGrantIssuer](doc/TrustedOAuth2JwtGrantIssuer.md)
 - [TrustedOAuth2JwtGrantJsonWebKey](doc/TrustedOAuth2JwtGrantJsonWebKey.md)
 - [UiContainer](doc/UiContainer.md)
 - [UiNode](doc/UiNode.md)
 - [UiNodeAnchorAttributes](doc/UiNodeAnchorAttributes.md)
 - [UiNodeAttributes](doc/UiNodeAttributes.md)
 - [UiNodeImageAttributes](doc/UiNodeImageAttributes.md)
 - [UiNodeInputAttributes](doc/UiNodeInputAttributes.md)
 - [UiNodeMeta](doc/UiNodeMeta.md)
 - [UiNodeScriptAttributes](doc/UiNodeScriptAttributes.md)
 - [UiNodeTextAttributes](doc/UiNodeTextAttributes.md)
 - [UiText](doc/UiText.md)
 - [UpdateCustomHostnameBody](doc/UpdateCustomHostnameBody.md)
 - [UpdateOAuth2ClientLifespans](doc/UpdateOAuth2ClientLifespans.md)
 - [UpdateProject](doc/UpdateProject.md)
 - [UpdateSubscriptionPayload](doc/UpdateSubscriptionPayload.md)
 - [VerifiableIdentityAddress](doc/VerifiableIdentityAddress.md)
 - [Version](doc/Version.md)
 - [Warning](doc/Warning.md)


## Documentation For Authorization


## basic

- **Type**: HTTP basic authentication

## bearer

- **Type**: HTTP basic authentication

## oauth2

- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://hydra.demo.ory.sh/oauth2/auth
- **Scopes**: 
 - **offline**: A scope required when requesting refresh tokens (alias for `offline_access`)
 - **offline_access**: A scope required when requesting refresh tokens
 - **openid**: Request an OpenID Connect ID Token

## oryAccessToken

- **Type**: HTTP basic authentication


## Author

support@ory.sh

