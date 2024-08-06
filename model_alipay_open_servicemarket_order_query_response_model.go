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

// checks if the AlipayOpenServicemarketOrderQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenServicemarketOrderQueryResponseModel{}

// AlipayOpenServicemarketOrderQueryResponseModel struct for AlipayOpenServicemarketOrderQueryResponseModel
type AlipayOpenServicemarketOrderQueryResponseModel struct {
	// true：开启；false：关闭；需要同步校验commodity_id，如果没有查询到订购信息的话，忽略该字段返回值
	AutoUpgrade *bool `json:"auto_upgrade,omitempty"`
	// 订购服务商品ID
	CommodityId *string `json:"commodity_id,omitempty"`
	// 当前查询页（本接口支持最多查询100条记录）
	CurrentPage *int32 `json:"current_page,omitempty"`
	// 订单明细列表
	OrderItems []OrderItem `json:"order_items,omitempty"`
	// 用于区分同一个服务的不同版本
	Specifications *string `json:"specifications,omitempty"`
	// MERCHANT_ORDED（待服务商接单）
	Status *string `json:"status,omitempty"`
	// 总记录数
	TotalSize *int32 `json:"total_size,omitempty"`
}

// NewAlipayOpenServicemarketOrderQueryResponseModel instantiates a new AlipayOpenServicemarketOrderQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenServicemarketOrderQueryResponseModel() *AlipayOpenServicemarketOrderQueryResponseModel {
	this := AlipayOpenServicemarketOrderQueryResponseModel{}
	return &this
}

// NewAlipayOpenServicemarketOrderQueryResponseModelWithDefaults instantiates a new AlipayOpenServicemarketOrderQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenServicemarketOrderQueryResponseModelWithDefaults() *AlipayOpenServicemarketOrderQueryResponseModel {
	this := AlipayOpenServicemarketOrderQueryResponseModel{}
	return &this
}

// GetAutoUpgrade returns the AutoUpgrade field value if set, zero value otherwise.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) GetAutoUpgrade() bool {
	if o == nil || IsNil(o.AutoUpgrade) {
		var ret bool
		return ret
	}
	return *o.AutoUpgrade
}

// GetAutoUpgradeOk returns a tuple with the AutoUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) GetAutoUpgradeOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoUpgrade) {
		return nil, false
	}
	return o.AutoUpgrade, true
}

// HasAutoUpgrade returns a boolean if a field has been set.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) HasAutoUpgrade() bool {
	if o != nil && !IsNil(o.AutoUpgrade) {
		return true
	}

	return false
}

// SetAutoUpgrade gets a reference to the given bool and assigns it to the AutoUpgrade field.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) SetAutoUpgrade(v bool) {
	o.AutoUpgrade = &v
}

// GetCommodityId returns the CommodityId field value if set, zero value otherwise.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) GetCommodityId() string {
	if o == nil || IsNil(o.CommodityId) {
		var ret string
		return ret
	}
	return *o.CommodityId
}

// GetCommodityIdOk returns a tuple with the CommodityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) GetCommodityIdOk() (*string, bool) {
	if o == nil || IsNil(o.CommodityId) {
		return nil, false
	}
	return o.CommodityId, true
}

// HasCommodityId returns a boolean if a field has been set.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) HasCommodityId() bool {
	if o != nil && !IsNil(o.CommodityId) {
		return true
	}

	return false
}

// SetCommodityId gets a reference to the given string and assigns it to the CommodityId field.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) SetCommodityId(v string) {
	o.CommodityId = &v
}

// GetCurrentPage returns the CurrentPage field value if set, zero value otherwise.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) GetCurrentPage() int32 {
	if o == nil || IsNil(o.CurrentPage) {
		var ret int32
		return ret
	}
	return *o.CurrentPage
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) GetCurrentPageOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrentPage) {
		return nil, false
	}
	return o.CurrentPage, true
}

