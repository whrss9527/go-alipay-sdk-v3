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

// checks if the AlipayMarketingCardQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingCardQueryResponseModel{}

// AlipayMarketingCardQueryResponseModel struct for AlipayMarketingCardQueryResponseModel
type AlipayMarketingCardQueryResponseModel struct {
	CardInfo *MerchantCard `json:"card_info,omitempty"`
	PaidOuterCardInfo *PaidOuterCardExtraInfoDTO `json:"paid_outer_card_info,omitempty"`
	// 商家会员卡页面跳转到支付宝卡券详情页的pass_id，对应schema地址中的参数p， 主要用于小程序跳往会员卡详情的链接拼接
	PassId *string `json:"pass_id,omitempty"`
	// 商家会员卡页面跳转到支付宝卡券详情页面的schema地址
	SchemaUrl *string `json:"schema_url,omitempty"`
}

// NewAlipayMarketingCardQueryResponseModel instantiates a new AlipayMarketingCardQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingCardQueryResponseModel() *AlipayMarketingCardQueryResponseModel {
	this := AlipayMarketingCardQueryResponseModel{}
	return &this
}

// NewAlipayMarketingCardQueryResponseModelWithDefaults instantiates a new AlipayMarketingCardQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingCardQueryResponseModelWithDefaults() *AlipayMarketingCardQueryResponseModel {
	this := AlipayMarketingCardQueryResponseModel{}
	return &this
}

// GetCardInfo returns the CardInfo field value if set, zero value otherwise.
func (o *AlipayMarketingCardQueryResponseModel) GetCardInfo() MerchantCard {
	if o == nil || IsNil(o.CardInfo) {
		var ret MerchantCard
		return ret
	}
	return *o.CardInfo
}

// GetCardInfoOk returns a tuple with the CardInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardQueryResponseModel) GetCardInfoOk() (*MerchantCard, bool) {
	if o == nil || IsNil(o.CardInfo) {
		return nil, false
	}
	return o.CardInfo, true
}

// HasCardInfo returns a boolean if a field has been set.
func (o *AlipayMarketingCardQueryResponseModel) HasCardInfo() bool {
	if o != nil && !IsNil(o.CardInfo) {
		return true
	}

	return false
}

// SetCardInfo gets a reference to the given MerchantCard and assigns it to the CardInfo field.
func (o *AlipayMarketingCardQueryResponseModel) SetCardInfo(v MerchantCard) {
	o.CardInfo = &v
}

// GetPaidOuterCardInfo returns the PaidOuterCardInfo field value if set, zero value otherwise.
func (o *AlipayMarketingCardQueryResponseModel) GetPaidOuterCardInfo() PaidOuterCardExtraInfoDTO {
	if o == nil || IsNil(o.PaidOuterCardInfo) {
		var ret PaidOuterCardExtraInfoDTO
		return ret
	}
	return *o.PaidOuterCardInfo
}

// GetPaidOuterCardInfoOk returns a tuple with the PaidOuterCardInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardQueryResponseModel) GetPaidOuterCardInfoOk() (*PaidOuterCardExtraInfoDTO, bool) {
	if o == nil || IsNil(o.PaidOuterCardInfo) {
		return nil, false
	}
	return o.PaidOuterCardInfo, true
}

// HasPaidOuterCardInfo returns a boolean if a field has been set.
func (o *AlipayMarketingCardQueryResponseModel) HasPaidOuterCardInfo() bool {
	if o != nil && !IsNil(o.PaidOuterCardInfo) {
		return true
	}

	return false
}

// SetPaidOuterCardInfo gets a reference to the given PaidOuterCardExtraInfoDTO and assigns it to the PaidOuterCardInfo field.
func (o *AlipayMarketingCardQueryResponseModel) SetPaidOuterCardInfo(v PaidOuterCardExtraInfoDTO) {
	o.PaidOuterCardInfo = &v
}

// GetPassId returns the PassId field value if set, zero value otherwise.
func (o *AlipayMarketingCardQueryResponseModel) GetPassId() string {
	if o == nil || IsNil(o.PassId) {
		var ret string
		return ret
	}
	return *o.PassId
}

// GetPassIdOk returns a tuple with the PassId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardQueryResponseModel) GetPassIdOk() (*string, bool) {
	if o == nil || IsNil(o.PassId) {
		return nil, false
	}
	return o.PassId, true
}

// HasPassId returns a boolean if a field has been set.
func (o *AlipayMarketingCardQueryResponseModel) HasPassId() bool {
	if o != nil && !IsNil(o.PassId) {
		return true
	}

	return false
}

// SetPassId gets a reference to the given string and assigns it to the PassId field.
func (o *AlipayMarketingCardQueryResponseModel) SetPassId(v string) {
	o.PassId = &v
}

// GetSchemaUrl returns the SchemaUrl field value if set, zero value otherwise.
func (o *AlipayMarketingCardQueryResponseModel) GetSchemaUrl() string {
	if o == nil || IsNil(o.SchemaUrl) {
		var ret string
		return ret
	}
	return *o.SchemaUrl
}

// GetSchemaUrlOk returns a tuple with the SchemaUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardQueryResponseModel) GetSchemaUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaUrl) {
		return nil, false
	}
	return o.SchemaUrl, true
}

// HasSchemaUrl returns a boolean if a field has been set.
func (o *AlipayMarketingCardQueryResponseModel) HasSchemaUrl() bool {
	if o != nil && !IsNil(o.SchemaUrl) {
		return true
	}

	return false
}

// SetSchemaUrl gets a reference to the given string and assigns it to the SchemaUrl field.
func (o *AlipayMarketingCardQueryResponseModel) SetSchemaUrl(v string) {
	o.SchemaUrl = &v
}

func (o AlipayMarketingCardQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingCardQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CardInfo) {
		toSerialize["card_info"] = o.CardInfo
	}
	if !IsNil(o.PaidOuterCardInfo) {
		toSerialize["paid_outer_card_info"] = o.PaidOuterCardInfo
	}
	if !IsNil(o.PassId) {
		toSerialize["pass_id"] = o.PassId
	}
	if !IsNil(o.SchemaUrl) {
		toSerialize["schema_url"] = o.SchemaUrl
	}
	return toSerialize, nil
}

type NullableAlipayMarketingCardQueryResponseModel struct {
	value *AlipayMarketingCardQueryResponseModel
	isSet bool
}

func (v NullableAlipayMarketingCardQueryResponseModel) Get() *AlipayMarketingCardQueryResponseModel {
	return v.value
}

func (v *NullableAlipayMarketingCardQueryResponseModel) Set(val *AlipayMarketingCardQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingCardQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingCardQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingCardQueryResponseModel(val *AlipayMarketingCardQueryResponseModel) *NullableAlipayMarketingCardQueryResponseModel {
	return &NullableAlipayMarketingCardQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingCardQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingCardQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


