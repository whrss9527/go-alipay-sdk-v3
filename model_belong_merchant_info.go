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

// checks if the BelongMerchantInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BelongMerchantInfo{}

// BelongMerchantInfo struct for BelongMerchantInfo
type BelongMerchantInfo struct {
	// 合作业务类型,默认为ISV_FOR_MERCHANT
	BusinessType *string `json:"business_type,omitempty"`
	// 商户PID,默认为当前接口调用商户。
	MerchantId *string `json:"merchant_id,omitempty"`
	// 商户id类型。特别说明：如果merchant_id_type选择SMID。则表明当前商家券活动的归属者是该SMID所绑定的支付宝PID。因此要求该SMID必须绑定相应的支付宝PID。
	MerchantIdType *string `json:"merchant_id_type,omitempty"`
}

// NewBelongMerchantInfo instantiates a new BelongMerchantInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBelongMerchantInfo() *BelongMerchantInfo {
	this := BelongMerchantInfo{}
	return &this
}

// NewBelongMerchantInfoWithDefaults instantiates a new BelongMerchantInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBelongMerchantInfoWithDefaults() *BelongMerchantInfo {
	this := BelongMerchantInfo{}
	return &this
}

// GetBusinessType returns the BusinessType field value if set, zero value otherwise.
func (o *BelongMerchantInfo) GetBusinessType() string {
	if o == nil || IsNil(o.BusinessType) {
		var ret string
		return ret
	}
	return *o.BusinessType
}

// GetBusinessTypeOk returns a tuple with the BusinessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BelongMerchantInfo) GetBusinessTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessType) {
		return nil, false
	}
	return o.BusinessType, true
}

// HasBusinessType returns a boolean if a field has been set.
func (o *BelongMerchantInfo) HasBusinessType() bool {
	if o != nil && !IsNil(o.BusinessType) {
		return true
	}

	return false
}

// SetBusinessType gets a reference to the given string and assigns it to the BusinessType field.
func (o *BelongMerchantInfo) SetBusinessType(v string) {
	o.BusinessType = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *BelongMerchantInfo) GetMerchantId() string {
	if o == nil || IsNil(o.MerchantId) {
		var ret string
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BelongMerchantInfo) GetMerchantIdOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *BelongMerchantInfo) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given string and assigns it to the MerchantId field.
func (o *BelongMerchantInfo) SetMerchantId(v string) {
	o.MerchantId = &v
}

// GetMerchantIdType returns the MerchantIdType field value if set, zero value otherwise.
func (o *BelongMerchantInfo) GetMerchantIdType() string {
	if o == nil || IsNil(o.MerchantIdType) {
		var ret string
		return ret
	}
	return *o.MerchantIdType
}

// GetMerchantIdTypeOk returns a tuple with the MerchantIdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BelongMerchantInfo) GetMerchantIdTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantIdType) {
		return nil, false
	}
	return o.MerchantIdType, true
}

// HasMerchantIdType returns a boolean if a field has been set.
func (o *BelongMerchantInfo) HasMerchantIdType() bool {
	if o != nil && !IsNil(o.MerchantIdType) {
		return true
	}

	return false
}

// SetMerchantIdType gets a reference to the given string and assigns it to the MerchantIdType field.
func (o *BelongMerchantInfo) SetMerchantIdType(v string) {
	o.MerchantIdType = &v
}

func (o BelongMerchantInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BelongMerchantInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusinessType) {
		toSerialize["business_type"] = o.BusinessType
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchant_id"] = o.MerchantId
	}
	if !IsNil(o.MerchantIdType) {
		toSerialize["merchant_id_type"] = o.MerchantIdType
	}
	return toSerialize, nil
}

type NullableBelongMerchantInfo struct {
	value *BelongMerchantInfo
	isSet bool
}

func (v NullableBelongMerchantInfo) Get() *BelongMerchantInfo {
	return v.value
}

func (v *NullableBelongMerchantInfo) Set(val *BelongMerchantInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBelongMerchantInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBelongMerchantInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBelongMerchantInfo(val *BelongMerchantInfo) *NullableBelongMerchantInfo {
	return &NullableBelongMerchantInfo{value: val, isSet: true}
}

func (v NullableBelongMerchantInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBelongMerchantInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


