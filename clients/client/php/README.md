# ory/client-php

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed
with a valid Personal Access Token. Public APIs are mostly used in browsers.



## Installation & Usage

### Requirements

PHP 7.3 and later.
Should also work with PHP 8.0 but has not been tested.

### Composer

To install the bindings via [Composer](https://getcomposer.org/), add the following to `composer.json`:

```json
{
  "repositories": [
    {
      "type": "vcs",
      "url": "https://github.com/ory/client-php.git"
    }
  ],
  "require": {
    "ory/client-php": "*@dev"
  }
}
```

Then run `composer install`

### Manual Installation

Download the files and include `autoload.php`:

```php
<?php
require_once('/path/to/ory/client-php/vendor/autoload.php');
```

## Getting Started

Please follow the [installation procedure](#installation--usage) and then run the following:

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');



// Configure Bearer authorization: oryAccessToken
$config = Ory\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new Ory\Client\Api\IdentityApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 'id_example'; // string | ID is the session's ID.

try {
    $apiInstance->disableSession($id);
} catch (Exception $e) {
    echo 'Exception when calling IdentityApi->disableSession: ', $e->getMessage(), PHP_EOL;
}

```

## API Endpoints

All URIs are relative to *https://playground.projects.oryapis.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*IdentityApi* | [**disableSession**](docs/Api/IdentityApi.md#disablesession) | **DELETE** /admin/sessions/{id} | Deactivate a Session
*JwkApi* | [**createJsonWebKeySet**](docs/Api/JwkApi.md#createjsonwebkeyset) | **POST** /admin/keys/{set} | Create JSON Web Key
*JwkApi* | [**deleteJsonWebKey**](docs/Api/JwkApi.md#deletejsonwebkey) | **DELETE** /admin/keys/{set}/{kid} | Delete JSON Web Key
*JwkApi* | [**deleteJsonWebKeySet**](docs/Api/JwkApi.md#deletejsonwebkeyset) | **DELETE** /admin/keys/{set} | Delete JSON Web Key Set
*JwkApi* | [**getJsonWebKey**](docs/Api/JwkApi.md#getjsonwebkey) | **GET** /admin/keys/{set}/{kid} | Get JSON Web Key
*JwkApi* | [**getJsonWebKeySet**](docs/Api/JwkApi.md#getjsonwebkeyset) | **GET** /admin/keys/{set} | Retrieve a JSON Web Key Set
*JwkApi* | [**setJsonWebKey**](docs/Api/JwkApi.md#setjsonwebkey) | **PUT** /admin/keys/{set}/{kid} | Set JSON Web Key
*JwkApi* | [**setJsonWebKeySet**](docs/Api/JwkApi.md#setjsonwebkeyset) | **PUT** /admin/keys/{set} | Update a JSON Web Key Set
*MetadataApi* | [**getVersion**](docs/Api/MetadataApi.md#getversion) | **GET** /version | Return Running Software Version.
*MetadataApi* | [**isAlive**](docs/Api/MetadataApi.md#isalive) | **GET** /health/alive | Check HTTP Server Status
*MetadataApi* | [**isReady**](docs/Api/MetadataApi.md#isready) | **GET** /health/ready | Check HTTP Server and Database Status
*NamespacesApi* | [**getNamespaces**](docs/Api/NamespacesApi.md#getnamespaces) | **GET** /namespaces | Query namespaces
*OAuth2Api* | [**acceptOAuth2ConsentRequest**](docs/Api/OAuth2Api.md#acceptoauth2consentrequest) | **PUT** /admin/oauth2/auth/requests/consent/accept | Accept OAuth 2.0 Consent Request
*OAuth2Api* | [**acceptOAuth2LoginRequest**](docs/Api/OAuth2Api.md#acceptoauth2loginrequest) | **PUT** /admin/oauth2/auth/requests/login/accept | Accept OAuth 2.0 Login Request
*OAuth2Api* | [**acceptOAuth2LogoutRequest**](docs/Api/OAuth2Api.md#acceptoauth2logoutrequest) | **PUT** /admin/oauth2/auth/requests/logout/accept | Accept OAuth 2.0 Session Logout Request
*OAuth2Api* | [**createOAuth2Client**](docs/Api/OAuth2Api.md#createoauth2client) | **POST** /admin/clients | Create OAuth 2.0 Client
*OAuth2Api* | [**deleteOAuth2Client**](docs/Api/OAuth2Api.md#deleteoauth2client) | **DELETE** /admin/clients/{id} | Delete OAuth 2.0 Client
*OAuth2Api* | [**deleteOAuth2Token**](docs/Api/OAuth2Api.md#deleteoauth2token) | **DELETE** /admin/oauth2/tokens | Delete OAuth 2.0 Access Tokens from specific OAuth 2.0 Client
*OAuth2Api* | [**deleteTrustedOAuth2JwtGrantIssuer**](docs/Api/OAuth2Api.md#deletetrustedoauth2jwtgrantissuer) | **DELETE** /admin/trust/grants/jwt-bearer/issuers/{id} | Delete Trusted OAuth2 JWT Bearer Grant Type Issuer
*OAuth2Api* | [**getOAuth2Client**](docs/Api/OAuth2Api.md#getoauth2client) | **GET** /admin/clients/{id} | Get an OAuth 2.0 Client
*OAuth2Api* | [**getOAuth2ConsentRequest**](docs/Api/OAuth2Api.md#getoauth2consentrequest) | **GET** /admin/oauth2/auth/requests/consent | Get OAuth 2.0 Consent Request
*OAuth2Api* | [**getOAuth2LoginRequest**](docs/Api/OAuth2Api.md#getoauth2loginrequest) | **GET** /admin/oauth2/auth/requests/login | Get OAuth 2.0 Login Request
*OAuth2Api* | [**getOAuth2LogoutRequest**](docs/Api/OAuth2Api.md#getoauth2logoutrequest) | **GET** /admin/oauth2/auth/requests/logout | Get OAuth 2.0 Session Logout Request
*OAuth2Api* | [**getTrustedOAuth2JwtGrantIssuer**](docs/Api/OAuth2Api.md#gettrustedoauth2jwtgrantissuer) | **GET** /admin/trust/grants/jwt-bearer/issuers/{id} | Get Trusted OAuth2 JWT Bearer Grant Type Issuer
*OAuth2Api* | [**introspectOAuth2Token**](docs/Api/OAuth2Api.md#introspectoauth2token) | **POST** /admin/oauth2/introspect | Introspect OAuth2 Access and Refresh Tokens
*OAuth2Api* | [**listOAuth2Clients**](docs/Api/OAuth2Api.md#listoauth2clients) | **GET** /admin/clients | List OAuth 2.0 Clients
*OAuth2Api* | [**listOAuth2ConsentSessions**](docs/Api/OAuth2Api.md#listoauth2consentsessions) | **GET** /admin/oauth2/auth/sessions/consent | List OAuth 2.0 Consent Sessions of a Subject
*OAuth2Api* | [**listTrustedOAuth2JwtGrantIssuers**](docs/Api/OAuth2Api.md#listtrustedoauth2jwtgrantissuers) | **GET** /admin/trust/grants/jwt-bearer/issuers | List Trusted OAuth2 JWT Bearer Grant Type Issuers
*OAuth2Api* | [**oAuth2Authorize**](docs/Api/OAuth2Api.md#oauth2authorize) | **GET** /oauth2/auth | OAuth 2.0 Authorize Endpoint
*OAuth2Api* | [**oauth2TokenExchange**](docs/Api/OAuth2Api.md#oauth2tokenexchange) | **POST** /oauth2/token | The OAuth 2.0 Token Endpoint
*OAuth2Api* | [**patchOAuth2Client**](docs/Api/OAuth2Api.md#patchoauth2client) | **PATCH** /admin/clients/{id} | Patch OAuth 2.0 Client
*OAuth2Api* | [**rejectOAuth2ConsentRequest**](docs/Api/OAuth2Api.md#rejectoauth2consentrequest) | **PUT** /admin/oauth2/auth/requests/consent/reject | Reject OAuth 2.0 Consent Request
*OAuth2Api* | [**rejectOAuth2LoginRequest**](docs/Api/OAuth2Api.md#rejectoauth2loginrequest) | **PUT** /admin/oauth2/auth/requests/login/reject | Reject OAuth 2.0 Login Request
*OAuth2Api* | [**rejectOAuth2LogoutRequest**](docs/Api/OAuth2Api.md#rejectoauth2logoutrequest) | **PUT** /admin/oauth2/auth/requests/logout/reject | Reject OAuth 2.0 Session Logout Request
*OAuth2Api* | [**revokeOAuth2ConsentSessions**](docs/Api/OAuth2Api.md#revokeoauth2consentsessions) | **DELETE** /admin/oauth2/auth/sessions/consent | Revoke OAuth 2.0 Consent Sessions of a Subject
*OAuth2Api* | [**revokeOAuth2LoginSessions**](docs/Api/OAuth2Api.md#revokeoauth2loginsessions) | **DELETE** /admin/oauth2/auth/sessions/login | Revokes All OAuth 2.0 Login Sessions of a Subject
*OAuth2Api* | [**revokeOAuth2Token**](docs/Api/OAuth2Api.md#revokeoauth2token) | **POST** /oauth2/revoke | Revoke OAuth 2.0 Access or Refresh Token
*OAuth2Api* | [**setOAuth2Client**](docs/Api/OAuth2Api.md#setoauth2client) | **PUT** /admin/clients/{id} | Set OAuth 2.0 Client
*OAuth2Api* | [**setOAuth2ClientLifespans**](docs/Api/OAuth2Api.md#setoauth2clientlifespans) | **PUT** /admin/clients/{id}/lifespans | Set OAuth2 Client Token Lifespans
*OAuth2Api* | [**trustOAuth2JwtGrantIssuer**](docs/Api/OAuth2Api.md#trustoauth2jwtgrantissuer) | **POST** /admin/trust/grants/jwt-bearer/issuers | Trust OAuth2 JWT Bearer Grant Type Issuer
*OidcApi* | [**createOidcDynamicClient**](docs/Api/OidcApi.md#createoidcdynamicclient) | **POST** /oauth2/register | Register OAuth2 Client using OpenID Dynamic Client Registration
*OidcApi* | [**deleteOidcDynamicClient**](docs/Api/OidcApi.md#deleteoidcdynamicclient) | **DELETE** /oauth2/register/{id} | Delete OAuth 2.0 Client using the OpenID Dynamic Client Registration Management Protocol
*OidcApi* | [**discoverOidcConfiguration**](docs/Api/OidcApi.md#discoveroidcconfiguration) | **GET** /.well-known/openid-configuration | OpenID Connect Discovery
*OidcApi* | [**getOidcDynamicClient**](docs/Api/OidcApi.md#getoidcdynamicclient) | **GET** /oauth2/register/{id} | Get OAuth2 Client using OpenID Dynamic Client Registration
*OidcApi* | [**getOidcUserInfo**](docs/Api/OidcApi.md#getoidcuserinfo) | **GET** /userinfo | OpenID Connect Userinfo
*OidcApi* | [**revokeOidcSession**](docs/Api/OidcApi.md#revokeoidcsession) | **GET** /oauth2/sessions/logout | OpenID Connect Front- and Back-channel Enabled Logout
*OidcApi* | [**setOidcDynamicClient**](docs/Api/OidcApi.md#setoidcdynamicclient) | **PUT** /oauth2/register/{id} | Set OAuth2 Client using OpenID Dynamic Client Registration
*ReadApi* | [**getCheck**](docs/Api/ReadApi.md#getcheck) | **GET** /relation-tuples/check/openapi | Check a relation tuple
*ReadApi* | [**getCheckMirrorStatus**](docs/Api/ReadApi.md#getcheckmirrorstatus) | **GET** /relation-tuples/check | Check a relation tuple
*ReadApi* | [**getExpand**](docs/Api/ReadApi.md#getexpand) | **GET** /relation-tuples/expand | Expand a Relation Tuple
*ReadApi* | [**getRelationTuples**](docs/Api/ReadApi.md#getrelationtuples) | **GET** /relation-tuples | Query relation tuples
*ReadApi* | [**postCheck**](docs/Api/ReadApi.md#postcheck) | **POST** /relation-tuples/check/openapi | Check a relation tuple
*ReadApi* | [**postCheckMirrorStatus**](docs/Api/ReadApi.md#postcheckmirrorstatus) | **POST** /relation-tuples/check | Check a relation tuple
*SyntaxApi* | [**checkOplSyntax**](docs/Api/SyntaxApi.md#checkoplsyntax) | **POST** /opl/syntax/check | Check the syntax of an OPL file
*V0alpha2Api* | [**adminCreateIdentity**](docs/Api/V0alpha2Api.md#admincreateidentity) | **POST** /admin/identities | Create an Identity
*V0alpha2Api* | [**adminCreateSelfServiceRecoveryCode**](docs/Api/V0alpha2Api.md#admincreateselfservicerecoverycode) | **POST** /admin/recovery/code | Create a Recovery Code
*V0alpha2Api* | [**adminCreateSelfServiceRecoveryLink**](docs/Api/V0alpha2Api.md#admincreateselfservicerecoverylink) | **POST** /admin/recovery/link | Create a Recovery Link
*V0alpha2Api* | [**adminDeleteIdentity**](docs/Api/V0alpha2Api.md#admindeleteidentity) | **DELETE** /admin/identities/{id} | Delete an Identity
*V0alpha2Api* | [**adminDeleteIdentitySessions**](docs/Api/V0alpha2Api.md#admindeleteidentitysessions) | **DELETE** /admin/identities/{id}/sessions | Delete &amp; Invalidate an Identity&#39;s Sessions
*V0alpha2Api* | [**adminExtendSession**](docs/Api/V0alpha2Api.md#adminextendsession) | **PATCH** /admin/sessions/{id}/extend | Extend a Session
*V0alpha2Api* | [**adminGetIdentity**](docs/Api/V0alpha2Api.md#admingetidentity) | **GET** /admin/identities/{id} | Get an Identity
*V0alpha2Api* | [**adminGetSession**](docs/Api/V0alpha2Api.md#admingetsession) | **GET** /admin/sessions/{id} | This endpoint returns the session object with expandables specified.
*V0alpha2Api* | [**adminListCourierMessages**](docs/Api/V0alpha2Api.md#adminlistcouriermessages) | **GET** /admin/courier/messages | List Messages
*V0alpha2Api* | [**adminListIdentities**](docs/Api/V0alpha2Api.md#adminlistidentities) | **GET** /admin/identities | List Identities
*V0alpha2Api* | [**adminListIdentitySessions**](docs/Api/V0alpha2Api.md#adminlistidentitysessions) | **GET** /admin/identities/{id}/sessions | List an Identity&#39;s Sessions
*V0alpha2Api* | [**adminListSessions**](docs/Api/V0alpha2Api.md#adminlistsessions) | **GET** /admin/sessions | This endpoint returns all sessions that exist.
*V0alpha2Api* | [**adminPatchIdentity**](docs/Api/V0alpha2Api.md#adminpatchidentity) | **PATCH** /admin/identities/{id} | Patch an Identity
*V0alpha2Api* | [**adminUpdateIdentity**](docs/Api/V0alpha2Api.md#adminupdateidentity) | **PUT** /admin/identities/{id} | Update an Identity
*V0alpha2Api* | [**createProject**](docs/Api/V0alpha2Api.md#createproject) | **POST** /projects | Create a Project
*V0alpha2Api* | [**createProjectApiKey**](docs/Api/V0alpha2Api.md#createprojectapikey) | **POST** /projects/{project}/tokens | Create API Token
*V0alpha2Api* | [**createSelfServiceLogoutFlowUrlForBrowsers**](docs/Api/V0alpha2Api.md#createselfservicelogoutflowurlforbrowsers) | **GET** /self-service/logout/browser | Create a Logout URL for Browsers
*V0alpha2Api* | [**deleteProjectApiKey**](docs/Api/V0alpha2Api.md#deleteprojectapikey) | **DELETE** /projects/{project}/tokens/{token_id} | Delete API Token
*V0alpha2Api* | [**getIdentitySchema**](docs/Api/V0alpha2Api.md#getidentityschema) | **GET** /schemas/{id} | 
*V0alpha2Api* | [**getProject**](docs/Api/V0alpha2Api.md#getproject) | **GET** /projects/{project_id} | Get a Project
*V0alpha2Api* | [**getProjectMembers**](docs/Api/V0alpha2Api.md#getprojectmembers) | **GET** /projects/{project_id}/members | Get all members associated with this project.
*V0alpha2Api* | [**getSelfServiceError**](docs/Api/V0alpha2Api.md#getselfserviceerror) | **GET** /self-service/errors | Get Self-Service Errors
*V0alpha2Api* | [**getSelfServiceLoginFlow**](docs/Api/V0alpha2Api.md#getselfserviceloginflow) | **GET** /self-service/login/flows | Get Login Flow
*V0alpha2Api* | [**getSelfServiceRecoveryFlow**](docs/Api/V0alpha2Api.md#getselfservicerecoveryflow) | **GET** /self-service/recovery/flows | Get Recovery Flow
*V0alpha2Api* | [**getSelfServiceRegistrationFlow**](docs/Api/V0alpha2Api.md#getselfserviceregistrationflow) | **GET** /self-service/registration/flows | Get Registration Flow
*V0alpha2Api* | [**getSelfServiceSettingsFlow**](docs/Api/V0alpha2Api.md#getselfservicesettingsflow) | **GET** /self-service/settings/flows | Get Settings Flow
*V0alpha2Api* | [**getSelfServiceVerificationFlow**](docs/Api/V0alpha2Api.md#getselfserviceverificationflow) | **GET** /self-service/verification/flows | Get Verification Flow
*V0alpha2Api* | [**getWebAuthnJavaScript**](docs/Api/V0alpha2Api.md#getwebauthnjavascript) | **GET** /.well-known/ory/webauthn.js | Get WebAuthn JavaScript
*V0alpha2Api* | [**initializeSelfServiceLoginFlowForBrowsers**](docs/Api/V0alpha2Api.md#initializeselfserviceloginflowforbrowsers) | **GET** /self-service/login/browser | Initialize Login Flow for Browsers
*V0alpha2Api* | [**initializeSelfServiceLoginFlowWithoutBrowser**](docs/Api/V0alpha2Api.md#initializeselfserviceloginflowwithoutbrowser) | **GET** /self-service/login/api | Initialize Login Flow for APIs, Services, Apps, ...
*V0alpha2Api* | [**initializeSelfServiceRecoveryFlowForBrowsers**](docs/Api/V0alpha2Api.md#initializeselfservicerecoveryflowforbrowsers) | **GET** /self-service/recovery/browser | Initialize Recovery Flow for Browsers
*V0alpha2Api* | [**initializeSelfServiceRecoveryFlowWithoutBrowser**](docs/Api/V0alpha2Api.md#initializeselfservicerecoveryflowwithoutbrowser) | **GET** /self-service/recovery/api | Initialize Recovery Flow for APIs, Services, Apps, ...
*V0alpha2Api* | [**initializeSelfServiceRegistrationFlowForBrowsers**](docs/Api/V0alpha2Api.md#initializeselfserviceregistrationflowforbrowsers) | **GET** /self-service/registration/browser | Initialize Registration Flow for Browsers
*V0alpha2Api* | [**initializeSelfServiceRegistrationFlowWithoutBrowser**](docs/Api/V0alpha2Api.md#initializeselfserviceregistrationflowwithoutbrowser) | **GET** /self-service/registration/api | Initialize Registration Flow for APIs, Services, Apps, ...
*V0alpha2Api* | [**initializeSelfServiceSettingsFlowForBrowsers**](docs/Api/V0alpha2Api.md#initializeselfservicesettingsflowforbrowsers) | **GET** /self-service/settings/browser | Initialize Settings Flow for Browsers
*V0alpha2Api* | [**initializeSelfServiceSettingsFlowWithoutBrowser**](docs/Api/V0alpha2Api.md#initializeselfservicesettingsflowwithoutbrowser) | **GET** /self-service/settings/api | Initialize Settings Flow for APIs, Services, Apps, ...
*V0alpha2Api* | [**initializeSelfServiceVerificationFlowForBrowsers**](docs/Api/V0alpha2Api.md#initializeselfserviceverificationflowforbrowsers) | **GET** /self-service/verification/browser | Initialize Verification Flow for Browser Clients
*V0alpha2Api* | [**initializeSelfServiceVerificationFlowWithoutBrowser**](docs/Api/V0alpha2Api.md#initializeselfserviceverificationflowwithoutbrowser) | **GET** /self-service/verification/api | Initialize Verification Flow for APIs, Services, Apps, ...
*V0alpha2Api* | [**listIdentitySchemas**](docs/Api/V0alpha2Api.md#listidentityschemas) | **GET** /schemas | 
*V0alpha2Api* | [**listProjectApiKeys**](docs/Api/V0alpha2Api.md#listprojectapikeys) | **GET** /projects/{project}/tokens | List a Project&#39;s API Tokens
*V0alpha2Api* | [**listProjects**](docs/Api/V0alpha2Api.md#listprojects) | **GET** /projects | List All Projects
*V0alpha2Api* | [**listSessions**](docs/Api/V0alpha2Api.md#listsessions) | **GET** /sessions | Get Active Sessions
*V0alpha2Api* | [**patchProject**](docs/Api/V0alpha2Api.md#patchproject) | **PATCH** /projects/{project_id} | Patch an Ory Cloud Project Configuration&#x60;
*V0alpha2Api* | [**purgeProject**](docs/Api/V0alpha2Api.md#purgeproject) | **DELETE** /projects/{project_id} | Irrecoverably Purge a Project
*V0alpha2Api* | [**removeProjectMember**](docs/Api/V0alpha2Api.md#removeprojectmember) | **DELETE** /projects/{project_id}/members/{member_id} | Remove a member associated with this project. This also sets their invite status to &#x60;REMOVED&#x60;.
*V0alpha2Api* | [**revokeSession**](docs/Api/V0alpha2Api.md#revokesession) | **DELETE** /sessions/{id} | Invalidate a Session
*V0alpha2Api* | [**revokeSessions**](docs/Api/V0alpha2Api.md#revokesessions) | **DELETE** /sessions | Invalidate all Other Sessions
*V0alpha2Api* | [**submitSelfServiceLoginFlow**](docs/Api/V0alpha2Api.md#submitselfserviceloginflow) | **POST** /self-service/login | Submit a Login Flow
*V0alpha2Api* | [**submitSelfServiceLogoutFlow**](docs/Api/V0alpha2Api.md#submitselfservicelogoutflow) | **GET** /self-service/logout | Complete Self-Service Logout
*V0alpha2Api* | [**submitSelfServiceLogoutFlowWithoutBrowser**](docs/Api/V0alpha2Api.md#submitselfservicelogoutflowwithoutbrowser) | **DELETE** /self-service/logout/api | Perform Logout for APIs, Services, Apps, ...
*V0alpha2Api* | [**submitSelfServiceRecoveryFlow**](docs/Api/V0alpha2Api.md#submitselfservicerecoveryflow) | **POST** /self-service/recovery | Complete Recovery Flow
*V0alpha2Api* | [**submitSelfServiceRegistrationFlow**](docs/Api/V0alpha2Api.md#submitselfserviceregistrationflow) | **POST** /self-service/registration | Submit a Registration Flow
*V0alpha2Api* | [**submitSelfServiceSettingsFlow**](docs/Api/V0alpha2Api.md#submitselfservicesettingsflow) | **POST** /self-service/settings | Complete Settings Flow
*V0alpha2Api* | [**submitSelfServiceVerificationFlow**](docs/Api/V0alpha2Api.md#submitselfserviceverificationflow) | **POST** /self-service/verification | Complete Verification Flow
*V0alpha2Api* | [**toSession**](docs/Api/V0alpha2Api.md#tosession) | **GET** /sessions/whoami | Check Who the Current HTTP Session Belongs To
*V0alpha2Api* | [**updateProject**](docs/Api/V0alpha2Api.md#updateproject) | **PUT** /projects/{project_id} | Update an Ory Cloud Project Configuration
*WellknownApi* | [**discoverJsonWebKeys**](docs/Api/WellknownApi.md#discoverjsonwebkeys) | **GET** /.well-known/jwks.json | Discover Well-Known JSON Web Keys
*WriteApi* | [**createRelationTuple**](docs/Api/WriteApi.md#createrelationtuple) | **PUT** /admin/relation-tuples | Create a Relation Tuple
*WriteApi* | [**deleteRelationTuples**](docs/Api/WriteApi.md#deleterelationtuples) | **DELETE** /admin/relation-tuples | Delete Relation Tuples
*WriteApi* | [**patchRelationTuples**](docs/Api/WriteApi.md#patchrelationtuples) | **PATCH** /admin/relation-tuples | Patch Multiple Relation Tuples

## Models

- [AcceptOAuth2ConsentRequest](docs/Model/AcceptOAuth2ConsentRequest.md)
- [AcceptOAuth2ConsentRequestSession](docs/Model/AcceptOAuth2ConsentRequestSession.md)
- [AcceptOAuth2LoginRequest](docs/Model/AcceptOAuth2LoginRequest.md)
- [ActiveProject](docs/Model/ActiveProject.md)
- [AdminCreateIdentityBody](docs/Model/AdminCreateIdentityBody.md)
- [AdminCreateIdentityImportCredentialsOidc](docs/Model/AdminCreateIdentityImportCredentialsOidc.md)
- [AdminCreateIdentityImportCredentialsOidcConfig](docs/Model/AdminCreateIdentityImportCredentialsOidcConfig.md)
- [AdminCreateIdentityImportCredentialsOidcProvider](docs/Model/AdminCreateIdentityImportCredentialsOidcProvider.md)
- [AdminCreateIdentityImportCredentialsPassword](docs/Model/AdminCreateIdentityImportCredentialsPassword.md)
- [AdminCreateIdentityImportCredentialsPasswordConfig](docs/Model/AdminCreateIdentityImportCredentialsPasswordConfig.md)
- [AdminCreateSelfServiceRecoveryCodeBody](docs/Model/AdminCreateSelfServiceRecoveryCodeBody.md)
- [AdminCreateSelfServiceRecoveryLinkBody](docs/Model/AdminCreateSelfServiceRecoveryLinkBody.md)
- [AdminIdentityImportCredentials](docs/Model/AdminIdentityImportCredentials.md)
- [AdminUpdateIdentityBody](docs/Model/AdminUpdateIdentityBody.md)
- [AuthenticatorAssuranceLevel](docs/Model/AuthenticatorAssuranceLevel.md)
- [CloudAccount](docs/Model/CloudAccount.md)
- [CnameSettings](docs/Model/CnameSettings.md)
- [CourierMessageStatus](docs/Model/CourierMessageStatus.md)
- [CourierMessageType](docs/Model/CourierMessageType.md)
- [CreateCustomHostnameBody](docs/Model/CreateCustomHostnameBody.md)
- [CreateJsonWebKeySet](docs/Model/CreateJsonWebKeySet.md)
- [CreateProjectBody](docs/Model/CreateProjectBody.md)
- [CreateSubscriptionPayload](docs/Model/CreateSubscriptionPayload.md)
- [ErrorAuthenticatorAssuranceLevelNotSatisfied](docs/Model/ErrorAuthenticatorAssuranceLevelNotSatisfied.md)
- [ErrorOAuth2](docs/Model/ErrorOAuth2.md)
- [ExpandTree](docs/Model/ExpandTree.md)
- [GenericError](docs/Model/GenericError.md)
- [GenericErrorContent](docs/Model/GenericErrorContent.md)
- [GetCheckResponse](docs/Model/GetCheckResponse.md)
- [GetManagedIdentitySchemaLocation](docs/Model/GetManagedIdentitySchemaLocation.md)
- [GetNamespacesResponse](docs/Model/GetNamespacesResponse.md)
- [GetRelationTuplesResponse](docs/Model/GetRelationTuplesResponse.md)
- [HealthNotReadyStatus](docs/Model/HealthNotReadyStatus.md)
- [HealthStatus](docs/Model/HealthStatus.md)
- [Identity](docs/Model/Identity.md)
- [IdentityCredentials](docs/Model/IdentityCredentials.md)
- [IdentityCredentialsOidc](docs/Model/IdentityCredentialsOidc.md)
- [IdentityCredentialsOidcProvider](docs/Model/IdentityCredentialsOidcProvider.md)
- [IdentityCredentialsPassword](docs/Model/IdentityCredentialsPassword.md)
- [IdentityCredentialsType](docs/Model/IdentityCredentialsType.md)
- [IdentitySchemaContainer](docs/Model/IdentitySchemaContainer.md)
- [IdentitySchemaPreset](docs/Model/IdentitySchemaPreset.md)
- [IdentityState](docs/Model/IdentityState.md)
- [InlineObject3](docs/Model/InlineObject3.md)
- [InlineResponse200](docs/Model/InlineResponse200.md)
- [InlineResponse2001](docs/Model/InlineResponse2001.md)
- [InlineResponse503](docs/Model/InlineResponse503.md)
- [IntrospectedOAuth2Token](docs/Model/IntrospectedOAuth2Token.md)
- [InvitePayload](docs/Model/InvitePayload.md)
- [IsOwnerForProjectBySlug](docs/Model/IsOwnerForProjectBySlug.md)
- [IsOwnerForProjectBySlugPayload](docs/Model/IsOwnerForProjectBySlugPayload.md)
- [JsonError](docs/Model/JsonError.md)
- [JsonPatch](docs/Model/JsonPatch.md)
- [JsonWebKey](docs/Model/JsonWebKey.md)
- [JsonWebKeySet](docs/Model/JsonWebKeySet.md)
- [KetoNamespace](docs/Model/KetoNamespace.md)
- [LoginRequest](docs/Model/LoginRequest.md)
- [ManagedIdentitySchema](docs/Model/ManagedIdentitySchema.md)
- [ManagedIdentitySchemaValidationResult](docs/Model/ManagedIdentitySchemaValidationResult.md)
- [Message](docs/Model/Message.md)
- [ModelNamespace](docs/Model/ModelNamespace.md)
- [NeedsPrivilegedSessionError](docs/Model/NeedsPrivilegedSessionError.md)
- [NormalizedProject](docs/Model/NormalizedProject.md)
- [NormalizedProjectRevision](docs/Model/NormalizedProjectRevision.md)
- [NormalizedProjectRevisionHook](docs/Model/NormalizedProjectRevisionHook.md)
- [NormalizedProjectRevisionIdentitySchema](docs/Model/NormalizedProjectRevisionIdentitySchema.md)
- [NormalizedProjectRevisionThirdPartyProvider](docs/Model/NormalizedProjectRevisionThirdPartyProvider.md)
- [NullPlan](docs/Model/NullPlan.md)
- [OAuth2Client](docs/Model/OAuth2Client.md)
- [OAuth2ClientTokenLifespans](docs/Model/OAuth2ClientTokenLifespans.md)
- [OAuth2ConsentRequest](docs/Model/OAuth2ConsentRequest.md)
- [OAuth2ConsentRequestOpenIDConnectContext](docs/Model/OAuth2ConsentRequestOpenIDConnectContext.md)
- [OAuth2ConsentSession](docs/Model/OAuth2ConsentSession.md)
- [OAuth2ConsentSessionExpiresAt](docs/Model/OAuth2ConsentSessionExpiresAt.md)
- [OAuth2LoginRequest](docs/Model/OAuth2LoginRequest.md)
- [OAuth2LogoutRequest](docs/Model/OAuth2LogoutRequest.md)
- [OAuth2RedirectTo](docs/Model/OAuth2RedirectTo.md)
- [OAuth2TokenExchange](docs/Model/OAuth2TokenExchange.md)
- [OidcConfiguration](docs/Model/OidcConfiguration.md)
- [OidcUserInfo](docs/Model/OidcUserInfo.md)
- [OpenIDConnectContext](docs/Model/OpenIDConnectContext.md)
- [Pagination](docs/Model/Pagination.md)
- [PaginationHeaders](docs/Model/PaginationHeaders.md)
- [ParseError](docs/Model/ParseError.md)
- [PatchDelta](docs/Model/PatchDelta.md)
- [PostCheckOplSyntaxResponse](docs/Model/PostCheckOplSyntaxResponse.md)
- [Project](docs/Model/Project.md)
- [ProjectApiKey](docs/Model/ProjectApiKey.md)
- [ProjectHost](docs/Model/ProjectHost.md)
- [ProjectInvite](docs/Model/ProjectInvite.md)
- [ProjectMetadata](docs/Model/ProjectMetadata.md)
- [ProjectServiceIdentity](docs/Model/ProjectServiceIdentity.md)
- [ProjectServiceOAuth2](docs/Model/ProjectServiceOAuth2.md)
- [ProjectServicePermission](docs/Model/ProjectServicePermission.md)
- [ProjectServices](docs/Model/ProjectServices.md)
- [ProvisionMockSubscriptionPayload](docs/Model/ProvisionMockSubscriptionPayload.md)
- [QuotaCustomDomains](docs/Model/QuotaCustomDomains.md)
- [QuotaProjectMemberSeats](docs/Model/QuotaProjectMemberSeats.md)
- [RecoveryIdentityAddress](docs/Model/RecoveryIdentityAddress.md)
- [RejectOAuth2Request](docs/Model/RejectOAuth2Request.md)
- [RelationQuery](docs/Model/RelationQuery.md)
- [RelationTuple](docs/Model/RelationTuple.md)
- [RevokedSessions](docs/Model/RevokedSessions.md)
- [SchemaPatch](docs/Model/SchemaPatch.md)
- [SelfServiceBrowserLocationChangeRequiredError](docs/Model/SelfServiceBrowserLocationChangeRequiredError.md)
- [SelfServiceError](docs/Model/SelfServiceError.md)
- [SelfServiceFlowExpiredError](docs/Model/SelfServiceFlowExpiredError.md)
- [SelfServiceLoginFlow](docs/Model/SelfServiceLoginFlow.md)
- [SelfServiceLogoutUrl](docs/Model/SelfServiceLogoutUrl.md)
- [SelfServiceRecoveryCode](docs/Model/SelfServiceRecoveryCode.md)
- [SelfServiceRecoveryFlow](docs/Model/SelfServiceRecoveryFlow.md)
- [SelfServiceRecoveryFlowState](docs/Model/SelfServiceRecoveryFlowState.md)
- [SelfServiceRecoveryLink](docs/Model/SelfServiceRecoveryLink.md)
- [SelfServiceRegistrationFlow](docs/Model/SelfServiceRegistrationFlow.md)
- [SelfServiceSettingsFlow](docs/Model/SelfServiceSettingsFlow.md)
- [SelfServiceSettingsFlowState](docs/Model/SelfServiceSettingsFlowState.md)
- [SelfServiceVerificationFlow](docs/Model/SelfServiceVerificationFlow.md)
- [SelfServiceVerificationFlowState](docs/Model/SelfServiceVerificationFlowState.md)
- [Session](docs/Model/Session.md)
- [SessionAuthenticationMethod](docs/Model/SessionAuthenticationMethod.md)
- [SessionDevice](docs/Model/SessionDevice.md)
- [SettingsProfileFormConfig](docs/Model/SettingsProfileFormConfig.md)
- [SourcePosition](docs/Model/SourcePosition.md)
- [StripeCustomerResponse](docs/Model/StripeCustomerResponse.md)
- [SubjectSet](docs/Model/SubjectSet.md)
- [SubmitSelfServiceFlowWithWebAuthnRegistrationMethod](docs/Model/SubmitSelfServiceFlowWithWebAuthnRegistrationMethod.md)
- [SubmitSelfServiceLoginFlowBody](docs/Model/SubmitSelfServiceLoginFlowBody.md)
- [SubmitSelfServiceLoginFlowWithLookupSecretMethodBody](docs/Model/SubmitSelfServiceLoginFlowWithLookupSecretMethodBody.md)
- [SubmitSelfServiceLoginFlowWithOidcMethodBody](docs/Model/SubmitSelfServiceLoginFlowWithOidcMethodBody.md)
- [SubmitSelfServiceLoginFlowWithPasswordMethodBody](docs/Model/SubmitSelfServiceLoginFlowWithPasswordMethodBody.md)
- [SubmitSelfServiceLoginFlowWithTotpMethodBody](docs/Model/SubmitSelfServiceLoginFlowWithTotpMethodBody.md)
- [SubmitSelfServiceLoginFlowWithWebAuthnMethodBody](docs/Model/SubmitSelfServiceLoginFlowWithWebAuthnMethodBody.md)
- [SubmitSelfServiceLogoutFlowWithoutBrowserBody](docs/Model/SubmitSelfServiceLogoutFlowWithoutBrowserBody.md)
- [SubmitSelfServiceRecoveryFlowBody](docs/Model/SubmitSelfServiceRecoveryFlowBody.md)
- [SubmitSelfServiceRecoveryFlowWithCodeMethodBody](docs/Model/SubmitSelfServiceRecoveryFlowWithCodeMethodBody.md)
- [SubmitSelfServiceRecoveryFlowWithLinkMethodBody](docs/Model/SubmitSelfServiceRecoveryFlowWithLinkMethodBody.md)
- [SubmitSelfServiceRegistrationFlowBody](docs/Model/SubmitSelfServiceRegistrationFlowBody.md)
- [SubmitSelfServiceRegistrationFlowWithOidcMethodBody](docs/Model/SubmitSelfServiceRegistrationFlowWithOidcMethodBody.md)
- [SubmitSelfServiceRegistrationFlowWithPasswordMethodBody](docs/Model/SubmitSelfServiceRegistrationFlowWithPasswordMethodBody.md)
- [SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody](docs/Model/SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody.md)
- [SubmitSelfServiceSettingsFlowBody](docs/Model/SubmitSelfServiceSettingsFlowBody.md)
- [SubmitSelfServiceSettingsFlowWithLookupMethodBody](docs/Model/SubmitSelfServiceSettingsFlowWithLookupMethodBody.md)
- [SubmitSelfServiceSettingsFlowWithOidcMethodBody](docs/Model/SubmitSelfServiceSettingsFlowWithOidcMethodBody.md)
- [SubmitSelfServiceSettingsFlowWithPasswordMethodBody](docs/Model/SubmitSelfServiceSettingsFlowWithPasswordMethodBody.md)
- [SubmitSelfServiceSettingsFlowWithProfileMethodBody](docs/Model/SubmitSelfServiceSettingsFlowWithProfileMethodBody.md)
- [SubmitSelfServiceSettingsFlowWithTotpMethodBody](docs/Model/SubmitSelfServiceSettingsFlowWithTotpMethodBody.md)
- [SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody](docs/Model/SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody.md)
- [SubmitSelfServiceVerificationFlowBody](docs/Model/SubmitSelfServiceVerificationFlowBody.md)
- [SubmitSelfServiceVerificationFlowWithLinkMethodBody](docs/Model/SubmitSelfServiceVerificationFlowWithLinkMethodBody.md)
- [Subscription](docs/Model/Subscription.md)
- [SuccessfulProjectUpdate](docs/Model/SuccessfulProjectUpdate.md)
- [SuccessfulSelfServiceLoginWithoutBrowser](docs/Model/SuccessfulSelfServiceLoginWithoutBrowser.md)
- [SuccessfulSelfServiceRegistrationWithoutBrowser](docs/Model/SuccessfulSelfServiceRegistrationWithoutBrowser.md)
- [TokenPagination](docs/Model/TokenPagination.md)
- [TokenPaginationHeaders](docs/Model/TokenPaginationHeaders.md)
- [TokenPaginationRequestParameters](docs/Model/TokenPaginationRequestParameters.md)
- [TokenPaginationResponseHeaders](docs/Model/TokenPaginationResponseHeaders.md)
- [TrustOAuth2JwtGrantIssuer](docs/Model/TrustOAuth2JwtGrantIssuer.md)
- [TrustedOAuth2JwtGrantIssuer](docs/Model/TrustedOAuth2JwtGrantIssuer.md)
- [TrustedOAuth2JwtGrantJsonWebKey](docs/Model/TrustedOAuth2JwtGrantJsonWebKey.md)
- [UiContainer](docs/Model/UiContainer.md)
- [UiNode](docs/Model/UiNode.md)
- [UiNodeAnchorAttributes](docs/Model/UiNodeAnchorAttributes.md)
- [UiNodeAttributes](docs/Model/UiNodeAttributes.md)
- [UiNodeImageAttributes](docs/Model/UiNodeImageAttributes.md)
- [UiNodeInputAttributes](docs/Model/UiNodeInputAttributes.md)
- [UiNodeMeta](docs/Model/UiNodeMeta.md)
- [UiNodeScriptAttributes](docs/Model/UiNodeScriptAttributes.md)
- [UiNodeTextAttributes](docs/Model/UiNodeTextAttributes.md)
- [UiText](docs/Model/UiText.md)
- [UpdateCustomHostnameBody](docs/Model/UpdateCustomHostnameBody.md)
- [UpdateProject](docs/Model/UpdateProject.md)
- [UpdateSubscriptionPayload](docs/Model/UpdateSubscriptionPayload.md)
- [VerifiableIdentityAddress](docs/Model/VerifiableIdentityAddress.md)
- [Version](docs/Model/Version.md)
- [Warning](docs/Model/Warning.md)

## Authorization

### basic

- **Type**: HTTP basic authentication


### bearer

- **Type**: Bearer authentication


### oauth2

- **Type**: `OAuth`
- **Flow**: `accessCode`
- **Authorization URL**: `https://hydra.demo.ory.sh/oauth2/auth`
- **Scopes**: 
    - **offline**: A scope required when requesting refresh tokens (alias for `offline_access`)
    - **offline_access**: A scope required when requesting refresh tokens
    - **openid**: Request an OpenID Connect ID Token


### oryAccessToken

- **Type**: Bearer authentication

## Tests

To run the tests, use:

```bash
composer install
vendor/bin/phpunit
```

## Author

support@ory.sh

## About this package

This PHP package is automatically generated by the [OpenAPI Generator](https://openapi-generator.tech) project:

- API version: `v0.3.0`
- Build package: `org.openapitools.codegen.languages.PhpClientCodegen`
