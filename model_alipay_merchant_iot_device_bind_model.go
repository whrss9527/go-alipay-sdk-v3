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

// checks if the AlipayMerchantIotDeviceBindModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMerchantIotDeviceBindModel{}

// AlipayMerchantIotDeviceBindModel struct for AlipayMerchantIotDeviceBindModel
type AlipayMerchantIotDeviceBindModel struct {
	// 设备 ID ，仅device_id_type 为 ID 时填写。
	BizTid *string `json:"biz_tid,omitempty"`
	// 可选方式。枚举支持： *  ID：表示使用 biztid 作为设备唯一识别标识， *  SN：表示使用 supplier_id、device_sn联合作为设备唯一识别标识。 注意：由于不同机型的supplier_id不同，推荐使用 ID 。
	DeviceIdType *string `json:"device_id_type,omitempty"`
	// 设备序列号 ，device_id_type 为 SN 时填写。需配合supplier_id使用。
	DeviceSn *string `json:"device_sn,omitempty"`
	// 商户类型。枚举支持： *  direct：直连商户； *  indirect：间连商户。
	MerchantType *string `json:"merchant_type,omitempty"`
	// 直连场景填写商户收单pid。间连场景填写smid绑定的pid（可通过代运营授权API获取）。
	Pid *string `json:"pid,omitempty"`
	// 商户pid名下的门店 ID。可通过门店管理API获取，需注意请勿使用smid创建门店 ID。
	ShopId *string `json:"shop_id,omitempty"`
	// 直连场景不填，间连场景填写商户收单smid
	Smid *string `json:"smid,omitempty"`
	// 设备供应商ID ，device_id_type 为 SN 时填写。需注意不同机型的供应商ID可能不同。
	SupplierId *string `json:"supplier_id,omitempty"`
}

// NewAlipayMerchantIotDeviceBindModel instantiates a new AlipayMerchantIotDeviceBindModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMerchantIotDeviceBindModel() *AlipayMerchantIotDeviceBindModel {
	this := AlipayMerchantIotDeviceBindModel{}
	return &this
}

// NewAlipayMerchantIotDeviceBindModelWithDefaults instantiates a new AlipayMerchantIotDeviceBindModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMerchantIotDeviceBindModelWithDefaults() *AlipayMerchantIotDeviceBindModel {
	this := AlipayMerchantIotDeviceBindModel{}
	return &this
}

// GetBizTid returns the BizTid field value if set, zero value otherwise.
func (o *AlipayMerchantIotDeviceBindModel) GetBizTid() string {
	if o == nil || IsNil(o.BizTid) {
		var ret string
		return ret
	}
	return *o.BizTid
}

// GetBizTidOk returns a tuple with the BizTid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantIotDeviceBindModel) GetBizTidOk() (*string, bool) {
	if o == nil || IsNil(o.BizTid) {
		return nil, false
	}
	return o.BizTid, true
}

// HasBizTid returns a boolean if a field has been set.
func (o *AlipayMerchantIotDeviceBindModel) HasBizTid() bool {
	if o != nil && !IsNil(o.BizTid) {
		return true
	}

	return false
}

// SetBizTid gets a reference to the given string and assigns it to the BizTid field.
func (o *AlipayMerchantIotDeviceBindModel) SetBizTid(v string) {
	o.BizTid = &v
}

// GetDeviceIdType returns the DeviceIdType field value if set, zero value otherwise.
func (o *AlipayMerchantIotDeviceBindModel) GetDeviceIdType() string {
	if o == nil || IsNil(o.DeviceIdType) {
		var ret string
		return ret
	}
	return *o.DeviceIdType
}

// GetDeviceIdTypeOk returns a tuple with the DeviceIdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantIotDeviceBindModel) GetDeviceIdTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceIdType) {
		return nil, false
	}
	return o.DeviceIdType, true
}

// HasDeviceIdType returns a boolean if a field has been set.
func (o *AlipayMerchantIotDeviceBindModel) HasDeviceIdType() bool {
	if o != nil && !IsNil(o.DeviceIdType) {
		return true
	}

	return false
}

// SetDeviceIdType gets a reference to the given string and assigns it to the DeviceIdType field.
func (o *AlipayMerchantIotDeviceBindModel) SetDeviceIdType(v string) {
	o.DeviceIdType = &v
}

// GetDeviceSn returns the DeviceSn field value if set, zero value otherwise.
func (o *AlipayMerchantIotDeviceBindModel) GetDeviceSn() string {
	if o == nil || IsNil(o.DeviceSn) {
		var ret string
		return ret
	}
	return *o.DeviceSn
}

// GetDeviceSnOk returns a tuple with the DeviceSn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantIotDeviceBindModel) GetDeviceSnOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceSn) {
		return nil, false
	}
	return o.DeviceSn, true
}

// HasDeviceSn returns a boolean if a field has been set.
func (o *AlipayMerchantIotDeviceBindModel) HasDeviceSn() bool {
	if o != nil && !IsNil(o.DeviceSn) {
		return true
	}

	return false
}

// SetDeviceSn gets a reference to the given string and assigns it to the DeviceSn field.
func (o *AlipayMerchantIotDeviceBindModel) SetDeviceSn(v string) {
	o.DeviceSn = &v
}

