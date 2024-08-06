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

// checks if the AlipayMarketingCardConsumeSyncResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingCardConsumeSyncResponseModel{}

// AlipayMarketingCardConsumeSyncResponseModel struct for AlipayMarketingCardConsumeSyncResponseModel
type AlipayMarketingCardConsumeSyncResponseModel struct {
	// 外部卡号
	ExternalCardNo *string `json:"external_card_no,omitempty"`
}

// NewAlipayMarketingCardConsumeSyncResponseModel instantiates a new AlipayMarketingCardConsumeSyncResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingCardConsumeSyncResponseModel() *AlipayMarketingCardConsumeSyncResponseModel {
	this := AlipayMarketingCardConsumeSyncResponseModel{}
	return &this
}

// NewAlipayMarketingCardConsumeSyncResponseModelWithDefaults instantiates a new AlipayMarketingCardConsumeSyncResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingCardConsumeSyncResponseModelWithDefaults() *AlipayMarketingCardConsumeSyncResponseModel {
	this := AlipayMarketingCardConsumeSyncResponseModel{}
	return &this
}

// GetExternalCardNo returns the ExternalCardNo field value if set, zero value otherwise.
func (o *AlipayMarketingCardConsumeSyncResponseModel) GetExternalCardNo() string {
	if o == nil || IsNil(o.ExternalCardNo) {
		var ret string
		return ret
	}
	return *o.ExternalCardNo
}

// GetExternalCardNoOk returns a tuple with the ExternalCardNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardConsumeSyncResponseModel) GetExternalCardNoOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalCardNo) {
		return nil, false
	}
	return o.ExternalCardNo, true
}

// HasExternalCardNo returns a boolean if a field has been set.
func (o *AlipayMarketingCardConsumeSyncResponseModel) HasExternalCardNo() bool {
	if o != nil && !IsNil(o.ExternalCardNo) {
		return true
	}

	return false
}

// SetExternalCardNo gets a reference to the given string and assigns it to the ExternalCardNo field.
func (o *AlipayMarketingCardConsumeSyncResponseModel) SetExternalCardNo(v string) {
	o.ExternalCardNo = &v
}

func (o AlipayMarketingCardConsumeSyncResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingCardConsumeSyncResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternalCardNo) {
		toSerialize["external_card_no"] = o.ExternalCardNo
	}
	return toSerialize, nil
}

type NullableAlipayMarketingCardConsumeSyncResponseModel struct {
	value *AlipayMarketingCardConsumeSyncResponseModel
	isSet bool
}

func (v NullableAlipayMarketingCardConsumeSyncResponseModel) Get() *AlipayMarketingCardConsumeSyncResponseModel {
	return v.value
}

func (v *NullableAlipayMarketingCardConsumeSyncResponseModel) Set(val *AlipayMarketingCardConsumeSyncResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingCardConsumeSyncResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingCardConsumeSyncResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingCardConsumeSyncResponseModel(val *AlipayMarketingCardConsumeSyncResponseModel) *NullableAlipayMarketingCardConsumeSyncResponseModel {
	return &NullableAlipayMarketingCardConsumeSyncResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingCardConsumeSyncResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingCardConsumeSyncResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


