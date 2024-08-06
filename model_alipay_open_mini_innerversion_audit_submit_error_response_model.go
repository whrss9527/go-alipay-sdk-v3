/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel{}

// AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel struct for AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel
type AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel struct {
	// 错误码
	Code string `json:"code"`
	// 解决方案链接
	Links *string `json:"links,omitempty"`
	// 错误描述
	Message string `json:"message"`
}

type _AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel

// NewAlipayOpenMiniInnerversionAuditSubmitErrorResponseModel instantiates a new AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniInnerversionAuditSubmitErrorResponseModel(code string, message string) *AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel {
	this := AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel{}
	this.Code = code
	this.Message = message
	return &this
}

// NewAlipayOpenMiniInnerversionAuditSubmitErrorResponseModelWithDefaults instantiates a new AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniInnerversionAuditSubmitErrorResponseModelWithDefaults() *AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel {
	this := AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel{}
	return &this
}

// GetCode returns the Code field value
func (o *AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel) SetCode(v string) {
	o.Code = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel) GetLinks() string {
	if o == nil || IsNil(o.Links) {
		var ret string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel) GetLinksOk() (*string, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given string and assigns it to the Links field.
func (o *AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel) SetLinks(v string) {
	o.Links = &v
}

// GetMessage returns the Message field value
func (o *AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel) SetMessage(v string) {
	o.Message = v
}

func (o AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"message",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAlipayOpenMiniInnerversionAuditSubmitErrorResponseModel := _AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlipayOpenMiniInnerversionAuditSubmitErrorResponseModel)

	if err != nil {
		return err
	}

	*o = AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel(varAlipayOpenMiniInnerversionAuditSubmitErrorResponseModel)

	return err
}

type NullableAlipayOpenMiniInnerversionAuditSubmitErrorResponseModel struct {
	value *AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel
	isSet bool
}

func (v NullableAlipayOpenMiniInnerversionAuditSubmitErrorResponseModel) Get() *AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerversionAuditSubmitErrorResponseModel) Set(val *AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerversionAuditSubmitErrorResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerversionAuditSubmitErrorResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerversionAuditSubmitErrorResponseModel(val *AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel) *NullableAlipayOpenMiniInnerversionAuditSubmitErrorResponseModel {
	return &NullableAlipayOpenMiniInnerversionAuditSubmitErrorResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerversionAuditSubmitErrorResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerversionAuditSubmitErrorResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


