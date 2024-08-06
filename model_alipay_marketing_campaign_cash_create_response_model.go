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

// checks if the AlipayMarketingCampaignCashCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingCampaignCashCreateResponseModel{}

// AlipayMarketingCampaignCashCreateResponseModel struct for AlipayMarketingCampaignCashCreateResponseModel
type AlipayMarketingCampaignCashCreateResponseModel struct {
	// 生成的现金红包活动号
	CrowdNo *string `json:"crowd_no,omitempty"`
	// 原始活动号,商户排查问题时提供的活动依据
	OriginCrowdNo *string `json:"origin_crowd_no,omitempty"`
	// 活动创建后的付款链接，返回的是urlencode编码后的字符串。需要先进行urldecode解码，然后在浏览器中进行访问，会先进行支付宝登录引导，然后商户进行付款，付款有效期24小时。
	PayUrl *string `json:"pay_url,omitempty"`
}

// NewAlipayMarketingCampaignCashCreateResponseModel instantiates a new AlipayMarketingCampaignCashCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingCampaignCashCreateResponseModel() *AlipayMarketingCampaignCashCreateResponseModel {
	this := AlipayMarketingCampaignCashCreateResponseModel{}
	return &this
}

// NewAlipayMarketingCampaignCashCreateResponseModelWithDefaults instantiates a new AlipayMarketingCampaignCashCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingCampaignCashCreateResponseModelWithDefaults() *AlipayMarketingCampaignCashCreateResponseModel {
	this := AlipayMarketingCampaignCashCreateResponseModel{}
	return &this
}

// GetCrowdNo returns the CrowdNo field value if set, zero value otherwise.
func (o *AlipayMarketingCampaignCashCreateResponseModel) GetCrowdNo() string {
	if o == nil || IsNil(o.CrowdNo) {
		var ret string
		return ret
	}
	return *o.CrowdNo
}

// GetCrowdNoOk returns a tuple with the CrowdNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCampaignCashCreateResponseModel) GetCrowdNoOk() (*string, bool) {
	if o == nil || IsNil(o.CrowdNo) {
		return nil, false
	}
	return o.CrowdNo, true
}

// HasCrowdNo returns a boolean if a field has been set.
func (o *AlipayMarketingCampaignCashCreateResponseModel) HasCrowdNo() bool {
	if o != nil && !IsNil(o.CrowdNo) {
		return true
	}

	return false
}

// SetCrowdNo gets a reference to the given string and assigns it to the CrowdNo field.
func (o *AlipayMarketingCampaignCashCreateResponseModel) SetCrowdNo(v string) {
	o.CrowdNo = &v
}

// GetOriginCrowdNo returns the OriginCrowdNo field value if set, zero value otherwise.
func (o *AlipayMarketingCampaignCashCreateResponseModel) GetOriginCrowdNo() string {
	if o == nil || IsNil(o.OriginCrowdNo) {
		var ret string
		return ret
	}
	return *o.OriginCrowdNo
}

// GetOriginCrowdNoOk returns a tuple with the OriginCrowdNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCampaignCashCreateResponseModel) GetOriginCrowdNoOk() (*string, bool) {
	if o == nil || IsNil(o.OriginCrowdNo) {
		return nil, false
	}
	return o.OriginCrowdNo, true
}

// HasOriginCrowdNo returns a boolean if a field has been set.
func (o *AlipayMarketingCampaignCashCreateResponseModel) HasOriginCrowdNo() bool {
	if o != nil && !IsNil(o.OriginCrowdNo) {
		return true
	}

	return false
}

// SetOriginCrowdNo gets a reference to the given string and assigns it to the OriginCrowdNo field.
func (o *AlipayMarketingCampaignCashCreateResponseModel) SetOriginCrowdNo(v string) {
	o.OriginCrowdNo = &v
}

// GetPayUrl returns the PayUrl field value if set, zero value otherwise.
func (o *AlipayMarketingCampaignCashCreateResponseModel) GetPayUrl() string {
	if o == nil || IsNil(o.PayUrl) {
		var ret string
		return ret
	}
	return *o.PayUrl
}

// GetPayUrlOk returns a tuple with the PayUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCampaignCashCreateResponseModel) GetPayUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PayUrl) {
		return nil, false
	}
	return o.PayUrl, true
}

// HasPayUrl returns a boolean if a field has been set.
func (o *AlipayMarketingCampaignCashCreateResponseModel) HasPayUrl() bool {
	if o != nil && !IsNil(o.PayUrl) {
		return true
	}

	return false
}

// SetPayUrl gets a reference to the given string and assigns it to the PayUrl field.
func (o *AlipayMarketingCampaignCashCreateResponseModel) SetPayUrl(v string) {
	o.PayUrl = &v
}

func (o AlipayMarketingCampaignCashCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingCampaignCashCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CrowdNo) {
		toSerialize["crowd_no"] = o.CrowdNo
	}
	if !IsNil(o.OriginCrowdNo) {
		toSerialize["origin_crowd_no"] = o.OriginCrowdNo
	}
	if !IsNil(o.PayUrl) {
		toSerialize["pay_url"] = o.PayUrl
	}
	return toSerialize, nil
}

type NullableAlipayMarketingCampaignCashCreateResponseModel struct {
	value *AlipayMarketingCampaignCashCreateResponseModel
	isSet bool
}

func (v NullableAlipayMarketingCampaignCashCreateResponseModel) Get() *AlipayMarketingCampaignCashCreateResponseModel {
	return v.value
}

func (v *NullableAlipayMarketingCampaignCashCreateResponseModel) Set(val *AlipayMarketingCampaignCashCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingCampaignCashCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingCampaignCashCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingCampaignCashCreateResponseModel(val *AlipayMarketingCampaignCashCreateResponseModel) *NullableAlipayMarketingCampaignCashCreateResponseModel {
	return &NullableAlipayMarketingCampaignCashCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingCampaignCashCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingCampaignCashCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
