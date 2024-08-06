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

// checks if the AlipayOpenMiniTemplateMarketingCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniTemplateMarketingCreateModel{}

// AlipayOpenMiniTemplateMarketingCreateModel struct for AlipayOpenMiniTemplateMarketingCreateModel
type AlipayOpenMiniTemplateMarketingCreateModel struct {
	// 营销活动id
	ActivityId *string `json:"activity_id,omitempty"`
	// 活动结束时间
	GmtEnd *string `json:"gmt_end,omitempty"`
	// 活动开始时间
	GmtStart *string `json:"gmt_start,omitempty"`
	// 消息模板id列表，最多50个模板id
	TemplateIds []string `json:"template_ids,omitempty"`
	// 消息运营位名称，不填默认使用券名称
	Title *string `json:"title,omitempty"`
}

// NewAlipayOpenMiniTemplateMarketingCreateModel instantiates a new AlipayOpenMiniTemplateMarketingCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniTemplateMarketingCreateModel() *AlipayOpenMiniTemplateMarketingCreateModel {
	this := AlipayOpenMiniTemplateMarketingCreateModel{}
	return &this
}

// NewAlipayOpenMiniTemplateMarketingCreateModelWithDefaults instantiates a new AlipayOpenMiniTemplateMarketingCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniTemplateMarketingCreateModelWithDefaults() *AlipayOpenMiniTemplateMarketingCreateModel {
	this := AlipayOpenMiniTemplateMarketingCreateModel{}
	return &this
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise.
func (o *AlipayOpenMiniTemplateMarketingCreateModel) GetActivityId() string {
	if o == nil || IsNil(o.ActivityId) {
		var ret string
		return ret
	}
	return *o.ActivityId
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniTemplateMarketingCreateModel) GetActivityIdOk() (*string, bool) {
	if o == nil || IsNil(o.ActivityId) {
		return nil, false
	}
	return o.ActivityId, true
}

// HasActivityId returns a boolean if a field has been set.
func (o *AlipayOpenMiniTemplateMarketingCreateModel) HasActivityId() bool {
	if o != nil && !IsNil(o.ActivityId) {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given string and assigns it to the ActivityId field.
func (o *AlipayOpenMiniTemplateMarketingCreateModel) SetActivityId(v string) {
	o.ActivityId = &v
}

// GetGmtEnd returns the GmtEnd field value if set, zero value otherwise.
func (o *AlipayOpenMiniTemplateMarketingCreateModel) GetGmtEnd() string {
	if o == nil || IsNil(o.GmtEnd) {
		var ret string
		return ret
	}
	return *o.GmtEnd
}

// GetGmtEndOk returns a tuple with the GmtEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniTemplateMarketingCreateModel) GetGmtEndOk() (*string, bool) {
	if o == nil || IsNil(o.GmtEnd) {
		return nil, false
	}
	return o.GmtEnd, true
}

// HasGmtEnd returns a boolean if a field has been set.
func (o *AlipayOpenMiniTemplateMarketingCreateModel) HasGmtEnd() bool {
	if o != nil && !IsNil(o.GmtEnd) {
		return true
	}

	return false
}

// SetGmtEnd gets a reference to the given string and assigns it to the GmtEnd field.
func (o *AlipayOpenMiniTemplateMarketingCreateModel) SetGmtEnd(v string) {
	o.GmtEnd = &v
}

// GetGmtStart returns the GmtStart field value if set, zero value otherwise.
func (o *AlipayOpenMiniTemplateMarketingCreateModel) GetGmtStart() string {
	if o == nil || IsNil(o.GmtStart) {
		var ret string
		return ret
	}
	return *o.GmtStart
}

// GetGmtStartOk returns a tuple with the GmtStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniTemplateMarketingCreateModel) GetGmtStartOk() (*string, bool) {
	if o == nil || IsNil(o.GmtStart) {
		return nil, false
	}
	return o.GmtStart, true
}

// HasGmtStart returns a boolean if a field has been set.
func (o *AlipayOpenMiniTemplateMarketingCreateModel) HasGmtStart() bool {
	if o != nil && !IsNil(o.GmtStart) {
		return true
	}

	return false
}

// SetGmtStart gets a reference to the given string and assigns it to the GmtStart field.
func (o *AlipayOpenMiniTemplateMarketingCreateModel) SetGmtStart(v string) {
	o.GmtStart = &v
}

// GetTemplateIds returns the TemplateIds field value if set, zero value otherwise.
func (o *AlipayOpenMiniTemplateMarketingCreateModel) GetTemplateIds() []string {
	if o == nil || IsNil(o.TemplateIds) {
		var ret []string
		return ret
	}
	return o.TemplateIds
}

// GetTemplateIdsOk returns a tuple with the TemplateIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniTemplateMarketingCreateModel) GetTemplateIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TemplateIds) {
		return nil, false
	}
	return o.TemplateIds, true
}

// HasTemplateIds returns a boolean if a field has been set.
func (o *AlipayOpenMiniTemplateMarketingCreateModel) HasTemplateIds() bool {
	if o != nil && !IsNil(o.TemplateIds) {
		return true
	}

	return false
}

// SetTemplateIds gets a reference to the given []string and assigns it to the TemplateIds field.
func (o *AlipayOpenMiniTemplateMarketingCreateModel) SetTemplateIds(v []string) {
	o.TemplateIds = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AlipayOpenMiniTemplateMarketingCreateModel) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniTemplateMarketingCreateModel) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AlipayOpenMiniTemplateMarketingCreateModel) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AlipayOpenMiniTemplateMarketingCreateModel) SetTitle(v string) {
	o.Title = &v
}

func (o AlipayOpenMiniTemplateMarketingCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniTemplateMarketingCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivityId) {
		toSerialize["activity_id"] = o.ActivityId
	}
	if !IsNil(o.GmtEnd) {
		toSerialize["gmt_end"] = o.GmtEnd
	}
	if !IsNil(o.GmtStart) {
		toSerialize["gmt_start"] = o.GmtStart
	}
	if !IsNil(o.TemplateIds) {
		toSerialize["template_ids"] = o.TemplateIds
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniTemplateMarketingCreateModel struct {
	value *AlipayOpenMiniTemplateMarketingCreateModel
	isSet bool
}

func (v NullableAlipayOpenMiniTemplateMarketingCreateModel) Get() *AlipayOpenMiniTemplateMarketingCreateModel {
	return v.value
}

func (v *NullableAlipayOpenMiniTemplateMarketingCreateModel) Set(val *AlipayOpenMiniTemplateMarketingCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniTemplateMarketingCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniTemplateMarketingCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniTemplateMarketingCreateModel(val *AlipayOpenMiniTemplateMarketingCreateModel) *NullableAlipayOpenMiniTemplateMarketingCreateModel {
	return &NullableAlipayOpenMiniTemplateMarketingCreateModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniTemplateMarketingCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniTemplateMarketingCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


