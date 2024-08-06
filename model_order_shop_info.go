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

// checks if the OrderShopInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderShopInfo{}

// OrderShopInfo struct for OrderShopInfo
type OrderShopInfo struct {
	// 店铺地址
	Address *string `json:"address,omitempty"`
	// 蚂蚁门店shop_id
	AlipayShopId *string `json:"alipay_shop_id,omitempty"`
	// 门店其他业务属性，不同业务场景KEY枚举值不同，使用前请参考产品文档
	ExtInfo []OrderExtInfo `json:"ext_info,omitempty"`
	// 商户门店id 支持英文、数字的组合
	MerchantShopId *string `json:"merchant_shop_id,omitempty"`
	// 店铺详情链接地址
	MerchantShopLinkPage *string `json:"merchant_shop_link_page,omitempty"`
	// 店铺名称
	Name *string `json:"name,omitempty"`
	// 联系电话-支持固话或手机号 仅支持数字、+、- 。例如 手机：1380***1111、固话：021-888**888
	PhoneNum *string `json:"phone_num,omitempty"`
	// 仅当alipay_shop_id字段值为非标准蚂蚁门店时使用，其他场景无需传入
	Type *string `json:"type,omitempty"`
}

// NewOrderShopInfo instantiates a new OrderShopInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderShopInfo() *OrderShopInfo {
	this := OrderShopInfo{}
	return &this
}

// NewOrderShopInfoWithDefaults instantiates a new OrderShopInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderShopInfoWithDefaults() *OrderShopInfo {
	this := OrderShopInfo{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *OrderShopInfo) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShopInfo) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *OrderShopInfo) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *OrderShopInfo) SetAddress(v string) {
	o.Address = &v
}

// GetAlipayShopId returns the AlipayShopId field value if set, zero value otherwise.
func (o *OrderShopInfo) GetAlipayShopId() string {
	if o == nil || IsNil(o.AlipayShopId) {
		var ret string
		return ret
	}
	return *o.AlipayShopId
}

// GetAlipayShopIdOk returns a tuple with the AlipayShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShopInfo) GetAlipayShopIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlipayShopId) {
		return nil, false
	}
	return o.AlipayShopId, true
}

// HasAlipayShopId returns a boolean if a field has been set.
func (o *OrderShopInfo) HasAlipayShopId() bool {
	if o != nil && !IsNil(o.AlipayShopId) {
		return true
	}

	return false
}

// SetAlipayShopId gets a reference to the given string and assigns it to the AlipayShopId field.
func (o *OrderShopInfo) SetAlipayShopId(v string) {
	o.AlipayShopId = &v
}

// GetExtInfo returns the ExtInfo field value if set, zero value otherwise.
func (o *OrderShopInfo) GetExtInfo() []OrderExtInfo {
	if o == nil || IsNil(o.ExtInfo) {
		var ret []OrderExtInfo
		return ret
	}
	return o.ExtInfo
}

// GetExtInfoOk returns a tuple with the ExtInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShopInfo) GetExtInfoOk() ([]OrderExtInfo, bool) {
	if o == nil || IsNil(o.ExtInfo) {
		return nil, false
	}
	return o.ExtInfo, true
}

// HasExtInfo returns a boolean if a field has been set.
func (o *OrderShopInfo) HasExtInfo() bool {
	if o != nil && !IsNil(o.ExtInfo) {
		return true
	}

	return false
}

// SetExtInfo gets a reference to the given []OrderExtInfo and assigns it to the ExtInfo field.
func (o *OrderShopInfo) SetExtInfo(v []OrderExtInfo) {
	o.ExtInfo = v
}

// GetMerchantShopId returns the MerchantShopId field value if set, zero value otherwise.
func (o *OrderShopInfo) GetMerchantShopId() string {
	if o == nil || IsNil(o.MerchantShopId) {
		var ret string
		return ret
	}
	return *o.MerchantShopId
}

