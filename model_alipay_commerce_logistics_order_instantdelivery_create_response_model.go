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

// checks if the AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel{}

// AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel struct for AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel
type AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel struct {
	// 优惠券费用
	CouponFee *string `json:"coupon_fee,omitempty"`
	// 运费
	DeliverFee *string `json:"deliver_fee,omitempty"`
	// 预计骑手还剩多久接单，单位：秒
	DispatchDuration *int32 `json:"dispatch_duration,omitempty"`
	// 配送距离
	Distance *int32 `json:"distance,omitempty"`
	// 实际运费
	Fee *string `json:"fee,omitempty"`
	// 收货码，骑手必须输入收货码才能完成订单妥投
	FinishCode *string `json:"finish_code,omitempty"`
	// 保价费用
	InsuranceFee *string `json:"insurance_fee,omitempty"`
	// 支付宝订单流水号
	OrderNo *string `json:"order_no,omitempty"`
	// 支付金额, 实际扣减的费用以此字段为准
	PayAmount *string `json:"pay_amount,omitempty"`
	// 取货码, 骑手必须输入取货码才能从商家取货
	PickupCode *string `json:"pickup_code,omitempty"`
	// 即时配送运单状态
	Status *string `json:"status,omitempty"`
	// 即时配送运单编号
	WaybillNo *string `json:"waybill_no,omitempty"`
}

// NewAlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel instantiates a new AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel() *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel {
	this := AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel{}
	return &this
}

// NewAlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModelWithDefaults instantiates a new AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModelWithDefaults() *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel {
	this := AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel{}
	return &this
}

// GetCouponFee returns the CouponFee field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) GetCouponFee() string {
	if o == nil || IsNil(o.CouponFee) {
		var ret string
		return ret
	}
	return *o.CouponFee
}

// GetCouponFeeOk returns a tuple with the CouponFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) GetCouponFeeOk() (*string, bool) {
	if o == nil || IsNil(o.CouponFee) {
		return nil, false
	}
	return o.CouponFee, true
}

// HasCouponFee returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) HasCouponFee() bool {
	if o != nil && !IsNil(o.CouponFee) {
		return true
	}

	return false
}

// SetCouponFee gets a reference to the given string and assigns it to the CouponFee field.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) SetCouponFee(v string) {
	o.CouponFee = &v
}

// GetDeliverFee returns the DeliverFee field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) GetDeliverFee() string {
	if o == nil || IsNil(o.DeliverFee) {
		var ret string
		return ret
	}
	return *o.DeliverFee
}

// GetDeliverFeeOk returns a tuple with the DeliverFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) GetDeliverFeeOk() (*string, bool) {
	if o == nil || IsNil(o.DeliverFee) {
		return nil, false
	}
	return o.DeliverFee, true
}

// HasDeliverFee returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) HasDeliverFee() bool {
	if o != nil && !IsNil(o.DeliverFee) {
		return true
	}

	return false
}

// SetDeliverFee gets a reference to the given string and assigns it to the DeliverFee field.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) SetDeliverFee(v string) {
	o.DeliverFee = &v
}

// GetDispatchDuration returns the DispatchDuration field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) GetDispatchDuration() int32 {
	if o == nil || IsNil(o.DispatchDuration) {
		var ret int32
		return ret
	}
	return *o.DispatchDuration
}

// GetDispatchDurationOk returns a tuple with the DispatchDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) GetDispatchDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.DispatchDuration) {
		return nil, false
	}
	return o.DispatchDuration, true
}

// HasDispatchDuration returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) HasDispatchDuration() bool {
	if o != nil && !IsNil(o.DispatchDuration) {
		return true
	}

	return false
}

// SetDispatchDuration gets a reference to the given int32 and assigns it to the DispatchDuration field.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) SetDispatchDuration(v int32) {
	o.DispatchDuration = &v
}

// GetDistance returns the Distance field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) GetDistance() int32 {
	if o == nil || IsNil(o.Distance) {
		var ret int32
		return ret
	}
	return *o.Distance
}

// GetDistanceOk returns a tuple with the Distance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) GetDistanceOk() (*int32, bool) {
	if o == nil || IsNil(o.Distance) {
		return nil, false
	}
	return o.Distance, true
}

// HasDistance returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) HasDistance() bool {
	if o != nil && !IsNil(o.Distance) {
		return true
	}

	return false
}

// SetDistance gets a reference to the given int32 and assigns it to the Distance field.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) SetDistance(v int32) {
	o.Distance = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) GetFee() string {
	if o == nil || IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) GetFeeOk() (*string, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) SetFee(v string) {
	o.Fee = &v
}

// GetFinishCode returns the FinishCode field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) GetFinishCode() string {
	if o == nil || IsNil(o.FinishCode) {
		var ret string
		return ret
	}
	return *o.FinishCode
}

// GetFinishCodeOk returns a tuple with the FinishCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) GetFinishCodeOk() (*string, bool) {
	if o == nil || IsNil(o.FinishCode) {
		return nil, false
	}
	return o.FinishCode, true
}

// HasFinishCode returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) HasFinishCode() bool {
	if o != nil && !IsNil(o.FinishCode) {
		return true
	}

	return false
}

