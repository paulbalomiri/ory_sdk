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

// SubmitSelfServiceSettingsFlowWithOidcMethodBody nolint:deadcode,unused
type SubmitSelfServiceSettingsFlowWithOidcMethodBody struct {
	// Flow ID is the flow's ID.  in: query
	Flow *string `json:"flow,omitempty"`
	// Link this provider  Either this or `unlink` must be set.  type: string in: body
	Link *string `json:"link,omitempty"`
	// Method  Should be set to profile when trying to update a profile.
	Method string `json:"method"`
	// The identity's traits  in: body
	Traits map[string]interface{} `json:"traits,omitempty"`
	// Unlink this provider  Either this or `link` must be set.  type: string in: body
	Unlink *string `json:"unlink,omitempty"`
}

// NewSubmitSelfServiceSettingsFlowWithOidcMethodBody instantiates a new SubmitSelfServiceSettingsFlowWithOidcMethodBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitSelfServiceSettingsFlowWithOidcMethodBody(method string) *SubmitSelfServiceSettingsFlowWithOidcMethodBody {
	this := SubmitSelfServiceSettingsFlowWithOidcMethodBody{}
	this.Method = method
	return &this
}

// NewSubmitSelfServiceSettingsFlowWithOidcMethodBodyWithDefaults instantiates a new SubmitSelfServiceSettingsFlowWithOidcMethodBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitSelfServiceSettingsFlowWithOidcMethodBodyWithDefaults() *SubmitSelfServiceSettingsFlowWithOidcMethodBody {
	this := SubmitSelfServiceSettingsFlowWithOidcMethodBody{}
	return &this
}

// GetFlow returns the Flow field value if set, zero value otherwise.
func (o *SubmitSelfServiceSettingsFlowWithOidcMethodBody) GetFlow() string {
	if o == nil || o.Flow == nil {
		var ret string
		return ret
	}
	return *o.Flow
}

// GetFlowOk returns a tuple with the Flow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceSettingsFlowWithOidcMethodBody) GetFlowOk() (*string, bool) {
	if o == nil || o.Flow == nil {
		return nil, false
	}
	return o.Flow, true
}

// HasFlow returns a boolean if a field has been set.
func (o *SubmitSelfServiceSettingsFlowWithOidcMethodBody) HasFlow() bool {
	if o != nil && o.Flow != nil {
		return true
	}

	return false
}

// SetFlow gets a reference to the given string and assigns it to the Flow field.
func (o *SubmitSelfServiceSettingsFlowWithOidcMethodBody) SetFlow(v string) {
	o.Flow = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *SubmitSelfServiceSettingsFlowWithOidcMethodBody) GetLink() string {
	if o == nil || o.Link == nil {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceSettingsFlowWithOidcMethodBody) GetLinkOk() (*string, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *SubmitSelfServiceSettingsFlowWithOidcMethodBody) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *SubmitSelfServiceSettingsFlowWithOidcMethodBody) SetLink(v string) {
	o.Link = &v
}

// GetMethod returns the Method field value
func (o *SubmitSelfServiceSettingsFlowWithOidcMethodBody) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceSettingsFlowWithOidcMethodBody) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *SubmitSelfServiceSettingsFlowWithOidcMethodBody) SetMethod(v string) {
	o.Method = v
}

// GetTraits returns the Traits field value if set, zero value otherwise.
func (o *SubmitSelfServiceSettingsFlowWithOidcMethodBody) GetTraits() map[string]interface{} {
	if o == nil || o.Traits == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Traits
}

// GetTraitsOk returns a tuple with the Traits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceSettingsFlowWithOidcMethodBody) GetTraitsOk() (map[string]interface{}, bool) {
	if o == nil || o.Traits == nil {
		return nil, false
	}
	return o.Traits, true
}

// HasTraits returns a boolean if a field has been set.
func (o *SubmitSelfServiceSettingsFlowWithOidcMethodBody) HasTraits() bool {
	if o != nil && o.Traits != nil {
		return true
	}

	return false
}

// SetTraits gets a reference to the given map[string]interface{} and assigns it to the Traits field.
func (o *SubmitSelfServiceSettingsFlowWithOidcMethodBody) SetTraits(v map[string]interface{}) {
	o.Traits = v
}

// GetUnlink returns the Unlink field value if set, zero value otherwise.
func (o *SubmitSelfServiceSettingsFlowWithOidcMethodBody) GetUnlink() string {
	if o == nil || o.Unlink == nil {
		var ret string
		return ret
	}
	return *o.Unlink
}

// GetUnlinkOk returns a tuple with the Unlink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceSettingsFlowWithOidcMethodBody) GetUnlinkOk() (*string, bool) {
	if o == nil || o.Unlink == nil {
		return nil, false
	}
	return o.Unlink, true
}

// HasUnlink returns a boolean if a field has been set.
func (o *SubmitSelfServiceSettingsFlowWithOidcMethodBody) HasUnlink() bool {
	if o != nil && o.Unlink != nil {
		return true
	}

	return false
}

// SetUnlink gets a reference to the given string and assigns it to the Unlink field.
func (o *SubmitSelfServiceSettingsFlowWithOidcMethodBody) SetUnlink(v string) {
	o.Unlink = &v
}

func (o SubmitSelfServiceSettingsFlowWithOidcMethodBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Flow != nil {
		toSerialize["flow"] = o.Flow
	}
	if o.Link != nil {
		toSerialize["link"] = o.Link
	}
	if true {
		toSerialize["method"] = o.Method
	}
	if o.Traits != nil {
		toSerialize["traits"] = o.Traits
	}
	if o.Unlink != nil {
		toSerialize["unlink"] = o.Unlink
	}
	return json.Marshal(toSerialize)
}

type NullableSubmitSelfServiceSettingsFlowWithOidcMethodBody struct {
	value *SubmitSelfServiceSettingsFlowWithOidcMethodBody
	isSet bool
}

func (v NullableSubmitSelfServiceSettingsFlowWithOidcMethodBody) Get() *SubmitSelfServiceSettingsFlowWithOidcMethodBody {
	return v.value
}

func (v *NullableSubmitSelfServiceSettingsFlowWithOidcMethodBody) Set(val *SubmitSelfServiceSettingsFlowWithOidcMethodBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitSelfServiceSettingsFlowWithOidcMethodBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitSelfServiceSettingsFlowWithOidcMethodBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitSelfServiceSettingsFlowWithOidcMethodBody(val *SubmitSelfServiceSettingsFlowWithOidcMethodBody) *NullableSubmitSelfServiceSettingsFlowWithOidcMethodBody {
	return &NullableSubmitSelfServiceSettingsFlowWithOidcMethodBody{value: val, isSet: true}
}

func (v NullableSubmitSelfServiceSettingsFlowWithOidcMethodBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitSelfServiceSettingsFlowWithOidcMethodBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


