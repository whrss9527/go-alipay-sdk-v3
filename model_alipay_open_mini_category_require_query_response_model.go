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

// checks if the AlipayOpenMiniCategoryRequireQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniCategoryRequireQueryResponseModel{}

// AlipayOpenMiniCategoryRequireQueryResponseModel struct for AlipayOpenMiniCategoryRequireQueryResponseModel
type AlipayOpenMiniCategoryRequireQueryResponseModel struct {
	// 类目对应资质列表
	CategoryRequireInfoList []CategoryRequireInfo `json:"category_require_info_list,omitempty"`
}

// NewAlipayOpenMiniCategoryRequireQueryResponseModel instantiates a new AlipayOpenMiniCategoryRequireQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniCategoryRequireQueryResponseModel() *AlipayOpenMiniCategoryRequireQueryResponseModel {
	this := AlipayOpenMiniCategoryRequireQueryResponseModel{}
	return &this
}

// NewAlipayOpenMiniCategoryRequireQueryResponseModelWithDefaults instantiates a new AlipayOpenMiniCategoryRequireQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniCategoryRequireQueryResponseModelWithDefaults() *AlipayOpenMiniCategoryRequireQueryResponseModel {
	this := AlipayOpenMiniCategoryRequireQueryResponseModel{}
	return &this
}

// GetCategoryRequireInfoList returns the CategoryRequireInfoList field value if set, zero value otherwise.
func (o *AlipayOpenMiniCategoryRequireQueryResponseModel) GetCategoryRequireInfoList() []CategoryRequireInfo {
	if o == nil || IsNil(o.CategoryRequireInfoList) {
		var ret []CategoryRequireInfo
		return ret
	}
	return o.CategoryRequireInfoList
}

// GetCategoryRequireInfoListOk returns a tuple with the CategoryRequireInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniCategoryRequireQueryResponseModel) GetCategoryRequireInfoListOk() ([]CategoryRequireInfo, bool) {
	if o == nil || IsNil(o.CategoryRequireInfoList) {
		return nil, false
	}
	return o.CategoryRequireInfoList, true
}

// HasCategoryRequireInfoList returns a boolean if a field has been set.
func (o *AlipayOpenMiniCategoryRequireQueryResponseModel) HasCategoryRequireInfoList() bool {
	if o != nil && !IsNil(o.CategoryRequireInfoList) {
		return true
	}

	return false
}

// SetCategoryRequireInfoList gets a reference to the given []CategoryRequireInfo and assigns it to the CategoryRequireInfoList field.
func (o *AlipayOpenMiniCategoryRequireQueryResponseModel) SetCategoryRequireInfoList(v []CategoryRequireInfo) {
	o.CategoryRequireInfoList = v
}

func (o AlipayOpenMiniCategoryRequireQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniCategoryRequireQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CategoryRequireInfoList) {
		toSerialize["category_require_info_list"] = o.CategoryRequireInfoList
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniCategoryRequireQueryResponseModel struct {
	value *AlipayOpenMiniCategoryRequireQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenMiniCategoryRequireQueryResponseModel) Get() *AlipayOpenMiniCategoryRequireQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenMiniCategoryRequireQueryResponseModel) Set(val *AlipayOpenMiniCategoryRequireQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniCategoryRequireQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniCategoryRequireQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniCategoryRequireQueryResponseModel(val *AlipayOpenMiniCategoryRequireQueryResponseModel) *NullableAlipayOpenMiniCategoryRequireQueryResponseModel {
	return &NullableAlipayOpenMiniCategoryRequireQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniCategoryRequireQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniCategoryRequireQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


