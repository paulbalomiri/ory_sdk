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

// IdentityCredentials Credentials represents a specific credential type
type IdentityCredentials struct {
	Config map[string]interface{} `json:"config,omitempty"`
	// CreatedAt is a helper struct field for gobuffalo.pop.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Identifiers represents a list of unique identifiers this credential type matches.
	Identifiers []string `json:"identifiers,omitempty"`
	Type *IdentityCredentialsType `json:"type,omitempty"`
	// UpdatedAt is a helper struct field for gobuffalo.pop.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Version refers to the version of the credential. Useful when changing the config schema.
	Version *int64 `json:"version,omitempty"`
}

// NewIdentityCredentials instantiates a new IdentityCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityCredentials() *IdentityCredentials {
	this := IdentityCredentials{}
	return &this
}

// NewIdentityCredentialsWithDefaults instantiates a new IdentityCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityCredentialsWithDefaults() *IdentityCredentials {
	this := IdentityCredentials{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *IdentityCredentials) GetConfig() map[string]interface{} {
	if o == nil || o.Config == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCredentials) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *IdentityCredentials) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *IdentityCredentials) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *IdentityCredentials) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCredentials) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *IdentityCredentials) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *IdentityCredentials) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetIdentifiers returns the Identifiers field value if set, zero value otherwise.
func (o *IdentityCredentials) GetIdentifiers() []string {
	if o == nil || o.Identifiers == nil {
		var ret []string
		return ret
	}
	return o.Identifiers
}

// GetIdentifiersOk returns a tuple with the Identifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCredentials) GetIdentifiersOk() ([]string, bool) {
	if o == nil || o.Identifiers == nil {
		return nil, false
	}
	return o.Identifiers, true
}

// HasIdentifiers returns a boolean if a field has been set.
func (o *IdentityCredentials) HasIdentifiers() bool {
	if o != nil && o.Identifiers != nil {
		return true
	}

	return false
}

// SetIdentifiers gets a reference to the given []string and assigns it to the Identifiers field.
func (o *IdentityCredentials) SetIdentifiers(v []string) {
	o.Identifiers = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IdentityCredentials) GetType() IdentityCredentialsType {
	if o == nil || o.Type == nil {
		var ret IdentityCredentialsType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCredentials) GetTypeOk() (*IdentityCredentialsType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IdentityCredentials) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given IdentityCredentialsType and assigns it to the Type field.
func (o *IdentityCredentials) SetType(v IdentityCredentialsType) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *IdentityCredentials) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCredentials) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *IdentityCredentials) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *IdentityCredentials) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *IdentityCredentials) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCredentials) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *IdentityCredentials) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *IdentityCredentials) SetVersion(v int64) {
	o.Version = &v
}

func (o IdentityCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Identifiers != nil {
		toSerialize["identifiers"] = o.Identifiers
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityCredentials struct {
	value *IdentityCredentials
	isSet bool
}

func (v NullableIdentityCredentials) Get() *IdentityCredentials {
	return v.value
}

func (v *NullableIdentityCredentials) Set(val *IdentityCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityCredentials(val *IdentityCredentials) *NullableIdentityCredentials {
	return &NullableIdentityCredentials{value: val, isSet: true}
}

func (v NullableIdentityCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


