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

// checks if the AlipayIserviceCcmSwOrderSyncModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayIserviceCcmSwOrderSyncModel{}

// AlipayIserviceCcmSwOrderSyncModel struct for AlipayIserviceCcmSwOrderSyncModel
type AlipayIserviceCcmSwOrderSyncModel struct {
	// 订单金额
	Amount *string `json:"amount,omitempty"`
	// 订单链接
	LinkUrl *string `json:"link_url,omitempty"`
	// 订单物流数量
	LogisticCount *int32 `json:"logistic_count,omitempty"`
	// 物流信息列表
	Logistics []LogisticInfo `json:"logistics,omitempty"`
	// 订单创建时间
	OrderCreateTime *string `json:"order_create_time,omitempty"`
	// 订单id
	OrderId *string `json:"order_id,omitempty"`
	// 普通订单：NORMAL 预售订单：PRE_SALE
	OrderType *string `json:"order_type,omitempty"`
	// 订单商品种类
	SpuCount *int32 `json:"spu_count,omitempty"`
	// 订单商品信息
	Spus []SpuInfo `json:"spus,omitempty"`
	// 订单状态，目前支持以下几种状态 下单未支付 WAIT_PAY 支付未发货 PAIED 已发货 IN_DELIVERY 售后中 IN_AFTER_SALE 订单完成 FINISHED
	Status *string `json:"status,omitempty"`
	// 订单子状态
	SubStatus *string `json:"sub_status,omitempty"`
	// 下单用户id(外部系统ID)
	UserId *string `json:"user_id,omitempty"`
}

// NewAlipayIserviceCcmSwOrderSyncModel instantiates a new AlipayIserviceCcmSwOrderSyncModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayIserviceCcmSwOrderSyncModel() *AlipayIserviceCcmSwOrderSyncModel {
	this := AlipayIserviceCcmSwOrderSyncModel{}
	return &this
}

// NewAlipayIserviceCcmSwOrderSyncModelWithDefaults instantiates a new AlipayIserviceCcmSwOrderSyncModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayIserviceCcmSwOrderSyncModelWithDefaults() *AlipayIserviceCcmSwOrderSyncModel {
	this := AlipayIserviceCcmSwOrderSyncModel{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwOrderSyncModel) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwOrderSyncModel) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwOrderSyncModel) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *AlipayIserviceCcmSwOrderSyncModel) SetAmount(v string) {
	o.Amount = &v
}

// GetLinkUrl returns the LinkUrl field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwOrderSyncModel) GetLinkUrl() string {
	if o == nil || IsNil(o.LinkUrl) {
		var ret string
		return ret
	}
	return *o.LinkUrl
}

// GetLinkUrlOk returns a tuple with the LinkUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwOrderSyncModel) GetLinkUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LinkUrl) {
		return nil, false
	}
	return o.LinkUrl, true
}

// HasLinkUrl returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwOrderSyncModel) HasLinkUrl() bool {
	if o != nil && !IsNil(o.LinkUrl) {
		return true
	}

	return false
}

// SetLinkUrl gets a reference to the given string and assigns it to the LinkUrl field.
func (o *AlipayIserviceCcmSwOrderSyncModel) SetLinkUrl(v string) {
	o.LinkUrl = &v
}

// GetLogisticCount returns the LogisticCount field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwOrderSyncModel) GetLogisticCount() int32 {
	if o == nil || IsNil(o.LogisticCount) {
		var ret int32
		return ret
	}
	return *o.LogisticCount
}

// GetLogisticCountOk returns a tuple with the LogisticCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwOrderSyncModel) GetLogisticCountOk() (*int32, bool) {
	if o == nil || IsNil(o.LogisticCount) {
		return nil, false
	}
	return o.LogisticCount, true
}

// HasLogisticCount returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwOrderSyncModel) HasLogisticCount() bool {
	if o != nil && !IsNil(o.LogisticCount) {
		return true
	}

	return false
}

// SetLogisticCount gets a reference to the given int32 and assigns it to the LogisticCount field.
func (o *AlipayIserviceCcmSwOrderSyncModel) SetLogisticCount(v int32) {
	o.LogisticCount = &v
}

// GetLogistics returns the Logistics field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwOrderSyncModel) GetLogistics() []LogisticInfo {
	if o == nil || IsNil(o.Logistics) {
		var ret []LogisticInfo
		return ret
	}
	return o.Logistics
}

// GetLogisticsOk returns a tuple with the Logistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwOrderSyncModel) GetLogisticsOk() ([]LogisticInfo, bool) {
	if o == nil || IsNil(o.Logistics) {
		return nil, false
	}
	return o.Logistics, true
}

// HasLogistics returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwOrderSyncModel) HasLogistics() bool {
	if o != nil && !IsNil(o.Logistics) {
		return true
	}

	return false
}

// SetLogistics gets a reference to the given []LogisticInfo and assigns it to the Logistics field.
func (o *AlipayIserviceCcmSwOrderSyncModel) SetLogistics(v []LogisticInfo) {
	o.Logistics = v
}

// GetOrderCreateTime returns the OrderCreateTime field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwOrderSyncModel) GetOrderCreateTime() string {
	if o == nil || IsNil(o.OrderCreateTime) {
		var ret string
		return ret
	}
	return *o.OrderCreateTime
}

// GetOrderCreateTimeOk returns a tuple with the OrderCreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwOrderSyncModel) GetOrderCreateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.OrderCreateTime) {
		return nil, false
	}
	return o.OrderCreateTime, true
}

// HasOrderCreateTime returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwOrderSyncModel) HasOrderCreateTime() bool {
	if o != nil && !IsNil(o.OrderCreateTime) {
		return true
	}

	return false
}

