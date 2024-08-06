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

// checks if the OrderDataDistributeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDataDistributeInfo{}

// OrderDataDistributeInfo struct for OrderDataDistributeInfo
type OrderDataDistributeInfo struct {
	// 未分发到场景的具体原因。开发者可根据具体原因定位解决问题后发起重试。
	NotDistributeReason *string `json:"not_distribute_reason,omitempty"`
	// 分发场景code。
	SceneCode *string `json:"scene_code,omitempty"`
	// 分发场景名，对应scene_code。
	SceneName *string `json:"scene_name,omitempty"`
}

// NewOrderDataDistributeInfo instantiates a new OrderDataDistributeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDataDistributeInfo() *OrderDataDistributeInfo {
	this := OrderDataDistributeInfo{}
	return &this
}

// NewOrderDataDistributeInfoWithDefaults instantiates a new OrderDataDistributeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDataDistributeInfoWithDefaults() *OrderDataDistributeInfo {
	this := OrderDataDistributeInfo{}
	return &this
}

// GetNotDistributeReason returns the NotDistributeReason field value if set, zero value otherwise.
func (o *OrderDataDistributeInfo) GetNotDistributeReason() string {
	if o == nil || IsNil(o.NotDistributeReason) {
		var ret string
		return ret
	}
	return *o.NotDistributeReason
}

// GetNotDistributeReasonOk returns a tuple with the NotDistributeReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataDistributeInfo) GetNotDistributeReasonOk() (*string, bool) {
	if o == nil || IsNil(o.NotDistributeReason) {
		return nil, false
	}
	return o.NotDistributeReason, true
}

// HasNotDistributeReason returns a boolean if a field has been set.
func (o *OrderDataDistributeInfo) HasNotDistributeReason() bool {
	if o != nil && !IsNil(o.NotDistributeReason) {
		return true
	}

	return false
}

// SetNotDistributeReason gets a reference to the given string and assigns it to the NotDistributeReason field.
func (o *OrderDataDistributeInfo) SetNotDistributeReason(v string) {
	o.NotDistributeReason = &v
}

// GetSceneCode returns the SceneCode field value if set, zero value otherwise.
func (o *OrderDataDistributeInfo) GetSceneCode() string {
	if o == nil || IsNil(o.SceneCode) {
		var ret string
		return ret
	}
	return *o.SceneCode
}

// GetSceneCodeOk returns a tuple with the SceneCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataDistributeInfo) GetSceneCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SceneCode) {
		return nil, false
	}
	return o.SceneCode, true
}

// HasSceneCode returns a boolean if a field has been set.
func (o *OrderDataDistributeInfo) HasSceneCode() bool {
	if o != nil && !IsNil(o.SceneCode) {
		return true
	}

	return false
}

// SetSceneCode gets a reference to the given string and assigns it to the SceneCode field.
func (o *OrderDataDistributeInfo) SetSceneCode(v string) {
	o.SceneCode = &v
}

// GetSceneName returns the SceneName field value if set, zero value otherwise.
func (o *OrderDataDistributeInfo) GetSceneName() string {
	if o == nil || IsNil(o.SceneName) {
		var ret string
		return ret
	}
	return *o.SceneName
}

// GetSceneNameOk returns a tuple with the SceneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataDistributeInfo) GetSceneNameOk() (*string, bool) {
	if o == nil || IsNil(o.SceneName) {
		return nil, false
	}
	return o.SceneName, true
}

// HasSceneName returns a boolean if a field has been set.
func (o *OrderDataDistributeInfo) HasSceneName() bool {
	if o != nil && !IsNil(o.SceneName) {
		return true
	}

	return false
}

// SetSceneName gets a reference to the given string and assigns it to the SceneName field.
func (o *OrderDataDistributeInfo) SetSceneName(v string) {
	o.SceneName = &v
}

func (o OrderDataDistributeInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDataDistributeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NotDistributeReason) {
		toSerialize["not_distribute_reason"] = o.NotDistributeReason
	}
	if !IsNil(o.SceneCode) {
		toSerialize["scene_code"] = o.SceneCode
	}
	if !IsNil(o.SceneName) {
		toSerialize["scene_name"] = o.SceneName
	}
	return toSerialize, nil
}

type NullableOrderDataDistributeInfo struct {
	value *OrderDataDistributeInfo
	isSet bool
}

func (v NullableOrderDataDistributeInfo) Get() *OrderDataDistributeInfo {
	return v.value
}

func (v *NullableOrderDataDistributeInfo) Set(val *OrderDataDistributeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDataDistributeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDataDistributeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDataDistributeInfo(val *OrderDataDistributeInfo) *NullableOrderDataDistributeInfo {
	return &NullableOrderDataDistributeInfo{value: val, isSet: true}
}

func (v NullableOrderDataDistributeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDataDistributeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
