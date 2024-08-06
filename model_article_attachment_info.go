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

// checks if the ArticleAttachmentInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArticleAttachmentInfo{}

// ArticleAttachmentInfo struct for ArticleAttachmentInfo
type ArticleAttachmentInfo struct {
	// 文章附件
	FileName *string `json:"file_name,omitempty"`
	// 附件完整Key
	Key *string `json:"key,omitempty"`
	// 附件的url
	Url *string `json:"url,omitempty"`
}

// NewArticleAttachmentInfo instantiates a new ArticleAttachmentInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArticleAttachmentInfo() *ArticleAttachmentInfo {
	this := ArticleAttachmentInfo{}
	return &this
}

// NewArticleAttachmentInfoWithDefaults instantiates a new ArticleAttachmentInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArticleAttachmentInfoWithDefaults() *ArticleAttachmentInfo {
	this := ArticleAttachmentInfo{}
	return &this
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *ArticleAttachmentInfo) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArticleAttachmentInfo) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *ArticleAttachmentInfo) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *ArticleAttachmentInfo) SetFileName(v string) {
	o.FileName = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ArticleAttachmentInfo) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArticleAttachmentInfo) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ArticleAttachmentInfo) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ArticleAttachmentInfo) SetKey(v string) {
	o.Key = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ArticleAttachmentInfo) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArticleAttachmentInfo) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ArticleAttachmentInfo) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ArticleAttachmentInfo) SetUrl(v string) {
	o.Url = &v
}

func (o ArticleAttachmentInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArticleAttachmentInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileName) {
		toSerialize["file_name"] = o.FileName
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableArticleAttachmentInfo struct {
	value *ArticleAttachmentInfo
	isSet bool
}

func (v NullableArticleAttachmentInfo) Get() *ArticleAttachmentInfo {
	return v.value
}

func (v *NullableArticleAttachmentInfo) Set(val *ArticleAttachmentInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableArticleAttachmentInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableArticleAttachmentInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArticleAttachmentInfo(val *ArticleAttachmentInfo) *NullableArticleAttachmentInfo {
	return &NullableArticleAttachmentInfo{value: val, isSet: true}
}

func (v NullableArticleAttachmentInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArticleAttachmentInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

