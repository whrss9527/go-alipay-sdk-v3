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

// checks if the AlipayOpenMiniIsvFastregisterCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniIsvFastregisterCreateResponseModel{}

// AlipayOpenMiniIsvFastregisterCreateResponseModel struct for AlipayOpenMiniIsvFastregisterCreateResponseModel
type AlipayOpenMiniIsvFastregisterCreateResponseModel struct {
	// 授权确认跳转url
	AuthorizeUrl *string `json:"authorize_url,omitempty"`
	// 代创建试用小程序单号
	OrderNo *string `json:"order_no,omitempty"`
}

// NewAlipayOpenMiniIsvFastregisterCreateResponseModel instantiates a new AlipayOpenMiniIsvFastregisterCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniIsvFastregisterCreateResponseModel() *AlipayOpenMiniIsvFastregisterCreateResponseModel {
	this := AlipayOpenMiniIsvFastregisterCreateResponseModel{}
	return &this
}

// NewAlipayOpenMiniIsvFastregisterCreateResponseModelWithDefaults instantiates a new AlipayOpenMiniIsvFastregisterCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniIsvFastregisterCreateResponseModelWithDefaults() *AlipayOpenMiniIsvFastregisterCreateResponseModel {
	this := AlipayOpenMiniIsvFastregisterCreateResponseModel{}
	return &this
}

// GetAuthorizeUrl returns the AuthorizeUrl field value if set, zero value otherwise.
func (o *AlipayOpenMiniIsvFastregisterCreateResponseModel) GetAuthorizeUrl() string {
	if o == nil || IsNil(o.AuthorizeUrl) {
		var ret string
		return ret
	}
	return *o.AuthorizeUrl
}

// GetAuthorizeUrlOk returns a tuple with the AuthorizeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniIsvFastregisterCreateResponseModel) GetAuthorizeUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorizeUrl) {
		return nil, false
	}
	return o.AuthorizeUrl, true
}

// HasAuthorizeUrl returns a boolean if a field has been set.
func (o *AlipayOpenMiniIsvFastregisterCreateResponseModel) HasAuthorizeUrl() bool {
	if o != nil && !IsNil(o.AuthorizeUrl) {
		return true
	}

	return false
}

// SetAuthorizeUrl gets a reference to the given string and assigns it to the AuthorizeUrl field.
func (o *AlipayOpenMiniIsvFastregisterCreateResponseModel) SetAuthorizeUrl(v string) {
	o.AuthorizeUrl = &v
}

// GetOrderNo returns the OrderNo field value if set, zero value otherwise.
func (o *AlipayOpenMiniIsvFastregisterCreateResponseModel) GetOrderNo() string {
	if o == nil || IsNil(o.OrderNo) {
		var ret string
		return ret
	}
	return *o.OrderNo
}

// GetOrderNoOk returns a tuple with the OrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniIsvFastregisterCreateResponseModel) GetOrderNoOk() (*string, bool) {
	if o == nil || IsNil(o.OrderNo) {
		return nil, false
	}
	return o.OrderNo, true
}

// HasOrderNo returns a boolean if a field has been set.
func (o *AlipayOpenMiniIsvFastregisterCreateResponseModel) HasOrderNo() bool {
	if o != nil && !IsNil(o.OrderNo) {
		return true
	}

	return false
}

// SetOrderNo gets a reference to the given string and assigns it to the OrderNo field.
func (o *AlipayOpenMiniIsvFastregisterCreateResponseModel) SetOrderNo(v string) {
	o.OrderNo = &v
}

func (o AlipayOpenMiniIsvFastregisterCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniIsvFastregisterCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthorizeUrl) {
		toSerialize["authorize_url"] = o.AuthorizeUrl
	}
	if !IsNil(o.OrderNo) {
		toSerialize["order_no"] = o.OrderNo
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniIsvFastregisterCreateResponseModel struct {
	value *AlipayOpenMiniIsvFastregisterCreateResponseModel
	isSet bool
}

func (v NullableAlipayOpenMiniIsvFastregisterCreateResponseModel) Get() *AlipayOpenMiniIsvFastregisterCreateResponseModel {
	return v.value
}

func (v *NullableAlipayOpenMiniIsvFastregisterCreateResponseModel) Set(val *AlipayOpenMiniIsvFastregisterCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniIsvFastregisterCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniIsvFastregisterCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniIsvFastregisterCreateResponseModel(val *AlipayOpenMiniIsvFastregisterCreateResponseModel) *NullableAlipayOpenMiniIsvFastregisterCreateResponseModel {
	return &NullableAlipayOpenMiniIsvFastregisterCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniIsvFastregisterCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniIsvFastregisterCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


