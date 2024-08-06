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

// checks if the JinyoutestopenidThree type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JinyoutestopenidThree{}

// JinyoutestopenidThree struct for JinyoutestopenidThree
type JinyoutestopenidThree struct {
	// 12
	A *string `json:"a,omitempty"`
	// 有openid,无枚举，有注意事项
	B *string `json:"b,omitempty"`
	// 有openid,无枚举，有注意事项
	BOpenId *string `json:"b_open_id,omitempty"`
	F *JinyoutestopenidTwo `json:"f,omitempty"`
}

// NewJinyoutestopenidThree instantiates a new JinyoutestopenidThree object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJinyoutestopenidThree() *JinyoutestopenidThree {
	this := JinyoutestopenidThree{}
	return &this
}

// NewJinyoutestopenidThreeWithDefaults instantiates a new JinyoutestopenidThree object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJinyoutestopenidThreeWithDefaults() *JinyoutestopenidThree {
	this := JinyoutestopenidThree{}
	return &this
}

// GetA returns the A field value if set, zero value otherwise.
func (o *JinyoutestopenidThree) GetA() string {
	if o == nil || IsNil(o.A) {
		var ret string
		return ret
	}
	return *o.A
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JinyoutestopenidThree) GetAOk() (*string, bool) {
	if o == nil || IsNil(o.A) {
		return nil, false
	}
	return o.A, true
}

// HasA returns a boolean if a field has been set.
func (o *JinyoutestopenidThree) HasA() bool {
	if o != nil && !IsNil(o.A) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *JinyoutestopenidThree) SetA(v string) {
	o.A = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *JinyoutestopenidThree) GetB() string {
	if o == nil || IsNil(o.B) {
		var ret string
		return ret
	}
	return *o.B
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JinyoutestopenidThree) GetBOk() (*string, bool) {
	if o == nil || IsNil(o.B) {
		return nil, false
	}
	return o.B, true
}

// HasB returns a boolean if a field has been set.
func (o *JinyoutestopenidThree) HasB() bool {
	if o != nil && !IsNil(o.B) {
		return true
	}

	return false
}

// SetB gets a reference to the given string and assigns it to the B field.
func (o *JinyoutestopenidThree) SetB(v string) {
	o.B = &v
}

// GetBOpenId returns the BOpenId field value if set, zero value otherwise.
func (o *JinyoutestopenidThree) GetBOpenId() string {
	if o == nil || IsNil(o.BOpenId) {
		var ret string
		return ret
	}
	return *o.BOpenId
}

// GetBOpenIdOk returns a tuple with the BOpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JinyoutestopenidThree) GetBOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.BOpenId) {
		return nil, false
	}
	return o.BOpenId, true
}

// HasBOpenId returns a boolean if a field has been set.
func (o *JinyoutestopenidThree) HasBOpenId() bool {
	if o != nil && !IsNil(o.BOpenId) {
		return true
	}

	return false
}

// SetBOpenId gets a reference to the given string and assigns it to the BOpenId field.
func (o *JinyoutestopenidThree) SetBOpenId(v string) {
	o.BOpenId = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *JinyoutestopenidThree) GetF() JinyoutestopenidTwo {
	if o == nil || IsNil(o.F) {
		var ret JinyoutestopenidTwo
		return ret
	}
	return *o.F
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JinyoutestopenidThree) GetFOk() (*JinyoutestopenidTwo, bool) {
	if o == nil || IsNil(o.F) {
		return nil, false
	}
	return o.F, true
}

// HasF returns a boolean if a field has been set.
func (o *JinyoutestopenidThree) HasF() bool {
	if o != nil && !IsNil(o.F) {
		return true
	}

	return false
}

// SetF gets a reference to the given JinyoutestopenidTwo and assigns it to the F field.
func (o *JinyoutestopenidThree) SetF(v JinyoutestopenidTwo) {
	o.F = &v
}

func (o JinyoutestopenidThree) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JinyoutestopenidThree) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.A) {
		toSerialize["a"] = o.A
	}
	if !IsNil(o.B) {
		toSerialize["b"] = o.B
	}
	if !IsNil(o.BOpenId) {
		toSerialize["b_open_id"] = o.BOpenId
	}
	if !IsNil(o.F) {
		toSerialize["f"] = o.F
	}
	return toSerialize, nil
}

type NullableJinyoutestopenidThree struct {
	value *JinyoutestopenidThree
	isSet bool
}

func (v NullableJinyoutestopenidThree) Get() *JinyoutestopenidThree {
	return v.value
}

func (v *NullableJinyoutestopenidThree) Set(val *JinyoutestopenidThree) {
	v.value = val
	v.isSet = true
}

func (v NullableJinyoutestopenidThree) IsSet() bool {
	return v.isSet
}

func (v *NullableJinyoutestopenidThree) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJinyoutestopenidThree(val *JinyoutestopenidThree) *NullableJinyoutestopenidThree {
	return &NullableJinyoutestopenidThree{value: val, isSet: true}
}

func (v NullableJinyoutestopenidThree) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJinyoutestopenidThree) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


