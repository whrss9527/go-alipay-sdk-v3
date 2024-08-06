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

// checks if the GavintestNewLeveaOne type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GavintestNewLeveaOne{}

// GavintestNewLeveaOne struct for GavintestNewLeveaOne
type GavintestNewLeveaOne struct {
	// 2
	Boolen *bool `json:"boolen,omitempty"`
	// 21
	Ces []int32 `json:"ces,omitempty"`
	// 12
	Des *string `json:"des,omitempty"`
	// 1
	Str []string `json:"str,omitempty"`
}

// NewGavintestNewLeveaOne instantiates a new GavintestNewLeveaOne object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGavintestNewLeveaOne() *GavintestNewLeveaOne {
	this := GavintestNewLeveaOne{}
	return &this
}

// NewGavintestNewLeveaOneWithDefaults instantiates a new GavintestNewLeveaOne object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGavintestNewLeveaOneWithDefaults() *GavintestNewLeveaOne {
	this := GavintestNewLeveaOne{}
	return &this
}

// GetBoolen returns the Boolen field value if set, zero value otherwise.
func (o *GavintestNewLeveaOne) GetBoolen() bool {
	if o == nil || IsNil(o.Boolen) {
		var ret bool
		return ret
	}
	return *o.Boolen
}

// GetBoolenOk returns a tuple with the Boolen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GavintestNewLeveaOne) GetBoolenOk() (*bool, bool) {
	if o == nil || IsNil(o.Boolen) {
		return nil, false
	}
	return o.Boolen, true
}

// HasBoolen returns a boolean if a field has been set.
func (o *GavintestNewLeveaOne) HasBoolen() bool {
	if o != nil && !IsNil(o.Boolen) {
		return true
	}

	return false
}

// SetBoolen gets a reference to the given bool and assigns it to the Boolen field.
func (o *GavintestNewLeveaOne) SetBoolen(v bool) {
	o.Boolen = &v
}

// GetCes returns the Ces field value if set, zero value otherwise.
func (o *GavintestNewLeveaOne) GetCes() []int32 {
	if o == nil || IsNil(o.Ces) {
		var ret []int32
		return ret
	}
	return o.Ces
}

// GetCesOk returns a tuple with the Ces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GavintestNewLeveaOne) GetCesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Ces) {
		return nil, false
	}
	return o.Ces, true
}

// HasCes returns a boolean if a field has been set.
func (o *GavintestNewLeveaOne) HasCes() bool {
	if o != nil && !IsNil(o.Ces) {
		return true
	}

	return false
}

// SetCes gets a reference to the given []int32 and assigns it to the Ces field.
func (o *GavintestNewLeveaOne) SetCes(v []int32) {
	o.Ces = v
}

// GetDes returns the Des field value if set, zero value otherwise.
func (o *GavintestNewLeveaOne) GetDes() string {
	if o == nil || IsNil(o.Des) {
		var ret string
		return ret
	}
	return *o.Des
}

// GetDesOk returns a tuple with the Des field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GavintestNewLeveaOne) GetDesOk() (*string, bool) {
	if o == nil || IsNil(o.Des) {
		return nil, false
	}
	return o.Des, true
}

// HasDes returns a boolean if a field has been set.
func (o *GavintestNewLeveaOne) HasDes() bool {
	if o != nil && !IsNil(o.Des) {
		return true
	}

	return false
}

// SetDes gets a reference to the given string and assigns it to the Des field.
func (o *GavintestNewLeveaOne) SetDes(v string) {
	o.Des = &v
}

// GetStr returns the Str field value if set, zero value otherwise.
func (o *GavintestNewLeveaOne) GetStr() []string {
	if o == nil || IsNil(o.Str) {
		var ret []string
		return ret
	}
	return o.Str
}

// GetStrOk returns a tuple with the Str field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GavintestNewLeveaOne) GetStrOk() ([]string, bool) {
	if o == nil || IsNil(o.Str) {
		return nil, false
	}
	return o.Str, true
}

// HasStr returns a boolean if a field has been set.
func (o *GavintestNewLeveaOne) HasStr() bool {
	if o != nil && !IsNil(o.Str) {
		return true
	}

	return false
}

// SetStr gets a reference to the given []string and assigns it to the Str field.
func (o *GavintestNewLeveaOne) SetStr(v []string) {
	o.Str = v
}

func (o GavintestNewLeveaOne) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GavintestNewLeveaOne) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Boolen) {
		toSerialize["boolen"] = o.Boolen
	}
	if !IsNil(o.Ces) {
		toSerialize["ces"] = o.Ces
	}
	if !IsNil(o.Des) {
		toSerialize["des"] = o.Des
	}
	if !IsNil(o.Str) {
		toSerialize["str"] = o.Str
	}
	return toSerialize, nil
}

type NullableGavintestNewLeveaOne struct {
	value *GavintestNewLeveaOne
	isSet bool
}

func (v NullableGavintestNewLeveaOne) Get() *GavintestNewLeveaOne {
	return v.value
}

func (v *NullableGavintestNewLeveaOne) Set(val *GavintestNewLeveaOne) {
	v.value = val
	v.isSet = true
}

func (v NullableGavintestNewLeveaOne) IsSet() bool {
	return v.isSet
}

func (v *NullableGavintestNewLeveaOne) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGavintestNewLeveaOne(val *GavintestNewLeveaOne) *NullableGavintestNewLeveaOne {
	return &NullableGavintestNewLeveaOne{value: val, isSet: true}
}

func (v NullableGavintestNewLeveaOne) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGavintestNewLeveaOne) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
