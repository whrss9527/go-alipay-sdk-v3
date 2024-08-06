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

// checks if the CloudResumeHeadPic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudResumeHeadPic{}

// CloudResumeHeadPic struct for CloudResumeHeadPic
type CloudResumeHeadPic struct {
	// 头像连接url
	PicUrl *string `json:"pic_url,omitempty"`
}

// NewCloudResumeHeadPic instantiates a new CloudResumeHeadPic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudResumeHeadPic() *CloudResumeHeadPic {
	this := CloudResumeHeadPic{}
	return &this
}

// NewCloudResumeHeadPicWithDefaults instantiates a new CloudResumeHeadPic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudResumeHeadPicWithDefaults() *CloudResumeHeadPic {
	this := CloudResumeHeadPic{}
	return &this
}

// GetPicUrl returns the PicUrl field value if set, zero value otherwise.
func (o *CloudResumeHeadPic) GetPicUrl() string {
	if o == nil || IsNil(o.PicUrl) {
		var ret string
		return ret
	}
	return *o.PicUrl
}

// GetPicUrlOk returns a tuple with the PicUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudResumeHeadPic) GetPicUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PicUrl) {
		return nil, false
	}
	return o.PicUrl, true
}

// HasPicUrl returns a boolean if a field has been set.
func (o *CloudResumeHeadPic) HasPicUrl() bool {
	if o != nil && !IsNil(o.PicUrl) {
		return true
	}

	return false
}

// SetPicUrl gets a reference to the given string and assigns it to the PicUrl field.
func (o *CloudResumeHeadPic) SetPicUrl(v string) {
	o.PicUrl = &v
}

func (o CloudResumeHeadPic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudResumeHeadPic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PicUrl) {
		toSerialize["pic_url"] = o.PicUrl
	}
	return toSerialize, nil
}

type NullableCloudResumeHeadPic struct {
	value *CloudResumeHeadPic
	isSet bool
}

func (v NullableCloudResumeHeadPic) Get() *CloudResumeHeadPic {
	return v.value
}

func (v *NullableCloudResumeHeadPic) Set(val *CloudResumeHeadPic) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudResumeHeadPic) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudResumeHeadPic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudResumeHeadPic(val *CloudResumeHeadPic) *NullableCloudResumeHeadPic {
	return &NullableCloudResumeHeadPic{value: val, isSet: true}
}

func (v NullableCloudResumeHeadPic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudResumeHeadPic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


