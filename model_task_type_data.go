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

// checks if the TaskTypeData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskTypeData{}

// TaskTypeData struct for TaskTypeData
type TaskTypeData struct {
	// 商户数据回传的任务名称，供C端展示。当累计的数据类型为任务类型时，此为必传字段
	Name *string `json:"name,omitempty"`
}

// NewTaskTypeData instantiates a new TaskTypeData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskTypeData() *TaskTypeData {
	this := TaskTypeData{}
	return &this
}

// NewTaskTypeDataWithDefaults instantiates a new TaskTypeData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskTypeDataWithDefaults() *TaskTypeData {
	this := TaskTypeData{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TaskTypeData) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskTypeData) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TaskTypeData) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TaskTypeData) SetName(v string) {
	o.Name = &v
}

func (o TaskTypeData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskTypeData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableTaskTypeData struct {
	value *TaskTypeData
	isSet bool
}

func (v NullableTaskTypeData) Get() *TaskTypeData {
	return v.value
}

func (v *NullableTaskTypeData) Set(val *TaskTypeData) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskTypeData) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskTypeData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskTypeData(val *TaskTypeData) *NullableTaskTypeData {
	return &NullableTaskTypeData{value: val, isSet: true}
}

func (v NullableTaskTypeData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskTypeData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
