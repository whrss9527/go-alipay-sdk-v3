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

// checks if the AlipayOpenSearchBoxactivityBatchqueryModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenSearchBoxactivityBatchqueryModel{}

// AlipayOpenSearchBoxactivityBatchqueryModel struct for AlipayOpenSearchBoxactivityBatchqueryModel
type AlipayOpenSearchBoxactivityBatchqueryModel struct {
	// 搜索直达id
	BoxId *string `json:"box_id,omitempty"`
	// 商户id，代运营模式下传入。代运营模式，需要服务商已获得商家\"运营支付宝小程序\"授权。
	MerchantId *string `json:"merchant_id,omitempty"`
	// 分页查询的当前页号,从1开始
	PageNumber *int32 `json:"page_number,omitempty"`
	// 每页查询的数量，默认10，不超过50
	PageSize *int32 `json:"page_size,omitempty"`
}

// NewAlipayOpenSearchBoxactivityBatchqueryModel instantiates a new AlipayOpenSearchBoxactivityBatchqueryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenSearchBoxactivityBatchqueryModel() *AlipayOpenSearchBoxactivityBatchqueryModel {
	this := AlipayOpenSearchBoxactivityBatchqueryModel{}
	return &this
}

// NewAlipayOpenSearchBoxactivityBatchqueryModelWithDefaults instantiates a new AlipayOpenSearchBoxactivityBatchqueryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenSearchBoxactivityBatchqueryModelWithDefaults() *AlipayOpenSearchBoxactivityBatchqueryModel {
	this := AlipayOpenSearchBoxactivityBatchqueryModel{}
	return &this
}

// GetBoxId returns the BoxId field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxactivityBatchqueryModel) GetBoxId() string {
	if o == nil || IsNil(o.BoxId) {
		var ret string
		return ret
	}
	return *o.BoxId
}

// GetBoxIdOk returns a tuple with the BoxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxactivityBatchqueryModel) GetBoxIdOk() (*string, bool) {
	if o == nil || IsNil(o.BoxId) {
		return nil, false
	}
	return o.BoxId, true
}

// HasBoxId returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxactivityBatchqueryModel) HasBoxId() bool {
	if o != nil && !IsNil(o.BoxId) {
		return true
	}

	return false
}

// SetBoxId gets a reference to the given string and assigns it to the BoxId field.
func (o *AlipayOpenSearchBoxactivityBatchqueryModel) SetBoxId(v string) {
	o.BoxId = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxactivityBatchqueryModel) GetMerchantId() string {
	if o == nil || IsNil(o.MerchantId) {
		var ret string
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxactivityBatchqueryModel) GetMerchantIdOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxactivityBatchqueryModel) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given string and assigns it to the MerchantId field.
func (o *AlipayOpenSearchBoxactivityBatchqueryModel) SetMerchantId(v string) {
	o.MerchantId = &v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxactivityBatchqueryModel) GetPageNumber() int32 {
	if o == nil || IsNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxactivityBatchqueryModel) GetPageNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxactivityBatchqueryModel) HasPageNumber() bool {
	if o != nil && !IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *AlipayOpenSearchBoxactivityBatchqueryModel) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxactivityBatchqueryModel) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxactivityBatchqueryModel) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxactivityBatchqueryModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *AlipayOpenSearchBoxactivityBatchqueryModel) SetPageSize(v int32) {
	o.PageSize = &v
}

func (o AlipayOpenSearchBoxactivityBatchqueryModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenSearchBoxactivityBatchqueryModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BoxId) {
		toSerialize["box_id"] = o.BoxId
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchant_id"] = o.MerchantId
	}
	if !IsNil(o.PageNumber) {
		toSerialize["page_number"] = o.PageNumber
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	return toSerialize, nil
}

type NullableAlipayOpenSearchBoxactivityBatchqueryModel struct {
	value *AlipayOpenSearchBoxactivityBatchqueryModel
	isSet bool
}

func (v NullableAlipayOpenSearchBoxactivityBatchqueryModel) Get() *AlipayOpenSearchBoxactivityBatchqueryModel {
	return v.value
}

func (v *NullableAlipayOpenSearchBoxactivityBatchqueryModel) Set(val *AlipayOpenSearchBoxactivityBatchqueryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSearchBoxactivityBatchqueryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSearchBoxactivityBatchqueryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSearchBoxactivityBatchqueryModel(val *AlipayOpenSearchBoxactivityBatchqueryModel) *NullableAlipayOpenSearchBoxactivityBatchqueryModel {
	return &NullableAlipayOpenSearchBoxactivityBatchqueryModel{value: val, isSet: true}
}

func (v NullableAlipayOpenSearchBoxactivityBatchqueryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSearchBoxactivityBatchqueryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


