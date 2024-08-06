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

// checks if the SearchOrderDetailDataBrandItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchOrderDetailDataBrandItems{}

// SearchOrderDetailDataBrandItems struct for SearchOrderDetailDataBrandItems
type SearchOrderDetailDataBrandItems struct {
	// 工单详情bizid
	BizId *string `json:"biz_id,omitempty"`
	// 上下架状态
	BoxStatus *string `json:"box_status,omitempty"`
	// 关键词信息
	BrandBoxKeywords *string `json:"brand_box_keywords,omitempty"`
	// 工单详情数据信息
	BrandDetailList []SearchOrderBrandDetail `json:"brand_detail_list,omitempty"`
	// 品牌展示模板类型
	BrandTemplateId *string `json:"brand_template_id,omitempty"`
	// 工单详情数据channel
	Channel *string `json:"channel,omitempty"`
	// 工单详情数据merchant_type
	MerchantType *string `json:"merchant_type,omitempty"`
	// 模板id
	TemplateId *string `json:"template_id,omitempty"`
}

// NewSearchOrderDetailDataBrandItems instantiates a new SearchOrderDetailDataBrandItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchOrderDetailDataBrandItems() *SearchOrderDetailDataBrandItems {
	this := SearchOrderDetailDataBrandItems{}
	return &this
}

// NewSearchOrderDetailDataBrandItemsWithDefaults instantiates a new SearchOrderDetailDataBrandItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchOrderDetailDataBrandItemsWithDefaults() *SearchOrderDetailDataBrandItems {
	this := SearchOrderDetailDataBrandItems{}
	return &this
}

// GetBizId returns the BizId field value if set, zero value otherwise.
func (o *SearchOrderDetailDataBrandItems) GetBizId() string {
	if o == nil || IsNil(o.BizId) {
		var ret string
		return ret
	}
	return *o.BizId
}

// GetBizIdOk returns a tuple with the BizId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderDetailDataBrandItems) GetBizIdOk() (*string, bool) {
	if o == nil || IsNil(o.BizId) {
		return nil, false
	}
	return o.BizId, true
}

// HasBizId returns a boolean if a field has been set.
func (o *SearchOrderDetailDataBrandItems) HasBizId() bool {
	if o != nil && !IsNil(o.BizId) {
		return true
	}

	return false
}

// SetBizId gets a reference to the given string and assigns it to the BizId field.
func (o *SearchOrderDetailDataBrandItems) SetBizId(v string) {
	o.BizId = &v
}

// GetBoxStatus returns the BoxStatus field value if set, zero value otherwise.
func (o *SearchOrderDetailDataBrandItems) GetBoxStatus() string {
	if o == nil || IsNil(o.BoxStatus) {
		var ret string
		return ret
	}
	return *o.BoxStatus
}

// GetBoxStatusOk returns a tuple with the BoxStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderDetailDataBrandItems) GetBoxStatusOk() (*string, bool) {
	if o == nil || IsNil(o.BoxStatus) {
		return nil, false
	}
	return o.BoxStatus, true
}

// HasBoxStatus returns a boolean if a field has been set.
func (o *SearchOrderDetailDataBrandItems) HasBoxStatus() bool {
	if o != nil && !IsNil(o.BoxStatus) {
		return true
	}

	return false
}

// SetBoxStatus gets a reference to the given string and assigns it to the BoxStatus field.
func (o *SearchOrderDetailDataBrandItems) SetBoxStatus(v string) {
	o.BoxStatus = &v
}

// GetBrandBoxKeywords returns the BrandBoxKeywords field value if set, zero value otherwise.
func (o *SearchOrderDetailDataBrandItems) GetBrandBoxKeywords() string {
	if o == nil || IsNil(o.BrandBoxKeywords) {
		var ret string
		return ret
	}
	return *o.BrandBoxKeywords
}

// GetBrandBoxKeywordsOk returns a tuple with the BrandBoxKeywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderDetailDataBrandItems) GetBrandBoxKeywordsOk() (*string, bool) {
	if o == nil || IsNil(o.BrandBoxKeywords) {
		return nil, false
	}
	return o.BrandBoxKeywords, true
}

// HasBrandBoxKeywords returns a boolean if a field has been set.
func (o *SearchOrderDetailDataBrandItems) HasBrandBoxKeywords() bool {
	if o != nil && !IsNil(o.BrandBoxKeywords) {
		return true
	}

	return false
}

// SetBrandBoxKeywords gets a reference to the given string and assigns it to the BrandBoxKeywords field.
func (o *SearchOrderDetailDataBrandItems) SetBrandBoxKeywords(v string) {
	o.BrandBoxKeywords = &v
}

// GetBrandDetailList returns the BrandDetailList field value if set, zero value otherwise.
func (o *SearchOrderDetailDataBrandItems) GetBrandDetailList() []SearchOrderBrandDetail {
	if o == nil || IsNil(o.BrandDetailList) {
		var ret []SearchOrderBrandDetail
		return ret
	}
	return o.BrandDetailList
}