// GetMerchantType returns the MerchantType field value if set, zero value otherwise.
func (o *AlipayMerchantIotDeviceBindModel) GetMerchantType() string {
	if o == nil || IsNil(o.MerchantType) {
		var ret string
		return ret
	}
	return *o.MerchantType
}

// GetMerchantTypeOk returns a tuple with the MerchantType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantIotDeviceBindModel) GetMerchantTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantType) {
		return nil, false
	}
	return o.MerchantType, true
}

// HasMerchantType returns a boolean if a field has been set.
func (o *AlipayMerchantIotDeviceBindModel) HasMerchantType() bool {
	if o != nil && !IsNil(o.MerchantType) {
		return true
	}

	return false
}

// SetMerchantType gets a reference to the given string and assigns it to the MerchantType field.
func (o *AlipayMerchantIotDeviceBindModel) SetMerchantType(v string) {
	o.MerchantType = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *AlipayMerchantIotDeviceBindModel) GetPid() string {
	if o == nil || IsNil(o.Pid) {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantIotDeviceBindModel) GetPidOk() (*string, bool) {
	if o == nil || IsNil(o.Pid) {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *AlipayMerchantIotDeviceBindModel) HasPid() bool {
	if o != nil && !IsNil(o.Pid) {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *AlipayMerchantIotDeviceBindModel) SetPid(v string) {
	o.Pid = &v
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *AlipayMerchantIotDeviceBindModel) GetShopId() string {
	if o == nil || IsNil(o.ShopId) {
		var ret string
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantIotDeviceBindModel) GetShopIdOk() (*string, bool) {
	if o == nil || IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *AlipayMerchantIotDeviceBindModel) HasShopId() bool {
	if o != nil && !IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given string and assigns it to the ShopId field.
func (o *AlipayMerchantIotDeviceBindModel) SetShopId(v string) {
	o.ShopId = &v
}

// GetSmid returns the Smid field value if set, zero value otherwise.
func (o *AlipayMerchantIotDeviceBindModel) GetSmid() string {
	if o == nil || IsNil(o.Smid) {
		var ret string
		return ret
	}
	return *o.Smid
}

// GetSmidOk returns a tuple with the Smid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantIotDeviceBindModel) GetSmidOk() (*string, bool) {
	if o == nil || IsNil(o.Smid) {
		return nil, false
	}
	return o.Smid, true
}

// HasSmid returns a boolean if a field has been set.
func (o *AlipayMerchantIotDeviceBindModel) HasSmid() bool {
	if o != nil && !IsNil(o.Smid) {
		return true
	}

	return false
}

// SetSmid gets a reference to the given string and assigns it to the Smid field.
func (o *AlipayMerchantIotDeviceBindModel) SetSmid(v string) {
	o.Smid = &v
}

// GetSupplierId returns the SupplierId field value if set, zero value otherwise.
func (o *AlipayMerchantIotDeviceBindModel) GetSupplierId() string {
	if o == nil || IsNil(o.SupplierId) {
		var ret string
		return ret
	}
	return *o.SupplierId
}

// GetSupplierIdOk returns a tuple with the SupplierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantIotDeviceBindModel) GetSupplierIdOk() (*string, bool) {
	if o == nil || IsNil(o.SupplierId) {
		return nil, false
	}
	return o.SupplierId, true
}

// HasSupplierId returns a boolean if a field has been set.
func (o *AlipayMerchantIotDeviceBindModel) HasSupplierId() bool {
	if o != nil && !IsNil(o.SupplierId) {
		return true
	}

	return false
}

// SetSupplierId gets a reference to the given string and assigns it to the SupplierId field.
func (o *AlipayMerchantIotDeviceBindModel) SetSupplierId(v string) {
	o.SupplierId = &v
}

func (o AlipayMerchantIotDeviceBindModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMerchantIotDeviceBindModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BizTid) {
		toSerialize["biz_tid"] = o.BizTid
	}
	if !IsNil(o.DeviceIdType) {
		toSerialize["device_id_type"] = o.DeviceIdType
	}
	if !IsNil(o.DeviceSn) {
		toSerialize["device_sn"] = o.DeviceSn
	}
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
	if !IsNil(o.SupplierId) {
		toSerialize["supplier_id"] = o.SupplierId
	}
	return toSerialize, nil
}

type NullableAlipayMerchantIotDeviceBindModel struct {
	value *AlipayMerchantIotDeviceBindModel
	isSet bool
}

func (v NullableAlipayMerchantIotDeviceBindModel) Get() *AlipayMerchantIotDeviceBindModel {
	return v.value
}

func (v *NullableAlipayMerchantIotDeviceBindModel) Set(val *AlipayMerchantIotDeviceBindModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMerchantIotDeviceBindModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMerchantIotDeviceBindModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMerchantIotDeviceBindModel(val *AlipayMerchantIotDeviceBindModel) *NullableAlipayMerchantIotDeviceBindModel {
	return &NullableAlipayMerchantIotDeviceBindModel{value: val, isSet: true}
}

func (v NullableAlipayMerchantIotDeviceBindModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMerchantIotDeviceBindModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


