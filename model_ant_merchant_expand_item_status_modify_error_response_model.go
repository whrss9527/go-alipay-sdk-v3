/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alipay

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AntMerchantExpandItemStatusModifyErrorResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AntMerchantExpandItemStatusModifyErrorResponseModel{}

// AntMerchantExpandItemStatusModifyErrorResponseModel struct for AntMerchantExpandItemStatusModifyErrorResponseModel
type AntMerchantExpandItemStatusModifyErrorResponseModel struct {
	// 错误码
	Code string `json:"code"`
	// 解决方案链接
	Links *string `json:"links,omitempty"`
	// 错误描述
	Message string `json:"message"`
}

type _AntMerchantExpandItemStatusModifyErrorResponseModel AntMerchantExpandItemStatusModifyErrorResponseModel

// NewAntMerchantExpandItemStatusModifyErrorResponseModel instantiates a new AntMerchantExpandItemStatusModifyErrorResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAntMerchantExpandItemStatusModifyErrorResponseModel(code string, message string) *AntMerchantExpandItemStatusModifyErrorResponseModel {
	this := AntMerchantExpandItemStatusModifyErrorResponseModel{}
	this.Code = code
	this.Message = message
	return &this
}

// NewAntMerchantExpandItemStatusModifyErrorResponseModelWithDefaults instantiates a new AntMerchantExpandItemStatusModifyErrorResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAntMerchantExpandItemStatusModifyErrorResponseModelWithDefaults() *AntMerchantExpandItemStatusModifyErrorResponseModel {
	this := AntMerchantExpandItemStatusModifyErrorResponseModel{}
	return &this
}

// GetCode returns the Code field value
func (o *AntMerchantExpandItemStatusModifyErrorResponseModel) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandItemStatusModifyErrorResponseModel) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *AntMerchantExpandItemStatusModifyErrorResponseModel) SetCode(v string) {
	o.Code = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AntMerchantExpandItemStatusModifyErrorResponseModel) GetLinks() string {
	if o == nil || IsNil(o.Links) {
		var ret string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandItemStatusModifyErrorResponseModel) GetLinksOk() (*string, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AntMerchantExpandItemStatusModifyErrorResponseModel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given string and assigns it to the Links field.
func (o *AntMerchantExpandItemStatusModifyErrorResponseModel) SetLinks(v string) {
	o.Links = &v
}

// GetMessage returns the Message field value
func (o *AntMerchantExpandItemStatusModifyErrorResponseModel) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandItemStatusModifyErrorResponseModel) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AntMerchantExpandItemStatusModifyErrorResponseModel) SetMessage(v string) {
	o.Message = v
}

func (o AntMerchantExpandItemStatusModifyErrorResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AntMerchantExpandItemStatusModifyErrorResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *AntMerchantExpandItemStatusModifyErrorResponseModel) UnmarshalJSON(data []byte) (err error) {
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
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAntMerchantExpandItemStatusModifyErrorResponseModel := _AntMerchantExpandItemStatusModifyErrorResponseModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAntMerchantExpandItemStatusModifyErrorResponseModel)

	if err != nil {
		return err
	}

	*o = AntMerchantExpandItemStatusModifyErrorResponseModel(varAntMerchantExpandItemStatusModifyErrorResponseModel)

	return err
}

type NullableAntMerchantExpandItemStatusModifyErrorResponseModel struct {
	value *AntMerchantExpandItemStatusModifyErrorResponseModel
	isSet bool
}

func (v NullableAntMerchantExpandItemStatusModifyErrorResponseModel) Get() *AntMerchantExpandItemStatusModifyErrorResponseModel {
	return v.value
}

func (v *NullableAntMerchantExpandItemStatusModifyErrorResponseModel) Set(val *AntMerchantExpandItemStatusModifyErrorResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAntMerchantExpandItemStatusModifyErrorResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAntMerchantExpandItemStatusModifyErrorResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntMerchantExpandItemStatusModifyErrorResponseModel(val *AntMerchantExpandItemStatusModifyErrorResponseModel) *NullableAntMerchantExpandItemStatusModifyErrorResponseModel {
	return &NullableAntMerchantExpandItemStatusModifyErrorResponseModel{value: val, isSet: true}
}

func (v NullableAntMerchantExpandItemStatusModifyErrorResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntMerchantExpandItemStatusModifyErrorResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
