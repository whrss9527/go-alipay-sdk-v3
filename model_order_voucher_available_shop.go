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

// checks if the OrderVoucherAvailableShop type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderVoucherAvailableShop{}

// OrderVoucherAvailableShop struct for OrderVoucherAvailableShop
type OrderVoucherAvailableShop struct {
	OrderVoucherMerchantAllShop *OrderVoucherMerchantAllShop `json:"order_voucher_merchant_all_shop,omitempty"`
	// 代运营商业关系门店列表，列表中的门店id是调用接口alipay.business.relation.shop.create创建门店返回的real_shop_id  接口参数是列表类型。
	RealShopIds []string `json:"real_shop_ids,omitempty"`
	// 券可使用的门店列表。列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店id  接口参数是列表类型。
	ShopIds []string `json:"shop_ids,omitempty"`
}

// NewOrderVoucherAvailableShop instantiates a new OrderVoucherAvailableShop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderVoucherAvailableShop() *OrderVoucherAvailableShop {
	this := OrderVoucherAvailableShop{}
	return &this
}

// NewOrderVoucherAvailableShopWithDefaults instantiates a new OrderVoucherAvailableShop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderVoucherAvailableShopWithDefaults() *OrderVoucherAvailableShop {
	this := OrderVoucherAvailableShop{}
	return &this
}

// GetOrderVoucherMerchantAllShop returns the OrderVoucherMerchantAllShop field value if set, zero value otherwise.
func (o *OrderVoucherAvailableShop) GetOrderVoucherMerchantAllShop() OrderVoucherMerchantAllShop {
	if o == nil || IsNil(o.OrderVoucherMerchantAllShop) {
		var ret OrderVoucherMerchantAllShop
		return ret
	}
	return *o.OrderVoucherMerchantAllShop
}

// GetOrderVoucherMerchantAllShopOk returns a tuple with the OrderVoucherMerchantAllShop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderVoucherAvailableShop) GetOrderVoucherMerchantAllShopOk() (*OrderVoucherMerchantAllShop, bool) {
	if o == nil || IsNil(o.OrderVoucherMerchantAllShop) {
		return nil, false
	}
	return o.OrderVoucherMerchantAllShop, true
}

// HasOrderVoucherMerchantAllShop returns a boolean if a field has been set.
func (o *OrderVoucherAvailableShop) HasOrderVoucherMerchantAllShop() bool {
	if o != nil && !IsNil(o.OrderVoucherMerchantAllShop) {
		return true
	}

	return false
}

// SetOrderVoucherMerchantAllShop gets a reference to the given OrderVoucherMerchantAllShop and assigns it to the OrderVoucherMerchantAllShop field.
func (o *OrderVoucherAvailableShop) SetOrderVoucherMerchantAllShop(v OrderVoucherMerchantAllShop) {
	o.OrderVoucherMerchantAllShop = &v
}

// GetRealShopIds returns the RealShopIds field value if set, zero value otherwise.
func (o *OrderVoucherAvailableShop) GetRealShopIds() []string {
	if o == nil || IsNil(o.RealShopIds) {
		var ret []string
		return ret
	}
	return o.RealShopIds
}

// GetRealShopIdsOk returns a tuple with the RealShopIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderVoucherAvailableShop) GetRealShopIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RealShopIds) {
		return nil, false
	}
	return o.RealShopIds, true
}

// HasRealShopIds returns a boolean if a field has been set.
func (o *OrderVoucherAvailableShop) HasRealShopIds() bool {
	if o != nil && !IsNil(o.RealShopIds) {
		return true
	}

	return false
}

// SetRealShopIds gets a reference to the given []string and assigns it to the RealShopIds field.
func (o *OrderVoucherAvailableShop) SetRealShopIds(v []string) {
	o.RealShopIds = v
}

// GetShopIds returns the ShopIds field value if set, zero value otherwise.
func (o *OrderVoucherAvailableShop) GetShopIds() []string {
	if o == nil || IsNil(o.ShopIds) {
		var ret []string
		return ret
	}
	return o.ShopIds
}

// GetShopIdsOk returns a tuple with the ShopIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderVoucherAvailableShop) GetShopIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ShopIds) {
		return nil, false
	}
	return o.ShopIds, true
}

// HasShopIds returns a boolean if a field has been set.
func (o *OrderVoucherAvailableShop) HasShopIds() bool {
	if o != nil && !IsNil(o.ShopIds) {
		return true
	}

	return false
}

// SetShopIds gets a reference to the given []string and assigns it to the ShopIds field.
func (o *OrderVoucherAvailableShop) SetShopIds(v []string) {
	o.ShopIds = v
}

func (o OrderVoucherAvailableShop) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderVoucherAvailableShop) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderVoucherMerchantAllShop) {
		toSerialize["order_voucher_merchant_all_shop"] = o.OrderVoucherMerchantAllShop
	}
	if !IsNil(o.RealShopIds) {
		toSerialize["real_shop_ids"] = o.RealShopIds
	}
	if !IsNil(o.ShopIds) {
		toSerialize["shop_ids"] = o.ShopIds
	}
	return toSerialize, nil
}

type NullableOrderVoucherAvailableShop struct {
	value *OrderVoucherAvailableShop
	isSet bool
}

func (v NullableOrderVoucherAvailableShop) Get() *OrderVoucherAvailableShop {
	return v.value
}

func (v *NullableOrderVoucherAvailableShop) Set(val *OrderVoucherAvailableShop) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderVoucherAvailableShop) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderVoucherAvailableShop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderVoucherAvailableShop(val *OrderVoucherAvailableShop) *NullableOrderVoucherAvailableShop {
	return &NullableOrderVoucherAvailableShop{value: val, isSet: true}
}

func (v NullableOrderVoucherAvailableShop) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderVoucherAvailableShop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
