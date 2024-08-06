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

// checks if the SearchOrderDetailDataBaseItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchOrderDetailDataBaseItems{}

// SearchOrderDetailDataBaseItems struct for SearchOrderDetailDataBaseItems
type SearchOrderDetailDataBaseItems struct {
	// 搜索可见性
	CanSearch *string `json:"can_search,omitempty"`
	// 描述信息
	Desc *string `json:"desc,omitempty"`
	// img地址
	Img *string `json:"img,omitempty"`
	// 关键词
	KeyWord *string `json:"key_word,omitempty"`
	// 应用名称
	Name *string `json:"name,omitempty"`
	// 基础工单详情region
	Region *string `json:"region,omitempty"`
	// 搜索可见性
	ServCanSearch *string `json:"serv_can_search,omitempty"`
	// 搜索关键词
	ServSearchKeywords *string `json:"serv_search_keywords,omitempty"`
	// 模板id
	TemplateId *string `json:"template_id,omitempty"`
	// 跳转url链接
	Url *string `json:"url,omitempty"`
}

// NewSearchOrderDetailDataBaseItems instantiates a new SearchOrderDetailDataBaseItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchOrderDetailDataBaseItems() *SearchOrderDetailDataBaseItems {
	this := SearchOrderDetailDataBaseItems{}
	return &this
}

// NewSearchOrderDetailDataBaseItemsWithDefaults instantiates a new SearchOrderDetailDataBaseItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchOrderDetailDataBaseItemsWithDefaults() *SearchOrderDetailDataBaseItems {
	this := SearchOrderDetailDataBaseItems{}
	return &this
}

// GetCanSearch returns the CanSearch field value if set, zero value otherwise.
func (o *SearchOrderDetailDataBaseItems) GetCanSearch() string {
	if o == nil || IsNil(o.CanSearch) {
		var ret string
		return ret
	}
	return *o.CanSearch
}

// GetCanSearchOk returns a tuple with the CanSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderDetailDataBaseItems) GetCanSearchOk() (*string, bool) {
	if o == nil || IsNil(o.CanSearch) {
		return nil, false
	}
	return o.CanSearch, true
}

// HasCanSearch returns a boolean if a field has been set.
func (o *SearchOrderDetailDataBaseItems) HasCanSearch() bool {
	if o != nil && !IsNil(o.CanSearch) {
		return true
	}

	return false
}

// SetCanSearch gets a reference to the given string and assigns it to the CanSearch field.
func (o *SearchOrderDetailDataBaseItems) SetCanSearch(v string) {
	o.CanSearch = &v
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *SearchOrderDetailDataBaseItems) GetDesc() string {
	if o == nil || IsNil(o.Desc) {
		var ret string
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderDetailDataBaseItems) GetDescOk() (*string, bool) {
	if o == nil || IsNil(o.Desc) {
		return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *SearchOrderDetailDataBaseItems) HasDesc() bool {
	if o != nil && !IsNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given string and assigns it to the Desc field.
func (o *SearchOrderDetailDataBaseItems) SetDesc(v string) {
	o.Desc = &v
}

// GetImg returns the Img field value if set, zero value otherwise.
func (o *SearchOrderDetailDataBaseItems) GetImg() string {
	if o == nil || IsNil(o.Img) {
		var ret string
		return ret
	}
	return *o.Img
}

// GetImgOk returns a tuple with the Img field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderDetailDataBaseItems) GetImgOk() (*string, bool) {
	if o == nil || IsNil(o.Img) {
		return nil, false
	}
	return o.Img, true
}

// HasImg returns a boolean if a field has been set.
func (o *SearchOrderDetailDataBaseItems) HasImg() bool {
	if o != nil && !IsNil(o.Img) {
		return true
	}

	return false
}

// SetImg gets a reference to the given string and assigns it to the Img field.
func (o *SearchOrderDetailDataBaseItems) SetImg(v string) {
	o.Img = &v
}

// GetKeyWord returns the KeyWord field value if set, zero value otherwise.
func (o *SearchOrderDetailDataBaseItems) GetKeyWord() string {
	if o == nil || IsNil(o.KeyWord) {
		var ret string
		return ret
	}
	return *o.KeyWord
}

// GetKeyWordOk returns a tuple with the KeyWord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderDetailDataBaseItems) GetKeyWordOk() (*string, bool) {
	if o == nil || IsNil(o.KeyWord) {
		return nil, false
	}
	return o.KeyWord, true
}

// HasKeyWord returns a boolean if a field has been set.
func (o *SearchOrderDetailDataBaseItems) HasKeyWord() bool {
	if o != nil && !IsNil(o.KeyWord) {
		return true
	}

	return false
}

// SetKeyWord gets a reference to the given string and assigns it to the KeyWord field.
func (o *SearchOrderDetailDataBaseItems) SetKeyWord(v string) {
	o.KeyWord = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SearchOrderDetailDataBaseItems) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderDetailDataBaseItems) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SearchOrderDetailDataBaseItems) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SearchOrderDetailDataBaseItems) SetName(v string) {
	o.Name = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *SearchOrderDetailDataBaseItems) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderDetailDataBaseItems) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *SearchOrderDetailDataBaseItems) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *SearchOrderDetailDataBaseItems) SetRegion(v string) {
	o.Region = &v
}

