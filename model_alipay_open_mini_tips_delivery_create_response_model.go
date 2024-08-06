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

// checks if the AlipayOpenMiniTipsDeliveryCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniTipsDeliveryCreateResponseModel{}

// AlipayOpenMiniTipsDeliveryCreateResponseModel struct for AlipayOpenMiniTipsDeliveryCreateResponseModel
type AlipayOpenMiniTipsDeliveryCreateResponseModel struct {
	// 收藏引导活动ID
	DeliveryId *string `json:"delivery_id,omitempty"`
}

// NewAlipayOpenMiniTipsDeliveryCreateResponseModel instantiates a new AlipayOpenMiniTipsDeliveryCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniTipsDeliveryCreateResponseModel() *AlipayOpenMiniTipsDeliveryCreateResponseModel {
	this := AlipayOpenMiniTipsDeliveryCreateResponseModel{}
	return &this
}

// NewAlipayOpenMiniTipsDeliveryCreateResponseModelWithDefaults instantiates a new AlipayOpenMiniTipsDeliveryCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniTipsDeliveryCreateResponseModelWithDefaults() *AlipayOpenMiniTipsDeliveryCreateResponseModel {
	this := AlipayOpenMiniTipsDeliveryCreateResponseModel{}
	return &this
}

// GetDeliveryId returns the DeliveryId field value if set, zero value otherwise.
func (o *AlipayOpenMiniTipsDeliveryCreateResponseModel) GetDeliveryId() string {
	if o == nil || IsNil(o.DeliveryId) {
		var ret string
		return ret
	}
	return *o.DeliveryId
}

// GetDeliveryIdOk returns a tuple with the DeliveryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniTipsDeliveryCreateResponseModel) GetDeliveryIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeliveryId) {
		return nil, false
	}
	return o.DeliveryId, true
}

// HasDeliveryId returns a boolean if a field has been set.
func (o *AlipayOpenMiniTipsDeliveryCreateResponseModel) HasDeliveryId() bool {
	if o != nil && !IsNil(o.DeliveryId) {
		return true
	}

	return false
}

// SetDeliveryId gets a reference to the given string and assigns it to the DeliveryId field.
func (o *AlipayOpenMiniTipsDeliveryCreateResponseModel) SetDeliveryId(v string) {
	o.DeliveryId = &v
}

func (o AlipayOpenMiniTipsDeliveryCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniTipsDeliveryCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeliveryId) {
		toSerialize["delivery_id"] = o.DeliveryId
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniTipsDeliveryCreateResponseModel struct {
	value *AlipayOpenMiniTipsDeliveryCreateResponseModel
	isSet bool
}

func (v NullableAlipayOpenMiniTipsDeliveryCreateResponseModel) Get() *AlipayOpenMiniTipsDeliveryCreateResponseModel {
	return v.value
}

func (v *NullableAlipayOpenMiniTipsDeliveryCreateResponseModel) Set(val *AlipayOpenMiniTipsDeliveryCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniTipsDeliveryCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniTipsDeliveryCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniTipsDeliveryCreateResponseModel(val *AlipayOpenMiniTipsDeliveryCreateResponseModel) *NullableAlipayOpenMiniTipsDeliveryCreateResponseModel {
	return &NullableAlipayOpenMiniTipsDeliveryCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniTipsDeliveryCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniTipsDeliveryCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