// HasCurrentPage returns a boolean if a field has been set.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) HasCurrentPage() bool {
	if o != nil && !IsNil(o.CurrentPage) {
		return true
	}

	return false
}

// SetCurrentPage gets a reference to the given int32 and assigns it to the CurrentPage field.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) SetCurrentPage(v int32) {
	o.CurrentPage = &v
}

// GetOrderItems returns the OrderItems field value if set, zero value otherwise.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) GetOrderItems() []OrderItem {
	if o == nil || IsNil(o.OrderItems) {
		var ret []OrderItem
		return ret
	}
	return o.OrderItems
}

// GetOrderItemsOk returns a tuple with the OrderItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) GetOrderItemsOk() ([]OrderItem, bool) {
	if o == nil || IsNil(o.OrderItems) {
		return nil, false
	}
	return o.OrderItems, true
}

// HasOrderItems returns a boolean if a field has been set.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) HasOrderItems() bool {
	if o != nil && !IsNil(o.OrderItems) {
		return true
	}

	return false
}

// SetOrderItems gets a reference to the given []OrderItem and assigns it to the OrderItems field.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) SetOrderItems(v []OrderItem) {
	o.OrderItems = v
}

// GetSpecifications returns the Specifications field value if set, zero value otherwise.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) GetSpecifications() string {
	if o == nil || IsNil(o.Specifications) {
		var ret string
		return ret
	}
	return *o.Specifications
}

// GetSpecificationsOk returns a tuple with the Specifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) GetSpecificationsOk() (*string, bool) {
	if o == nil || IsNil(o.Specifications) {
		return nil, false
	}
	return o.Specifications, true
}

// HasSpecifications returns a boolean if a field has been set.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) HasSpecifications() bool {
	if o != nil && !IsNil(o.Specifications) {
		return true
	}

	return false
}

// SetSpecifications gets a reference to the given string and assigns it to the Specifications field.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) SetSpecifications(v string) {
	o.Specifications = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) SetStatus(v string) {
	o.Status = &v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) GetTotalSize() int32 {
	if o == nil || IsNil(o.TotalSize) {
		var ret int32
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) GetTotalSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalSize) {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) HasTotalSize() bool {
	if o != nil && !IsNil(o.TotalSize) {
		return true
	}

	return false
}

// SetTotalSize gets a reference to the given int32 and assigns it to the TotalSize field.
func (o *AlipayOpenServicemarketOrderQueryResponseModel) SetTotalSize(v int32) {
	o.TotalSize = &v
}

func (o AlipayOpenServicemarketOrderQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenServicemarketOrderQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoUpgrade) {
		toSerialize["auto_upgrade"] = o.AutoUpgrade
	}
	if !IsNil(o.CommodityId) {
		toSerialize["commodity_id"] = o.CommodityId
	}
	if !IsNil(o.CurrentPage) {
		toSerialize["current_page"] = o.CurrentPage
	}
	if !IsNil(o.OrderItems) {
		toSerialize["order_items"] = o.OrderItems
	}
	if !IsNil(o.Specifications) {
		toSerialize["specifications"] = o.Specifications
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TotalSize) {
		toSerialize["total_size"] = o.TotalSize
	}
	return toSerialize, nil
}

type NullableAlipayOpenServicemarketOrderQueryResponseModel struct {
	value *AlipayOpenServicemarketOrderQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenServicemarketOrderQueryResponseModel) Get() *AlipayOpenServicemarketOrderQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenServicemarketOrderQueryResponseModel) Set(val *AlipayOpenServicemarketOrderQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenServicemarketOrderQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenServicemarketOrderQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenServicemarketOrderQueryResponseModel(val *AlipayOpenServicemarketOrderQueryResponseModel) *NullableAlipayOpenServicemarketOrderQueryResponseModel {
	return &NullableAlipayOpenServicemarketOrderQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenServicemarketOrderQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenServicemarketOrderQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

