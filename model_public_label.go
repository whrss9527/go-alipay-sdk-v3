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

// checks if the PublicLabel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicLabel{}

// PublicLabel struct for PublicLabel
type PublicLabel struct {
	// 标签用户量
	Count *int32 `json:"count,omitempty"`
	// 标签编号
	Id *string `json:"id,omitempty"`
	// 标签名称
	Name *string `json:"name,omitempty"`
}

// NewPublicLabel instantiates a new PublicLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicLabel() *PublicLabel {
	this := PublicLabel{}
	return &this
}

// NewPublicLabelWithDefaults instantiates a new PublicLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicLabelWithDefaults() *PublicLabel {
	this := PublicLabel{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PublicLabel) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicLabel) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PublicLabel) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PublicLabel) SetCount(v int32) {
	o.Count = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PublicLabel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicLabel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PublicLabel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PublicLabel) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PublicLabel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicLabel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PublicLabel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PublicLabel) SetName(v string) {
	o.Name = &v
}

func (o PublicLabel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicLabel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullablePublicLabel struct {
	value *PublicLabel
	isSet bool
}

func (v NullablePublicLabel) Get() *PublicLabel {
	return v.value
}

func (v *NullablePublicLabel) Set(val *PublicLabel) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicLabel) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicLabel(val *PublicLabel) *NullablePublicLabel {
	return &NullablePublicLabel{value: val, isSet: true}
}

func (v NullablePublicLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
