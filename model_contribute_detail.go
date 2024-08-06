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

// checks if the ContributeDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContributeDetail{}

// ContributeDetail struct for ContributeDetail
type ContributeDetail struct {
	// 出资方金额
	ContributeAmount *string `json:"contribute_amount,omitempty"`
	// 出资方类型
	ContributeType *string `json:"contribute_type,omitempty"`
}

// NewContributeDetail instantiates a new ContributeDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContributeDetail() *ContributeDetail {
	this := ContributeDetail{}
	return &this
}

// NewContributeDetailWithDefaults instantiates a new ContributeDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContributeDetailWithDefaults() *ContributeDetail {
	this := ContributeDetail{}
	return &this
}

// GetContributeAmount returns the ContributeAmount field value if set, zero value otherwise.
func (o *ContributeDetail) GetContributeAmount() string {
	if o == nil || IsNil(o.ContributeAmount) {
		var ret string
		return ret
	}
	return *o.ContributeAmount
}

// GetContributeAmountOk returns a tuple with the ContributeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContributeDetail) GetContributeAmountOk() (*string, bool) {
	if o == nil || IsNil(o.ContributeAmount) {
		return nil, false
	}
	return o.ContributeAmount, true
}

// HasContributeAmount returns a boolean if a field has been set.
func (o *ContributeDetail) HasContributeAmount() bool {
	if o != nil && !IsNil(o.ContributeAmount) {
		return true
	}

	return false
}

// SetContributeAmount gets a reference to the given string and assigns it to the ContributeAmount field.
func (o *ContributeDetail) SetContributeAmount(v string) {
	o.ContributeAmount = &v
}

// GetContributeType returns the ContributeType field value if set, zero value otherwise.
func (o *ContributeDetail) GetContributeType() string {
	if o == nil || IsNil(o.ContributeType) {
		var ret string
		return ret
	}
	return *o.ContributeType
}

// GetContributeTypeOk returns a tuple with the ContributeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContributeDetail) GetContributeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContributeType) {
		return nil, false
	}
	return o.ContributeType, true
}

// HasContributeType returns a boolean if a field has been set.
func (o *ContributeDetail) HasContributeType() bool {
	if o != nil && !IsNil(o.ContributeType) {
		return true
	}

	return false
}

// SetContributeType gets a reference to the given string and assigns it to the ContributeType field.
func (o *ContributeDetail) SetContributeType(v string) {
	o.ContributeType = &v
}

func (o ContributeDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContributeDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContributeAmount) {
		toSerialize["contribute_amount"] = o.ContributeAmount
	}
	if !IsNil(o.ContributeType) {
		toSerialize["contribute_type"] = o.ContributeType
	}
	return toSerialize, nil
}

type NullableContributeDetail struct {
	value *ContributeDetail
	isSet bool
}

func (v NullableContributeDetail) Get() *ContributeDetail {
	return v.value
}

func (v *NullableContributeDetail) Set(val *ContributeDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableContributeDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableContributeDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContributeDetail(val *ContributeDetail) *NullableContributeDetail {
	return &NullableContributeDetail{value: val, isSet: true}
}

func (v NullableContributeDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContributeDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

