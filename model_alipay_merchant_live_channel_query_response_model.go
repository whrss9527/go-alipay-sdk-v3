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

// checks if the AlipayMerchantLiveChannelQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMerchantLiveChannelQueryResponseModel{}

// AlipayMerchantLiveChannelQueryResponseModel struct for AlipayMerchantLiveChannelQueryResponseModel
type AlipayMerchantLiveChannelQueryResponseModel struct {
	// 渠道内容，包含主播、文章的上游信息。字符串内容为Map，需要转换
	ChannelContent *string `json:"channel_content,omitempty"`
	// 渠道业务标识
	ChannelIdentity *string `json:"channel_identity,omitempty"`
	// 渠道密文
	ChannelSecret *string `json:"channel_secret,omitempty"`
	// 渠道类型
	ChannelType *string `json:"channel_type,omitempty"`
}

// NewAlipayMerchantLiveChannelQueryResponseModel instantiates a new AlipayMerchantLiveChannelQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMerchantLiveChannelQueryResponseModel() *AlipayMerchantLiveChannelQueryResponseModel {
	this := AlipayMerchantLiveChannelQueryResponseModel{}
	return &this
}

// NewAlipayMerchantLiveChannelQueryResponseModelWithDefaults instantiates a new AlipayMerchantLiveChannelQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMerchantLiveChannelQueryResponseModelWithDefaults() *AlipayMerchantLiveChannelQueryResponseModel {
	this := AlipayMerchantLiveChannelQueryResponseModel{}
	return &this
}

// GetChannelContent returns the ChannelContent field value if set, zero value otherwise.
func (o *AlipayMerchantLiveChannelQueryResponseModel) GetChannelContent() string {
	if o == nil || IsNil(o.ChannelContent) {
		var ret string
		return ret
	}
	return *o.ChannelContent
}

// GetChannelContentOk returns a tuple with the ChannelContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantLiveChannelQueryResponseModel) GetChannelContentOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelContent) {
		return nil, false
	}
	return o.ChannelContent, true
}

// HasChannelContent returns a boolean if a field has been set.
func (o *AlipayMerchantLiveChannelQueryResponseModel) HasChannelContent() bool {
	if o != nil && !IsNil(o.ChannelContent) {
		return true
	}

	return false
}

// SetChannelContent gets a reference to the given string and assigns it to the ChannelContent field.
func (o *AlipayMerchantLiveChannelQueryResponseModel) SetChannelContent(v string) {
	o.ChannelContent = &v
}

// GetChannelIdentity returns the ChannelIdentity field value if set, zero value otherwise.
func (o *AlipayMerchantLiveChannelQueryResponseModel) GetChannelIdentity() string {
	if o == nil || IsNil(o.ChannelIdentity) {
		var ret string
		return ret
	}
	return *o.ChannelIdentity
}

// GetChannelIdentityOk returns a tuple with the ChannelIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantLiveChannelQueryResponseModel) GetChannelIdentityOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelIdentity) {
		return nil, false
	}
	return o.ChannelIdentity, true
}

// HasChannelIdentity returns a boolean if a field has been set.
func (o *AlipayMerchantLiveChannelQueryResponseModel) HasChannelIdentity() bool {
	if o != nil && !IsNil(o.ChannelIdentity) {
		return true
	}

	return false
}

// SetChannelIdentity gets a reference to the given string and assigns it to the ChannelIdentity field.
func (o *AlipayMerchantLiveChannelQueryResponseModel) SetChannelIdentity(v string) {
	o.ChannelIdentity = &v
}

// GetChannelSecret returns the ChannelSecret field value if set, zero value otherwise.
func (o *AlipayMerchantLiveChannelQueryResponseModel) GetChannelSecret() string {
	if o == nil || IsNil(o.ChannelSecret) {
		var ret string
		return ret
	}
	return *o.ChannelSecret
}

// GetChannelSecretOk returns a tuple with the ChannelSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantLiveChannelQueryResponseModel) GetChannelSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelSecret) {
		return nil, false
	}
	return o.ChannelSecret, true
}

// HasChannelSecret returns a boolean if a field has been set.
func (o *AlipayMerchantLiveChannelQueryResponseModel) HasChannelSecret() bool {
	if o != nil && !IsNil(o.ChannelSecret) {
		return true
	}

	return false
}

// SetChannelSecret gets a reference to the given string and assigns it to the ChannelSecret field.
func (o *AlipayMerchantLiveChannelQueryResponseModel) SetChannelSecret(v string) {
	o.ChannelSecret = &v
}

// GetChannelType returns the ChannelType field value if set, zero value otherwise.
func (o *AlipayMerchantLiveChannelQueryResponseModel) GetChannelType() string {
	if o == nil || IsNil(o.ChannelType) {
		var ret string
		return ret
	}
	return *o.ChannelType
}

// GetChannelTypeOk returns a tuple with the ChannelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantLiveChannelQueryResponseModel) GetChannelTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelType) {
		return nil, false
	}
	return o.ChannelType, true
}

// HasChannelType returns a boolean if a field has been set.
func (o *AlipayMerchantLiveChannelQueryResponseModel) HasChannelType() bool {
	if o != nil && !IsNil(o.ChannelType) {
		return true
	}

	return false
}

// SetChannelType gets a reference to the given string and assigns it to the ChannelType field.
func (o *AlipayMerchantLiveChannelQueryResponseModel) SetChannelType(v string) {
	o.ChannelType = &v
}

func (o AlipayMerchantLiveChannelQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMerchantLiveChannelQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChannelContent) {
		toSerialize["channel_content"] = o.ChannelContent
	}
	if !IsNil(o.ChannelIdentity) {
		toSerialize["channel_identity"] = o.ChannelIdentity
	}
	if !IsNil(o.ChannelSecret) {
		toSerialize["channel_secret"] = o.ChannelSecret
	}
	if !IsNil(o.ChannelType) {
		toSerialize["channel_type"] = o.ChannelType
	}
	return toSerialize, nil
}

type NullableAlipayMerchantLiveChannelQueryResponseModel struct {
	value *AlipayMerchantLiveChannelQueryResponseModel
	isSet bool
}

func (v NullableAlipayMerchantLiveChannelQueryResponseModel) Get() *AlipayMerchantLiveChannelQueryResponseModel {
	return v.value
}

func (v *NullableAlipayMerchantLiveChannelQueryResponseModel) Set(val *AlipayMerchantLiveChannelQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMerchantLiveChannelQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMerchantLiveChannelQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMerchantLiveChannelQueryResponseModel(val *AlipayMerchantLiveChannelQueryResponseModel) *NullableAlipayMerchantLiveChannelQueryResponseModel {
	return &NullableAlipayMerchantLiveChannelQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMerchantLiveChannelQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMerchantLiveChannelQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


