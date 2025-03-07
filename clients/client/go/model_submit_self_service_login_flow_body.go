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
	"fmt"
)

// SubmitSelfServiceLoginFlowBody - struct for SubmitSelfServiceLoginFlowBody
type SubmitSelfServiceLoginFlowBody struct {
	SubmitSelfServiceLoginFlowWithLookupSecretMethodBody *SubmitSelfServiceLoginFlowWithLookupSecretMethodBody
	SubmitSelfServiceLoginFlowWithOidcMethodBody *SubmitSelfServiceLoginFlowWithOidcMethodBody
	SubmitSelfServiceLoginFlowWithPasswordMethodBody *SubmitSelfServiceLoginFlowWithPasswordMethodBody
	SubmitSelfServiceLoginFlowWithTotpMethodBody *SubmitSelfServiceLoginFlowWithTotpMethodBody
	SubmitSelfServiceLoginFlowWithWebAuthnMethodBody *SubmitSelfServiceLoginFlowWithWebAuthnMethodBody
}

// SubmitSelfServiceLoginFlowWithLookupSecretMethodBodyAsSubmitSelfServiceLoginFlowBody is a convenience function that returns SubmitSelfServiceLoginFlowWithLookupSecretMethodBody wrapped in SubmitSelfServiceLoginFlowBody
func SubmitSelfServiceLoginFlowWithLookupSecretMethodBodyAsSubmitSelfServiceLoginFlowBody(v *SubmitSelfServiceLoginFlowWithLookupSecretMethodBody) SubmitSelfServiceLoginFlowBody {
	return SubmitSelfServiceLoginFlowBody{
		SubmitSelfServiceLoginFlowWithLookupSecretMethodBody: v,
	}
}

// SubmitSelfServiceLoginFlowWithOidcMethodBodyAsSubmitSelfServiceLoginFlowBody is a convenience function that returns SubmitSelfServiceLoginFlowWithOidcMethodBody wrapped in SubmitSelfServiceLoginFlowBody
func SubmitSelfServiceLoginFlowWithOidcMethodBodyAsSubmitSelfServiceLoginFlowBody(v *SubmitSelfServiceLoginFlowWithOidcMethodBody) SubmitSelfServiceLoginFlowBody {
	return SubmitSelfServiceLoginFlowBody{
		SubmitSelfServiceLoginFlowWithOidcMethodBody: v,
	}
}

// SubmitSelfServiceLoginFlowWithPasswordMethodBodyAsSubmitSelfServiceLoginFlowBody is a convenience function that returns SubmitSelfServiceLoginFlowWithPasswordMethodBody wrapped in SubmitSelfServiceLoginFlowBody
func SubmitSelfServiceLoginFlowWithPasswordMethodBodyAsSubmitSelfServiceLoginFlowBody(v *SubmitSelfServiceLoginFlowWithPasswordMethodBody) SubmitSelfServiceLoginFlowBody {
	return SubmitSelfServiceLoginFlowBody{
		SubmitSelfServiceLoginFlowWithPasswordMethodBody: v,
	}
}

// SubmitSelfServiceLoginFlowWithTotpMethodBodyAsSubmitSelfServiceLoginFlowBody is a convenience function that returns SubmitSelfServiceLoginFlowWithTotpMethodBody wrapped in SubmitSelfServiceLoginFlowBody
func SubmitSelfServiceLoginFlowWithTotpMethodBodyAsSubmitSelfServiceLoginFlowBody(v *SubmitSelfServiceLoginFlowWithTotpMethodBody) SubmitSelfServiceLoginFlowBody {
	return SubmitSelfServiceLoginFlowBody{
		SubmitSelfServiceLoginFlowWithTotpMethodBody: v,
	}
}

