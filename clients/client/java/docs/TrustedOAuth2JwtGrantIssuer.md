

# TrustedOAuth2JwtGrantIssuer


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**allowAnySubject** | **Boolean** | The \&quot;allow_any_subject\&quot; indicates that the issuer is allowed to have any principal as the subject of the JWT. |  [optional] |
|**createdAt** | **OffsetDateTime** | The \&quot;created_at\&quot; indicates, when grant was created. |  [optional] |
|**expiresAt** | **OffsetDateTime** | The \&quot;expires_at\&quot; indicates, when grant will expire, so we will reject assertion from \&quot;issuer\&quot; targeting \&quot;subject\&quot;. |  [optional] |
|**id** | **String** |  |  [optional] |
|**issuer** | **String** | The \&quot;issuer\&quot; identifies the principal that issued the JWT assertion (same as \&quot;iss\&quot; claim in JWT). |  [optional] |
|**publicKey** | [**TrustedOAuth2JwtGrantJsonWebKey**](TrustedOAuth2JwtGrantJsonWebKey.md) |  |  [optional] |
|**scope** | **List&lt;String&gt;** | The \&quot;scope\&quot; contains list of scope values (as described in Section 3.3 of OAuth 2.0 [RFC6749]) |  [optional] |
|**subject** | **String** | The \&quot;subject\&quot; identifies the principal that is the subject of the JWT. |  [optional] |



