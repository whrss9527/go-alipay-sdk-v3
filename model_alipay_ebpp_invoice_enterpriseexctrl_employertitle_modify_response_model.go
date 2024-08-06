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

// checks if the AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel{}

// AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel struct for AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel
type AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel struct {
	// 返回码 10000-成功 其他都是失败
	Code *string `json:"code,omitempty"`
	// 返回信息
	Msg *string `json:"msg,omitempty"`
	// 是否成功
	Success *bool `json:"success,omitempty"`
	// 抬头ID
	TitleId *string `json:"title_id,omitempty"`
}

// NewAlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel instantiates a new AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel() *AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel {
	this := AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel{}
	return &this
}

// NewAlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModelWithDefaults instantiates a new AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModelWithDefaults() *AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel {
	this := AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel) SetCode(v string) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel) SetMsg(v string) {
	o.Msg = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel) SetSuccess(v bool) {
	o.Success = &v
}

// GetTitleId returns the TitleId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel) GetTitleId() string {
	if o == nil || IsNil(o.TitleId) {
		var ret string
		return ret
	}
	return *o.TitleId
}

// GetTitleIdOk returns a tuple with the TitleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel) GetTitleIdOk() (*string, bool) {
	if o == nil || IsNil(o.TitleId) {
		return nil, false
	}
	return o.TitleId, true
}

// HasTitleId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel) HasTitleId() bool {
	if o != nil && !IsNil(o.TitleId) {
		return true
	}

	return false
}

// SetTitleId gets a reference to the given string and assigns it to the TitleId field.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel) SetTitleId(v string) {
	o.TitleId = &v
}

func (o AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.TitleId) {
		toSerialize["title_id"] = o.TitleId
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel struct {
	value *AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel) Get() *AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel) Set(val *AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel(val *AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel) *NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel {
	return &NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


