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

// GetManagedIdentitySchemaLocation struct for GetManagedIdentitySchemaLocation
type GetManagedIdentitySchemaLocation struct {
	Location *string `json:"location,omitempty"`
}

// NewGetManagedIdentitySchemaLocation instantiates a new GetManagedIdentitySchemaLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetManagedIdentitySchemaLocation() *GetManagedIdentitySchemaLocation {
	this := GetManagedIdentitySchemaLocation{}
	return &this
}

// NewGetManagedIdentitySchemaLocationWithDefaults instantiates a new GetManagedIdentitySchemaLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetManagedIdentitySchemaLocationWithDefaults() *GetManagedIdentitySchemaLocation {
	this := GetManagedIdentitySchemaLocation{}
	return &this
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *GetManagedIdentitySchemaLocation) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedIdentitySchemaLocation) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *GetManagedIdentitySchemaLocation) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *GetManagedIdentitySchemaLocation) SetLocation(v string) {
	o.Location = &v
}

func (o GetManagedIdentitySchemaLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	return json.Marshal(toSerialize)
}

type NullableGetManagedIdentitySchemaLocation struct {
	value *GetManagedIdentitySchemaLocation
	isSet bool
}

func (v NullableGetManagedIdentitySchemaLocation) Get() *GetManagedIdentitySchemaLocation {
	return v.value
}

func (v *NullableGetManagedIdentitySchemaLocation) Set(val *GetManagedIdentitySchemaLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableGetManagedIdentitySchemaLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableGetManagedIdentitySchemaLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetManagedIdentitySchemaLocation(val *GetManagedIdentitySchemaLocation) *NullableGetManagedIdentitySchemaLocation {
	return &NullableGetManagedIdentitySchemaLocation{value: val, isSet: true}
}

func (v NullableGetManagedIdentitySchemaLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetManagedIdentitySchemaLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


