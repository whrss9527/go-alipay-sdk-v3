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

// checks if the MiniAppServiceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MiniAppServiceInfo{}

// MiniAppServiceInfo struct for MiniAppServiceInfo
type MiniAppServiceInfo struct {
	// 插件发布状态码，暂存100，风控审核200，运营审核300，等待上架400，已预发上架500，已上架501，已下架600，已驳回700
	BizStatus *string `json:"biz_status,omitempty"`
	// 是否是内部标，true/false
	IsInner *bool `json:"is_inner,omitempty"`
	// 是否订购，true/false
	IsOrder *bool `json:"is_order,omitempty"`
	// 三方应用appid
	IsvAppId *string `json:"isv_app_id,omitempty"`
	// 应用id
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// 测试插件
	MiniAppName *string `json:"mini_app_name,omitempty"`
	// 卖家pid
	SellerId *string `json:"seller_id,omitempty"`
	// 卖家名
	SellerName *string `json:"seller_name,omitempty"`
	// 商品CODE
	ServiceCode *string `json:"service_code,omitempty"`
	// 服务图标
	ServiceLogo *string `json:"service_logo,omitempty"`
	// 服务名
	ServiceName *string `json:"service_name,omitempty"`
	// 服务简介
	ServiceSlogan *string `json:"service_slogan,omitempty"`
	// 是否在服务市场透出，SHOW展示、HIDE隐藏
	ShowType *string `json:"show_type,omitempty"`
}

// NewMiniAppServiceInfo instantiates a new MiniAppServiceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMiniAppServiceInfo() *MiniAppServiceInfo {
	this := MiniAppServiceInfo{}
	return &this
}

// NewMiniAppServiceInfoWithDefaults instantiates a new MiniAppServiceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMiniAppServiceInfoWithDefaults() *MiniAppServiceInfo {
	this := MiniAppServiceInfo{}
	return &this
}

// GetBizStatus returns the BizStatus field value if set, zero value otherwise.
func (o *MiniAppServiceInfo) GetBizStatus() string {
	if o == nil || IsNil(o.BizStatus) {
		var ret string
		return ret
	}
	return *o.BizStatus
}

// GetBizStatusOk returns a tuple with the BizStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppServiceInfo) GetBizStatusOk() (*string, bool) {
	if o == nil || IsNil(o.BizStatus) {
		return nil, false
	}
	return o.BizStatus, true
}

// HasBizStatus returns a boolean if a field has been set.
func (o *MiniAppServiceInfo) HasBizStatus() bool {
	if o != nil && !IsNil(o.BizStatus) {
		return true
	}

	return false
}

// SetBizStatus gets a reference to the given string and assigns it to the BizStatus field.
func (o *MiniAppServiceInfo) SetBizStatus(v string) {
	o.BizStatus = &v
}

// GetIsInner returns the IsInner field value if set, zero value otherwise.
func (o *MiniAppServiceInfo) GetIsInner() bool {
	if o == nil || IsNil(o.IsInner) {
		var ret bool
		return ret
	}
	return *o.IsInner
}

// GetIsInnerOk returns a tuple with the IsInner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppServiceInfo) GetIsInnerOk() (*bool, bool) {
	if o == nil || IsNil(o.IsInner) {
		return nil, false
	}
	return o.IsInner, true
}

// HasIsInner returns a boolean if a field has been set.
func (o *MiniAppServiceInfo) HasIsInner() bool {
	if o != nil && !IsNil(o.IsInner) {
		return true
	}

	return false
}

// SetIsInner gets a reference to the given bool and assigns it to the IsInner field.
func (o *MiniAppServiceInfo) SetIsInner(v bool) {
	o.IsInner = &v
}

// GetIsOrder returns the IsOrder field value if set, zero value otherwise.
func (o *MiniAppServiceInfo) GetIsOrder() bool {
	if o == nil || IsNil(o.IsOrder) {
		var ret bool
		return ret
	}
	return *o.IsOrder
}

// GetIsOrderOk returns a tuple with the IsOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppServiceInfo) GetIsOrderOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOrder) {
		return nil, false
	}
	return o.IsOrder, true
}

// HasIsOrder returns a boolean if a field has been set.
func (o *MiniAppServiceInfo) HasIsOrder() bool {
	if o != nil && !IsNil(o.IsOrder) {
		return true
	}

	return false
}

// SetIsOrder gets a reference to the given bool and assigns it to the IsOrder field.
func (o *MiniAppServiceInfo) SetIsOrder(v bool) {
	o.IsOrder = &v
}

// GetIsvAppId returns the IsvAppId field value if set, zero value otherwise.
func (o *MiniAppServiceInfo) GetIsvAppId() string {
	if o == nil || IsNil(o.IsvAppId) {
		var ret string
		return ret
	}
	return *o.IsvAppId
}

// GetIsvAppIdOk returns a tuple with the IsvAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppServiceInfo) GetIsvAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.IsvAppId) {
		return nil, false
	}
	return o.IsvAppId, true
}

