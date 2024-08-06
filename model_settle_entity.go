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

// checks if the SettleEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettleEntity{}

// SettleEntity struct for SettleEntity
type SettleEntity struct {
	// 结算主体资产业务类型，settled 表示已结算资产
	SettleEntityBizType *string `json:"settle_entity_biz_type,omitempty"`
	// 结算主体账号 当结算主体类型为SecondMerchant，本参数为二级商户的SecondMerchantID；当结算主体类型为MerchantStore，本参数为商户门店ID。
	SettleEntityId *string `json:"settle_entity_id,omitempty"`
	// 结算主体类型 SecondMerchant：结算主体为二级商户；Store：结算主体为商户门店
	SettleEntityType *string `json:"settle_entity_type,omitempty"`
}

// NewSettleEntity instantiates a new SettleEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettleEntity() *SettleEntity {
	this := SettleEntity{}
	return &this
}

// NewSettleEntityWithDefaults instantiates a new SettleEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettleEntityWithDefaults() *SettleEntity {
	this := SettleEntity{}
	return &this
}

// GetSettleEntityBizType returns the SettleEntityBizType field value if set, zero value otherwise.
func (o *SettleEntity) GetSettleEntityBizType() string {
	if o == nil || IsNil(o.SettleEntityBizType) {
		var ret string
		return ret
	}
	return *o.SettleEntityBizType
}

// GetSettleEntityBizTypeOk returns a tuple with the SettleEntityBizType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettleEntity) GetSettleEntityBizTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SettleEntityBizType) {
		return nil, false
	}
	return o.SettleEntityBizType, true
}

// HasSettleEntityBizType returns a boolean if a field has been set.
func (o *SettleEntity) HasSettleEntityBizType() bool {
	if o != nil && !IsNil(o.SettleEntityBizType) {
		return true
	}

	return false
}

// SetSettleEntityBizType gets a reference to the given string and assigns it to the SettleEntityBizType field.
func (o *SettleEntity) SetSettleEntityBizType(v string) {
	o.SettleEntityBizType = &v
}

// GetSettleEntityId returns the SettleEntityId field value if set, zero value otherwise.
func (o *SettleEntity) GetSettleEntityId() string {
	if o == nil || IsNil(o.SettleEntityId) {
		var ret string
		return ret
	}
	return *o.SettleEntityId
}

// GetSettleEntityIdOk returns a tuple with the SettleEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettleEntity) GetSettleEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.SettleEntityId) {
		return nil, false
	}
	return o.SettleEntityId, true
}

// HasSettleEntityId returns a boolean if a field has been set.
func (o *SettleEntity) HasSettleEntityId() bool {
	if o != nil && !IsNil(o.SettleEntityId) {
		return true
	}

	return false
}

// SetSettleEntityId gets a reference to the given string and assigns it to the SettleEntityId field.
func (o *SettleEntity) SetSettleEntityId(v string) {
	o.SettleEntityId = &v
}

// GetSettleEntityType returns the SettleEntityType field value if set, zero value otherwise.
func (o *SettleEntity) GetSettleEntityType() string {
	if o == nil || IsNil(o.SettleEntityType) {
		var ret string
		return ret
	}
	return *o.SettleEntityType
}

// GetSettleEntityTypeOk returns a tuple with the SettleEntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettleEntity) GetSettleEntityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SettleEntityType) {
		return nil, false
	}
	return o.SettleEntityType, true
}

// HasSettleEntityType returns a boolean if a field has been set.
func (o *SettleEntity) HasSettleEntityType() bool {
	if o != nil && !IsNil(o.SettleEntityType) {
		return true
	}

	return false
}

// SetSettleEntityType gets a reference to the given string and assigns it to the SettleEntityType field.
func (o *SettleEntity) SetSettleEntityType(v string) {
	o.SettleEntityType = &v
}

func (o SettleEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettleEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SettleEntityBizType) {
		toSerialize["settle_entity_biz_type"] = o.SettleEntityBizType
	}
	if !IsNil(o.SettleEntityId) {
		toSerialize["settle_entity_id"] = o.SettleEntityId
	}
	if !IsNil(o.SettleEntityType) {
		toSerialize["settle_entity_type"] = o.SettleEntityType
	}
	return toSerialize, nil
}

type NullableSettleEntity struct {
	value *SettleEntity
	isSet bool
}

func (v NullableSettleEntity) Get() *SettleEntity {
	return v.value
}

func (v *NullableSettleEntity) Set(val *SettleEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableSettleEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableSettleEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettleEntity(val *SettleEntity) *NullableSettleEntity {
	return &NullableSettleEntity{value: val, isSet: true}
}

func (v NullableSettleEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettleEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


