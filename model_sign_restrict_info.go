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

// checks if the SignRestrictInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignRestrictInfo{}

// SignRestrictInfo struct for SignRestrictInfo
type SignRestrictInfo struct {
	// 产品码，支付宝定义的产品码
	RestrictProduct *string `json:"restrict_product,omitempty"`
	// 受限的原因
	RestrictReason *string `json:"restrict_reason,omitempty"`
	// 受限原因编码
	RestrictReasonCode *string `json:"restrict_reason_code,omitempty"`
	// 受限版本，02-受限版 04-极速版
	RestrictVersion *string `json:"restrict_version,omitempty"`
}

// NewSignRestrictInfo instantiates a new SignRestrictInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignRestrictInfo() *SignRestrictInfo {
	this := SignRestrictInfo{}
	return &this
}

// NewSignRestrictInfoWithDefaults instantiates a new SignRestrictInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignRestrictInfoWithDefaults() *SignRestrictInfo {
	this := SignRestrictInfo{}
	return &this
}

// GetRestrictProduct returns the RestrictProduct field value if set, zero value otherwise.
func (o *SignRestrictInfo) GetRestrictProduct() string {
	if o == nil || IsNil(o.RestrictProduct) {
		var ret string
		return ret
	}
	return *o.RestrictProduct
}

// GetRestrictProductOk returns a tuple with the RestrictProduct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignRestrictInfo) GetRestrictProductOk() (*string, bool) {
	if o == nil || IsNil(o.RestrictProduct) {
		return nil, false
	}
	return o.RestrictProduct, true
}

// HasRestrictProduct returns a boolean if a field has been set.
func (o *SignRestrictInfo) HasRestrictProduct() bool {
	if o != nil && !IsNil(o.RestrictProduct) {
		return true
	}

	return false
}

// SetRestrictProduct gets a reference to the given string and assigns it to the RestrictProduct field.
func (o *SignRestrictInfo) SetRestrictProduct(v string) {
	o.RestrictProduct = &v
}

// GetRestrictReason returns the RestrictReason field value if set, zero value otherwise.
func (o *SignRestrictInfo) GetRestrictReason() string {
	if o == nil || IsNil(o.RestrictReason) {
		var ret string
		return ret
	}
	return *o.RestrictReason
}

// GetRestrictReasonOk returns a tuple with the RestrictReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignRestrictInfo) GetRestrictReasonOk() (*string, bool) {
	if o == nil || IsNil(o.RestrictReason) {
		return nil, false
	}
	return o.RestrictReason, true
}

// HasRestrictReason returns a boolean if a field has been set.
func (o *SignRestrictInfo) HasRestrictReason() bool {
	if o != nil && !IsNil(o.RestrictReason) {
		return true
	}

	return false
}

// SetRestrictReason gets a reference to the given string and assigns it to the RestrictReason field.
func (o *SignRestrictInfo) SetRestrictReason(v string) {
	o.RestrictReason = &v
}

// GetRestrictReasonCode returns the RestrictReasonCode field value if set, zero value otherwise.
func (o *SignRestrictInfo) GetRestrictReasonCode() string {
	if o == nil || IsNil(o.RestrictReasonCode) {
		var ret string
		return ret
	}
	return *o.RestrictReasonCode
}

// GetRestrictReasonCodeOk returns a tuple with the RestrictReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignRestrictInfo) GetRestrictReasonCodeOk() (*string, bool) {
	if o == nil || IsNil(o.RestrictReasonCode) {
		return nil, false
	}
	return o.RestrictReasonCode, true
}

// HasRestrictReasonCode returns a boolean if a field has been set.
func (o *SignRestrictInfo) HasRestrictReasonCode() bool {
	if o != nil && !IsNil(o.RestrictReasonCode) {
		return true
	}

	return false
}

// SetRestrictReasonCode gets a reference to the given string and assigns it to the RestrictReasonCode field.
func (o *SignRestrictInfo) SetRestrictReasonCode(v string) {
	o.RestrictReasonCode = &v
}

// GetRestrictVersion returns the RestrictVersion field value if set, zero value otherwise.
func (o *SignRestrictInfo) GetRestrictVersion() string {
	if o == nil || IsNil(o.RestrictVersion) {
		var ret string
		return ret
	}
	return *o.RestrictVersion
}

// GetRestrictVersionOk returns a tuple with the RestrictVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignRestrictInfo) GetRestrictVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RestrictVersion) {
		return nil, false
	}
	return o.RestrictVersion, true
}

// HasRestrictVersion returns a boolean if a field has been set.
func (o *SignRestrictInfo) HasRestrictVersion() bool {
	if o != nil && !IsNil(o.RestrictVersion) {
		return true
	}

	return false
}

// SetRestrictVersion gets a reference to the given string and assigns it to the RestrictVersion field.
func (o *SignRestrictInfo) SetRestrictVersion(v string) {
	o.RestrictVersion = &v
}

func (o SignRestrictInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignRestrictInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RestrictProduct) {
		toSerialize["restrict_product"] = o.RestrictProduct
	}
	if !IsNil(o.RestrictReason) {
		toSerialize["restrict_reason"] = o.RestrictReason
	}
	if !IsNil(o.RestrictReasonCode) {
		toSerialize["restrict_reason_code"] = o.RestrictReasonCode
	}
	if !IsNil(o.RestrictVersion) {
		toSerialize["restrict_version"] = o.RestrictVersion
	}
	return toSerialize, nil
}

type NullableSignRestrictInfo struct {
	value *SignRestrictInfo
	isSet bool
}

func (v NullableSignRestrictInfo) Get() *SignRestrictInfo {
	return v.value
}

func (v *NullableSignRestrictInfo) Set(val *SignRestrictInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSignRestrictInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSignRestrictInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignRestrictInfo(val *SignRestrictInfo) *NullableSignRestrictInfo {
	return &NullableSignRestrictInfo{value: val, isSet: true}
}

func (v NullableSignRestrictInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignRestrictInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


