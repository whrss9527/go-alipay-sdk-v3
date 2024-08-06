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

// checks if the AlipayCommerceEducateCampusInstitutionsQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayCommerceEducateCampusInstitutionsQueryResponseModel{}

// AlipayCommerceEducateCampusInstitutionsQueryResponseModel struct for AlipayCommerceEducateCampusInstitutionsQueryResponseModel
type AlipayCommerceEducateCampusInstitutionsQueryResponseModel struct {
	SchoolInfo *SchoolSimpleInfo `json:"school_info,omitempty"`
	// 学校基础信息列表
	SchoolInfoList []SchoolBaseInfo `json:"school_info_list,omitempty"`
}

// NewAlipayCommerceEducateCampusInstitutionsQueryResponseModel instantiates a new AlipayCommerceEducateCampusInstitutionsQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayCommerceEducateCampusInstitutionsQueryResponseModel() *AlipayCommerceEducateCampusInstitutionsQueryResponseModel {
	this := AlipayCommerceEducateCampusInstitutionsQueryResponseModel{}
	return &this
}

// NewAlipayCommerceEducateCampusInstitutionsQueryResponseModelWithDefaults instantiates a new AlipayCommerceEducateCampusInstitutionsQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayCommerceEducateCampusInstitutionsQueryResponseModelWithDefaults() *AlipayCommerceEducateCampusInstitutionsQueryResponseModel {
	this := AlipayCommerceEducateCampusInstitutionsQueryResponseModel{}
	return &this
}

// GetSchoolInfo returns the SchoolInfo field value if set, zero value otherwise.
func (o *AlipayCommerceEducateCampusInstitutionsQueryResponseModel) GetSchoolInfo() SchoolSimpleInfo {
	if o == nil || IsNil(o.SchoolInfo) {
		var ret SchoolSimpleInfo
		return ret
	}
	return *o.SchoolInfo
}

// GetSchoolInfoOk returns a tuple with the SchoolInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEducateCampusInstitutionsQueryResponseModel) GetSchoolInfoOk() (*SchoolSimpleInfo, bool) {
	if o == nil || IsNil(o.SchoolInfo) {
		return nil, false
	}
	return o.SchoolInfo, true
}

// HasSchoolInfo returns a boolean if a field has been set.
func (o *AlipayCommerceEducateCampusInstitutionsQueryResponseModel) HasSchoolInfo() bool {
	if o != nil && !IsNil(o.SchoolInfo) {
		return true
	}

	return false
}

// SetSchoolInfo gets a reference to the given SchoolSimpleInfo and assigns it to the SchoolInfo field.
func (o *AlipayCommerceEducateCampusInstitutionsQueryResponseModel) SetSchoolInfo(v SchoolSimpleInfo) {
	o.SchoolInfo = &v
}

// GetSchoolInfoList returns the SchoolInfoList field value if set, zero value otherwise.
func (o *AlipayCommerceEducateCampusInstitutionsQueryResponseModel) GetSchoolInfoList() []SchoolBaseInfo {
	if o == nil || IsNil(o.SchoolInfoList) {
		var ret []SchoolBaseInfo
		return ret
	}
	return o.SchoolInfoList
}

// GetSchoolInfoListOk returns a tuple with the SchoolInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEducateCampusInstitutionsQueryResponseModel) GetSchoolInfoListOk() ([]SchoolBaseInfo, bool) {
	if o == nil || IsNil(o.SchoolInfoList) {
		return nil, false
	}
	return o.SchoolInfoList, true
}

// HasSchoolInfoList returns a boolean if a field has been set.
func (o *AlipayCommerceEducateCampusInstitutionsQueryResponseModel) HasSchoolInfoList() bool {
	if o != nil && !IsNil(o.SchoolInfoList) {
		return true
	}

	return false
}

// SetSchoolInfoList gets a reference to the given []SchoolBaseInfo and assigns it to the SchoolInfoList field.
func (o *AlipayCommerceEducateCampusInstitutionsQueryResponseModel) SetSchoolInfoList(v []SchoolBaseInfo) {
	o.SchoolInfoList = v
}

func (o AlipayCommerceEducateCampusInstitutionsQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayCommerceEducateCampusInstitutionsQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SchoolInfo) {
		toSerialize["school_info"] = o.SchoolInfo
	}
	if !IsNil(o.SchoolInfoList) {
		toSerialize["school_info_list"] = o.SchoolInfoList
	}
	return toSerialize, nil
}

type NullableAlipayCommerceEducateCampusInstitutionsQueryResponseModel struct {
	value *AlipayCommerceEducateCampusInstitutionsQueryResponseModel
	isSet bool
}

func (v NullableAlipayCommerceEducateCampusInstitutionsQueryResponseModel) Get() *AlipayCommerceEducateCampusInstitutionsQueryResponseModel {
	return v.value
}

func (v *NullableAlipayCommerceEducateCampusInstitutionsQueryResponseModel) Set(val *AlipayCommerceEducateCampusInstitutionsQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceEducateCampusInstitutionsQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceEducateCampusInstitutionsQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceEducateCampusInstitutionsQueryResponseModel(val *AlipayCommerceEducateCampusInstitutionsQueryResponseModel) *NullableAlipayCommerceEducateCampusInstitutionsQueryResponseModel {
	return &NullableAlipayCommerceEducateCampusInstitutionsQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayCommerceEducateCampusInstitutionsQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceEducateCampusInstitutionsQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

