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

// checks if the AntMerchantExpandShopCloseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AntMerchantExpandShopCloseModel{}

// AntMerchantExpandShopCloseModel struct for AntMerchantExpandShopCloseModel
type AntMerchantExpandShopCloseModel struct {
	// 商户角色id，表示将要开的店属于哪个商户角色。对于直连开店场景，填写商户pid；对于间连开店场景（线上、线下、直付通），填写商户smid。若未传入shop_id 则本参数与store_id均必填。
	IpRoleId *string `json:"ip_role_id,omitempty"`
	// 支付宝侧蚂蚁店铺 id。传入本参数后可不填 store_id 及 ip_role_id。
	ShopId *string `json:"shop_id,omitempty"`
	// 商户侧门店编号。表示该门店在该商户角色id(直连pid，间连smid)下，由商户自己定义的外部门店编号。若未传入 shop_id 则本参数与与ip_role_id均必填。
	StoreId *string `json:"store_id,omitempty"`
}

// NewAntMerchantExpandShopCloseModel instantiates a new AntMerchantExpandShopCloseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAntMerchantExpandShopCloseModel() *AntMerchantExpandShopCloseModel {
	this := AntMerchantExpandShopCloseModel{}
	return &this
}

// NewAntMerchantExpandShopCloseModelWithDefaults instantiates a new AntMerchantExpandShopCloseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAntMerchantExpandShopCloseModelWithDefaults() *AntMerchantExpandShopCloseModel {
	this := AntMerchantExpandShopCloseModel{}
	return &this
}

// GetIpRoleId returns the IpRoleId field value if set, zero value otherwise.
func (o *AntMerchantExpandShopCloseModel) GetIpRoleId() string {
	if o == nil || IsNil(o.IpRoleId) {
		var ret string
		return ret
	}
	return *o.IpRoleId
}

// GetIpRoleIdOk returns a tuple with the IpRoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandShopCloseModel) GetIpRoleIdOk() (*string, bool) {
	if o == nil || IsNil(o.IpRoleId) {
		return nil, false
	}
	return o.IpRoleId, true
}

// HasIpRoleId returns a boolean if a field has been set.
func (o *AntMerchantExpandShopCloseModel) HasIpRoleId() bool {
	if o != nil && !IsNil(o.IpRoleId) {
		return true
	}

	return false
}

// SetIpRoleId gets a reference to the given string and assigns it to the IpRoleId field.
func (o *AntMerchantExpandShopCloseModel) SetIpRoleId(v string) {
	o.IpRoleId = &v
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *AntMerchantExpandShopCloseModel) GetShopId() string {
	if o == nil || IsNil(o.ShopId) {
		var ret string
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandShopCloseModel) GetShopIdOk() (*string, bool) {
	if o == nil || IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *AntMerchantExpandShopCloseModel) HasShopId() bool {
	if o != nil && !IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given string and assigns it to the ShopId field.
func (o *AntMerchantExpandShopCloseModel) SetShopId(v string) {
	o.ShopId = &v
}

// GetStoreId returns the StoreId field value if set, zero value otherwise.
func (o *AntMerchantExpandShopCloseModel) GetStoreId() string {
	if o == nil || IsNil(o.StoreId) {
		var ret string
		return ret
	}
	return *o.StoreId
}

// GetStoreIdOk returns a tuple with the StoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandShopCloseModel) GetStoreIdOk() (*string, bool) {
	if o == nil || IsNil(o.StoreId) {
		return nil, false
	}
	return o.StoreId, true
}

// HasStoreId returns a boolean if a field has been set.
func (o *AntMerchantExpandShopCloseModel) HasStoreId() bool {
	if o != nil && !IsNil(o.StoreId) {
		return true
	}

	return false
}

// SetStoreId gets a reference to the given string and assigns it to the StoreId field.
func (o *AntMerchantExpandShopCloseModel) SetStoreId(v string) {
	o.StoreId = &v
}

func (o AntMerchantExpandShopCloseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AntMerchantExpandShopCloseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IpRoleId) {
		toSerialize["ip_role_id"] = o.IpRoleId
	}
	if !IsNil(o.ShopId) {
		toSerialize["shop_id"] = o.ShopId
	}
	if !IsNil(o.StoreId) {
		toSerialize["store_id"] = o.StoreId
	}
	return toSerialize, nil
}

type NullableAntMerchantExpandShopCloseModel struct {
	value *AntMerchantExpandShopCloseModel
	isSet bool
}

func (v NullableAntMerchantExpandShopCloseModel) Get() *AntMerchantExpandShopCloseModel {
	return v.value
}

func (v *NullableAntMerchantExpandShopCloseModel) Set(val *AntMerchantExpandShopCloseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAntMerchantExpandShopCloseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAntMerchantExpandShopCloseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntMerchantExpandShopCloseModel(val *AntMerchantExpandShopCloseModel) *NullableAntMerchantExpandShopCloseModel {
	return &NullableAntMerchantExpandShopCloseModel{value: val, isSet: true}
}

func (v NullableAntMerchantExpandShopCloseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntMerchantExpandShopCloseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


