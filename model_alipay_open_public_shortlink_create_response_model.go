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

// checks if the AlipayOpenPublicShortlinkCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicShortlinkCreateResponseModel{}

// AlipayOpenPublicShortlinkCreateResponseModel struct for AlipayOpenPublicShortlinkCreateResponseModel
type AlipayOpenPublicShortlinkCreateResponseModel struct {
	// 生成的带参推广短链接
	Shortlink *string `json:"shortlink,omitempty"`
}

// NewAlipayOpenPublicShortlinkCreateResponseModel instantiates a new AlipayOpenPublicShortlinkCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicShortlinkCreateResponseModel() *AlipayOpenPublicShortlinkCreateResponseModel {
	this := AlipayOpenPublicShortlinkCreateResponseModel{}
	return &this
}

// NewAlipayOpenPublicShortlinkCreateResponseModelWithDefaults instantiates a new AlipayOpenPublicShortlinkCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicShortlinkCreateResponseModelWithDefaults() *AlipayOpenPublicShortlinkCreateResponseModel {
	this := AlipayOpenPublicShortlinkCreateResponseModel{}
	return &this
}

// GetShortlink returns the Shortlink field value if set, zero value otherwise.
func (o *AlipayOpenPublicShortlinkCreateResponseModel) GetShortlink() string {
	if o == nil || IsNil(o.Shortlink) {
		var ret string
		return ret
	}
	return *o.Shortlink
}

// GetShortlinkOk returns a tuple with the Shortlink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicShortlinkCreateResponseModel) GetShortlinkOk() (*string, bool) {
	if o == nil || IsNil(o.Shortlink) {
		return nil, false
	}
	return o.Shortlink, true
}

// HasShortlink returns a boolean if a field has been set.
func (o *AlipayOpenPublicShortlinkCreateResponseModel) HasShortlink() bool {
	if o != nil && !IsNil(o.Shortlink) {
		return true
	}

	return false
}

// SetShortlink gets a reference to the given string and assigns it to the Shortlink field.
func (o *AlipayOpenPublicShortlinkCreateResponseModel) SetShortlink(v string) {
	o.Shortlink = &v
}

func (o AlipayOpenPublicShortlinkCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicShortlinkCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Shortlink) {
		toSerialize["shortlink"] = o.Shortlink
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicShortlinkCreateResponseModel struct {
	value *AlipayOpenPublicShortlinkCreateResponseModel
	isSet bool
}

func (v NullableAlipayOpenPublicShortlinkCreateResponseModel) Get() *AlipayOpenPublicShortlinkCreateResponseModel {
	return v.value
}

func (v *NullableAlipayOpenPublicShortlinkCreateResponseModel) Set(val *AlipayOpenPublicShortlinkCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicShortlinkCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicShortlinkCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicShortlinkCreateResponseModel(val *AlipayOpenPublicShortlinkCreateResponseModel) *NullableAlipayOpenPublicShortlinkCreateResponseModel {
	return &NullableAlipayOpenPublicShortlinkCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicShortlinkCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicShortlinkCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


