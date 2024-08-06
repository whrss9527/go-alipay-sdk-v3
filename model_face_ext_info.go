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

// checks if the FaceExtInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FaceExtInfo{}

// FaceExtInfo struct for FaceExtInfo
type FaceExtInfo struct {
	// 年龄区间判断的上限，闭区间
	MaxAge *string `json:"max_age,omitempty"`
	// 年龄区间判断的下限，闭区间
	MinAge *string `json:"min_age,omitempty"`
	// query_type不填, 返回uid  query_type=1, 返回手机号  query_type=2, 返回图片
	QueryType *string `json:"query_type,omitempty"`
}

// NewFaceExtInfo instantiates a new FaceExtInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFaceExtInfo() *FaceExtInfo {
	this := FaceExtInfo{}
	return &this
}

// NewFaceExtInfoWithDefaults instantiates a new FaceExtInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFaceExtInfoWithDefaults() *FaceExtInfo {
	this := FaceExtInfo{}
	return &this
}

// GetMaxAge returns the MaxAge field value if set, zero value otherwise.
func (o *FaceExtInfo) GetMaxAge() string {
	if o == nil || IsNil(o.MaxAge) {
		var ret string
		return ret
	}
	return *o.MaxAge
}

// GetMaxAgeOk returns a tuple with the MaxAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaceExtInfo) GetMaxAgeOk() (*string, bool) {
	if o == nil || IsNil(o.MaxAge) {
		return nil, false
	}
	return o.MaxAge, true
}

// HasMaxAge returns a boolean if a field has been set.
func (o *FaceExtInfo) HasMaxAge() bool {
	if o != nil && !IsNil(o.MaxAge) {
		return true
	}

	return false
}

// SetMaxAge gets a reference to the given string and assigns it to the MaxAge field.
func (o *FaceExtInfo) SetMaxAge(v string) {
	o.MaxAge = &v
}

// GetMinAge returns the MinAge field value if set, zero value otherwise.
func (o *FaceExtInfo) GetMinAge() string {
	if o == nil || IsNil(o.MinAge) {
		var ret string
		return ret
	}
	return *o.MinAge
}

// GetMinAgeOk returns a tuple with the MinAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaceExtInfo) GetMinAgeOk() (*string, bool) {
	if o == nil || IsNil(o.MinAge) {
		return nil, false
	}
	return o.MinAge, true
}

// HasMinAge returns a boolean if a field has been set.
func (o *FaceExtInfo) HasMinAge() bool {
	if o != nil && !IsNil(o.MinAge) {
		return true
	}

	return false
}

// SetMinAge gets a reference to the given string and assigns it to the MinAge field.
func (o *FaceExtInfo) SetMinAge(v string) {
	o.MinAge = &v
}

// GetQueryType returns the QueryType field value if set, zero value otherwise.
func (o *FaceExtInfo) GetQueryType() string {
	if o == nil || IsNil(o.QueryType) {
		var ret string
		return ret
	}
	return *o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaceExtInfo) GetQueryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.QueryType) {
		return nil, false
	}
	return o.QueryType, true
}

// HasQueryType returns a boolean if a field has been set.
func (o *FaceExtInfo) HasQueryType() bool {
	if o != nil && !IsNil(o.QueryType) {
		return true
	}

	return false
}

// SetQueryType gets a reference to the given string and assigns it to the QueryType field.
func (o *FaceExtInfo) SetQueryType(v string) {
	o.QueryType = &v
}

func (o FaceExtInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FaceExtInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxAge) {
		toSerialize["max_age"] = o.MaxAge
	}
	if !IsNil(o.MinAge) {
		toSerialize["min_age"] = o.MinAge
	}
	if !IsNil(o.QueryType) {
		toSerialize["query_type"] = o.QueryType
	}
	return toSerialize, nil
}

type NullableFaceExtInfo struct {
	value *FaceExtInfo
	isSet bool
}

func (v NullableFaceExtInfo) Get() *FaceExtInfo {
	return v.value
}

func (v *NullableFaceExtInfo) Set(val *FaceExtInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFaceExtInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFaceExtInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFaceExtInfo(val *FaceExtInfo) *NullableFaceExtInfo {
	return &NullableFaceExtInfo{value: val, isSet: true}
}

func (v NullableFaceExtInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFaceExtInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

