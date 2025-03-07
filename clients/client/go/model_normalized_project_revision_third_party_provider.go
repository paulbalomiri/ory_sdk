/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.60
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// NormalizedProjectRevisionThirdPartyProvider struct for NormalizedProjectRevisionThirdPartyProvider
type NormalizedProjectRevisionThirdPartyProvider struct {
	ApplePrivateKey NullableString `json:"apple_private_key,omitempty"`
	// Apple Private Key Identifier  Sign In with Apple Private Key Identifier needed for generating a JWT token for client secret
	ApplePrivateKeyId *string `json:"apple_private_key_id,omitempty"`
	// Apple Developer Team ID  Apple Developer Team ID needed for generating a JWT token for client secret
	AppleTeamId *string `json:"apple_team_id,omitempty"`
	// AuthURL is the authorize url, typically something like: https://example.org/oauth2/auth Should only be used when the OAuth2 / OpenID Connect server is not supporting OpenID Connect Discovery and when `provider` is set to `generic`.
	AuthUrl *string `json:"auth_url,omitempty"`
	// Tenant is the Azure AD Tenant to use for authentication, and must be set when `provider` is set to `microsoft`.  Can be either `common`, `organizations`, `consumers` for a multitenant application or a specific tenant like `8eaef023-2b34-4da1-9baa-8bc8c9d6a490` or `contoso.onmicrosoft.com`.
	AzureTenant *string `json:"azure_tenant,omitempty"`
	// ClientID is the application's Client ID.
	ClientId *string `json:"client_id,omitempty"`
	ClientSecret NullableString `json:"client_secret,omitempty"`
	// The Project's Revision Creation Date
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Id *string `json:"id,omitempty"`
	// IssuerURL is the OpenID Connect Server URL. You can leave this empty if `provider` is not set to `generic`. If set, neither `auth_url` nor `token_url` are required.
	IssuerUrl *string `json:"issuer_url,omitempty"`
	// Label represents an optional label which can be used in the UI generation.
	Label *string `json:"label,omitempty"`
	// Mapper specifies the JSONNet code snippet which uses the OpenID Connect Provider's data (e.g. GitHub or Google profile information) to hydrate the identity's data.  It can be either a URL (file://, http(s)://, base64://) or an inline JSONNet code snippet.
	MapperUrl *string `json:"mapper_url,omitempty"`
	// The Revision's ID this schema belongs to
	ProjectRevisionId *string `json:"project_revision_id,omitempty"`
	// Provider is either \"generic\" for a generic OAuth 2.0 / OpenID Connect Provider or one of: generic google github gitlab microsoft discord slack facebook vk yandex apple
	Provider *string `json:"provider,omitempty"`
	// ID is the provider's ID
	ProviderId *string `json:"provider_id,omitempty"`
	RequestedClaims map[string]interface{} `json:"requested_claims,omitempty"`
	Scope []string `json:"scope,omitempty"`
	// TokenURL is the token url, typically something like: https://example.org/oauth2/token  Should only be used when the OAuth2 / OpenID Connect server is not supporting OpenID Connect Discovery and when `provider` is set to `generic`.
	TokenUrl *string `json:"token_url,omitempty"`
	// Last Time Project's Revision was Updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewNormalizedProjectRevisionThirdPartyProvider instantiates a new NormalizedProjectRevisionThirdPartyProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNormalizedProjectRevisionThirdPartyProvider() *NormalizedProjectRevisionThirdPartyProvider {
	this := NormalizedProjectRevisionThirdPartyProvider{}
	return &this
}

// NewNormalizedProjectRevisionThirdPartyProviderWithDefaults instantiates a new NormalizedProjectRevisionThirdPartyProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNormalizedProjectRevisionThirdPartyProviderWithDefaults() *NormalizedProjectRevisionThirdPartyProvider {
	this := NormalizedProjectRevisionThirdPartyProvider{}
	return &this
}

// GetApplePrivateKey returns the ApplePrivateKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NormalizedProjectRevisionThirdPartyProvider) GetApplePrivateKey() string {
	if o == nil || o.ApplePrivateKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.ApplePrivateKey.Get()
}

// GetApplePrivateKeyOk returns a tuple with the ApplePrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NormalizedProjectRevisionThirdPartyProvider) GetApplePrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplePrivateKey.Get(), o.ApplePrivateKey.IsSet()
}

