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

// checks if the AlipayMerchantTradecomplainSupplementSubmitModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMerchantTradecomplainSupplementSubmitModel{}

// AlipayMerchantTradecomplainSupplementSubmitModel struct for AlipayMerchantTradecomplainSupplementSubmitModel
type AlipayMerchantTradecomplainSupplementSubmitModel struct {
	// 支付宝侧投诉单号
	ComplainEventId *string `json:"complain_event_id,omitempty"`
	// 回复内容，最多不超过200个字
	SupplementContent *string `json:"supplement_content,omitempty"`
	// 商家补充凭证时的图片id，多个逗号隔开（图片id可以通过\"商户上传处理图片\"接口获取）
	SupplementImages *string `json:"supplement_images,omitempty"`
}

// NewAlipayMerchantTradecomplainSupplementSubmitModel instantiates a new AlipayMerchantTradecomplainSupplementSubmitModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMerchantTradecomplainSupplementSubmitModel() *AlipayMerchantTradecomplainSupplementSubmitModel {
	this := AlipayMerchantTradecomplainSupplementSubmitModel{}
	return &this
}

// NewAlipayMerchantTradecomplainSupplementSubmitModelWithDefaults instantiates a new AlipayMerchantTradecomplainSupplementSubmitModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMerchantTradecomplainSupplementSubmitModelWithDefaults() *AlipayMerchantTradecomplainSupplementSubmitModel {
	this := AlipayMerchantTradecomplainSupplementSubmitModel{}
	return &this
}

// GetComplainEventId returns the ComplainEventId field value if set, zero value otherwise.
func (o *AlipayMerchantTradecomplainSupplementSubmitModel) GetComplainEventId() string {
	if o == nil || IsNil(o.ComplainEventId) {
		var ret string
		return ret
	}
	return *o.ComplainEventId
}

// GetComplainEventIdOk returns a tuple with the ComplainEventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantTradecomplainSupplementSubmitModel) GetComplainEventIdOk() (*string, bool) {
	if o == nil || IsNil(o.ComplainEventId) {
		return nil, false
	}
	return o.ComplainEventId, true
}

// HasComplainEventId returns a boolean if a field has been set.
func (o *AlipayMerchantTradecomplainSupplementSubmitModel) HasComplainEventId() bool {
	if o != nil && !IsNil(o.ComplainEventId) {
		return true
	}

	return false
}

// SetComplainEventId gets a reference to the given string and assigns it to the ComplainEventId field.
func (o *AlipayMerchantTradecomplainSupplementSubmitModel) SetComplainEventId(v string) {
	o.ComplainEventId = &v
}

// GetSupplementContent returns the SupplementContent field value if set, zero value otherwise.
func (o *AlipayMerchantTradecomplainSupplementSubmitModel) GetSupplementContent() string {
	if o == nil || IsNil(o.SupplementContent) {
		var ret string
		return ret
	}
	return *o.SupplementContent
}

// GetSupplementContentOk returns a tuple with the SupplementContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantTradecomplainSupplementSubmitModel) GetSupplementContentOk() (*string, bool) {
	if o == nil || IsNil(o.SupplementContent) {
		return nil, false
	}
	return o.SupplementContent, true
}

// HasSupplementContent returns a boolean if a field has been set.
func (o *AlipayMerchantTradecomplainSupplementSubmitModel) HasSupplementContent() bool {
	if o != nil && !IsNil(o.SupplementContent) {
		return true
	}

	return false
}

// SetSupplementContent gets a reference to the given string and assigns it to the SupplementContent field.
func (o *AlipayMerchantTradecomplainSupplementSubmitModel) SetSupplementContent(v string) {
	o.SupplementContent = &v
}

// GetSupplementImages returns the SupplementImages field value if set, zero value otherwise.
func (o *AlipayMerchantTradecomplainSupplementSubmitModel) GetSupplementImages() string {
	if o == nil || IsNil(o.SupplementImages) {
		var ret string
		return ret
	}
	return *o.SupplementImages
}

// GetSupplementImagesOk returns a tuple with the SupplementImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantTradecomplainSupplementSubmitModel) GetSupplementImagesOk() (*string, bool) {
	if o == nil || IsNil(o.SupplementImages) {
		return nil, false
	}
	return o.SupplementImages, true
}

// HasSupplementImages returns a boolean if a field has been set.
func (o *AlipayMerchantTradecomplainSupplementSubmitModel) HasSupplementImages() bool {
	if o != nil && !IsNil(o.SupplementImages) {
		return true
	}

	return false
}

// SetSupplementImages gets a reference to the given string and assigns it to the SupplementImages field.
func (o *AlipayMerchantTradecomplainSupplementSubmitModel) SetSupplementImages(v string) {
	o.SupplementImages = &v
}

func (o AlipayMerchantTradecomplainSupplementSubmitModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMerchantTradecomplainSupplementSubmitModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComplainEventId) {
		toSerialize["complain_event_id"] = o.ComplainEventId
	}
	if !IsNil(o.SupplementContent) {
		toSerialize["supplement_content"] = o.SupplementContent
	}
	if !IsNil(o.SupplementImages) {
		toSerialize["supplement_images"] = o.SupplementImages
	}
	return toSerialize, nil
}

type NullableAlipayMerchantTradecomplainSupplementSubmitModel struct {
	value *AlipayMerchantTradecomplainSupplementSubmitModel
	isSet bool
}

func (v NullableAlipayMerchantTradecomplainSupplementSubmitModel) Get() *AlipayMerchantTradecomplainSupplementSubmitModel {
	return v.value
}

func (v *NullableAlipayMerchantTradecomplainSupplementSubmitModel) Set(val *AlipayMerchantTradecomplainSupplementSubmitModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMerchantTradecomplainSupplementSubmitModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMerchantTradecomplainSupplementSubmitModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMerchantTradecomplainSupplementSubmitModel(val *AlipayMerchantTradecomplainSupplementSubmitModel) *NullableAlipayMerchantTradecomplainSupplementSubmitModel {
	return &NullableAlipayMerchantTradecomplainSupplementSubmitModel{value: val, isSet: true}
}

func (v NullableAlipayMerchantTradecomplainSupplementSubmitModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMerchantTradecomplainSupplementSubmitModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
