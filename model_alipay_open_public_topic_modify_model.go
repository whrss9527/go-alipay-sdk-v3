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

// checks if the AlipayOpenPublicTopicModifyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicTopicModifyModel{}

// AlipayOpenPublicTopicModifyModel struct for AlipayOpenPublicTopicModifyModel
type AlipayOpenPublicTopicModifyModel struct {
	// 营销位图片url, 尺寸为996*450，最大不超过5M，支持格式:.jpg、.png ，请先调用<a href=\"https://docs.open.alipay.com/api_3/alipay.offline.material.image.upload\"> 图片上传接口</a>获得图片url。营销位需要展示头图时，必须填写该参数。
	ImgUrl *string `json:"img_url,omitempty"`
	// 跳转类型，网页:HTTP、小程序:APP，不传默认HTTP
	LinkType *string `json:"link_type,omitempty"`
	// 营销位跳转地址，点击营销位头图跳到的链接url。营销位需要展示头图时，必须填写该参数。
	LinkUrl *string `json:"link_url,omitempty"`
	// 营销位描述。营销位需要展示头图时，必须填写该参数
	SubTitle *string `json:"sub_title,omitempty"`
	// 营销位名称
	Title *string `json:"title,omitempty"`
	// 营销位id
	TopicId *string `json:"topic_id,omitempty"`
	TopicItems *TopicItem `json:"topic_items,omitempty"`
}

// NewAlipayOpenPublicTopicModifyModel instantiates a new AlipayOpenPublicTopicModifyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicTopicModifyModel() *AlipayOpenPublicTopicModifyModel {
	this := AlipayOpenPublicTopicModifyModel{}
	return &this
}

// NewAlipayOpenPublicTopicModifyModelWithDefaults instantiates a new AlipayOpenPublicTopicModifyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicTopicModifyModelWithDefaults() *AlipayOpenPublicTopicModifyModel {
	this := AlipayOpenPublicTopicModifyModel{}
	return &this
}

// GetImgUrl returns the ImgUrl field value if set, zero value otherwise.
func (o *AlipayOpenPublicTopicModifyModel) GetImgUrl() string {
	if o == nil || IsNil(o.ImgUrl) {
		var ret string
		return ret
	}
	return *o.ImgUrl
}

// GetImgUrlOk returns a tuple with the ImgUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicTopicModifyModel) GetImgUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImgUrl) {
		return nil, false
	}
	return o.ImgUrl, true
}

// HasImgUrl returns a boolean if a field has been set.
func (o *AlipayOpenPublicTopicModifyModel) HasImgUrl() bool {
	if o != nil && !IsNil(o.ImgUrl) {
		return true
	}

	return false
}

// SetImgUrl gets a reference to the given string and assigns it to the ImgUrl field.
func (o *AlipayOpenPublicTopicModifyModel) SetImgUrl(v string) {
	o.ImgUrl = &v
}

// GetLinkType returns the LinkType field value if set, zero value otherwise.
func (o *AlipayOpenPublicTopicModifyModel) GetLinkType() string {
	if o == nil || IsNil(o.LinkType) {
		var ret string
		return ret
	}
	return *o.LinkType
}

// GetLinkTypeOk returns a tuple with the LinkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicTopicModifyModel) GetLinkTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LinkType) {
		return nil, false
	}
	return o.LinkType, true
}

// HasLinkType returns a boolean if a field has been set.
func (o *AlipayOpenPublicTopicModifyModel) HasLinkType() bool {
	if o != nil && !IsNil(o.LinkType) {
		return true
	}

	return false
}

// SetLinkType gets a reference to the given string and assigns it to the LinkType field.
func (o *AlipayOpenPublicTopicModifyModel) SetLinkType(v string) {
	o.LinkType = &v
}

// GetLinkUrl returns the LinkUrl field value if set, zero value otherwise.
func (o *AlipayOpenPublicTopicModifyModel) GetLinkUrl() string {
	if o == nil || IsNil(o.LinkUrl) {
		var ret string
		return ret
	}
	return *o.LinkUrl
}

// GetLinkUrlOk returns a tuple with the LinkUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicTopicModifyModel) GetLinkUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LinkUrl) {
		return nil, false
	}
	return o.LinkUrl, true
}

// HasLinkUrl returns a boolean if a field has been set.
func (o *AlipayOpenPublicTopicModifyModel) HasLinkUrl() bool {
	if o != nil && !IsNil(o.LinkUrl) {
		return true
	}

	return false
}

