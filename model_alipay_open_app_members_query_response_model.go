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

// checks if the AlipayOpenAppMembersQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenAppMembersQueryResponseModel{}

// AlipayOpenAppMembersQueryResponseModel struct for AlipayOpenAppMembersQueryResponseModel
type AlipayOpenAppMembersQueryResponseModel struct {
	// 小程序成员模型
	AppMemberInfoList []AppMemberInfo `json:"app_member_info_list,omitempty"`
}

// NewAlipayOpenAppMembersQueryResponseModel instantiates a new AlipayOpenAppMembersQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenAppMembersQueryResponseModel() *AlipayOpenAppMembersQueryResponseModel {
	this := AlipayOpenAppMembersQueryResponseModel{}
	return &this
}

// NewAlipayOpenAppMembersQueryResponseModelWithDefaults instantiates a new AlipayOpenAppMembersQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenAppMembersQueryResponseModelWithDefaults() *AlipayOpenAppMembersQueryResponseModel {
	this := AlipayOpenAppMembersQueryResponseModel{}
	return &this
}

// GetAppMemberInfoList returns the AppMemberInfoList field value if set, zero value otherwise.
func (o *AlipayOpenAppMembersQueryResponseModel) GetAppMemberInfoList() []AppMemberInfo {
	if o == nil || IsNil(o.AppMemberInfoList) {
		var ret []AppMemberInfo
		return ret
	}
	return o.AppMemberInfoList
}

// GetAppMemberInfoListOk returns a tuple with the AppMemberInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAppMembersQueryResponseModel) GetAppMemberInfoListOk() ([]AppMemberInfo, bool) {
	if o == nil || IsNil(o.AppMemberInfoList) {
		return nil, false
	}
	return o.AppMemberInfoList, true
}

// HasAppMemberInfoList returns a boolean if a field has been set.
func (o *AlipayOpenAppMembersQueryResponseModel) HasAppMemberInfoList() bool {
	if o != nil && !IsNil(o.AppMemberInfoList) {
		return true
	}

	return false
}

// SetAppMemberInfoList gets a reference to the given []AppMemberInfo and assigns it to the AppMemberInfoList field.
func (o *AlipayOpenAppMembersQueryResponseModel) SetAppMemberInfoList(v []AppMemberInfo) {
	o.AppMemberInfoList = v
}

func (o AlipayOpenAppMembersQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenAppMembersQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppMemberInfoList) {
		toSerialize["app_member_info_list"] = o.AppMemberInfoList
	}
	return toSerialize, nil
}

type NullableAlipayOpenAppMembersQueryResponseModel struct {
	value *AlipayOpenAppMembersQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenAppMembersQueryResponseModel) Get() *AlipayOpenAppMembersQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenAppMembersQueryResponseModel) Set(val *AlipayOpenAppMembersQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenAppMembersQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenAppMembersQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenAppMembersQueryResponseModel(val *AlipayOpenAppMembersQueryResponseModel) *NullableAlipayOpenAppMembersQueryResponseModel {
	return &NullableAlipayOpenAppMembersQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenAppMembersQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenAppMembersQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

