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

// checks if the ContactFollower type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactFollower{}

// ContactFollower struct for ContactFollower
type ContactFollower struct {
	// 支付宝头像
	Avatar *string `json:"avatar,omitempty"`
	// 默认头像
	DefaultAvatar *string `json:"default_avatar,omitempty"`
	// false
	EachRecordFlag *string `json:"each_record_flag,omitempty"`
	// 用户id
	UserId *string `json:"user_id,omitempty"`
}

// NewContactFollower instantiates a new ContactFollower object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactFollower() *ContactFollower {
	this := ContactFollower{}
	return &this
}

// NewContactFollowerWithDefaults instantiates a new ContactFollower object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactFollowerWithDefaults() *ContactFollower {
	this := ContactFollower{}
	return &this
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *ContactFollower) GetAvatar() string {
	if o == nil || IsNil(o.Avatar) {
		var ret string
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactFollower) GetAvatarOk() (*string, bool) {
	if o == nil || IsNil(o.Avatar) {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *ContactFollower) HasAvatar() bool {
	if o != nil && !IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given string and assigns it to the Avatar field.
func (o *ContactFollower) SetAvatar(v string) {
	o.Avatar = &v
}

// GetDefaultAvatar returns the DefaultAvatar field value if set, zero value otherwise.
func (o *ContactFollower) GetDefaultAvatar() string {
	if o == nil || IsNil(o.DefaultAvatar) {
		var ret string
		return ret
	}
	return *o.DefaultAvatar
}

// GetDefaultAvatarOk returns a tuple with the DefaultAvatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactFollower) GetDefaultAvatarOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultAvatar) {
		return nil, false
	}
	return o.DefaultAvatar, true
}

// HasDefaultAvatar returns a boolean if a field has been set.
func (o *ContactFollower) HasDefaultAvatar() bool {
	if o != nil && !IsNil(o.DefaultAvatar) {
		return true
	}

	return false
}

// SetDefaultAvatar gets a reference to the given string and assigns it to the DefaultAvatar field.
func (o *ContactFollower) SetDefaultAvatar(v string) {
	o.DefaultAvatar = &v
}

// GetEachRecordFlag returns the EachRecordFlag field value if set, zero value otherwise.
func (o *ContactFollower) GetEachRecordFlag() string {
	if o == nil || IsNil(o.EachRecordFlag) {
		var ret string
		return ret
	}
	return *o.EachRecordFlag
}

// GetEachRecordFlagOk returns a tuple with the EachRecordFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactFollower) GetEachRecordFlagOk() (*string, bool) {
	if o == nil || IsNil(o.EachRecordFlag) {
		return nil, false
	}
	return o.EachRecordFlag, true
}

// HasEachRecordFlag returns a boolean if a field has been set.
func (o *ContactFollower) HasEachRecordFlag() bool {
	if o != nil && !IsNil(o.EachRecordFlag) {
		return true
	}

	return false
}

// SetEachRecordFlag gets a reference to the given string and assigns it to the EachRecordFlag field.
func (o *ContactFollower) SetEachRecordFlag(v string) {
	o.EachRecordFlag = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ContactFollower) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactFollower) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ContactFollower) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *ContactFollower) SetUserId(v string) {
	o.UserId = &v
}

func (o ContactFollower) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactFollower) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Avatar) {
		toSerialize["avatar"] = o.Avatar
	}
	if !IsNil(o.DefaultAvatar) {
		toSerialize["default_avatar"] = o.DefaultAvatar
	}
	if !IsNil(o.EachRecordFlag) {
		toSerialize["each_record_flag"] = o.EachRecordFlag
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableContactFollower struct {
	value *ContactFollower
	isSet bool
}

func (v NullableContactFollower) Get() *ContactFollower {
	return v.value
}

func (v *NullableContactFollower) Set(val *ContactFollower) {
	v.value = val
	v.isSet = true
}

func (v NullableContactFollower) IsSet() bool {
	return v.isSet
}

func (v *NullableContactFollower) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactFollower(val *ContactFollower) *NullableContactFollower {
	return &NullableContactFollower{value: val, isSet: true}
}

func (v NullableContactFollower) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactFollower) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
