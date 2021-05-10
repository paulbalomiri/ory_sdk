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

// ContainerCreateCreatedBody ContainerCreateCreatedBody OK response to ContainerCreate operation
type ContainerCreateCreatedBody struct {
	// The ID of the created container
	Id string `json:"Id"`
	// Warnings encountered when creating the container
	Warnings []string `json:"Warnings"`
}

// NewContainerCreateCreatedBody instantiates a new ContainerCreateCreatedBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerCreateCreatedBody(id string, warnings []string) *ContainerCreateCreatedBody {
	this := ContainerCreateCreatedBody{}
	this.Id = id
	this.Warnings = warnings
	return &this
}

// NewContainerCreateCreatedBodyWithDefaults instantiates a new ContainerCreateCreatedBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerCreateCreatedBodyWithDefaults() *ContainerCreateCreatedBody {
	this := ContainerCreateCreatedBody{}
	return &this
}

// GetId returns the Id field value
func (o *ContainerCreateCreatedBody) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ContainerCreateCreatedBody) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ContainerCreateCreatedBody) SetId(v string) {
	o.Id = v
}

// GetWarnings returns the Warnings field value
func (o *ContainerCreateCreatedBody) GetWarnings() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
func (o *ContainerCreateCreatedBody) GetWarningsOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Warnings, true
}

// SetWarnings sets field value
func (o *ContainerCreateCreatedBody) SetWarnings(v []string) {
	o.Warnings = v
}

func (o ContainerCreateCreatedBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Id"] = o.Id
	}
	if true {
		toSerialize["Warnings"] = o.Warnings
	}
	return json.Marshal(toSerialize)
}

type NullableContainerCreateCreatedBody struct {
	value *ContainerCreateCreatedBody
	isSet bool
}

func (v NullableContainerCreateCreatedBody) Get() *ContainerCreateCreatedBody {
	return v.value
}

func (v *NullableContainerCreateCreatedBody) Set(val *ContainerCreateCreatedBody) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerCreateCreatedBody) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerCreateCreatedBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerCreateCreatedBody(val *ContainerCreateCreatedBody) *NullableContainerCreateCreatedBody {
	return &NullableContainerCreateCreatedBody{value: val, isSet: true}
}

func (v NullableContainerCreateCreatedBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerCreateCreatedBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


