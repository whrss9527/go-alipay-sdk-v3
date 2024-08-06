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

// checks if the DeliveryConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryConfig{}

// DeliveryConfig struct for DeliveryConfig
type DeliveryConfig struct {
	// 活动id
	ActivityId *string `json:"activity_id,omitempty"`
	// 展位码:boothCode表达的是具体渠道所属的展位码。例如：支付结果页PAYMENT_RESULT
	BoothCode *string `json:"booth_code,omitempty"`
	// 可投放的渠道标识.
	Channel *string `json:"channel,omitempty"`
}

// NewDeliveryConfig instantiates a new DeliveryConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryConfig() *DeliveryConfig {
	this := DeliveryConfig{}
	return &this
}

// NewDeliveryConfigWithDefaults instantiates a new DeliveryConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryConfigWithDefaults() *DeliveryConfig {
	this := DeliveryConfig{}
	return &this
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise.
func (o *DeliveryConfig) GetActivityId() string {
	if o == nil || IsNil(o.ActivityId) {
		var ret string
		return ret
	}
	return *o.ActivityId
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryConfig) GetActivityIdOk() (*string, bool) {
	if o == nil || IsNil(o.ActivityId) {
		return nil, false
	}
	return o.ActivityId, true
}

// HasActivityId returns a boolean if a field has been set.
func (o *DeliveryConfig) HasActivityId() bool {
	if o != nil && !IsNil(o.ActivityId) {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given string and assigns it to the ActivityId field.
func (o *DeliveryConfig) SetActivityId(v string) {
	o.ActivityId = &v
}

// GetBoothCode returns the BoothCode field value if set, zero value otherwise.
func (o *DeliveryConfig) GetBoothCode() string {
	if o == nil || IsNil(o.BoothCode) {
		var ret string
		return ret
	}
	return *o.BoothCode
}

// GetBoothCodeOk returns a tuple with the BoothCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryConfig) GetBoothCodeOk() (*string, bool) {
	if o == nil || IsNil(o.BoothCode) {
		return nil, false
	}
	return o.BoothCode, true
}

// HasBoothCode returns a boolean if a field has been set.
func (o *DeliveryConfig) HasBoothCode() bool {
	if o != nil && !IsNil(o.BoothCode) {
		return true
	}

	return false
}

// SetBoothCode gets a reference to the given string and assigns it to the BoothCode field.
func (o *DeliveryConfig) SetBoothCode(v string) {
	o.BoothCode = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *DeliveryConfig) GetChannel() string {
	if o == nil || IsNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryConfig) GetChannelOk() (*string, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *DeliveryConfig) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *DeliveryConfig) SetChannel(v string) {
	o.Channel = &v
}

func (o DeliveryConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivityId) {
		toSerialize["activity_id"] = o.ActivityId
	}
	if !IsNil(o.BoothCode) {
		toSerialize["booth_code"] = o.BoothCode
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	return toSerialize, nil
}

type NullableDeliveryConfig struct {
	value *DeliveryConfig
	isSet bool
}

func (v NullableDeliveryConfig) Get() *DeliveryConfig {
	return v.value
}

func (v *NullableDeliveryConfig) Set(val *DeliveryConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryConfig(val *DeliveryConfig) *NullableDeliveryConfig {
	return &NullableDeliveryConfig{value: val, isSet: true}
}

func (v NullableDeliveryConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


