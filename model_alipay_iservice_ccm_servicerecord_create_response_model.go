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

// checks if the AlipayIserviceCcmServicerecordCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayIserviceCcmServicerecordCreateResponseModel{}

// AlipayIserviceCcmServicerecordCreateResponseModel struct for AlipayIserviceCcmServicerecordCreateResponseModel
type AlipayIserviceCcmServicerecordCreateResponseModel struct {
	// 服务记录ID
	ServiceRecordId *string `json:"service_record_id,omitempty"`
}

// NewAlipayIserviceCcmServicerecordCreateResponseModel instantiates a new AlipayIserviceCcmServicerecordCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayIserviceCcmServicerecordCreateResponseModel() *AlipayIserviceCcmServicerecordCreateResponseModel {
	this := AlipayIserviceCcmServicerecordCreateResponseModel{}
	return &this
}

// NewAlipayIserviceCcmServicerecordCreateResponseModelWithDefaults instantiates a new AlipayIserviceCcmServicerecordCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayIserviceCcmServicerecordCreateResponseModelWithDefaults() *AlipayIserviceCcmServicerecordCreateResponseModel {
	this := AlipayIserviceCcmServicerecordCreateResponseModel{}
	return &this
}

// GetServiceRecordId returns the ServiceRecordId field value if set, zero value otherwise.
func (o *AlipayIserviceCcmServicerecordCreateResponseModel) GetServiceRecordId() string {
	if o == nil || IsNil(o.ServiceRecordId) {
		var ret string
		return ret
	}
	return *o.ServiceRecordId
}

// GetServiceRecordIdOk returns a tuple with the ServiceRecordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmServicerecordCreateResponseModel) GetServiceRecordIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceRecordId) {
		return nil, false
	}
	return o.ServiceRecordId, true
}

// HasServiceRecordId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmServicerecordCreateResponseModel) HasServiceRecordId() bool {
	if o != nil && !IsNil(o.ServiceRecordId) {
		return true
	}

	return false
}

// SetServiceRecordId gets a reference to the given string and assigns it to the ServiceRecordId field.
func (o *AlipayIserviceCcmServicerecordCreateResponseModel) SetServiceRecordId(v string) {
	o.ServiceRecordId = &v
}

func (o AlipayIserviceCcmServicerecordCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayIserviceCcmServicerecordCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServiceRecordId) {
		toSerialize["service_record_id"] = o.ServiceRecordId
	}
	return toSerialize, nil
}

type NullableAlipayIserviceCcmServicerecordCreateResponseModel struct {
	value *AlipayIserviceCcmServicerecordCreateResponseModel
	isSet bool
}

func (v NullableAlipayIserviceCcmServicerecordCreateResponseModel) Get() *AlipayIserviceCcmServicerecordCreateResponseModel {
	return v.value
}

func (v *NullableAlipayIserviceCcmServicerecordCreateResponseModel) Set(val *AlipayIserviceCcmServicerecordCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayIserviceCcmServicerecordCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayIserviceCcmServicerecordCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayIserviceCcmServicerecordCreateResponseModel(val *AlipayIserviceCcmServicerecordCreateResponseModel) *NullableAlipayIserviceCcmServicerecordCreateResponseModel {
	return &NullableAlipayIserviceCcmServicerecordCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayIserviceCcmServicerecordCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayIserviceCcmServicerecordCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
