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

// checks if the AlipayEcoMycarParkingSpaceinfoSyncResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEcoMycarParkingSpaceinfoSyncResponseModel{}

// AlipayEcoMycarParkingSpaceinfoSyncResponseModel struct for AlipayEcoMycarParkingSpaceinfoSyncResponseModel
type AlipayEcoMycarParkingSpaceinfoSyncResponseModel struct {
	// 同步结果：0 成功，1 失败
	SyncResult *string `json:"sync_result,omitempty"`
}

// NewAlipayEcoMycarParkingSpaceinfoSyncResponseModel instantiates a new AlipayEcoMycarParkingSpaceinfoSyncResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEcoMycarParkingSpaceinfoSyncResponseModel() *AlipayEcoMycarParkingSpaceinfoSyncResponseModel {
	this := AlipayEcoMycarParkingSpaceinfoSyncResponseModel{}
	return &this
}

// NewAlipayEcoMycarParkingSpaceinfoSyncResponseModelWithDefaults instantiates a new AlipayEcoMycarParkingSpaceinfoSyncResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEcoMycarParkingSpaceinfoSyncResponseModelWithDefaults() *AlipayEcoMycarParkingSpaceinfoSyncResponseModel {
	this := AlipayEcoMycarParkingSpaceinfoSyncResponseModel{}
	return &this
}

// GetSyncResult returns the SyncResult field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingSpaceinfoSyncResponseModel) GetSyncResult() string {
	if o == nil || IsNil(o.SyncResult) {
		var ret string
		return ret
	}
	return *o.SyncResult
}

// GetSyncResultOk returns a tuple with the SyncResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingSpaceinfoSyncResponseModel) GetSyncResultOk() (*string, bool) {
	if o == nil || IsNil(o.SyncResult) {
		return nil, false
	}
	return o.SyncResult, true
}

// HasSyncResult returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingSpaceinfoSyncResponseModel) HasSyncResult() bool {
	if o != nil && !IsNil(o.SyncResult) {
		return true
	}

	return false
}

// SetSyncResult gets a reference to the given string and assigns it to the SyncResult field.
func (o *AlipayEcoMycarParkingSpaceinfoSyncResponseModel) SetSyncResult(v string) {
	o.SyncResult = &v
}

func (o AlipayEcoMycarParkingSpaceinfoSyncResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEcoMycarParkingSpaceinfoSyncResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SyncResult) {
		toSerialize["sync_result"] = o.SyncResult
	}
	return toSerialize, nil
}

type NullableAlipayEcoMycarParkingSpaceinfoSyncResponseModel struct {
	value *AlipayEcoMycarParkingSpaceinfoSyncResponseModel
	isSet bool
}

func (v NullableAlipayEcoMycarParkingSpaceinfoSyncResponseModel) Get() *AlipayEcoMycarParkingSpaceinfoSyncResponseModel {
	return v.value
}

func (v *NullableAlipayEcoMycarParkingSpaceinfoSyncResponseModel) Set(val *AlipayEcoMycarParkingSpaceinfoSyncResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEcoMycarParkingSpaceinfoSyncResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEcoMycarParkingSpaceinfoSyncResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEcoMycarParkingSpaceinfoSyncResponseModel(val *AlipayEcoMycarParkingSpaceinfoSyncResponseModel) *NullableAlipayEcoMycarParkingSpaceinfoSyncResponseModel {
	return &NullableAlipayEcoMycarParkingSpaceinfoSyncResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEcoMycarParkingSpaceinfoSyncResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEcoMycarParkingSpaceinfoSyncResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


