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

// checks if the AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel{}

// AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel struct for AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel
type AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel struct {
	// 门店收单pid
	Pid *string `json:"pid,omitempty"`
	// 门店直间连类型
	RoleType *string `json:"role_type,omitempty"`
	// 门店ID
	ShopId *string `json:"shop_id,omitempty"`
	// 是否成功
	Success *bool `json:"success,omitempty"`
}

// NewAlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel instantiates a new AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel() *AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel {
	this := AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel{}
	return &this
}

// NewAlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModelWithDefaults instantiates a new AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModelWithDefaults() *AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel {
	this := AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel{}
	return &this
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel) GetPid() string {
	if o == nil || IsNil(o.Pid) {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel) GetPidOk() (*string, bool) {
	if o == nil || IsNil(o.Pid) {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel) HasPid() bool {
	if o != nil && !IsNil(o.Pid) {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel) SetPid(v string) {
	o.Pid = &v
}

// GetRoleType returns the RoleType field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel) GetRoleType() string {
	if o == nil || IsNil(o.RoleType) {
		var ret string
		return ret
	}
	return *o.RoleType
}

// GetRoleTypeOk returns a tuple with the RoleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel) GetRoleTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RoleType) {
		return nil, false
	}
	return o.RoleType, true
}

// HasRoleType returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel) HasRoleType() bool {
	if o != nil && !IsNil(o.RoleType) {
		return true
	}

	return false
}

// SetRoleType gets a reference to the given string and assigns it to the RoleType field.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel) SetRoleType(v string) {
	o.RoleType = &v
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel) GetShopId() string {
	if o == nil || IsNil(o.ShopId) {
		var ret string
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel) GetShopIdOk() (*string, bool) {
	if o == nil || IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel) HasShopId() bool {
	if o != nil && !IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given string and assigns it to the ShopId field.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel) SetShopId(v string) {
	o.ShopId = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel) SetSuccess(v bool) {
	o.Success = &v
}

func (o AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pid) {
		toSerialize["pid"] = o.Pid
	}
	if !IsNil(o.RoleType) {
		toSerialize["role_type"] = o.RoleType
	}
	if !IsNil(o.ShopId) {
		toSerialize["shop_id"] = o.ShopId
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel struct {
	value *AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel) Get() *AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel) Set(val *AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel(val *AlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel) *NullableAlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel {
	return &NullableAlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceEnterpriseMerchantrelationCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