// HasApplePrivateKey returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) HasApplePrivateKey() bool {
	if o != nil && o.ApplePrivateKey.IsSet() {
		return true
	}

	return false
}

// SetApplePrivateKey gets a reference to the given NullableString and assigns it to the ApplePrivateKey field.
func (o *NormalizedProjectRevisionThirdPartyProvider) SetApplePrivateKey(v string) {
	o.ApplePrivateKey.Set(&v)
}
// SetApplePrivateKeyNil sets the value for ApplePrivateKey to be an explicit nil
func (o *NormalizedProjectRevisionThirdPartyProvider) SetApplePrivateKeyNil() {
	o.ApplePrivateKey.Set(nil)
}

// UnsetApplePrivateKey ensures that no value is present for ApplePrivateKey, not even an explicit nil
func (o *NormalizedProjectRevisionThirdPartyProvider) UnsetApplePrivateKey() {
	o.ApplePrivateKey.Unset()
}

// GetApplePrivateKeyId returns the ApplePrivateKeyId field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetApplePrivateKeyId() string {
	if o == nil || o.ApplePrivateKeyId == nil {
		var ret string
		return ret
	}
	return *o.ApplePrivateKeyId
}

// GetApplePrivateKeyIdOk returns a tuple with the ApplePrivateKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetApplePrivateKeyIdOk() (*string, bool) {
	if o == nil || o.ApplePrivateKeyId == nil {
		return nil, false
	}
	return o.ApplePrivateKeyId, true
}

// HasApplePrivateKeyId returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) HasApplePrivateKeyId() bool {
	if o != nil && o.ApplePrivateKeyId != nil {
		return true
	}

	return false
}

// SetApplePrivateKeyId gets a reference to the given string and assigns it to the ApplePrivateKeyId field.
func (o *NormalizedProjectRevisionThirdPartyProvider) SetApplePrivateKeyId(v string) {
	o.ApplePrivateKeyId = &v
}

// GetAppleTeamId returns the AppleTeamId field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetAppleTeamId() string {
	if o == nil || o.AppleTeamId == nil {
		var ret string
		return ret
	}
	return *o.AppleTeamId
}

// GetAppleTeamIdOk returns a tuple with the AppleTeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetAppleTeamIdOk() (*string, bool) {
	if o == nil || o.AppleTeamId == nil {
		return nil, false
	}
	return o.AppleTeamId, true
}

// HasAppleTeamId returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) HasAppleTeamId() bool {
	if o != nil && o.AppleTeamId != nil {
		return true
	}

	return false
}

// SetAppleTeamId gets a reference to the given string and assigns it to the AppleTeamId field.
func (o *NormalizedProjectRevisionThirdPartyProvider) SetAppleTeamId(v string) {
	o.AppleTeamId = &v
}

// GetAuthUrl returns the AuthUrl field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetAuthUrl() string {
	if o == nil || o.AuthUrl == nil {
		var ret string
		return ret
	}
	return *o.AuthUrl
}

// GetAuthUrlOk returns a tuple with the AuthUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetAuthUrlOk() (*string, bool) {
	if o == nil || o.AuthUrl == nil {
		return nil, false
	}
	return o.AuthUrl, true
}

// HasAuthUrl returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) HasAuthUrl() bool {
	if o != nil && o.AuthUrl != nil {
		return true
	}

	return false
}

// SetAuthUrl gets a reference to the given string and assigns it to the AuthUrl field.
func (o *NormalizedProjectRevisionThirdPartyProvider) SetAuthUrl(v string) {
	o.AuthUrl = &v
}

// GetAzureTenant returns the AzureTenant field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetAzureTenant() string {
	if o == nil || o.AzureTenant == nil {
		var ret string
		return ret
	}
	return *o.AzureTenant
}

// GetAzureTenantOk returns a tuple with the AzureTenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetAzureTenantOk() (*string, bool) {
	if o == nil || o.AzureTenant == nil {
		return nil, false
	}
	return o.AzureTenant, true
}

// HasAzureTenant returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) HasAzureTenant() bool {
	if o != nil && o.AzureTenant != nil {
		return true
	}

	return false
}

