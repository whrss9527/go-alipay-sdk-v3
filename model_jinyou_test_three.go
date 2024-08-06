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

// checks if the JinyouTestThree type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JinyouTestThree{}

// JinyouTestThree struct for JinyouTestThree
type JinyouTestThree struct {
	Th1F *JinyouTestTwo `json:"th_1_f,omitempty"`
	// 无枚举值
	Th2N *string `json:"th_2_n,omitempty"`
	// 1
	Th3N *bool `json:"th_3_n,omitempty"`
}

// NewJinyouTestThree instantiates a new JinyouTestThree object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJinyouTestThree() *JinyouTestThree {
	this := JinyouTestThree{}
	return &this
}

// NewJinyouTestThreeWithDefaults instantiates a new JinyouTestThree object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJinyouTestThreeWithDefaults() *JinyouTestThree {
	this := JinyouTestThree{}
	return &this
}

// GetTh1F returns the Th1F field value if set, zero value otherwise.
func (o *JinyouTestThree) GetTh1F() JinyouTestTwo {
	if o == nil || IsNil(o.Th1F) {
		var ret JinyouTestTwo
		return ret
	}
	return *o.Th1F
}

// GetTh1FOk returns a tuple with the Th1F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JinyouTestThree) GetTh1FOk() (*JinyouTestTwo, bool) {
	if o == nil || IsNil(o.Th1F) {
		return nil, false
	}
	return o.Th1F, true
}

// HasTh1F returns a boolean if a field has been set.
func (o *JinyouTestThree) HasTh1F() bool {
	if o != nil && !IsNil(o.Th1F) {
		return true
	}

	return false
}

// SetTh1F gets a reference to the given JinyouTestTwo and assigns it to the Th1F field.
func (o *JinyouTestThree) SetTh1F(v JinyouTestTwo) {
	o.Th1F = &v
}

// GetTh2N returns the Th2N field value if set, zero value otherwise.
func (o *JinyouTestThree) GetTh2N() string {
	if o == nil || IsNil(o.Th2N) {
		var ret string
		return ret
	}
	return *o.Th2N
}

// GetTh2NOk returns a tuple with the Th2N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JinyouTestThree) GetTh2NOk() (*string, bool) {
	if o == nil || IsNil(o.Th2N) {
		return nil, false
	}
	return o.Th2N, true
}

// HasTh2N returns a boolean if a field has been set.
func (o *JinyouTestThree) HasTh2N() bool {
	if o != nil && !IsNil(o.Th2N) {
		return true
	}

	return false
}

// SetTh2N gets a reference to the given string and assigns it to the Th2N field.
func (o *JinyouTestThree) SetTh2N(v string) {
	o.Th2N = &v
}

// GetTh3N returns the Th3N field value if set, zero value otherwise.
func (o *JinyouTestThree) GetTh3N() bool {
	if o == nil || IsNil(o.Th3N) {
		var ret bool
		return ret
	}
	return *o.Th3N
}

// GetTh3NOk returns a tuple with the Th3N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JinyouTestThree) GetTh3NOk() (*bool, bool) {
	if o == nil || IsNil(o.Th3N) {
		return nil, false
	}
	return o.Th3N, true
}

// HasTh3N returns a boolean if a field has been set.
func (o *JinyouTestThree) HasTh3N() bool {
	if o != nil && !IsNil(o.Th3N) {
		return true
	}

	return false
}

// SetTh3N gets a reference to the given bool and assigns it to the Th3N field.
func (o *JinyouTestThree) SetTh3N(v bool) {
	o.Th3N = &v
}

func (o JinyouTestThree) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JinyouTestThree) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Th1F) {
		toSerialize["th_1_f"] = o.Th1F
	}
	if !IsNil(o.Th2N) {
		toSerialize["th_2_n"] = o.Th2N
	}
	if !IsNil(o.Th3N) {
		toSerialize["th_3_n"] = o.Th3N
	}
	return toSerialize, nil
}

type NullableJinyouTestThree struct {
	value *JinyouTestThree
	isSet bool
}

func (v NullableJinyouTestThree) Get() *JinyouTestThree {
	return v.value
}

func (v *NullableJinyouTestThree) Set(val *JinyouTestThree) {
	v.value = val
	v.isSet = true
}

func (v NullableJinyouTestThree) IsSet() bool {
	return v.isSet
}

func (v *NullableJinyouTestThree) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJinyouTestThree(val *JinyouTestThree) *NullableJinyouTestThree {
	return &NullableJinyouTestThree{value: val, isSet: true}
}

func (v NullableJinyouTestThree) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJinyouTestThree) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

