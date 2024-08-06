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

// checks if the CustomerGuide type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerGuide{}

// CustomerGuide struct for CustomerGuide
type CustomerGuide struct {
	// 券可用的小程序appId，卡包详情页可跳转到该appId
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// 指定跳转到mini_app_id时的具体页面路径。  限制：  1、只有mini_app_id有值时该值传入才会有效 2、该小程序路径是相对路径。详情参见 <a href=\"https://opendocs.alipay.com/support/01rb18\">小程序scheme链接介绍</a>
	MiniAppPath *string `json:"mini_app_path,omitempty"`
	// 代运营商业关系门店列表，列表中的门店id是调用接口alipay.business.relation.shop.create创建门店返回的real_shop_id。接口参数是列表类型。
	RealShopIds []string `json:"real_shop_ids,omitempty"`
	// 小程序服务编码，通过 alipay.open.app.appcontent.function.create(小程序服务创建)接口创建服务后获取。
	ServiceCodes []string `json:"service_codes,omitempty"`
	// 券可使用的门店列表。列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店id  接口参数是列表类型。
	ShopIds []string `json:"shop_ids,omitempty"`
	// 该字段后续废弃。券可使用的门店列表。列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店id。接口参数是列表类型。
	StoreIds []string `json:"store_ids,omitempty"`
	VoucherSendGuide *VoucherSendGuide `json:"voucher_send_guide,omitempty"`
	VoucherUseGuide *VoucherUseGuide `json:"voucher_use_guide,omitempty"`
}

// NewCustomerGuide instantiates a new CustomerGuide object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerGuide() *CustomerGuide {
	this := CustomerGuide{}
	return &this
}

// NewCustomerGuideWithDefaults instantiates a new CustomerGuide object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerGuideWithDefaults() *CustomerGuide {
	this := CustomerGuide{}
	return &this
}

// GetMiniAppId returns the MiniAppId field value if set, zero value otherwise.
func (o *CustomerGuide) GetMiniAppId() string {
	if o == nil || IsNil(o.MiniAppId) {
		var ret string
		return ret
	}
	return *o.MiniAppId
}

// GetMiniAppIdOk returns a tuple with the MiniAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGuide) GetMiniAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.MiniAppId) {
		return nil, false
	}
	return o.MiniAppId, true
}

// HasMiniAppId returns a boolean if a field has been set.
func (o *CustomerGuide) HasMiniAppId() bool {
	if o != nil && !IsNil(o.MiniAppId) {
		return true
	}

	return false
}

// SetMiniAppId gets a reference to the given string and assigns it to the MiniAppId field.
func (o *CustomerGuide) SetMiniAppId(v string) {
	o.MiniAppId = &v
}

// GetMiniAppPath returns the MiniAppPath field value if set, zero value otherwise.
func (o *CustomerGuide) GetMiniAppPath() string {
	if o == nil || IsNil(o.MiniAppPath) {
		var ret string
		return ret
	}
	return *o.MiniAppPath
}

// GetMiniAppPathOk returns a tuple with the MiniAppPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGuide) GetMiniAppPathOk() (*string, bool) {
	if o == nil || IsNil(o.MiniAppPath) {
		return nil, false
	}
	return o.MiniAppPath, true
}

// HasMiniAppPath returns a boolean if a field has been set.
func (o *CustomerGuide) HasMiniAppPath() bool {
	if o != nil && !IsNil(o.MiniAppPath) {
		return true
	}

	return false
}

// SetMiniAppPath gets a reference to the given string and assigns it to the MiniAppPath field.
func (o *CustomerGuide) SetMiniAppPath(v string) {
	o.MiniAppPath = &v
}

// GetRealShopIds returns the RealShopIds field value if set, zero value otherwise.
func (o *CustomerGuide) GetRealShopIds() []string {
	if o == nil || IsNil(o.RealShopIds) {
		var ret []string
		return ret
	}
	return o.RealShopIds
}

// GetRealShopIdsOk returns a tuple with the RealShopIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGuide) GetRealShopIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RealShopIds) {
		return nil, false
	}
	return o.RealShopIds, true
}

// HasRealShopIds returns a boolean if a field has been set.
func (o *CustomerGuide) HasRealShopIds() bool {
	if o != nil && !IsNil(o.RealShopIds) {
		return true
	}

	return false
}

// SetRealShopIds gets a reference to the given []string and assigns it to the RealShopIds field.
func (o *CustomerGuide) SetRealShopIds(v []string) {
	o.RealShopIds = v
}

// GetServiceCodes returns the ServiceCodes field value if set, zero value otherwise.
func (o *CustomerGuide) GetServiceCodes() []string {
	if o == nil || IsNil(o.ServiceCodes) {
		var ret []string
		return ret
	}
	return o.ServiceCodes
}

// GetServiceCodesOk returns a tuple with the ServiceCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGuide) GetServiceCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.ServiceCodes) {
		return nil, false
	}
	return o.ServiceCodes, true
}

// HasServiceCodes returns a boolean if a field has been set.
func (o *CustomerGuide) HasServiceCodes() bool {
	if o != nil && !IsNil(o.ServiceCodes) {
		return true
	}

	return false
}

// SetServiceCodes gets a reference to the given []string and assigns it to the ServiceCodes field.
func (o *CustomerGuide) SetServiceCodes(v []string) {
	o.ServiceCodes = v
}

