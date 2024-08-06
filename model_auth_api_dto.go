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

// checks if the AuthApiDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthApiDTO{}

// AuthApiDTO struct for AuthApiDTO
type AuthApiDTO struct {
	// 敏感信息可申请接口名
	ApiName *string `json:"api_name,omitempty"`
	// 敏感信息申请字段
	FieldName *string `json:"field_name,omitempty"`
	// 敏感信息申请能力code值
	PackageCode *string `json:"package_code,omitempty"`
}

// NewAuthApiDTO instantiates a new AuthApiDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthApiDTO() *AuthApiDTO {
	this := AuthApiDTO{}
	return &this
}

// NewAuthApiDTOWithDefaults instantiates a new AuthApiDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthApiDTOWithDefaults() *AuthApiDTO {
	this := AuthApiDTO{}
	return &this
}

// GetApiName returns the ApiName field value if set, zero value otherwise.
func (o *AuthApiDTO) GetApiName() string {
	if o == nil || IsNil(o.ApiName) {
		var ret string
		return ret
	}
	return *o.ApiName
}

// GetApiNameOk returns a tuple with the ApiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthApiDTO) GetApiNameOk() (*string, bool) {
	if o == nil || IsNil(o.ApiName) {
		return nil, false
	}
	return o.ApiName, true
}

// HasApiName returns a boolean if a field has been set.
func (o *AuthApiDTO) HasApiName() bool {
	if o != nil && !IsNil(o.ApiName) {
		return true
	}

	return false
}

// SetApiName gets a reference to the given string and assigns it to the ApiName field.
func (o *AuthApiDTO) SetApiName(v string) {
	o.ApiName = &v
}

// GetFieldName returns the FieldName field value if set, zero value otherwise.
func (o *AuthApiDTO) GetFieldName() string {
	if o == nil || IsNil(o.FieldName) {
		var ret string
		return ret
	}
	return *o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthApiDTO) GetFieldNameOk() (*string, bool) {
	if o == nil || IsNil(o.FieldName) {
		return nil, false
	}
	return o.FieldName, true
}

// HasFieldName returns a boolean if a field has been set.
func (o *AuthApiDTO) HasFieldName() bool {
	if o != nil && !IsNil(o.FieldName) {
		return true
	}

	return false
}

// SetFieldName gets a reference to the given string and assigns it to the FieldName field.
func (o *AuthApiDTO) SetFieldName(v string) {
	o.FieldName = &v
}

// GetPackageCode returns the PackageCode field value if set, zero value otherwise.
func (o *AuthApiDTO) GetPackageCode() string {
	if o == nil || IsNil(o.PackageCode) {
		var ret string
		return ret
	}
	return *o.PackageCode
}

// GetPackageCodeOk returns a tuple with the PackageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthApiDTO) GetPackageCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PackageCode) {
		return nil, false
	}
	return o.PackageCode, true
}

// HasPackageCode returns a boolean if a field has been set.
func (o *AuthApiDTO) HasPackageCode() bool {
	if o != nil && !IsNil(o.PackageCode) {
		return true
	}

	return false
}

// SetPackageCode gets a reference to the given string and assigns it to the PackageCode field.
func (o *AuthApiDTO) SetPackageCode(v string) {
	o.PackageCode = &v
}

func (o AuthApiDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthApiDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiName) {
		toSerialize["api_name"] = o.ApiName
	}
	if !IsNil(o.FieldName) {
		toSerialize["field_name"] = o.FieldName
	}
	if !IsNil(o.PackageCode) {
		toSerialize["package_code"] = o.PackageCode
	}
	return toSerialize, nil
}

type NullableAuthApiDTO struct {
	value *AuthApiDTO
	isSet bool
}

func (v NullableAuthApiDTO) Get() *AuthApiDTO {
	return v.value
}

func (v *NullableAuthApiDTO) Set(val *AuthApiDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthApiDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthApiDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthApiDTO(val *AuthApiDTO) *NullableAuthApiDTO {
	return &NullableAuthApiDTO{value: val, isSet: true}
}

func (v NullableAuthApiDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthApiDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


