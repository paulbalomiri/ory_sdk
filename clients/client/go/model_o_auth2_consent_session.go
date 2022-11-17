/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.3.0
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// OAuth2ConsentSession A completed OAuth 2.0 Consent Session.
type OAuth2ConsentSession struct {
	ConsentRequest *OAuth2ConsentRequest `json:"consent_request,omitempty"`
	ExpiresAt *OAuth2ConsentSessionExpiresAt `json:"expires_at,omitempty"`
	GrantAccessTokenAudience []string `json:"grant_access_token_audience,omitempty"`
	GrantScope []string `json:"grant_scope,omitempty"`
	HandledAt *time.Time `json:"handled_at,omitempty"`
	// Remember Consent  Remember, if set to true, tells ORY Hydra to remember this consent authorization and reuse it if the same client asks the same user for the same, or a subset of, scope.
	Remember *bool `json:"remember,omitempty"`
	// Remember Consent For  RememberFor sets how long the consent authorization should be remembered for in seconds. If set to `0`, the authorization will be remembered indefinitely.
	RememberFor *int64 `json:"remember_for,omitempty"`
	Session *AcceptOAuth2ConsentRequestSession `json:"session,omitempty"`
}

// NewOAuth2ConsentSession instantiates a new OAuth2ConsentSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ConsentSession() *OAuth2ConsentSession {
	this := OAuth2ConsentSession{}
	return &this
}

// NewOAuth2ConsentSessionWithDefaults instantiates a new OAuth2ConsentSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ConsentSessionWithDefaults() *OAuth2ConsentSession {
	this := OAuth2ConsentSession{}
	return &this
}

// GetConsentRequest returns the ConsentRequest field value if set, zero value otherwise.
func (o *OAuth2ConsentSession) GetConsentRequest() OAuth2ConsentRequest {
	if o == nil || o.ConsentRequest == nil {
		var ret OAuth2ConsentRequest
		return ret
	}
	return *o.ConsentRequest
}

// GetConsentRequestOk returns a tuple with the ConsentRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ConsentSession) GetConsentRequestOk() (*OAuth2ConsentRequest, bool) {
	if o == nil || o.ConsentRequest == nil {
		return nil, false
	}
	return o.ConsentRequest, true
}

// HasConsentRequest returns a boolean if a field has been set.
func (o *OAuth2ConsentSession) HasConsentRequest() bool {
	if o != nil && o.ConsentRequest != nil {
		return true
	}

	return false
}

// SetConsentRequest gets a reference to the given OAuth2ConsentRequest and assigns it to the ConsentRequest field.
func (o *OAuth2ConsentSession) SetConsentRequest(v OAuth2ConsentRequest) {
	o.ConsentRequest = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *OAuth2ConsentSession) GetExpiresAt() OAuth2ConsentSessionExpiresAt {
	if o == nil || o.ExpiresAt == nil {
		var ret OAuth2ConsentSessionExpiresAt
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ConsentSession) GetExpiresAtOk() (*OAuth2ConsentSessionExpiresAt, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *OAuth2ConsentSession) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given OAuth2ConsentSessionExpiresAt and assigns it to the ExpiresAt field.
func (o *OAuth2ConsentSession) SetExpiresAt(v OAuth2ConsentSessionExpiresAt) {
	o.ExpiresAt = &v
}

// GetGrantAccessTokenAudience returns the GrantAccessTokenAudience field value if set, zero value otherwise.
func (o *OAuth2ConsentSession) GetGrantAccessTokenAudience() []string {
	if o == nil || o.GrantAccessTokenAudience == nil {
		var ret []string
		return ret
	}
	return o.GrantAccessTokenAudience
}

// GetGrantAccessTokenAudienceOk returns a tuple with the GrantAccessTokenAudience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ConsentSession) GetGrantAccessTokenAudienceOk() ([]string, bool) {
	if o == nil || o.GrantAccessTokenAudience == nil {
		return nil, false
	}
	return o.GrantAccessTokenAudience, true
}

// HasGrantAccessTokenAudience returns a boolean if a field has been set.
func (o *OAuth2ConsentSession) HasGrantAccessTokenAudience() bool {
	if o != nil && o.GrantAccessTokenAudience != nil {
		return true
	}

	return false
}

// SetGrantAccessTokenAudience gets a reference to the given []string and assigns it to the GrantAccessTokenAudience field.
func (o *OAuth2ConsentSession) SetGrantAccessTokenAudience(v []string) {
	o.GrantAccessTokenAudience = v
}

// GetGrantScope returns the GrantScope field value if set, zero value otherwise.
func (o *OAuth2ConsentSession) GetGrantScope() []string {
	if o == nil || o.GrantScope == nil {
		var ret []string
		return ret
	}
	return o.GrantScope
}

// GetGrantScopeOk returns a tuple with the GrantScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ConsentSession) GetGrantScopeOk() ([]string, bool) {
	if o == nil || o.GrantScope == nil {
		return nil, false
	}
	return o.GrantScope, true
}

