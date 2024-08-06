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

// checks if the AlipayUserAlipaypointSendResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayUserAlipaypointSendResponseModel{}

// AlipayUserAlipaypointSendResponseModel struct for AlipayUserAlipaypointSendResponseModel
type AlipayUserAlipaypointSendResponseModel struct {
	// 发放记录号。
	RecordId *string `json:"record_id,omitempty"`
	// 对账中心关联订单，一般场景无需关注
	TransactionId *string `json:"transaction_id,omitempty"`
}

// NewAlipayUserAlipaypointSendResponseModel instantiates a new AlipayUserAlipaypointSendResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayUserAlipaypointSendResponseModel() *AlipayUserAlipaypointSendResponseModel {
	this := AlipayUserAlipaypointSendResponseModel{}
	return &this
}

// NewAlipayUserAlipaypointSendResponseModelWithDefaults instantiates a new AlipayUserAlipaypointSendResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayUserAlipaypointSendResponseModelWithDefaults() *AlipayUserAlipaypointSendResponseModel {
	this := AlipayUserAlipaypointSendResponseModel{}
	return &this
}

// GetRecordId returns the RecordId field value if set, zero value otherwise.
func (o *AlipayUserAlipaypointSendResponseModel) GetRecordId() string {
	if o == nil || IsNil(o.RecordId) {
		var ret string
		return ret
	}
	return *o.RecordId
}

// GetRecordIdOk returns a tuple with the RecordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayUserAlipaypointSendResponseModel) GetRecordIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecordId) {
		return nil, false
	}
	return o.RecordId, true
}

// HasRecordId returns a boolean if a field has been set.
func (o *AlipayUserAlipaypointSendResponseModel) HasRecordId() bool {
	if o != nil && !IsNil(o.RecordId) {
		return true
	}

	return false
}

// SetRecordId gets a reference to the given string and assigns it to the RecordId field.
func (o *AlipayUserAlipaypointSendResponseModel) SetRecordId(v string) {
	o.RecordId = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *AlipayUserAlipaypointSendResponseModel) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayUserAlipaypointSendResponseModel) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *AlipayUserAlipaypointSendResponseModel) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *AlipayUserAlipaypointSendResponseModel) SetTransactionId(v string) {
	o.TransactionId = &v
}

func (o AlipayUserAlipaypointSendResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayUserAlipaypointSendResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RecordId) {
		toSerialize["record_id"] = o.RecordId
	}
	if !IsNil(o.TransactionId) {
		toSerialize["transaction_id"] = o.TransactionId
	}
	return toSerialize, nil
}

type NullableAlipayUserAlipaypointSendResponseModel struct {
	value *AlipayUserAlipaypointSendResponseModel
	isSet bool
}

func (v NullableAlipayUserAlipaypointSendResponseModel) Get() *AlipayUserAlipaypointSendResponseModel {
	return v.value
}

func (v *NullableAlipayUserAlipaypointSendResponseModel) Set(val *AlipayUserAlipaypointSendResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayUserAlipaypointSendResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayUserAlipaypointSendResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayUserAlipaypointSendResponseModel(val *AlipayUserAlipaypointSendResponseModel) *NullableAlipayUserAlipaypointSendResponseModel {
	return &NullableAlipayUserAlipaypointSendResponseModel{value: val, isSet: true}
}

func (v NullableAlipayUserAlipaypointSendResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayUserAlipaypointSendResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


