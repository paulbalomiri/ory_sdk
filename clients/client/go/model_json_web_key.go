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

// JsonWebKey struct for JsonWebKey
type JsonWebKey struct {
	// The \"alg\" (algorithm) parameter identifies the algorithm intended for use with the key.  The values used should either be registered in the IANA \"JSON Web Signature and Encryption Algorithms\" registry established by [JWA] or be a value that contains a Collision- Resistant Name.
	Alg string `json:"alg"`
	Crv *string `json:"crv,omitempty"`
	D *string `json:"d,omitempty"`
	Dp *string `json:"dp,omitempty"`
	Dq *string `json:"dq,omitempty"`
	E *string `json:"e,omitempty"`
	K *string `json:"k,omitempty"`
	// The \"kid\" (key ID) parameter is used to match a specific key.  This is used, for instance, to choose among a set of keys within a JWK Set during key rollover.  The structure of the \"kid\" value is unspecified.  When \"kid\" values are used within a JWK Set, different keys within the JWK Set SHOULD use distinct \"kid\" values.  (One example in which different keys might use the same \"kid\" value is if they have different \"kty\" (key type) values but are considered to be equivalent alternatives by the application using them.)  The \"kid\" value is a case-sensitive string.
	Kid string `json:"kid"`
	// The \"kty\" (key type) parameter identifies the cryptographic algorithm family used with the key, such as \"RSA\" or \"EC\". \"kty\" values should either be registered in the IANA \"JSON Web Key Types\" registry established by [JWA] or be a value that contains a Collision- Resistant Name.  The \"kty\" value is a case-sensitive string.
	Kty string `json:"kty"`
	N *string `json:"n,omitempty"`
	P *string `json:"p,omitempty"`
	Q *string `json:"q,omitempty"`
	Qi *string `json:"qi,omitempty"`
	// Use (\"public key use\") identifies the intended use of the public key. The \"use\" parameter is employed to indicate whether a public key is used for encrypting data or verifying the signature on data. Values are commonly \"sig\" (signature) or \"enc\" (encryption).
	Use string `json:"use"`
	X *string `json:"x,omitempty"`
	// The \"x5c\" (X.509 certificate chain) parameter contains a chain of one or more PKIX certificates [RFC5280].  The certificate chain is represented as a JSON array of certificate value strings.  Each string in the array is a base64-encoded (Section 4 of [RFC4648] -- not base64url-encoded) DER [ITU.X690.1994] PKIX certificate value. The PKIX certificate containing the key value MUST be the first certificate.
	X5c []string `json:"x5c,omitempty"`
	Y *string `json:"y,omitempty"`
}

// NewJsonWebKey instantiates a new JsonWebKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonWebKey(alg string, kid string, kty string, use string) *JsonWebKey {
	this := JsonWebKey{}
	this.Alg = alg
	this.Kid = kid
	this.Kty = kty
	this.Use = use
	return &this
}

// NewJsonWebKeyWithDefaults instantiates a new JsonWebKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonWebKeyWithDefaults() *JsonWebKey {
	this := JsonWebKey{}
	return &this
}

// GetAlg returns the Alg field value
func (o *JsonWebKey) GetAlg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alg
}

// GetAlgOk returns a tuple with the Alg field value
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetAlgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alg, true
}

// SetAlg sets field value
func (o *JsonWebKey) SetAlg(v string) {
	o.Alg = v
}

// GetCrv returns the Crv field value if set, zero value otherwise.
func (o *JsonWebKey) GetCrv() string {
	if o == nil || o.Crv == nil {
		var ret string
		return ret
	}
	return *o.Crv
}

// GetCrvOk returns a tuple with the Crv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetCrvOk() (*string, bool) {
	if o == nil || o.Crv == nil {
		return nil, false
	}
	return o.Crv, true
}

// HasCrv returns a boolean if a field has been set.
func (o *JsonWebKey) HasCrv() bool {
	if o != nil && o.Crv != nil {
		return true
	}

	return false
}

// SetCrv gets a reference to the given string and assigns it to the Crv field.
func (o *JsonWebKey) SetCrv(v string) {
	o.Crv = &v
}

// GetD returns the D field value if set, zero value otherwise.
func (o *JsonWebKey) GetD() string {
	if o == nil || o.D == nil {
		var ret string
		return ret
	}
	return *o.D
}

// GetDOk returns a tuple with the D field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetDOk() (*string, bool) {
	if o == nil || o.D == nil {
		return nil, false
	}
	return o.D, true
}

