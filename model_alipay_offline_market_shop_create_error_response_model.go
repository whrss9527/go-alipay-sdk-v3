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

// checks if the AlipayOfflineMarketShopCreateErrorResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOfflineMarketShopCreateErrorResponseModel{}

// AlipayOfflineMarketShopCreateErrorResponseModel struct for AlipayOfflineMarketShopCreateErrorResponseModel
type AlipayOfflineMarketShopCreateErrorResponseModel struct {
	// 错误码
	Code string `json:"code"`
	// 解决方案链接
	Links *string `json:"links,omitempty"`
	// 错误描述
	Message string `json:"message"`
}

type _AlipayOfflineMarketShopCreateErrorResponseModel AlipayOfflineMarketShopCreateErrorResponseModel

// NewAlipayOfflineMarketShopCreateErrorResponseModel instantiates a new AlipayOfflineMarketShopCreateErrorResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOfflineMarketShopCreateErrorResponseModel(code string, message string) *AlipayOfflineMarketShopCreateErrorResponseModel {
	this := AlipayOfflineMarketShopCreateErrorResponseModel{}
	this.Code = code
	this.Message = message
	return &this
}

// NewAlipayOfflineMarketShopCreateErrorResponseModelWithDefaults instantiates a new AlipayOfflineMarketShopCreateErrorResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOfflineMarketShopCreateErrorResponseModelWithDefaults() *AlipayOfflineMarketShopCreateErrorResponseModel {
	this := AlipayOfflineMarketShopCreateErrorResponseModel{}
	return &this
}

// GetCode returns the Code field value
func (o *AlipayOfflineMarketShopCreateErrorResponseModel) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AlipayOfflineMarketShopCreateErrorResponseModel) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *AlipayOfflineMarketShopCreateErrorResponseModel) SetCode(v string) {
	o.Code = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AlipayOfflineMarketShopCreateErrorResponseModel) GetLinks() string {
	if o == nil || IsNil(o.Links) {
		var ret string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOfflineMarketShopCreateErrorResponseModel) GetLinksOk() (*string, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AlipayOfflineMarketShopCreateErrorResponseModel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given string and assigns it to the Links field.
func (o *AlipayOfflineMarketShopCreateErrorResponseModel) SetLinks(v string) {
	o.Links = &v
}

// GetMessage returns the Message field value
func (o *AlipayOfflineMarketShopCreateErrorResponseModel) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AlipayOfflineMarketShopCreateErrorResponseModel) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AlipayOfflineMarketShopCreateErrorResponseModel) SetMessage(v string) {
	o.Message = v
}

func (o AlipayOfflineMarketShopCreateErrorResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOfflineMarketShopCreateErrorResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *AlipayOfflineMarketShopCreateErrorResponseModel) UnmarshalJSON(data []byte) (err error) {
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

	varAlipayOfflineMarketShopCreateErrorResponseModel := _AlipayOfflineMarketShopCreateErrorResponseModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlipayOfflineMarketShopCreateErrorResponseModel)

	if err != nil {
		return err
	}

	*o = AlipayOfflineMarketShopCreateErrorResponseModel(varAlipayOfflineMarketShopCreateErrorResponseModel)

	return err
}

type NullableAlipayOfflineMarketShopCreateErrorResponseModel struct {
	value *AlipayOfflineMarketShopCreateErrorResponseModel
	isSet bool
}

func (v NullableAlipayOfflineMarketShopCreateErrorResponseModel) Get() *AlipayOfflineMarketShopCreateErrorResponseModel {
	return v.value
}

func (v *NullableAlipayOfflineMarketShopCreateErrorResponseModel) Set(val *AlipayOfflineMarketShopCreateErrorResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOfflineMarketShopCreateErrorResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOfflineMarketShopCreateErrorResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOfflineMarketShopCreateErrorResponseModel(val *AlipayOfflineMarketShopCreateErrorResponseModel) *NullableAlipayOfflineMarketShopCreateErrorResponseModel {
	return &NullableAlipayOfflineMarketShopCreateErrorResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOfflineMarketShopCreateErrorResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOfflineMarketShopCreateErrorResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

