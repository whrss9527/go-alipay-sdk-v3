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

// checks if the ShopCategoryConfigInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShopCategoryConfigInfo{}

// ShopCategoryConfigInfo struct for ShopCategoryConfigInfo
type ShopCategoryConfigInfo struct {
	// 类目ID
	Id *string `json:"id,omitempty"`
	// 是否是叶子节点
	IsLeaf *string `json:"is_leaf,omitempty"`
	// 类目层级
	Level *string `json:"level,omitempty"`
	// 类目层级路径
	Link *string `json:"link,omitempty"`
	// 类目名称
	Nm *string `json:"nm,omitempty"`
}

// NewShopCategoryConfigInfo instantiates a new ShopCategoryConfigInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShopCategoryConfigInfo() *ShopCategoryConfigInfo {
	this := ShopCategoryConfigInfo{}
	return &this
}

// NewShopCategoryConfigInfoWithDefaults instantiates a new ShopCategoryConfigInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShopCategoryConfigInfoWithDefaults() *ShopCategoryConfigInfo {
	this := ShopCategoryConfigInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ShopCategoryConfigInfo) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopCategoryConfigInfo) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ShopCategoryConfigInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ShopCategoryConfigInfo) SetId(v string) {
	o.Id = &v
}

// GetIsLeaf returns the IsLeaf field value if set, zero value otherwise.
func (o *ShopCategoryConfigInfo) GetIsLeaf() string {
	if o == nil || IsNil(o.IsLeaf) {
		var ret string
		return ret
	}
	return *o.IsLeaf
}

// GetIsLeafOk returns a tuple with the IsLeaf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopCategoryConfigInfo) GetIsLeafOk() (*string, bool) {
	if o == nil || IsNil(o.IsLeaf) {
		return nil, false
	}
	return o.IsLeaf, true
}

// HasIsLeaf returns a boolean if a field has been set.
func (o *ShopCategoryConfigInfo) HasIsLeaf() bool {
	if o != nil && !IsNil(o.IsLeaf) {
		return true
	}

	return false
}

// SetIsLeaf gets a reference to the given string and assigns it to the IsLeaf field.
func (o *ShopCategoryConfigInfo) SetIsLeaf(v string) {
	o.IsLeaf = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *ShopCategoryConfigInfo) GetLevel() string {
	if o == nil || IsNil(o.Level) {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopCategoryConfigInfo) GetLevelOk() (*string, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *ShopCategoryConfigInfo) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *ShopCategoryConfigInfo) SetLevel(v string) {
	o.Level = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *ShopCategoryConfigInfo) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopCategoryConfigInfo) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *ShopCategoryConfigInfo) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *ShopCategoryConfigInfo) SetLink(v string) {
	o.Link = &v
}

// GetNm returns the Nm field value if set, zero value otherwise.
func (o *ShopCategoryConfigInfo) GetNm() string {
	if o == nil || IsNil(o.Nm) {
		var ret string
		return ret
	}
	return *o.Nm
}

// GetNmOk returns a tuple with the Nm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopCategoryConfigInfo) GetNmOk() (*string, bool) {
	if o == nil || IsNil(o.Nm) {
		return nil, false
	}
	return o.Nm, true
}

// HasNm returns a boolean if a field has been set.
func (o *ShopCategoryConfigInfo) HasNm() bool {
	if o != nil && !IsNil(o.Nm) {
		return true
	}

	return false
}

// SetNm gets a reference to the given string and assigns it to the Nm field.
func (o *ShopCategoryConfigInfo) SetNm(v string) {
	o.Nm = &v
}

func (o ShopCategoryConfigInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShopCategoryConfigInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsLeaf) {
		toSerialize["is_leaf"] = o.IsLeaf
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.Nm) {
		toSerialize["nm"] = o.Nm
	}
	return toSerialize, nil
}

type NullableShopCategoryConfigInfo struct {
	value *ShopCategoryConfigInfo
	isSet bool
}

func (v NullableShopCategoryConfigInfo) Get() *ShopCategoryConfigInfo {
	return v.value
}

func (v *NullableShopCategoryConfigInfo) Set(val *ShopCategoryConfigInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableShopCategoryConfigInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableShopCategoryConfigInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShopCategoryConfigInfo(val *ShopCategoryConfigInfo) *NullableShopCategoryConfigInfo {
	return &NullableShopCategoryConfigInfo{value: val, isSet: true}
}

func (v NullableShopCategoryConfigInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShopCategoryConfigInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