// HasD returns a boolean if a field has been set.
func (o *JsonWebKey) HasD() bool {
	if o != nil && o.D != nil {
		return true
	}

	return false
}

// SetD gets a reference to the given string and assigns it to the D field.
func (o *JsonWebKey) SetD(v string) {
	o.D = &v
}

// GetDp returns the Dp field value if set, zero value otherwise.
func (o *JsonWebKey) GetDp() string {
	if o == nil || o.Dp == nil {
		var ret string
		return ret
	}
	return *o.Dp
}

// GetDpOk returns a tuple with the Dp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetDpOk() (*string, bool) {
	if o == nil || o.Dp == nil {
		return nil, false
	}
	return o.Dp, true
}

// HasDp returns a boolean if a field has been set.
func (o *JsonWebKey) HasDp() bool {
	if o != nil && o.Dp != nil {
		return true
	}

	return false
}

// SetDp gets a reference to the given string and assigns it to the Dp field.
func (o *JsonWebKey) SetDp(v string) {
	o.Dp = &v
}

// GetDq returns the Dq field value if set, zero value otherwise.
func (o *JsonWebKey) GetDq() string {
	if o == nil || o.Dq == nil {
		var ret string
		return ret
	}
	return *o.Dq
}

// GetDqOk returns a tuple with the Dq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetDqOk() (*string, bool) {
	if o == nil || o.Dq == nil {
		return nil, false
	}
	return o.Dq, true
}

// HasDq returns a boolean if a field has been set.
func (o *JsonWebKey) HasDq() bool {
	if o != nil && o.Dq != nil {
		return true
	}

	return false
}

// SetDq gets a reference to the given string and assigns it to the Dq field.
func (o *JsonWebKey) SetDq(v string) {
	o.Dq = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *JsonWebKey) GetE() string {
	if o == nil || o.E == nil {
		var ret string
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetEOk() (*string, bool) {
	if o == nil || o.E == nil {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *JsonWebKey) HasE() bool {
	if o != nil && o.E != nil {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *JsonWebKey) SetE(v string) {
	o.E = &v
}

// GetK returns the K field value if set, zero value otherwise.
func (o *JsonWebKey) GetK() string {
	if o == nil || o.K == nil {
		var ret string
		return ret
	}
	return *o.K
}

// GetKOk returns a tuple with the K field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetKOk() (*string, bool) {
	if o == nil || o.K == nil {
		return nil, false
	}
	return o.K, true
}

// HasK returns a boolean if a field has been set.
func (o *JsonWebKey) HasK() bool {
	if o != nil && o.K != nil {
		return true
	}

	return false
}

// SetK gets a reference to the given string and assigns it to the K field.
func (o *JsonWebKey) SetK(v string) {
	o.K = &v
}

// GetKid returns the Kid field value
func (o *JsonWebKey) GetKid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kid
}

// GetKidOk returns a tuple with the Kid field value
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetKidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kid, true
}

// SetKid sets field value
func (o *JsonWebKey) SetKid(v string) {
	o.Kid = v
}

// GetKty returns the Kty field value
func (o *JsonWebKey) GetKty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kty
}

// GetKtyOk returns a tuple with the Kty field value
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetKtyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kty, true
}

