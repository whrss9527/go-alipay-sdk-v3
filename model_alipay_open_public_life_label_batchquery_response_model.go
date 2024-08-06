/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alipay

import (
	"encoding/json"
)

// checks if the AlipayOpenPublicLifeLabelBatchqueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicLifeLabelBatchqueryResponseModel{}

// AlipayOpenPublicLifeLabelBatchqueryResponseModel struct for AlipayOpenPublicLifeLabelBatchqueryResponseModel
type AlipayOpenPublicLifeLabelBatchqueryResponseModel struct {
	// 标签列表
	Labels []LifeLabel `json:"labels,omitempty"`
}

// NewAlipayOpenPublicLifeLabelBatchqueryResponseModel instantiates a new AlipayOpenPublicLifeLabelBatchqueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicLifeLabelBatchqueryResponseModel() *AlipayOpenPublicLifeLabelBatchqueryResponseModel {
	this := AlipayOpenPublicLifeLabelBatchqueryResponseModel{}
	return &this
}

// NewAlipayOpenPublicLifeLabelBatchqueryResponseModelWithDefaults instantiates a new AlipayOpenPublicLifeLabelBatchqueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicLifeLabelBatchqueryResponseModelWithDefaults() *AlipayOpenPublicLifeLabelBatchqueryResponseModel {
	this := AlipayOpenPublicLifeLabelBatchqueryResponseModel{}
	return &this
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *AlipayOpenPublicLifeLabelBatchqueryResponseModel) GetLabels() []LifeLabel {
	if o == nil || IsNil(o.Labels) {
		var ret []LifeLabel
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicLifeLabelBatchqueryResponseModel) GetLabelsOk() ([]LifeLabel, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *AlipayOpenPublicLifeLabelBatchqueryResponseModel) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []LifeLabel and assigns it to the Labels field.
func (o *AlipayOpenPublicLifeLabelBatchqueryResponseModel) SetLabels(v []LifeLabel) {
	o.Labels = v
}

func (o AlipayOpenPublicLifeLabelBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicLifeLabelBatchqueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicLifeLabelBatchqueryResponseModel struct {
	value *AlipayOpenPublicLifeLabelBatchqueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenPublicLifeLabelBatchqueryResponseModel) Get() *AlipayOpenPublicLifeLabelBatchqueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenPublicLifeLabelBatchqueryResponseModel) Set(val *AlipayOpenPublicLifeLabelBatchqueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicLifeLabelBatchqueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicLifeLabelBatchqueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicLifeLabelBatchqueryResponseModel(val *AlipayOpenPublicLifeLabelBatchqueryResponseModel) *NullableAlipayOpenPublicLifeLabelBatchqueryResponseModel {
	return &NullableAlipayOpenPublicLifeLabelBatchqueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicLifeLabelBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicLifeLabelBatchqueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
