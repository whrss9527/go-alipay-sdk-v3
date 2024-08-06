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

// checks if the SourceMediaInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceMediaInfo{}

// SourceMediaInfo struct for SourceMediaInfo
type SourceMediaInfo struct {
	// 素材ID，对应“支付宝文件上传接口”获取的file_id
	MediaId *string `json:"media_id,omitempty"`
	// 素材类型。  image：图片  video：视频 cover_static：静态封面
	MediaType *string `json:"media_type,omitempty"`
}

// NewSourceMediaInfo instantiates a new SourceMediaInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceMediaInfo() *SourceMediaInfo {
	this := SourceMediaInfo{}
	return &this
}

// NewSourceMediaInfoWithDefaults instantiates a new SourceMediaInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceMediaInfoWithDefaults() *SourceMediaInfo {
	this := SourceMediaInfo{}
	return &this
}

// GetMediaId returns the MediaId field value if set, zero value otherwise.
func (o *SourceMediaInfo) GetMediaId() string {
	if o == nil || IsNil(o.MediaId) {
		var ret string
		return ret
	}
	return *o.MediaId
}

// GetMediaIdOk returns a tuple with the MediaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceMediaInfo) GetMediaIdOk() (*string, bool) {
	if o == nil || IsNil(o.MediaId) {
		return nil, false
	}
	return o.MediaId, true
}

// HasMediaId returns a boolean if a field has been set.
func (o *SourceMediaInfo) HasMediaId() bool {
	if o != nil && !IsNil(o.MediaId) {
		return true
	}

	return false
}

// SetMediaId gets a reference to the given string and assigns it to the MediaId field.
func (o *SourceMediaInfo) SetMediaId(v string) {
	o.MediaId = &v
}

// GetMediaType returns the MediaType field value if set, zero value otherwise.
func (o *SourceMediaInfo) GetMediaType() string {
	if o == nil || IsNil(o.MediaType) {
		var ret string
		return ret
	}
	return *o.MediaType
}

// GetMediaTypeOk returns a tuple with the MediaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceMediaInfo) GetMediaTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MediaType) {
		return nil, false
	}
	return o.MediaType, true
}

// HasMediaType returns a boolean if a field has been set.
func (o *SourceMediaInfo) HasMediaType() bool {
	if o != nil && !IsNil(o.MediaType) {
		return true
	}

	return false
}

// SetMediaType gets a reference to the given string and assigns it to the MediaType field.
func (o *SourceMediaInfo) SetMediaType(v string) {
	o.MediaType = &v
}

func (o SourceMediaInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceMediaInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MediaId) {
		toSerialize["media_id"] = o.MediaId
	}
	if !IsNil(o.MediaType) {
		toSerialize["media_type"] = o.MediaType
	}
	return toSerialize, nil
}

type NullableSourceMediaInfo struct {
	value *SourceMediaInfo
	isSet bool
}

func (v NullableSourceMediaInfo) Get() *SourceMediaInfo {
	return v.value
}

func (v *NullableSourceMediaInfo) Set(val *SourceMediaInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceMediaInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceMediaInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceMediaInfo(val *SourceMediaInfo) *NullableSourceMediaInfo {
	return &NullableSourceMediaInfo{value: val, isSet: true}
}

func (v NullableSourceMediaInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceMediaInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