// GetShopIds returns the ShopIds field value if set, zero value otherwise.
func (o *CustomerGuide) GetShopIds() []string {
	if o == nil || IsNil(o.ShopIds) {
		var ret []string
		return ret
	}
	return o.ShopIds
}

// GetShopIdsOk returns a tuple with the ShopIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGuide) GetShopIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ShopIds) {
		return nil, false
	}
	return o.ShopIds, true
}

// HasShopIds returns a boolean if a field has been set.
func (o *CustomerGuide) HasShopIds() bool {
	if o != nil && !IsNil(o.ShopIds) {
		return true
	}

	return false
}

// SetShopIds gets a reference to the given []string and assigns it to the ShopIds field.
func (o *CustomerGuide) SetShopIds(v []string) {
	o.ShopIds = v
}

// GetStoreIds returns the StoreIds field value if set, zero value otherwise.
func (o *CustomerGuide) GetStoreIds() []string {
	if o == nil || IsNil(o.StoreIds) {
		var ret []string
		return ret
	}
	return o.StoreIds
}

// GetStoreIdsOk returns a tuple with the StoreIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGuide) GetStoreIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.StoreIds) {
		return nil, false
	}
	return o.StoreIds, true
}

// HasStoreIds returns a boolean if a field has been set.
func (o *CustomerGuide) HasStoreIds() bool {
	if o != nil && !IsNil(o.StoreIds) {
		return true
	}

	return false
}

// SetStoreIds gets a reference to the given []string and assigns it to the StoreIds field.
func (o *CustomerGuide) SetStoreIds(v []string) {
	o.StoreIds = v
}

// GetVoucherSendGuide returns the VoucherSendGuide field value if set, zero value otherwise.
func (o *CustomerGuide) GetVoucherSendGuide() VoucherSendGuide {
	if o == nil || IsNil(o.VoucherSendGuide) {
		var ret VoucherSendGuide
		return ret
	}
	return *o.VoucherSendGuide
}

// GetVoucherSendGuideOk returns a tuple with the VoucherSendGuide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGuide) GetVoucherSendGuideOk() (*VoucherSendGuide, bool) {
	if o == nil || IsNil(o.VoucherSendGuide) {
		return nil, false
	}
	return o.VoucherSendGuide, true
}

// HasVoucherSendGuide returns a boolean if a field has been set.
func (o *CustomerGuide) HasVoucherSendGuide() bool {
	if o != nil && !IsNil(o.VoucherSendGuide) {
		return true
	}

	return false
}

// SetVoucherSendGuide gets a reference to the given VoucherSendGuide and assigns it to the VoucherSendGuide field.
func (o *CustomerGuide) SetVoucherSendGuide(v VoucherSendGuide) {
	o.VoucherSendGuide = &v
}

// GetVoucherUseGuide returns the VoucherUseGuide field value if set, zero value otherwise.
func (o *CustomerGuide) GetVoucherUseGuide() VoucherUseGuide {
	if o == nil || IsNil(o.VoucherUseGuide) {
		var ret VoucherUseGuide
		return ret
	}
	return *o.VoucherUseGuide
}

// GetVoucherUseGuideOk returns a tuple with the VoucherUseGuide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGuide) GetVoucherUseGuideOk() (*VoucherUseGuide, bool) {
	if o == nil || IsNil(o.VoucherUseGuide) {
		return nil, false
	}
	return o.VoucherUseGuide, true
}

// HasVoucherUseGuide returns a boolean if a field has been set.
func (o *CustomerGuide) HasVoucherUseGuide() bool {
	if o != nil && !IsNil(o.VoucherUseGuide) {
		return true
	}

	return false
}

// SetVoucherUseGuide gets a reference to the given VoucherUseGuide and assigns it to the VoucherUseGuide field.
func (o *CustomerGuide) SetVoucherUseGuide(v VoucherUseGuide) {
	o.VoucherUseGuide = &v
}

func (o CustomerGuide) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerGuide) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MiniAppId) {
		toSerialize["mini_app_id"] = o.MiniAppId
	}
	if !IsNil(o.MiniAppPath) {
		toSerialize["mini_app_path"] = o.MiniAppPath
	}
	if !IsNil(o.RealShopIds) {
		toSerialize["real_shop_ids"] = o.RealShopIds
	}
	if !IsNil(o.ServiceCodes) {
		toSerialize["service_codes"] = o.ServiceCodes
	}
	if !IsNil(o.ShopIds) {
		toSerialize["shop_ids"] = o.ShopIds
	}
	if !IsNil(o.StoreIds) {
		toSerialize["store_ids"] = o.StoreIds
	}
	if !IsNil(o.VoucherSendGuide) {
		toSerialize["voucher_send_guide"] = o.VoucherSendGuide
	}
	if !IsNil(o.VoucherUseGuide) {
		toSerialize["voucher_use_guide"] = o.VoucherUseGuide
	}
	return toSerialize, nil
}

type NullableCustomerGuide struct {
	value *CustomerGuide
	isSet bool
}

func (v NullableCustomerGuide) Get() *CustomerGuide {
	return v.value
}

func (v *NullableCustomerGuide) Set(val *CustomerGuide) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerGuide) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerGuide) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerGuide(val *CustomerGuide) *NullableCustomerGuide {
	return &NullableCustomerGuide{value: val, isSet: true}
}

func (v NullableCustomerGuide) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerGuide) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


