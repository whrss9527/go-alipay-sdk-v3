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

// checks if the AlipayOpenPublicAdvertBatchqueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicAdvertBatchqueryResponseModel{}

// AlipayOpenPublicAdvertBatchqueryResponseModel struct for AlipayOpenPublicAdvertBatchqueryResponseModel
type AlipayOpenPublicAdvertBatchqueryResponseModel struct {
	// 广告位列表 。目前只有一个广告位。
	AdvertList []Advert `json:"advert_list,omitempty"`
	// 广告组数量
	Count *int32 `json:"count,omitempty"`
}

// NewAlipayOpenPublicAdvertBatchqueryResponseModel instantiates a new AlipayOpenPublicAdvertBatchqueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicAdvertBatchqueryResponseModel() *AlipayOpenPublicAdvertBatchqueryResponseModel {
	this := AlipayOpenPublicAdvertBatchqueryResponseModel{}
	return &this
}

// NewAlipayOpenPublicAdvertBatchqueryResponseModelWithDefaults instantiates a new AlipayOpenPublicAdvertBatchqueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicAdvertBatchqueryResponseModelWithDefaults() *AlipayOpenPublicAdvertBatchqueryResponseModel {
	this := AlipayOpenPublicAdvertBatchqueryResponseModel{}
	return &this
}

// GetAdvertList returns the AdvertList field value if set, zero value otherwise.
func (o *AlipayOpenPublicAdvertBatchqueryResponseModel) GetAdvertList() []Advert {
	if o == nil || IsNil(o.AdvertList) {
		var ret []Advert
		return ret
	}
	return o.AdvertList
}

// GetAdvertListOk returns a tuple with the AdvertList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicAdvertBatchqueryResponseModel) GetAdvertListOk() ([]Advert, bool) {
	if o == nil || IsNil(o.AdvertList) {
		return nil, false
	}
	return o.AdvertList, true
}

// HasAdvertList returns a boolean if a field has been set.
func (o *AlipayOpenPublicAdvertBatchqueryResponseModel) HasAdvertList() bool {
	if o != nil && !IsNil(o.AdvertList) {
		return true
	}

	return false
}

// SetAdvertList gets a reference to the given []Advert and assigns it to the AdvertList field.
func (o *AlipayOpenPublicAdvertBatchqueryResponseModel) SetAdvertList(v []Advert) {
	o.AdvertList = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *AlipayOpenPublicAdvertBatchqueryResponseModel) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicAdvertBatchqueryResponseModel) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *AlipayOpenPublicAdvertBatchqueryResponseModel) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *AlipayOpenPublicAdvertBatchqueryResponseModel) SetCount(v int32) {
	o.Count = &v
}

func (o AlipayOpenPublicAdvertBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicAdvertBatchqueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdvertList) {
		toSerialize["advert_list"] = o.AdvertList
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicAdvertBatchqueryResponseModel struct {
	value *AlipayOpenPublicAdvertBatchqueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenPublicAdvertBatchqueryResponseModel) Get() *AlipayOpenPublicAdvertBatchqueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenPublicAdvertBatchqueryResponseModel) Set(val *AlipayOpenPublicAdvertBatchqueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicAdvertBatchqueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicAdvertBatchqueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicAdvertBatchqueryResponseModel(val *AlipayOpenPublicAdvertBatchqueryResponseModel) *NullableAlipayOpenPublicAdvertBatchqueryResponseModel {
	return &NullableAlipayOpenPublicAdvertBatchqueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicAdvertBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicAdvertBatchqueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


