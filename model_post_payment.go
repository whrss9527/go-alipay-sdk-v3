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

// checks if the PostPayment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostPayment{}

// PostPayment struct for PostPayment
type PostPayment struct {
	// 后付费金额，单位为：元（人民币），精确到小数点后两位。
	Amount *string `json:"amount,omitempty"`
	// 计费说明
	Description *string `json:"description,omitempty"`
	// 后付费项目名称
	Name *string `json:"name,omitempty"`
}

// NewPostPayment instantiates a new PostPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostPayment() *PostPayment {
	this := PostPayment{}
	return &this
}

// NewPostPaymentWithDefaults instantiates a new PostPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostPaymentWithDefaults() *PostPayment {
	this := PostPayment{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PostPayment) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPayment) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PostPayment) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *PostPayment) SetAmount(v string) {
	o.Amount = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PostPayment) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPayment) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PostPayment) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PostPayment) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PostPayment) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPayment) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PostPayment) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PostPayment) SetName(v string) {
	o.Name = &v
}

func (o PostPayment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostPayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullablePostPayment struct {
	value *PostPayment
	isSet bool
}

func (v NullablePostPayment) Get() *PostPayment {
	return v.value
}

func (v *NullablePostPayment) Set(val *PostPayment) {
	v.value = val
	v.isSet = true
}

func (v NullablePostPayment) IsSet() bool {
	return v.isSet
}

func (v *NullablePostPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostPayment(val *PostPayment) *NullablePostPayment {
	return &NullablePostPayment{value: val, isSet: true}
}

func (v NullablePostPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