// SubmitSelfServiceLoginFlowWithWebAuthnMethodBodyAsSubmitSelfServiceLoginFlowBody is a convenience function that returns SubmitSelfServiceLoginFlowWithWebAuthnMethodBody wrapped in SubmitSelfServiceLoginFlowBody
func SubmitSelfServiceLoginFlowWithWebAuthnMethodBodyAsSubmitSelfServiceLoginFlowBody(v *SubmitSelfServiceLoginFlowWithWebAuthnMethodBody) SubmitSelfServiceLoginFlowBody {
	return SubmitSelfServiceLoginFlowBody{
		SubmitSelfServiceLoginFlowWithWebAuthnMethodBody: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SubmitSelfServiceLoginFlowBody) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SubmitSelfServiceLoginFlowWithLookupSecretMethodBody
	err = newStrictDecoder(data).Decode(&dst.SubmitSelfServiceLoginFlowWithLookupSecretMethodBody)
	if err == nil {
		jsonSubmitSelfServiceLoginFlowWithLookupSecretMethodBody, _ := json.Marshal(dst.SubmitSelfServiceLoginFlowWithLookupSecretMethodBody)
		if string(jsonSubmitSelfServiceLoginFlowWithLookupSecretMethodBody) == "{}" { // empty struct
			dst.SubmitSelfServiceLoginFlowWithLookupSecretMethodBody = nil
		} else {
			match++
		}
	} else {
		dst.SubmitSelfServiceLoginFlowWithLookupSecretMethodBody = nil
	}

	// try to unmarshal data into SubmitSelfServiceLoginFlowWithOidcMethodBody
	err = newStrictDecoder(data).Decode(&dst.SubmitSelfServiceLoginFlowWithOidcMethodBody)
	if err == nil {
		jsonSubmitSelfServiceLoginFlowWithOidcMethodBody, _ := json.Marshal(dst.SubmitSelfServiceLoginFlowWithOidcMethodBody)
		if string(jsonSubmitSelfServiceLoginFlowWithOidcMethodBody) == "{}" { // empty struct
			dst.SubmitSelfServiceLoginFlowWithOidcMethodBody = nil
		} else {
			match++
		}
	} else {
		dst.SubmitSelfServiceLoginFlowWithOidcMethodBody = nil
	}

	// try to unmarshal data into SubmitSelfServiceLoginFlowWithPasswordMethodBody
	err = newStrictDecoder(data).Decode(&dst.SubmitSelfServiceLoginFlowWithPasswordMethodBody)
	if err == nil {
		jsonSubmitSelfServiceLoginFlowWithPasswordMethodBody, _ := json.Marshal(dst.SubmitSelfServiceLoginFlowWithPasswordMethodBody)
		if string(jsonSubmitSelfServiceLoginFlowWithPasswordMethodBody) == "{}" { // empty struct
			dst.SubmitSelfServiceLoginFlowWithPasswordMethodBody = nil
		} else {
			match++
		}
	} else {
		dst.SubmitSelfServiceLoginFlowWithPasswordMethodBody = nil
	}

	// try to unmarshal data into SubmitSelfServiceLoginFlowWithTotpMethodBody
	err = newStrictDecoder(data).Decode(&dst.SubmitSelfServiceLoginFlowWithTotpMethodBody)
	if err == nil {
		jsonSubmitSelfServiceLoginFlowWithTotpMethodBody, _ := json.Marshal(dst.SubmitSelfServiceLoginFlowWithTotpMethodBody)
		if string(jsonSubmitSelfServiceLoginFlowWithTotpMethodBody) == "{}" { // empty struct
			dst.SubmitSelfServiceLoginFlowWithTotpMethodBody = nil
		} else {
			match++
		}
	} else {
		dst.SubmitSelfServiceLoginFlowWithTotpMethodBody = nil
	}

	// try to unmarshal data into SubmitSelfServiceLoginFlowWithWebAuthnMethodBody
	err = newStrictDecoder(data).Decode(&dst.SubmitSelfServiceLoginFlowWithWebAuthnMethodBody)
	if err == nil {
		jsonSubmitSelfServiceLoginFlowWithWebAuthnMethodBody, _ := json.Marshal(dst.SubmitSelfServiceLoginFlowWithWebAuthnMethodBody)
		if string(jsonSubmitSelfServiceLoginFlowWithWebAuthnMethodBody) == "{}" { // empty struct
			dst.SubmitSelfServiceLoginFlowWithWebAuthnMethodBody = nil
		} else {
			match++
		}
	} else {
		dst.SubmitSelfServiceLoginFlowWithWebAuthnMethodBody = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SubmitSelfServiceLoginFlowWithLookupSecretMethodBody = nil
		dst.SubmitSelfServiceLoginFlowWithOidcMethodBody = nil
		dst.SubmitSelfServiceLoginFlowWithPasswordMethodBody = nil
		dst.SubmitSelfServiceLoginFlowWithTotpMethodBody = nil
		dst.SubmitSelfServiceLoginFlowWithWebAuthnMethodBody = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(SubmitSelfServiceLoginFlowBody)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(SubmitSelfServiceLoginFlowBody)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SubmitSelfServiceLoginFlowBody) MarshalJSON() ([]byte, error) {
	if src.SubmitSelfServiceLoginFlowWithLookupSecretMethodBody != nil {
		return json.Marshal(&src.SubmitSelfServiceLoginFlowWithLookupSecretMethodBody)
	}

	if src.SubmitSelfServiceLoginFlowWithOidcMethodBody != nil {
		return json.Marshal(&src.SubmitSelfServiceLoginFlowWithOidcMethodBody)
	}

	if src.SubmitSelfServiceLoginFlowWithPasswordMethodBody != nil {
		return json.Marshal(&src.SubmitSelfServiceLoginFlowWithPasswordMethodBody)
	}

	if src.SubmitSelfServiceLoginFlowWithTotpMethodBody != nil {
		return json.Marshal(&src.SubmitSelfServiceLoginFlowWithTotpMethodBody)
	}

	if src.SubmitSelfServiceLoginFlowWithWebAuthnMethodBody != nil {
		return json.Marshal(&src.SubmitSelfServiceLoginFlowWithWebAuthnMethodBody)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SubmitSelfServiceLoginFlowBody) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.SubmitSelfServiceLoginFlowWithLookupSecretMethodBody != nil {
		return obj.SubmitSelfServiceLoginFlowWithLookupSecretMethodBody
	}

	if obj.SubmitSelfServiceLoginFlowWithOidcMethodBody != nil {
		return obj.SubmitSelfServiceLoginFlowWithOidcMethodBody
	}

	if obj.SubmitSelfServiceLoginFlowWithPasswordMethodBody != nil {
		return obj.SubmitSelfServiceLoginFlowWithPasswordMethodBody
	}

	if obj.SubmitSelfServiceLoginFlowWithTotpMethodBody != nil {
		return obj.SubmitSelfServiceLoginFlowWithTotpMethodBody
	}

	if obj.SubmitSelfServiceLoginFlowWithWebAuthnMethodBody != nil {
		return obj.SubmitSelfServiceLoginFlowWithWebAuthnMethodBody
	}

	// all schemas are nil
	return nil
}

type NullableSubmitSelfServiceLoginFlowBody struct {
	value *SubmitSelfServiceLoginFlowBody
	isSet bool
}

func (v NullableSubmitSelfServiceLoginFlowBody) Get() *SubmitSelfServiceLoginFlowBody {
	return v.value
}

func (v *NullableSubmitSelfServiceLoginFlowBody) Set(val *SubmitSelfServiceLoginFlowBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitSelfServiceLoginFlowBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitSelfServiceLoginFlowBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitSelfServiceLoginFlowBody(val *SubmitSelfServiceLoginFlowBody) *NullableSubmitSelfServiceLoginFlowBody {
	return &NullableSubmitSelfServiceLoginFlowBody{value: val, isSet: true}
}

func (v NullableSubmitSelfServiceLoginFlowBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitSelfServiceLoginFlowBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


