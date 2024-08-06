/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the LabelContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LabelContext{}

// LabelContext struct for LabelContext
type LabelContext struct {
	A *LabelFilter `json:"a,omitempty"`
}

// NewLabelContext instantiates a new LabelContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelContext() *LabelContext {
	this := LabelContext{}
	return &this
}

// NewLabelContextWithDefaults instantiates a new LabelContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelContextWithDefaults() *LabelContext {
	this := LabelContext{}
	return &this
}

// GetA returns the A field value if set, zero value otherwise.
func (o *LabelContext) GetA() LabelFilter {
	if o == nil || IsNil(o.A) {
		var ret LabelFilter
		return ret
	}
	return *o.A
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelContext) GetAOk() (*LabelFilter, bool) {
	if o == nil || IsNil(o.A) {
		return nil, false
	}
	return o.A, true
}

// HasA returns a boolean if a field has been set.
func (o *LabelContext) HasA() bool {
	if o != nil && !IsNil(o.A) {
		return true
	}

	return false
}

// SetA gets a reference to the given LabelFilter and assigns it to the A field.
func (o *LabelContext) SetA(v LabelFilter) {
	o.A = &v
}

func (o LabelContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LabelContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.A) {
		toSerialize["a"] = o.A
	}
	return toSerialize, nil
}

type NullableLabelContext struct {
	value *LabelContext
	isSet bool
}

func (v NullableLabelContext) Get() *LabelContext {
	return v.value
}

func (v *NullableLabelContext) Set(val *LabelContext) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelContext) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelContext(val *LabelContext) *NullableLabelContext {
	return &NullableLabelContext{value: val, isSet: true}
}

func (v NullableLabelContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


