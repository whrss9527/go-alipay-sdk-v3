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

// checks if the ZhimaCustomerJobworthJobdataAddResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZhimaCustomerJobworthJobdataAddResponseModel{}

// ZhimaCustomerJobworthJobdataAddResponseModel struct for ZhimaCustomerJobworthJobdataAddResponseModel
type ZhimaCustomerJobworthJobdataAddResponseModel struct {
	// 受理台单号，用来调用工作证受理台
	AcceptanceId *string `json:"acceptance_id,omitempty"`
	// 错误码
	SubCode *string `json:"sub_code,omitempty"`
	// 业务中文结果信息
	SubMsg *string `json:"sub_msg,omitempty"`
	// 职得工作证跳转小程序链接
	Url *string `json:"url,omitempty"`
}

// NewZhimaCustomerJobworthJobdataAddResponseModel instantiates a new ZhimaCustomerJobworthJobdataAddResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZhimaCustomerJobworthJobdataAddResponseModel() *ZhimaCustomerJobworthJobdataAddResponseModel {
	this := ZhimaCustomerJobworthJobdataAddResponseModel{}
	return &this
}

// NewZhimaCustomerJobworthJobdataAddResponseModelWithDefaults instantiates a new ZhimaCustomerJobworthJobdataAddResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZhimaCustomerJobworthJobdataAddResponseModelWithDefaults() *ZhimaCustomerJobworthJobdataAddResponseModel {
	this := ZhimaCustomerJobworthJobdataAddResponseModel{}
	return &this
}

// GetAcceptanceId returns the AcceptanceId field value if set, zero value otherwise.
func (o *ZhimaCustomerJobworthJobdataAddResponseModel) GetAcceptanceId() string {
	if o == nil || IsNil(o.AcceptanceId) {
		var ret string
		return ret
	}
	return *o.AcceptanceId
}

// GetAcceptanceIdOk returns a tuple with the AcceptanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCustomerJobworthJobdataAddResponseModel) GetAcceptanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.AcceptanceId) {
		return nil, false
	}
	return o.AcceptanceId, true
}

// HasAcceptanceId returns a boolean if a field has been set.
func (o *ZhimaCustomerJobworthJobdataAddResponseModel) HasAcceptanceId() bool {
	if o != nil && !IsNil(o.AcceptanceId) {
		return true
	}

	return false
}

// SetAcceptanceId gets a reference to the given string and assigns it to the AcceptanceId field.
func (o *ZhimaCustomerJobworthJobdataAddResponseModel) SetAcceptanceId(v string) {
	o.AcceptanceId = &v
}

// GetSubCode returns the SubCode field value if set, zero value otherwise.
func (o *ZhimaCustomerJobworthJobdataAddResponseModel) GetSubCode() string {
	if o == nil || IsNil(o.SubCode) {
		var ret string
		return ret
	}
	return *o.SubCode
}

// GetSubCodeOk returns a tuple with the SubCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCustomerJobworthJobdataAddResponseModel) GetSubCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SubCode) {
		return nil, false
	}
	return o.SubCode, true
}

// HasSubCode returns a boolean if a field has been set.
func (o *ZhimaCustomerJobworthJobdataAddResponseModel) HasSubCode() bool {
	if o != nil && !IsNil(o.SubCode) {
		return true
	}

	return false
}

// SetSubCode gets a reference to the given string and assigns it to the SubCode field.
func (o *ZhimaCustomerJobworthJobdataAddResponseModel) SetSubCode(v string) {
	o.SubCode = &v
}

// GetSubMsg returns the SubMsg field value if set, zero value otherwise.
func (o *ZhimaCustomerJobworthJobdataAddResponseModel) GetSubMsg() string {
	if o == nil || IsNil(o.SubMsg) {
		var ret string
		return ret
	}
	return *o.SubMsg
}

// GetSubMsgOk returns a tuple with the SubMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCustomerJobworthJobdataAddResponseModel) GetSubMsgOk() (*string, bool) {
	if o == nil || IsNil(o.SubMsg) {
		return nil, false
	}
	return o.SubMsg, true
}

// HasSubMsg returns a boolean if a field has been set.
func (o *ZhimaCustomerJobworthJobdataAddResponseModel) HasSubMsg() bool {
	if o != nil && !IsNil(o.SubMsg) {
		return true
	}

	return false
}

// SetSubMsg gets a reference to the given string and assigns it to the SubMsg field.
func (o *ZhimaCustomerJobworthJobdataAddResponseModel) SetSubMsg(v string) {
	o.SubMsg = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ZhimaCustomerJobworthJobdataAddResponseModel) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCustomerJobworthJobdataAddResponseModel) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ZhimaCustomerJobworthJobdataAddResponseModel) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ZhimaCustomerJobworthJobdataAddResponseModel) SetUrl(v string) {
	o.Url = &v
}

func (o ZhimaCustomerJobworthJobdataAddResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZhimaCustomerJobworthJobdataAddResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcceptanceId) {
		toSerialize["acceptance_id"] = o.AcceptanceId
	}
	if !IsNil(o.SubCode) {
		toSerialize["sub_code"] = o.SubCode
	}
	if !IsNil(o.SubMsg) {
		toSerialize["sub_msg"] = o.SubMsg
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableZhimaCustomerJobworthJobdataAddResponseModel struct {
	value *ZhimaCustomerJobworthJobdataAddResponseModel
	isSet bool
}

func (v NullableZhimaCustomerJobworthJobdataAddResponseModel) Get() *ZhimaCustomerJobworthJobdataAddResponseModel {
	return v.value
}

func (v *NullableZhimaCustomerJobworthJobdataAddResponseModel) Set(val *ZhimaCustomerJobworthJobdataAddResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableZhimaCustomerJobworthJobdataAddResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableZhimaCustomerJobworthJobdataAddResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZhimaCustomerJobworthJobdataAddResponseModel(val *ZhimaCustomerJobworthJobdataAddResponseModel) *NullableZhimaCustomerJobworthJobdataAddResponseModel {
	return &NullableZhimaCustomerJobworthJobdataAddResponseModel{value: val, isSet: true}
}

func (v NullableZhimaCustomerJobworthJobdataAddResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZhimaCustomerJobworthJobdataAddResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