// SetOrderCreateTime gets a reference to the given string and assigns it to the OrderCreateTime field.
func (o *AlipayIserviceCcmSwOrderSyncModel) SetOrderCreateTime(v string) {
	o.OrderCreateTime = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwOrderSyncModel) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwOrderSyncModel) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwOrderSyncModel) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *AlipayIserviceCcmSwOrderSyncModel) SetOrderId(v string) {
	o.OrderId = &v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwOrderSyncModel) GetOrderType() string {
	if o == nil || IsNil(o.OrderType) {
		var ret string
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwOrderSyncModel) GetOrderTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwOrderSyncModel) HasOrderType() bool {
	if o != nil && !IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given string and assigns it to the OrderType field.
func (o *AlipayIserviceCcmSwOrderSyncModel) SetOrderType(v string) {
	o.OrderType = &v
}

// GetSpuCount returns the SpuCount field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwOrderSyncModel) GetSpuCount() int32 {
	if o == nil || IsNil(o.SpuCount) {
		var ret int32
		return ret
	}
	return *o.SpuCount
}

// GetSpuCountOk returns a tuple with the SpuCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwOrderSyncModel) GetSpuCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SpuCount) {
		return nil, false
	}
	return o.SpuCount, true
}

// HasSpuCount returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwOrderSyncModel) HasSpuCount() bool {
	if o != nil && !IsNil(o.SpuCount) {
		return true
	}

	return false
}

// SetSpuCount gets a reference to the given int32 and assigns it to the SpuCount field.
func (o *AlipayIserviceCcmSwOrderSyncModel) SetSpuCount(v int32) {
	o.SpuCount = &v
}

// GetSpus returns the Spus field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwOrderSyncModel) GetSpus() []SpuInfo {
	if o == nil || IsNil(o.Spus) {
		var ret []SpuInfo
		return ret
	}
	return o.Spus
}

// GetSpusOk returns a tuple with the Spus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwOrderSyncModel) GetSpusOk() ([]SpuInfo, bool) {
	if o == nil || IsNil(o.Spus) {
		return nil, false
	}
	return o.Spus, true
}

// HasSpus returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwOrderSyncModel) HasSpus() bool {
	if o != nil && !IsNil(o.Spus) {
		return true
	}

	return false
}

// SetSpus gets a reference to the given []SpuInfo and assigns it to the Spus field.
func (o *AlipayIserviceCcmSwOrderSyncModel) SetSpus(v []SpuInfo) {
	o.Spus = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwOrderSyncModel) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwOrderSyncModel) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwOrderSyncModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlipayIserviceCcmSwOrderSyncModel) SetStatus(v string) {
	o.Status = &v
}

// GetSubStatus returns the SubStatus field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwOrderSyncModel) GetSubStatus() string {
	if o == nil || IsNil(o.SubStatus) {
		var ret string
		return ret
	}
	return *o.SubStatus
}

// GetSubStatusOk returns a tuple with the SubStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwOrderSyncModel) GetSubStatusOk() (*string, bool) {
	if o == nil || IsNil(o.SubStatus) {
		return nil, false
	}
	return o.SubStatus, true
}

// HasSubStatus returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwOrderSyncModel) HasSubStatus() bool {
	if o != nil && !IsNil(o.SubStatus) {
		return true
	}

	return false
}

// SetSubStatus gets a reference to the given string and assigns it to the SubStatus field.
func (o *AlipayIserviceCcmSwOrderSyncModel) SetSubStatus(v string) {
	o.SubStatus = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwOrderSyncModel) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwOrderSyncModel) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwOrderSyncModel) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *AlipayIserviceCcmSwOrderSyncModel) SetUserId(v string) {
	o.UserId = &v
}

func (o AlipayIserviceCcmSwOrderSyncModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayIserviceCcmSwOrderSyncModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.LinkUrl) {
		toSerialize["link_url"] = o.LinkUrl
	}
	if !IsNil(o.LogisticCount) {
		toSerialize["logistic_count"] = o.LogisticCount
	}
	if !IsNil(o.Logistics) {
		toSerialize["logistics"] = o.Logistics
	}
	if !IsNil(o.OrderCreateTime) {
		toSerialize["order_create_time"] = o.OrderCreateTime
	}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.OrderType) {
		toSerialize["order_type"] = o.OrderType
	}
	if !IsNil(o.SpuCount) {
		toSerialize["spu_count"] = o.SpuCount
	}
	if !IsNil(o.Spus) {
		toSerialize["spus"] = o.Spus
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SubStatus) {
		toSerialize["sub_status"] = o.SubStatus
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableAlipayIserviceCcmSwOrderSyncModel struct {
	value *AlipayIserviceCcmSwOrderSyncModel
	isSet bool
}

func (v NullableAlipayIserviceCcmSwOrderSyncModel) Get() *AlipayIserviceCcmSwOrderSyncModel {
	return v.value
}

func (v *NullableAlipayIserviceCcmSwOrderSyncModel) Set(val *AlipayIserviceCcmSwOrderSyncModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayIserviceCcmSwOrderSyncModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayIserviceCcmSwOrderSyncModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayIserviceCcmSwOrderSyncModel(val *AlipayIserviceCcmSwOrderSyncModel) *NullableAlipayIserviceCcmSwOrderSyncModel {
	return &NullableAlipayIserviceCcmSwOrderSyncModel{value: val, isSet: true}
}

func (v NullableAlipayIserviceCcmSwOrderSyncModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayIserviceCcmSwOrderSyncModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
