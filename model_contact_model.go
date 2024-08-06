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

// checks if the ContactModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactModel{}

// ContactModel struct for ContactModel
type ContactModel struct {
	// 联系人邮箱
	ContactEmail *string `json:"contact_email,omitempty"`
	// 联系人手机号码
	ContactMobile *string `json:"contact_mobile,omitempty"`
	// 联系人名称
	ContactName *string `json:"contact_name,omitempty"`
}

// NewContactModel instantiates a new ContactModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactModel() *ContactModel {
	this := ContactModel{}
	return &this
}

// NewContactModelWithDefaults instantiates a new ContactModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactModelWithDefaults() *ContactModel {
	this := ContactModel{}
	return &this
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise.
func (o *ContactModel) GetContactEmail() string {
	if o == nil || IsNil(o.ContactEmail) {
		var ret string
		return ret
	}
	return *o.ContactEmail
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactModel) GetContactEmailOk() (*string, bool) {
	if o == nil || IsNil(o.ContactEmail) {
		return nil, false
	}
	return o.ContactEmail, true
}

// HasContactEmail returns a boolean if a field has been set.
func (o *ContactModel) HasContactEmail() bool {
	if o != nil && !IsNil(o.ContactEmail) {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given string and assigns it to the ContactEmail field.
func (o *ContactModel) SetContactEmail(v string) {
	o.ContactEmail = &v
}

// GetContactMobile returns the ContactMobile field value if set, zero value otherwise.
func (o *ContactModel) GetContactMobile() string {
	if o == nil || IsNil(o.ContactMobile) {
		var ret string
		return ret
	}
	return *o.ContactMobile
}

// GetContactMobileOk returns a tuple with the ContactMobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactModel) GetContactMobileOk() (*string, bool) {
	if o == nil || IsNil(o.ContactMobile) {
		return nil, false
	}
	return o.ContactMobile, true
}

// HasContactMobile returns a boolean if a field has been set.
func (o *ContactModel) HasContactMobile() bool {
	if o != nil && !IsNil(o.ContactMobile) {
		return true
	}

	return false
}

// SetContactMobile gets a reference to the given string and assigns it to the ContactMobile field.
func (o *ContactModel) SetContactMobile(v string) {
	o.ContactMobile = &v
}

// GetContactName returns the ContactName field value if set, zero value otherwise.
func (o *ContactModel) GetContactName() string {
	if o == nil || IsNil(o.ContactName) {
		var ret string
		return ret
	}
	return *o.ContactName
}

// GetContactNameOk returns a tuple with the ContactName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactModel) GetContactNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContactName) {
		return nil, false
	}
	return o.ContactName, true
}

// HasContactName returns a boolean if a field has been set.
func (o *ContactModel) HasContactName() bool {
	if o != nil && !IsNil(o.ContactName) {
		return true
	}

	return false
}

// SetContactName gets a reference to the given string and assigns it to the ContactName field.
func (o *ContactModel) SetContactName(v string) {
	o.ContactName = &v
}

func (o ContactModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContactEmail) {
		toSerialize["contact_email"] = o.ContactEmail
	}
	if !IsNil(o.ContactMobile) {
		toSerialize["contact_mobile"] = o.ContactMobile
	}
	if !IsNil(o.ContactName) {
		toSerialize["contact_name"] = o.ContactName
	}
	return toSerialize, nil
}

type NullableContactModel struct {
	value *ContactModel
	isSet bool
}

func (v NullableContactModel) Get() *ContactModel {
	return v.value
}

func (v *NullableContactModel) Set(val *ContactModel) {
	v.value = val
	v.isSet = true
}

func (v NullableContactModel) IsSet() bool {
	return v.isSet
}

func (v *NullableContactModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactModel(val *ContactModel) *NullableContactModel {
	return &NullableContactModel{value: val, isSet: true}
}

func (v NullableContactModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