// SetAzureTenant gets a reference to the given string and assigns it to the AzureTenant field.
func (o *NormalizedProjectRevisionThirdPartyProvider) SetAzureTenant(v string) {
	o.AzureTenant = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *NormalizedProjectRevisionThirdPartyProvider) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NormalizedProjectRevisionThirdPartyProvider) GetClientSecret() string {
	if o == nil || o.ClientSecret.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret.Get()
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NormalizedProjectRevisionThirdPartyProvider) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientSecret.Get(), o.ClientSecret.IsSet()
}

// HasClientSecret returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) HasClientSecret() bool {
	if o != nil && o.ClientSecret.IsSet() {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given NullableString and assigns it to the ClientSecret field.
func (o *NormalizedProjectRevisionThirdPartyProvider) SetClientSecret(v string) {
	o.ClientSecret.Set(&v)
}
// SetClientSecretNil sets the value for ClientSecret to be an explicit nil
func (o *NormalizedProjectRevisionThirdPartyProvider) SetClientSecretNil() {
	o.ClientSecret.Set(nil)
}

// UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
func (o *NormalizedProjectRevisionThirdPartyProvider) UnsetClientSecret() {
	o.ClientSecret.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *NormalizedProjectRevisionThirdPartyProvider) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NormalizedProjectRevisionThirdPartyProvider) SetId(v string) {
	o.Id = &v
}

// GetIssuerUrl returns the IssuerUrl field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetIssuerUrl() string {
	if o == nil || o.IssuerUrl == nil {
		var ret string
		return ret
	}
	return *o.IssuerUrl
}

// GetIssuerUrlOk returns a tuple with the IssuerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetIssuerUrlOk() (*string, bool) {
	if o == nil || o.IssuerUrl == nil {
		return nil, false
	}
	return o.IssuerUrl, true
}

// HasIssuerUrl returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) HasIssuerUrl() bool {
	if o != nil && o.IssuerUrl != nil {
		return true
	}

	return false
}

// SetIssuerUrl gets a reference to the given string and assigns it to the IssuerUrl field.
func (o *NormalizedProjectRevisionThirdPartyProvider) SetIssuerUrl(v string) {
	o.IssuerUrl = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *NormalizedProjectRevisionThirdPartyProvider) SetLabel(v string) {
	o.Label = &v
}

// GetMapperUrl returns the MapperUrl field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetMapperUrl() string {
	if o == nil || o.MapperUrl == nil {
		var ret string
		return ret
	}
	return *o.MapperUrl
}

// GetMapperUrlOk returns a tuple with the MapperUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetMapperUrlOk() (*string, bool) {
	if o == nil || o.MapperUrl == nil {
		return nil, false
	}
	return o.MapperUrl, true
}

// HasMapperUrl returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) HasMapperUrl() bool {
	if o != nil && o.MapperUrl != nil {
		return true
	}

	return false
}

// SetMapperUrl gets a reference to the given string and assigns it to the MapperUrl field.
func (o *NormalizedProjectRevisionThirdPartyProvider) SetMapperUrl(v string) {
	o.MapperUrl = &v
}

// GetProjectRevisionId returns the ProjectRevisionId field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetProjectRevisionId() string {
	if o == nil || o.ProjectRevisionId == nil {
		var ret string
		return ret
	}
	return *o.ProjectRevisionId
}

// GetProjectRevisionIdOk returns a tuple with the ProjectRevisionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetProjectRevisionIdOk() (*string, bool) {
	if o == nil || o.ProjectRevisionId == nil {
		return nil, false
	}
	return o.ProjectRevisionId, true
}

// HasProjectRevisionId returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) HasProjectRevisionId() bool {
	if o != nil && o.ProjectRevisionId != nil {
		return true
	}

	return false
}

