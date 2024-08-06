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

// checks if the AlipayOpenPublicLifeLabelModifyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicLifeLabelModifyModel{}

// AlipayOpenPublicLifeLabelModifyModel struct for AlipayOpenPublicLifeLabelModifyModel
type AlipayOpenPublicLifeLabelModifyModel struct {
	// 标签 id，只支持生活号自定义标签。通过 <a href=\"https://opendocs.alipay.com/apis/api_6/alipay.open.public.life.label.create\">alipay.open.public.life.label.create</a>(创建标签接口)创建自定义标签后获取。
	LabelId *string `json:"label_id,omitempty"`
	// 标签名
	LabelName *string `json:"label_name,omitempty"`
}

// NewAlipayOpenPublicLifeLabelModifyModel instantiates a new AlipayOpenPublicLifeLabelModifyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicLifeLabelModifyModel() *AlipayOpenPublicLifeLabelModifyModel {
	this := AlipayOpenPublicLifeLabelModifyModel{}
	return &this
}

// NewAlipayOpenPublicLifeLabelModifyModelWithDefaults instantiates a new AlipayOpenPublicLifeLabelModifyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicLifeLabelModifyModelWithDefaults() *AlipayOpenPublicLifeLabelModifyModel {
	this := AlipayOpenPublicLifeLabelModifyModel{}
	return &this
}

// GetLabelId returns the LabelId field value if set, zero value otherwise.
func (o *AlipayOpenPublicLifeLabelModifyModel) GetLabelId() string {
	if o == nil || IsNil(o.LabelId) {
		var ret string
		return ret
	}
	return *o.LabelId
}

// GetLabelIdOk returns a tuple with the LabelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicLifeLabelModifyModel) GetLabelIdOk() (*string, bool) {
	if o == nil || IsNil(o.LabelId) {
		return nil, false
	}
	return o.LabelId, true
}

// HasLabelId returns a boolean if a field has been set.
func (o *AlipayOpenPublicLifeLabelModifyModel) HasLabelId() bool {
	if o != nil && !IsNil(o.LabelId) {
		return true
	}

	return false
}

// SetLabelId gets a reference to the given string and assigns it to the LabelId field.
func (o *AlipayOpenPublicLifeLabelModifyModel) SetLabelId(v string) {
	o.LabelId = &v
}

// GetLabelName returns the LabelName field value if set, zero value otherwise.
func (o *AlipayOpenPublicLifeLabelModifyModel) GetLabelName() string {
	if o == nil || IsNil(o.LabelName) {
		var ret string
		return ret
	}
	return *o.LabelName
}

// GetLabelNameOk returns a tuple with the LabelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicLifeLabelModifyModel) GetLabelNameOk() (*string, bool) {
	if o == nil || IsNil(o.LabelName) {
		return nil, false
	}
	return o.LabelName, true
}

// HasLabelName returns a boolean if a field has been set.
func (o *AlipayOpenPublicLifeLabelModifyModel) HasLabelName() bool {
	if o != nil && !IsNil(o.LabelName) {
		return true
	}

	return false
}

// SetLabelName gets a reference to the given string and assigns it to the LabelName field.
func (o *AlipayOpenPublicLifeLabelModifyModel) SetLabelName(v string) {
	o.LabelName = &v
}

func (o AlipayOpenPublicLifeLabelModifyModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicLifeLabelModifyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LabelId) {
		toSerialize["label_id"] = o.LabelId
	}
	if !IsNil(o.LabelName) {
		toSerialize["label_name"] = o.LabelName
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicLifeLabelModifyModel struct {
	value *AlipayOpenPublicLifeLabelModifyModel
	isSet bool
}

func (v NullableAlipayOpenPublicLifeLabelModifyModel) Get() *AlipayOpenPublicLifeLabelModifyModel {
	return v.value
}

func (v *NullableAlipayOpenPublicLifeLabelModifyModel) Set(val *AlipayOpenPublicLifeLabelModifyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicLifeLabelModifyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicLifeLabelModifyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicLifeLabelModifyModel(val *AlipayOpenPublicLifeLabelModifyModel) *NullableAlipayOpenPublicLifeLabelModifyModel {
	return &NullableAlipayOpenPublicLifeLabelModifyModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicLifeLabelModifyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicLifeLabelModifyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


