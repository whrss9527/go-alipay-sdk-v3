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

// checks if the Material type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Material{}

// Material struct for Material
type Material struct {
	// 图文消息子消息项集合，单条消息最多6个子项，否则会发送失败
	Articles []Article `json:"articles,omitempty"`
	// 消息类型，text：文本类型，image-text：图文类型。当消息类型为text时，text参数必传，当消息类型为image-text时，articles参数必传
	MsgType *string `json:"msg_type,omitempty"`
	Text *Text `json:"text,omitempty"`
}

// NewMaterial instantiates a new Material object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaterial() *Material {
	this := Material{}
	return &this
}

// NewMaterialWithDefaults instantiates a new Material object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaterialWithDefaults() *Material {
	this := Material{}
	return &this
}

// GetArticles returns the Articles field value if set, zero value otherwise.
func (o *Material) GetArticles() []Article {
	if o == nil || IsNil(o.Articles) {
		var ret []Article
		return ret
	}
	return o.Articles
}

// GetArticlesOk returns a tuple with the Articles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Material) GetArticlesOk() ([]Article, bool) {
	if o == nil || IsNil(o.Articles) {
		return nil, false
	}
	return o.Articles, true
}

// HasArticles returns a boolean if a field has been set.
func (o *Material) HasArticles() bool {
	if o != nil && !IsNil(o.Articles) {
		return true
	}

	return false
}

// SetArticles gets a reference to the given []Article and assigns it to the Articles field.
func (o *Material) SetArticles(v []Article) {
	o.Articles = v
}

// GetMsgType returns the MsgType field value if set, zero value otherwise.
func (o *Material) GetMsgType() string {
	if o == nil || IsNil(o.MsgType) {
		var ret string
		return ret
	}
	return *o.MsgType
}

// GetMsgTypeOk returns a tuple with the MsgType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Material) GetMsgTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MsgType) {
		return nil, false
	}
	return o.MsgType, true
}

// HasMsgType returns a boolean if a field has been set.
func (o *Material) HasMsgType() bool {
	if o != nil && !IsNil(o.MsgType) {
		return true
	}

	return false
}

// SetMsgType gets a reference to the given string and assigns it to the MsgType field.
func (o *Material) SetMsgType(v string) {
	o.MsgType = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *Material) GetText() Text {
	if o == nil || IsNil(o.Text) {
		var ret Text
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Material) GetTextOk() (*Text, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *Material) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Text and assigns it to the Text field.
func (o *Material) SetText(v Text) {
	o.Text = &v
}

func (o Material) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Material) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Articles) {
		toSerialize["articles"] = o.Articles
	}
	if !IsNil(o.MsgType) {
		toSerialize["msg_type"] = o.MsgType
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return toSerialize, nil
}

type NullableMaterial struct {
	value *Material
	isSet bool
}

func (v NullableMaterial) Get() *Material {
	return v.value
}

func (v *NullableMaterial) Set(val *Material) {
	v.value = val
	v.isSet = true
}

func (v NullableMaterial) IsSet() bool {
	return v.isSet
}

func (v *NullableMaterial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaterial(val *Material) *NullableMaterial {
	return &NullableMaterial{value: val, isSet: true}
}

func (v NullableMaterial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaterial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


