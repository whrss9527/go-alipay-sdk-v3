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

// checks if the AlipayOpenPublicLifeAboardApplyResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicLifeAboardApplyResponseModel{}

// AlipayOpenPublicLifeAboardApplyResponseModel struct for AlipayOpenPublicLifeAboardApplyResponseModel
type AlipayOpenPublicLifeAboardApplyResponseModel struct {
	// 上架成功后返回的提示
	Result *string `json:"result,omitempty"`
}

// NewAlipayOpenPublicLifeAboardApplyResponseModel instantiates a new AlipayOpenPublicLifeAboardApplyResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicLifeAboardApplyResponseModel() *AlipayOpenPublicLifeAboardApplyResponseModel {
	this := AlipayOpenPublicLifeAboardApplyResponseModel{}
	return &this
}

// NewAlipayOpenPublicLifeAboardApplyResponseModelWithDefaults instantiates a new AlipayOpenPublicLifeAboardApplyResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicLifeAboardApplyResponseModelWithDefaults() *AlipayOpenPublicLifeAboardApplyResponseModel {
	this := AlipayOpenPublicLifeAboardApplyResponseModel{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *AlipayOpenPublicLifeAboardApplyResponseModel) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicLifeAboardApplyResponseModel) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *AlipayOpenPublicLifeAboardApplyResponseModel) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *AlipayOpenPublicLifeAboardApplyResponseModel) SetResult(v string) {
	o.Result = &v
}

func (o AlipayOpenPublicLifeAboardApplyResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicLifeAboardApplyResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicLifeAboardApplyResponseModel struct {
	value *AlipayOpenPublicLifeAboardApplyResponseModel
	isSet bool
}

func (v NullableAlipayOpenPublicLifeAboardApplyResponseModel) Get() *AlipayOpenPublicLifeAboardApplyResponseModel {
	return v.value
}

func (v *NullableAlipayOpenPublicLifeAboardApplyResponseModel) Set(val *AlipayOpenPublicLifeAboardApplyResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicLifeAboardApplyResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicLifeAboardApplyResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicLifeAboardApplyResponseModel(val *AlipayOpenPublicLifeAboardApplyResponseModel) *NullableAlipayOpenPublicLifeAboardApplyResponseModel {
	return &NullableAlipayOpenPublicLifeAboardApplyResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicLifeAboardApplyResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicLifeAboardApplyResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


