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

// checks if the ActivityDiscountVoucher type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivityDiscountVoucher{}

// ActivityDiscountVoucher struct for ActivityDiscountVoucher
type ActivityDiscountVoucher struct {
	// 封顶金额。
	CeilingAmount *string `json:"ceiling_amount,omitempty"`
	// 折扣率。
	Discount *string `json:"discount,omitempty"`
	// 门槛金额。说明：该字段可不填，认为无门槛;
	FloorAmount *string `json:"floor_amount,omitempty"`
	// 商品名称。说明：该字段可不填，填入 origin_amount 则必填。
	GoodsName *string `json:"goods_name,omitempty"`
	// 原价。说明：该字段可不填，填入商品名称goods_name则必填;
	OriginAmount *string `json:"origin_amount,omitempty"`
}

// NewActivityDiscountVoucher instantiates a new ActivityDiscountVoucher object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivityDiscountVoucher() *ActivityDiscountVoucher {
	this := ActivityDiscountVoucher{}
	return &this
}

// NewActivityDiscountVoucherWithDefaults instantiates a new ActivityDiscountVoucher object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityDiscountVoucherWithDefaults() *ActivityDiscountVoucher {
	this := ActivityDiscountVoucher{}
	return &this
}

// GetCeilingAmount returns the CeilingAmount field value if set, zero value otherwise.
func (o *ActivityDiscountVoucher) GetCeilingAmount() string {
	if o == nil || IsNil(o.CeilingAmount) {
		var ret string
		return ret
	}
	return *o.CeilingAmount
}

// GetCeilingAmountOk returns a tuple with the CeilingAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityDiscountVoucher) GetCeilingAmountOk() (*string, bool) {
	if o == nil || IsNil(o.CeilingAmount) {
		return nil, false
	}
	return o.CeilingAmount, true
}

// HasCeilingAmount returns a boolean if a field has been set.
func (o *ActivityDiscountVoucher) HasCeilingAmount() bool {
	if o != nil && !IsNil(o.CeilingAmount) {
		return true
	}

	return false
}

// SetCeilingAmount gets a reference to the given string and assigns it to the CeilingAmount field.
func (o *ActivityDiscountVoucher) SetCeilingAmount(v string) {
	o.CeilingAmount = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *ActivityDiscountVoucher) GetDiscount() string {
	if o == nil || IsNil(o.Discount) {
		var ret string
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityDiscountVoucher) GetDiscountOk() (*string, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *ActivityDiscountVoucher) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given string and assigns it to the Discount field.
func (o *ActivityDiscountVoucher) SetDiscount(v string) {
	o.Discount = &v
}

// GetFloorAmount returns the FloorAmount field value if set, zero value otherwise.
func (o *ActivityDiscountVoucher) GetFloorAmount() string {
	if o == nil || IsNil(o.FloorAmount) {
		var ret string
		return ret
	}
	return *o.FloorAmount
}

// GetFloorAmountOk returns a tuple with the FloorAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityDiscountVoucher) GetFloorAmountOk() (*string, bool) {
	if o == nil || IsNil(o.FloorAmount) {
		return nil, false
	}
	return o.FloorAmount, true
}

// HasFloorAmount returns a boolean if a field has been set.
func (o *ActivityDiscountVoucher) HasFloorAmount() bool {
	if o != nil && !IsNil(o.FloorAmount) {
		return true
	}

	return false
}

// SetFloorAmount gets a reference to the given string and assigns it to the FloorAmount field.
func (o *ActivityDiscountVoucher) SetFloorAmount(v string) {
	o.FloorAmount = &v
}

// GetGoodsName returns the GoodsName field value if set, zero value otherwise.
func (o *ActivityDiscountVoucher) GetGoodsName() string {
	if o == nil || IsNil(o.GoodsName) {
		var ret string
		return ret
	}
	return *o.GoodsName
}

// GetGoodsNameOk returns a tuple with the GoodsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityDiscountVoucher) GetGoodsNameOk() (*string, bool) {
	if o == nil || IsNil(o.GoodsName) {
		return nil, false
	}
	return o.GoodsName, true
}

// HasGoodsName returns a boolean if a field has been set.
func (o *ActivityDiscountVoucher) HasGoodsName() bool {
	if o != nil && !IsNil(o.GoodsName) {
		return true
	}

	return false
}

// SetGoodsName gets a reference to the given string and assigns it to the GoodsName field.
func (o *ActivityDiscountVoucher) SetGoodsName(v string) {
	o.GoodsName = &v
}

// GetOriginAmount returns the OriginAmount field value if set, zero value otherwise.
func (o *ActivityDiscountVoucher) GetOriginAmount() string {
	if o == nil || IsNil(o.OriginAmount) {
		var ret string
		return ret
	}
	return *o.OriginAmount
}

// GetOriginAmountOk returns a tuple with the OriginAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityDiscountVoucher) GetOriginAmountOk() (*string, bool) {
	if o == nil || IsNil(o.OriginAmount) {
		return nil, false
	}
	return o.OriginAmount, true
}

// HasOriginAmount returns a boolean if a field has been set.
func (o *ActivityDiscountVoucher) HasOriginAmount() bool {
	if o != nil && !IsNil(o.OriginAmount) {
		return true
	}

	return false
}

// SetOriginAmount gets a reference to the given string and assigns it to the OriginAmount field.
func (o *ActivityDiscountVoucher) SetOriginAmount(v string) {
	o.OriginAmount = &v
}

func (o ActivityDiscountVoucher) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivityDiscountVoucher) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CeilingAmount) {
		toSerialize["ceiling_amount"] = o.CeilingAmount
	}
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	if !IsNil(o.FloorAmount) {
		toSerialize["floor_amount"] = o.FloorAmount
	}
	if !IsNil(o.GoodsName) {
		toSerialize["goods_name"] = o.GoodsName
	}
	if !IsNil(o.OriginAmount) {
		toSerialize["origin_amount"] = o.OriginAmount
	}
	return toSerialize, nil
}

type NullableActivityDiscountVoucher struct {
	value *ActivityDiscountVoucher
	isSet bool
}

func (v NullableActivityDiscountVoucher) Get() *ActivityDiscountVoucher {
	return v.value
}

func (v *NullableActivityDiscountVoucher) Set(val *ActivityDiscountVoucher) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityDiscountVoucher) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityDiscountVoucher) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityDiscountVoucher(val *ActivityDiscountVoucher) *NullableActivityDiscountVoucher {
	return &NullableActivityDiscountVoucher{value: val, isSet: true}
}

func (v NullableActivityDiscountVoucher) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityDiscountVoucher) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