// HasGrantScope returns a boolean if a field has been set.
func (o *OAuth2ConsentSession) HasGrantScope() bool {
	if o != nil && o.GrantScope != nil {
		return true
	}

	return false
}

// SetGrantScope gets a reference to the given []string and assigns it to the GrantScope field.
func (o *OAuth2ConsentSession) SetGrantScope(v []string) {
	o.GrantScope = v
}

// GetHandledAt returns the HandledAt field value if set, zero value otherwise.
func (o *OAuth2ConsentSession) GetHandledAt() time.Time {
	if o == nil || o.HandledAt == nil {
		var ret time.Time
		return ret
	}
	return *o.HandledAt
}

// GetHandledAtOk returns a tuple with the HandledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ConsentSession) GetHandledAtOk() (*time.Time, bool) {
	if o == nil || o.HandledAt == nil {
		return nil, false
	}
	return o.HandledAt, true
}

// HasHandledAt returns a boolean if a field has been set.
func (o *OAuth2ConsentSession) HasHandledAt() bool {
	if o != nil && o.HandledAt != nil {
		return true
	}

	return false
}

// SetHandledAt gets a reference to the given time.Time and assigns it to the HandledAt field.
func (o *OAuth2ConsentSession) SetHandledAt(v time.Time) {
	o.HandledAt = &v
}

// GetRemember returns the Remember field value if set, zero value otherwise.
func (o *OAuth2ConsentSession) GetRemember() bool {
	if o == nil || o.Remember == nil {
		var ret bool
		return ret
	}
	return *o.Remember
}

// GetRememberOk returns a tuple with the Remember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ConsentSession) GetRememberOk() (*bool, bool) {
	if o == nil || o.Remember == nil {
		return nil, false
	}
	return o.Remember, true
}

// HasRemember returns a boolean if a field has been set.
func (o *OAuth2ConsentSession) HasRemember() bool {
	if o != nil && o.Remember != nil {
		return true
	}

	return false
}

// SetRemember gets a reference to the given bool and assigns it to the Remember field.
func (o *OAuth2ConsentSession) SetRemember(v bool) {
	o.Remember = &v
}

// GetRememberFor returns the RememberFor field value if set, zero value otherwise.
func (o *OAuth2ConsentSession) GetRememberFor() int64 {
	if o == nil || o.RememberFor == nil {
		var ret int64
		return ret
	}
	return *o.RememberFor
}

// GetRememberForOk returns a tuple with the RememberFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ConsentSession) GetRememberForOk() (*int64, bool) {
	if o == nil || o.RememberFor == nil {
		return nil, false
	}
	return o.RememberFor, true
}

// HasRememberFor returns a boolean if a field has been set.
func (o *OAuth2ConsentSession) HasRememberFor() bool {
	if o != nil && o.RememberFor != nil {
		return true
	}

	return false
}

// SetRememberFor gets a reference to the given int64 and assigns it to the RememberFor field.
func (o *OAuth2ConsentSession) SetRememberFor(v int64) {
	o.RememberFor = &v
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *OAuth2ConsentSession) GetSession() AcceptOAuth2ConsentRequestSession {
	if o == nil || o.Session == nil {
		var ret AcceptOAuth2ConsentRequestSession
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ConsentSession) GetSessionOk() (*AcceptOAuth2ConsentRequestSession, bool) {
	if o == nil || o.Session == nil {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *OAuth2ConsentSession) HasSession() bool {
	if o != nil && o.Session != nil {
		return true
	}

	return false
}

// SetSession gets a reference to the given AcceptOAuth2ConsentRequestSession and assigns it to the Session field.
func (o *OAuth2ConsentSession) SetSession(v AcceptOAuth2ConsentRequestSession) {
	o.Session = &v
}

func (o OAuth2ConsentSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConsentRequest != nil {
		toSerialize["consent_request"] = o.ConsentRequest
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.GrantAccessTokenAudience != nil {
		toSerialize["grant_access_token_audience"] = o.GrantAccessTokenAudience
	}
	if o.GrantScope != nil {
		toSerialize["grant_scope"] = o.GrantScope
	}
	if o.HandledAt != nil {
		toSerialize["handled_at"] = o.HandledAt
	}
	if o.Remember != nil {
		toSerialize["remember"] = o.Remember
	}
	if o.RememberFor != nil {
		toSerialize["remember_for"] = o.RememberFor
	}
	if o.Session != nil {
		toSerialize["session"] = o.Session
	}
	return json.Marshal(toSerialize)
}

type NullableOAuth2ConsentSession struct {
	value *OAuth2ConsentSession
	isSet bool
}

func (v NullableOAuth2ConsentSession) Get() *OAuth2ConsentSession {
	return v.value
}

func (v *NullableOAuth2ConsentSession) Set(val *OAuth2ConsentSession) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ConsentSession) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ConsentSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ConsentSession(val *OAuth2ConsentSession) *NullableOAuth2ConsentSession {
	return &NullableOAuth2ConsentSession{value: val, isSet: true}
}

func (v NullableOAuth2ConsentSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ConsentSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


