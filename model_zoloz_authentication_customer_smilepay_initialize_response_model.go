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

// checks if the ZolozAuthenticationCustomerSmilepayInitializeResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZolozAuthenticationCustomerSmilepayInitializeResponseModel{}

// ZolozAuthenticationCustomerSmilepayInitializeResponseModel struct for ZolozAuthenticationCustomerSmilepayInitializeResponseModel
type ZolozAuthenticationCustomerSmilepayInitializeResponseModel struct {
	// 返回值
	Result *string `json:"result,omitempty"`
}

// NewZolozAuthenticationCustomerSmilepayInitializeResponseModel instantiates a new ZolozAuthenticationCustomerSmilepayInitializeResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZolozAuthenticationCustomerSmilepayInitializeResponseModel() *ZolozAuthenticationCustomerSmilepayInitializeResponseModel {
	this := ZolozAuthenticationCustomerSmilepayInitializeResponseModel{}
	return &this
}

// NewZolozAuthenticationCustomerSmilepayInitializeResponseModelWithDefaults instantiates a new ZolozAuthenticationCustomerSmilepayInitializeResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZolozAuthenticationCustomerSmilepayInitializeResponseModelWithDefaults() *ZolozAuthenticationCustomerSmilepayInitializeResponseModel {
	this := ZolozAuthenticationCustomerSmilepayInitializeResponseModel{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ZolozAuthenticationCustomerSmilepayInitializeResponseModel) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZolozAuthenticationCustomerSmilepayInitializeResponseModel) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ZolozAuthenticationCustomerSmilepayInitializeResponseModel) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *ZolozAuthenticationCustomerSmilepayInitializeResponseModel) SetResult(v string) {
	o.Result = &v
}

func (o ZolozAuthenticationCustomerSmilepayInitializeResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZolozAuthenticationCustomerSmilepayInitializeResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableZolozAuthenticationCustomerSmilepayInitializeResponseModel struct {
	value *ZolozAuthenticationCustomerSmilepayInitializeResponseModel
	isSet bool
}

func (v NullableZolozAuthenticationCustomerSmilepayInitializeResponseModel) Get() *ZolozAuthenticationCustomerSmilepayInitializeResponseModel {
	return v.value
}

func (v *NullableZolozAuthenticationCustomerSmilepayInitializeResponseModel) Set(val *ZolozAuthenticationCustomerSmilepayInitializeResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableZolozAuthenticationCustomerSmilepayInitializeResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableZolozAuthenticationCustomerSmilepayInitializeResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZolozAuthenticationCustomerSmilepayInitializeResponseModel(val *ZolozAuthenticationCustomerSmilepayInitializeResponseModel) *NullableZolozAuthenticationCustomerSmilepayInitializeResponseModel {
	return &NullableZolozAuthenticationCustomerSmilepayInitializeResponseModel{value: val, isSet: true}
}

func (v NullableZolozAuthenticationCustomerSmilepayInitializeResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZolozAuthenticationCustomerSmilepayInitializeResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


