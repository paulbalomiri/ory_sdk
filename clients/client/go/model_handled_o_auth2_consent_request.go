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

// HandledOAuth2ConsentRequest struct for HandledOAuth2ConsentRequest
type HandledOAuth2ConsentRequest struct {
	// Original request URL to which you should redirect the user if request was already handled.
	RedirectTo string `json:"redirect_to"`
}

// NewHandledOAuth2ConsentRequest instantiates a new HandledOAuth2ConsentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandledOAuth2ConsentRequest(redirectTo string) *HandledOAuth2ConsentRequest {
	this := HandledOAuth2ConsentRequest{}
	this.RedirectTo = redirectTo
	return &this
}

// NewHandledOAuth2ConsentRequestWithDefaults instantiates a new HandledOAuth2ConsentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandledOAuth2ConsentRequestWithDefaults() *HandledOAuth2ConsentRequest {
	this := HandledOAuth2ConsentRequest{}
	return &this
}

// GetRedirectTo returns the RedirectTo field value
func (o *HandledOAuth2ConsentRequest) GetRedirectTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedirectTo
}

// GetRedirectToOk returns a tuple with the RedirectTo field value
// and a boolean to check if the value has been set.
func (o *HandledOAuth2ConsentRequest) GetRedirectToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedirectTo, true
}

// SetRedirectTo sets field value
func (o *HandledOAuth2ConsentRequest) SetRedirectTo(v string) {
	o.RedirectTo = v
}

func (o HandledOAuth2ConsentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["redirect_to"] = o.RedirectTo
	}
	return json.Marshal(toSerialize)
}

type NullableHandledOAuth2ConsentRequest struct {
	value *HandledOAuth2ConsentRequest
	isSet bool
}

func (v NullableHandledOAuth2ConsentRequest) Get() *HandledOAuth2ConsentRequest {
	return v.value
}

func (v *NullableHandledOAuth2ConsentRequest) Set(val *HandledOAuth2ConsentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHandledOAuth2ConsentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHandledOAuth2ConsentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandledOAuth2ConsentRequest(val *HandledOAuth2ConsentRequest) *NullableHandledOAuth2ConsentRequest {
	return &NullableHandledOAuth2ConsentRequest{value: val, isSet: true}
}

func (v NullableHandledOAuth2ConsentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandledOAuth2ConsentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


