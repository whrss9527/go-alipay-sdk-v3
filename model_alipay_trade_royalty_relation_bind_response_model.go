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

// checks if the AlipayTradeRoyaltyRelationBindResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayTradeRoyaltyRelationBindResponseModel{}

// AlipayTradeRoyaltyRelationBindResponseModel struct for AlipayTradeRoyaltyRelationBindResponseModel
type AlipayTradeRoyaltyRelationBindResponseModel struct {
	// SUCCESS：分账关系绑定成功； FAIL：分账关系绑定失败。
	ResultCode *string `json:"result_code,omitempty"`
}

// NewAlipayTradeRoyaltyRelationBindResponseModel instantiates a new AlipayTradeRoyaltyRelationBindResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayTradeRoyaltyRelationBindResponseModel() *AlipayTradeRoyaltyRelationBindResponseModel {
	this := AlipayTradeRoyaltyRelationBindResponseModel{}
	return &this
}

// NewAlipayTradeRoyaltyRelationBindResponseModelWithDefaults instantiates a new AlipayTradeRoyaltyRelationBindResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayTradeRoyaltyRelationBindResponseModelWithDefaults() *AlipayTradeRoyaltyRelationBindResponseModel {
	this := AlipayTradeRoyaltyRelationBindResponseModel{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *AlipayTradeRoyaltyRelationBindResponseModel) GetResultCode() string {
	if o == nil || IsNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeRoyaltyRelationBindResponseModel) GetResultCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *AlipayTradeRoyaltyRelationBindResponseModel) HasResultCode() bool {
	if o != nil && !IsNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *AlipayTradeRoyaltyRelationBindResponseModel) SetResultCode(v string) {
	o.ResultCode = &v
}

func (o AlipayTradeRoyaltyRelationBindResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayTradeRoyaltyRelationBindResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResultCode) {
		toSerialize["result_code"] = o.ResultCode
	}
	return toSerialize, nil
}

type NullableAlipayTradeRoyaltyRelationBindResponseModel struct {
	value *AlipayTradeRoyaltyRelationBindResponseModel
	isSet bool
}

func (v NullableAlipayTradeRoyaltyRelationBindResponseModel) Get() *AlipayTradeRoyaltyRelationBindResponseModel {
	return v.value
}

func (v *NullableAlipayTradeRoyaltyRelationBindResponseModel) Set(val *AlipayTradeRoyaltyRelationBindResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayTradeRoyaltyRelationBindResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayTradeRoyaltyRelationBindResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayTradeRoyaltyRelationBindResponseModel(val *AlipayTradeRoyaltyRelationBindResponseModel) *NullableAlipayTradeRoyaltyRelationBindResponseModel {
	return &NullableAlipayTradeRoyaltyRelationBindResponseModel{value: val, isSet: true}
}

func (v NullableAlipayTradeRoyaltyRelationBindResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayTradeRoyaltyRelationBindResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
