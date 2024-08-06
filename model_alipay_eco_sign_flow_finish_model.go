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

// checks if the AlipayEcoSignFlowFinishModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEcoSignFlowFinishModel{}

// AlipayEcoSignFlowFinishModel struct for AlipayEcoSignFlowFinishModel
type AlipayEcoSignFlowFinishModel struct {
	// 流程id
	FlowId *string `json:"flow_id,omitempty"`
}

// NewAlipayEcoSignFlowFinishModel instantiates a new AlipayEcoSignFlowFinishModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEcoSignFlowFinishModel() *AlipayEcoSignFlowFinishModel {
	this := AlipayEcoSignFlowFinishModel{}
	return &this
}

// NewAlipayEcoSignFlowFinishModelWithDefaults instantiates a new AlipayEcoSignFlowFinishModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEcoSignFlowFinishModelWithDefaults() *AlipayEcoSignFlowFinishModel {
	this := AlipayEcoSignFlowFinishModel{}
	return &this
}

// GetFlowId returns the FlowId field value if set, zero value otherwise.
func (o *AlipayEcoSignFlowFinishModel) GetFlowId() string {
	if o == nil || IsNil(o.FlowId) {
		var ret string
		return ret
	}
	return *o.FlowId
}

// GetFlowIdOk returns a tuple with the FlowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoSignFlowFinishModel) GetFlowIdOk() (*string, bool) {
	if o == nil || IsNil(o.FlowId) {
		return nil, false
	}
	return o.FlowId, true
}

// HasFlowId returns a boolean if a field has been set.
func (o *AlipayEcoSignFlowFinishModel) HasFlowId() bool {
	if o != nil && !IsNil(o.FlowId) {
		return true
	}

	return false
}

// SetFlowId gets a reference to the given string and assigns it to the FlowId field.
func (o *AlipayEcoSignFlowFinishModel) SetFlowId(v string) {
	o.FlowId = &v
}

func (o AlipayEcoSignFlowFinishModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEcoSignFlowFinishModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FlowId) {
		toSerialize["flow_id"] = o.FlowId
	}
	return toSerialize, nil
}

type NullableAlipayEcoSignFlowFinishModel struct {
	value *AlipayEcoSignFlowFinishModel
	isSet bool
}

func (v NullableAlipayEcoSignFlowFinishModel) Get() *AlipayEcoSignFlowFinishModel {
	return v.value
}

func (v *NullableAlipayEcoSignFlowFinishModel) Set(val *AlipayEcoSignFlowFinishModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEcoSignFlowFinishModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEcoSignFlowFinishModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEcoSignFlowFinishModel(val *AlipayEcoSignFlowFinishModel) *NullableAlipayEcoSignFlowFinishModel {
	return &NullableAlipayEcoSignFlowFinishModel{value: val, isSet: true}
}

func (v NullableAlipayEcoSignFlowFinishModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEcoSignFlowFinishModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