// SetProjectRevisionId gets a reference to the given string and assigns it to the ProjectRevisionId field.
func (o *NormalizedProjectRevisionThirdPartyProvider) SetProjectRevisionId(v string) {
	o.ProjectRevisionId = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *NormalizedProjectRevisionThirdPartyProvider) SetProvider(v string) {
	o.Provider = &v
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetProviderId() string {
	if o == nil || o.ProviderId == nil {
		var ret string
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetProviderIdOk() (*string, bool) {
	if o == nil || o.ProviderId == nil {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) HasProviderId() bool {
	if o != nil && o.ProviderId != nil {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given string and assigns it to the ProviderId field.
func (o *NormalizedProjectRevisionThirdPartyProvider) SetProviderId(v string) {
	o.ProviderId = &v
}

// GetRequestedClaims returns the RequestedClaims field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetRequestedClaims() map[string]interface{} {
	if o == nil || o.RequestedClaims == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.RequestedClaims
}

// GetRequestedClaimsOk returns a tuple with the RequestedClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetRequestedClaimsOk() (map[string]interface{}, bool) {
	if o == nil || o.RequestedClaims == nil {
		return nil, false
	}
	return o.RequestedClaims, true
}

// HasRequestedClaims returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) HasRequestedClaims() bool {
	if o != nil && o.RequestedClaims != nil {
		return true
	}

	return false
}

// SetRequestedClaims gets a reference to the given map[string]interface{} and assigns it to the RequestedClaims field.
func (o *NormalizedProjectRevisionThirdPartyProvider) SetRequestedClaims(v map[string]interface{}) {
	o.RequestedClaims = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetScope() []string {
	if o == nil || o.Scope == nil {
		var ret []string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetScopeOk() ([]string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given []string and assigns it to the Scope field.
func (o *NormalizedProjectRevisionThirdPartyProvider) SetScope(v []string) {
	o.Scope = v
}

// GetTokenUrl returns the TokenUrl field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetTokenUrl() string {
	if o == nil || o.TokenUrl == nil {
		var ret string
		return ret
	}
	return *o.TokenUrl
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetTokenUrlOk() (*string, bool) {
	if o == nil || o.TokenUrl == nil {
		return nil, false
	}
	return o.TokenUrl, true
}

// HasTokenUrl returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) HasTokenUrl() bool {
	if o != nil && o.TokenUrl != nil {
		return true
	}

	return false
}

// SetTokenUrl gets a reference to the given string and assigns it to the TokenUrl field.
func (o *NormalizedProjectRevisionThirdPartyProvider) SetTokenUrl(v string) {
	o.TokenUrl = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionThirdPartyProvider) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *NormalizedProjectRevisionThirdPartyProvider) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o NormalizedProjectRevisionThirdPartyProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplePrivateKey.IsSet() {
		toSerialize["apple_private_key"] = o.ApplePrivateKey.Get()
	}
	if o.ApplePrivateKeyId != nil {
		toSerialize["apple_private_key_id"] = o.ApplePrivateKeyId
	}
	if o.AppleTeamId != nil {
		toSerialize["apple_team_id"] = o.AppleTeamId
	}
	if o.AuthUrl != nil {
		toSerialize["auth_url"] = o.AuthUrl
	}
	if o.AzureTenant != nil {
		toSerialize["azure_tenant"] = o.AzureTenant
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientSecret.IsSet() {
		toSerialize["client_secret"] = o.ClientSecret.Get()
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IssuerUrl != nil {
		toSerialize["issuer_url"] = o.IssuerUrl
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.MapperUrl != nil {
		toSerialize["mapper_url"] = o.MapperUrl
	}
	if o.ProjectRevisionId != nil {
		toSerialize["project_revision_id"] = o.ProjectRevisionId
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.ProviderId != nil {
		toSerialize["provider_id"] = o.ProviderId
	}
	if o.RequestedClaims != nil {
		toSerialize["requested_claims"] = o.RequestedClaims
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.TokenUrl != nil {
		toSerialize["token_url"] = o.TokenUrl
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableNormalizedProjectRevisionThirdPartyProvider struct {
	value *NormalizedProjectRevisionThirdPartyProvider
	isSet bool
}

func (v NullableNormalizedProjectRevisionThirdPartyProvider) Get() *NormalizedProjectRevisionThirdPartyProvider {
	return v.value
}

func (v *NullableNormalizedProjectRevisionThirdPartyProvider) Set(val *NormalizedProjectRevisionThirdPartyProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableNormalizedProjectRevisionThirdPartyProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableNormalizedProjectRevisionThirdPartyProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNormalizedProjectRevisionThirdPartyProvider(val *NormalizedProjectRevisionThirdPartyProvider) *NullableNormalizedProjectRevisionThirdPartyProvider {
	return &NullableNormalizedProjectRevisionThirdPartyProvider{value: val, isSet: true}
}

func (v NullableNormalizedProjectRevisionThirdPartyProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNormalizedProjectRevisionThirdPartyProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


