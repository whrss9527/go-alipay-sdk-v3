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

// checks if the AlipayOpenMiniInnerappPluginservicePublishResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniInnerappPluginservicePublishResponseModel{}

// AlipayOpenMiniInnerappPluginservicePublishResponseModel struct for AlipayOpenMiniInnerappPluginservicePublishResponseModel
type AlipayOpenMiniInnerappPluginservicePublishResponseModel struct {
	// PL2012211102012056
	AbilityId *string `json:"ability_id,omitempty"`
}

// NewAlipayOpenMiniInnerappPluginservicePublishResponseModel instantiates a new AlipayOpenMiniInnerappPluginservicePublishResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniInnerappPluginservicePublishResponseModel() *AlipayOpenMiniInnerappPluginservicePublishResponseModel {
	this := AlipayOpenMiniInnerappPluginservicePublishResponseModel{}
	return &this
}

// NewAlipayOpenMiniInnerappPluginservicePublishResponseModelWithDefaults instantiates a new AlipayOpenMiniInnerappPluginservicePublishResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniInnerappPluginservicePublishResponseModelWithDefaults() *AlipayOpenMiniInnerappPluginservicePublishResponseModel {
	this := AlipayOpenMiniInnerappPluginservicePublishResponseModel{}
	return &this
}

// GetAbilityId returns the AbilityId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerappPluginservicePublishResponseModel) GetAbilityId() string {
	if o == nil || IsNil(o.AbilityId) {
		var ret string
		return ret
	}
	return *o.AbilityId
}

// GetAbilityIdOk returns a tuple with the AbilityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerappPluginservicePublishResponseModel) GetAbilityIdOk() (*string, bool) {
	if o == nil || IsNil(o.AbilityId) {
		return nil, false
	}
	return o.AbilityId, true
}

// HasAbilityId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerappPluginservicePublishResponseModel) HasAbilityId() bool {
	if o != nil && !IsNil(o.AbilityId) {
		return true
	}

	return false
}

// SetAbilityId gets a reference to the given string and assigns it to the AbilityId field.
func (o *AlipayOpenMiniInnerappPluginservicePublishResponseModel) SetAbilityId(v string) {
	o.AbilityId = &v
}

func (o AlipayOpenMiniInnerappPluginservicePublishResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniInnerappPluginservicePublishResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AbilityId) {
		toSerialize["ability_id"] = o.AbilityId
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniInnerappPluginservicePublishResponseModel struct {
	value *AlipayOpenMiniInnerappPluginservicePublishResponseModel
	isSet bool
}

func (v NullableAlipayOpenMiniInnerappPluginservicePublishResponseModel) Get() *AlipayOpenMiniInnerappPluginservicePublishResponseModel {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerappPluginservicePublishResponseModel) Set(val *AlipayOpenMiniInnerappPluginservicePublishResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerappPluginservicePublishResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerappPluginservicePublishResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerappPluginservicePublishResponseModel(val *AlipayOpenMiniInnerappPluginservicePublishResponseModel) *NullableAlipayOpenMiniInnerappPluginservicePublishResponseModel {
	return &NullableAlipayOpenMiniInnerappPluginservicePublishResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerappPluginservicePublishResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerappPluginservicePublishResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
