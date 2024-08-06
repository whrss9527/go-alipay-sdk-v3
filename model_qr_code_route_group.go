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

// checks if the QrCodeRouteGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QrCodeRouteGroup{}

// QrCodeRouteGroup struct for QrCodeRouteGroup
type QrCodeRouteGroup struct {
	// 匹配规则，EXACT（精确匹配）、FUZZY（模糊匹配）、PATTERN（模式匹配）\\\\（如：配置二维码地址为https://www.alipay.com/my?id=123，当用户扫这个地址的二维码可唤起小程序）。 模糊匹配：根据填写的二维码地址模糊匹配，只要地址前缀匹配即可唤起小程序（如：配置二维码地址为https://www.alipay.com/my/，当用户扫的二维码地址为https://www.alipay.com/my/id=123,可唤起小程序）。 。 模式匹配：根据填写的二维码地址可变变量进行匹配，只要地址的变量位置自定义，变量之外的部分匹配即可唤起小程序(如：配置的二维码地址为https://www.alipay.com/{0}/my/{1},当用户扫的二维码地址为：https://www.alipay.com/user/my/scan，可唤起小程序)
	Mode *string `json:"mode,omitempty"`
	// 路由组id（参数说明：该参数可用于alipay.open.mini.qrcode.unbind接口入参route_group，进行二维码解绑）
	RouteGroup *string `json:"route_group,omitempty"`
	// 规则路由地址
	RouteUrl *string `json:"route_url,omitempty"`
}

// NewQrCodeRouteGroup instantiates a new QrCodeRouteGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQrCodeRouteGroup() *QrCodeRouteGroup {
	this := QrCodeRouteGroup{}
	return &this
}

// NewQrCodeRouteGroupWithDefaults instantiates a new QrCodeRouteGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQrCodeRouteGroupWithDefaults() *QrCodeRouteGroup {
	this := QrCodeRouteGroup{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *QrCodeRouteGroup) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QrCodeRouteGroup) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *QrCodeRouteGroup) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *QrCodeRouteGroup) SetMode(v string) {
	o.Mode = &v
}

// GetRouteGroup returns the RouteGroup field value if set, zero value otherwise.
func (o *QrCodeRouteGroup) GetRouteGroup() string {
	if o == nil || IsNil(o.RouteGroup) {
		var ret string
		return ret
	}
	return *o.RouteGroup
}

// GetRouteGroupOk returns a tuple with the RouteGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QrCodeRouteGroup) GetRouteGroupOk() (*string, bool) {
	if o == nil || IsNil(o.RouteGroup) {
		return nil, false
	}
	return o.RouteGroup, true
}

// HasRouteGroup returns a boolean if a field has been set.
func (o *QrCodeRouteGroup) HasRouteGroup() bool {
	if o != nil && !IsNil(o.RouteGroup) {
		return true
	}

	return false
}

// SetRouteGroup gets a reference to the given string and assigns it to the RouteGroup field.
func (o *QrCodeRouteGroup) SetRouteGroup(v string) {
	o.RouteGroup = &v
}

// GetRouteUrl returns the RouteUrl field value if set, zero value otherwise.
func (o *QrCodeRouteGroup) GetRouteUrl() string {
	if o == nil || IsNil(o.RouteUrl) {
		var ret string
		return ret
	}
	return *o.RouteUrl
}

// GetRouteUrlOk returns a tuple with the RouteUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QrCodeRouteGroup) GetRouteUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RouteUrl) {
		return nil, false
	}
	return o.RouteUrl, true
}

// HasRouteUrl returns a boolean if a field has been set.
func (o *QrCodeRouteGroup) HasRouteUrl() bool {
	if o != nil && !IsNil(o.RouteUrl) {
		return true
	}

	return false
}

// SetRouteUrl gets a reference to the given string and assigns it to the RouteUrl field.
func (o *QrCodeRouteGroup) SetRouteUrl(v string) {
	o.RouteUrl = &v
}

func (o QrCodeRouteGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QrCodeRouteGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.RouteGroup) {
		toSerialize["route_group"] = o.RouteGroup
	}
	if !IsNil(o.RouteUrl) {
		toSerialize["route_url"] = o.RouteUrl
	}
	return toSerialize, nil
}

type NullableQrCodeRouteGroup struct {
	value *QrCodeRouteGroup
	isSet bool
}

func (v NullableQrCodeRouteGroup) Get() *QrCodeRouteGroup {
	return v.value
}

func (v *NullableQrCodeRouteGroup) Set(val *QrCodeRouteGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableQrCodeRouteGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableQrCodeRouteGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQrCodeRouteGroup(val *QrCodeRouteGroup) *NullableQrCodeRouteGroup {
	return &NullableQrCodeRouteGroup{value: val, isSet: true}
}

func (v NullableQrCodeRouteGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQrCodeRouteGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

