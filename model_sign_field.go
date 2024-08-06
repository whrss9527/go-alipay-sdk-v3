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

// checks if the SignField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignField{}

// SignField struct for SignField
type SignField struct {
	// 平台自动签
	AutoExecute *string `json:"auto_execute,omitempty"`
	Signer      *Signer `json:"signer,omitempty"`
	// 模板组件id
	StructKey *string `json:"struct_key,omitempty"`
}

// NewSignField instantiates a new SignField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignField() *SignField {
	this := SignField{}
	return &this
}

// NewSignFieldWithDefaults instantiates a new SignField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignFieldWithDefaults() *SignField {
	this := SignField{}
	return &this
}

// GetAutoExecute returns the AutoExecute field value if set, zero value otherwise.
func (o *SignField) GetAutoExecute() string {
	if o == nil || IsNil(o.AutoExecute) {
		var ret string
		return ret
	}
	return *o.AutoExecute
}

// GetAutoExecuteOk returns a tuple with the AutoExecute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignField) GetAutoExecuteOk() (*string, bool) {
	if o == nil || IsNil(o.AutoExecute) {
		return nil, false
	}
	return o.AutoExecute, true
}

// HasAutoExecute returns a boolean if a field has been set.
func (o *SignField) HasAutoExecute() bool {
	if o != nil && !IsNil(o.AutoExecute) {
		return true
	}

	return false
}

// SetAutoExecute gets a reference to the given string and assigns it to the AutoExecute field.
func (o *SignField) SetAutoExecute(v string) {
	o.AutoExecute = &v
}

// GetSigner returns the Signer field value if set, zero value otherwise.
func (o *SignField) GetSigner() Signer {
	if o == nil || IsNil(o.Signer) {
		var ret Signer
		return ret
	}
	return *o.Signer
}

// GetSignerOk returns a tuple with the Signer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignField) GetSignerOk() (*Signer, bool) {
	if o == nil || IsNil(o.Signer) {
		return nil, false
	}
	return o.Signer, true
}

// HasSigner returns a boolean if a field has been set.
func (o *SignField) HasSigner() bool {
	if o != nil && !IsNil(o.Signer) {
		return true
	}

	return false
}

// SetSigner gets a reference to the given Signer and assigns it to the Signer field.
func (o *SignField) SetSigner(v Signer) {
	o.Signer = &v
}

// GetStructKey returns the StructKey field value if set, zero value otherwise.
func (o *SignField) GetStructKey() string {
	if o == nil || IsNil(o.StructKey) {
		var ret string
		return ret
	}
	return *o.StructKey
}

// GetStructKeyOk returns a tuple with the StructKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignField) GetStructKeyOk() (*string, bool) {
	if o == nil || IsNil(o.StructKey) {
		return nil, false
	}
	return o.StructKey, true
}

// HasStructKey returns a boolean if a field has been set.
func (o *SignField) HasStructKey() bool {
	if o != nil && !IsNil(o.StructKey) {
		return true
	}

	return false
}

// SetStructKey gets a reference to the given string and assigns it to the StructKey field.
func (o *SignField) SetStructKey(v string) {
	o.StructKey = &v
}

func (o SignField) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoExecute) {
		toSerialize["auto_execute"] = o.AutoExecute
	}
	if !IsNil(o.Signer) {
		toSerialize["signer"] = o.Signer
	}
	if !IsNil(o.StructKey) {
		toSerialize["struct_key"] = o.StructKey
	}
	return toSerialize, nil
}

type NullableSignField struct {
	value *SignField
	isSet bool
}

func (v NullableSignField) Get() *SignField {
	return v.value
}

func (v *NullableSignField) Set(val *SignField) {
	v.value = val
	v.isSet = true
}

func (v NullableSignField) IsSet() bool {
	return v.isSet
}

func (v *NullableSignField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignField(val *SignField) *NullableSignField {
	return &NullableSignField{value: val, isSet: true}
}

func (v NullableSignField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
