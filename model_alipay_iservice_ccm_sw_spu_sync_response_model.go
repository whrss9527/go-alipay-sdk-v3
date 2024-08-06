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

// checks if the AlipayIserviceCcmSwSpuSyncResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayIserviceCcmSwSpuSyncResponseModel{}

// AlipayIserviceCcmSwSpuSyncResponseModel struct for AlipayIserviceCcmSwSpuSyncResponseModel
type AlipayIserviceCcmSwSpuSyncResponseModel struct {
	// 商品id
	SpuId *string `json:"spu_id,omitempty"`
}

// NewAlipayIserviceCcmSwSpuSyncResponseModel instantiates a new AlipayIserviceCcmSwSpuSyncResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayIserviceCcmSwSpuSyncResponseModel() *AlipayIserviceCcmSwSpuSyncResponseModel {
	this := AlipayIserviceCcmSwSpuSyncResponseModel{}
	return &this
}

// NewAlipayIserviceCcmSwSpuSyncResponseModelWithDefaults instantiates a new AlipayIserviceCcmSwSpuSyncResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayIserviceCcmSwSpuSyncResponseModelWithDefaults() *AlipayIserviceCcmSwSpuSyncResponseModel {
	this := AlipayIserviceCcmSwSpuSyncResponseModel{}
	return &this
}

// GetSpuId returns the SpuId field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwSpuSyncResponseModel) GetSpuId() string {
	if o == nil || IsNil(o.SpuId) {
		var ret string
		return ret
	}
	return *o.SpuId
}

// GetSpuIdOk returns a tuple with the SpuId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwSpuSyncResponseModel) GetSpuIdOk() (*string, bool) {
	if o == nil || IsNil(o.SpuId) {
		return nil, false
	}
	return o.SpuId, true
}

// HasSpuId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwSpuSyncResponseModel) HasSpuId() bool {
	if o != nil && !IsNil(o.SpuId) {
		return true
	}

	return false
}

// SetSpuId gets a reference to the given string and assigns it to the SpuId field.
func (o *AlipayIserviceCcmSwSpuSyncResponseModel) SetSpuId(v string) {
	o.SpuId = &v
}

func (o AlipayIserviceCcmSwSpuSyncResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayIserviceCcmSwSpuSyncResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SpuId) {
		toSerialize["spu_id"] = o.SpuId
	}
	return toSerialize, nil
}

type NullableAlipayIserviceCcmSwSpuSyncResponseModel struct {
	value *AlipayIserviceCcmSwSpuSyncResponseModel
	isSet bool
}

func (v NullableAlipayIserviceCcmSwSpuSyncResponseModel) Get() *AlipayIserviceCcmSwSpuSyncResponseModel {
	return v.value
}

func (v *NullableAlipayIserviceCcmSwSpuSyncResponseModel) Set(val *AlipayIserviceCcmSwSpuSyncResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayIserviceCcmSwSpuSyncResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayIserviceCcmSwSpuSyncResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayIserviceCcmSwSpuSyncResponseModel(val *AlipayIserviceCcmSwSpuSyncResponseModel) *NullableAlipayIserviceCcmSwSpuSyncResponseModel {
	return &NullableAlipayIserviceCcmSwSpuSyncResponseModel{value: val, isSet: true}
}

func (v NullableAlipayIserviceCcmSwSpuSyncResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayIserviceCcmSwSpuSyncResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


