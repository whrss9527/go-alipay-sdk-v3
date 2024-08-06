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

// checks if the AlipayOpenMiniResourcePromotionsourceNotifyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniResourcePromotionsourceNotifyModel{}

// AlipayOpenMiniResourcePromotionsourceNotifyModel struct for AlipayOpenMiniResourcePromotionsourceNotifyModel
type AlipayOpenMiniResourcePromotionsourceNotifyModel struct {
	// 媒体唤起时传入的支付宝id
	AuthorId *string `json:"author_id,omitempty"`
	// 媒体唤起时提供的额外参数值列表，这里可能有多个值，打平以后拼入。即url_decode(authcbparams)
	Params *string `json:"params,omitempty"`
	// 推广位id
	PromotionId *string `json:"promotion_id,omitempty"`
	// 推广位名称
	PromotionName *string `json:"promotion_name,omitempty"`
	// 媒体来源，标识调用方
	Source *string `json:"source,omitempty"`
}

// NewAlipayOpenMiniResourcePromotionsourceNotifyModel instantiates a new AlipayOpenMiniResourcePromotionsourceNotifyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniResourcePromotionsourceNotifyModel() *AlipayOpenMiniResourcePromotionsourceNotifyModel {
	this := AlipayOpenMiniResourcePromotionsourceNotifyModel{}
	return &this
}

// NewAlipayOpenMiniResourcePromotionsourceNotifyModelWithDefaults instantiates a new AlipayOpenMiniResourcePromotionsourceNotifyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniResourcePromotionsourceNotifyModelWithDefaults() *AlipayOpenMiniResourcePromotionsourceNotifyModel {
	this := AlipayOpenMiniResourcePromotionsourceNotifyModel{}
	return &this
}

// GetAuthorId returns the AuthorId field value if set, zero value otherwise.
func (o *AlipayOpenMiniResourcePromotionsourceNotifyModel) GetAuthorId() string {
	if o == nil || IsNil(o.AuthorId) {
		var ret string
		return ret
	}
	return *o.AuthorId
}

// GetAuthorIdOk returns a tuple with the AuthorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniResourcePromotionsourceNotifyModel) GetAuthorIdOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorId) {
		return nil, false
	}
	return o.AuthorId, true
}

// HasAuthorId returns a boolean if a field has been set.
func (o *AlipayOpenMiniResourcePromotionsourceNotifyModel) HasAuthorId() bool {
	if o != nil && !IsNil(o.AuthorId) {
		return true
	}

	return false
}

// SetAuthorId gets a reference to the given string and assigns it to the AuthorId field.
func (o *AlipayOpenMiniResourcePromotionsourceNotifyModel) SetAuthorId(v string) {
	o.AuthorId = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *AlipayOpenMiniResourcePromotionsourceNotifyModel) GetParams() string {
	if o == nil || IsNil(o.Params) {
		var ret string
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniResourcePromotionsourceNotifyModel) GetParamsOk() (*string, bool) {
	if o == nil || IsNil(o.Params) {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *AlipayOpenMiniResourcePromotionsourceNotifyModel) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given string and assigns it to the Params field.
func (o *AlipayOpenMiniResourcePromotionsourceNotifyModel) SetParams(v string) {
	o.Params = &v
}

// GetPromotionId returns the PromotionId field value if set, zero value otherwise.
func (o *AlipayOpenMiniResourcePromotionsourceNotifyModel) GetPromotionId() string {
	if o == nil || IsNil(o.PromotionId) {
		var ret string
		return ret
	}
	return *o.PromotionId
}

// GetPromotionIdOk returns a tuple with the PromotionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniResourcePromotionsourceNotifyModel) GetPromotionIdOk() (*string, bool) {
	if o == nil || IsNil(o.PromotionId) {
		return nil, false
	}
	return o.PromotionId, true
}

// HasPromotionId returns a boolean if a field has been set.
func (o *AlipayOpenMiniResourcePromotionsourceNotifyModel) HasPromotionId() bool {
	if o != nil && !IsNil(o.PromotionId) {
		return true
	}

	return false
}

// SetPromotionId gets a reference to the given string and assigns it to the PromotionId field.
func (o *AlipayOpenMiniResourcePromotionsourceNotifyModel) SetPromotionId(v string) {
	o.PromotionId = &v
}

// GetPromotionName returns the PromotionName field value if set, zero value otherwise.
func (o *AlipayOpenMiniResourcePromotionsourceNotifyModel) GetPromotionName() string {
	if o == nil || IsNil(o.PromotionName) {
		var ret string
		return ret
	}
	return *o.PromotionName
}

// GetPromotionNameOk returns a tuple with the PromotionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniResourcePromotionsourceNotifyModel) GetPromotionNameOk() (*string, bool) {
	if o == nil || IsNil(o.PromotionName) {
		return nil, false
	}
	return o.PromotionName, true
}

// HasPromotionName returns a boolean if a field has been set.
func (o *AlipayOpenMiniResourcePromotionsourceNotifyModel) HasPromotionName() bool {
	if o != nil && !IsNil(o.PromotionName) {
		return true
	}

	return false
}

// SetPromotionName gets a reference to the given string and assigns it to the PromotionName field.
func (o *AlipayOpenMiniResourcePromotionsourceNotifyModel) SetPromotionName(v string) {
	o.PromotionName = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *AlipayOpenMiniResourcePromotionsourceNotifyModel) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniResourcePromotionsourceNotifyModel) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *AlipayOpenMiniResourcePromotionsourceNotifyModel) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *AlipayOpenMiniResourcePromotionsourceNotifyModel) SetSource(v string) {
	o.Source = &v
}

func (o AlipayOpenMiniResourcePromotionsourceNotifyModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniResourcePromotionsourceNotifyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthorId) {
		toSerialize["author_id"] = o.AuthorId
	}
	if !IsNil(o.Params) {
		toSerialize["params"] = o.Params
	}
	if !IsNil(o.PromotionId) {
		toSerialize["promotion_id"] = o.PromotionId
	}
	if !IsNil(o.PromotionName) {
		toSerialize["promotion_name"] = o.PromotionName
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniResourcePromotionsourceNotifyModel struct {
	value *AlipayOpenMiniResourcePromotionsourceNotifyModel
	isSet bool
}

func (v NullableAlipayOpenMiniResourcePromotionsourceNotifyModel) Get() *AlipayOpenMiniResourcePromotionsourceNotifyModel {
	return v.value
}

func (v *NullableAlipayOpenMiniResourcePromotionsourceNotifyModel) Set(val *AlipayOpenMiniResourcePromotionsourceNotifyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniResourcePromotionsourceNotifyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniResourcePromotionsourceNotifyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniResourcePromotionsourceNotifyModel(val *AlipayOpenMiniResourcePromotionsourceNotifyModel) *NullableAlipayOpenMiniResourcePromotionsourceNotifyModel {
	return &NullableAlipayOpenMiniResourcePromotionsourceNotifyModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniResourcePromotionsourceNotifyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniResourcePromotionsourceNotifyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