// SetLinkUrl gets a reference to the given string and assigns it to the LinkUrl field.
func (o *AlipayOpenPublicTopicModifyModel) SetLinkUrl(v string) {
	o.LinkUrl = &v
}

// GetSubTitle returns the SubTitle field value if set, zero value otherwise.
func (o *AlipayOpenPublicTopicModifyModel) GetSubTitle() string {
	if o == nil || IsNil(o.SubTitle) {
		var ret string
		return ret
	}
	return *o.SubTitle
}

// GetSubTitleOk returns a tuple with the SubTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicTopicModifyModel) GetSubTitleOk() (*string, bool) {
	if o == nil || IsNil(o.SubTitle) {
		return nil, false
	}
	return o.SubTitle, true
}

// HasSubTitle returns a boolean if a field has been set.
func (o *AlipayOpenPublicTopicModifyModel) HasSubTitle() bool {
	if o != nil && !IsNil(o.SubTitle) {
		return true
	}

	return false
}

// SetSubTitle gets a reference to the given string and assigns it to the SubTitle field.
func (o *AlipayOpenPublicTopicModifyModel) SetSubTitle(v string) {
	o.SubTitle = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AlipayOpenPublicTopicModifyModel) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicTopicModifyModel) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AlipayOpenPublicTopicModifyModel) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AlipayOpenPublicTopicModifyModel) SetTitle(v string) {
	o.Title = &v
}

// GetTopicId returns the TopicId field value if set, zero value otherwise.
func (o *AlipayOpenPublicTopicModifyModel) GetTopicId() string {
	if o == nil || IsNil(o.TopicId) {
		var ret string
		return ret
	}
	return *o.TopicId
}

// GetTopicIdOk returns a tuple with the TopicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicTopicModifyModel) GetTopicIdOk() (*string, bool) {
	if o == nil || IsNil(o.TopicId) {
		return nil, false
	}
	return o.TopicId, true
}

// HasTopicId returns a boolean if a field has been set.
func (o *AlipayOpenPublicTopicModifyModel) HasTopicId() bool {
	if o != nil && !IsNil(o.TopicId) {
		return true
	}

	return false
}

// SetTopicId gets a reference to the given string and assigns it to the TopicId field.
func (o *AlipayOpenPublicTopicModifyModel) SetTopicId(v string) {
	o.TopicId = &v
}

// GetTopicItems returns the TopicItems field value if set, zero value otherwise.
func (o *AlipayOpenPublicTopicModifyModel) GetTopicItems() TopicItem {
	if o == nil || IsNil(o.TopicItems) {
		var ret TopicItem
		return ret
	}
	return *o.TopicItems
}

// GetTopicItemsOk returns a tuple with the TopicItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicTopicModifyModel) GetTopicItemsOk() (*TopicItem, bool) {
	if o == nil || IsNil(o.TopicItems) {
		return nil, false
	}
	return o.TopicItems, true
}

// HasTopicItems returns a boolean if a field has been set.
func (o *AlipayOpenPublicTopicModifyModel) HasTopicItems() bool {
	if o != nil && !IsNil(o.TopicItems) {
		return true
	}

	return false
}

// SetTopicItems gets a reference to the given TopicItem and assigns it to the TopicItems field.
func (o *AlipayOpenPublicTopicModifyModel) SetTopicItems(v TopicItem) {
	o.TopicItems = &v
}

func (o AlipayOpenPublicTopicModifyModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicTopicModifyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImgUrl) {
		toSerialize["img_url"] = o.ImgUrl
	}
	if !IsNil(o.LinkType) {
		toSerialize["link_type"] = o.LinkType
	}
	if !IsNil(o.LinkUrl) {
		toSerialize["link_url"] = o.LinkUrl
	}
	if !IsNil(o.SubTitle) {
		toSerialize["sub_title"] = o.SubTitle
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.TopicId) {
		toSerialize["topic_id"] = o.TopicId
	}
	if !IsNil(o.TopicItems) {
		toSerialize["topic_items"] = o.TopicItems
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicTopicModifyModel struct {
	value *AlipayOpenPublicTopicModifyModel
	isSet bool
}

func (v NullableAlipayOpenPublicTopicModifyModel) Get() *AlipayOpenPublicTopicModifyModel {
	return v.value
}

func (v *NullableAlipayOpenPublicTopicModifyModel) Set(val *AlipayOpenPublicTopicModifyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicTopicModifyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicTopicModifyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicTopicModifyModel(val *AlipayOpenPublicTopicModifyModel) *NullableAlipayOpenPublicTopicModifyModel {
	return &NullableAlipayOpenPublicTopicModifyModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicTopicModifyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicTopicModifyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


