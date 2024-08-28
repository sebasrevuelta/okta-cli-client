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

// LogStreamPutSchema struct for LogStreamPutSchema
type LogStreamPutSchema struct {
	// Unique name for the Log Stream object
	Name string `json:"name"`
	// Specifies the streaming provider used  Supported providers:   * `aws_eventbridge` ([AWS EventBridge](https://aws.amazon.com/eventbridge))   * `splunk_cloud_logstreaming` ([Splunk Cloud](https://www.splunk.com/en_us/software/splunk-cloud-platform.html))  Select the provider type to see provider-specific configurations in the `settings` property:
	Type string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _LogStreamPutSchema LogStreamPutSchema

// NewLogStreamPutSchema instantiates a new LogStreamPutSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogStreamPutSchema(name string, type_ string) *LogStreamPutSchema {
	this := LogStreamPutSchema{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewLogStreamPutSchemaWithDefaults instantiates a new LogStreamPutSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogStreamPutSchemaWithDefaults() *LogStreamPutSchema {
	this := LogStreamPutSchema{}
	return &this
}

// GetName returns the Name field value
func (o *LogStreamPutSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LogStreamPutSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LogStreamPutSchema) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *LogStreamPutSchema) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogStreamPutSchema) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LogStreamPutSchema) SetType(v string) {
	o.Type = v
}

func (o LogStreamPutSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LogStreamPutSchema) UnmarshalJSON(bytes []byte) (err error) {
	varLogStreamPutSchema := _LogStreamPutSchema{}

	err = json.Unmarshal(bytes, &varLogStreamPutSchema)
	if err == nil {
		*o = LogStreamPutSchema(varLogStreamPutSchema)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLogStreamPutSchema struct {
	value *LogStreamPutSchema
	isSet bool
}

func (v NullableLogStreamPutSchema) Get() *LogStreamPutSchema {
	return v.value
}

func (v *NullableLogStreamPutSchema) Set(val *LogStreamPutSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableLogStreamPutSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableLogStreamPutSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogStreamPutSchema(val *LogStreamPutSchema) *NullableLogStreamPutSchema {
	return &NullableLogStreamPutSchema{value: val, isSet: true}
}

func (v NullableLogStreamPutSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogStreamPutSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

