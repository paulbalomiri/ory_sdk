/*
Ory Kratos API

Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 

API version: v0.10.1
Contact: hi@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SubmitSelfServiceFlowWithWebAuthnRegistrationMethod struct for SubmitSelfServiceFlowWithWebAuthnRegistrationMethod
type SubmitSelfServiceFlowWithWebAuthnRegistrationMethod struct {
	// Register a WebAuthn Security Key  It is expected that the JSON returned by the WebAuthn registration process is included here.
	WebauthnRegister *string `json:"webauthn_register,omitempty"`
	// Name of the WebAuthn Security Key to be Added  A human-readable name for the security key which will be added.
	WebauthnRegisterDisplayname *string `json:"webauthn_register_displayname,omitempty"`
}

// NewSubmitSelfServiceFlowWithWebAuthnRegistrationMethod instantiates a new SubmitSelfServiceFlowWithWebAuthnRegistrationMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitSelfServiceFlowWithWebAuthnRegistrationMethod() *SubmitSelfServiceFlowWithWebAuthnRegistrationMethod {
	this := SubmitSelfServiceFlowWithWebAuthnRegistrationMethod{}
	return &this
}

// NewSubmitSelfServiceFlowWithWebAuthnRegistrationMethodWithDefaults instantiates a new SubmitSelfServiceFlowWithWebAuthnRegistrationMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitSelfServiceFlowWithWebAuthnRegistrationMethodWithDefaults() *SubmitSelfServiceFlowWithWebAuthnRegistrationMethod {
	this := SubmitSelfServiceFlowWithWebAuthnRegistrationMethod{}
	return &this
}

// GetWebauthnRegister returns the WebauthnRegister field value if set, zero value otherwise.
func (o *SubmitSelfServiceFlowWithWebAuthnRegistrationMethod) GetWebauthnRegister() string {
	if o == nil || o.WebauthnRegister == nil {
		var ret string
		return ret
	}
	return *o.WebauthnRegister
}

// GetWebauthnRegisterOk returns a tuple with the WebauthnRegister field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceFlowWithWebAuthnRegistrationMethod) GetWebauthnRegisterOk() (*string, bool) {
	if o == nil || o.WebauthnRegister == nil {
		return nil, false
	}
	return o.WebauthnRegister, true
}

// HasWebauthnRegister returns a boolean if a field has been set.
func (o *SubmitSelfServiceFlowWithWebAuthnRegistrationMethod) HasWebauthnRegister() bool {
	if o != nil && o.WebauthnRegister != nil {
		return true
	}

	return false
}

// SetWebauthnRegister gets a reference to the given string and assigns it to the WebauthnRegister field.
func (o *SubmitSelfServiceFlowWithWebAuthnRegistrationMethod) SetWebauthnRegister(v string) {
	o.WebauthnRegister = &v
}

// GetWebauthnRegisterDisplayname returns the WebauthnRegisterDisplayname field value if set, zero value otherwise.
func (o *SubmitSelfServiceFlowWithWebAuthnRegistrationMethod) GetWebauthnRegisterDisplayname() string {
	if o == nil || o.WebauthnRegisterDisplayname == nil {
		var ret string
		return ret
	}
	return *o.WebauthnRegisterDisplayname
}

// GetWebauthnRegisterDisplaynameOk returns a tuple with the WebauthnRegisterDisplayname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceFlowWithWebAuthnRegistrationMethod) GetWebauthnRegisterDisplaynameOk() (*string, bool) {
	if o == nil || o.WebauthnRegisterDisplayname == nil {
		return nil, false
	}
	return o.WebauthnRegisterDisplayname, true
}

// HasWebauthnRegisterDisplayname returns a boolean if a field has been set.
func (o *SubmitSelfServiceFlowWithWebAuthnRegistrationMethod) HasWebauthnRegisterDisplayname() bool {
	if o != nil && o.WebauthnRegisterDisplayname != nil {
		return true
	}

	return false
}

// SetWebauthnRegisterDisplayname gets a reference to the given string and assigns it to the WebauthnRegisterDisplayname field.
func (o *SubmitSelfServiceFlowWithWebAuthnRegistrationMethod) SetWebauthnRegisterDisplayname(v string) {
	o.WebauthnRegisterDisplayname = &v
}

func (o SubmitSelfServiceFlowWithWebAuthnRegistrationMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WebauthnRegister != nil {
		toSerialize["webauthn_register"] = o.WebauthnRegister
	}
	if o.WebauthnRegisterDisplayname != nil {
		toSerialize["webauthn_register_displayname"] = o.WebauthnRegisterDisplayname
	}
	return json.Marshal(toSerialize)
}

type NullableSubmitSelfServiceFlowWithWebAuthnRegistrationMethod struct {
	value *SubmitSelfServiceFlowWithWebAuthnRegistrationMethod
	isSet bool
}

func (v NullableSubmitSelfServiceFlowWithWebAuthnRegistrationMethod) Get() *SubmitSelfServiceFlowWithWebAuthnRegistrationMethod {
	return v.value
}

func (v *NullableSubmitSelfServiceFlowWithWebAuthnRegistrationMethod) Set(val *SubmitSelfServiceFlowWithWebAuthnRegistrationMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitSelfServiceFlowWithWebAuthnRegistrationMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitSelfServiceFlowWithWebAuthnRegistrationMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitSelfServiceFlowWithWebAuthnRegistrationMethod(val *SubmitSelfServiceFlowWithWebAuthnRegistrationMethod) *NullableSubmitSelfServiceFlowWithWebAuthnRegistrationMethod {
	return &NullableSubmitSelfServiceFlowWithWebAuthnRegistrationMethod{value: val, isSet: true}
}

func (v NullableSubmitSelfServiceFlowWithWebAuthnRegistrationMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitSelfServiceFlowWithWebAuthnRegistrationMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


