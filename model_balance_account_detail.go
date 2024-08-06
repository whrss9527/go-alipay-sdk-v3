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

// checks if the BalanceAccountDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BalanceAccountDetail{}

// BalanceAccountDetail struct for BalanceAccountDetail
type BalanceAccountDetail struct {
	// acs余额，单位：元
	Acs *string `json:"acs,omitempty"`
	// bank余额，单位：元
	Bank *string `json:"bank,omitempty"`
}

// NewBalanceAccountDetail instantiates a new BalanceAccountDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalanceAccountDetail() *BalanceAccountDetail {
	this := BalanceAccountDetail{}
	return &this
}

// NewBalanceAccountDetailWithDefaults instantiates a new BalanceAccountDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceAccountDetailWithDefaults() *BalanceAccountDetail {
	this := BalanceAccountDetail{}
	return &this
}

// GetAcs returns the Acs field value if set, zero value otherwise.
func (o *BalanceAccountDetail) GetAcs() string {
	if o == nil || IsNil(o.Acs) {
		var ret string
		return ret
	}
	return *o.Acs
}

// GetAcsOk returns a tuple with the Acs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceAccountDetail) GetAcsOk() (*string, bool) {
	if o == nil || IsNil(o.Acs) {
		return nil, false
	}
	return o.Acs, true
}

// HasAcs returns a boolean if a field has been set.
func (o *BalanceAccountDetail) HasAcs() bool {
	if o != nil && !IsNil(o.Acs) {
		return true
	}

	return false
}

// SetAcs gets a reference to the given string and assigns it to the Acs field.
func (o *BalanceAccountDetail) SetAcs(v string) {
	o.Acs = &v
}

// GetBank returns the Bank field value if set, zero value otherwise.
func (o *BalanceAccountDetail) GetBank() string {
	if o == nil || IsNil(o.Bank) {
		var ret string
		return ret
	}
	return *o.Bank
}

// GetBankOk returns a tuple with the Bank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceAccountDetail) GetBankOk() (*string, bool) {
	if o == nil || IsNil(o.Bank) {
		return nil, false
	}
	return o.Bank, true
}

// HasBank returns a boolean if a field has been set.
func (o *BalanceAccountDetail) HasBank() bool {
	if o != nil && !IsNil(o.Bank) {
		return true
	}

	return false
}

// SetBank gets a reference to the given string and assigns it to the Bank field.
func (o *BalanceAccountDetail) SetBank(v string) {
	o.Bank = &v
}

func (o BalanceAccountDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BalanceAccountDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Acs) {
		toSerialize["acs"] = o.Acs
	}
	if !IsNil(o.Bank) {
		toSerialize["bank"] = o.Bank
	}
	return toSerialize, nil
}

type NullableBalanceAccountDetail struct {
	value *BalanceAccountDetail
	isSet bool
}

func (v NullableBalanceAccountDetail) Get() *BalanceAccountDetail {
	return v.value
}

func (v *NullableBalanceAccountDetail) Set(val *BalanceAccountDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceAccountDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceAccountDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceAccountDetail(val *BalanceAccountDetail) *NullableBalanceAccountDetail {
	return &NullableBalanceAccountDetail{value: val, isSet: true}
}

func (v NullableBalanceAccountDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalanceAccountDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
