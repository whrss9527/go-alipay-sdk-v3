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

// checks if the AlipayOpenAppApiFieldQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenAppApiFieldQueryResponseModel{}

// AlipayOpenAppApiFieldQueryResponseModel struct for AlipayOpenAppApiFieldQueryResponseModel
type AlipayOpenAppApiFieldQueryResponseModel struct {
	AuthFieldResponse *AuthFieldResponse `json:"auth_field_response,omitempty"`
}

// NewAlipayOpenAppApiFieldQueryResponseModel instantiates a new AlipayOpenAppApiFieldQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenAppApiFieldQueryResponseModel() *AlipayOpenAppApiFieldQueryResponseModel {
	this := AlipayOpenAppApiFieldQueryResponseModel{}
	return &this
}

// NewAlipayOpenAppApiFieldQueryResponseModelWithDefaults instantiates a new AlipayOpenAppApiFieldQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenAppApiFieldQueryResponseModelWithDefaults() *AlipayOpenAppApiFieldQueryResponseModel {
	this := AlipayOpenAppApiFieldQueryResponseModel{}
	return &this
}

// GetAuthFieldResponse returns the AuthFieldResponse field value if set, zero value otherwise.
func (o *AlipayOpenAppApiFieldQueryResponseModel) GetAuthFieldResponse() AuthFieldResponse {
	if o == nil || IsNil(o.AuthFieldResponse) {
		var ret AuthFieldResponse
		return ret
	}
	return *o.AuthFieldResponse
}

// GetAuthFieldResponseOk returns a tuple with the AuthFieldResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAppApiFieldQueryResponseModel) GetAuthFieldResponseOk() (*AuthFieldResponse, bool) {
	if o == nil || IsNil(o.AuthFieldResponse) {
		return nil, false
	}
	return o.AuthFieldResponse, true
}

// HasAuthFieldResponse returns a boolean if a field has been set.
func (o *AlipayOpenAppApiFieldQueryResponseModel) HasAuthFieldResponse() bool {
	if o != nil && !IsNil(o.AuthFieldResponse) {
		return true
	}

	return false
}

// SetAuthFieldResponse gets a reference to the given AuthFieldResponse and assigns it to the AuthFieldResponse field.
func (o *AlipayOpenAppApiFieldQueryResponseModel) SetAuthFieldResponse(v AuthFieldResponse) {
	o.AuthFieldResponse = &v
}

func (o AlipayOpenAppApiFieldQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenAppApiFieldQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthFieldResponse) {
		toSerialize["auth_field_response"] = o.AuthFieldResponse
	}
	return toSerialize, nil
}

type NullableAlipayOpenAppApiFieldQueryResponseModel struct {
	value *AlipayOpenAppApiFieldQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenAppApiFieldQueryResponseModel) Get() *AlipayOpenAppApiFieldQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenAppApiFieldQueryResponseModel) Set(val *AlipayOpenAppApiFieldQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenAppApiFieldQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenAppApiFieldQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenAppApiFieldQueryResponseModel(val *AlipayOpenAppApiFieldQueryResponseModel) *NullableAlipayOpenAppApiFieldQueryResponseModel {
	return &NullableAlipayOpenAppApiFieldQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenAppApiFieldQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenAppApiFieldQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


