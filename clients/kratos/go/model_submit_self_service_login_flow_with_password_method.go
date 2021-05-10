/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * API version: v0.6.0-alpha.17
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SubmitSelfServiceLoginFlowWithPasswordMethod struct for SubmitSelfServiceLoginFlowWithPasswordMethod
type SubmitSelfServiceLoginFlowWithPasswordMethod struct {
	// Sending the anti-csrf token is only required for browser login flows.
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Method should be set to \"password\" when logging in using the identifier and password strategy.
	Method *string `json:"method,omitempty"`
	// The user's password.
	Password *string `json:"password,omitempty"`
	// Identifier is the email or username of the user trying to log in.
	PasswordIdentifier *string `json:"password_identifier,omitempty"`
}

// NewSubmitSelfServiceLoginFlowWithPasswordMethod instantiates a new SubmitSelfServiceLoginFlowWithPasswordMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitSelfServiceLoginFlowWithPasswordMethod() *SubmitSelfServiceLoginFlowWithPasswordMethod {
	this := SubmitSelfServiceLoginFlowWithPasswordMethod{}
	return &this
}

// NewSubmitSelfServiceLoginFlowWithPasswordMethodWithDefaults instantiates a new SubmitSelfServiceLoginFlowWithPasswordMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitSelfServiceLoginFlowWithPasswordMethodWithDefaults() *SubmitSelfServiceLoginFlowWithPasswordMethod {
	this := SubmitSelfServiceLoginFlowWithPasswordMethod{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *SubmitSelfServiceLoginFlowWithPasswordMethod) GetCsrfToken() string {
	if o == nil || o.CsrfToken == nil {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceLoginFlowWithPasswordMethod) GetCsrfTokenOk() (*string, bool) {
	if o == nil || o.CsrfToken == nil {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *SubmitSelfServiceLoginFlowWithPasswordMethod) HasCsrfToken() bool {
	if o != nil && o.CsrfToken != nil {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *SubmitSelfServiceLoginFlowWithPasswordMethod) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *SubmitSelfServiceLoginFlowWithPasswordMethod) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceLoginFlowWithPasswordMethod) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *SubmitSelfServiceLoginFlowWithPasswordMethod) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *SubmitSelfServiceLoginFlowWithPasswordMethod) SetMethod(v string) {
	o.Method = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SubmitSelfServiceLoginFlowWithPasswordMethod) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceLoginFlowWithPasswordMethod) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SubmitSelfServiceLoginFlowWithPasswordMethod) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SubmitSelfServiceLoginFlowWithPasswordMethod) SetPassword(v string) {
	o.Password = &v
}

// GetPasswordIdentifier returns the PasswordIdentifier field value if set, zero value otherwise.
func (o *SubmitSelfServiceLoginFlowWithPasswordMethod) GetPasswordIdentifier() string {
	if o == nil || o.PasswordIdentifier == nil {
		var ret string
		return ret
	}
	return *o.PasswordIdentifier
}

// GetPasswordIdentifierOk returns a tuple with the PasswordIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceLoginFlowWithPasswordMethod) GetPasswordIdentifierOk() (*string, bool) {
	if o == nil || o.PasswordIdentifier == nil {
		return nil, false
	}
	return o.PasswordIdentifier, true
}

// HasPasswordIdentifier returns a boolean if a field has been set.
func (o *SubmitSelfServiceLoginFlowWithPasswordMethod) HasPasswordIdentifier() bool {
	if o != nil && o.PasswordIdentifier != nil {
		return true
	}

	return false
}

// SetPasswordIdentifier gets a reference to the given string and assigns it to the PasswordIdentifier field.
func (o *SubmitSelfServiceLoginFlowWithPasswordMethod) SetPasswordIdentifier(v string) {
	o.PasswordIdentifier = &v
}

func (o SubmitSelfServiceLoginFlowWithPasswordMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CsrfToken != nil {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.PasswordIdentifier != nil {
		toSerialize["password_identifier"] = o.PasswordIdentifier
	}
	return json.Marshal(toSerialize)
}

type NullableSubmitSelfServiceLoginFlowWithPasswordMethod struct {
	value *SubmitSelfServiceLoginFlowWithPasswordMethod
	isSet bool
}

func (v NullableSubmitSelfServiceLoginFlowWithPasswordMethod) Get() *SubmitSelfServiceLoginFlowWithPasswordMethod {
	return v.value
}

func (v *NullableSubmitSelfServiceLoginFlowWithPasswordMethod) Set(val *SubmitSelfServiceLoginFlowWithPasswordMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitSelfServiceLoginFlowWithPasswordMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitSelfServiceLoginFlowWithPasswordMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitSelfServiceLoginFlowWithPasswordMethod(val *SubmitSelfServiceLoginFlowWithPasswordMethod) *NullableSubmitSelfServiceLoginFlowWithPasswordMethod {
	return &NullableSubmitSelfServiceLoginFlowWithPasswordMethod{value: val, isSet: true}
}

func (v NullableSubmitSelfServiceLoginFlowWithPasswordMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitSelfServiceLoginFlowWithPasswordMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