// GetMerchantShopIdOk returns a tuple with the MerchantShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShopInfo) GetMerchantShopIdOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantShopId) {
		return nil, false
	}
	return o.MerchantShopId, true
}

// HasMerchantShopId returns a boolean if a field has been set.
func (o *OrderShopInfo) HasMerchantShopId() bool {
	if o != nil && !IsNil(o.MerchantShopId) {
		return true
	}

	return false
}

// SetMerchantShopId gets a reference to the given string and assigns it to the MerchantShopId field.
func (o *OrderShopInfo) SetMerchantShopId(v string) {
	o.MerchantShopId = &v
}

// GetMerchantShopLinkPage returns the MerchantShopLinkPage field value if set, zero value otherwise.
func (o *OrderShopInfo) GetMerchantShopLinkPage() string {
	if o == nil || IsNil(o.MerchantShopLinkPage) {
		var ret string
		return ret
	}
	return *o.MerchantShopLinkPage
}

// GetMerchantShopLinkPageOk returns a tuple with the MerchantShopLinkPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShopInfo) GetMerchantShopLinkPageOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantShopLinkPage) {
		return nil, false
	}
	return o.MerchantShopLinkPage, true
}

// HasMerchantShopLinkPage returns a boolean if a field has been set.
func (o *OrderShopInfo) HasMerchantShopLinkPage() bool {
	if o != nil && !IsNil(o.MerchantShopLinkPage) {
		return true
	}

	return false
}

// SetMerchantShopLinkPage gets a reference to the given string and assigns it to the MerchantShopLinkPage field.
func (o *OrderShopInfo) SetMerchantShopLinkPage(v string) {
	o.MerchantShopLinkPage = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrderShopInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShopInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrderShopInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrderShopInfo) SetName(v string) {
	o.Name = &v
}

// GetPhoneNum returns the PhoneNum field value if set, zero value otherwise.
func (o *OrderShopInfo) GetPhoneNum() string {
	if o == nil || IsNil(o.PhoneNum) {
		var ret string
		return ret
	}
	return *o.PhoneNum
}

// GetPhoneNumOk returns a tuple with the PhoneNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShopInfo) GetPhoneNumOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNum) {
		return nil, false
	}
	return o.PhoneNum, true
}

// HasPhoneNum returns a boolean if a field has been set.
func (o *OrderShopInfo) HasPhoneNum() bool {
	if o != nil && !IsNil(o.PhoneNum) {
		return true
	}

	return false
}

// SetPhoneNum gets a reference to the given string and assigns it to the PhoneNum field.
func (o *OrderShopInfo) SetPhoneNum(v string) {
	o.PhoneNum = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrderShopInfo) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShopInfo) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrderShopInfo) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrderShopInfo) SetType(v string) {
	o.Type = &v
}

func (o OrderShopInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderShopInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.AlipayShopId) {
		toSerialize["alipay_shop_id"] = o.AlipayShopId
	}
	if !IsNil(o.ExtInfo) {
		toSerialize["ext_info"] = o.ExtInfo
	}
	if !IsNil(o.MerchantShopId) {
		toSerialize["merchant_shop_id"] = o.MerchantShopId
	}
	if !IsNil(o.MerchantShopLinkPage) {
		toSerialize["merchant_shop_link_page"] = o.MerchantShopLinkPage
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PhoneNum) {
		toSerialize["phone_num"] = o.PhoneNum
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableOrderShopInfo struct {
	value *OrderShopInfo
	isSet bool
}

func (v NullableOrderShopInfo) Get() *OrderShopInfo {
	return v.value
}

func (v *NullableOrderShopInfo) Set(val *OrderShopInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderShopInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderShopInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderShopInfo(val *OrderShopInfo) *NullableOrderShopInfo {
	return &NullableOrderShopInfo{value: val, isSet: true}
}

func (v NullableOrderShopInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderShopInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
