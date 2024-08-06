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

// checks if the Article type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Article{}

// Article struct for Article
type Article struct {
	// 链接文字
	ActionName *string `json:"action_name,omitempty"`
	// 图文消息内容
	Desc *string `json:"desc,omitempty"`
	// 图片链接，对于多条图文消息的第一条消息，该字段不能为空; 请先调用<a href=\"https://docs.open.alipay.com/api_3/alipay.offline.material.image.upload\"> 图片上传接口</a>获得图片url
	ImageUrl *string `json:"image_url,omitempty"`
	// 图文消息标题
	Title *string `json:"title,omitempty"`
	// 点击图文消息跳转的链接
	Url *string `json:"url,omitempty"`
}

// NewArticle instantiates a new Article object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArticle() *Article {
	this := Article{}
	return &this
}

// NewArticleWithDefaults instantiates a new Article object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArticleWithDefaults() *Article {
	this := Article{}
	return &this
}

// GetActionName returns the ActionName field value if set, zero value otherwise.
func (o *Article) GetActionName() string {
	if o == nil || IsNil(o.ActionName) {
		var ret string
		return ret
	}
	return *o.ActionName
}

// GetActionNameOk returns a tuple with the ActionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Article) GetActionNameOk() (*string, bool) {
	if o == nil || IsNil(o.ActionName) {
		return nil, false
	}
	return o.ActionName, true
}

// HasActionName returns a boolean if a field has been set.
func (o *Article) HasActionName() bool {
	if o != nil && !IsNil(o.ActionName) {
		return true
	}

	return false
}

// SetActionName gets a reference to the given string and assigns it to the ActionName field.
func (o *Article) SetActionName(v string) {
	o.ActionName = &v
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *Article) GetDesc() string {
	if o == nil || IsNil(o.Desc) {
		var ret string
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Article) GetDescOk() (*string, bool) {
	if o == nil || IsNil(o.Desc) {
		return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *Article) HasDesc() bool {
	if o != nil && !IsNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given string and assigns it to the Desc field.
func (o *Article) SetDesc(v string) {
	o.Desc = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *Article) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Article) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *Article) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *Article) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Article) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Article) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Article) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Article) SetTitle(v string) {
	o.Title = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Article) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Article) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Article) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Article) SetUrl(v string) {
	o.Url = &v
}

func (o Article) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Article) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionName) {
		toSerialize["action_name"] = o.ActionName
	}
	if !IsNil(o.Desc) {
		toSerialize["desc"] = o.Desc
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["image_url"] = o.ImageUrl
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableArticle struct {
	value *Article
	isSet bool
}

func (v NullableArticle) Get() *Article {
	return v.value
}

func (v *NullableArticle) Set(val *Article) {
	v.value = val
	v.isSet = true
}

func (v NullableArticle) IsSet() bool {
	return v.isSet
}

func (v *NullableArticle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArticle(val *Article) *NullableArticle {
	return &NullableArticle{value: val, isSet: true}
}

func (v NullableArticle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArticle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
