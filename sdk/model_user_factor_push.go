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
	"reflect"
	"strings"
	"time"
)

// UserFactorPush struct for UserFactorPush
type UserFactorPush struct {
	UserFactor
	// Timestamp indicating when the Factor verification attempt expires
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// Result of a Factor verification attempt
	FactorResult         *string                `json:"factorResult,omitempty"`
	Profile              *UserFactorPushProfile `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorPush UserFactorPush

// NewUserFactorPush instantiates a new UserFactorPush object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorPush() *UserFactorPush {
	this := UserFactorPush{}
	return &this
}

// NewUserFactorPushWithDefaults instantiates a new UserFactorPush object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorPushWithDefaults() *UserFactorPush {
	this := UserFactorPush{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *UserFactorPush) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPush) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *UserFactorPush) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *UserFactorPush) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetFactorResult returns the FactorResult field value if set, zero value otherwise.
func (o *UserFactorPush) GetFactorResult() string {
	if o == nil || o.FactorResult == nil {
		var ret string
		return ret
	}
	return *o.FactorResult
}

// GetFactorResultOk returns a tuple with the FactorResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPush) GetFactorResultOk() (*string, bool) {
	if o == nil || o.FactorResult == nil {
		return nil, false
	}
	return o.FactorResult, true
}

// HasFactorResult returns a boolean if a field has been set.
func (o *UserFactorPush) HasFactorResult() bool {
	if o != nil && o.FactorResult != nil {
		return true
	}

	return false
}

// SetFactorResult gets a reference to the given string and assigns it to the FactorResult field.
func (o *UserFactorPush) SetFactorResult(v string) {
	o.FactorResult = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UserFactorPush) GetProfile() UserFactorPushProfile {
	if o == nil || o.Profile == nil {
		var ret UserFactorPushProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPush) GetProfileOk() (*UserFactorPushProfile, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UserFactorPush) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given UserFactorPushProfile and assigns it to the Profile field.
func (o *UserFactorPush) SetProfile(v UserFactorPushProfile) {
	o.Profile = &v
}

func (o UserFactorPush) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedUserFactor, errUserFactor := json.Marshal(o.UserFactor)
	if errUserFactor != nil {
		return []byte{}, errUserFactor
	}
	errUserFactor = json.Unmarshal([]byte(serializedUserFactor), &toSerialize)
	if errUserFactor != nil {
		return []byte{}, errUserFactor
	}
	if o.ExpiresAt != nil {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if o.FactorResult != nil {
		toSerialize["factorResult"] = o.FactorResult
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserFactorPush) UnmarshalJSON(bytes []byte) (err error) {
	type UserFactorPushWithoutEmbeddedStruct struct {
		// Timestamp indicating when the Factor verification attempt expires
		ExpiresAt *time.Time `json:"expiresAt,omitempty"`
		// Result of a Factor verification attempt
		FactorResult *string                `json:"factorResult,omitempty"`
		Profile      *UserFactorPushProfile `json:"profile,omitempty"`
	}

	varUserFactorPushWithoutEmbeddedStruct := UserFactorPushWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varUserFactorPushWithoutEmbeddedStruct)
	if err == nil {
		varUserFactorPush := _UserFactorPush{}
		varUserFactorPush.ExpiresAt = varUserFactorPushWithoutEmbeddedStruct.ExpiresAt
		varUserFactorPush.FactorResult = varUserFactorPushWithoutEmbeddedStruct.FactorResult
		varUserFactorPush.Profile = varUserFactorPushWithoutEmbeddedStruct.Profile
		*o = UserFactorPush(varUserFactorPush)
	} else {
		return err
	}

	varUserFactorPush := _UserFactorPush{}

	err = json.Unmarshal(bytes, &varUserFactorPush)
	if err == nil {
		o.UserFactor = varUserFactorPush.UserFactor
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "expiresAt")
		delete(additionalProperties, "factorResult")
		delete(additionalProperties, "profile")

		// remove fields from embedded structs
		reflectUserFactor := reflect.ValueOf(o.UserFactor)
		for i := 0; i < reflectUserFactor.Type().NumField(); i++ {
			t := reflectUserFactor.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserFactorPush struct {
	value *UserFactorPush
	isSet bool
}

func (v NullableUserFactorPush) Get() *UserFactorPush {
	return v.value
}

func (v *NullableUserFactorPush) Set(val *UserFactorPush) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorPush) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorPush) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorPush(val *UserFactorPush) *NullableUserFactorPush {
	return &NullableUserFactorPush{value: val, isSet: true}
}

func (v NullableUserFactorPush) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorPush) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}