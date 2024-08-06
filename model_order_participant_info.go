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

// checks if the OrderParticipantInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderParticipantInfo{}

// OrderParticipantInfo struct for OrderParticipantInfo
type OrderParticipantInfo struct {
	// 扩展信息
	ExtInfo []OrderExtInfo `json:"ext_info,omitempty"`
	// 参与方logo素材id 开发者通过 alipay.merchant.item.file.upload(商品文件上传接口)上传图片，获取到对应的素材 ID( material_id )
	LogoMaterialId *string `json:"logo_material_id,omitempty"`
	// 参与方名称
	Name *string `json:"name,omitempty"`
	// 参与方简称
	ShortName *string `json:"short_name,omitempty"`
	// 参与方类型
	Type *string `json:"type,omitempty"`
	// 参与方支付宝uid
	Uid *string `json:"uid,omitempty"`
}

// NewOrderParticipantInfo instantiates a new OrderParticipantInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderParticipantInfo() *OrderParticipantInfo {
	this := OrderParticipantInfo{}
	return &this
}

// NewOrderParticipantInfoWithDefaults instantiates a new OrderParticipantInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderParticipantInfoWithDefaults() *OrderParticipantInfo {
	this := OrderParticipantInfo{}
	return &this
}

// GetExtInfo returns the ExtInfo field value if set, zero value otherwise.
func (o *OrderParticipantInfo) GetExtInfo() []OrderExtInfo {
	if o == nil || IsNil(o.ExtInfo) {
		var ret []OrderExtInfo
		return ret
	}
	return o.ExtInfo
}

// GetExtInfoOk returns a tuple with the ExtInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderParticipantInfo) GetExtInfoOk() ([]OrderExtInfo, bool) {
	if o == nil || IsNil(o.ExtInfo) {
		return nil, false
	}
	return o.ExtInfo, true
}

// HasExtInfo returns a boolean if a field has been set.
func (o *OrderParticipantInfo) HasExtInfo() bool {
	if o != nil && !IsNil(o.ExtInfo) {
		return true
	}

	return false
}

// SetExtInfo gets a reference to the given []OrderExtInfo and assigns it to the ExtInfo field.
func (o *OrderParticipantInfo) SetExtInfo(v []OrderExtInfo) {
	o.ExtInfo = v
}

// GetLogoMaterialId returns the LogoMaterialId field value if set, zero value otherwise.
func (o *OrderParticipantInfo) GetLogoMaterialId() string {
	if o == nil || IsNil(o.LogoMaterialId) {
		var ret string
		return ret
	}
	return *o.LogoMaterialId
}

// GetLogoMaterialIdOk returns a tuple with the LogoMaterialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderParticipantInfo) GetLogoMaterialIdOk() (*string, bool) {
	if o == nil || IsNil(o.LogoMaterialId) {
		return nil, false
	}
	return o.LogoMaterialId, true
}

// HasLogoMaterialId returns a boolean if a field has been set.
func (o *OrderParticipantInfo) HasLogoMaterialId() bool {
	if o != nil && !IsNil(o.LogoMaterialId) {
		return true
	}

	return false
}

// SetLogoMaterialId gets a reference to the given string and assigns it to the LogoMaterialId field.
func (o *OrderParticipantInfo) SetLogoMaterialId(v string) {
	o.LogoMaterialId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrderParticipantInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderParticipantInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrderParticipantInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrderParticipantInfo) SetName(v string) {
	o.Name = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *OrderParticipantInfo) GetShortName() string {
	if o == nil || IsNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderParticipantInfo) GetShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShortName) {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *OrderParticipantInfo) HasShortName() bool {
	if o != nil && !IsNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *OrderParticipantInfo) SetShortName(v string) {
	o.ShortName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrderParticipantInfo) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderParticipantInfo) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrderParticipantInfo) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrderParticipantInfo) SetType(v string) {
	o.Type = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *OrderParticipantInfo) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderParticipantInfo) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *OrderParticipantInfo) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *OrderParticipantInfo) SetUid(v string) {
	o.Uid = &v
}

func (o OrderParticipantInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderParticipantInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExtInfo) {
		toSerialize["ext_info"] = o.ExtInfo
	}
	if !IsNil(o.LogoMaterialId) {
		toSerialize["logo_material_id"] = o.LogoMaterialId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ShortName) {
		toSerialize["short_name"] = o.ShortName
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	return toSerialize, nil
}

type NullableOrderParticipantInfo struct {
	value *OrderParticipantInfo
	isSet bool
}

func (v NullableOrderParticipantInfo) Get() *OrderParticipantInfo {
	return v.value
}

func (v *NullableOrderParticipantInfo) Set(val *OrderParticipantInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderParticipantInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderParticipantInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderParticipantInfo(val *OrderParticipantInfo) *NullableOrderParticipantInfo {
	return &NullableOrderParticipantInfo{value: val, isSet: true}
}

func (v NullableOrderParticipantInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderParticipantInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


