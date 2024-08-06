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

// checks if the AlipayMarketingActivityConsultResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingActivityConsultResponseModel{}

// AlipayMarketingActivityConsultResponseModel struct for AlipayMarketingActivityConsultResponseModel
type AlipayMarketingActivityConsultResponseModel struct {
	// 咨询后的活动结果信息
	ConsultResultInfoList []ConsultActivityResultInfo `json:"consult_result_info_list,omitempty"`
	// 领券的用户openId
	OpenId *string `json:"open_id,omitempty"`
	// 领券的用户uid
	UserId *string `json:"user_id,omitempty"`
}

// NewAlipayMarketingActivityConsultResponseModel instantiates a new AlipayMarketingActivityConsultResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingActivityConsultResponseModel() *AlipayMarketingActivityConsultResponseModel {
	this := AlipayMarketingActivityConsultResponseModel{}
	return &this
}

// NewAlipayMarketingActivityConsultResponseModelWithDefaults instantiates a new AlipayMarketingActivityConsultResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingActivityConsultResponseModelWithDefaults() *AlipayMarketingActivityConsultResponseModel {
	this := AlipayMarketingActivityConsultResponseModel{}
	return &this
}

// GetConsultResultInfoList returns the ConsultResultInfoList field value if set, zero value otherwise.
func (o *AlipayMarketingActivityConsultResponseModel) GetConsultResultInfoList() []ConsultActivityResultInfo {
	if o == nil || IsNil(o.ConsultResultInfoList) {
		var ret []ConsultActivityResultInfo
		return ret
	}
	return o.ConsultResultInfoList
}

// GetConsultResultInfoListOk returns a tuple with the ConsultResultInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityConsultResponseModel) GetConsultResultInfoListOk() ([]ConsultActivityResultInfo, bool) {
	if o == nil || IsNil(o.ConsultResultInfoList) {
		return nil, false
	}
	return o.ConsultResultInfoList, true
}

// HasConsultResultInfoList returns a boolean if a field has been set.
func (o *AlipayMarketingActivityConsultResponseModel) HasConsultResultInfoList() bool {
	if o != nil && !IsNil(o.ConsultResultInfoList) {
		return true
	}

	return false
}

// SetConsultResultInfoList gets a reference to the given []ConsultActivityResultInfo and assigns it to the ConsultResultInfoList field.
func (o *AlipayMarketingActivityConsultResponseModel) SetConsultResultInfoList(v []ConsultActivityResultInfo) {
	o.ConsultResultInfoList = v
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *AlipayMarketingActivityConsultResponseModel) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityConsultResponseModel) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *AlipayMarketingActivityConsultResponseModel) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *AlipayMarketingActivityConsultResponseModel) SetOpenId(v string) {
	o.OpenId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AlipayMarketingActivityConsultResponseModel) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityConsultResponseModel) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AlipayMarketingActivityConsultResponseModel) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *AlipayMarketingActivityConsultResponseModel) SetUserId(v string) {
	o.UserId = &v
}

func (o AlipayMarketingActivityConsultResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingActivityConsultResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConsultResultInfoList) {
		toSerialize["consult_result_info_list"] = o.ConsultResultInfoList
	}
	if !IsNil(o.OpenId) {
		toSerialize["open_id"] = o.OpenId
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableAlipayMarketingActivityConsultResponseModel struct {
	value *AlipayMarketingActivityConsultResponseModel
	isSet bool
}

func (v NullableAlipayMarketingActivityConsultResponseModel) Get() *AlipayMarketingActivityConsultResponseModel {
	return v.value
}

func (v *NullableAlipayMarketingActivityConsultResponseModel) Set(val *AlipayMarketingActivityConsultResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityConsultResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityConsultResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityConsultResponseModel(val *AlipayMarketingActivityConsultResponseModel) *NullableAlipayMarketingActivityConsultResponseModel {
	return &NullableAlipayMarketingActivityConsultResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityConsultResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityConsultResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


