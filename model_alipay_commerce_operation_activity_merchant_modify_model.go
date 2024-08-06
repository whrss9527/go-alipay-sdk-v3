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

// checks if the AlipayCommerceOperationActivityMerchantModifyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayCommerceOperationActivityMerchantModifyModel{}

// AlipayCommerceOperationActivityMerchantModifyModel struct for AlipayCommerceOperationActivityMerchantModifyModel
type AlipayCommerceOperationActivityMerchantModifyModel struct {
	// 消费门槛，单位分
	ConsumptionThreshold *int32 `json:"consumption_threshold,omitempty"`
	// 优惠金额，单位分
	DiscountAmount *int32 `json:"discount_amount,omitempty"`
	// NORMAL 常规类型，不限制, RESTRICTED 限制报名条件
	Type *string `json:"type,omitempty"`
}

// NewAlipayCommerceOperationActivityMerchantModifyModel instantiates a new AlipayCommerceOperationActivityMerchantModifyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayCommerceOperationActivityMerchantModifyModel() *AlipayCommerceOperationActivityMerchantModifyModel {
	this := AlipayCommerceOperationActivityMerchantModifyModel{}
	return &this
}

// NewAlipayCommerceOperationActivityMerchantModifyModelWithDefaults instantiates a new AlipayCommerceOperationActivityMerchantModifyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayCommerceOperationActivityMerchantModifyModelWithDefaults() *AlipayCommerceOperationActivityMerchantModifyModel {
	this := AlipayCommerceOperationActivityMerchantModifyModel{}
	return &this
}

// GetConsumptionThreshold returns the ConsumptionThreshold field value if set, zero value otherwise.
func (o *AlipayCommerceOperationActivityMerchantModifyModel) GetConsumptionThreshold() int32 {
	if o == nil || IsNil(o.ConsumptionThreshold) {
		var ret int32
		return ret
	}
	return *o.ConsumptionThreshold
}

// GetConsumptionThresholdOk returns a tuple with the ConsumptionThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceOperationActivityMerchantModifyModel) GetConsumptionThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.ConsumptionThreshold) {
		return nil, false
	}
	return o.ConsumptionThreshold, true
}

// HasConsumptionThreshold returns a boolean if a field has been set.
func (o *AlipayCommerceOperationActivityMerchantModifyModel) HasConsumptionThreshold() bool {
	if o != nil && !IsNil(o.ConsumptionThreshold) {
		return true
	}

	return false
}

// SetConsumptionThreshold gets a reference to the given int32 and assigns it to the ConsumptionThreshold field.
func (o *AlipayCommerceOperationActivityMerchantModifyModel) SetConsumptionThreshold(v int32) {
	o.ConsumptionThreshold = &v
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise.
func (o *AlipayCommerceOperationActivityMerchantModifyModel) GetDiscountAmount() int32 {
	if o == nil || IsNil(o.DiscountAmount) {
		var ret int32
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceOperationActivityMerchantModifyModel) GetDiscountAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.DiscountAmount) {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *AlipayCommerceOperationActivityMerchantModifyModel) HasDiscountAmount() bool {
	if o != nil && !IsNil(o.DiscountAmount) {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given int32 and assigns it to the DiscountAmount field.
func (o *AlipayCommerceOperationActivityMerchantModifyModel) SetDiscountAmount(v int32) {
	o.DiscountAmount = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AlipayCommerceOperationActivityMerchantModifyModel) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceOperationActivityMerchantModifyModel) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AlipayCommerceOperationActivityMerchantModifyModel) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AlipayCommerceOperationActivityMerchantModifyModel) SetType(v string) {
	o.Type = &v
}

func (o AlipayCommerceOperationActivityMerchantModifyModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayCommerceOperationActivityMerchantModifyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConsumptionThreshold) {
		toSerialize["consumption_threshold"] = o.ConsumptionThreshold
	}
	if !IsNil(o.DiscountAmount) {
		toSerialize["discount_amount"] = o.DiscountAmount
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAlipayCommerceOperationActivityMerchantModifyModel struct {
	value *AlipayCommerceOperationActivityMerchantModifyModel
	isSet bool
}

func (v NullableAlipayCommerceOperationActivityMerchantModifyModel) Get() *AlipayCommerceOperationActivityMerchantModifyModel {
	return v.value
}

func (v *NullableAlipayCommerceOperationActivityMerchantModifyModel) Set(val *AlipayCommerceOperationActivityMerchantModifyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceOperationActivityMerchantModifyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceOperationActivityMerchantModifyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceOperationActivityMerchantModifyModel(val *AlipayCommerceOperationActivityMerchantModifyModel) *NullableAlipayCommerceOperationActivityMerchantModifyModel {
	return &NullableAlipayCommerceOperationActivityMerchantModifyModel{value: val, isSet: true}
}

func (v NullableAlipayCommerceOperationActivityMerchantModifyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceOperationActivityMerchantModifyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
