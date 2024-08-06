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

// checks if the AlipayIserviceCcmServicerecordModifyErrorResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayIserviceCcmServicerecordModifyErrorResponseModel{}

// AlipayIserviceCcmServicerecordModifyErrorResponseModel struct for AlipayIserviceCcmServicerecordModifyErrorResponseModel
type AlipayIserviceCcmServicerecordModifyErrorResponseModel struct {
	// 错误码
	Code string `json:"code"`
	// 解决方案链接
	Links *string `json:"links,omitempty"`
	// 错误描述
	Message string `json:"message"`
}

type _AlipayIserviceCcmServicerecordModifyErrorResponseModel AlipayIserviceCcmServicerecordModifyErrorResponseModel

// NewAlipayIserviceCcmServicerecordModifyErrorResponseModel instantiates a new AlipayIserviceCcmServicerecordModifyErrorResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayIserviceCcmServicerecordModifyErrorResponseModel(code string, message string) *AlipayIserviceCcmServicerecordModifyErrorResponseModel {
	this := AlipayIserviceCcmServicerecordModifyErrorResponseModel{}
	this.Code = code
	this.Message = message
	return &this
}

// NewAlipayIserviceCcmServicerecordModifyErrorResponseModelWithDefaults instantiates a new AlipayIserviceCcmServicerecordModifyErrorResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayIserviceCcmServicerecordModifyErrorResponseModelWithDefaults() *AlipayIserviceCcmServicerecordModifyErrorResponseModel {
	this := AlipayIserviceCcmServicerecordModifyErrorResponseModel{}
	return &this
}

// GetCode returns the Code field value
func (o *AlipayIserviceCcmServicerecordModifyErrorResponseModel) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmServicerecordModifyErrorResponseModel) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *AlipayIserviceCcmServicerecordModifyErrorResponseModel) SetCode(v string) {
	o.Code = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AlipayIserviceCcmServicerecordModifyErrorResponseModel) GetLinks() string {
	if o == nil || IsNil(o.Links) {
		var ret string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmServicerecordModifyErrorResponseModel) GetLinksOk() (*string, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AlipayIserviceCcmServicerecordModifyErrorResponseModel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given string and assigns it to the Links field.
func (o *AlipayIserviceCcmServicerecordModifyErrorResponseModel) SetLinks(v string) {
	o.Links = &v
}

// GetMessage returns the Message field value
func (o *AlipayIserviceCcmServicerecordModifyErrorResponseModel) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmServicerecordModifyErrorResponseModel) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AlipayIserviceCcmServicerecordModifyErrorResponseModel) SetMessage(v string) {
	o.Message = v
}

func (o AlipayIserviceCcmServicerecordModifyErrorResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayIserviceCcmServicerecordModifyErrorResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *AlipayIserviceCcmServicerecordModifyErrorResponseModel) UnmarshalJSON(data []byte) (err error) {
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

	varAlipayIserviceCcmServicerecordModifyErrorResponseModel := _AlipayIserviceCcmServicerecordModifyErrorResponseModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlipayIserviceCcmServicerecordModifyErrorResponseModel)

	if err != nil {
		return err
	}

	*o = AlipayIserviceCcmServicerecordModifyErrorResponseModel(varAlipayIserviceCcmServicerecordModifyErrorResponseModel)

	return err
}

type NullableAlipayIserviceCcmServicerecordModifyErrorResponseModel struct {
	value *AlipayIserviceCcmServicerecordModifyErrorResponseModel
	isSet bool
}

func (v NullableAlipayIserviceCcmServicerecordModifyErrorResponseModel) Get() *AlipayIserviceCcmServicerecordModifyErrorResponseModel {
	return v.value
}

func (v *NullableAlipayIserviceCcmServicerecordModifyErrorResponseModel) Set(val *AlipayIserviceCcmServicerecordModifyErrorResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayIserviceCcmServicerecordModifyErrorResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayIserviceCcmServicerecordModifyErrorResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayIserviceCcmServicerecordModifyErrorResponseModel(val *AlipayIserviceCcmServicerecordModifyErrorResponseModel) *NullableAlipayIserviceCcmServicerecordModifyErrorResponseModel {
	return &NullableAlipayIserviceCcmServicerecordModifyErrorResponseModel{value: val, isSet: true}
}

func (v NullableAlipayIserviceCcmServicerecordModifyErrorResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayIserviceCcmServicerecordModifyErrorResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


