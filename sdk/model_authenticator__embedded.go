/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

API version: 5.1.0
Contact: devex-public@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// AuthenticatorEmbedded struct for AuthenticatorEmbedded
type AuthenticatorEmbedded struct {
	Methods              []AuthenticatorMethodBase `json:"methods,omitempty"`
	Policies             []Policy                  `json:"policies,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorEmbedded AuthenticatorEmbedded

// NewAuthenticatorEmbedded instantiates a new AuthenticatorEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorEmbedded() *AuthenticatorEmbedded {
	this := AuthenticatorEmbedded{}
	return &this
}

// NewAuthenticatorEmbeddedWithDefaults instantiates a new AuthenticatorEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorEmbeddedWithDefaults() *AuthenticatorEmbedded {
	this := AuthenticatorEmbedded{}
	return &this
}

// GetMethods returns the Methods field value if set, zero value otherwise.
func (o *AuthenticatorEmbedded) GetMethods() []AuthenticatorMethodBase {
	if o == nil || o.Methods == nil {
		var ret []AuthenticatorMethodBase
		return ret
	}
	return o.Methods
}

// GetMethodsOk returns a tuple with the Methods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEmbedded) GetMethodsOk() ([]AuthenticatorMethodBase, bool) {
	if o == nil || o.Methods == nil {
		return nil, false
	}
	return o.Methods, true
}

// HasMethods returns a boolean if a field has been set.
func (o *AuthenticatorEmbedded) HasMethods() bool {
	if o != nil && o.Methods != nil {
		return true
	}

	return false
}

// SetMethods gets a reference to the given []AuthenticatorMethodBase and assigns it to the Methods field.
func (o *AuthenticatorEmbedded) SetMethods(v []AuthenticatorMethodBase) {
	o.Methods = v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *AuthenticatorEmbedded) GetPolicies() []Policy {
	if o == nil || o.Policies == nil {
		var ret []Policy
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEmbedded) GetPoliciesOk() ([]Policy, bool) {
	if o == nil || o.Policies == nil {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *AuthenticatorEmbedded) HasPolicies() bool {
	if o != nil && o.Policies != nil {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []Policy and assigns it to the Policies field.
func (o *AuthenticatorEmbedded) SetPolicies(v []Policy) {
	o.Policies = v
}

func (o AuthenticatorEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Methods != nil {
		toSerialize["methods"] = o.Methods
	}
	if o.Policies != nil {
		toSerialize["policies"] = o.Policies
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticatorEmbedded) UnmarshalJSON(bytes []byte) (err error) {
	varAuthenticatorEmbedded := _AuthenticatorEmbedded{}

	err = json.Unmarshal(bytes, &varAuthenticatorEmbedded)
	if err == nil {
		*o = AuthenticatorEmbedded(varAuthenticatorEmbedded)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "methods")
		delete(additionalProperties, "policies")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthenticatorEmbedded struct {
	value *AuthenticatorEmbedded
	isSet bool
}

func (v NullableAuthenticatorEmbedded) Get() *AuthenticatorEmbedded {
	return v.value
}

func (v *NullableAuthenticatorEmbedded) Set(val *AuthenticatorEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorEmbedded(val *AuthenticatorEmbedded) *NullableAuthenticatorEmbedded {
	return &NullableAuthenticatorEmbedded{value: val, isSet: true}
}

func (v NullableAuthenticatorEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}