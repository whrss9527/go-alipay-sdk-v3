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

// checks if the SkuCreateInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkuCreateInfo{}

// SkuCreateInfo struct for SkuCreateInfo
type SkuCreateInfo struct {
	// 库存
	Inventory *int32 `json:"inventory,omitempty"`
	// SKU素材列表（最多3个）
	MaterialList []MaterialCreateInfo `json:"material_list,omitempty"`
	// 标价，单位分
	OriginalPrice *int32 `json:"original_price,omitempty"`
	// 售价，单位分
	Price *int32 `json:"price,omitempty"`
	// SKU属性列表
	PropertyList []ItemSkuPropertyInfo `json:"property_list,omitempty"`
}

// NewSkuCreateInfo instantiates a new SkuCreateInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuCreateInfo() *SkuCreateInfo {
	this := SkuCreateInfo{}
	return &this
}

// NewSkuCreateInfoWithDefaults instantiates a new SkuCreateInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuCreateInfoWithDefaults() *SkuCreateInfo {
	this := SkuCreateInfo{}
	return &this
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *SkuCreateInfo) GetInventory() int32 {
	if o == nil || IsNil(o.Inventory) {
		var ret int32
		return ret
	}
	return *o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuCreateInfo) GetInventoryOk() (*int32, bool) {
	if o == nil || IsNil(o.Inventory) {
		return nil, false
	}
	return o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *SkuCreateInfo) HasInventory() bool {
	if o != nil && !IsNil(o.Inventory) {
		return true
	}

	return false
}

// SetInventory gets a reference to the given int32 and assigns it to the Inventory field.
func (o *SkuCreateInfo) SetInventory(v int32) {
	o.Inventory = &v
}

// GetMaterialList returns the MaterialList field value if set, zero value otherwise.
func (o *SkuCreateInfo) GetMaterialList() []MaterialCreateInfo {
	if o == nil || IsNil(o.MaterialList) {
		var ret []MaterialCreateInfo
		return ret
	}
	return o.MaterialList
}

// GetMaterialListOk returns a tuple with the MaterialList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuCreateInfo) GetMaterialListOk() ([]MaterialCreateInfo, bool) {
	if o == nil || IsNil(o.MaterialList) {
		return nil, false
	}
	return o.MaterialList, true
}

// HasMaterialList returns a boolean if a field has been set.
func (o *SkuCreateInfo) HasMaterialList() bool {
	if o != nil && !IsNil(o.MaterialList) {
		return true
	}

	return false
}

// SetMaterialList gets a reference to the given []MaterialCreateInfo and assigns it to the MaterialList field.
func (o *SkuCreateInfo) SetMaterialList(v []MaterialCreateInfo) {
	o.MaterialList = v
}

// GetOriginalPrice returns the OriginalPrice field value if set, zero value otherwise.
func (o *SkuCreateInfo) GetOriginalPrice() int32 {
	if o == nil || IsNil(o.OriginalPrice) {
		var ret int32
		return ret
	}
	return *o.OriginalPrice
}

// GetOriginalPriceOk returns a tuple with the OriginalPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuCreateInfo) GetOriginalPriceOk() (*int32, bool) {
	if o == nil || IsNil(o.OriginalPrice) {
		return nil, false
	}
	return o.OriginalPrice, true
}

// HasOriginalPrice returns a boolean if a field has been set.
func (o *SkuCreateInfo) HasOriginalPrice() bool {
	if o != nil && !IsNil(o.OriginalPrice) {
		return true
	}

	return false
}

// SetOriginalPrice gets a reference to the given int32 and assigns it to the OriginalPrice field.
func (o *SkuCreateInfo) SetOriginalPrice(v int32) {
	o.OriginalPrice = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *SkuCreateInfo) GetPrice() int32 {
	if o == nil || IsNil(o.Price) {
		var ret int32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuCreateInfo) GetPriceOk() (*int32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *SkuCreateInfo) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given int32 and assigns it to the Price field.
func (o *SkuCreateInfo) SetPrice(v int32) {
	o.Price = &v
}

// GetPropertyList returns the PropertyList field value if set, zero value otherwise.
func (o *SkuCreateInfo) GetPropertyList() []ItemSkuPropertyInfo {
	if o == nil || IsNil(o.PropertyList) {
		var ret []ItemSkuPropertyInfo
		return ret
	}
	return o.PropertyList
}

// GetPropertyListOk returns a tuple with the PropertyList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuCreateInfo) GetPropertyListOk() ([]ItemSkuPropertyInfo, bool) {
	if o == nil || IsNil(o.PropertyList) {
		return nil, false
	}
	return o.PropertyList, true
}

// HasPropertyList returns a boolean if a field has been set.
func (o *SkuCreateInfo) HasPropertyList() bool {
	if o != nil && !IsNil(o.PropertyList) {
		return true
	}

	return false
}

// SetPropertyList gets a reference to the given []ItemSkuPropertyInfo and assigns it to the PropertyList field.
func (o *SkuCreateInfo) SetPropertyList(v []ItemSkuPropertyInfo) {
	o.PropertyList = v
}

func (o SkuCreateInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkuCreateInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Inventory) {
		toSerialize["inventory"] = o.Inventory
	}
	if !IsNil(o.MaterialList) {
		toSerialize["material_list"] = o.MaterialList
	}
	if !IsNil(o.OriginalPrice) {
		toSerialize["original_price"] = o.OriginalPrice
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.PropertyList) {
		toSerialize["property_list"] = o.PropertyList
	}
	return toSerialize, nil
}

type NullableSkuCreateInfo struct {
	value *SkuCreateInfo
	isSet bool
}

func (v NullableSkuCreateInfo) Get() *SkuCreateInfo {
	return v.value
}

func (v *NullableSkuCreateInfo) Set(val *SkuCreateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuCreateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuCreateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuCreateInfo(val *SkuCreateInfo) *NullableSkuCreateInfo {
	return &NullableSkuCreateInfo{value: val, isSet: true}
}

func (v NullableSkuCreateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuCreateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
