/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * API version: v0.10.1
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SuccessfulSelfServiceLoginWithoutBrowser The Response for Login Flows via API
type SuccessfulSelfServiceLoginWithoutBrowser struct {
	Session Session `json:"session"`
	// The Session Token  A session token is equivalent to a session cookie, but it can be sent in the HTTP Authorization Header:  Authorization: bearer ${session-token}  The session token is only issued for API flows, not for Browser flows!
	SessionToken *string `json:"session_token,omitempty"`
}

// NewSuccessfulSelfServiceLoginWithoutBrowser instantiates a new SuccessfulSelfServiceLoginWithoutBrowser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessfulSelfServiceLoginWithoutBrowser(session Session) *SuccessfulSelfServiceLoginWithoutBrowser {
	this := SuccessfulSelfServiceLoginWithoutBrowser{}
	this.Session = session
	return &this
}

// NewSuccessfulSelfServiceLoginWithoutBrowserWithDefaults instantiates a new SuccessfulSelfServiceLoginWithoutBrowser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessfulSelfServiceLoginWithoutBrowserWithDefaults() *SuccessfulSelfServiceLoginWithoutBrowser {
	this := SuccessfulSelfServiceLoginWithoutBrowser{}
	return &this
}

// GetSession returns the Session field value
func (o *SuccessfulSelfServiceLoginWithoutBrowser) GetSession() Session {
	if o == nil {
		var ret Session
		return ret
	}

	return o.Session
}

// GetSessionOk returns a tuple with the Session field value
// and a boolean to check if the value has been set.
func (o *SuccessfulSelfServiceLoginWithoutBrowser) GetSessionOk() (*Session, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Session, true
}

// SetSession sets field value
func (o *SuccessfulSelfServiceLoginWithoutBrowser) SetSession(v Session) {
	o.Session = v
}

// GetSessionToken returns the SessionToken field value if set, zero value otherwise.
func (o *SuccessfulSelfServiceLoginWithoutBrowser) GetSessionToken() string {
	if o == nil || o.SessionToken == nil {
		var ret string
		return ret
	}
	return *o.SessionToken
}

// GetSessionTokenOk returns a tuple with the SessionToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessfulSelfServiceLoginWithoutBrowser) GetSessionTokenOk() (*string, bool) {
	if o == nil || o.SessionToken == nil {
		return nil, false
	}
	return o.SessionToken, true
}

// HasSessionToken returns a boolean if a field has been set.
func (o *SuccessfulSelfServiceLoginWithoutBrowser) HasSessionToken() bool {
	if o != nil && o.SessionToken != nil {
		return true
	}

	return false
}

// SetSessionToken gets a reference to the given string and assigns it to the SessionToken field.
func (o *SuccessfulSelfServiceLoginWithoutBrowser) SetSessionToken(v string) {
	o.SessionToken = &v
}

func (o SuccessfulSelfServiceLoginWithoutBrowser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["session"] = o.Session
	}
	if o.SessionToken != nil {
		toSerialize["session_token"] = o.SessionToken
	}
	return json.Marshal(toSerialize)
}

type NullableSuccessfulSelfServiceLoginWithoutBrowser struct {
	value *SuccessfulSelfServiceLoginWithoutBrowser
	isSet bool
}

func (v NullableSuccessfulSelfServiceLoginWithoutBrowser) Get() *SuccessfulSelfServiceLoginWithoutBrowser {
	return v.value
}

func (v *NullableSuccessfulSelfServiceLoginWithoutBrowser) Set(val *SuccessfulSelfServiceLoginWithoutBrowser) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessfulSelfServiceLoginWithoutBrowser) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessfulSelfServiceLoginWithoutBrowser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessfulSelfServiceLoginWithoutBrowser(val *SuccessfulSelfServiceLoginWithoutBrowser) *NullableSuccessfulSelfServiceLoginWithoutBrowser {
	return &NullableSuccessfulSelfServiceLoginWithoutBrowser{value: val, isSet: true}
}

func (v NullableSuccessfulSelfServiceLoginWithoutBrowser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessfulSelfServiceLoginWithoutBrowser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


