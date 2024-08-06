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

// checks if the AlipayEcoMycarParkingParkinglotinfoCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEcoMycarParkingParkinglotinfoCreateResponseModel{}

// AlipayEcoMycarParkingParkinglotinfoCreateResponseModel struct for AlipayEcoMycarParkingParkinglotinfoCreateResponseModel
type AlipayEcoMycarParkingParkinglotinfoCreateResponseModel struct {
	// 支付宝返回停车场id。成功不为空，失败返回空
	ParkingId *string `json:"parking_id,omitempty"`
}

// NewAlipayEcoMycarParkingParkinglotinfoCreateResponseModel instantiates a new AlipayEcoMycarParkingParkinglotinfoCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEcoMycarParkingParkinglotinfoCreateResponseModel() *AlipayEcoMycarParkingParkinglotinfoCreateResponseModel {
	this := AlipayEcoMycarParkingParkinglotinfoCreateResponseModel{}
	return &this
}

// NewAlipayEcoMycarParkingParkinglotinfoCreateResponseModelWithDefaults instantiates a new AlipayEcoMycarParkingParkinglotinfoCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEcoMycarParkingParkinglotinfoCreateResponseModelWithDefaults() *AlipayEcoMycarParkingParkinglotinfoCreateResponseModel {
	this := AlipayEcoMycarParkingParkinglotinfoCreateResponseModel{}
	return &this
}

// GetParkingId returns the ParkingId field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingParkinglotinfoCreateResponseModel) GetParkingId() string {
	if o == nil || IsNil(o.ParkingId) {
		var ret string
		return ret
	}
	return *o.ParkingId
}

// GetParkingIdOk returns a tuple with the ParkingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoCreateResponseModel) GetParkingIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParkingId) {
		return nil, false
	}
	return o.ParkingId, true
}

// HasParkingId returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingParkinglotinfoCreateResponseModel) HasParkingId() bool {
	if o != nil && !IsNil(o.ParkingId) {
		return true
	}

	return false
}

// SetParkingId gets a reference to the given string and assigns it to the ParkingId field.
func (o *AlipayEcoMycarParkingParkinglotinfoCreateResponseModel) SetParkingId(v string) {
	o.ParkingId = &v
}

func (o AlipayEcoMycarParkingParkinglotinfoCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEcoMycarParkingParkinglotinfoCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ParkingId) {
		toSerialize["parking_id"] = o.ParkingId
	}
	return toSerialize, nil
}

type NullableAlipayEcoMycarParkingParkinglotinfoCreateResponseModel struct {
	value *AlipayEcoMycarParkingParkinglotinfoCreateResponseModel
	isSet bool
}

func (v NullableAlipayEcoMycarParkingParkinglotinfoCreateResponseModel) Get() *AlipayEcoMycarParkingParkinglotinfoCreateResponseModel {
	return v.value
}

func (v *NullableAlipayEcoMycarParkingParkinglotinfoCreateResponseModel) Set(val *AlipayEcoMycarParkingParkinglotinfoCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEcoMycarParkingParkinglotinfoCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEcoMycarParkingParkinglotinfoCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEcoMycarParkingParkinglotinfoCreateResponseModel(val *AlipayEcoMycarParkingParkinglotinfoCreateResponseModel) *NullableAlipayEcoMycarParkingParkinglotinfoCreateResponseModel {
	return &NullableAlipayEcoMycarParkingParkinglotinfoCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEcoMycarParkingParkinglotinfoCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEcoMycarParkingParkinglotinfoCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


