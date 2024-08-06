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

// checks if the AlipayOpenPublicGroupBatchqueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicGroupBatchqueryResponseModel{}

// AlipayOpenPublicGroupBatchqueryResponseModel struct for AlipayOpenPublicGroupBatchqueryResponseModel
type AlipayOpenPublicGroupBatchqueryResponseModel struct {
	// 用户分组列表，包含每个分组的id、name、以及规则模型
	Groups []QueryGroup `json:"groups,omitempty"`
}

// NewAlipayOpenPublicGroupBatchqueryResponseModel instantiates a new AlipayOpenPublicGroupBatchqueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicGroupBatchqueryResponseModel() *AlipayOpenPublicGroupBatchqueryResponseModel {
	this := AlipayOpenPublicGroupBatchqueryResponseModel{}
	return &this
}

// NewAlipayOpenPublicGroupBatchqueryResponseModelWithDefaults instantiates a new AlipayOpenPublicGroupBatchqueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicGroupBatchqueryResponseModelWithDefaults() *AlipayOpenPublicGroupBatchqueryResponseModel {
	this := AlipayOpenPublicGroupBatchqueryResponseModel{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *AlipayOpenPublicGroupBatchqueryResponseModel) GetGroups() []QueryGroup {
	if o == nil || IsNil(o.Groups) {
		var ret []QueryGroup
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicGroupBatchqueryResponseModel) GetGroupsOk() ([]QueryGroup, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *AlipayOpenPublicGroupBatchqueryResponseModel) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []QueryGroup and assigns it to the Groups field.
func (o *AlipayOpenPublicGroupBatchqueryResponseModel) SetGroups(v []QueryGroup) {
	o.Groups = v
}

func (o AlipayOpenPublicGroupBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicGroupBatchqueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicGroupBatchqueryResponseModel struct {
	value *AlipayOpenPublicGroupBatchqueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenPublicGroupBatchqueryResponseModel) Get() *AlipayOpenPublicGroupBatchqueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenPublicGroupBatchqueryResponseModel) Set(val *AlipayOpenPublicGroupBatchqueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicGroupBatchqueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicGroupBatchqueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicGroupBatchqueryResponseModel(val *AlipayOpenPublicGroupBatchqueryResponseModel) *NullableAlipayOpenPublicGroupBatchqueryResponseModel {
	return &NullableAlipayOpenPublicGroupBatchqueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicGroupBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicGroupBatchqueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
