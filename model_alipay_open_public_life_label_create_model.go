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

// checks if the AlipayOpenPublicLifeLabelCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicLifeLabelCreateModel{}

// AlipayOpenPublicLifeLabelCreateModel struct for AlipayOpenPublicLifeLabelCreateModel
type AlipayOpenPublicLifeLabelCreateModel struct {
	// 标签值类型，不填默认为 string 类型。 注意：目前只支持 string（字符串类型）。
	DataType *string `json:"data_type,omitempty"`
	// 自定义标签名。 注意：每个生活号最多创建 100 个自定义标签。
	LabelName *string `json:"label_name,omitempty"`
}

// NewAlipayOpenPublicLifeLabelCreateModel instantiates a new AlipayOpenPublicLifeLabelCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicLifeLabelCreateModel() *AlipayOpenPublicLifeLabelCreateModel {
	this := AlipayOpenPublicLifeLabelCreateModel{}
	return &this
}

// NewAlipayOpenPublicLifeLabelCreateModelWithDefaults instantiates a new AlipayOpenPublicLifeLabelCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicLifeLabelCreateModelWithDefaults() *AlipayOpenPublicLifeLabelCreateModel {
	this := AlipayOpenPublicLifeLabelCreateModel{}
	return &this
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *AlipayOpenPublicLifeLabelCreateModel) GetDataType() string {
	if o == nil || IsNil(o.DataType) {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicLifeLabelCreateModel) GetDataTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DataType) {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *AlipayOpenPublicLifeLabelCreateModel) HasDataType() bool {
	if o != nil && !IsNil(o.DataType) {
		return true
	}

	return false
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *AlipayOpenPublicLifeLabelCreateModel) SetDataType(v string) {
	o.DataType = &v
}

// GetLabelName returns the LabelName field value if set, zero value otherwise.
func (o *AlipayOpenPublicLifeLabelCreateModel) GetLabelName() string {
	if o == nil || IsNil(o.LabelName) {
		var ret string
		return ret
	}
	return *o.LabelName
}

// GetLabelNameOk returns a tuple with the LabelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicLifeLabelCreateModel) GetLabelNameOk() (*string, bool) {
	if o == nil || IsNil(o.LabelName) {
		return nil, false
	}
	return o.LabelName, true
}

// HasLabelName returns a boolean if a field has been set.
func (o *AlipayOpenPublicLifeLabelCreateModel) HasLabelName() bool {
	if o != nil && !IsNil(o.LabelName) {
		return true
	}

	return false
}

// SetLabelName gets a reference to the given string and assigns it to the LabelName field.
func (o *AlipayOpenPublicLifeLabelCreateModel) SetLabelName(v string) {
	o.LabelName = &v
}

func (o AlipayOpenPublicLifeLabelCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicLifeLabelCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataType) {
		toSerialize["data_type"] = o.DataType
	}
	if !IsNil(o.LabelName) {
		toSerialize["label_name"] = o.LabelName
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicLifeLabelCreateModel struct {
	value *AlipayOpenPublicLifeLabelCreateModel
	isSet bool
}

func (v NullableAlipayOpenPublicLifeLabelCreateModel) Get() *AlipayOpenPublicLifeLabelCreateModel {
	return v.value
}

func (v *NullableAlipayOpenPublicLifeLabelCreateModel) Set(val *AlipayOpenPublicLifeLabelCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicLifeLabelCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicLifeLabelCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicLifeLabelCreateModel(val *AlipayOpenPublicLifeLabelCreateModel) *NullableAlipayOpenPublicLifeLabelCreateModel {
	return &NullableAlipayOpenPublicLifeLabelCreateModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicLifeLabelCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicLifeLabelCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