// GetServCanSearch returns the ServCanSearch field value if set, zero value otherwise.
func (o *SearchOrderDetailDataBaseItems) GetServCanSearch() string {
	if o == nil || IsNil(o.ServCanSearch) {
		var ret string
		return ret
	}
	return *o.ServCanSearch
}

// GetServCanSearchOk returns a tuple with the ServCanSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderDetailDataBaseItems) GetServCanSearchOk() (*string, bool) {
	if o == nil || IsNil(o.ServCanSearch) {
		return nil, false
	}
	return o.ServCanSearch, true
}

// HasServCanSearch returns a boolean if a field has been set.
func (o *SearchOrderDetailDataBaseItems) HasServCanSearch() bool {
	if o != nil && !IsNil(o.ServCanSearch) {
		return true
	}

	return false
}

// SetServCanSearch gets a reference to the given string and assigns it to the ServCanSearch field.
func (o *SearchOrderDetailDataBaseItems) SetServCanSearch(v string) {
	o.ServCanSearch = &v
}

// GetServSearchKeywords returns the ServSearchKeywords field value if set, zero value otherwise.
func (o *SearchOrderDetailDataBaseItems) GetServSearchKeywords() string {
	if o == nil || IsNil(o.ServSearchKeywords) {
		var ret string
		return ret
	}
	return *o.ServSearchKeywords
}

// GetServSearchKeywordsOk returns a tuple with the ServSearchKeywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderDetailDataBaseItems) GetServSearchKeywordsOk() (*string, bool) {
	if o == nil || IsNil(o.ServSearchKeywords) {
		return nil, false
	}
	return o.ServSearchKeywords, true
}

// HasServSearchKeywords returns a boolean if a field has been set.
func (o *SearchOrderDetailDataBaseItems) HasServSearchKeywords() bool {
	if o != nil && !IsNil(o.ServSearchKeywords) {
		return true
	}

	return false
}

// SetServSearchKeywords gets a reference to the given string and assigns it to the ServSearchKeywords field.
func (o *SearchOrderDetailDataBaseItems) SetServSearchKeywords(v string) {
	o.ServSearchKeywords = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *SearchOrderDetailDataBaseItems) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId) {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderDetailDataBaseItems) GetTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *SearchOrderDetailDataBaseItems) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *SearchOrderDetailDataBaseItems) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SearchOrderDetailDataBaseItems) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderDetailDataBaseItems) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SearchOrderDetailDataBaseItems) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SearchOrderDetailDataBaseItems) SetUrl(v string) {
	o.Url = &v
}

func (o SearchOrderDetailDataBaseItems) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchOrderDetailDataBaseItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CanSearch) {
		toSerialize["can_search"] = o.CanSearch
	}
	if !IsNil(o.Desc) {
		toSerialize["desc"] = o.Desc
	}
	if !IsNil(o.Img) {
		toSerialize["img"] = o.Img
	}
	if !IsNil(o.KeyWord) {
		toSerialize["key_word"] = o.KeyWord
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.ServCanSearch) {
		toSerialize["serv_can_search"] = o.ServCanSearch
	}
	if !IsNil(o.ServSearchKeywords) {
		toSerialize["serv_search_keywords"] = o.ServSearchKeywords
	}
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableSearchOrderDetailDataBaseItems struct {
	value *SearchOrderDetailDataBaseItems
	isSet bool
}

func (v NullableSearchOrderDetailDataBaseItems) Get() *SearchOrderDetailDataBaseItems {
	return v.value
}

func (v *NullableSearchOrderDetailDataBaseItems) Set(val *SearchOrderDetailDataBaseItems) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchOrderDetailDataBaseItems) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchOrderDetailDataBaseItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchOrderDetailDataBaseItems(val *SearchOrderDetailDataBaseItems) *NullableSearchOrderDetailDataBaseItems {
	return &NullableSearchOrderDetailDataBaseItems{value: val, isSet: true}
}

func (v NullableSearchOrderDetailDataBaseItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchOrderDetailDataBaseItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
