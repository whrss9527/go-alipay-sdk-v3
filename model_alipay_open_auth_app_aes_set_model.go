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

// checks if the AlipayOpenAuthAppAesSetModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenAuthAppAesSetModel{}

// AlipayOpenAuthAppAesSetModel struct for AlipayOpenAuthAppAesSetModel
type AlipayOpenAuthAppAesSetModel struct {
	// 商家应用appId
	MerchantAppId *string `json:"merchant_app_id,omitempty"`
}

// NewAlipayOpenAuthAppAesSetModel instantiates a new AlipayOpenAuthAppAesSetModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenAuthAppAesSetModel() *AlipayOpenAuthAppAesSetModel {
	this := AlipayOpenAuthAppAesSetModel{}
	return &this
}

// NewAlipayOpenAuthAppAesSetModelWithDefaults instantiates a new AlipayOpenAuthAppAesSetModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenAuthAppAesSetModelWithDefaults() *AlipayOpenAuthAppAesSetModel {
	this := AlipayOpenAuthAppAesSetModel{}
	return &this
}

// GetMerchantAppId returns the MerchantAppId field value if set, zero value otherwise.
func (o *AlipayOpenAuthAppAesSetModel) GetMerchantAppId() string {
	if o == nil || IsNil(o.MerchantAppId) {
		var ret string
		return ret
	}
	return *o.MerchantAppId
}

// GetMerchantAppIdOk returns a tuple with the MerchantAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAuthAppAesSetModel) GetMerchantAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantAppId) {
		return nil, false
	}
	return o.MerchantAppId, true
}

// HasMerchantAppId returns a boolean if a field has been set.
func (o *AlipayOpenAuthAppAesSetModel) HasMerchantAppId() bool {
	if o != nil && !IsNil(o.MerchantAppId) {
		return true
	}

	return false
}

// SetMerchantAppId gets a reference to the given string and assigns it to the MerchantAppId field.
func (o *AlipayOpenAuthAppAesSetModel) SetMerchantAppId(v string) {
	o.MerchantAppId = &v
}

func (o AlipayOpenAuthAppAesSetModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenAuthAppAesSetModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantAppId) {
		toSerialize["merchant_app_id"] = o.MerchantAppId
	}
	return toSerialize, nil
}

type NullableAlipayOpenAuthAppAesSetModel struct {
	value *AlipayOpenAuthAppAesSetModel
	isSet bool
}

func (v NullableAlipayOpenAuthAppAesSetModel) Get() *AlipayOpenAuthAppAesSetModel {
	return v.value
}

func (v *NullableAlipayOpenAuthAppAesSetModel) Set(val *AlipayOpenAuthAppAesSetModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenAuthAppAesSetModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenAuthAppAesSetModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenAuthAppAesSetModel(val *AlipayOpenAuthAppAesSetModel) *NullableAlipayOpenAuthAppAesSetModel {
	return &NullableAlipayOpenAuthAppAesSetModel{value: val, isSet: true}
}

func (v NullableAlipayOpenAuthAppAesSetModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenAuthAppAesSetModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