// HasIsvAppId returns a boolean if a field has been set.
func (o *MiniAppServiceInfo) HasIsvAppId() bool {
	if o != nil && !IsNil(o.IsvAppId) {
		return true
	}

	return false
}

// SetIsvAppId gets a reference to the given string and assigns it to the IsvAppId field.
func (o *MiniAppServiceInfo) SetIsvAppId(v string) {
	o.IsvAppId = &v
}

// GetMiniAppId returns the MiniAppId field value if set, zero value otherwise.
func (o *MiniAppServiceInfo) GetMiniAppId() string {
	if o == nil || IsNil(o.MiniAppId) {
		var ret string
		return ret
	}
	return *o.MiniAppId
}

// GetMiniAppIdOk returns a tuple with the MiniAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppServiceInfo) GetMiniAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.MiniAppId) {
		return nil, false
	}
	return o.MiniAppId, true
}

// HasMiniAppId returns a boolean if a field has been set.
func (o *MiniAppServiceInfo) HasMiniAppId() bool {
	if o != nil && !IsNil(o.MiniAppId) {
		return true
	}

	return false
}

// SetMiniAppId gets a reference to the given string and assigns it to the MiniAppId field.
func (o *MiniAppServiceInfo) SetMiniAppId(v string) {
	o.MiniAppId = &v
}

// GetMiniAppName returns the MiniAppName field value if set, zero value otherwise.
func (o *MiniAppServiceInfo) GetMiniAppName() string {
	if o == nil || IsNil(o.MiniAppName) {
		var ret string
		return ret
	}
	return *o.MiniAppName
}

// GetMiniAppNameOk returns a tuple with the MiniAppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppServiceInfo) GetMiniAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.MiniAppName) {
		return nil, false
	}
	return o.MiniAppName, true
}

// HasMiniAppName returns a boolean if a field has been set.
func (o *MiniAppServiceInfo) HasMiniAppName() bool {
	if o != nil && !IsNil(o.MiniAppName) {
		return true
	}

	return false
}

// SetMiniAppName gets a reference to the given string and assigns it to the MiniAppName field.
func (o *MiniAppServiceInfo) SetMiniAppName(v string) {
	o.MiniAppName = &v
}

// GetSellerId returns the SellerId field value if set, zero value otherwise.
func (o *MiniAppServiceInfo) GetSellerId() string {
	if o == nil || IsNil(o.SellerId) {
		var ret string
		return ret
	}
	return *o.SellerId
}

// GetSellerIdOk returns a tuple with the SellerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppServiceInfo) GetSellerIdOk() (*string, bool) {
	if o == nil || IsNil(o.SellerId) {
		return nil, false
	}
	return o.SellerId, true
}

// HasSellerId returns a boolean if a field has been set.
func (o *MiniAppServiceInfo) HasSellerId() bool {
	if o != nil && !IsNil(o.SellerId) {
		return true
	}

	return false
}

// SetSellerId gets a reference to the given string and assigns it to the SellerId field.
func (o *MiniAppServiceInfo) SetSellerId(v string) {
	o.SellerId = &v
}

// GetSellerName returns the SellerName field value if set, zero value otherwise.
func (o *MiniAppServiceInfo) GetSellerName() string {
	if o == nil || IsNil(o.SellerName) {
		var ret string
		return ret
	}
	return *o.SellerName
}

// GetSellerNameOk returns a tuple with the SellerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppServiceInfo) GetSellerNameOk() (*string, bool) {
	if o == nil || IsNil(o.SellerName) {
		return nil, false
	}
	return o.SellerName, true
}

// HasSellerName returns a boolean if a field has been set.
func (o *MiniAppServiceInfo) HasSellerName() bool {
	if o != nil && !IsNil(o.SellerName) {
		return true
	}

	return false
}

// SetSellerName gets a reference to the given string and assigns it to the SellerName field.
func (o *MiniAppServiceInfo) SetSellerName(v string) {
	o.SellerName = &v
}

// GetServiceCode returns the ServiceCode field value if set, zero value otherwise.
func (o *MiniAppServiceInfo) GetServiceCode() string {
	if o == nil || IsNil(o.ServiceCode) {
		var ret string
		return ret
	}
	return *o.ServiceCode
}

// GetServiceCodeOk returns a tuple with the ServiceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppServiceInfo) GetServiceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceCode) {
		return nil, false
	}
	return o.ServiceCode, true
}

// HasServiceCode returns a boolean if a field has been set.
func (o *MiniAppServiceInfo) HasServiceCode() bool {
	if o != nil && !IsNil(o.ServiceCode) {
		return true
	}

	return false
}

// SetServiceCode gets a reference to the given string and assigns it to the ServiceCode field.
func (o *MiniAppServiceInfo) SetServiceCode(v string) {
	o.ServiceCode = &v
}

