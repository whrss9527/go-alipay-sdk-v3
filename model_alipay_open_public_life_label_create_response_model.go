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

// checks if the AlipayOpenPublicLifeLabelCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicLifeLabelCreateResponseModel{}

// AlipayOpenPublicLifeLabelCreateResponseModel struct for AlipayOpenPublicLifeLabelCreateResponseModel
type AlipayOpenPublicLifeLabelCreateResponseModel struct {
	// 标签id
	LabelId *string `json:"label_id,omitempty"`
}

// NewAlipayOpenPublicLifeLabelCreateResponseModel instantiates a new AlipayOpenPublicLifeLabelCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicLifeLabelCreateResponseModel() *AlipayOpenPublicLifeLabelCreateResponseModel {
	this := AlipayOpenPublicLifeLabelCreateResponseModel{}
	return &this
}

// NewAlipayOpenPublicLifeLabelCreateResponseModelWithDefaults instantiates a new AlipayOpenPublicLifeLabelCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicLifeLabelCreateResponseModelWithDefaults() *AlipayOpenPublicLifeLabelCreateResponseModel {
	this := AlipayOpenPublicLifeLabelCreateResponseModel{}
	return &this
}

// GetLabelId returns the LabelId field value if set, zero value otherwise.
func (o *AlipayOpenPublicLifeLabelCreateResponseModel) GetLabelId() string {
	if o == nil || IsNil(o.LabelId) {
		var ret string
		return ret
	}
	return *o.LabelId
}

// GetLabelIdOk returns a tuple with the LabelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicLifeLabelCreateResponseModel) GetLabelIdOk() (*string, bool) {
	if o == nil || IsNil(o.LabelId) {
		return nil, false
	}
	return o.LabelId, true
}

// HasLabelId returns a boolean if a field has been set.
func (o *AlipayOpenPublicLifeLabelCreateResponseModel) HasLabelId() bool {
	if o != nil && !IsNil(o.LabelId) {
		return true
	}

	return false
}

// SetLabelId gets a reference to the given string and assigns it to the LabelId field.
func (o *AlipayOpenPublicLifeLabelCreateResponseModel) SetLabelId(v string) {
	o.LabelId = &v
}

func (o AlipayOpenPublicLifeLabelCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicLifeLabelCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LabelId) {
		toSerialize["label_id"] = o.LabelId
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicLifeLabelCreateResponseModel struct {
	value *AlipayOpenPublicLifeLabelCreateResponseModel
	isSet bool
}

func (v NullableAlipayOpenPublicLifeLabelCreateResponseModel) Get() *AlipayOpenPublicLifeLabelCreateResponseModel {
	return v.value
}

func (v *NullableAlipayOpenPublicLifeLabelCreateResponseModel) Set(val *AlipayOpenPublicLifeLabelCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicLifeLabelCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicLifeLabelCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicLifeLabelCreateResponseModel(val *AlipayOpenPublicLifeLabelCreateResponseModel) *NullableAlipayOpenPublicLifeLabelCreateResponseModel {
	return &NullableAlipayOpenPublicLifeLabelCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicLifeLabelCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicLifeLabelCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


