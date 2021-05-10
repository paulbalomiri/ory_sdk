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

// GenericErrorPayload struct for GenericErrorPayload
type GenericErrorPayload struct {
	// Code represents the error status code (404, 403, 401, ...).
	Code *int64 `json:"code,omitempty"`
	// Debug contains debug information. This is usually not available and has to be enabled.
	Debug *string `json:"debug,omitempty"`
	Details map[string]map[string]interface{} `json:"details,omitempty"`
	Message *string `json:"message,omitempty"`
	Reason *string `json:"reason,omitempty"`
	Request *string `json:"request,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewGenericErrorPayload instantiates a new GenericErrorPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericErrorPayload() *GenericErrorPayload {
	this := GenericErrorPayload{}
	return &this
}

// NewGenericErrorPayloadWithDefaults instantiates a new GenericErrorPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericErrorPayloadWithDefaults() *GenericErrorPayload {
	this := GenericErrorPayload{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GenericErrorPayload) GetCode() int64 {
	if o == nil || o.Code == nil {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericErrorPayload) GetCodeOk() (*int64, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GenericErrorPayload) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *GenericErrorPayload) SetCode(v int64) {
	o.Code = &v
}

// GetDebug returns the Debug field value if set, zero value otherwise.
func (o *GenericErrorPayload) GetDebug() string {
	if o == nil || o.Debug == nil {
		var ret string
		return ret
	}
	return *o.Debug
}

// GetDebugOk returns a tuple with the Debug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericErrorPayload) GetDebugOk() (*string, bool) {
	if o == nil || o.Debug == nil {
		return nil, false
	}
	return o.Debug, true
}

// HasDebug returns a boolean if a field has been set.
func (o *GenericErrorPayload) HasDebug() bool {
	if o != nil && o.Debug != nil {
		return true
	}

	return false
}

// SetDebug gets a reference to the given string and assigns it to the Debug field.
func (o *GenericErrorPayload) SetDebug(v string) {
	o.Debug = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *GenericErrorPayload) GetDetails() map[string]map[string]interface{} {
	if o == nil || o.Details == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericErrorPayload) GetDetailsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *GenericErrorPayload) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given map[string]map[string]interface{} and assigns it to the Details field.
func (o *GenericErrorPayload) SetDetails(v map[string]map[string]interface{}) {
	o.Details = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GenericErrorPayload) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericErrorPayload) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GenericErrorPayload) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GenericErrorPayload) SetMessage(v string) {
	o.Message = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *GenericErrorPayload) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericErrorPayload) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *GenericErrorPayload) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *GenericErrorPayload) SetReason(v string) {
	o.Reason = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *GenericErrorPayload) GetRequest() string {
	if o == nil || o.Request == nil {
		var ret string
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericErrorPayload) GetRequestOk() (*string, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *GenericErrorPayload) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given string and assigns it to the Request field.
func (o *GenericErrorPayload) SetRequest(v string) {
	o.Request = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GenericErrorPayload) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericErrorPayload) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GenericErrorPayload) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GenericErrorPayload) SetStatus(v string) {
	o.Status = &v
}

func (o GenericErrorPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Debug != nil {
		toSerialize["debug"] = o.Debug
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableGenericErrorPayload struct {
	value *GenericErrorPayload
	isSet bool
}

func (v NullableGenericErrorPayload) Get() *GenericErrorPayload {
	return v.value
}

func (v *NullableGenericErrorPayload) Set(val *GenericErrorPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericErrorPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericErrorPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericErrorPayload(val *GenericErrorPayload) *NullableGenericErrorPayload {
	return &NullableGenericErrorPayload{value: val, isSet: true}
}

func (v NullableGenericErrorPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericErrorPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


