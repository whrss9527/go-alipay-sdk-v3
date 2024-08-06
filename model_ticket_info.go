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

// checks if the TicketInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TicketInfo{}

// TicketInfo struct for TicketInfo
type TicketInfo struct {
	// 店铺地址
	Address *string `json:"address,omitempty"`
	// 截止时间
	EndTime *string `json:"end_time,omitempty"`
	// 凭证预约类型，一般不需要传入。可选值为INSTANT/RESERVATION，其中INSTANT代表是实时凭证，RESERVATION代表是预约凭证，不传入默认为实时凭证。区别在于预约凭证一般不是当场可取，而是用户下单后的很多天之后才可以凭借凭证提取
	ResvType *string `json:"resv_type,omitempty"`
	// 凭证可核销门店/货品自提门店，如果自提门店与购买门店不一致，可传入该字段提示用户自提门店
	Shop *string `json:"shop,omitempty"`
	// 单据号 不同类型下单据号含义不同。若类型为MEAL_NUM，则ticket_no表示取餐号；若类型为PICKUP_CODE，则ticket_no表示取件码；其他类型查看小程序订单产品文档
	TicketNo *string `json:"ticket_no,omitempty"`
	// 时间
	Time *string `json:"time,omitempty"`
	// 凭证类型 具体类型查看产品文档
	Type *string `json:"type,omitempty"`
}

// NewTicketInfo instantiates a new TicketInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTicketInfo() *TicketInfo {
	this := TicketInfo{}
	return &this
}

// NewTicketInfoWithDefaults instantiates a new TicketInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTicketInfoWithDefaults() *TicketInfo {
	this := TicketInfo{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *TicketInfo) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TicketInfo) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *TicketInfo) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *TicketInfo) SetAddress(v string) {
	o.Address = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *TicketInfo) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TicketInfo) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *TicketInfo) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *TicketInfo) SetEndTime(v string) {
	o.EndTime = &v
}

// GetResvType returns the ResvType field value if set, zero value otherwise.
func (o *TicketInfo) GetResvType() string {
	if o == nil || IsNil(o.ResvType) {
		var ret string
		return ret
	}
	return *o.ResvType
}

// GetResvTypeOk returns a tuple with the ResvType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TicketInfo) GetResvTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResvType) {
		return nil, false
	}
	return o.ResvType, true
}

// HasResvType returns a boolean if a field has been set.
func (o *TicketInfo) HasResvType() bool {
	if o != nil && !IsNil(o.ResvType) {
		return true
	}

	return false
}

// SetResvType gets a reference to the given string and assigns it to the ResvType field.
func (o *TicketInfo) SetResvType(v string) {
	o.ResvType = &v
}

// GetShop returns the Shop field value if set, zero value otherwise.
func (o *TicketInfo) GetShop() string {
	if o == nil || IsNil(o.Shop) {
		var ret string
		return ret
	}
	return *o.Shop
}

// GetShopOk returns a tuple with the Shop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TicketInfo) GetShopOk() (*string, bool) {
	if o == nil || IsNil(o.Shop) {
		return nil, false
	}
	return o.Shop, true
}

// HasShop returns a boolean if a field has been set.
func (o *TicketInfo) HasShop() bool {
	if o != nil && !IsNil(o.Shop) {
		return true
	}

	return false
}

// SetShop gets a reference to the given string and assigns it to the Shop field.
func (o *TicketInfo) SetShop(v string) {
	o.Shop = &v
}

// GetTicketNo returns the TicketNo field value if set, zero value otherwise.
func (o *TicketInfo) GetTicketNo() string {
	if o == nil || IsNil(o.TicketNo) {
		var ret string
		return ret
	}
	return *o.TicketNo
}

// GetTicketNoOk returns a tuple with the TicketNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TicketInfo) GetTicketNoOk() (*string, bool) {
	if o == nil || IsNil(o.TicketNo) {
		return nil, false
	}
	return o.TicketNo, true
}

// HasTicketNo returns a boolean if a field has been set.
func (o *TicketInfo) HasTicketNo() bool {
	if o != nil && !IsNil(o.TicketNo) {
		return true
	}

	return false
}

// SetTicketNo gets a reference to the given string and assigns it to the TicketNo field.
func (o *TicketInfo) SetTicketNo(v string) {
	o.TicketNo = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *TicketInfo) GetTime() string {
	if o == nil || IsNil(o.Time) {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TicketInfo) GetTimeOk() (*string, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *TicketInfo) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *TicketInfo) SetTime(v string) {
	o.Time = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TicketInfo) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TicketInfo) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TicketInfo) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TicketInfo) SetType(v string) {
	o.Type = &v
}

func (o TicketInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TicketInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.ResvType) {
		toSerialize["resv_type"] = o.ResvType
	}
	if !IsNil(o.Shop) {
		toSerialize["shop"] = o.Shop
	}
	if !IsNil(o.TicketNo) {
		toSerialize["ticket_no"] = o.TicketNo
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableTicketInfo struct {
	value *TicketInfo
	isSet bool
}

func (v NullableTicketInfo) Get() *TicketInfo {
	return v.value
}

func (v *NullableTicketInfo) Set(val *TicketInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTicketInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTicketInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTicketInfo(val *TicketInfo) *NullableTicketInfo {
	return &NullableTicketInfo{value: val, isSet: true}
}

func (v NullableTicketInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTicketInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

