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

// BouncesRemoveListError struct for BouncesRemoveListError
type BouncesRemoveListError struct {
	EmailAddress         *string `json:"emailAddress,omitempty"`
	Reason               *string `json:"reason,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BouncesRemoveListError BouncesRemoveListError

// NewBouncesRemoveListError instantiates a new BouncesRemoveListError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBouncesRemoveListError() *BouncesRemoveListError {
	this := BouncesRemoveListError{}
	return &this
}

// NewBouncesRemoveListErrorWithDefaults instantiates a new BouncesRemoveListError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBouncesRemoveListErrorWithDefaults() *BouncesRemoveListError {
	this := BouncesRemoveListError{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *BouncesRemoveListError) GetEmailAddress() string {
	if o == nil || o.EmailAddress == nil {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BouncesRemoveListError) GetEmailAddressOk() (*string, bool) {
	if o == nil || o.EmailAddress == nil {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *BouncesRemoveListError) HasEmailAddress() bool {
	if o != nil && o.EmailAddress != nil {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *BouncesRemoveListError) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *BouncesRemoveListError) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BouncesRemoveListError) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *BouncesRemoveListError) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *BouncesRemoveListError) SetReason(v string) {
	o.Reason = &v
}

func (o BouncesRemoveListError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EmailAddress != nil {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BouncesRemoveListError) UnmarshalJSON(bytes []byte) (err error) {
	varBouncesRemoveListError := _BouncesRemoveListError{}

	err = json.Unmarshal(bytes, &varBouncesRemoveListError)
	if err == nil {
		*o = BouncesRemoveListError(varBouncesRemoveListError)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "emailAddress")
		delete(additionalProperties, "reason")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableBouncesRemoveListError struct {
	value *BouncesRemoveListError
	isSet bool
}

func (v NullableBouncesRemoveListError) Get() *BouncesRemoveListError {
	return v.value
}

func (v *NullableBouncesRemoveListError) Set(val *BouncesRemoveListError) {
	v.value = val
	v.isSet = true
}

func (v NullableBouncesRemoveListError) IsSet() bool {
	return v.isSet
}

func (v *NullableBouncesRemoveListError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBouncesRemoveListError(val *BouncesRemoveListError) *NullableBouncesRemoveListError {
	return &NullableBouncesRemoveListError{value: val, isSet: true}
}

func (v NullableBouncesRemoveListError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBouncesRemoveListError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}