// SetFinishCode gets a reference to the given string and assigns it to the FinishCode field.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) SetFinishCode(v string) {
	o.FinishCode = &v
}

// GetInsuranceFee returns the InsuranceFee field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) GetInsuranceFee() string {
	if o == nil || IsNil(o.InsuranceFee) {
		var ret string
		return ret
	}
	return *o.InsuranceFee
}

// GetInsuranceFeeOk returns a tuple with the InsuranceFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) GetInsuranceFeeOk() (*string, bool) {
	if o == nil || IsNil(o.InsuranceFee) {
		return nil, false
	}
	return o.InsuranceFee, true
}

// HasInsuranceFee returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) HasInsuranceFee() bool {
	if o != nil && !IsNil(o.InsuranceFee) {
		return true
	}

	return false
}

// SetInsuranceFee gets a reference to the given string and assigns it to the InsuranceFee field.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) SetInsuranceFee(v string) {
	o.InsuranceFee = &v
}

// GetOrderNo returns the OrderNo field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) GetOrderNo() string {
	if o == nil || IsNil(o.OrderNo) {
		var ret string
		return ret
	}
	return *o.OrderNo
}

// GetOrderNoOk returns a tuple with the OrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) GetOrderNoOk() (*string, bool) {
	if o == nil || IsNil(o.OrderNo) {
		return nil, false
	}
	return o.OrderNo, true
}

// HasOrderNo returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) HasOrderNo() bool {
	if o != nil && !IsNil(o.OrderNo) {
		return true
	}

	return false
}

// SetOrderNo gets a reference to the given string and assigns it to the OrderNo field.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) SetOrderNo(v string) {
	o.OrderNo = &v
}

// GetPayAmount returns the PayAmount field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) GetPayAmount() string {
	if o == nil || IsNil(o.PayAmount) {
		var ret string
		return ret
	}
	return *o.PayAmount
}

// GetPayAmountOk returns a tuple with the PayAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) GetPayAmountOk() (*string, bool) {
	if o == nil || IsNil(o.PayAmount) {
		return nil, false
	}
	return o.PayAmount, true
}

// HasPayAmount returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) HasPayAmount() bool {
	if o != nil && !IsNil(o.PayAmount) {
		return true
	}

	return false
}

// SetPayAmount gets a reference to the given string and assigns it to the PayAmount field.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) SetPayAmount(v string) {
	o.PayAmount = &v
}

// GetPickupCode returns the PickupCode field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) GetPickupCode() string {
	if o == nil || IsNil(o.PickupCode) {
		var ret string
		return ret
	}
	return *o.PickupCode
}

// GetPickupCodeOk returns a tuple with the PickupCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) GetPickupCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PickupCode) {
		return nil, false
	}
	return o.PickupCode, true
}

// HasPickupCode returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) HasPickupCode() bool {
	if o != nil && !IsNil(o.PickupCode) {
		return true
	}

	return false
}

// SetPickupCode gets a reference to the given string and assigns it to the PickupCode field.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) SetPickupCode(v string) {
	o.PickupCode = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) SetStatus(v string) {
	o.Status = &v
}

// GetWaybillNo returns the WaybillNo field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) GetWaybillNo() string {
	if o == nil || IsNil(o.WaybillNo) {
		var ret string
		return ret
	}
	return *o.WaybillNo
}

// GetWaybillNoOk returns a tuple with the WaybillNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) GetWaybillNoOk() (*string, bool) {
	if o == nil || IsNil(o.WaybillNo) {
		return nil, false
	}
	return o.WaybillNo, true
}

// HasWaybillNo returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) HasWaybillNo() bool {
	if o != nil && !IsNil(o.WaybillNo) {
		return true
	}

	return false
}

// SetWaybillNo gets a reference to the given string and assigns it to the WaybillNo field.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) SetWaybillNo(v string) {
	o.WaybillNo = &v
}

func (o AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CouponFee) {
		toSerialize["coupon_fee"] = o.CouponFee
	}
	if !IsNil(o.DeliverFee) {
		toSerialize["deliver_fee"] = o.DeliverFee
	}
	if !IsNil(o.DispatchDuration) {
		toSerialize["dispatch_duration"] = o.DispatchDuration
	}
	if !IsNil(o.Distance) {
		toSerialize["distance"] = o.Distance
	}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !IsNil(o.FinishCode) {
		toSerialize["finish_code"] = o.FinishCode
	}
	if !IsNil(o.InsuranceFee) {
		toSerialize["insurance_fee"] = o.InsuranceFee
	}
	if !IsNil(o.OrderNo) {
		toSerialize["order_no"] = o.OrderNo
	}
	if !IsNil(o.PayAmount) {
		toSerialize["pay_amount"] = o.PayAmount
	}
	if !IsNil(o.PickupCode) {
		toSerialize["pickup_code"] = o.PickupCode
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.WaybillNo) {
		toSerialize["waybill_no"] = o.WaybillNo
	}
	return toSerialize, nil
}

type NullableAlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel struct {
	value *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel
	isSet bool
}

func (v NullableAlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) Get() *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel {
	return v.value
}

func (v *NullableAlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) Set(val *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel(val *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) *NullableAlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel {
	return &NullableAlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