// GetServiceLogo returns the ServiceLogo field value if set, zero value otherwise.
func (o *MiniAppServiceInfo) GetServiceLogo() string {
	if o == nil || IsNil(o.ServiceLogo) {
		var ret string
		return ret
	}
	return *o.ServiceLogo
}

// GetServiceLogoOk returns a tuple with the ServiceLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppServiceInfo) GetServiceLogoOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceLogo) {
		return nil, false
	}
	return o.ServiceLogo, true
}

// HasServiceLogo returns a boolean if a field has been set.
func (o *MiniAppServiceInfo) HasServiceLogo() bool {
	if o != nil && !IsNil(o.ServiceLogo) {
		return true
	}

	return false
}

// SetServiceLogo gets a reference to the given string and assigns it to the ServiceLogo field.
func (o *MiniAppServiceInfo) SetServiceLogo(v string) {
	o.ServiceLogo = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *MiniAppServiceInfo) GetServiceName() string {
	if o == nil || IsNil(o.ServiceName) {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppServiceInfo) GetServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceName) {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *MiniAppServiceInfo) HasServiceName() bool {
	if o != nil && !IsNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *MiniAppServiceInfo) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetServiceSlogan returns the ServiceSlogan field value if set, zero value otherwise.
func (o *MiniAppServiceInfo) GetServiceSlogan() string {
	if o == nil || IsNil(o.ServiceSlogan) {
		var ret string
		return ret
	}
	return *o.ServiceSlogan
}

// GetServiceSloganOk returns a tuple with the ServiceSlogan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppServiceInfo) GetServiceSloganOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceSlogan) {
		return nil, false
	}
	return o.ServiceSlogan, true
}

// HasServiceSlogan returns a boolean if a field has been set.
func (o *MiniAppServiceInfo) HasServiceSlogan() bool {
	if o != nil && !IsNil(o.ServiceSlogan) {
		return true
	}

	return false
}

// SetServiceSlogan gets a reference to the given string and assigns it to the ServiceSlogan field.
func (o *MiniAppServiceInfo) SetServiceSlogan(v string) {
	o.ServiceSlogan = &v
}

// GetShowType returns the ShowType field value if set, zero value otherwise.
func (o *MiniAppServiceInfo) GetShowType() string {
	if o == nil || IsNil(o.ShowType) {
		var ret string
		return ret
	}
	return *o.ShowType
}

// GetShowTypeOk returns a tuple with the ShowType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppServiceInfo) GetShowTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ShowType) {
		return nil, false
	}
	return o.ShowType, true
}

// HasShowType returns a boolean if a field has been set.
func (o *MiniAppServiceInfo) HasShowType() bool {
	if o != nil && !IsNil(o.ShowType) {
		return true
	}

	return false
}

// SetShowType gets a reference to the given string and assigns it to the ShowType field.
func (o *MiniAppServiceInfo) SetShowType(v string) {
	o.ShowType = &v
}

func (o MiniAppServiceInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MiniAppServiceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BizStatus) {
		toSerialize["biz_status"] = o.BizStatus
	}
	if !IsNil(o.IsInner) {
		toSerialize["is_inner"] = o.IsInner
	}
	if !IsNil(o.IsOrder) {
		toSerialize["is_order"] = o.IsOrder
	}
	if !IsNil(o.IsvAppId) {
		toSerialize["isv_app_id"] = o.IsvAppId
	}
	if !IsNil(o.MiniAppId) {
		toSerialize["mini_app_id"] = o.MiniAppId
	}
	if !IsNil(o.MiniAppName) {
		toSerialize["mini_app_name"] = o.MiniAppName
	}
	if !IsNil(o.SellerId) {
		toSerialize["seller_id"] = o.SellerId
	}
	if !IsNil(o.SellerName) {
		toSerialize["seller_name"] = o.SellerName
	}
	if !IsNil(o.ServiceCode) {
		toSerialize["service_code"] = o.ServiceCode
	}
	if !IsNil(o.ServiceLogo) {
		toSerialize["service_logo"] = o.ServiceLogo
	}
	if !IsNil(o.ServiceName) {
		toSerialize["service_name"] = o.ServiceName
	}
	if !IsNil(o.ServiceSlogan) {
		toSerialize["service_slogan"] = o.ServiceSlogan
	}
	if !IsNil(o.ShowType) {
		toSerialize["show_type"] = o.ShowType
	}
	return toSerialize, nil
}

type NullableMiniAppServiceInfo struct {
	value *MiniAppServiceInfo
	isSet bool
}

func (v NullableMiniAppServiceInfo) Get() *MiniAppServiceInfo {
	return v.value
}

func (v *NullableMiniAppServiceInfo) Set(val *MiniAppServiceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMiniAppServiceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMiniAppServiceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMiniAppServiceInfo(val *MiniAppServiceInfo) *NullableMiniAppServiceInfo {
	return &NullableMiniAppServiceInfo{value: val, isSet: true}
}

func (v NullableMiniAppServiceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMiniAppServiceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


