/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alipay

import (
	"encoding/json"
)

// checks if the MaterialField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaterialField{}

// MaterialField struct for MaterialField
type MaterialField struct {
	// 素材字段名称，来源于素材规范定义的字段名称。注意一次素材提报中字段名称不能重复。
	FieldName *string `json:"field_name,omitempty"`
	// 素材字段值，素材字段提报的实际值，支持多值，具体的要求请查看素材规范或对应的产品文档。
	FieldValue []string `json:"field_value,omitempty"`
}

// NewMaterialField instantiates a new MaterialField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaterialField() *MaterialField {
	this := MaterialField{}
	return &this
}

// NewMaterialFieldWithDefaults instantiates a new MaterialField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaterialFieldWithDefaults() *MaterialField {
	this := MaterialField{}
	return &this
}

// GetFieldName returns the FieldName field value if set, zero value otherwise.
func (o *MaterialField) GetFieldName() string {
	if o == nil || IsNil(o.FieldName) {
		var ret string
		return ret
	}
	return *o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaterialField) GetFieldNameOk() (*string, bool) {
	if o == nil || IsNil(o.FieldName) {
		return nil, false
	}
	return o.FieldName, true
}

// HasFieldName returns a boolean if a field has been set.
func (o *MaterialField) HasFieldName() bool {
	if o != nil && !IsNil(o.FieldName) {
		return true
	}

	return false
}

// SetFieldName gets a reference to the given string and assigns it to the FieldName field.
func (o *MaterialField) SetFieldName(v string) {
	o.FieldName = &v
}

// GetFieldValue returns the FieldValue field value if set, zero value otherwise.
func (o *MaterialField) GetFieldValue() []string {
	if o == nil || IsNil(o.FieldValue) {
		var ret []string
		return ret
	}
	return o.FieldValue
}

// GetFieldValueOk returns a tuple with the FieldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaterialField) GetFieldValueOk() ([]string, bool) {
	if o == nil || IsNil(o.FieldValue) {
		return nil, false
	}
	return o.FieldValue, true
}

// HasFieldValue returns a boolean if a field has been set.
func (o *MaterialField) HasFieldValue() bool {
	if o != nil && !IsNil(o.FieldValue) {
		return true
	}

	return false
}

// SetFieldValue gets a reference to the given []string and assigns it to the FieldValue field.
func (o *MaterialField) SetFieldValue(v []string) {
	o.FieldValue = v
}

func (o MaterialField) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaterialField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FieldName) {
		toSerialize["field_name"] = o.FieldName
	}
	if !IsNil(o.FieldValue) {
		toSerialize["field_value"] = o.FieldValue
	}
	return toSerialize, nil
}

type NullableMaterialField struct {
	value *MaterialField
	isSet bool
}

func (v NullableMaterialField) Get() *MaterialField {
	return v.value
}

func (v *NullableMaterialField) Set(val *MaterialField) {
	v.value = val
	v.isSet = true
}

func (v NullableMaterialField) IsSet() bool {
	return v.isSet
}

func (v *NullableMaterialField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaterialField(val *MaterialField) *NullableMaterialField {
	return &NullableMaterialField{value: val, isSet: true}
}

func (v NullableMaterialField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaterialField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
