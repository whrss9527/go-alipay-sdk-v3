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

// checks if the MiniAppCategoryInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MiniAppCategoryInfo{}

// MiniAppCategoryInfo struct for MiniAppCategoryInfo
type MiniAppCategoryInfo struct {
	// 一级类目id
	FirstCategoryId *string `json:"first_category_id,omitempty"`
	// 一级类目名称
	FirstCategoryName *string `json:"first_category_name,omitempty"`
	// 二级类目id
	SecondCategoryId *string `json:"second_category_id,omitempty"`
	// 二级类目名称
	SecondCategoryName *string `json:"second_category_name,omitempty"`
	// 三级类目id，可空
	ThirdCategoryId *string `json:"third_category_id,omitempty"`
	// 三级类目名称，可空
	ThirdCategoryName *string `json:"third_category_name,omitempty"`
}

// NewMiniAppCategoryInfo instantiates a new MiniAppCategoryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMiniAppCategoryInfo() *MiniAppCategoryInfo {
	this := MiniAppCategoryInfo{}
	return &this
}

// NewMiniAppCategoryInfoWithDefaults instantiates a new MiniAppCategoryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMiniAppCategoryInfoWithDefaults() *MiniAppCategoryInfo {
	this := MiniAppCategoryInfo{}
	return &this
}

// GetFirstCategoryId returns the FirstCategoryId field value if set, zero value otherwise.
func (o *MiniAppCategoryInfo) GetFirstCategoryId() string {
	if o == nil || IsNil(o.FirstCategoryId) {
		var ret string
		return ret
	}
	return *o.FirstCategoryId
}

// GetFirstCategoryIdOk returns a tuple with the FirstCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppCategoryInfo) GetFirstCategoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.FirstCategoryId) {
		return nil, false
	}
	return o.FirstCategoryId, true
}

// HasFirstCategoryId returns a boolean if a field has been set.
func (o *MiniAppCategoryInfo) HasFirstCategoryId() bool {
	if o != nil && !IsNil(o.FirstCategoryId) {
		return true
	}

	return false
}

// SetFirstCategoryId gets a reference to the given string and assigns it to the FirstCategoryId field.
func (o *MiniAppCategoryInfo) SetFirstCategoryId(v string) {
	o.FirstCategoryId = &v
}

// GetFirstCategoryName returns the FirstCategoryName field value if set, zero value otherwise.
func (o *MiniAppCategoryInfo) GetFirstCategoryName() string {
	if o == nil || IsNil(o.FirstCategoryName) {
		var ret string
		return ret
	}
	return *o.FirstCategoryName
}

// GetFirstCategoryNameOk returns a tuple with the FirstCategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppCategoryInfo) GetFirstCategoryNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstCategoryName) {
		return nil, false
	}
	return o.FirstCategoryName, true
}

// HasFirstCategoryName returns a boolean if a field has been set.
func (o *MiniAppCategoryInfo) HasFirstCategoryName() bool {
	if o != nil && !IsNil(o.FirstCategoryName) {
		return true
	}

	return false
}

// SetFirstCategoryName gets a reference to the given string and assigns it to the FirstCategoryName field.
func (o *MiniAppCategoryInfo) SetFirstCategoryName(v string) {
	o.FirstCategoryName = &v
}

// GetSecondCategoryId returns the SecondCategoryId field value if set, zero value otherwise.
func (o *MiniAppCategoryInfo) GetSecondCategoryId() string {
	if o == nil || IsNil(o.SecondCategoryId) {
		var ret string
		return ret
	}
	return *o.SecondCategoryId
}

// GetSecondCategoryIdOk returns a tuple with the SecondCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppCategoryInfo) GetSecondCategoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.SecondCategoryId) {
		return nil, false
	}
	return o.SecondCategoryId, true
}

// HasSecondCategoryId returns a boolean if a field has been set.
func (o *MiniAppCategoryInfo) HasSecondCategoryId() bool {
	if o != nil && !IsNil(o.SecondCategoryId) {
		return true
	}

	return false
}

// SetSecondCategoryId gets a reference to the given string and assigns it to the SecondCategoryId field.
func (o *MiniAppCategoryInfo) SetSecondCategoryId(v string) {
	o.SecondCategoryId = &v
}

// GetSecondCategoryName returns the SecondCategoryName field value if set, zero value otherwise.
func (o *MiniAppCategoryInfo) GetSecondCategoryName() string {
	if o == nil || IsNil(o.SecondCategoryName) {
		var ret string
		return ret
	}
	return *o.SecondCategoryName
}

