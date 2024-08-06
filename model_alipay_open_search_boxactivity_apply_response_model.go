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

// checks if the AlipayOpenSearchBoxactivityApplyResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenSearchBoxactivityApplyResponseModel{}

// AlipayOpenSearchBoxactivityApplyResponseModel struct for AlipayOpenSearchBoxactivityApplyResponseModel
type AlipayOpenSearchBoxactivityApplyResponseModel struct {
	// 搜索直达活动ID
	BoxActivityId *string `json:"box_activity_id,omitempty"`
}

// NewAlipayOpenSearchBoxactivityApplyResponseModel instantiates a new AlipayOpenSearchBoxactivityApplyResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenSearchBoxactivityApplyResponseModel() *AlipayOpenSearchBoxactivityApplyResponseModel {
	this := AlipayOpenSearchBoxactivityApplyResponseModel{}
	return &this
}

// NewAlipayOpenSearchBoxactivityApplyResponseModelWithDefaults instantiates a new AlipayOpenSearchBoxactivityApplyResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenSearchBoxactivityApplyResponseModelWithDefaults() *AlipayOpenSearchBoxactivityApplyResponseModel {
	this := AlipayOpenSearchBoxactivityApplyResponseModel{}
	return &this
}

// GetBoxActivityId returns the BoxActivityId field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxactivityApplyResponseModel) GetBoxActivityId() string {
	if o == nil || IsNil(o.BoxActivityId) {
		var ret string
		return ret
	}
	return *o.BoxActivityId
}

// GetBoxActivityIdOk returns a tuple with the BoxActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxactivityApplyResponseModel) GetBoxActivityIdOk() (*string, bool) {
	if o == nil || IsNil(o.BoxActivityId) {
		return nil, false
	}
	return o.BoxActivityId, true
}

// HasBoxActivityId returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxactivityApplyResponseModel) HasBoxActivityId() bool {
	if o != nil && !IsNil(o.BoxActivityId) {
		return true
	}

	return false
}

// SetBoxActivityId gets a reference to the given string and assigns it to the BoxActivityId field.
func (o *AlipayOpenSearchBoxactivityApplyResponseModel) SetBoxActivityId(v string) {
	o.BoxActivityId = &v
}

func (o AlipayOpenSearchBoxactivityApplyResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenSearchBoxactivityApplyResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BoxActivityId) {
		toSerialize["box_activity_id"] = o.BoxActivityId
	}
	return toSerialize, nil
}

type NullableAlipayOpenSearchBoxactivityApplyResponseModel struct {
	value *AlipayOpenSearchBoxactivityApplyResponseModel
	isSet bool
}

func (v NullableAlipayOpenSearchBoxactivityApplyResponseModel) Get() *AlipayOpenSearchBoxactivityApplyResponseModel {
	return v.value
}

func (v *NullableAlipayOpenSearchBoxactivityApplyResponseModel) Set(val *AlipayOpenSearchBoxactivityApplyResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSearchBoxactivityApplyResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSearchBoxactivityApplyResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSearchBoxactivityApplyResponseModel(val *AlipayOpenSearchBoxactivityApplyResponseModel) *NullableAlipayOpenSearchBoxactivityApplyResponseModel {
	return &NullableAlipayOpenSearchBoxactivityApplyResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenSearchBoxactivityApplyResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSearchBoxactivityApplyResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


