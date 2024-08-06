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

// checks if the SignMerchantParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignMerchantParams{}

// SignMerchantParams struct for SignMerchantParams
type SignMerchantParams struct {
	// 子商户的商户id
	SubMerchantId *string `json:"sub_merchant_id,omitempty"`
	// 子商户的商户名称
	SubMerchantName *string `json:"sub_merchant_name,omitempty"`
	// 子商户的服务描述
	SubMerchantServiceDescription *string `json:"sub_merchant_service_description,omitempty"`
	// 子商户的服务名称
	SubMerchantServiceName *string `json:"sub_merchant_service_name,omitempty"`
}

// NewSignMerchantParams instantiates a new SignMerchantParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignMerchantParams() *SignMerchantParams {
	this := SignMerchantParams{}
	return &this
}

// NewSignMerchantParamsWithDefaults instantiates a new SignMerchantParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignMerchantParamsWithDefaults() *SignMerchantParams {
	this := SignMerchantParams{}
	return &this
}

// GetSubMerchantId returns the SubMerchantId field value if set, zero value otherwise.
func (o *SignMerchantParams) GetSubMerchantId() string {
	if o == nil || IsNil(o.SubMerchantId) {
		var ret string
		return ret
	}
	return *o.SubMerchantId
}

// GetSubMerchantIdOk returns a tuple with the SubMerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignMerchantParams) GetSubMerchantIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubMerchantId) {
		return nil, false
	}
	return o.SubMerchantId, true
}

// HasSubMerchantId returns a boolean if a field has been set.
func (o *SignMerchantParams) HasSubMerchantId() bool {
	if o != nil && !IsNil(o.SubMerchantId) {
		return true
	}

	return false
}

// SetSubMerchantId gets a reference to the given string and assigns it to the SubMerchantId field.
func (o *SignMerchantParams) SetSubMerchantId(v string) {
	o.SubMerchantId = &v
}

// GetSubMerchantName returns the SubMerchantName field value if set, zero value otherwise.
func (o *SignMerchantParams) GetSubMerchantName() string {
	if o == nil || IsNil(o.SubMerchantName) {
		var ret string
		return ret
	}
	return *o.SubMerchantName
}

// GetSubMerchantNameOk returns a tuple with the SubMerchantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignMerchantParams) GetSubMerchantNameOk() (*string, bool) {
	if o == nil || IsNil(o.SubMerchantName) {
		return nil, false
	}
	return o.SubMerchantName, true
}

// HasSubMerchantName returns a boolean if a field has been set.
func (o *SignMerchantParams) HasSubMerchantName() bool {
	if o != nil && !IsNil(o.SubMerchantName) {
		return true
	}

	return false
}

// SetSubMerchantName gets a reference to the given string and assigns it to the SubMerchantName field.
func (o *SignMerchantParams) SetSubMerchantName(v string) {
	o.SubMerchantName = &v
}

// GetSubMerchantServiceDescription returns the SubMerchantServiceDescription field value if set, zero value otherwise.
func (o *SignMerchantParams) GetSubMerchantServiceDescription() string {
	if o == nil || IsNil(o.SubMerchantServiceDescription) {
		var ret string
		return ret
	}
	return *o.SubMerchantServiceDescription
}

// GetSubMerchantServiceDescriptionOk returns a tuple with the SubMerchantServiceDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignMerchantParams) GetSubMerchantServiceDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.SubMerchantServiceDescription) {
		return nil, false
	}
	return o.SubMerchantServiceDescription, true
}

// HasSubMerchantServiceDescription returns a boolean if a field has been set.
func (o *SignMerchantParams) HasSubMerchantServiceDescription() bool {
	if o != nil && !IsNil(o.SubMerchantServiceDescription) {
		return true
	}

	return false
}

// SetSubMerchantServiceDescription gets a reference to the given string and assigns it to the SubMerchantServiceDescription field.
func (o *SignMerchantParams) SetSubMerchantServiceDescription(v string) {
	o.SubMerchantServiceDescription = &v
}

// GetSubMerchantServiceName returns the SubMerchantServiceName field value if set, zero value otherwise.
func (o *SignMerchantParams) GetSubMerchantServiceName() string {
	if o == nil || IsNil(o.SubMerchantServiceName) {
		var ret string
		return ret
	}
	return *o.SubMerchantServiceName
}

// GetSubMerchantServiceNameOk returns a tuple with the SubMerchantServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignMerchantParams) GetSubMerchantServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.SubMerchantServiceName) {
		return nil, false
	}
	return o.SubMerchantServiceName, true
}

// HasSubMerchantServiceName returns a boolean if a field has been set.
func (o *SignMerchantParams) HasSubMerchantServiceName() bool {
	if o != nil && !IsNil(o.SubMerchantServiceName) {
		return true
	}

	return false
}

// SetSubMerchantServiceName gets a reference to the given string and assigns it to the SubMerchantServiceName field.
func (o *SignMerchantParams) SetSubMerchantServiceName(v string) {
	o.SubMerchantServiceName = &v
}

func (o SignMerchantParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignMerchantParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubMerchantId) {
		toSerialize["sub_merchant_id"] = o.SubMerchantId
	}
	if !IsNil(o.SubMerchantName) {
		toSerialize["sub_merchant_name"] = o.SubMerchantName
	}
	if !IsNil(o.SubMerchantServiceDescription) {
		toSerialize["sub_merchant_service_description"] = o.SubMerchantServiceDescription
	}
	if !IsNil(o.SubMerchantServiceName) {
		toSerialize["sub_merchant_service_name"] = o.SubMerchantServiceName
	}
	return toSerialize, nil
}

type NullableSignMerchantParams struct {
	value *SignMerchantParams
	isSet bool
}

func (v NullableSignMerchantParams) Get() *SignMerchantParams {
	return v.value
}

func (v *NullableSignMerchantParams) Set(val *SignMerchantParams) {
	v.value = val
	v.isSet = true
}

func (v NullableSignMerchantParams) IsSet() bool {
	return v.isSet
}

func (v *NullableSignMerchantParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignMerchantParams(val *SignMerchantParams) *NullableSignMerchantParams {
	return &NullableSignMerchantParams{value: val, isSet: true}
}

func (v NullableSignMerchantParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignMerchantParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
