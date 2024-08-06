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

// checks if the ProductSignStatusInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductSignStatusInfo{}

// ProductSignStatusInfo struct for ProductSignStatusInfo
type ProductSignStatusInfo struct {
	// 产品编码
	ProductCode *string `json:"product_code,omitempty"`
	// 产品名称
	ProductName *string `json:"product_name,omitempty"`
	// none:未签约，表示还没有签约该产品  valid:已生效，表示合约已经生效，不需要再签约了  restrictValid:受限生效，表示合约已经生效，但是资料不全，功能受限  audit:审核中，已经有合约在审核中，请等待审核完成  waitConfirm:待商户确认协议，合约已经审核通过，需要商户确认后合约才生效  auditReject:审核未通过  invalid:合约失效，曾经签过合约，但已经失效了，可以重新发起签约  restrictInvalid:受限失效，受限合约失效了
	Status *string `json:"status,omitempty"`
}

// NewProductSignStatusInfo instantiates a new ProductSignStatusInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductSignStatusInfo() *ProductSignStatusInfo {
	this := ProductSignStatusInfo{}
	return &this
}

// NewProductSignStatusInfoWithDefaults instantiates a new ProductSignStatusInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductSignStatusInfoWithDefaults() *ProductSignStatusInfo {
	this := ProductSignStatusInfo{}
	return &this
}

// GetProductCode returns the ProductCode field value if set, zero value otherwise.
func (o *ProductSignStatusInfo) GetProductCode() string {
	if o == nil || IsNil(o.ProductCode) {
		var ret string
		return ret
	}
	return *o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductSignStatusInfo) GetProductCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductCode) {
		return nil, false
	}
	return o.ProductCode, true
}

// HasProductCode returns a boolean if a field has been set.
func (o *ProductSignStatusInfo) HasProductCode() bool {
	if o != nil && !IsNil(o.ProductCode) {
		return true
	}

	return false
}

// SetProductCode gets a reference to the given string and assigns it to the ProductCode field.
func (o *ProductSignStatusInfo) SetProductCode(v string) {
	o.ProductCode = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *ProductSignStatusInfo) GetProductName() string {
	if o == nil || IsNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductSignStatusInfo) GetProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *ProductSignStatusInfo) HasProductName() bool {
	if o != nil && !IsNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *ProductSignStatusInfo) SetProductName(v string) {
	o.ProductName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProductSignStatusInfo) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductSignStatusInfo) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProductSignStatusInfo) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ProductSignStatusInfo) SetStatus(v string) {
	o.Status = &v
}

func (o ProductSignStatusInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductSignStatusInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductCode) {
		toSerialize["product_code"] = o.ProductCode
	}
	if !IsNil(o.ProductName) {
		toSerialize["product_name"] = o.ProductName
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableProductSignStatusInfo struct {
	value *ProductSignStatusInfo
	isSet bool
}

func (v NullableProductSignStatusInfo) Get() *ProductSignStatusInfo {
	return v.value
}

func (v *NullableProductSignStatusInfo) Set(val *ProductSignStatusInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableProductSignStatusInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableProductSignStatusInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductSignStatusInfo(val *ProductSignStatusInfo) *NullableProductSignStatusInfo {
	return &NullableProductSignStatusInfo{value: val, isSet: true}
}

func (v NullableProductSignStatusInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductSignStatusInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


