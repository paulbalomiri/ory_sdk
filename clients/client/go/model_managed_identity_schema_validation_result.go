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
)

// ManagedIdentitySchemaValidationResult struct for ManagedIdentitySchemaValidationResult
type ManagedIdentitySchemaValidationResult struct {
	Message *string `json:"message,omitempty"`
	Valid *bool `json:"valid,omitempty"`
}

// NewManagedIdentitySchemaValidationResult instantiates a new ManagedIdentitySchemaValidationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedIdentitySchemaValidationResult() *ManagedIdentitySchemaValidationResult {
	this := ManagedIdentitySchemaValidationResult{}
	return &this
}

// NewManagedIdentitySchemaValidationResultWithDefaults instantiates a new ManagedIdentitySchemaValidationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedIdentitySchemaValidationResultWithDefaults() *ManagedIdentitySchemaValidationResult {
	this := ManagedIdentitySchemaValidationResult{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ManagedIdentitySchemaValidationResult) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedIdentitySchemaValidationResult) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ManagedIdentitySchemaValidationResult) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ManagedIdentitySchemaValidationResult) SetMessage(v string) {
	o.Message = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *ManagedIdentitySchemaValidationResult) GetValid() bool {
	if o == nil || o.Valid == nil {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedIdentitySchemaValidationResult) GetValidOk() (*bool, bool) {
	if o == nil || o.Valid == nil {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *ManagedIdentitySchemaValidationResult) HasValid() bool {
	if o != nil && o.Valid != nil {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *ManagedIdentitySchemaValidationResult) SetValid(v bool) {
	o.Valid = &v
}

func (o ManagedIdentitySchemaValidationResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Valid != nil {
		toSerialize["valid"] = o.Valid
	}
	return json.Marshal(toSerialize)
}

type NullableManagedIdentitySchemaValidationResult struct {
	value *ManagedIdentitySchemaValidationResult
	isSet bool
}

func (v NullableManagedIdentitySchemaValidationResult) Get() *ManagedIdentitySchemaValidationResult {
	return v.value
}

func (v *NullableManagedIdentitySchemaValidationResult) Set(val *ManagedIdentitySchemaValidationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedIdentitySchemaValidationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedIdentitySchemaValidationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedIdentitySchemaValidationResult(val *ManagedIdentitySchemaValidationResult) *NullableManagedIdentitySchemaValidationResult {
	return &NullableManagedIdentitySchemaValidationResult{value: val, isSet: true}
}

func (v NullableManagedIdentitySchemaValidationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedIdentitySchemaValidationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


