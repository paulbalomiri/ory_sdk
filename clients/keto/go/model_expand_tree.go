/*
Ory Keto API

Documentation for all of Ory Keto's REST APIs. gRPC is documented separately. 

API version: v0.10.0-alpha.0
Contact: hi@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ExpandTree struct for ExpandTree
type ExpandTree struct {
	// The children of the node, possibly none.
	Children []ExpandTree `json:"children,omitempty"`
	Tuple *RelationTuple `json:"tuple,omitempty"`
	// The type of the node. union TreeNodeUnion exclusion TreeNodeExclusion intersection TreeNodeIntersection leaf TreeNodeLeaf tuple_to_subject_set TreeNodeTupleToSubjectSet computed_subject_set TreeNodeComputedSubjectSet not TreeNodeNot unspecified TreeNodeUnspecified
	Type string `json:"type"`
}

// NewExpandTree instantiates a new ExpandTree object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpandTree(type_ string) *ExpandTree {
	this := ExpandTree{}
	this.Type = type_
	return &this
}

// NewExpandTreeWithDefaults instantiates a new ExpandTree object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpandTreeWithDefaults() *ExpandTree {
	this := ExpandTree{}
	return &this
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *ExpandTree) GetChildren() []ExpandTree {
	if o == nil || o.Children == nil {
		var ret []ExpandTree
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandTree) GetChildrenOk() ([]ExpandTree, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *ExpandTree) HasChildren() bool {
	if o != nil && o.Children != nil {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []ExpandTree and assigns it to the Children field.
func (o *ExpandTree) SetChildren(v []ExpandTree) {
	o.Children = v
}

// GetTuple returns the Tuple field value if set, zero value otherwise.
func (o *ExpandTree) GetTuple() RelationTuple {
	if o == nil || o.Tuple == nil {
		var ret RelationTuple
		return ret
	}
	return *o.Tuple
}

// GetTupleOk returns a tuple with the Tuple field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandTree) GetTupleOk() (*RelationTuple, bool) {
	if o == nil || o.Tuple == nil {
		return nil, false
	}
	return o.Tuple, true
}

// HasTuple returns a boolean if a field has been set.
func (o *ExpandTree) HasTuple() bool {
	if o != nil && o.Tuple != nil {
		return true
	}

	return false
}

// SetTuple gets a reference to the given RelationTuple and assigns it to the Tuple field.
func (o *ExpandTree) SetTuple(v RelationTuple) {
	o.Tuple = &v
}

// GetType returns the Type field value
func (o *ExpandTree) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExpandTree) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExpandTree) SetType(v string) {
	o.Type = v
}

func (o ExpandTree) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}
	if o.Tuple != nil {
		toSerialize["tuple"] = o.Tuple
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableExpandTree struct {
	value *ExpandTree
	isSet bool
}

func (v NullableExpandTree) Get() *ExpandTree {
	return v.value
}

func (v *NullableExpandTree) Set(val *ExpandTree) {
	v.value = val
	v.isSet = true
}

func (v NullableExpandTree) IsSet() bool {
	return v.isSet
}

func (v *NullableExpandTree) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpandTree(val *ExpandTree) *NullableExpandTree {
	return &NullableExpandTree{value: val, isSet: true}
}

func (v NullableExpandTree) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpandTree) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