// GetBrandDetailListOk returns a tuple with the BrandDetailList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderDetailDataBrandItems) GetBrandDetailListOk() ([]SearchOrderBrandDetail, bool) {
	if o == nil || IsNil(o.BrandDetailList) {
		return nil, false
	}
	return o.BrandDetailList, true
}

// HasBrandDetailList returns a boolean if a field has been set.
func (o *SearchOrderDetailDataBrandItems) HasBrandDetailList() bool {
	if o != nil && !IsNil(o.BrandDetailList) {
		return true
	}

	return false
}

// SetBrandDetailList gets a reference to the given []SearchOrderBrandDetail and assigns it to the BrandDetailList field.
func (o *SearchOrderDetailDataBrandItems) SetBrandDetailList(v []SearchOrderBrandDetail) {
	o.BrandDetailList = v
}

// GetBrandTemplateId returns the BrandTemplateId field value if set, zero value otherwise.
func (o *SearchOrderDetailDataBrandItems) GetBrandTemplateId() string {
	if o == nil || IsNil(o.BrandTemplateId) {
		var ret string
		return ret
	}
	return *o.BrandTemplateId
}

// GetBrandTemplateIdOk returns a tuple with the BrandTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderDetailDataBrandItems) GetBrandTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.BrandTemplateId) {
		return nil, false
	}
	return o.BrandTemplateId, true
}

// HasBrandTemplateId returns a boolean if a field has been set.
func (o *SearchOrderDetailDataBrandItems) HasBrandTemplateId() bool {
	if o != nil && !IsNil(o.BrandTemplateId) {
		return true
	}

	return false
}

// SetBrandTemplateId gets a reference to the given string and assigns it to the BrandTemplateId field.
func (o *SearchOrderDetailDataBrandItems) SetBrandTemplateId(v string) {
	o.BrandTemplateId = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *SearchOrderDetailDataBrandItems) GetChannel() string {
	if o == nil || IsNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderDetailDataBrandItems) GetChannelOk() (*string, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *SearchOrderDetailDataBrandItems) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *SearchOrderDetailDataBrandItems) SetChannel(v string) {
	o.Channel = &v
}

// GetMerchantType returns the MerchantType field value if set, zero value otherwise.
func (o *SearchOrderDetailDataBrandItems) GetMerchantType() string {
	if o == nil || IsNil(o.MerchantType) {
		var ret string
		return ret
	}
	return *o.MerchantType
}

// GetMerchantTypeOk returns a tuple with the MerchantType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderDetailDataBrandItems) GetMerchantTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantType) {
		return nil, false
	}
	return o.MerchantType, true
}

// HasMerchantType returns a boolean if a field has been set.
func (o *SearchOrderDetailDataBrandItems) HasMerchantType() bool {
	if o != nil && !IsNil(o.MerchantType) {
		return true
	}

	return false
}

// SetMerchantType gets a reference to the given string and assigns it to the MerchantType field.
func (o *SearchOrderDetailDataBrandItems) SetMerchantType(v string) {
	o.MerchantType = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *SearchOrderDetailDataBrandItems) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId) {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderDetailDataBrandItems) GetTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *SearchOrderDetailDataBrandItems) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *SearchOrderDetailDataBrandItems) SetTemplateId(v string) {
	o.TemplateId = &v
}

func (o SearchOrderDetailDataBrandItems) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchOrderDetailDataBrandItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BizId) {
		toSerialize["biz_id"] = o.BizId
	}
	if !IsNil(o.BoxStatus) {
		toSerialize["box_status"] = o.BoxStatus
	}
	if !IsNil(o.BrandBoxKeywords) {
		toSerialize["brand_box_keywords"] = o.BrandBoxKeywords
	}
	if !IsNil(o.BrandDetailList) {
		toSerialize["brand_detail_list"] = o.BrandDetailList
	}
	if !IsNil(o.BrandTemplateId) {
		toSerialize["brand_template_id"] = o.BrandTemplateId
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.MerchantType) {
		toSerialize["merchant_type"] = o.MerchantType
	}
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	return toSerialize, nil
}

type NullableSearchOrderDetailDataBrandItems struct {
	value *SearchOrderDetailDataBrandItems
	isSet bool
}

func (v NullableSearchOrderDetailDataBrandItems) Get() *SearchOrderDetailDataBrandItems {
	return v.value
}

func (v *NullableSearchOrderDetailDataBrandItems) Set(val *SearchOrderDetailDataBrandItems) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchOrderDetailDataBrandItems) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchOrderDetailDataBrandItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchOrderDetailDataBrandItems(val *SearchOrderDetailDataBrandItems) *NullableSearchOrderDetailDataBrandItems {
	return &NullableSearchOrderDetailDataBrandItems{value: val, isSet: true}
}

func (v NullableSearchOrderDetailDataBrandItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchOrderDetailDataBrandItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
