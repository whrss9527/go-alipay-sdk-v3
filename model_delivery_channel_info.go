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

// checks if the DeliveryChannelInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryChannelInfo{}

// DeliveryChannelInfo struct for DeliveryChannelInfo
type DeliveryChannelInfo struct {
	// 展位码:boothCode表达的是具体渠道所属的展位码。例如：支付结果页PAY_RESULT
	BoothCode *string `json:"booth_code,omitempty"`
	// 可投放的渠道标识.  说明： 针对不同的boot_code，channel值各不相同.接口使用者可以认为channel是不同booth_code展位下的渠道的唯一标识。  例如： booth_code为PAYMENT_RESULT。 channel为某个商户的pid。
	Channel *string `json:"channel,omitempty"`
	// 渠道的名称。  说明： booth_code：PAYMENT_RESULT，channel_name为商户名称。
	ChannelName *string `json:"channel_name,omitempty"`
}

// NewDeliveryChannelInfo instantiates a new DeliveryChannelInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryChannelInfo() *DeliveryChannelInfo {
	this := DeliveryChannelInfo{}
	return &this
}

// NewDeliveryChannelInfoWithDefaults instantiates a new DeliveryChannelInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryChannelInfoWithDefaults() *DeliveryChannelInfo {
	this := DeliveryChannelInfo{}
	return &this
}

// GetBoothCode returns the BoothCode field value if set, zero value otherwise.
func (o *DeliveryChannelInfo) GetBoothCode() string {
	if o == nil || IsNil(o.BoothCode) {
		var ret string
		return ret
	}
	return *o.BoothCode
}

// GetBoothCodeOk returns a tuple with the BoothCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryChannelInfo) GetBoothCodeOk() (*string, bool) {
	if o == nil || IsNil(o.BoothCode) {
		return nil, false
	}
	return o.BoothCode, true
}

// HasBoothCode returns a boolean if a field has been set.
func (o *DeliveryChannelInfo) HasBoothCode() bool {
	if o != nil && !IsNil(o.BoothCode) {
		return true
	}

	return false
}

// SetBoothCode gets a reference to the given string and assigns it to the BoothCode field.
func (o *DeliveryChannelInfo) SetBoothCode(v string) {
	o.BoothCode = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *DeliveryChannelInfo) GetChannel() string {
	if o == nil || IsNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryChannelInfo) GetChannelOk() (*string, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *DeliveryChannelInfo) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *DeliveryChannelInfo) SetChannel(v string) {
	o.Channel = &v
}

// GetChannelName returns the ChannelName field value if set, zero value otherwise.
func (o *DeliveryChannelInfo) GetChannelName() string {
	if o == nil || IsNil(o.ChannelName) {
		var ret string
		return ret
	}
	return *o.ChannelName
}

// GetChannelNameOk returns a tuple with the ChannelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryChannelInfo) GetChannelNameOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelName) {
		return nil, false
	}
	return o.ChannelName, true
}

// HasChannelName returns a boolean if a field has been set.
func (o *DeliveryChannelInfo) HasChannelName() bool {
	if o != nil && !IsNil(o.ChannelName) {
		return true
	}

	return false
}

// SetChannelName gets a reference to the given string and assigns it to the ChannelName field.
func (o *DeliveryChannelInfo) SetChannelName(v string) {
	o.ChannelName = &v
}

func (o DeliveryChannelInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryChannelInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BoothCode) {
		toSerialize["booth_code"] = o.BoothCode
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.ChannelName) {
		toSerialize["channel_name"] = o.ChannelName
	}
	return toSerialize, nil
}

type NullableDeliveryChannelInfo struct {
	value *DeliveryChannelInfo
	isSet bool
}

func (v NullableDeliveryChannelInfo) Get() *DeliveryChannelInfo {
	return v.value
}

func (v *NullableDeliveryChannelInfo) Set(val *DeliveryChannelInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryChannelInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryChannelInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryChannelInfo(val *DeliveryChannelInfo) *NullableDeliveryChannelInfo {
	return &NullableDeliveryChannelInfo{value: val, isSet: true}
}

func (v NullableDeliveryChannelInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryChannelInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

