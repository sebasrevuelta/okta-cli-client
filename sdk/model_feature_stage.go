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

// FeatureStage Current release cycle stage of a feature  If a feature's stage value is `EA`, the state is `null` and not returned. If the value is `BETA`, the state is `OPEN` or `CLOSED` depending on whether the `BETA` feature is manageable.  > **Note:** If a feature's stage is `OPEN BETA`, you can update it only in Preview cells. If a feature's stage is `CLOSED BETA`, you can disable it only in Preview cells.
type FeatureStage struct {
	// Indicates the release state of the feature
	State *string `json:"state,omitempty"`
	// Current release stage of the feature
	Value                *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FeatureStage FeatureStage

// NewFeatureStage instantiates a new FeatureStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureStage() *FeatureStage {
	this := FeatureStage{}
	return &this
}

// NewFeatureStageWithDefaults instantiates a new FeatureStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureStageWithDefaults() *FeatureStage {
	this := FeatureStage{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *FeatureStage) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureStage) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *FeatureStage) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *FeatureStage) SetState(v string) {
	o.State = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FeatureStage) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureStage) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FeatureStage) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *FeatureStage) SetValue(v string) {
	o.Value = &v
}

func (o FeatureStage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FeatureStage) UnmarshalJSON(bytes []byte) (err error) {
	varFeatureStage := _FeatureStage{}

	err = json.Unmarshal(bytes, &varFeatureStage)
	if err == nil {
		*o = FeatureStage(varFeatureStage)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "state")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableFeatureStage struct {
	value *FeatureStage
	isSet bool
}

func (v NullableFeatureStage) Get() *FeatureStage {
	return v.value
}

func (v *NullableFeatureStage) Set(val *FeatureStage) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureStage) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureStage(val *FeatureStage) *NullableFeatureStage {
	return &NullableFeatureStage{value: val, isSet: true}
}

func (v NullableFeatureStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}