// SetKty sets field value
func (o *JsonWebKey) SetKty(v string) {
	o.Kty = v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *JsonWebKey) GetN() string {
	if o == nil || o.N == nil {
		var ret string
		return ret
	}
	return *o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetNOk() (*string, bool) {
	if o == nil || o.N == nil {
		return nil, false
	}
	return o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *JsonWebKey) HasN() bool {
	if o != nil && o.N != nil {
		return true
	}

	return false
}

// SetN gets a reference to the given string and assigns it to the N field.
func (o *JsonWebKey) SetN(v string) {
	o.N = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *JsonWebKey) GetP() string {
	if o == nil || o.P == nil {
		var ret string
		return ret
	}
	return *o.P
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetPOk() (*string, bool) {
	if o == nil || o.P == nil {
		return nil, false
	}
	return o.P, true
}

// HasP returns a boolean if a field has been set.
func (o *JsonWebKey) HasP() bool {
	if o != nil && o.P != nil {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *JsonWebKey) SetP(v string) {
	o.P = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *JsonWebKey) GetQ() string {
	if o == nil || o.Q == nil {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetQOk() (*string, bool) {
	if o == nil || o.Q == nil {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *JsonWebKey) HasQ() bool {
	if o != nil && o.Q != nil {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *JsonWebKey) SetQ(v string) {
	o.Q = &v
}

// GetQi returns the Qi field value if set, zero value otherwise.
func (o *JsonWebKey) GetQi() string {
	if o == nil || o.Qi == nil {
		var ret string
		return ret
	}
	return *o.Qi
}

// GetQiOk returns a tuple with the Qi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetQiOk() (*string, bool) {
	if o == nil || o.Qi == nil {
		return nil, false
	}
	return o.Qi, true
}

// HasQi returns a boolean if a field has been set.
func (o *JsonWebKey) HasQi() bool {
	if o != nil && o.Qi != nil {
		return true
	}

	return false
}

// SetQi gets a reference to the given string and assigns it to the Qi field.
func (o *JsonWebKey) SetQi(v string) {
	o.Qi = &v
}

// GetUse returns the Use field value
func (o *JsonWebKey) GetUse() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Use
}

// GetUseOk returns a tuple with the Use field value
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetUseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Use, true
}

// SetUse sets field value
func (o *JsonWebKey) SetUse(v string) {
	o.Use = v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *JsonWebKey) GetX() string {
	if o == nil || o.X == nil {
		var ret string
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetXOk() (*string, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *JsonWebKey) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given string and assigns it to the X field.
func (o *JsonWebKey) SetX(v string) {
	o.X = &v
}

// GetX5c returns the X5c field value if set, zero value otherwise.
func (o *JsonWebKey) GetX5c() []string {
	if o == nil || o.X5c == nil {
		var ret []string
		return ret
	}
	return o.X5c
}

// GetX5cOk returns a tuple with the X5c field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetX5cOk() ([]string, bool) {
	if o == nil || o.X5c == nil {
		return nil, false
	}
	return o.X5c, true
}

// HasX5c returns a boolean if a field has been set.
func (o *JsonWebKey) HasX5c() bool {
	if o != nil && o.X5c != nil {
		return true
	}

	return false
}

// SetX5c gets a reference to the given []string and assigns it to the X5c field.
func (o *JsonWebKey) SetX5c(v []string) {
	o.X5c = v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *JsonWebKey) GetY() string {
	if o == nil || o.Y == nil {
		var ret string
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetYOk() (*string, bool) {
	if o == nil || o.Y == nil {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *JsonWebKey) HasY() bool {
	if o != nil && o.Y != nil {
		return true
	}

	return false
}

// SetY gets a reference to the given string and assigns it to the Y field.
func (o *JsonWebKey) SetY(v string) {
	o.Y = &v
}

func (o JsonWebKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["alg"] = o.Alg
	}
	if o.Crv != nil {
		toSerialize["crv"] = o.Crv
	}
	if o.D != nil {
		toSerialize["d"] = o.D
	}
	if o.Dp != nil {
		toSerialize["dp"] = o.Dp
	}
	if o.Dq != nil {
		toSerialize["dq"] = o.Dq
	}
	if o.E != nil {
		toSerialize["e"] = o.E
	}
	if o.K != nil {
		toSerialize["k"] = o.K
	}
	if true {
		toSerialize["kid"] = o.Kid
	}
	if true {
		toSerialize["kty"] = o.Kty
	}
	if o.N != nil {
		toSerialize["n"] = o.N
	}
	if o.P != nil {
		toSerialize["p"] = o.P
	}
	if o.Q != nil {
		toSerialize["q"] = o.Q
	}
	if o.Qi != nil {
		toSerialize["qi"] = o.Qi
	}
	if true {
		toSerialize["use"] = o.Use
	}
	if o.X != nil {
		toSerialize["x"] = o.X
	}
	if o.X5c != nil {
		toSerialize["x5c"] = o.X5c
	}
	if o.Y != nil {
		toSerialize["y"] = o.Y
	}
	return json.Marshal(toSerialize)
}

type NullableJsonWebKey struct {
	value *JsonWebKey
	isSet bool
}

func (v NullableJsonWebKey) Get() *JsonWebKey {
	return v.value
}

func (v *NullableJsonWebKey) Set(val *JsonWebKey) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonWebKey) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonWebKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonWebKey(val *JsonWebKey) *NullableJsonWebKey {
	return &NullableJsonWebKey{value: val, isSet: true}
}

func (v NullableJsonWebKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonWebKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


