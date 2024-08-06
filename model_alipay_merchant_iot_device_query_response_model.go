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

// checks if the AlipayMerchantIotDeviceQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMerchantIotDeviceQueryResponseModel{}

// AlipayMerchantIotDeviceQueryResponseModel struct for AlipayMerchantIotDeviceQueryResponseModel
type AlipayMerchantIotDeviceQueryResponseModel struct {
	// 商户类型。direct表示直连商户，indirect表示间连商户。
	MerchantType *string `json:"merchant_type,omitempty"`
	// 设备绑定的商户支付宝pid。直连场景为商户收单pid，间连场景为商户smid绑定的pid。
	Pid *string `json:"pid,omitempty"`
	// 设备绑定的门店ID。
	ShopId *string `json:"shop_id,omitempty"`
	// 商户smid。直连场景为空。
	Smid *string `json:"smid,omitempty"`
}

// NewAlipayMerchantIotDeviceQueryResponseModel instantiates a new AlipayMerchantIotDeviceQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMerchantIotDeviceQueryResponseModel() *AlipayMerchantIotDeviceQueryResponseModel {
	this := AlipayMerchantIotDeviceQueryResponseModel{}
	return &this
}

// NewAlipayMerchantIotDeviceQueryResponseModelWithDefaults instantiates a new AlipayMerchantIotDeviceQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMerchantIotDeviceQueryResponseModelWithDefaults() *AlipayMerchantIotDeviceQueryResponseModel {
	this := AlipayMerchantIotDeviceQueryResponseModel{}
	return &this
}

// GetMerchantType returns the MerchantType field value if set, zero value otherwise.
func (o *AlipayMerchantIotDeviceQueryResponseModel) GetMerchantType() string {
	if o == nil || IsNil(o.MerchantType) {
		var ret string
		return ret
	}
	return *o.MerchantType
}

// GetMerchantTypeOk returns a tuple with the MerchantType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantIotDeviceQueryResponseModel) GetMerchantTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantType) {
		return nil, false
	}
	return o.MerchantType, true
}

// HasMerchantType returns a boolean if a field has been set.
func (o *AlipayMerchantIotDeviceQueryResponseModel) HasMerchantType() bool {
	if o != nil && !IsNil(o.MerchantType) {
		return true
	}

	return false
}

// SetMerchantType gets a reference to the given string and assigns it to the MerchantType field.
func (o *AlipayMerchantIotDeviceQueryResponseModel) SetMerchantType(v string) {
	o.MerchantType = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *AlipayMerchantIotDeviceQueryResponseModel) GetPid() string {
	if o == nil || IsNil(o.Pid) {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantIotDeviceQueryResponseModel) GetPidOk() (*string, bool) {
	if o == nil || IsNil(o.Pid) {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *AlipayMerchantIotDeviceQueryResponseModel) HasPid() bool {
	if o != nil && !IsNil(o.Pid) {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *AlipayMerchantIotDeviceQueryResponseModel) SetPid(v string) {
	o.Pid = &v
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *AlipayMerchantIotDeviceQueryResponseModel) GetShopId() string {
	if o == nil || IsNil(o.ShopId) {
		var ret string
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantIotDeviceQueryResponseModel) GetShopIdOk() (*string, bool) {
	if o == nil || IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *AlipayMerchantIotDeviceQueryResponseModel) HasShopId() bool {
	if o != nil && !IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given string and assigns it to the ShopId field.
func (o *AlipayMerchantIotDeviceQueryResponseModel) SetShopId(v string) {
	o.ShopId = &v
}

// GetSmid returns the Smid field value if set, zero value otherwise.
func (o *AlipayMerchantIotDeviceQueryResponseModel) GetSmid() string {
	if o == nil || IsNil(o.Smid) {
		var ret string
		return ret
	}
	return *o.Smid
}

// GetSmidOk returns a tuple with the Smid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantIotDeviceQueryResponseModel) GetSmidOk() (*string, bool) {
	if o == nil || IsNil(o.Smid) {
		return nil, false
	}
	return o.Smid, true
}

// HasSmid returns a boolean if a field has been set.
func (o *AlipayMerchantIotDeviceQueryResponseModel) HasSmid() bool {
	if o != nil && !IsNil(o.Smid) {
		return true
	}

	return false
}

// SetSmid gets a reference to the given string and assigns it to the Smid field.
func (o *AlipayMerchantIotDeviceQueryResponseModel) SetSmid(v string) {
	o.Smid = &v
}

func (o AlipayMerchantIotDeviceQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMerchantIotDeviceQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantType) {
		toSerialize["merchant_type"] = o.MerchantType
	}
	if !IsNil(o.Pid) {
		toSerialize["pid"] = o.Pid
	}
	if !IsNil(o.ShopId) {
		toSerialize["shop_id"] = o.ShopId
	}
	if !IsNil(o.Smid) {
		toSerialize["smid"] = o.Smid
	}
	return toSerialize, nil
}

type NullableAlipayMerchantIotDeviceQueryResponseModel struct {
	value *AlipayMerchantIotDeviceQueryResponseModel
	isSet bool
}

func (v NullableAlipayMerchantIotDeviceQueryResponseModel) Get() *AlipayMerchantIotDeviceQueryResponseModel {
	return v.value
}

func (v *NullableAlipayMerchantIotDeviceQueryResponseModel) Set(val *AlipayMerchantIotDeviceQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMerchantIotDeviceQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMerchantIotDeviceQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMerchantIotDeviceQueryResponseModel(val *AlipayMerchantIotDeviceQueryResponseModel) *NullableAlipayMerchantIotDeviceQueryResponseModel {
	return &NullableAlipayMerchantIotDeviceQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMerchantIotDeviceQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMerchantIotDeviceQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
