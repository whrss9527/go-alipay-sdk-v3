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

// checks if the AlipayMarketingActivityVoucherCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingActivityVoucherCreateResponseModel{}

// AlipayMarketingActivityVoucherCreateResponseModel struct for AlipayMarketingActivityVoucherCreateResponseModel
type AlipayMarketingActivityVoucherCreateResponseModel struct {
	// 活动id
	ActivityId *string `json:"activity_id,omitempty"`
}

// NewAlipayMarketingActivityVoucherCreateResponseModel instantiates a new AlipayMarketingActivityVoucherCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingActivityVoucherCreateResponseModel() *AlipayMarketingActivityVoucherCreateResponseModel {
	this := AlipayMarketingActivityVoucherCreateResponseModel{}
	return &this
}

// NewAlipayMarketingActivityVoucherCreateResponseModelWithDefaults instantiates a new AlipayMarketingActivityVoucherCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingActivityVoucherCreateResponseModelWithDefaults() *AlipayMarketingActivityVoucherCreateResponseModel {
	this := AlipayMarketingActivityVoucherCreateResponseModel{}
	return &this
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherCreateResponseModel) GetActivityId() string {
	if o == nil || IsNil(o.ActivityId) {
		var ret string
		return ret
	}
	return *o.ActivityId
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherCreateResponseModel) GetActivityIdOk() (*string, bool) {
	if o == nil || IsNil(o.ActivityId) {
		return nil, false
	}
	return o.ActivityId, true
}

// HasActivityId returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherCreateResponseModel) HasActivityId() bool {
	if o != nil && !IsNil(o.ActivityId) {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given string and assigns it to the ActivityId field.
func (o *AlipayMarketingActivityVoucherCreateResponseModel) SetActivityId(v string) {
	o.ActivityId = &v
}

func (o AlipayMarketingActivityVoucherCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingActivityVoucherCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivityId) {
		toSerialize["activity_id"] = o.ActivityId
	}
	return toSerialize, nil
}

type NullableAlipayMarketingActivityVoucherCreateResponseModel struct {
	value *AlipayMarketingActivityVoucherCreateResponseModel
	isSet bool
}

func (v NullableAlipayMarketingActivityVoucherCreateResponseModel) Get() *AlipayMarketingActivityVoucherCreateResponseModel {
	return v.value
}

func (v *NullableAlipayMarketingActivityVoucherCreateResponseModel) Set(val *AlipayMarketingActivityVoucherCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityVoucherCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityVoucherCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityVoucherCreateResponseModel(val *AlipayMarketingActivityVoucherCreateResponseModel) *NullableAlipayMarketingActivityVoucherCreateResponseModel {
	return &NullableAlipayMarketingActivityVoucherCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityVoucherCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityVoucherCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


