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

// checks if the AlipayMarketingActivityDeliveryQueryModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingActivityDeliveryQueryModel{}

// AlipayMarketingActivityDeliveryQueryModel struct for AlipayMarketingActivityDeliveryQueryModel
type AlipayMarketingActivityDeliveryQueryModel struct {
	BelongMerchantInfo *DeliveryAgencyMerchantInfo `json:"belong_merchant_info,omitempty"`
	// \"[已废弃] 待查询的投放配置列表。 最大数量限制20个。\"
	DeliveryConfigList []DeliveryConfig `json:"delivery_config_list,omitempty"`
	// 商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
}

// NewAlipayMarketingActivityDeliveryQueryModel instantiates a new AlipayMarketingActivityDeliveryQueryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingActivityDeliveryQueryModel() *AlipayMarketingActivityDeliveryQueryModel {
	this := AlipayMarketingActivityDeliveryQueryModel{}
	return &this
}

// NewAlipayMarketingActivityDeliveryQueryModelWithDefaults instantiates a new AlipayMarketingActivityDeliveryQueryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingActivityDeliveryQueryModelWithDefaults() *AlipayMarketingActivityDeliveryQueryModel {
	this := AlipayMarketingActivityDeliveryQueryModel{}
	return &this
}

// GetBelongMerchantInfo returns the BelongMerchantInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityDeliveryQueryModel) GetBelongMerchantInfo() DeliveryAgencyMerchantInfo {
	if o == nil || IsNil(o.BelongMerchantInfo) {
		var ret DeliveryAgencyMerchantInfo
		return ret
	}
	return *o.BelongMerchantInfo
}

// GetBelongMerchantInfoOk returns a tuple with the BelongMerchantInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityDeliveryQueryModel) GetBelongMerchantInfoOk() (*DeliveryAgencyMerchantInfo, bool) {
	if o == nil || IsNil(o.BelongMerchantInfo) {
		return nil, false
	}
	return o.BelongMerchantInfo, true
}

// HasBelongMerchantInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityDeliveryQueryModel) HasBelongMerchantInfo() bool {
	if o != nil && !IsNil(o.BelongMerchantInfo) {
		return true
	}

	return false
}

// SetBelongMerchantInfo gets a reference to the given DeliveryAgencyMerchantInfo and assigns it to the BelongMerchantInfo field.
func (o *AlipayMarketingActivityDeliveryQueryModel) SetBelongMerchantInfo(v DeliveryAgencyMerchantInfo) {
	o.BelongMerchantInfo = &v
}

// GetDeliveryConfigList returns the DeliveryConfigList field value if set, zero value otherwise.
func (o *AlipayMarketingActivityDeliveryQueryModel) GetDeliveryConfigList() []DeliveryConfig {
	if o == nil || IsNil(o.DeliveryConfigList) {
		var ret []DeliveryConfig
		return ret
	}
	return o.DeliveryConfigList
}

// GetDeliveryConfigListOk returns a tuple with the DeliveryConfigList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityDeliveryQueryModel) GetDeliveryConfigListOk() ([]DeliveryConfig, bool) {
	if o == nil || IsNil(o.DeliveryConfigList) {
		return nil, false
	}
	return o.DeliveryConfigList, true
}

// HasDeliveryConfigList returns a boolean if a field has been set.
func (o *AlipayMarketingActivityDeliveryQueryModel) HasDeliveryConfigList() bool {
	if o != nil && !IsNil(o.DeliveryConfigList) {
		return true
	}

	return false
}

// SetDeliveryConfigList gets a reference to the given []DeliveryConfig and assigns it to the DeliveryConfigList field.
func (o *AlipayMarketingActivityDeliveryQueryModel) SetDeliveryConfigList(v []DeliveryConfig) {
	o.DeliveryConfigList = v
}

// GetMerchantAccessMode returns the MerchantAccessMode field value if set, zero value otherwise.
func (o *AlipayMarketingActivityDeliveryQueryModel) GetMerchantAccessMode() string {
	if o == nil || IsNil(o.MerchantAccessMode) {
		var ret string
		return ret
	}
	return *o.MerchantAccessMode
}

// GetMerchantAccessModeOk returns a tuple with the MerchantAccessMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityDeliveryQueryModel) GetMerchantAccessModeOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantAccessMode) {
		return nil, false
	}
	return o.MerchantAccessMode, true
}

// HasMerchantAccessMode returns a boolean if a field has been set.
func (o *AlipayMarketingActivityDeliveryQueryModel) HasMerchantAccessMode() bool {
	if o != nil && !IsNil(o.MerchantAccessMode) {
		return true
	}

	return false
}

// SetMerchantAccessMode gets a reference to the given string and assigns it to the MerchantAccessMode field.
func (o *AlipayMarketingActivityDeliveryQueryModel) SetMerchantAccessMode(v string) {
	o.MerchantAccessMode = &v
}

func (o AlipayMarketingActivityDeliveryQueryModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingActivityDeliveryQueryModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BelongMerchantInfo) {
		toSerialize["belong_merchant_info"] = o.BelongMerchantInfo
	}
	if !IsNil(o.DeliveryConfigList) {
		toSerialize["delivery_config_list"] = o.DeliveryConfigList
	}
	if !IsNil(o.MerchantAccessMode) {
		toSerialize["merchant_access_mode"] = o.MerchantAccessMode
	}
	return toSerialize, nil
}

type NullableAlipayMarketingActivityDeliveryQueryModel struct {
	value *AlipayMarketingActivityDeliveryQueryModel
	isSet bool
}

func (v NullableAlipayMarketingActivityDeliveryQueryModel) Get() *AlipayMarketingActivityDeliveryQueryModel {
	return v.value
}

func (v *NullableAlipayMarketingActivityDeliveryQueryModel) Set(val *AlipayMarketingActivityDeliveryQueryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityDeliveryQueryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityDeliveryQueryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityDeliveryQueryModel(val *AlipayMarketingActivityDeliveryQueryModel) *NullableAlipayMarketingActivityDeliveryQueryModel {
	return &NullableAlipayMarketingActivityDeliveryQueryModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityDeliveryQueryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityDeliveryQueryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
