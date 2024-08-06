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

// checks if the AlipayIserviceCcmSwTreecategoryCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayIserviceCcmSwTreecategoryCreateResponseModel{}

// AlipayIserviceCcmSwTreecategoryCreateResponseModel struct for AlipayIserviceCcmSwTreecategoryCreateResponseModel
type AlipayIserviceCcmSwTreecategoryCreateResponseModel struct {
	// 节点ID
	Id *int32 `json:"id,omitempty"`
}

// NewAlipayIserviceCcmSwTreecategoryCreateResponseModel instantiates a new AlipayIserviceCcmSwTreecategoryCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayIserviceCcmSwTreecategoryCreateResponseModel() *AlipayIserviceCcmSwTreecategoryCreateResponseModel {
	this := AlipayIserviceCcmSwTreecategoryCreateResponseModel{}
	return &this
}

// NewAlipayIserviceCcmSwTreecategoryCreateResponseModelWithDefaults instantiates a new AlipayIserviceCcmSwTreecategoryCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayIserviceCcmSwTreecategoryCreateResponseModelWithDefaults() *AlipayIserviceCcmSwTreecategoryCreateResponseModel {
	this := AlipayIserviceCcmSwTreecategoryCreateResponseModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwTreecategoryCreateResponseModel) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwTreecategoryCreateResponseModel) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwTreecategoryCreateResponseModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AlipayIserviceCcmSwTreecategoryCreateResponseModel) SetId(v int32) {
	o.Id = &v
}

func (o AlipayIserviceCcmSwTreecategoryCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayIserviceCcmSwTreecategoryCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableAlipayIserviceCcmSwTreecategoryCreateResponseModel struct {
	value *AlipayIserviceCcmSwTreecategoryCreateResponseModel
	isSet bool
}

func (v NullableAlipayIserviceCcmSwTreecategoryCreateResponseModel) Get() *AlipayIserviceCcmSwTreecategoryCreateResponseModel {
	return v.value
}

func (v *NullableAlipayIserviceCcmSwTreecategoryCreateResponseModel) Set(val *AlipayIserviceCcmSwTreecategoryCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayIserviceCcmSwTreecategoryCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayIserviceCcmSwTreecategoryCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayIserviceCcmSwTreecategoryCreateResponseModel(val *AlipayIserviceCcmSwTreecategoryCreateResponseModel) *NullableAlipayIserviceCcmSwTreecategoryCreateResponseModel {
	return &NullableAlipayIserviceCcmSwTreecategoryCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayIserviceCcmSwTreecategoryCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayIserviceCcmSwTreecategoryCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


