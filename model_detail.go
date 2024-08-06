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

// checks if the Detail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Detail{}

// Detail struct for Detail
type Detail struct {
	// 英文描述（collectAttachement值为true时必填） 以下是系统固定code,对上传的图片做相应的校验。 身份证校验情况： 身份证正面照片：ESIGN_IDCARD 身份证背面照片：ESIGN_BACK_IDCARD
	Code *string `json:"code,omitempty"`
	// 请上传手机正面照
	Description *string `json:"description,omitempty"`
	// 顺序（collectAttachement值为true时必填）
	Order *int32 `json:"order,omitempty"`
}

// NewDetail instantiates a new Detail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetail() *Detail {
	this := Detail{}
	return &this
}

// NewDetailWithDefaults instantiates a new Detail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailWithDefaults() *Detail {
	this := Detail{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Detail) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Detail) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Detail) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Detail) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Detail) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Detail) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Detail) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Detail) SetDescription(v string) {
	o.Description = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *Detail) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Detail) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *Detail) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *Detail) SetOrder(v int32) {
	o.Order = &v
}

func (o Detail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Detail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	return toSerialize, nil
}

type NullableDetail struct {
	value *Detail
	isSet bool
}

func (v NullableDetail) Get() *Detail {
	return v.value
}

func (v *NullableDetail) Set(val *Detail) {
	v.value = val
	v.isSet = true
}

func (v NullableDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetail(val *Detail) *NullableDetail {
	return &NullableDetail{value: val, isSet: true}
}

func (v NullableDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
