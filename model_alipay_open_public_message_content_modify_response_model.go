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

// checks if the AlipayOpenPublicMessageContentModifyResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicMessageContentModifyResponseModel{}

// AlipayOpenPublicMessageContentModifyResponseModel struct for AlipayOpenPublicMessageContentModifyResponseModel
type AlipayOpenPublicMessageContentModifyResponseModel struct {
	// 内容id
	ContentId *string `json:"content_id,omitempty"`
	// 内容详情内链url
	ContentUrl *string `json:"content_url,omitempty"`
}

// NewAlipayOpenPublicMessageContentModifyResponseModel instantiates a new AlipayOpenPublicMessageContentModifyResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicMessageContentModifyResponseModel() *AlipayOpenPublicMessageContentModifyResponseModel {
	this := AlipayOpenPublicMessageContentModifyResponseModel{}
	return &this
}

// NewAlipayOpenPublicMessageContentModifyResponseModelWithDefaults instantiates a new AlipayOpenPublicMessageContentModifyResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicMessageContentModifyResponseModelWithDefaults() *AlipayOpenPublicMessageContentModifyResponseModel {
	this := AlipayOpenPublicMessageContentModifyResponseModel{}
	return &this
}

// GetContentId returns the ContentId field value if set, zero value otherwise.
func (o *AlipayOpenPublicMessageContentModifyResponseModel) GetContentId() string {
	if o == nil || IsNil(o.ContentId) {
		var ret string
		return ret
	}
	return *o.ContentId
}

// GetContentIdOk returns a tuple with the ContentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicMessageContentModifyResponseModel) GetContentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContentId) {
		return nil, false
	}
	return o.ContentId, true
}

// HasContentId returns a boolean if a field has been set.
func (o *AlipayOpenPublicMessageContentModifyResponseModel) HasContentId() bool {
	if o != nil && !IsNil(o.ContentId) {
		return true
	}

	return false
}

// SetContentId gets a reference to the given string and assigns it to the ContentId field.
func (o *AlipayOpenPublicMessageContentModifyResponseModel) SetContentId(v string) {
	o.ContentId = &v
}

// GetContentUrl returns the ContentUrl field value if set, zero value otherwise.
func (o *AlipayOpenPublicMessageContentModifyResponseModel) GetContentUrl() string {
	if o == nil || IsNil(o.ContentUrl) {
		var ret string
		return ret
	}
	return *o.ContentUrl
}

// GetContentUrlOk returns a tuple with the ContentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicMessageContentModifyResponseModel) GetContentUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ContentUrl) {
		return nil, false
	}
	return o.ContentUrl, true
}

// HasContentUrl returns a boolean if a field has been set.
func (o *AlipayOpenPublicMessageContentModifyResponseModel) HasContentUrl() bool {
	if o != nil && !IsNil(o.ContentUrl) {
		return true
	}

	return false
}

// SetContentUrl gets a reference to the given string and assigns it to the ContentUrl field.
func (o *AlipayOpenPublicMessageContentModifyResponseModel) SetContentUrl(v string) {
	o.ContentUrl = &v
}

func (o AlipayOpenPublicMessageContentModifyResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicMessageContentModifyResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContentId) {
		toSerialize["content_id"] = o.ContentId
	}
	if !IsNil(o.ContentUrl) {
		toSerialize["content_url"] = o.ContentUrl
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicMessageContentModifyResponseModel struct {
	value *AlipayOpenPublicMessageContentModifyResponseModel
	isSet bool
}

func (v NullableAlipayOpenPublicMessageContentModifyResponseModel) Get() *AlipayOpenPublicMessageContentModifyResponseModel {
	return v.value
}

func (v *NullableAlipayOpenPublicMessageContentModifyResponseModel) Set(val *AlipayOpenPublicMessageContentModifyResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicMessageContentModifyResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicMessageContentModifyResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicMessageContentModifyResponseModel(val *AlipayOpenPublicMessageContentModifyResponseModel) *NullableAlipayOpenPublicMessageContentModifyResponseModel {
	return &NullableAlipayOpenPublicMessageContentModifyResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicMessageContentModifyResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicMessageContentModifyResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


