/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * API version: v0.6.0-alpha.4
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SubmitSelfServiceRecoveryFlowWithLinkMethod struct for SubmitSelfServiceRecoveryFlowWithLinkMethod
type SubmitSelfServiceRecoveryFlowWithLinkMethod struct {
	// Sending the anti-csrf token is only required for browser login flows.
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Email to Recover  Needs to be set when initiating the flow. If the email is a registered recovery email, a recovery link will be sent. If the email is not known, a email with details on what happened will be sent instead.  format: email in: body
	Email *string `json:"email,omitempty"`
}

// NewSubmitSelfServiceRecoveryFlowWithLinkMethod instantiates a new SubmitSelfServiceRecoveryFlowWithLinkMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitSelfServiceRecoveryFlowWithLinkMethod() *SubmitSelfServiceRecoveryFlowWithLinkMethod {
	this := SubmitSelfServiceRecoveryFlowWithLinkMethod{}
	return &this
}

// NewSubmitSelfServiceRecoveryFlowWithLinkMethodWithDefaults instantiates a new SubmitSelfServiceRecoveryFlowWithLinkMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitSelfServiceRecoveryFlowWithLinkMethodWithDefaults() *SubmitSelfServiceRecoveryFlowWithLinkMethod {
	this := SubmitSelfServiceRecoveryFlowWithLinkMethod{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *SubmitSelfServiceRecoveryFlowWithLinkMethod) GetCsrfToken() string {
	if o == nil || o.CsrfToken == nil {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceRecoveryFlowWithLinkMethod) GetCsrfTokenOk() (*string, bool) {
	if o == nil || o.CsrfToken == nil {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *SubmitSelfServiceRecoveryFlowWithLinkMethod) HasCsrfToken() bool {
	if o != nil && o.CsrfToken != nil {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *SubmitSelfServiceRecoveryFlowWithLinkMethod) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SubmitSelfServiceRecoveryFlowWithLinkMethod) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceRecoveryFlowWithLinkMethod) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SubmitSelfServiceRecoveryFlowWithLinkMethod) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *SubmitSelfServiceRecoveryFlowWithLinkMethod) SetEmail(v string) {
	o.Email = &v
}

func (o SubmitSelfServiceRecoveryFlowWithLinkMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CsrfToken != nil {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableSubmitSelfServiceRecoveryFlowWithLinkMethod struct {
	value *SubmitSelfServiceRecoveryFlowWithLinkMethod
	isSet bool
}

func (v NullableSubmitSelfServiceRecoveryFlowWithLinkMethod) Get() *SubmitSelfServiceRecoveryFlowWithLinkMethod {
	return v.value
}

func (v *NullableSubmitSelfServiceRecoveryFlowWithLinkMethod) Set(val *SubmitSelfServiceRecoveryFlowWithLinkMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitSelfServiceRecoveryFlowWithLinkMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitSelfServiceRecoveryFlowWithLinkMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitSelfServiceRecoveryFlowWithLinkMethod(val *SubmitSelfServiceRecoveryFlowWithLinkMethod) *NullableSubmitSelfServiceRecoveryFlowWithLinkMethod {
	return &NullableSubmitSelfServiceRecoveryFlowWithLinkMethod{value: val, isSet: true}
}

func (v NullableSubmitSelfServiceRecoveryFlowWithLinkMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitSelfServiceRecoveryFlowWithLinkMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


