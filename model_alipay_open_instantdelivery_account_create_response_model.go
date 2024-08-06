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

// checks if the AlipayOpenInstantdeliveryAccountCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenInstantdeliveryAccountCreateResponseModel{}

// AlipayOpenInstantdeliveryAccountCreateResponseModel struct for AlipayOpenInstantdeliveryAccountCreateResponseModel
type AlipayOpenInstantdeliveryAccountCreateResponseModel struct {
	// 配送公司账户创建结果列表
	LogisticsAccountStatus []LogisticsAccountStatusDTO `json:"logistics_account_status,omitempty"`
}

// NewAlipayOpenInstantdeliveryAccountCreateResponseModel instantiates a new AlipayOpenInstantdeliveryAccountCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenInstantdeliveryAccountCreateResponseModel() *AlipayOpenInstantdeliveryAccountCreateResponseModel {
	this := AlipayOpenInstantdeliveryAccountCreateResponseModel{}
	return &this
}

// NewAlipayOpenInstantdeliveryAccountCreateResponseModelWithDefaults instantiates a new AlipayOpenInstantdeliveryAccountCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenInstantdeliveryAccountCreateResponseModelWithDefaults() *AlipayOpenInstantdeliveryAccountCreateResponseModel {
	this := AlipayOpenInstantdeliveryAccountCreateResponseModel{}
	return &this
}

// GetLogisticsAccountStatus returns the LogisticsAccountStatus field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryAccountCreateResponseModel) GetLogisticsAccountStatus() []LogisticsAccountStatusDTO {
	if o == nil || IsNil(o.LogisticsAccountStatus) {
		var ret []LogisticsAccountStatusDTO
		return ret
	}
	return o.LogisticsAccountStatus
}

// GetLogisticsAccountStatusOk returns a tuple with the LogisticsAccountStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryAccountCreateResponseModel) GetLogisticsAccountStatusOk() ([]LogisticsAccountStatusDTO, bool) {
	if o == nil || IsNil(o.LogisticsAccountStatus) {
		return nil, false
	}
	return o.LogisticsAccountStatus, true
}

// HasLogisticsAccountStatus returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryAccountCreateResponseModel) HasLogisticsAccountStatus() bool {
	if o != nil && !IsNil(o.LogisticsAccountStatus) {
		return true
	}

	return false
}

// SetLogisticsAccountStatus gets a reference to the given []LogisticsAccountStatusDTO and assigns it to the LogisticsAccountStatus field.
func (o *AlipayOpenInstantdeliveryAccountCreateResponseModel) SetLogisticsAccountStatus(v []LogisticsAccountStatusDTO) {
	o.LogisticsAccountStatus = v
}

func (o AlipayOpenInstantdeliveryAccountCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenInstantdeliveryAccountCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LogisticsAccountStatus) {
		toSerialize["logistics_account_status"] = o.LogisticsAccountStatus
	}
	return toSerialize, nil
}

type NullableAlipayOpenInstantdeliveryAccountCreateResponseModel struct {
	value *AlipayOpenInstantdeliveryAccountCreateResponseModel
	isSet bool
}

func (v NullableAlipayOpenInstantdeliveryAccountCreateResponseModel) Get() *AlipayOpenInstantdeliveryAccountCreateResponseModel {
	return v.value
}

func (v *NullableAlipayOpenInstantdeliveryAccountCreateResponseModel) Set(val *AlipayOpenInstantdeliveryAccountCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenInstantdeliveryAccountCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenInstantdeliveryAccountCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenInstantdeliveryAccountCreateResponseModel(val *AlipayOpenInstantdeliveryAccountCreateResponseModel) *NullableAlipayOpenInstantdeliveryAccountCreateResponseModel {
	return &NullableAlipayOpenInstantdeliveryAccountCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenInstantdeliveryAccountCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenInstantdeliveryAccountCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