// GetSecondCategoryNameOk returns a tuple with the SecondCategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppCategoryInfo) GetSecondCategoryNameOk() (*string, bool) {
	if o == nil || IsNil(o.SecondCategoryName) {
		return nil, false
	}
	return o.SecondCategoryName, true
}

// HasSecondCategoryName returns a boolean if a field has been set.
func (o *MiniAppCategoryInfo) HasSecondCategoryName() bool {
	if o != nil && !IsNil(o.SecondCategoryName) {
		return true
	}

	return false
}

// SetSecondCategoryName gets a reference to the given string and assigns it to the SecondCategoryName field.
func (o *MiniAppCategoryInfo) SetSecondCategoryName(v string) {
	o.SecondCategoryName = &v
}

// GetThirdCategoryId returns the ThirdCategoryId field value if set, zero value otherwise.
func (o *MiniAppCategoryInfo) GetThirdCategoryId() string {
	if o == nil || IsNil(o.ThirdCategoryId) {
		var ret string
		return ret
	}
	return *o.ThirdCategoryId
}

// GetThirdCategoryIdOk returns a tuple with the ThirdCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppCategoryInfo) GetThirdCategoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.ThirdCategoryId) {
		return nil, false
	}
	return o.ThirdCategoryId, true
}

// HasThirdCategoryId returns a boolean if a field has been set.
func (o *MiniAppCategoryInfo) HasThirdCategoryId() bool {
	if o != nil && !IsNil(o.ThirdCategoryId) {
		return true
	}

	return false
}

// SetThirdCategoryId gets a reference to the given string and assigns it to the ThirdCategoryId field.
func (o *MiniAppCategoryInfo) SetThirdCategoryId(v string) {
	o.ThirdCategoryId = &v
}

// GetThirdCategoryName returns the ThirdCategoryName field value if set, zero value otherwise.
func (o *MiniAppCategoryInfo) GetThirdCategoryName() string {
	if o == nil || IsNil(o.ThirdCategoryName) {
		var ret string
		return ret
	}
	return *o.ThirdCategoryName
}

// GetThirdCategoryNameOk returns a tuple with the ThirdCategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppCategoryInfo) GetThirdCategoryNameOk() (*string, bool) {
	if o == nil || IsNil(o.ThirdCategoryName) {
		return nil, false
	}
	return o.ThirdCategoryName, true
}

// HasThirdCategoryName returns a boolean if a field has been set.
func (o *MiniAppCategoryInfo) HasThirdCategoryName() bool {
	if o != nil && !IsNil(o.ThirdCategoryName) {
		return true
	}

	return false
}

// SetThirdCategoryName gets a reference to the given string and assigns it to the ThirdCategoryName field.
func (o *MiniAppCategoryInfo) SetThirdCategoryName(v string) {
	o.ThirdCategoryName = &v
}

func (o MiniAppCategoryInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MiniAppCategoryInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirstCategoryId) {
		toSerialize["first_category_id"] = o.FirstCategoryId
	}
	if !IsNil(o.FirstCategoryName) {
		toSerialize["first_category_name"] = o.FirstCategoryName
	}
	if !IsNil(o.SecondCategoryId) {
		toSerialize["second_category_id"] = o.SecondCategoryId
	}
	if !IsNil(o.SecondCategoryName) {
		toSerialize["second_category_name"] = o.SecondCategoryName
	}
	if !IsNil(o.ThirdCategoryId) {
		toSerialize["third_category_id"] = o.ThirdCategoryId
	}
	if !IsNil(o.ThirdCategoryName) {
		toSerialize["third_category_name"] = o.ThirdCategoryName
	}
	return toSerialize, nil
}

type NullableMiniAppCategoryInfo struct {
	value *MiniAppCategoryInfo
	isSet bool
}

func (v NullableMiniAppCategoryInfo) Get() *MiniAppCategoryInfo {
	return v.value
}

func (v *NullableMiniAppCategoryInfo) Set(val *MiniAppCategoryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMiniAppCategoryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMiniAppCategoryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMiniAppCategoryInfo(val *MiniAppCategoryInfo) *NullableMiniAppCategoryInfo {
	return &NullableMiniAppCategoryInfo{value: val, isSet: true}
}

func (v NullableMiniAppCategoryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMiniAppCategoryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


