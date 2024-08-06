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

// checks if the StandardServiceBaseInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardServiceBaseInfo{}

// StandardServiceBaseInfo struct for StandardServiceBaseInfo
type StandardServiceBaseInfo struct {
	// 服务状态
	BizStatus *string `json:"biz_status,omitempty"`
	// 类目id
	CategoryId *string `json:"category_id,omitempty"`
	// 服务code
	ServiceCode *string `json:"service_code,omitempty"`
	// 服务名称
	ServiceName *string `json:"service_name,omitempty"`
}

// NewStandardServiceBaseInfo instantiates a new StandardServiceBaseInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardServiceBaseInfo() *StandardServiceBaseInfo {
	this := StandardServiceBaseInfo{}
	return &this
}

// NewStandardServiceBaseInfoWithDefaults instantiates a new StandardServiceBaseInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardServiceBaseInfoWithDefaults() *StandardServiceBaseInfo {
	this := StandardServiceBaseInfo{}
	return &this
}

// GetBizStatus returns the BizStatus field value if set, zero value otherwise.
func (o *StandardServiceBaseInfo) GetBizStatus() string {
	if o == nil || IsNil(o.BizStatus) {
		var ret string
		return ret
	}
	return *o.BizStatus
}

// GetBizStatusOk returns a tuple with the BizStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardServiceBaseInfo) GetBizStatusOk() (*string, bool) {
	if o == nil || IsNil(o.BizStatus) {
		return nil, false
	}
	return o.BizStatus, true
}

// HasBizStatus returns a boolean if a field has been set.
func (o *StandardServiceBaseInfo) HasBizStatus() bool {
	if o != nil && !IsNil(o.BizStatus) {
		return true
	}

	return false
}

// SetBizStatus gets a reference to the given string and assigns it to the BizStatus field.
func (o *StandardServiceBaseInfo) SetBizStatus(v string) {
	o.BizStatus = &v
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *StandardServiceBaseInfo) GetCategoryId() string {
	if o == nil || IsNil(o.CategoryId) {
		var ret string
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardServiceBaseInfo) GetCategoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryId) {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *StandardServiceBaseInfo) HasCategoryId() bool {
	if o != nil && !IsNil(o.CategoryId) {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given string and assigns it to the CategoryId field.
func (o *StandardServiceBaseInfo) SetCategoryId(v string) {
	o.CategoryId = &v
}

// GetServiceCode returns the ServiceCode field value if set, zero value otherwise.
func (o *StandardServiceBaseInfo) GetServiceCode() string {
	if o == nil || IsNil(o.ServiceCode) {
		var ret string
		return ret
	}
	return *o.ServiceCode
}

// GetServiceCodeOk returns a tuple with the ServiceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardServiceBaseInfo) GetServiceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceCode) {
		return nil, false
	}
	return o.ServiceCode, true
}

// HasServiceCode returns a boolean if a field has been set.
func (o *StandardServiceBaseInfo) HasServiceCode() bool {
	if o != nil && !IsNil(o.ServiceCode) {
		return true
	}

	return false
}

// SetServiceCode gets a reference to the given string and assigns it to the ServiceCode field.
func (o *StandardServiceBaseInfo) SetServiceCode(v string) {
	o.ServiceCode = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *StandardServiceBaseInfo) GetServiceName() string {
	if o == nil || IsNil(o.ServiceName) {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardServiceBaseInfo) GetServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceName) {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *StandardServiceBaseInfo) HasServiceName() bool {
	if o != nil && !IsNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *StandardServiceBaseInfo) SetServiceName(v string) {
	o.ServiceName = &v
}

func (o StandardServiceBaseInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandardServiceBaseInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BizStatus) {
		toSerialize["biz_status"] = o.BizStatus
	}
	if !IsNil(o.CategoryId) {
		toSerialize["category_id"] = o.CategoryId
	}
	if !IsNil(o.ServiceCode) {
		toSerialize["service_code"] = o.ServiceCode
	}
	if !IsNil(o.ServiceName) {
		toSerialize["service_name"] = o.ServiceName
	}
	return toSerialize, nil
}

type NullableStandardServiceBaseInfo struct {
	value *StandardServiceBaseInfo
	isSet bool
}

func (v NullableStandardServiceBaseInfo) Get() *StandardServiceBaseInfo {
	return v.value
}

func (v *NullableStandardServiceBaseInfo) Set(val *StandardServiceBaseInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardServiceBaseInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardServiceBaseInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardServiceBaseInfo(val *StandardServiceBaseInfo) *NullableStandardServiceBaseInfo {
	return &NullableStandardServiceBaseInfo{value: val, isSet: true}
}

func (v NullableStandardServiceBaseInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandardServiceBaseInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


